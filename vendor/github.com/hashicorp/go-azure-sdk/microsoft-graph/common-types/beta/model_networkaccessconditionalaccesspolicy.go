package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessConditionalAccessPolicy{}

type NetworkaccessConditionalAccessPolicy struct {
	// Indicates the date and time the conditional access policy was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Provides a summary of the conditional access policy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Represents the human-readable name or title assigned to the conditional access policy.
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates the date and time when the conditional access policy was last modified.
	ModifiedDateTime *string `json:"modifiedDateTime,omitempty"`

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

func (s NetworkaccessConditionalAccessPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessConditionalAccessPolicy{}

func (s NetworkaccessConditionalAccessPolicy) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessConditionalAccessPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessConditionalAccessPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessConditionalAccessPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.conditionalAccessPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessConditionalAccessPolicy: %+v", err)
	}

	return encoded, nil
}
