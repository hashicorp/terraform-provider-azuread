package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PropertyType string

const (
	PropertyType_Boolean            PropertyType = "boolean"
	PropertyType_DateTime           PropertyType = "dateTime"
	PropertyType_DateTimeCollection PropertyType = "dateTimeCollection"
	PropertyType_Double             PropertyType = "double"
	PropertyType_DoubleCollection   PropertyType = "doubleCollection"
	PropertyType_Int64              PropertyType = "int64"
	PropertyType_Int64Collection    PropertyType = "int64Collection"
	PropertyType_String             PropertyType = "string"
	PropertyType_StringCollection   PropertyType = "stringCollection"
)

func PossibleValuesForPropertyType() []string {
	return []string{
		string(PropertyType_Boolean),
		string(PropertyType_DateTime),
		string(PropertyType_DateTimeCollection),
		string(PropertyType_Double),
		string(PropertyType_DoubleCollection),
		string(PropertyType_Int64),
		string(PropertyType_Int64Collection),
		string(PropertyType_String),
		string(PropertyType_StringCollection),
	}
}

func (s *PropertyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePropertyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePropertyType(input string) (*PropertyType, error) {
	vals := map[string]PropertyType{
		"boolean":            PropertyType_Boolean,
		"datetime":           PropertyType_DateTime,
		"datetimecollection": PropertyType_DateTimeCollection,
		"double":             PropertyType_Double,
		"doublecollection":   PropertyType_DoubleCollection,
		"int64":              PropertyType_Int64,
		"int64collection":    PropertyType_Int64Collection,
		"string":             PropertyType_String,
		"stringcollection":   PropertyType_StringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PropertyType(input)
	return &out, nil
}
