package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = LicenseDetails{}

type LicenseDetails struct {
	// Information about the service plans assigned with the license. Read-only. Not nullable.
	ServicePlans *[]ServicePlanInfo `json:"servicePlans,omitempty"`

	// Unique identifier (GUID) for the service SKU. Equal to the skuId property on the related subscribedSku object.
	// Read-only.
	SkuId nullable.Type[string] `json:"skuId,omitempty"`

	// Unique SKU display name. Equal to the skuPartNumber on the related subscribedSku object; for example, AAD_Premium.
	// Read-only.
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

func (s LicenseDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = LicenseDetails{}

func (s LicenseDetails) MarshalJSON() ([]byte, error) {
	type wrapper LicenseDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LicenseDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LicenseDetails: %+v", err)
	}

	delete(decoded, "servicePlans")
	delete(decoded, "skuId")
	delete(decoded, "skuPartNumber")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.licenseDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LicenseDetails: %+v", err)
	}

	return encoded, nil
}
