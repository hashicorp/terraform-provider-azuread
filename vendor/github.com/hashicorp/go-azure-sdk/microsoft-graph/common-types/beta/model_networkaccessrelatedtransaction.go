package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessRelatedResource = NetworkaccessRelatedTransaction{}

type NetworkaccessRelatedTransaction struct {
	TransactionId *string `json:"transactionId,omitempty"`

	// Fields inherited from NetworkaccessRelatedResource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NetworkaccessRelatedTransaction) NetworkaccessRelatedResource() BaseNetworkaccessRelatedResourceImpl {
	return BaseNetworkaccessRelatedResourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessRelatedTransaction{}

func (s NetworkaccessRelatedTransaction) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessRelatedTransaction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessRelatedTransaction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessRelatedTransaction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.relatedTransaction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessRelatedTransaction: %+v", err)
	}

	return encoded, nil
}
