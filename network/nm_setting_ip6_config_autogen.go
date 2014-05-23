// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingIp6ConfigKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_IP6_CONFIG_METHOD:
		t = ktypeString
	case NM_SETTING_IP6_CONFIG_ADDRESSES:
		t = ktypeWrapperIpv6Addresses
	case NM_SETTING_IP6_CONFIG_DNS:
		t = ktypeWrapperIpv6Dns
	case NM_SETTING_IP6_CONFIG_DNS_SEARCH:
		t = ktypeArrayString
	case NM_SETTING_IP6_CONFIG_ROUTES:
		t = ktypeWrapperIpv6Routes
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES:
		t = ktypeBoolean
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS:
		t = ktypeBoolean
	case NM_SETTING_IP6_CONFIG_NEVER_DEFAULT:
		t = ktypeBoolean
	case NM_SETTING_IP6_CONFIG_MAY_FAIL:
		t = ktypeBoolean
	case NM_SETTING_IP6_CONFIG_IP6_PRIVACY:
		t = ktypeInt32
	case NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME:
		t = ktypeString
	}
	return
}

// Check is key in current setting field
func isKeyInSettingIp6Config(key string) bool {
	switch key {
	case NM_SETTING_IP6_CONFIG_METHOD:
		return true
	case NM_SETTING_IP6_CONFIG_ADDRESSES:
		return true
	case NM_SETTING_IP6_CONFIG_DNS:
		return true
	case NM_SETTING_IP6_CONFIG_DNS_SEARCH:
		return true
	case NM_SETTING_IP6_CONFIG_ROUTES:
		return true
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES:
		return true
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS:
		return true
	case NM_SETTING_IP6_CONFIG_NEVER_DEFAULT:
		return true
	case NM_SETTING_IP6_CONFIG_MAY_FAIL:
		return true
	case NM_SETTING_IP6_CONFIG_IP6_PRIVACY:
		return true
	case NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME:
		return true
	}
	return false
}

// Get key's default value
func getSettingIp6ConfigDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_IP6_CONFIG_METHOD:
		value = ""
	case NM_SETTING_IP6_CONFIG_ADDRESSES:
		value = nil
	case NM_SETTING_IP6_CONFIG_DNS:
		value = nil
	case NM_SETTING_IP6_CONFIG_DNS_SEARCH:
		value = nil
	case NM_SETTING_IP6_CONFIG_ROUTES:
		value = nil
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES:
		value = false
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS:
		value = false
	case NM_SETTING_IP6_CONFIG_NEVER_DEFAULT:
		value = false
	case NM_SETTING_IP6_CONFIG_MAY_FAIL:
		value = true
	case NM_SETTING_IP6_CONFIG_IP6_PRIVACY:
		value = -1
	case NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME:
		value = ""
	}
	return
}

// Get JSON value generally
func generalGetSettingIp6ConfigKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingIp6ConfigKeyJSON: invalide key", key)
	case NM_SETTING_IP6_CONFIG_METHOD:
		value = getSettingIp6ConfigMethodJSON(data)
	case NM_SETTING_IP6_CONFIG_ADDRESSES:
		value = getSettingIp6ConfigAddressesJSON(data)
	case NM_SETTING_IP6_CONFIG_DNS:
		value = getSettingIp6ConfigDnsJSON(data)
	case NM_SETTING_IP6_CONFIG_DNS_SEARCH:
		value = getSettingIp6ConfigDnsSearchJSON(data)
	case NM_SETTING_IP6_CONFIG_ROUTES:
		value = getSettingIp6ConfigRoutesJSON(data)
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES:
		value = getSettingIp6ConfigIgnoreAutoRoutesJSON(data)
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS:
		value = getSettingIp6ConfigIgnoreAutoDnsJSON(data)
	case NM_SETTING_IP6_CONFIG_NEVER_DEFAULT:
		value = getSettingIp6ConfigNeverDefaultJSON(data)
	case NM_SETTING_IP6_CONFIG_MAY_FAIL:
		value = getSettingIp6ConfigMayFailJSON(data)
	case NM_SETTING_IP6_CONFIG_IP6_PRIVACY:
		value = getSettingIp6ConfigIp6PrivacyJSON(data)
	case NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME:
		value = getSettingIp6ConfigDhcpHostnameJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingIp6ConfigKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingIp6ConfigKeyJSON: invalide key", key)
	case NM_SETTING_IP6_CONFIG_METHOD:
		err = logicSetSettingIp6ConfigMethodJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_ADDRESSES:
		err = setSettingIp6ConfigAddressesJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_DNS:
		err = setSettingIp6ConfigDnsJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_DNS_SEARCH:
		err = setSettingIp6ConfigDnsSearchJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_ROUTES:
		err = setSettingIp6ConfigRoutesJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES:
		err = setSettingIp6ConfigIgnoreAutoRoutesJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS:
		err = setSettingIp6ConfigIgnoreAutoDnsJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_NEVER_DEFAULT:
		err = setSettingIp6ConfigNeverDefaultJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_MAY_FAIL:
		err = setSettingIp6ConfigMayFailJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_IP6_PRIVACY:
		err = setSettingIp6ConfigIp6PrivacyJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME:
		err = setSettingIp6ConfigDhcpHostnameJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingIp6ConfigMethodExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD)
}
func isSettingIp6ConfigAddressesExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES)
}
func isSettingIp6ConfigDnsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS)
}
func isSettingIp6ConfigDnsSearchExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH)
}
func isSettingIp6ConfigRoutesExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES)
}
func isSettingIp6ConfigIgnoreAutoRoutesExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES)
}
func isSettingIp6ConfigIgnoreAutoDnsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS)
}
func isSettingIp6ConfigNeverDefaultExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT)
}
func isSettingIp6ConfigMayFailExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL)
}
func isSettingIp6ConfigIp6PrivacyExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY)
}
func isSettingIp6ConfigDhcpHostnameExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME)
}

// Ensure field and key exists and not empty
func ensureFieldSettingIp6ConfigExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_IP6_CONFIG_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_IP6_CONFIG_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_IP6_CONFIG_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_IP6_CONFIG_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_IP6_CONFIG_SETTING_NAME))
	}
}
func ensureSettingIp6ConfigMethodNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp6ConfigMethodExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigMethod(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp6ConfigAddressesNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp6ConfigAddressesExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigAddresses(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp6ConfigDnsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp6ConfigDnsExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigDns(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp6ConfigDnsSearchNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp6ConfigDnsSearchExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigDnsSearch(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp6ConfigRoutesNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp6ConfigRoutesExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigRoutes(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp6ConfigIgnoreAutoRoutesNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp6ConfigIgnoreAutoRoutesExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp6ConfigIgnoreAutoDnsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp6ConfigIgnoreAutoDnsExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp6ConfigNeverDefaultNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp6ConfigNeverDefaultExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp6ConfigMayFailNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp6ConfigMayFailExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp6ConfigIp6PrivacyNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp6ConfigIp6PrivacyExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp6ConfigDhcpHostnameNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp6ConfigDhcpHostnameExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigDhcpHostname(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME, NM_KEY_ERROR_EMPTY_VALUE)
	}
}

// Getter
func getSettingIp6ConfigMethod(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD)
	value = interfaceToString(ivalue)
	return
}
func getSettingIp6ConfigAddresses(data connectionData) (value ipv6Addresses) {
	ivalue := getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES)
	value = interfaceToIpv6Addresses(ivalue)
	return
}
func getSettingIp6ConfigDns(data connectionData) (value [][]byte) {
	ivalue := getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS)
	value = interfaceToArrayArrayByte(ivalue)
	return
}
func getSettingIp6ConfigDnsSearch(data connectionData) (value []string) {
	ivalue := getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH)
	value = interfaceToArrayString(ivalue)
	return
}
func getSettingIp6ConfigRoutes(data connectionData) (value ipv6Routes) {
	ivalue := getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES)
	value = interfaceToIpv6Routes(ivalue)
	return
}
func getSettingIp6ConfigIgnoreAutoRoutes(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES)
	value = interfaceToBoolean(ivalue)
	return
}
func getSettingIp6ConfigIgnoreAutoDns(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS)
	value = interfaceToBoolean(ivalue)
	return
}
func getSettingIp6ConfigNeverDefault(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT)
	value = interfaceToBoolean(ivalue)
	return
}
func getSettingIp6ConfigMayFail(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL)
	value = interfaceToBoolean(ivalue)
	return
}
func getSettingIp6ConfigIp6Privacy(data connectionData) (value int32) {
	ivalue := getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY)
	value = interfaceToInt32(ivalue)
	return
}
func getSettingIp6ConfigDhcpHostname(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME)
	value = interfaceToString(ivalue)
	return
}

// Setter
func setSettingIp6ConfigMethod(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD, value)
}
func setSettingIp6ConfigAddresses(data connectionData, value ipv6Addresses) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, value)
}
func setSettingIp6ConfigDns(data connectionData, value [][]byte) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS, value)
}
func setSettingIp6ConfigDnsSearch(data connectionData, value []string) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH, value)
}
func setSettingIp6ConfigRoutes(data connectionData, value ipv6Routes) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, value)
}
func setSettingIp6ConfigIgnoreAutoRoutes(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES, value)
}
func setSettingIp6ConfigIgnoreAutoDns(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS, value)
}
func setSettingIp6ConfigNeverDefault(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT, value)
}
func setSettingIp6ConfigMayFail(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL, value)
}
func setSettingIp6ConfigIp6Privacy(data connectionData, value int32) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY, value)
}
func setSettingIp6ConfigDhcpHostname(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME, value)
}

// JSON Getter
func getSettingIp6ConfigMethodJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_METHOD))
	return
}
func getSettingIp6ConfigAddressesJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_ADDRESSES))
	return
}
func getSettingIp6ConfigDnsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DNS))
	return
}
func getSettingIp6ConfigDnsSearchJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DNS_SEARCH))
	return
}
func getSettingIp6ConfigRoutesJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_ROUTES))
	return
}
func getSettingIp6ConfigIgnoreAutoRoutesJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES))
	return
}
func getSettingIp6ConfigIgnoreAutoDnsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS))
	return
}
func getSettingIp6ConfigNeverDefaultJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_NEVER_DEFAULT))
	return
}
func getSettingIp6ConfigMayFailJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_MAY_FAIL))
	return
}
func getSettingIp6ConfigIp6PrivacyJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IP6_PRIVACY))
	return
}
func getSettingIp6ConfigDhcpHostnameJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME))
	return
}

// JSON Setter
func setSettingIp6ConfigMethodJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_METHOD))
}
func setSettingIp6ConfigAddressesJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_ADDRESSES))
}
func setSettingIp6ConfigDnsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DNS))
}
func setSettingIp6ConfigDnsSearchJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DNS_SEARCH))
}
func setSettingIp6ConfigRoutesJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_ROUTES))
}
func setSettingIp6ConfigIgnoreAutoRoutesJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES))
}
func setSettingIp6ConfigIgnoreAutoDnsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS))
}
func setSettingIp6ConfigNeverDefaultJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_NEVER_DEFAULT))
}
func setSettingIp6ConfigMayFailJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_MAY_FAIL))
}
func setSettingIp6ConfigIp6PrivacyJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IP6_PRIVACY))
}
func setSettingIp6ConfigDhcpHostnameJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME))
}

// Logic JSON Setter
func logicSetSettingIp6ConfigMethodJSON(data connectionData, valueJSON string) (err error) {
	err = setSettingIp6ConfigMethodJSON(data, valueJSON)
	if err != nil {
		return
	}
	if isSettingIp6ConfigMethodExists(data) {
		value := getSettingIp6ConfigMethod(data)
		err = logicSetSettingIp6ConfigMethod(data, value)
	}
	return
}

// Remover
func removeSettingIp6ConfigMethod(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD)
}
func removeSettingIp6ConfigAddresses(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES)
}
func removeSettingIp6ConfigDns(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS)
}
func removeSettingIp6ConfigDnsSearch(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH)
}
func removeSettingIp6ConfigRoutes(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES)
}
func removeSettingIp6ConfigIgnoreAutoRoutes(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES)
}
func removeSettingIp6ConfigIgnoreAutoDns(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS)
}
func removeSettingIp6ConfigNeverDefault(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT)
}
func removeSettingIp6ConfigMayFail(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL)
}
func removeSettingIp6ConfigIp6Privacy(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY)
}
func removeSettingIp6ConfigDhcpHostname(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME)
}
