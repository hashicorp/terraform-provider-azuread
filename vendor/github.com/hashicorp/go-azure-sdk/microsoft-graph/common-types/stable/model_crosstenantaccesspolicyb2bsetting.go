package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyB2BSetting interface {
	CrossTenantAccessPolicyB2BSetting() BaseCrossTenantAccessPolicyB2BSettingImpl
}

var _ CrossTenantAccessPolicyB2BSetting = BaseCrossTenantAccessPolicyB2BSettingImpl{}

type BaseCrossTenantAccessPolicyB2BSettingImpl struct {
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

func (s BaseCrossTenantAccessPolicyB2BSettingImpl) CrossTenantAccessPolicyB2BSetting() BaseCrossTenantAccessPolicyB2BSettingImpl {
	return s
}

var _ CrossTenantAccessPolicyB2BSetting = RawCrossTenantAccessPolicyB2BSettingImpl{}

// RawCrossTenantAccessPolicyB2BSettingImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawCrossTenantAccessPolicyB2BSettingImpl struct {
	crossTenantAccessPolicyB2BSetting BaseCrossTenantAccessPolicyB2BSettingImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawCrossTenantAccessPolicyB2BSettingImpl) CrossTenantAccessPolicyB2BSetting() BaseCrossTenantAccessPolicyB2BSettingImpl {
	return s.crossTenantAccessPolicyB2BSetting
}

func (s RawCrossTenantAccessPolicyB2BSettingImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(input []byte) (CrossTenantAccessPolicyB2BSetting, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CrossTenantAccessPolicyB2BSetting into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.crossTenantAccessPolicyTenantRestrictions") {
		var out CrossTenantAccessPolicyTenantRestrictions
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CrossTenantAccessPolicyTenantRestrictions: %+v", err)
		}
		return out, nil
	}

	var parent BaseCrossTenantAccessPolicyB2BSettingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCrossTenantAccessPolicyB2BSettingImpl: %+v", err)
	}

	return RawCrossTenantAccessPolicyB2BSettingImpl{
		crossTenantAccessPolicyB2BSetting: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
