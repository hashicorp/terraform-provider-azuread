package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRulesValidationError string

const (
	DeviceComplianceScriptRulesValidationError_DatatypeMissing                         DeviceComplianceScriptRulesValidationError = "datatypeMissing"
	DeviceComplianceScriptRulesValidationError_DatatypeNotSupported                    DeviceComplianceScriptRulesValidationError = "datatypeNotSupported"
	DeviceComplianceScriptRulesValidationError_DescriptionInvalid                      DeviceComplianceScriptRulesValidationError = "descriptionInvalid"
	DeviceComplianceScriptRulesValidationError_DescriptionMissing                      DeviceComplianceScriptRulesValidationError = "descriptionMissing"
	DeviceComplianceScriptRulesValidationError_DescriptionTooLarge                     DeviceComplianceScriptRulesValidationError = "descriptionTooLarge"
	DeviceComplianceScriptRulesValidationError_DuplicateLocales                        DeviceComplianceScriptRulesValidationError = "duplicateLocales"
	DeviceComplianceScriptRulesValidationError_DuplicateRules                          DeviceComplianceScriptRulesValidationError = "duplicateRules"
	DeviceComplianceScriptRulesValidationError_EnglishLocaleMissing                    DeviceComplianceScriptRulesValidationError = "englishLocaleMissing"
	DeviceComplianceScriptRulesValidationError_JsonFileInvalid                         DeviceComplianceScriptRulesValidationError = "jsonFileInvalid"
	DeviceComplianceScriptRulesValidationError_JsonFileMissing                         DeviceComplianceScriptRulesValidationError = "jsonFileMissing"
	DeviceComplianceScriptRulesValidationError_JsonFileTooLarge                        DeviceComplianceScriptRulesValidationError = "jsonFileTooLarge"
	DeviceComplianceScriptRulesValidationError_MoreInfoUriInvalid                      DeviceComplianceScriptRulesValidationError = "moreInfoUriInvalid"
	DeviceComplianceScriptRulesValidationError_MoreInfoUriMissing                      DeviceComplianceScriptRulesValidationError = "moreInfoUriMissing"
	DeviceComplianceScriptRulesValidationError_MoreInfoUriTooLarge                     DeviceComplianceScriptRulesValidationError = "moreInfoUriTooLarge"
	DeviceComplianceScriptRulesValidationError_None                                    DeviceComplianceScriptRulesValidationError = "none"
	DeviceComplianceScriptRulesValidationError_OperandInvalid                          DeviceComplianceScriptRulesValidationError = "operandInvalid"
	DeviceComplianceScriptRulesValidationError_OperandMissing                          DeviceComplianceScriptRulesValidationError = "operandMissing"
	DeviceComplianceScriptRulesValidationError_OperandTooLarge                         DeviceComplianceScriptRulesValidationError = "operandTooLarge"
	DeviceComplianceScriptRulesValidationError_OperatorDataTypeCombinationNotSupported DeviceComplianceScriptRulesValidationError = "operatorDataTypeCombinationNotSupported"
	DeviceComplianceScriptRulesValidationError_OperatorMissing                         DeviceComplianceScriptRulesValidationError = "operatorMissing"
	DeviceComplianceScriptRulesValidationError_OperatorNotSupported                    DeviceComplianceScriptRulesValidationError = "operatorNotSupported"
	DeviceComplianceScriptRulesValidationError_RemediationStringsMissing               DeviceComplianceScriptRulesValidationError = "remediationStringsMissing"
	DeviceComplianceScriptRulesValidationError_RulesMissing                            DeviceComplianceScriptRulesValidationError = "rulesMissing"
	DeviceComplianceScriptRulesValidationError_SettingNameInvalid                      DeviceComplianceScriptRulesValidationError = "settingNameInvalid"
	DeviceComplianceScriptRulesValidationError_SettingNameMissing                      DeviceComplianceScriptRulesValidationError = "settingNameMissing"
	DeviceComplianceScriptRulesValidationError_SettingNameTooLarge                     DeviceComplianceScriptRulesValidationError = "settingNameTooLarge"
	DeviceComplianceScriptRulesValidationError_TitleInvalid                            DeviceComplianceScriptRulesValidationError = "titleInvalid"
	DeviceComplianceScriptRulesValidationError_TitleMissing                            DeviceComplianceScriptRulesValidationError = "titleMissing"
	DeviceComplianceScriptRulesValidationError_TitleTooLarge                           DeviceComplianceScriptRulesValidationError = "titleTooLarge"
	DeviceComplianceScriptRulesValidationError_TooManyRulesSpecified                   DeviceComplianceScriptRulesValidationError = "tooManyRulesSpecified"
	DeviceComplianceScriptRulesValidationError_Unknown                                 DeviceComplianceScriptRulesValidationError = "unknown"
	DeviceComplianceScriptRulesValidationError_UnrecognizedLocale                      DeviceComplianceScriptRulesValidationError = "unrecognizedLocale"
)

func PossibleValuesForDeviceComplianceScriptRulesValidationError() []string {
	return []string{
		string(DeviceComplianceScriptRulesValidationError_DatatypeMissing),
		string(DeviceComplianceScriptRulesValidationError_DatatypeNotSupported),
		string(DeviceComplianceScriptRulesValidationError_DescriptionInvalid),
		string(DeviceComplianceScriptRulesValidationError_DescriptionMissing),
		string(DeviceComplianceScriptRulesValidationError_DescriptionTooLarge),
		string(DeviceComplianceScriptRulesValidationError_DuplicateLocales),
		string(DeviceComplianceScriptRulesValidationError_DuplicateRules),
		string(DeviceComplianceScriptRulesValidationError_EnglishLocaleMissing),
		string(DeviceComplianceScriptRulesValidationError_JsonFileInvalid),
		string(DeviceComplianceScriptRulesValidationError_JsonFileMissing),
		string(DeviceComplianceScriptRulesValidationError_JsonFileTooLarge),
		string(DeviceComplianceScriptRulesValidationError_MoreInfoUriInvalid),
		string(DeviceComplianceScriptRulesValidationError_MoreInfoUriMissing),
		string(DeviceComplianceScriptRulesValidationError_MoreInfoUriTooLarge),
		string(DeviceComplianceScriptRulesValidationError_None),
		string(DeviceComplianceScriptRulesValidationError_OperandInvalid),
		string(DeviceComplianceScriptRulesValidationError_OperandMissing),
		string(DeviceComplianceScriptRulesValidationError_OperandTooLarge),
		string(DeviceComplianceScriptRulesValidationError_OperatorDataTypeCombinationNotSupported),
		string(DeviceComplianceScriptRulesValidationError_OperatorMissing),
		string(DeviceComplianceScriptRulesValidationError_OperatorNotSupported),
		string(DeviceComplianceScriptRulesValidationError_RemediationStringsMissing),
		string(DeviceComplianceScriptRulesValidationError_RulesMissing),
		string(DeviceComplianceScriptRulesValidationError_SettingNameInvalid),
		string(DeviceComplianceScriptRulesValidationError_SettingNameMissing),
		string(DeviceComplianceScriptRulesValidationError_SettingNameTooLarge),
		string(DeviceComplianceScriptRulesValidationError_TitleInvalid),
		string(DeviceComplianceScriptRulesValidationError_TitleMissing),
		string(DeviceComplianceScriptRulesValidationError_TitleTooLarge),
		string(DeviceComplianceScriptRulesValidationError_TooManyRulesSpecified),
		string(DeviceComplianceScriptRulesValidationError_Unknown),
		string(DeviceComplianceScriptRulesValidationError_UnrecognizedLocale),
	}
}

func (s *DeviceComplianceScriptRulesValidationError) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptRulesValidationError(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptRulesValidationError(input string) (*DeviceComplianceScriptRulesValidationError, error) {
	vals := map[string]DeviceComplianceScriptRulesValidationError{
		"datatypemissing":                         DeviceComplianceScriptRulesValidationError_DatatypeMissing,
		"datatypenotsupported":                    DeviceComplianceScriptRulesValidationError_DatatypeNotSupported,
		"descriptioninvalid":                      DeviceComplianceScriptRulesValidationError_DescriptionInvalid,
		"descriptionmissing":                      DeviceComplianceScriptRulesValidationError_DescriptionMissing,
		"descriptiontoolarge":                     DeviceComplianceScriptRulesValidationError_DescriptionTooLarge,
		"duplicatelocales":                        DeviceComplianceScriptRulesValidationError_DuplicateLocales,
		"duplicaterules":                          DeviceComplianceScriptRulesValidationError_DuplicateRules,
		"englishlocalemissing":                    DeviceComplianceScriptRulesValidationError_EnglishLocaleMissing,
		"jsonfileinvalid":                         DeviceComplianceScriptRulesValidationError_JsonFileInvalid,
		"jsonfilemissing":                         DeviceComplianceScriptRulesValidationError_JsonFileMissing,
		"jsonfiletoolarge":                        DeviceComplianceScriptRulesValidationError_JsonFileTooLarge,
		"moreinfouriinvalid":                      DeviceComplianceScriptRulesValidationError_MoreInfoUriInvalid,
		"moreinfourimissing":                      DeviceComplianceScriptRulesValidationError_MoreInfoUriMissing,
		"moreinfouritoolarge":                     DeviceComplianceScriptRulesValidationError_MoreInfoUriTooLarge,
		"none":                                    DeviceComplianceScriptRulesValidationError_None,
		"operandinvalid":                          DeviceComplianceScriptRulesValidationError_OperandInvalid,
		"operandmissing":                          DeviceComplianceScriptRulesValidationError_OperandMissing,
		"operandtoolarge":                         DeviceComplianceScriptRulesValidationError_OperandTooLarge,
		"operatordatatypecombinationnotsupported": DeviceComplianceScriptRulesValidationError_OperatorDataTypeCombinationNotSupported,
		"operatormissing":                         DeviceComplianceScriptRulesValidationError_OperatorMissing,
		"operatornotsupported":                    DeviceComplianceScriptRulesValidationError_OperatorNotSupported,
		"remediationstringsmissing":               DeviceComplianceScriptRulesValidationError_RemediationStringsMissing,
		"rulesmissing":                            DeviceComplianceScriptRulesValidationError_RulesMissing,
		"settingnameinvalid":                      DeviceComplianceScriptRulesValidationError_SettingNameInvalid,
		"settingnamemissing":                      DeviceComplianceScriptRulesValidationError_SettingNameMissing,
		"settingnametoolarge":                     DeviceComplianceScriptRulesValidationError_SettingNameTooLarge,
		"titleinvalid":                            DeviceComplianceScriptRulesValidationError_TitleInvalid,
		"titlemissing":                            DeviceComplianceScriptRulesValidationError_TitleMissing,
		"titletoolarge":                           DeviceComplianceScriptRulesValidationError_TitleTooLarge,
		"toomanyrulesspecified":                   DeviceComplianceScriptRulesValidationError_TooManyRulesSpecified,
		"unknown":                                 DeviceComplianceScriptRulesValidationError_Unknown,
		"unrecognizedlocale":                      DeviceComplianceScriptRulesValidationError_UnrecognizedLocale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRulesValidationError(input)
	return &out, nil
}
