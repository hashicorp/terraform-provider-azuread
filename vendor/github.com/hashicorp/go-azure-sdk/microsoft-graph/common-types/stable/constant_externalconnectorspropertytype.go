package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsPropertyType string

const (
	ExternalConnectorsPropertyType_Boolean            ExternalConnectorsPropertyType = "boolean"
	ExternalConnectorsPropertyType_DateTime           ExternalConnectorsPropertyType = "dateTime"
	ExternalConnectorsPropertyType_DateTimeCollection ExternalConnectorsPropertyType = "dateTimeCollection"
	ExternalConnectorsPropertyType_Double             ExternalConnectorsPropertyType = "double"
	ExternalConnectorsPropertyType_DoubleCollection   ExternalConnectorsPropertyType = "doubleCollection"
	ExternalConnectorsPropertyType_Int64              ExternalConnectorsPropertyType = "int64"
	ExternalConnectorsPropertyType_Int64Collection    ExternalConnectorsPropertyType = "int64Collection"
	ExternalConnectorsPropertyType_String             ExternalConnectorsPropertyType = "string"
	ExternalConnectorsPropertyType_StringCollection   ExternalConnectorsPropertyType = "stringCollection"
)

func PossibleValuesForExternalConnectorsPropertyType() []string {
	return []string{
		string(ExternalConnectorsPropertyType_Boolean),
		string(ExternalConnectorsPropertyType_DateTime),
		string(ExternalConnectorsPropertyType_DateTimeCollection),
		string(ExternalConnectorsPropertyType_Double),
		string(ExternalConnectorsPropertyType_DoubleCollection),
		string(ExternalConnectorsPropertyType_Int64),
		string(ExternalConnectorsPropertyType_Int64Collection),
		string(ExternalConnectorsPropertyType_String),
		string(ExternalConnectorsPropertyType_StringCollection),
	}
}

func (s *ExternalConnectorsPropertyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsPropertyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsPropertyType(input string) (*ExternalConnectorsPropertyType, error) {
	vals := map[string]ExternalConnectorsPropertyType{
		"boolean":            ExternalConnectorsPropertyType_Boolean,
		"datetime":           ExternalConnectorsPropertyType_DateTime,
		"datetimecollection": ExternalConnectorsPropertyType_DateTimeCollection,
		"double":             ExternalConnectorsPropertyType_Double,
		"doublecollection":   ExternalConnectorsPropertyType_DoubleCollection,
		"int64":              ExternalConnectorsPropertyType_Int64,
		"int64collection":    ExternalConnectorsPropertyType_Int64Collection,
		"string":             ExternalConnectorsPropertyType_String,
		"stringcollection":   ExternalConnectorsPropertyType_StringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsPropertyType(input)
	return &out, nil
}
