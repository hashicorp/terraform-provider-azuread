package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleDataType string

const (
	DeviceComplianceScriptRuleDataType_Base64        DeviceComplianceScriptRuleDataType = "base64"
	DeviceComplianceScriptRuleDataType_Boolean       DeviceComplianceScriptRuleDataType = "boolean"
	DeviceComplianceScriptRuleDataType_BooleanArray  DeviceComplianceScriptRuleDataType = "booleanArray"
	DeviceComplianceScriptRuleDataType_DateTime      DeviceComplianceScriptRuleDataType = "dateTime"
	DeviceComplianceScriptRuleDataType_DateTimeArray DeviceComplianceScriptRuleDataType = "dateTimeArray"
	DeviceComplianceScriptRuleDataType_Double        DeviceComplianceScriptRuleDataType = "double"
	DeviceComplianceScriptRuleDataType_DoubleArray   DeviceComplianceScriptRuleDataType = "doubleArray"
	DeviceComplianceScriptRuleDataType_Int64         DeviceComplianceScriptRuleDataType = "int64"
	DeviceComplianceScriptRuleDataType_Int64Array    DeviceComplianceScriptRuleDataType = "int64Array"
	DeviceComplianceScriptRuleDataType_None          DeviceComplianceScriptRuleDataType = "none"
	DeviceComplianceScriptRuleDataType_String        DeviceComplianceScriptRuleDataType = "string"
	DeviceComplianceScriptRuleDataType_StringArray   DeviceComplianceScriptRuleDataType = "stringArray"
	DeviceComplianceScriptRuleDataType_Version       DeviceComplianceScriptRuleDataType = "version"
	DeviceComplianceScriptRuleDataType_VersionArray  DeviceComplianceScriptRuleDataType = "versionArray"
	DeviceComplianceScriptRuleDataType_Xml           DeviceComplianceScriptRuleDataType = "xml"
)

func PossibleValuesForDeviceComplianceScriptRuleDataType() []string {
	return []string{
		string(DeviceComplianceScriptRuleDataType_Base64),
		string(DeviceComplianceScriptRuleDataType_Boolean),
		string(DeviceComplianceScriptRuleDataType_BooleanArray),
		string(DeviceComplianceScriptRuleDataType_DateTime),
		string(DeviceComplianceScriptRuleDataType_DateTimeArray),
		string(DeviceComplianceScriptRuleDataType_Double),
		string(DeviceComplianceScriptRuleDataType_DoubleArray),
		string(DeviceComplianceScriptRuleDataType_Int64),
		string(DeviceComplianceScriptRuleDataType_Int64Array),
		string(DeviceComplianceScriptRuleDataType_None),
		string(DeviceComplianceScriptRuleDataType_String),
		string(DeviceComplianceScriptRuleDataType_StringArray),
		string(DeviceComplianceScriptRuleDataType_Version),
		string(DeviceComplianceScriptRuleDataType_VersionArray),
		string(DeviceComplianceScriptRuleDataType_Xml),
	}
}

func (s *DeviceComplianceScriptRuleDataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptRuleDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptRuleDataType(input string) (*DeviceComplianceScriptRuleDataType, error) {
	vals := map[string]DeviceComplianceScriptRuleDataType{
		"base64":        DeviceComplianceScriptRuleDataType_Base64,
		"boolean":       DeviceComplianceScriptRuleDataType_Boolean,
		"booleanarray":  DeviceComplianceScriptRuleDataType_BooleanArray,
		"datetime":      DeviceComplianceScriptRuleDataType_DateTime,
		"datetimearray": DeviceComplianceScriptRuleDataType_DateTimeArray,
		"double":        DeviceComplianceScriptRuleDataType_Double,
		"doublearray":   DeviceComplianceScriptRuleDataType_DoubleArray,
		"int64":         DeviceComplianceScriptRuleDataType_Int64,
		"int64array":    DeviceComplianceScriptRuleDataType_Int64Array,
		"none":          DeviceComplianceScriptRuleDataType_None,
		"string":        DeviceComplianceScriptRuleDataType_String,
		"stringarray":   DeviceComplianceScriptRuleDataType_StringArray,
		"version":       DeviceComplianceScriptRuleDataType_Version,
		"versionarray":  DeviceComplianceScriptRuleDataType_VersionArray,
		"xml":           DeviceComplianceScriptRuleDataType_Xml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleDataType(input)
	return &out, nil
}
