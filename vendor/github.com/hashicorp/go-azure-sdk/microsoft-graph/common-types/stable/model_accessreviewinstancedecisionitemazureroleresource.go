package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessReviewInstanceDecisionItemResource = AccessReviewInstanceDecisionItemAzureRoleResource{}

type AccessReviewInstanceDecisionItemAzureRoleResource struct {
	// Details of the scope this role is associated with.
	Scope AccessReviewInstanceDecisionItemResource `json:"scope"`

	// Fields inherited from AccessReviewInstanceDecisionItemResource

	// Display name of the resource
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Identifier of the resource
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of resource. Types include: Group, ServicePrincipal, DirectoryRole, AzureRole, AccessPackageAssignmentPolicy.
	Type nullable.Type[string] `json:"type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessReviewInstanceDecisionItemAzureRoleResource) AccessReviewInstanceDecisionItemResource() BaseAccessReviewInstanceDecisionItemResourceImpl {
	return BaseAccessReviewInstanceDecisionItemResourceImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		Type:        s.Type,
	}
}

var _ json.Marshaler = AccessReviewInstanceDecisionItemAzureRoleResource{}

func (s AccessReviewInstanceDecisionItemAzureRoleResource) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewInstanceDecisionItemAzureRoleResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewInstanceDecisionItemAzureRoleResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewInstanceDecisionItemAzureRoleResource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewInstanceDecisionItemAzureRoleResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewInstanceDecisionItemAzureRoleResource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessReviewInstanceDecisionItemAzureRoleResource{}

func (s *AccessReviewInstanceDecisionItemAzureRoleResource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName nullable.Type[string] `json:"displayName,omitempty"`
		Id          nullable.Type[string] `json:"id,omitempty"`
		ODataId     *string               `json:"@odata.id,omitempty"`
		ODataType   *string               `json:"@odata.type,omitempty"`
		Type        nullable.Type[string] `json:"type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessReviewInstanceDecisionItemAzureRoleResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["scope"]; ok {
		impl, err := UnmarshalAccessReviewInstanceDecisionItemResourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Scope' for 'AccessReviewInstanceDecisionItemAzureRoleResource': %+v", err)
		}
		s.Scope = impl
	}

	return nil
}
