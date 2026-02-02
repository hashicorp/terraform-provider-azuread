package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomClaimAttributeBase = SourcedAttribute{}

type SourcedAttribute struct {
	// The identifier of the attribute on the specified source.
	Id *string `json:"id,omitempty"`

	// A flag that indicates if the name specified is that of an extension attribute.
	IsExtensionAttribute *bool `json:"isExtensionAttribute,omitempty"`

	// The source where the claim is going to retrieve its value. Valid sources include user, application, resource,
	// audience and company.
	Source nullable.Type[string] `json:"source,omitempty"`

	// Fields inherited from CustomClaimAttributeBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SourcedAttribute) CustomClaimAttributeBase() BaseCustomClaimAttributeBaseImpl {
	return BaseCustomClaimAttributeBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SourcedAttribute{}

func (s SourcedAttribute) MarshalJSON() ([]byte, error) {
	type wrapper SourcedAttribute
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SourcedAttribute: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SourcedAttribute: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sourcedAttribute"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SourcedAttribute: %+v", err)
	}

	return encoded, nil
}
