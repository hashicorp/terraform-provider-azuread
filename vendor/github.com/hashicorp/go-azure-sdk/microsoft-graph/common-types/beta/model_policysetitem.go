package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySetItem interface {
	Entity
	PolicySetItem() BasePolicySetItemImpl
}

var _ PolicySetItem = BasePolicySetItemImpl{}

type BasePolicySetItemImpl struct {
	// Creation time of the PolicySetItem.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// DisplayName of the PolicySetItem.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	ErrorCode *ErrorCode `json:"errorCode,omitempty"`

	// Tags of the guided deployment
	GuidedDeploymentTags *[]string `json:"guidedDeploymentTags,omitempty"`

	// policySetType of the PolicySetItem.
	ItemType nullable.Type[string] `json:"itemType,omitempty"`

	// Last modified time of the PolicySetItem.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// PayloadId of the PolicySetItem.
	PayloadId *string `json:"payloadId,omitempty"`

	// The enum to specify the status of PolicySet.
	Status *PolicySetStatus `json:"status,omitempty"`

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

func (s BasePolicySetItemImpl) PolicySetItem() BasePolicySetItemImpl {
	return s
}

func (s BasePolicySetItemImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PolicySetItem = RawPolicySetItemImpl{}

// RawPolicySetItemImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPolicySetItemImpl struct {
	policySetItem BasePolicySetItemImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawPolicySetItemImpl) PolicySetItem() BasePolicySetItemImpl {
	return s.policySetItem
}

func (s RawPolicySetItemImpl) Entity() BaseEntityImpl {
	return s.policySetItem.Entity()
}

var _ json.Marshaler = BasePolicySetItemImpl{}

func (s BasePolicySetItemImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePolicySetItemImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePolicySetItemImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePolicySetItemImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.policySetItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePolicySetItemImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPolicySetItemImplementation(input []byte) (PolicySetItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PolicySetItem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceCompliancePolicyPolicySetItem") {
		var out DeviceCompliancePolicyPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCompliancePolicyPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceConfigurationPolicySetItem") {
		var out DeviceConfigurationPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceConfigurationPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationPolicyPolicySetItem") {
		var out DeviceManagementConfigurationPolicyPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationPolicyPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementScriptPolicySetItem") {
		var out DeviceManagementScriptPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementScriptPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enrollmentRestrictionsConfigurationPolicySetItem") {
		var out EnrollmentRestrictionsConfigurationPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnrollmentRestrictionsConfigurationPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosLobAppProvisioningConfigurationPolicySetItem") {
		var out IosLobAppProvisioningConfigurationPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosLobAppProvisioningConfigurationPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAppProtectionPolicySetItem") {
		var out ManagedAppProtectionPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAppProtectionPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedDeviceMobileAppConfigurationPolicySetItem") {
		var out ManagedDeviceMobileAppConfigurationPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mdmWindowsInformationProtectionPolicyPolicySetItem") {
		var out MdmWindowsInformationProtectionPolicyPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MdmWindowsInformationProtectionPolicyPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppPolicySetItem") {
		var out MobileAppPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetedManagedAppConfigurationPolicySetItem") {
		var out TargetedManagedAppConfigurationPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetedManagedAppConfigurationPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10EnrollmentCompletionPageConfigurationPolicySetItem") {
		var out Windows10EnrollmentCompletionPageConfigurationPolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10EnrollmentCompletionPageConfigurationPolicySetItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsAutopilotDeploymentProfilePolicySetItem") {
		var out WindowsAutopilotDeploymentProfilePolicySetItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAutopilotDeploymentProfilePolicySetItem: %+v", err)
		}
		return out, nil
	}

	var parent BasePolicySetItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePolicySetItemImpl: %+v", err)
	}

	return RawPolicySetItemImpl{
		policySetItem: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
