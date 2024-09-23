package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Code string

const (
	Code_DatatypeMissing                         Code = "datatypeMissing"
	Code_DatatypeNotSupported                    Code = "datatypeNotSupported"
	Code_DescriptionInvalid                      Code = "descriptionInvalid"
	Code_DescriptionMissing                      Code = "descriptionMissing"
	Code_DescriptionTooLarge                     Code = "descriptionTooLarge"
	Code_DuplicateLocales                        Code = "duplicateLocales"
	Code_DuplicateRules                          Code = "duplicateRules"
	Code_EnglishLocaleMissing                    Code = "englishLocaleMissing"
	Code_JsonFileInvalid                         Code = "jsonFileInvalid"
	Code_JsonFileMissing                         Code = "jsonFileMissing"
	Code_JsonFileTooLarge                        Code = "jsonFileTooLarge"
	Code_MoreInfoUriInvalid                      Code = "moreInfoUriInvalid"
	Code_MoreInfoUriMissing                      Code = "moreInfoUriMissing"
	Code_MoreInfoUriTooLarge                     Code = "moreInfoUriTooLarge"
	Code_None                                    Code = "none"
	Code_OperandInvalid                          Code = "operandInvalid"
	Code_OperandMissing                          Code = "operandMissing"
	Code_OperandTooLarge                         Code = "operandTooLarge"
	Code_OperatorDataTypeCombinationNotSupported Code = "operatorDataTypeCombinationNotSupported"
	Code_OperatorMissing                         Code = "operatorMissing"
	Code_OperatorNotSupported                    Code = "operatorNotSupported"
	Code_RemediationStringsMissing               Code = "remediationStringsMissing"
	Code_RulesMissing                            Code = "rulesMissing"
	Code_SettingNameInvalid                      Code = "settingNameInvalid"
	Code_SettingNameMissing                      Code = "settingNameMissing"
	Code_SettingNameTooLarge                     Code = "settingNameTooLarge"
	Code_TitleInvalid                            Code = "titleInvalid"
	Code_TitleMissing                            Code = "titleMissing"
	Code_TitleTooLarge                           Code = "titleTooLarge"
	Code_TooManyRulesSpecified                   Code = "tooManyRulesSpecified"
	Code_Unknown                                 Code = "unknown"
	Code_UnrecognizedLocale                      Code = "unrecognizedLocale"
)

func PossibleValuesForCode() []string {
	return []string{
		string(Code_DatatypeMissing),
		string(Code_DatatypeNotSupported),
		string(Code_DescriptionInvalid),
		string(Code_DescriptionMissing),
		string(Code_DescriptionTooLarge),
		string(Code_DuplicateLocales),
		string(Code_DuplicateRules),
		string(Code_EnglishLocaleMissing),
		string(Code_JsonFileInvalid),
		string(Code_JsonFileMissing),
		string(Code_JsonFileTooLarge),
		string(Code_MoreInfoUriInvalid),
		string(Code_MoreInfoUriMissing),
		string(Code_MoreInfoUriTooLarge),
		string(Code_None),
		string(Code_OperandInvalid),
		string(Code_OperandMissing),
		string(Code_OperandTooLarge),
		string(Code_OperatorDataTypeCombinationNotSupported),
		string(Code_OperatorMissing),
		string(Code_OperatorNotSupported),
		string(Code_RemediationStringsMissing),
		string(Code_RulesMissing),
		string(Code_SettingNameInvalid),
		string(Code_SettingNameMissing),
		string(Code_SettingNameTooLarge),
		string(Code_TitleInvalid),
		string(Code_TitleMissing),
		string(Code_TitleTooLarge),
		string(Code_TooManyRulesSpecified),
		string(Code_Unknown),
		string(Code_UnrecognizedLocale),
	}
}

func (s *Code) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCode(input string) (*Code, error) {
	vals := map[string]Code{
		"datatypemissing":                         Code_DatatypeMissing,
		"datatypenotsupported":                    Code_DatatypeNotSupported,
		"descriptioninvalid":                      Code_DescriptionInvalid,
		"descriptionmissing":                      Code_DescriptionMissing,
		"descriptiontoolarge":                     Code_DescriptionTooLarge,
		"duplicatelocales":                        Code_DuplicateLocales,
		"duplicaterules":                          Code_DuplicateRules,
		"englishlocalemissing":                    Code_EnglishLocaleMissing,
		"jsonfileinvalid":                         Code_JsonFileInvalid,
		"jsonfilemissing":                         Code_JsonFileMissing,
		"jsonfiletoolarge":                        Code_JsonFileTooLarge,
		"moreinfouriinvalid":                      Code_MoreInfoUriInvalid,
		"moreinfourimissing":                      Code_MoreInfoUriMissing,
		"moreinfouritoolarge":                     Code_MoreInfoUriTooLarge,
		"none":                                    Code_None,
		"operandinvalid":                          Code_OperandInvalid,
		"operandmissing":                          Code_OperandMissing,
		"operandtoolarge":                         Code_OperandTooLarge,
		"operatordatatypecombinationnotsupported": Code_OperatorDataTypeCombinationNotSupported,
		"operatormissing":                         Code_OperatorMissing,
		"operatornotsupported":                    Code_OperatorNotSupported,
		"remediationstringsmissing":               Code_RemediationStringsMissing,
		"rulesmissing":                            Code_RulesMissing,
		"settingnameinvalid":                      Code_SettingNameInvalid,
		"settingnamemissing":                      Code_SettingNameMissing,
		"settingnametoolarge":                     Code_SettingNameTooLarge,
		"titleinvalid":                            Code_TitleInvalid,
		"titlemissing":                            Code_TitleMissing,
		"titletoolarge":                           Code_TitleTooLarge,
		"toomanyrulesspecified":                   Code_TooManyRulesSpecified,
		"unknown":                                 Code_Unknown,
		"unrecognizedlocale":                      Code_UnrecognizedLocale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Code(input)
	return &out, nil
}
