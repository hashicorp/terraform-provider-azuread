package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryDefinitionDiscoverabilities string

const (
	DirectoryDefinitionDiscoverabilities_AttributeDataTypes  DirectoryDefinitionDiscoverabilities = "AttributeDataTypes"
	DirectoryDefinitionDiscoverabilities_AttributeNames      DirectoryDefinitionDiscoverabilities = "AttributeNames"
	DirectoryDefinitionDiscoverabilities_AttributeReadOnly   DirectoryDefinitionDiscoverabilities = "AttributeReadOnly"
	DirectoryDefinitionDiscoverabilities_None                DirectoryDefinitionDiscoverabilities = "None"
	DirectoryDefinitionDiscoverabilities_ReferenceAttributes DirectoryDefinitionDiscoverabilities = "ReferenceAttributes"
)

func PossibleValuesForDirectoryDefinitionDiscoverabilities() []string {
	return []string{
		string(DirectoryDefinitionDiscoverabilities_AttributeDataTypes),
		string(DirectoryDefinitionDiscoverabilities_AttributeNames),
		string(DirectoryDefinitionDiscoverabilities_AttributeReadOnly),
		string(DirectoryDefinitionDiscoverabilities_None),
		string(DirectoryDefinitionDiscoverabilities_ReferenceAttributes),
	}
}

func (s *DirectoryDefinitionDiscoverabilities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDirectoryDefinitionDiscoverabilities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDirectoryDefinitionDiscoverabilities(input string) (*DirectoryDefinitionDiscoverabilities, error) {
	vals := map[string]DirectoryDefinitionDiscoverabilities{
		"attributedatatypes":  DirectoryDefinitionDiscoverabilities_AttributeDataTypes,
		"attributenames":      DirectoryDefinitionDiscoverabilities_AttributeNames,
		"attributereadonly":   DirectoryDefinitionDiscoverabilities_AttributeReadOnly,
		"none":                DirectoryDefinitionDiscoverabilities_None,
		"referenceattributes": DirectoryDefinitionDiscoverabilities_ReferenceAttributes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DirectoryDefinitionDiscoverabilities(input)
	return &out, nil
}
