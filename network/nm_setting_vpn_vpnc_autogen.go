// This file is automatically generated, please don't edit manully.
package network

import (
	"fmt"
)

// Get key type
func getSettingVpnVpncKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_VPN_VPNC_KEY_GATEWAY:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_USER:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS:
		t = ktypeUint32
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_ID:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS:
		t = ktypeUint32
	case NM_SETTING_VPN_VPNC_KEY_SECRET:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_AUTHMODE:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_CA_FILE:
		t = ktypeString
	}
	return
}

// Check is key in current setting section
func isKeyInSettingVpnVpnc(key string) bool {
	switch key {
	case NM_SETTING_VPN_VPNC_KEY_GATEWAY:
		return true
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_USER:
		return true
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE:
		return true
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS:
		return true
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD:
		return true
	case NM_SETTING_VPN_VPNC_KEY_ID:
		return true
	case NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE:
		return true
	case NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS:
		return true
	case NM_SETTING_VPN_VPNC_KEY_SECRET:
		return true
	case NM_SETTING_VPN_VPNC_KEY_AUTHMODE:
		return true
	case NM_SETTING_VPN_VPNC_KEY_CA_FILE:
		return true
	}
	return false
}

// Get key's default value
func getSettingVpnVpncDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_VPN_VPNC_KEY_GATEWAY:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_USER:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS:
		value = uint32(0)
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_ID:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS:
		value = uint32(0)
	case NM_SETTING_VPN_VPNC_KEY_SECRET:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_AUTHMODE:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_CA_FILE:
		value = ""
	}
	return
}

// Get JSON value generally
func generalGetSettingVpnVpncKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingVpnVpncKeyJSON: invalide key", key)
	case NM_SETTING_VPN_VPNC_KEY_GATEWAY:
		value = getSettingVpnVpncKeyGatewayJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_USER:
		value = getSettingVpnVpncKeyXauthUserJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE:
		value = getSettingVpnVpncKeyXauthPasswordTypeJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS:
		value = getSettingVpnVpncKeyXauthPasswordFlagsJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD:
		value = getSettingVpnVpncKeyXauthPasswordJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_ID:
		value = getSettingVpnVpncKeyIdJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE:
		value = getSettingVpnVpncKeySecretTypeJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS:
		value = getSettingVpnVpncKeySecretFlagsJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_SECRET:
		value = getSettingVpnVpncKeySecretJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_AUTHMODE:
		value = getSettingVpnVpncKeyAuthmodeJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_CA_FILE:
		value = getSettingVpnVpncKeyCaFileJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingVpnVpncKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingVpnVpncKeyJSON: invalide key", key)
	case NM_SETTING_VPN_VPNC_KEY_GATEWAY:
		err = setSettingVpnVpncKeyGatewayJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_USER:
		err = setSettingVpnVpncKeyXauthUserJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE:
		err = setSettingVpnVpncKeyXauthPasswordTypeJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS:
		err = logicSetSettingVpnVpncKeyXauthPasswordFlagsJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD:
		err = setSettingVpnVpncKeyXauthPasswordJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_ID:
		err = setSettingVpnVpncKeyIdJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE:
		err = setSettingVpnVpncKeySecretTypeJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS:
		err = logicSetSettingVpnVpncKeySecretFlagsJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_SECRET:
		err = setSettingVpnVpncKeySecretJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_AUTHMODE:
		err = setSettingVpnVpncKeyAuthmodeJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_CA_FILE:
		err = setSettingVpnVpncKeyCaFileJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingVpnVpncKeyGatewayExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_GATEWAY)
}
func isSettingVpnVpncKeyXauthUserExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_USER)
}
func isSettingVpnVpncKeyXauthPasswordTypeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE)
}
func isSettingVpnVpncKeyXauthPasswordFlagsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS)
}
func isSettingVpnVpncKeyXauthPasswordExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD)
}
func isSettingVpnVpncKeyIdExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_ID)
}
func isSettingVpnVpncKeySecretTypeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE)
}
func isSettingVpnVpncKeySecretFlagsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS)
}
func isSettingVpnVpncKeySecretExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET)
}
func isSettingVpnVpncKeyAuthmodeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_AUTHMODE)
}
func isSettingVpnVpncKeyCaFileExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CA_FILE)
}

// Ensure section and key exists and not empty
func ensureSectionSettingVpnVpncExists(data connectionData, errs sectionErrors, relatedKey string) {
	if !isSettingSectionExists(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME))
	}
	sectionData, _ := data[NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME]
	if len(sectionData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME))
	}
}
func ensureSettingVpnVpncKeyGatewayNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyGatewayExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_GATEWAY, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyGateway(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_GATEWAY, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeyXauthUserNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyXauthUserExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_USER, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyXauthUser(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_USER, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeyXauthPasswordTypeNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyXauthPasswordTypeExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyXauthPasswordType(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeyXauthPasswordFlagsNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyXauthPasswordFlagsExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnVpncKeyXauthPasswordNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyXauthPasswordExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyXauthPassword(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeyIdNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyIdExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_ID, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyId(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_ID, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeySecretTypeNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeySecretTypeExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeySecretType(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeySecretFlagsNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeySecretFlagsExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnVpncKeySecretNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeySecretExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeySecret(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeyAuthmodeNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyAuthmodeExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_AUTHMODE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyAuthmode(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_AUTHMODE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeyCaFileNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyCaFileExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CA_FILE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyCaFile(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CA_FILE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}

// Getter
func getSettingVpnVpncKeyGateway(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_GATEWAY)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeyXauthUser(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_USER)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeyXauthPasswordType(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeyXauthPasswordFlags(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS)
	value = interfaceToUint32(ivalue)
	return
}
func getSettingVpnVpncKeyXauthPassword(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeyId(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_ID)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeySecretType(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeySecretFlags(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS)
	value = interfaceToUint32(ivalue)
	return
}
func getSettingVpnVpncKeySecret(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeyAuthmode(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_AUTHMODE)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeyCaFile(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CA_FILE)
	value = interfaceToString(ivalue)
	return
}

// Setter
func setSettingVpnVpncKeyGateway(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_GATEWAY, value)
}
func setSettingVpnVpncKeyXauthUser(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_USER, value)
}
func setSettingVpnVpncKeyXauthPasswordType(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE, value)
}
func setSettingVpnVpncKeyXauthPasswordFlags(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS, value)
}
func setSettingVpnVpncKeyXauthPassword(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD, value)
}
func setSettingVpnVpncKeyId(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_ID, value)
}
func setSettingVpnVpncKeySecretType(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE, value)
}
func setSettingVpnVpncKeySecretFlags(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS, value)
}
func setSettingVpnVpncKeySecret(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET, value)
}
func setSettingVpnVpncKeyAuthmode(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_AUTHMODE, value)
}
func setSettingVpnVpncKeyCaFile(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CA_FILE, value)
}

// JSON Getter
func getSettingVpnVpncKeyGatewayJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_GATEWAY, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_GATEWAY))
	return
}
func getSettingVpnVpncKeyXauthUserJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_USER, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_XAUTH_USER))
	return
}
func getSettingVpnVpncKeyXauthPasswordTypeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE))
	return
}
func getSettingVpnVpncKeyXauthPasswordFlagsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS))
	return
}
func getSettingVpnVpncKeyXauthPasswordJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD))
	return
}
func getSettingVpnVpncKeyIdJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_ID, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_ID))
	return
}
func getSettingVpnVpncKeySecretTypeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE))
	return
}
func getSettingVpnVpncKeySecretFlagsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS))
	return
}
func getSettingVpnVpncKeySecretJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_SECRET))
	return
}
func getSettingVpnVpncKeyAuthmodeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_AUTHMODE, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_AUTHMODE))
	return
}
func getSettingVpnVpncKeyCaFileJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CA_FILE, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_CA_FILE))
	return
}

// JSON Setter
func setSettingVpnVpncKeyGatewayJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_GATEWAY, valueJSON, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_GATEWAY))
}
func setSettingVpnVpncKeyXauthUserJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_USER, valueJSON, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_XAUTH_USER))
}
func setSettingVpnVpncKeyXauthPasswordTypeJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE, valueJSON, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE))
}
func setSettingVpnVpncKeyXauthPasswordFlagsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS, valueJSON, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS))
}
func setSettingVpnVpncKeyXauthPasswordJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD, valueJSON, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD))
}
func setSettingVpnVpncKeyIdJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_ID, valueJSON, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_ID))
}
func setSettingVpnVpncKeySecretTypeJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE, valueJSON, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE))
}
func setSettingVpnVpncKeySecretFlagsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS, valueJSON, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS))
}
func setSettingVpnVpncKeySecretJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET, valueJSON, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_SECRET))
}
func setSettingVpnVpncKeyAuthmodeJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_AUTHMODE, valueJSON, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_AUTHMODE))
}
func setSettingVpnVpncKeyCaFileJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CA_FILE, valueJSON, getSettingVpnVpncKeyType(NM_SETTING_VPN_VPNC_KEY_CA_FILE))
}

// Logic JSON Setter
func logicSetSettingVpnVpncKeyXauthPasswordFlagsJSON(data connectionData, valueJSON string) (err error) {
	err = setSettingVpnVpncKeyXauthPasswordFlagsJSON(data, valueJSON)
	if err != nil {
		return
	}
	if isSettingVpnVpncKeyXauthPasswordFlagsExists(data) {
		value := getSettingVpnVpncKeyXauthPasswordFlags(data)
		err = logicSetSettingVpnVpncKeyXauthPasswordFlags(data, value)
	}
	return
}
func logicSetSettingVpnVpncKeySecretFlagsJSON(data connectionData, valueJSON string) (err error) {
	err = setSettingVpnVpncKeySecretFlagsJSON(data, valueJSON)
	if err != nil {
		return
	}
	if isSettingVpnVpncKeySecretFlagsExists(data) {
		value := getSettingVpnVpncKeySecretFlags(data)
		err = logicSetSettingVpnVpncKeySecretFlags(data, value)
	}
	return
}

// Remover
func removeSettingVpnVpncKeyGateway(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_GATEWAY)
}
func removeSettingVpnVpncKeyXauthUser(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_USER)
}
func removeSettingVpnVpncKeyXauthPasswordType(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_TYPE)
}
func removeSettingVpnVpncKeyXauthPasswordFlags(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD_FLAGS)
}
func removeSettingVpnVpncKeyXauthPassword(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_XAUTH_PASSWORD)
}
func removeSettingVpnVpncKeyId(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_ID)
}
func removeSettingVpnVpncKeySecretType(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_TYPE)
}
func removeSettingVpnVpncKeySecretFlags(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET_FLAGS)
}
func removeSettingVpnVpncKeySecret(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SECRET)
}
func removeSettingVpnVpncKeyAuthmode(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_AUTHMODE)
}
func removeSettingVpnVpncKeyCaFile(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CA_FILE)
}
