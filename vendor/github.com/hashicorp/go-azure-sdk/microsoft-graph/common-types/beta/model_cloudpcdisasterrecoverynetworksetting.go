package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDisasterRecoveryNetworkSetting interface {
	CloudPCDisasterRecoveryNetworkSetting() BaseCloudPCDisasterRecoveryNetworkSettingImpl
}

var _ CloudPCDisasterRecoveryNetworkSetting = BaseCloudPCDisasterRecoveryNetworkSettingImpl{}

type BaseCloudPCDisasterRecoveryNetworkSettingImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCloudPCDisasterRecoveryNetworkSettingImpl) CloudPCDisasterRecoveryNetworkSetting() BaseCloudPCDisasterRecoveryNetworkSettingImpl {
	return s
}

var _ CloudPCDisasterRecoveryNetworkSetting = RawCloudPCDisasterRecoveryNetworkSettingImpl{}

// RawCloudPCDisasterRecoveryNetworkSettingImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawCloudPCDisasterRecoveryNetworkSettingImpl struct {
	cloudPCDisasterRecoveryNetworkSetting BaseCloudPCDisasterRecoveryNetworkSettingImpl
	Type                                  string
	Values                                map[string]interface{}
}

func (s RawCloudPCDisasterRecoveryNetworkSettingImpl) CloudPCDisasterRecoveryNetworkSetting() BaseCloudPCDisasterRecoveryNetworkSettingImpl {
	return s.cloudPCDisasterRecoveryNetworkSetting
}

func (s RawCloudPCDisasterRecoveryNetworkSettingImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalCloudPCDisasterRecoveryNetworkSettingImplementation(input []byte) (CloudPCDisasterRecoveryNetworkSetting, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCDisasterRecoveryNetworkSetting into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcDisasterRecoveryAzureConnectionSetting") {
		var out CloudPCDisasterRecoveryAzureConnectionSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCDisasterRecoveryAzureConnectionSetting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcDisasterRecoveryMicrosoftHostedNetworkSetting") {
		var out CloudPCDisasterRecoveryMicrosoftHostedNetworkSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCDisasterRecoveryMicrosoftHostedNetworkSetting: %+v", err)
		}
		return out, nil
	}

	var parent BaseCloudPCDisasterRecoveryNetworkSettingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCloudPCDisasterRecoveryNetworkSettingImpl: %+v", err)
	}

	return RawCloudPCDisasterRecoveryNetworkSettingImpl{
		cloudPCDisasterRecoveryNetworkSetting: parent,
		Type:                                  value,
		Values:                                temp,
	}, nil

}
