package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataType string

const (
	DataType_Base64        DataType = "base64"
	DataType_Boolean       DataType = "boolean"
	DataType_BooleanArray  DataType = "booleanArray"
	DataType_DateTime      DataType = "dateTime"
	DataType_DateTimeArray DataType = "dateTimeArray"
	DataType_Double        DataType = "double"
	DataType_DoubleArray   DataType = "doubleArray"
	DataType_Int64         DataType = "int64"
	DataType_Int64Array    DataType = "int64Array"
	DataType_None          DataType = "none"
	DataType_String        DataType = "string"
	DataType_StringArray   DataType = "stringArray"
	DataType_Version       DataType = "version"
	DataType_VersionArray  DataType = "versionArray"
	DataType_Xml           DataType = "xml"
)

func PossibleValuesForDataType() []string {
	return []string{
		string(DataType_Base64),
		string(DataType_Boolean),
		string(DataType_BooleanArray),
		string(DataType_DateTime),
		string(DataType_DateTimeArray),
		string(DataType_Double),
		string(DataType_DoubleArray),
		string(DataType_Int64),
		string(DataType_Int64Array),
		string(DataType_None),
		string(DataType_String),
		string(DataType_StringArray),
		string(DataType_Version),
		string(DataType_VersionArray),
		string(DataType_Xml),
	}
}

func (s *DataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataType(input string) (*DataType, error) {
	vals := map[string]DataType{
		"base64":        DataType_Base64,
		"boolean":       DataType_Boolean,
		"booleanarray":  DataType_BooleanArray,
		"datetime":      DataType_DateTime,
		"datetimearray": DataType_DateTimeArray,
		"double":        DataType_Double,
		"doublearray":   DataType_DoubleArray,
		"int64":         DataType_Int64,
		"int64array":    DataType_Int64Array,
		"none":          DataType_None,
		"string":        DataType_String,
		"stringarray":   DataType_StringArray,
		"version":       DataType_Version,
		"versionarray":  DataType_VersionArray,
		"xml":           DataType_Xml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataType(input)
	return &out, nil
}
