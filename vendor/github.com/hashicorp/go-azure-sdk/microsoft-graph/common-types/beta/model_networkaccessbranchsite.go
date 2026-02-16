package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessBranchSite{}

type NetworkaccessBranchSite struct {
	// Determines the maximum allowed Mbps (megabits per second) bandwidth from a branch site. The possible values
	// are:250,500,750,1000.
	BandwidthCapacity nullable.Type[int64] `json:"bandwidthCapacity,omitempty"`

	// Specifies the connectivity details of all device links associated with a branch.
	ConnectivityConfiguration *NetworkaccessBranchConnectivityConfiguration `json:"connectivityConfiguration,omitempty"`

	// Determines the branch site status. The possible values are: pending, connected, inactive, error.
	ConnectivityState *NetworkaccessConnectivityState `json:"connectivityState,omitempty"`

	// The branch site is created in the specified country. DO NOT USE.
	Country nullable.Type[string] `json:"country,omitempty"`

	// Each unique CPE device associated with a branch is specified. Supports $expand.
	DeviceLinks *[]NetworkaccessDeviceLink `json:"deviceLinks,omitempty"`

	// Each forwarding profile associated with a branch site is specified. Supports $expand.
	ForwardingProfiles *[]NetworkaccessForwardingProfile `json:"forwardingProfiles,omitempty"`

	// last modified time.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Name.
	Name *string `json:"name,omitempty"`

	Region *NetworkaccessRegion `json:"region,omitempty"`

	// The branch version.
	Version *string `json:"version,omitempty"`

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

func (s NetworkaccessBranchSite) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessBranchSite{}

func (s NetworkaccessBranchSite) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessBranchSite
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessBranchSite: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessBranchSite: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.branchSite"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessBranchSite: %+v", err)
	}

	return encoded, nil
}
