package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectDefinitionMetadata string

const (
	ObjectDefinitionMetadata_BaseObjectName               ObjectDefinitionMetadata = "BaseObjectName"
	ObjectDefinitionMetadata_ConnectorDataStorageRequired ObjectDefinitionMetadata = "ConnectorDataStorageRequired"
	ObjectDefinitionMetadata_Extensions                   ObjectDefinitionMetadata = "Extensions"
	ObjectDefinitionMetadata_IsSoftDeletionSupported      ObjectDefinitionMetadata = "IsSoftDeletionSupported"
	ObjectDefinitionMetadata_IsSynchronizeAllSupported    ObjectDefinitionMetadata = "IsSynchronizeAllSupported"
	ObjectDefinitionMetadata_PropertyNameAccountEnabled   ObjectDefinitionMetadata = "PropertyNameAccountEnabled"
	ObjectDefinitionMetadata_PropertyNameSoftDeleted      ObjectDefinitionMetadata = "PropertyNameSoftDeleted"
)

func PossibleValuesForObjectDefinitionMetadata() []string {
	return []string{
		string(ObjectDefinitionMetadata_BaseObjectName),
		string(ObjectDefinitionMetadata_ConnectorDataStorageRequired),
		string(ObjectDefinitionMetadata_Extensions),
		string(ObjectDefinitionMetadata_IsSoftDeletionSupported),
		string(ObjectDefinitionMetadata_IsSynchronizeAllSupported),
		string(ObjectDefinitionMetadata_PropertyNameAccountEnabled),
		string(ObjectDefinitionMetadata_PropertyNameSoftDeleted),
	}
}

func (s *ObjectDefinitionMetadata) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObjectDefinitionMetadata(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObjectDefinitionMetadata(input string) (*ObjectDefinitionMetadata, error) {
	vals := map[string]ObjectDefinitionMetadata{
		"baseobjectname":               ObjectDefinitionMetadata_BaseObjectName,
		"connectordatastoragerequired": ObjectDefinitionMetadata_ConnectorDataStorageRequired,
		"extensions":                   ObjectDefinitionMetadata_Extensions,
		"issoftdeletionsupported":      ObjectDefinitionMetadata_IsSoftDeletionSupported,
		"issynchronizeallsupported":    ObjectDefinitionMetadata_IsSynchronizeAllSupported,
		"propertynameaccountenabled":   ObjectDefinitionMetadata_PropertyNameAccountEnabled,
		"propertynamesoftdeleted":      ObjectDefinitionMetadata_PropertyNameSoftDeleted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectDefinitionMetadata(input)
	return &out, nil
}
