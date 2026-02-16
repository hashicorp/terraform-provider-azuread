package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCCrossRegionDisasterRecoverySetting struct {
	// True if an end user is allowed to set up cross-region disaster recovery for Cloud PC; otherwise, false. The default
	// value is false. This property is deprecated and will no longer be supported effective February 11, 2025. For
	// scenarios where crossRegionDisasterRecoveryEnabled is true, set disasterRecoveryType to crossRegion. For scenarios
	// where crossRegionDisasterRecoveryEnabled is false, set disasterRecoveryType to notconfigured.
	CrossRegionDisasterRecoveryEnabled *bool `json:"crossRegionDisasterRecoveryEnabled,omitempty"`

	// Indicates the network settings of the Cloud PC during a cross-region disaster recovery operation.
	DisasterRecoveryNetworkSetting CloudPCDisasterRecoveryNetworkSetting `json:"disasterRecoveryNetworkSetting"`

	// Indicates the type of disaster recovery to perform when a disaster occurs on the user's Cloud PC. The possible values
	// are: notConfigured, crossRegion, premium, unknownFutureValue. The default value is notConfigured.
	DisasterRecoveryType *CloudPCDisasterRecoveryType `json:"disasterRecoveryType,omitempty"`

	// Indicates whether Windows 365 maintain the cross-region disaster recovery function generated restore points. If true,
	// the Windows 365 stored restore points; false indicates that Windows 365 doesn't generate or keep the restore point
	// from the original Cloud PC. If a disaster occurs, the new Cloud PC can only be provisioned using the initial image.
	// This limitation can result in the loss of some user data on the original Cloud PC. The default value is false.
	MaintainCrossRegionRestorePointEnabled nullable.Type[bool] `json:"maintainCrossRegionRestorePointEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether the client allows the end user to initiate a disaster recovery activation. True indicates that the
	// client includes the option for the end user to activate Backup Cloud PC. When false, the end user doesn't have the
	// option to activate disaster recovery. The default value is false. Currently, only premium disaster recovery is
	// supported.
	UserInitiatedDisasterRecoveryAllowed nullable.Type[bool] `json:"userInitiatedDisasterRecoveryAllowed,omitempty"`
}

var _ json.Unmarshaler = &CloudPCCrossRegionDisasterRecoverySetting{}

func (s *CloudPCCrossRegionDisasterRecoverySetting) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CrossRegionDisasterRecoveryEnabled     *bool                        `json:"crossRegionDisasterRecoveryEnabled,omitempty"`
		DisasterRecoveryType                   *CloudPCDisasterRecoveryType `json:"disasterRecoveryType,omitempty"`
		MaintainCrossRegionRestorePointEnabled nullable.Type[bool]          `json:"maintainCrossRegionRestorePointEnabled,omitempty"`
		ODataId                                *string                      `json:"@odata.id,omitempty"`
		ODataType                              *string                      `json:"@odata.type,omitempty"`
		UserInitiatedDisasterRecoveryAllowed   nullable.Type[bool]          `json:"userInitiatedDisasterRecoveryAllowed,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CrossRegionDisasterRecoveryEnabled = decoded.CrossRegionDisasterRecoveryEnabled
	s.DisasterRecoveryType = decoded.DisasterRecoveryType
	s.MaintainCrossRegionRestorePointEnabled = decoded.MaintainCrossRegionRestorePointEnabled
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.UserInitiatedDisasterRecoveryAllowed = decoded.UserInitiatedDisasterRecoveryAllowed

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
