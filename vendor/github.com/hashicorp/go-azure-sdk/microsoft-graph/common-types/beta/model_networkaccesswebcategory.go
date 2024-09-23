package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessRuleDestination = NetworkaccessWebCategory{}

type NetworkaccessWebCategory struct {
	// The display name for the web category.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The group or category to which the web category belongs.
	Group nullable.Type[string] `json:"group,omitempty"`

	// The unique name that is associated with the web category.
	Name *string `json:"name,omitempty"`

	// Fields inherited from NetworkaccessRuleDestination

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NetworkaccessWebCategory) NetworkaccessRuleDestination() BaseNetworkaccessRuleDestinationImpl {
	return BaseNetworkaccessRuleDestinationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessWebCategory{}

func (s NetworkaccessWebCategory) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessWebCategory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessWebCategory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessWebCategory: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.webCategory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessWebCategory: %+v", err)
	}

	return encoded, nil
}
