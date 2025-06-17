package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudPCDisasterRecoveryNetworkSetting = CloudPCDisasterRecoveryAzureConnectionSetting{}

type CloudPCDisasterRecoveryAzureConnectionSetting struct {
	// Indicates the unique ID of the virtual network that the new Cloud PC joins.
	OnPremisesConnectionId *string `json:"onPremisesConnectionId,omitempty"`

	// Fields inherited from CloudPCDisasterRecoveryNetworkSetting

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CloudPCDisasterRecoveryAzureConnectionSetting) CloudPCDisasterRecoveryNetworkSetting() BaseCloudPCDisasterRecoveryNetworkSettingImpl {
	return BaseCloudPCDisasterRecoveryNetworkSettingImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCDisasterRecoveryAzureConnectionSetting{}

func (s CloudPCDisasterRecoveryAzureConnectionSetting) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCDisasterRecoveryAzureConnectionSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCDisasterRecoveryAzureConnectionSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCDisasterRecoveryAzureConnectionSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcDisasterRecoveryAzureConnectionSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCDisasterRecoveryAzureConnectionSetting: %+v", err)
	}

	return encoded, nil
}
