package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessRemoteNetwork{}

type NetworkaccessRemoteNetwork struct {
	// Specifies the connectivity details of all device links associated with a remote network.
	ConnectivityConfiguration *NetworkaccessRemoteNetworkConnectivityConfiguration `json:"connectivityConfiguration,omitempty"`

	// Each unique CPE device associated with a remote network is specified. Supports $expand.
	DeviceLinks *[]NetworkaccessDeviceLink `json:"deviceLinks,omitempty"`

	// Each forwarding profile associated with a remote network is specified. Supports $expand and $select.
	ForwardingProfiles *[]NetworkaccessForwardingProfile `json:"forwardingProfiles,omitempty"`

	// last modified time.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Name.
	Name *string `json:"name,omitempty"`

	Region *NetworkaccessRegion `json:"region,omitempty"`

	// Remote network version.
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

func (s NetworkaccessRemoteNetwork) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessRemoteNetwork{}

func (s NetworkaccessRemoteNetwork) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessRemoteNetwork
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessRemoteNetwork: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessRemoteNetwork: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.remoteNetwork"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessRemoteNetwork: %+v", err)
	}

	return encoded, nil
}
