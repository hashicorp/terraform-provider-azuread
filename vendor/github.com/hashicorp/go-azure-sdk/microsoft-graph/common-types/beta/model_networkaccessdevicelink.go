package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessDeviceLink{}

type NetworkaccessDeviceLink struct {
	// Determines the maximum allowed Mbps (megabits per second) bandwidth from a device link. The possible values
	// are:250,500,750,1000.
	BandwidthCapacityInMbps *NetworkaccessBandwidthCapacityInMbps `json:"bandwidthCapacityInMbps,omitempty"`

	BgpConfiguration *NetworkaccessBgpConfiguration `json:"bgpConfiguration,omitempty"`
	DeviceVendor     *NetworkaccessDeviceVendor     `json:"deviceVendor,omitempty"`

	// The public IP address of your CPE (customer premise equipment) device.
	IPAddress *string `json:"ipAddress,omitempty"`

	// last modified time.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Name.
	Name *string `json:"name,omitempty"`

	RedundancyConfiguration *NetworkaccessRedundancyConfiguration `json:"redundancyConfiguration,omitempty"`
	TunnelConfiguration     NetworkaccessTunnelConfiguration      `json:"tunnelConfiguration"`

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

func (s NetworkaccessDeviceLink) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessDeviceLink{}

func (s NetworkaccessDeviceLink) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessDeviceLink
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessDeviceLink: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessDeviceLink: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.deviceLink"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessDeviceLink: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &NetworkaccessDeviceLink{}

func (s *NetworkaccessDeviceLink) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		BandwidthCapacityInMbps *NetworkaccessBandwidthCapacityInMbps `json:"bandwidthCapacityInMbps,omitempty"`
		BgpConfiguration        *NetworkaccessBgpConfiguration        `json:"bgpConfiguration,omitempty"`
		DeviceVendor            *NetworkaccessDeviceVendor            `json:"deviceVendor,omitempty"`
		IPAddress               *string                               `json:"ipAddress,omitempty"`
		LastModifiedDateTime    *string                               `json:"lastModifiedDateTime,omitempty"`
		Name                    *string                               `json:"name,omitempty"`
		RedundancyConfiguration *NetworkaccessRedundancyConfiguration `json:"redundancyConfiguration,omitempty"`
		Id                      *string                               `json:"id,omitempty"`
		ODataId                 *string                               `json:"@odata.id,omitempty"`
		ODataType               *string                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.BandwidthCapacityInMbps = decoded.BandwidthCapacityInMbps
	s.BgpConfiguration = decoded.BgpConfiguration
	s.DeviceVendor = decoded.DeviceVendor
	s.IPAddress = decoded.IPAddress
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Name = decoded.Name
	s.RedundancyConfiguration = decoded.RedundancyConfiguration
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NetworkaccessDeviceLink into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["tunnelConfiguration"]; ok {
		impl, err := UnmarshalNetworkaccessTunnelConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'TunnelConfiguration' for 'NetworkaccessDeviceLink': %+v", err)
		}
		s.TunnelConfiguration = impl
	}

	return nil
}
