package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementParameterValueType string

const (
	ManagedTenantsManagementParameterValueType_Boolean           ManagedTenantsManagementParameterValueType = "boolean"
	ManagedTenantsManagementParameterValueType_BooleanCollection ManagedTenantsManagementParameterValueType = "booleanCollection"
	ManagedTenantsManagementParameterValueType_Guid              ManagedTenantsManagementParameterValueType = "guid"
	ManagedTenantsManagementParameterValueType_GuidCollection    ManagedTenantsManagementParameterValueType = "guidCollection"
	ManagedTenantsManagementParameterValueType_Integer           ManagedTenantsManagementParameterValueType = "integer"
	ManagedTenantsManagementParameterValueType_IntegerCollection ManagedTenantsManagementParameterValueType = "integerCollection"
	ManagedTenantsManagementParameterValueType_String            ManagedTenantsManagementParameterValueType = "string"
	ManagedTenantsManagementParameterValueType_StringCollection  ManagedTenantsManagementParameterValueType = "stringCollection"
)

func PossibleValuesForManagedTenantsManagementParameterValueType() []string {
	return []string{
		string(ManagedTenantsManagementParameterValueType_Boolean),
		string(ManagedTenantsManagementParameterValueType_BooleanCollection),
		string(ManagedTenantsManagementParameterValueType_Guid),
		string(ManagedTenantsManagementParameterValueType_GuidCollection),
		string(ManagedTenantsManagementParameterValueType_Integer),
		string(ManagedTenantsManagementParameterValueType_IntegerCollection),
		string(ManagedTenantsManagementParameterValueType_String),
		string(ManagedTenantsManagementParameterValueType_StringCollection),
	}
}

func (s *ManagedTenantsManagementParameterValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementParameterValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementParameterValueType(input string) (*ManagedTenantsManagementParameterValueType, error) {
	vals := map[string]ManagedTenantsManagementParameterValueType{
		"boolean":           ManagedTenantsManagementParameterValueType_Boolean,
		"booleancollection": ManagedTenantsManagementParameterValueType_BooleanCollection,
		"guid":              ManagedTenantsManagementParameterValueType_Guid,
		"guidcollection":    ManagedTenantsManagementParameterValueType_GuidCollection,
		"integer":           ManagedTenantsManagementParameterValueType_Integer,
		"integercollection": ManagedTenantsManagementParameterValueType_IntegerCollection,
		"string":            ManagedTenantsManagementParameterValueType_String,
		"stringcollection":  ManagedTenantsManagementParameterValueType_StringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementParameterValueType(input)
	return &out, nil
}
