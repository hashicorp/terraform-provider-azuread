package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CrossTenantAccessPolicyB2BSetting = CrossTenantAccessPolicyTenantRestrictions{}

type CrossTenantAccessPolicyTenantRestrictions struct {
	// Defines the rule for filtering devices and whether devices that satisfy the rule should be allowed or blocked. This
	// property isn't supported on the server side yet.
	Devices *DevicesFilter `json:"devices,omitempty"`

	// Fields inherited from CrossTenantAccessPolicyB2BSetting

	// The list of applications targeted with your cross-tenant access policy.
	Applications *CrossTenantAccessPolicyTargetConfiguration `json:"applications,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The list of users and groups targeted with your cross-tenant access policy.
	UsersAndGroups *CrossTenantAccessPolicyTargetConfiguration `json:"usersAndGroups,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CrossTenantAccessPolicyTenantRestrictions) CrossTenantAccessPolicyB2BSetting() BaseCrossTenantAccessPolicyB2BSettingImpl {
	return BaseCrossTenantAccessPolicyB2BSettingImpl{
		Applications:   s.Applications,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
		UsersAndGroups: s.UsersAndGroups,
	}
}

var _ json.Marshaler = CrossTenantAccessPolicyTenantRestrictions{}

func (s CrossTenantAccessPolicyTenantRestrictions) MarshalJSON() ([]byte, error) {
	type wrapper CrossTenantAccessPolicyTenantRestrictions
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CrossTenantAccessPolicyTenantRestrictions: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CrossTenantAccessPolicyTenantRestrictions: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.crossTenantAccessPolicyTenantRestrictions"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CrossTenantAccessPolicyTenantRestrictions: %+v", err)
	}

	return encoded, nil
}
