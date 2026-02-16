package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EntitlementsDataCollectionInfo = EntitlementsDataCollection{}

type EntitlementsDataCollection struct {
	// Last transformation time of entitlements.
	LastCollectionDateTime *string `json:"lastCollectionDateTime,omitempty"`

	PermissionsModificationCapability *PermissionsModificationCapability `json:"permissionsModificationCapability,omitempty"`
	Status                            *DataCollectionStatus              `json:"status,omitempty"`

	// Fields inherited from EntitlementsDataCollectionInfo

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EntitlementsDataCollection) EntitlementsDataCollectionInfo() BaseEntitlementsDataCollectionInfoImpl {
	return BaseEntitlementsDataCollectionInfoImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EntitlementsDataCollection{}

func (s EntitlementsDataCollection) MarshalJSON() ([]byte, error) {
	type wrapper EntitlementsDataCollection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EntitlementsDataCollection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EntitlementsDataCollection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.entitlementsDataCollection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EntitlementsDataCollection: %+v", err)
	}

	return encoded, nil
}
