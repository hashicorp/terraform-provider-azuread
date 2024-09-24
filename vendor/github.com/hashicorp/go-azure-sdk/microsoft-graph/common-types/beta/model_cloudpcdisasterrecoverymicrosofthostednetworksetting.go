package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudPCDisasterRecoveryNetworkSetting = CloudPCDisasterRecoveryMicrosoftHostedNetworkSetting{}

type CloudPCDisasterRecoveryMicrosoftHostedNetworkSetting struct {
	RegionGroup *CloudPCRegionGroup `json:"regionGroup,omitempty"`
	RegionName  *string             `json:"regionName,omitempty"`

	// Fields inherited from CloudPCDisasterRecoveryNetworkSetting

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CloudPCDisasterRecoveryMicrosoftHostedNetworkSetting) CloudPCDisasterRecoveryNetworkSetting() BaseCloudPCDisasterRecoveryNetworkSettingImpl {
	return BaseCloudPCDisasterRecoveryNetworkSettingImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCDisasterRecoveryMicrosoftHostedNetworkSetting{}

func (s CloudPCDisasterRecoveryMicrosoftHostedNetworkSetting) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCDisasterRecoveryMicrosoftHostedNetworkSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCDisasterRecoveryMicrosoftHostedNetworkSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCDisasterRecoveryMicrosoftHostedNetworkSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcDisasterRecoveryMicrosoftHostedNetworkSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCDisasterRecoveryMicrosoftHostedNetworkSetting: %+v", err)
	}

	return encoded, nil
}
