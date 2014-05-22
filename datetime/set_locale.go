/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package main

import (
	"dlib/utils"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	DMRC_FILE      = ".dmrc"
	DMRC_KEY_GROUP = "Desktop"
	PAM_ENV_FILE   = ".pam_environment"

	DEFAULT_LOCALE_FILE = "/etc/default/locale"
)

func (obj *Manager) listenLocaleChange() {
	setDate.ConnectGenLocaleStatus(func(ok bool, locale string) {
		if ok && changeLocaleFlag {
			//setLocaleDmrc(locale)
			setLocalePamEnv(locale)
			changeLocaleFlag = false
			obj.setPropName("CurrentLocale")
		}
		obj.LocaleStatus(ok, locale)
	})
}

func setLocaleDmrc(locale string) {
	homeDir, ok := getHomeDir()
	if !ok {
		return
	}
	utilsObj := utils.NewUtils()
	filePath := path.Join(homeDir, DMRC_FILE)
	utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
		"LANG", locale)
	utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
		"LANGUAGE", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_CTYPE", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_NUMERIC", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_TIME", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_COLLATE", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_MONETARY", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_MESSAGES", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_PAPER", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_NAME", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_ADDRESS", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_TELEPHONE", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_MEASUREMENT", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_IDENTIFICATION", locale)
	//utilsObj.WriteKeyToKeyFile(filePath, DMRC_KEY_GROUP,
	//"LC_ALL", "")
}

func genPamContents(locale string) string {
	contents := ""
	tmp := "LANG=" + locale + "\n"
	contents += tmp
	tmp = "LANGUAGE=" + locale + "\n"
	contents += tmp
	//tmp = "LC_CTYPE=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_NUMERIC=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_TIME=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_COLLATE=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_MONETARY=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_MESSAGES=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_PAPER=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_NAME=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_ADDRESS=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_TELEPHONE=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_MEASUREMENT=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_IDENTIFICATION=" + locale + "\n"
	//contents += tmp
	//tmp = "LC_ALL=\n"
	//contents += tmp

	return contents
}

func setLocalePamEnv(locale string) {
	homeDir, ok := getHomeDir()
	if !ok {
		return
	}
	filePath := path.Join(homeDir, PAM_ENV_FILE)

	fp, err := os.Create(filePath + "~")
	if err != nil {
		logger.Errorf("Create '%s' failed: %v", filePath+"~", err)
		return
	}
	defer fp.Close()

	if _, err = fp.WriteString(genPamContents(locale)); err != nil {
		logger.Errorf("Write '%s' failed: %v", filePath+"~", err)
		return
	}
	fp.Sync()
	os.Rename(filePath+"~", filePath)
}

func getDefaultLocale() (string, bool) {
	if !objUtils.IsFileExist(DEFAULT_LOCALE_FILE) {
		logger.Errorf("'%s' not exist", DEFAULT_LOCALE_FILE)
		return "", false
	}

	contents, err := ioutil.ReadFile(DEFAULT_LOCALE_FILE)
	if err != nil {
		logger.Errorf("ReadFile '%s' failed: %v", DEFAULT_LOCALE_FILE, err)
		return "", false
	}

	retStr := ""
	retOk := false
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		strs := strings.Split(line, "=")
		if len(strs) != 2 {
			continue
		}

		if strs[0] != "LANG" {
			continue
		}

		retStr = strings.Trim(strs[1], "\"")
		retOk = true
		break
	}

	return retStr, retOk
}

func getUserLocale() (string, bool) {
	homeDir, _ := objUtils.GetHomeDir()
	filePath := path.Join(homeDir, PAM_ENV_FILE)
	if !objUtils.IsFileExist(filePath) {
		logger.Warningf("'%s' not exist", filePath)
		return "", false
	}

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Warningf("ReadFile '%s' failed: %v", filePath, err)
		return "", false
	}

	retStr := ""
	retOk := false
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		strs := strings.Split(line, "=")
		if len(strs) != 2 {
			continue
		}

		if strs[0] != "LANG" {
			continue
		}

		retStr = strs[1]
		retOk = true
		break
	}

	return retStr, retOk
}