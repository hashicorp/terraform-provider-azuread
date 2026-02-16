package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCSupportedRegion{}

type CloudPCSupportedRegion struct {
	// The name for the supported region. Read-only.
	DisplayName *string `json:"displayName,omitempty"`

	RegionGroup *CloudPCRegionGroup `json:"regionGroup,omitempty"`

	// When the region isn't available, all region restrictions are set to true. These restrictions apply to three
	// properties: cPURestricted, gPURestricted, and nestedVirtualizationRestricted. cPURestricted indicates whether the
	// region is available for CPU, gPURestricted indicates whether the region is available for GPU, and
	// nestedVirtualizationRestricted indicates whether the region is available for nested virtualization. Read-only.
	RegionRestrictionDetail *CloudPCSupportedRegionRestrictionDetail `json:"regionRestrictionDetail,omitempty"`

	// The status of the supported region. Possible values are: available, restricted, unavailable, unknownFutureValue.
	// Read-only.
	RegionStatus *CloudPCSupportedRegionStatus `json:"regionStatus,omitempty"`

	SupportedSolution *CloudPCManagementService `json:"supportedSolution,omitempty"`

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

func (s CloudPCSupportedRegion) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCSupportedRegion{}

func (s CloudPCSupportedRegion) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCSupportedRegion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCSupportedRegion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCSupportedRegion: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "regionRestrictionDetail")
	delete(decoded, "regionStatus")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcSupportedRegion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCSupportedRegion: %+v", err)
	}

	return encoded, nil
}
