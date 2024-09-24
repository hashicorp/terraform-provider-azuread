package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityUserFlowAttributeAssignment{}

type IdentityUserFlowAttributeAssignment struct {
	// The display name of the identityUserFlowAttribute within a user flow.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Determines whether the identityUserFlowAttribute is optional. true means the user doesn't have to provide a value.
	// false means the user can't complete sign-up without providing a value.
	IsOptional *bool `json:"isOptional,omitempty"`

	// Determines whether the identityUserFlowAttribute requires verification, and is only used for verifying the user's
	// phone number or email address.
	RequiresVerification *bool `json:"requiresVerification,omitempty"`

	// The user attribute that you want to add to your user flow.
	UserAttribute *IdentityUserFlowAttribute `json:"userAttribute,omitempty"`

	// The input options for the user flow attribute. Only applicable when the userInputType is radioSingleSelect,
	// dropdownSingleSelect, or checkboxMultiSelect.
	UserAttributeValues *[]UserAttributeValuesItem `json:"userAttributeValues,omitempty"`

	UserInputType *IdentityUserFlowAttributeInputType `json:"userInputType,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityUserFlowAttributeAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityUserFlowAttributeAssignment{}

func (s IdentityUserFlowAttributeAssignment) MarshalJSON() ([]byte, error) {
	type wrapper IdentityUserFlowAttributeAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityUserFlowAttributeAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityUserFlowAttributeAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityUserFlowAttributeAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityUserFlowAttributeAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IdentityUserFlowAttributeAssignment{}

func (s *IdentityUserFlowAttributeAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName          nullable.Type[string]               `json:"displayName,omitempty"`
		IsOptional           *bool                               `json:"isOptional,omitempty"`
		RequiresVerification *bool                               `json:"requiresVerification,omitempty"`
		UserAttributeValues  *[]UserAttributeValuesItem          `json:"userAttributeValues,omitempty"`
		UserInputType        *IdentityUserFlowAttributeInputType `json:"userInputType,omitempty"`
		Id                   *string                             `json:"id,omitempty"`
		ODataId              *string                             `json:"@odata.id,omitempty"`
		ODataType            *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.IsOptional = decoded.IsOptional
	s.RequiresVerification = decoded.RequiresVerification
	s.UserAttributeValues = decoded.UserAttributeValues
	s.UserInputType = decoded.UserInputType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IdentityUserFlowAttributeAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["userAttribute"]; ok {
		impl, err := UnmarshalIdentityUserFlowAttributeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'UserAttribute' for 'IdentityUserFlowAttributeAssignment': %+v", err)
		}
		s.UserAttribute = &impl
	}

	return nil
}
