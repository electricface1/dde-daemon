package main

import (
	"bytes"
	"flag"
	"fmt"
	"path"
	"text/template"
)

var funcMap = template.FuncMap{
	"ToCaplitalize":          ToCaplitalize,
	"ToFieldFuncBaseName":    ToFieldFuncBaseName,
	"ToKeyFuncBaseName":      ToKeyFuncBaseName,
	"ToKeyTypeRealData":      ToKeyTypeRealData,
	"ToKeyTypeDefaultValue":  ToKeyTypeDefaultValue,
	"IfNeedCheckValueLength": IfNeedCheckValueLength,
	"GetAllVkFields":         GetAllVkFields,
	"GetAllVkFieldKeys":      GetAllVkFieldKeys,
	// "IsVkNeedLogicSetter":       IsVkNeedLogicSetter,
	"ToKeyTypeShortName":          ToKeyTypeShortName,
	"ToKeyDisplayName":            ToKeyDisplayName,
	"ToKeyValue":                  ToKeyValue,
	"IsKeyUsedByFrontEnd":         IsKeyUsedByFrontEnd,
	"ToFrontEndWidget":            ToFrontEndWidget,
	"ToClassName":                 ToClassName,
	"GetAllKeysInPage":            GetAllKeysInPage,
	"GetKeyWidgetProp":            GetKeyWidgetProp,
	"ToKeyTypeInterfaceConverter": ToKeyTypeInterfaceConverter,
}

const (
	backEndDir            = ".."
	frontEndDir           = "../../../dss/modules/network/components_autogen/"
	nmSettingsJSONFile    = "./nm_settings.json"
	nmSettingVkJSONFile   = "./nm_setting_vk.json"
	nmSettingPageJSONFile = "./nm_setting_page.json"
)

var (
	argWriteOutput       bool
	argBackEnd           bool
	argFrontEnd          bool
	nmSettingUtilsFile   = path.Join(backEndDir, "nm_setting_general_autogen.go")
	nmSettingVkFile      = path.Join(backEndDir, "nm_setting_virtual_key_autogen.go")
	frontEndConnPropFile = path.Join(frontEndDir, "BaseConnectionProperties.qml")
	nmSettings           []NMSettingStruct
	nmSettingVks         []NMSettingVkStruct
	nmSettingPages       []NMSettingPageStruct
)

type NMSettingStruct struct {
	FieldName  string // such as "NM_SETTING_CONNECTION_SETTING_NAME"
	FieldValue string // such as "connection"
	Keys       []NMSettingKeyStruct
}

type NMSettingKeyStruct struct {
	Name           string            // such as "NM_SETTING_CONNECTION_ID"
	Value          string            // such as "id"
	Type           string            // such as "ktypeString"
	Default        string            // such as "<default>", "<null>" or "true"
	UsedByBackEnd  bool              // determine if this key will be used by back-end(golang code)
	UsedByFrontEnd bool              // determine if this key will be used by front-end(qml code)
	LogicSet       bool              // determine if this key should to generate a logic setter
	DisplayName    string            // such as "Connection name"
	FrontEndWidget string            // such as "EditLinePasswordInput"
	WidgetProp     map[string]string // properties for front end widget, such as "WidgetProp":{"alwaysUpdate":"true"}
}

type NMSettingVkStruct struct {
	Name           string // such as "NM_SETTING_VK_802_1X_EAP"
	Value          string // such as "vk-eap"
	Type           string // such as "ktypeString"
	RelatedField   string // such as "NM_SETTING_802_1X_SETTING_NAME"
	RelatedKey     string // such as "NM_SETTING_802_1X_EAP"
	EnableWrapper  bool   // check if the virtual key is a wrapper just to enable target key
	UsedByFrontEnd bool   // check if is used by front-end
	Optional       bool   // if key is optional, will ignore error for it
	DisplayName    string
	FrontEndWidget string            // such as "EditLinePasswordInput"
	WidgetProp     map[string]string // properties for front end widget, such as "WidgetProp":{"alwaysUpdate":"true"}
}

type NMSettingPageStruct struct {
	Ignore        bool
	Name          string
	DisplayName   string
	RelatedFields []string
}

func genNMSettingCode(nmSetting NMSettingStruct) (content string) {
	content = fileHeader
	content += genTpl(nmSetting, tplGetKeyType)          // get key type
	content += genTpl(nmSetting, tplIsKeyInSettingField) // check is key in current field
	content += genTpl(nmSetting, tplGetDefaultValue)     // get default value
	content += genTpl(nmSetting, tplGeneralGetterJSON)   // general json getter
	content += genTpl(nmSetting, tplGeneralSetterJSON)   // general json setter
	content += genTpl(nmSetting, tplCheckExists)         // check if key exists
	content += genTpl(nmSetting, tplEnsureNoEmpty)       // ensure field and key exists and not empty
	content += genTpl(nmSetting, tplGetter)              // getter
	content += genTpl(nmSetting, tplSetter)              // setter
	content += genTpl(nmSetting, tplJSONGetter)          // json getter
	content += genTpl(nmSetting, tplJSONSetter)          // json setter
	content += genTpl(nmSetting, tplLogicJSONSetter)     // logic json setter
	content += genTpl(nmSetting, tplRemover)             // remover
	// TODO get avaiable values
	return
}

func genNMSettingGeneralUtilsCode(nmSettings []NMSettingStruct) (content string) {
	content = genTpl(nmSettings, tplGeneralSettingUtils) // general setting utils
	return
}

func genNMSettingVirtualKeyCode(nmSettings []NMSettingStruct, nmSettingVks []NMSettingVkStruct) (content string) {
	content = genTpl(nmSettingVks, tplVirtualKey) // general setting utils
	return
}

func genFrontEndConnPropCode(nmPages []NMSettingPageStruct) (content string) {
	content = genTpl(nmPages, tplFrontEndConnProp)
	return
}

func genFrontEndSectionCode(nmPage NMSettingPageStruct) (content string) {
	content = genTpl(nmPage, tplFrontEndSection)
	return
}

func genTpl(data interface{}, tplstr string) (content string) {
	templator := template.New("nm autogen").Funcs(funcMap)
	tpl, err := templator.Parse(tplstr)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := bytes.NewBufferString("")
	err = tpl.Execute(buf, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	content = string(buf.Bytes())
	return
}

func genBackEndCode() {
	// back-end code, echo nm setting fields
	for _, nmSetting := range nmSettings {
		autogenContent := genNMSettingCode(nmSetting)
		backEndFile := getBackEndFilePath(nmSetting.FieldName)
		writeOrDisplayResultForBackEnd(backEndFile, autogenContent)
	}

	// back-end code, general setting utils
	autogenContent := genNMSettingGeneralUtilsCode(nmSettings)
	writeOrDisplayResultForBackEnd(nmSettingUtilsFile, autogenContent)

	// back-end code, virtual key
	autogenContent = genNMSettingVirtualKeyCode(nmSettings, nmSettingVks)
	writeOrDisplayResultForBackEnd(nmSettingVkFile, autogenContent)
}

func genFrontEndCode() {
	// front-end code, BaseConnectionProperties.qml
	autogenContent := genFrontEndConnPropCode(nmSettingPages)
	writeOrDisplayResultForFrontEnd(frontEndConnPropFile, autogenContent)

	for _, nmPage := range nmSettingPages {
		if nmPage.Ignore {
			continue
		}
		autogenContent = genFrontEndSectionCode(nmPage)
		frontEndFile := getFrontEndFilePath(nmPage.Name)
		writeOrDisplayResultForFrontEnd(frontEndFile, autogenContent)
	}
}

func main() {
	flag.BoolVar(&argWriteOutput, "w", false, "write to file")
	flag.BoolVar(&argBackEnd, "b", false, "generate back-end code")
	flag.BoolVar(&argFrontEnd, "f", false, "generate front-end code")
	flag.Parse()

	unmarshalJSONFile(nmSettingsJSONFile, &nmSettings)
	unmarshalJSONFile(nmSettingVkJSONFile, &nmSettingVks)
	unmarshalJSONFile(nmSettingPageJSONFile, &nmSettingPages)
	if argBackEnd {
		genBackEndCode()
	}
	if argFrontEnd {
		genFrontEndCode()
	}
}
