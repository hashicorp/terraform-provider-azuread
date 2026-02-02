package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSupportedRegionRestrictionDetail struct {
	// Indicates that the region is restricted for Cloud PC CPU provisioning. True indicates that Cloud PC provisioning with
	// CPU isn't available in this region. false indicates that it's available. The default value is false. Read-only.
	CPURestricted nullable.Type[bool] `json:"cPURestricted,omitempty"`

	// Indicates that the region is restricted for Cloud PC GPU provisioning. True indicates that Cloud PC provisioning with
	// GPU isn't available in this region. false indicates that it's available. The default value is false. Read-only.
	GPURestricted nullable.Type[bool] `json:"gPURestricted,omitempty"`

	// Indicates that the region is restricted for Cloud PC nested virtualization provisioning. True indicates that Cloud PC
	// provisioning with nested virtualization isn't available in this region; false indicates that it's available. The
	// default value is false. Read-only.
	NestedVirtualizationRestricted nullable.Type[bool] `json:"nestedVirtualizationRestricted,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = CloudPCSupportedRegionRestrictionDetail{}

func (s CloudPCSupportedRegionRestrictionDetail) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCSupportedRegionRestrictionDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCSupportedRegionRestrictionDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCSupportedRegionRestrictionDetail: %+v", err)
	}

	delete(decoded, "cPURestricted")
	delete(decoded, "gPURestricted")
	delete(decoded, "nestedVirtualizationRestricted")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCSupportedRegionRestrictionDetail: %+v", err)
	}

	return encoded, nil
}
