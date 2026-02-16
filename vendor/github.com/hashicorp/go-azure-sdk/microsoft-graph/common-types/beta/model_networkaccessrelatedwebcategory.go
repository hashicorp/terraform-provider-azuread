package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessRelatedResource = NetworkaccessRelatedWebCategory{}

type NetworkaccessRelatedWebCategory struct {
	WebCategoryName *string `json:"webCategoryName,omitempty"`

	// Fields inherited from NetworkaccessRelatedResource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NetworkaccessRelatedWebCategory) NetworkaccessRelatedResource() BaseNetworkaccessRelatedResourceImpl {
	return BaseNetworkaccessRelatedResourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessRelatedWebCategory{}

func (s NetworkaccessRelatedWebCategory) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessRelatedWebCategory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessRelatedWebCategory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessRelatedWebCategory: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.relatedWebCategory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessRelatedWebCategory: %+v", err)
	}

	return encoded, nil
}
