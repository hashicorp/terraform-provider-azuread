package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceAttribute struct {
	// Information about how to set the attribute, currently a accessPackageUserDirectoryAttributeStore object type.
	AttributeDestination AccessPackageResourceAttributeDestination `json:"attributeDestination"`

	// The name of the attribute in the end system. If the destination is accessPackageUserDirectoryAttributeStore, then a
	// user property such as jobTitle or a directory schema extension for the user object type, such as
	// extension2b676109c7c74ae2b41549205f1947edpersonalTitle.
	AttributeName nullable.Type[string] `json:"attributeName,omitempty"`

	// Information about how to populate the attribute value when an accessPackageAssignmentRequest is being fulfilled,
	// currently a accessPackageResourceAttributeQuestion object type.
	AttributeSource AccessPackageResourceAttributeSource `json:"attributeSource"`

	// Unique identifier for the attribute on the access package resource. Read-only.
	Id nullable.Type[string] `json:"id,omitempty"`

	// Specifies whether or not an existing attribute value can be edited by the requester.
	IsEditable nullable.Type[bool] `json:"isEditable,omitempty"`

	// Specifies whether the attribute will remain in the end system after an assignment ends.
	IsPersistedOnAssignmentRemoval nullable.Type[bool] `json:"isPersistedOnAssignmentRemoval,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = AccessPackageResourceAttribute{}

func (s AccessPackageResourceAttribute) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageResourceAttribute
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageResourceAttribute: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageResourceAttribute: %+v", err)
	}

	delete(decoded, "id")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageResourceAttribute: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessPackageResourceAttribute{}

func (s *AccessPackageResourceAttribute) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AttributeName                  nullable.Type[string] `json:"attributeName,omitempty"`
		Id                             nullable.Type[string] `json:"id,omitempty"`
		IsEditable                     nullable.Type[bool]   `json:"isEditable,omitempty"`
		IsPersistedOnAssignmentRemoval nullable.Type[bool]   `json:"isPersistedOnAssignmentRemoval,omitempty"`
		ODataId                        *string               `json:"@odata.id,omitempty"`
		ODataType                      *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AttributeName = decoded.AttributeName
	s.Id = decoded.Id
	s.IsEditable = decoded.IsEditable
	s.IsPersistedOnAssignmentRemoval = decoded.IsPersistedOnAssignmentRemoval
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageResourceAttribute into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attributeDestination"]; ok {
		impl, err := UnmarshalAccessPackageResourceAttributeDestinationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AttributeDestination' for 'AccessPackageResourceAttribute': %+v", err)
		}
		s.AttributeDestination = impl
	}

	if v, ok := temp["attributeSource"]; ok {
		impl, err := UnmarshalAccessPackageResourceAttributeSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AttributeSource' for 'AccessPackageResourceAttribute': %+v", err)
		}
		s.AttributeSource = impl
	}

	return nil
}
