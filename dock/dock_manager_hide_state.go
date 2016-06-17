/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package dock

import (
	"errors"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/icccm"
	"github.com/BurntSushi/xgbutil/xrect"
	"github.com/BurntSushi/xgbutil/xwindow"
	"pkg.deepin.io/lib/dbus"
	"time"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func hasIntersection(rectA, rectB xrect.Rect) bool {
	if rectA == nil || rectB == nil {
		logger.Warning("hasIntersection rectA or rectB is nil")
		return false
	}
	x, y, w, h := rectA.Pieces()
	x1, y1, w1, h1 := rectB.Pieces()
	ax := max(x, x1)
	ay := max(y, y1)
	bx := min(x+w, x1+w1)
	by := min(y+h, y1+h1)
	return ax <= bx && ay <= by
}

func (m *DockManager) isWindowDockOverlap(win xproto.Window) (bool, error) {
	// overlap condition:  window showing and  on current workspace,
	// window dock rect has intersection
	window := xwindow.New(XU, win)
	if isHiddenPre(win) || (!onCurrentWorkspacePre(win)) {
		logger.Debugf("window %v is hidden or not on current workspace", win)
		return false, nil
	}

	winRect, err := window.DecorGeometry()
	if err != nil {
		logger.Warning("Get target window geometry failed", err)
		return false, err
	}

	dockWindow := xwindow.New(XU, m.frontendWindow)
	dockRect, err := dockWindow.DecorGeometry()
	if err != nil {
		logger.Warning("Get dock window geometry failed:", err)
		return false, err
	}

	logger.Debug("window rect:", winRect)
	logger.Debug("dock rect:", dockRect)
	result := hasIntersection(winRect, dockRect)
	logger.Debug("window dock overlap:", result)
	return result, nil
}

const (
	DDELauncher = "dde-launcher"
)

func (m *DockManager) isDeepinLauncherShown() bool {
	winClass, err := icccm.WmClassGet(XU, m.activeWindow)
	if err != nil {
		logger.Debug(err)
		return false
	}
	return winClass.Instance == DDELauncher
}

func (m *DockManager) shouldHideOnSmartHideMode() (bool, error) {
	if m.activeWindow == 0 {
		logger.Debug("shouldHideOnSmartHideMode: activeWindow is 0")
		return false, errors.New("activeWindow is 0")
	}
	if m.isDeepinLauncherShown() {
		logger.Debug("launcher is shown")
		return false, nil
	}
	return m.isWindowDockOverlap(m.activeWindow)
}

func (m *DockManager) smartHideModeTimerExpired() {
	logger.Debug("smartHideModeTimer expired!")
	shouldHide, err := m.shouldHideOnSmartHideMode()
	if err != nil {
		m.setPropHideState(HideStateUnknown)
		return
	}

	if shouldHide {
		m.setPropHideState(HideStateHide)
	} else {
		m.setPropHideState(HideStateShow)
	}
}

func (m *DockManager) resetSmartHideModeTimer(delay time.Duration) {
	m.smartHideModeMutex.Lock()
	defer m.smartHideModeMutex.Unlock()

	m.smartHideModeTimer.Reset(delay)
	logger.Debug("reset smart hide mode timer ", delay)
}

func (m *DockManager) cancelSmartHideModeTimer() {
	m.smartHideModeMutex.Lock()
	defer m.smartHideModeMutex.Unlock()

	m.smartHideModeTimer.Stop()
	logger.Debug("cancel smart hide mode timer ")
}

func (m *DockManager) smartHideModeDelayHandle() {
	shouldHide, err := m.shouldHideOnSmartHideMode()
	if err != nil {
		m.cancelSmartHideModeTimer()
		m.setPropHideState(HideStateUnknown)
		return
	}

	switch m.HideState {
	case HideStateShow:
		if shouldHide {
			logger.Debug("smartHideModeDelayHandle: show -> hide")
			m.resetSmartHideModeTimer(time.Millisecond * 500)
		} else {
			logger.Debug("smartHideModeDelayHandle: show -> show")
			m.cancelSmartHideModeTimer()
		}
	case HideStateHide:
		if shouldHide {
			logger.Debug("smartHideModeDelayHandle: hide -> hide")
			m.cancelSmartHideModeTimer()
		} else {
			logger.Debug("smartHideModeDelayHandle: hide -> show")
			m.resetSmartHideModeTimer(time.Millisecond * 500)
		}
	default:
		m.resetSmartHideModeTimer(0)
	}
}

func (m *DockManager) updateHideState(delay bool) {
	if m.isDeepinLauncherShown() {
		logger.Debug("updateHideState: launcher is shown, show dock")
		m.setPropHideState(HideStateShow)
		return
	}

	hideMode := HideModeType(m.HideMode.Get())
	logger.Debug("updateHideState: mode is", hideMode)
	switch hideMode {
	case HideModeKeepShowing:
		m.setPropHideState(HideStateShow)

	case HideModeKeepHidden:
		m.setPropHideState(HideStateHide)

	case HideModeSmartHide:
		if delay {
			m.smartHideModeDelayHandle()
		} else {
			m.smartHideModeTimer.Reset(0)
		}
	}
}

func (m *DockManager) updateHideStateWithDelay() {
	m.updateHideState(true)
}

func (m *DockManager) updateHideStateWithoutDelay() {
	m.updateHideState(false)
}

func (m *DockManager) setPropHideState(hideState HideStateType) {
	logger.Debug("setPropHideState", hideState)
	if m.HideState != hideState {
		logger.Debugf("HideState %v => %v", m.HideState, hideState)
		m.HideState = hideState
		dbus.NotifyChange(m, "HideState")
	}
}