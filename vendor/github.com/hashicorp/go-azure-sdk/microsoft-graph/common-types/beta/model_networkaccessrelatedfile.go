package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessRelatedResource = NetworkaccessRelatedFile{}

type NetworkaccessRelatedFile struct {
	Directory   nullable.Type[string] `json:"directory,omitempty"`
	Name        nullable.Type[string] `json:"name,omitempty"`
	SizeInBytes nullable.Type[int64]  `json:"sizeInBytes,omitempty"`

	// Fields inherited from NetworkaccessRelatedResource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NetworkaccessRelatedFile) NetworkaccessRelatedResource() BaseNetworkaccessRelatedResourceImpl {
	return BaseNetworkaccessRelatedResourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessRelatedFile{}

func (s NetworkaccessRelatedFile) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessRelatedFile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessRelatedFile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessRelatedFile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.relatedFile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessRelatedFile: %+v", err)
	}

	return encoded, nil
}
