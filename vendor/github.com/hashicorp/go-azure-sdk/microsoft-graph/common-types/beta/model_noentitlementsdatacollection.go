package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EntitlementsDataCollectionInfo = NoEntitlementsDataCollection{}

type NoEntitlementsDataCollection struct {

	// Fields inherited from EntitlementsDataCollectionInfo

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NoEntitlementsDataCollection) EntitlementsDataCollectionInfo() BaseEntitlementsDataCollectionInfoImpl {
	return BaseEntitlementsDataCollectionInfoImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NoEntitlementsDataCollection{}

func (s NoEntitlementsDataCollection) MarshalJSON() ([]byte, error) {
	type wrapper NoEntitlementsDataCollection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NoEntitlementsDataCollection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NoEntitlementsDataCollection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.noEntitlementsDataCollection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NoEntitlementsDataCollection: %+v", err)
	}

	return encoded, nil
}
