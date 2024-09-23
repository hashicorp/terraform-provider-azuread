package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCCrossRegionDisasterRecoverySetting struct {
	CrossRegionDisasterRecoveryEnabled     *bool                                 `json:"crossRegionDisasterRecoveryEnabled,omitempty"`
	DisasterRecoveryNetworkSetting         CloudPCDisasterRecoveryNetworkSetting `json:"disasterRecoveryNetworkSetting"`
	MaintainCrossRegionRestorePointEnabled nullable.Type[bool]                   `json:"maintainCrossRegionRestorePointEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &CloudPCCrossRegionDisasterRecoverySetting{}

func (s *CloudPCCrossRegionDisasterRecoverySetting) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CrossRegionDisasterRecoveryEnabled     *bool               `json:"crossRegionDisasterRecoveryEnabled,omitempty"`
		MaintainCrossRegionRestorePointEnabled nullable.Type[bool] `json:"maintainCrossRegionRestorePointEnabled,omitempty"`
		ODataId                                *string             `json:"@odata.id,omitempty"`
		ODataType                              *string             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CrossRegionDisasterRecoveryEnabled = decoded.CrossRegionDisasterRecoveryEnabled
	s.MaintainCrossRegionRestorePointEnabled = decoded.MaintainCrossRegionRestorePointEnabled
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CloudPCCrossRegionDisasterRecoverySetting into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["disasterRecoveryNetworkSetting"]; ok {
		impl, err := UnmarshalCloudPCDisasterRecoveryNetworkSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DisasterRecoveryNetworkSetting' for 'CloudPCCrossRegionDisasterRecoverySetting': %+v", err)
		}
		s.DisasterRecoveryNetworkSetting = impl
	}

	return nil
}
