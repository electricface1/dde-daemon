// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingCdmaKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_CDMA_NUMBER:
		t = ktypeString
	case NM_SETTING_CDMA_USERNAME:
		t = ktypeString
	case NM_SETTING_CDMA_PASSWORD:
		t = ktypeString
	case NM_SETTING_CDMA_PASSWORD_FLAGS:
		t = ktypeUint32
	}
	return
}

// Check is key in current setting field
func isKeyInSettingCdma(key string) bool {
	switch key {
	case NM_SETTING_CDMA_NUMBER:
		return true
	case NM_SETTING_CDMA_USERNAME:
		return true
	case NM_SETTING_CDMA_PASSWORD:
		return true
	case NM_SETTING_CDMA_PASSWORD_FLAGS:
		return true
	}
	return false
}

// Get key's default value
func getSettingCdmaDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_CDMA_NUMBER:
		value = "#777"
	case NM_SETTING_CDMA_USERNAME:
		value = ""
	case NM_SETTING_CDMA_PASSWORD:
		value = ""
	case NM_SETTING_CDMA_PASSWORD_FLAGS:
		value = nil
	}
	return
}

// Get JSON value generally
func generalGetSettingCdmaKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingCdmaKeyJSON: invalide key", key)
	case NM_SETTING_CDMA_NUMBER:
		value = getSettingCdmaNumberJSON(data)
	case NM_SETTING_CDMA_USERNAME:
		value = getSettingCdmaUsernameJSON(data)
	case NM_SETTING_CDMA_PASSWORD:
		value = getSettingCdmaPasswordJSON(data)
	case NM_SETTING_CDMA_PASSWORD_FLAGS:
		value = getSettingCdmaPasswordFlagsJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingCdmaKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingCdmaKeyJSON: invalide key", key)
	case NM_SETTING_CDMA_NUMBER:
		err = setSettingCdmaNumberJSON(data, valueJSON)
	case NM_SETTING_CDMA_USERNAME:
		err = setSettingCdmaUsernameJSON(data, valueJSON)
	case NM_SETTING_CDMA_PASSWORD:
		err = setSettingCdmaPasswordJSON(data, valueJSON)
	case NM_SETTING_CDMA_PASSWORD_FLAGS:
		err = setSettingCdmaPasswordFlagsJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingCdmaNumberExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_NUMBER)
}
func isSettingCdmaUsernameExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_USERNAME)
}
func isSettingCdmaPasswordExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD)
}
func isSettingCdmaPasswordFlagsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD_FLAGS)
}

// Ensure field and key exists and not empty
func ensureFieldSettingCdmaExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_CDMA_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_CDMA_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_CDMA_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_CDMA_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_CDMA_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_CDMA_SETTING_NAME))
	}
}
func ensureSettingCdmaNumberNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingCdmaNumberExists(data) {
		rememberError(errs, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_NUMBER, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingCdmaNumber(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_NUMBER, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingCdmaUsernameNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingCdmaUsernameExists(data) {
		rememberError(errs, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_USERNAME, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingCdmaUsername(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_USERNAME, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingCdmaPasswordNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingCdmaPasswordExists(data) {
		rememberError(errs, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingCdmaPassword(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingCdmaPasswordFlagsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingCdmaPasswordFlagsExists(data) {
		rememberError(errs, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD_FLAGS, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingCdmaNumber(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_NUMBER)
	value = interfaceToString(ivalue)
	return
}
func getSettingCdmaUsername(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_USERNAME)
	value = interfaceToString(ivalue)
	return
}
func getSettingCdmaPassword(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD)
	value = interfaceToString(ivalue)
	return
}
func getSettingCdmaPasswordFlags(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD_FLAGS)
	value = interfaceToUint32(ivalue)
	return
}

// Setter
func setSettingCdmaNumber(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_NUMBER, value)
}
func setSettingCdmaUsername(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_USERNAME, value)
}
func setSettingCdmaPassword(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD, value)
}
func setSettingCdmaPasswordFlags(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD_FLAGS, value)
}

// JSON Getter
func getSettingCdmaNumberJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_NUMBER, getSettingCdmaKeyType(NM_SETTING_CDMA_NUMBER))
	return
}
func getSettingCdmaUsernameJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_USERNAME, getSettingCdmaKeyType(NM_SETTING_CDMA_USERNAME))
	return
}
func getSettingCdmaPasswordJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD, getSettingCdmaKeyType(NM_SETTING_CDMA_PASSWORD))
	return
}
func getSettingCdmaPasswordFlagsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD_FLAGS, getSettingCdmaKeyType(NM_SETTING_CDMA_PASSWORD_FLAGS))
	return
}

// JSON Setter
func setSettingCdmaNumberJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_NUMBER, valueJSON, getSettingCdmaKeyType(NM_SETTING_CDMA_NUMBER))
}
func setSettingCdmaUsernameJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_USERNAME, valueJSON, getSettingCdmaKeyType(NM_SETTING_CDMA_USERNAME))
}
func setSettingCdmaPasswordJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD, valueJSON, getSettingCdmaKeyType(NM_SETTING_CDMA_PASSWORD))
}
func setSettingCdmaPasswordFlagsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD_FLAGS, valueJSON, getSettingCdmaKeyType(NM_SETTING_CDMA_PASSWORD_FLAGS))
}

// Logic JSON Setter

// Remover
func removeSettingCdmaNumber(data connectionData) {
	removeSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_NUMBER)
}
func removeSettingCdmaUsername(data connectionData) {
	removeSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_USERNAME)
}
func removeSettingCdmaPassword(data connectionData) {
	removeSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD)
}
func removeSettingCdmaPasswordFlags(data connectionData) {
	removeSettingKey(data, NM_SETTING_CDMA_SETTING_NAME, NM_SETTING_CDMA_PASSWORD_FLAGS)
}
