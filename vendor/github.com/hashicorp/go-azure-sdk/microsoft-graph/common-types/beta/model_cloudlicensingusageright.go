package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudLicensingUsageRight{}

type CloudLicensingUsageRight struct {
	// Information about the services associated with the usageRight. Not nullable. Read-only. Supports $filter on the
	// planId property.
	Services *[]CloudLicensingService `json:"services,omitempty"`

	// Unique identifier (GUID) for the service SKU that is equal to the skuId property on the related subscribedSku object.
	// Read-only. Supports $filter.
	SkuId nullable.Type[string] `json:"skuId,omitempty"`

	// Unique SKU display name that is equal to the skuPartNumber on the related subscribedSku object; for example,
	// AAD_Premium. Read-only.
	SkuPartNumber nullable.Type[string] `json:"skuPartNumber,omitempty"`

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

func (s CloudLicensingUsageRight) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudLicensingUsageRight{}

func (s CloudLicensingUsageRight) MarshalJSON() ([]byte, error) {
	type wrapper CloudLicensingUsageRight
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudLicensingUsageRight: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudLicensingUsageRight: %+v", err)
	}

	delete(decoded, "services")
	delete(decoded, "skuId")
	delete(decoded, "skuPartNumber")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudLicensing.usageRight"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudLicensingUsageRight: %+v", err)
	}

	return encoded, nil
}
