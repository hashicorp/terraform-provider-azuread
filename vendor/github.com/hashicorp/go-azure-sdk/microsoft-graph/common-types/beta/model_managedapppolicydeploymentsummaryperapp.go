package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppPolicyDeploymentSummaryPerApp struct {
	// Number of users the policy is applied.
	ConfigurationAppliedUserCount *int64 `json:"configurationAppliedUserCount,omitempty"`

	// Deployment of an app.
	MobileAppIdentifier MobileAppIdentifier `json:"mobileAppIdentifier"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &ManagedAppPolicyDeploymentSummaryPerApp{}

func (s *ManagedAppPolicyDeploymentSummaryPerApp) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ConfigurationAppliedUserCount *int64  `json:"configurationAppliedUserCount,omitempty"`
		ODataId                       *string `json:"@odata.id,omitempty"`
		ODataType                     *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ConfigurationAppliedUserCount = decoded.ConfigurationAppliedUserCount
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ManagedAppPolicyDeploymentSummaryPerApp into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["mobileAppIdentifier"]; ok {
		impl, err := UnmarshalMobileAppIdentifierImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MobileAppIdentifier' for 'ManagedAppPolicyDeploymentSummaryPerApp': %+v", err)
		}
		s.MobileAppIdentifier = impl
	}

	return nil
}
