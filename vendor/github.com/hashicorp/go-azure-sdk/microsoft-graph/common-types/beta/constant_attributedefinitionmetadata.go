package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeDefinitionMetadata string

const (
	AttributeDefinitionMetadata_BaseAttributeName       AttributeDefinitionMetadata = "BaseAttributeName"
	AttributeDefinitionMetadata_ComplexObjectDefinition AttributeDefinitionMetadata = "ComplexObjectDefinition"
	AttributeDefinitionMetadata_IsContainer             AttributeDefinitionMetadata = "IsContainer"
	AttributeDefinitionMetadata_IsCustomerDefined       AttributeDefinitionMetadata = "IsCustomerDefined"
	AttributeDefinitionMetadata_IsDomainQualified       AttributeDefinitionMetadata = "IsDomainQualified"
	AttributeDefinitionMetadata_LinkPropertyNames       AttributeDefinitionMetadata = "LinkPropertyNames"
	AttributeDefinitionMetadata_LinkTypeName            AttributeDefinitionMetadata = "LinkTypeName"
	AttributeDefinitionMetadata_MaximumLength           AttributeDefinitionMetadata = "MaximumLength"
	AttributeDefinitionMetadata_ReferencedProperty      AttributeDefinitionMetadata = "ReferencedProperty"
)

func PossibleValuesForAttributeDefinitionMetadata() []string {
	return []string{
		string(AttributeDefinitionMetadata_BaseAttributeName),
		string(AttributeDefinitionMetadata_ComplexObjectDefinition),
		string(AttributeDefinitionMetadata_IsContainer),
		string(AttributeDefinitionMetadata_IsCustomerDefined),
		string(AttributeDefinitionMetadata_IsDomainQualified),
		string(AttributeDefinitionMetadata_LinkPropertyNames),
		string(AttributeDefinitionMetadata_LinkTypeName),
		string(AttributeDefinitionMetadata_MaximumLength),
		string(AttributeDefinitionMetadata_ReferencedProperty),
	}
}

func (s *AttributeDefinitionMetadata) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeDefinitionMetadata(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeDefinitionMetadata(input string) (*AttributeDefinitionMetadata, error) {
	vals := map[string]AttributeDefinitionMetadata{
		"baseattributename":       AttributeDefinitionMetadata_BaseAttributeName,
		"complexobjectdefinition": AttributeDefinitionMetadata_ComplexObjectDefinition,
		"iscontainer":             AttributeDefinitionMetadata_IsContainer,
		"iscustomerdefined":       AttributeDefinitionMetadata_IsCustomerDefined,
		"isdomainqualified":       AttributeDefinitionMetadata_IsDomainQualified,
		"linkpropertynames":       AttributeDefinitionMetadata_LinkPropertyNames,
		"linktypename":            AttributeDefinitionMetadata_LinkTypeName,
		"maximumlength":           AttributeDefinitionMetadata_MaximumLength,
		"referencedproperty":      AttributeDefinitionMetadata_ReferencedProperty,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionMetadata(input)
	return &out, nil
}
