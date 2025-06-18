package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCServicePlan{}

type CloudPCServicePlan struct {
	// The name for the service plan. Read-only.
	DisplayName *string `json:"displayName,omitempty"`

	// Specifies the type of license used when provisioning Cloud PCs. By default, the license type is dedicated. Possible
	// values are: dedicated, shared, unknownFutureValue, sharedByUser, sharedByEntraGroup. Use the Prefer:
	// include-unknown-enum-members request header to get the following values from this evolvable enum: sharedByUser,
	// sharedByEntraGroup. The shared member is deprecated and will stop returning on April 30, 2027; going forward, use the
	// sharedByUser member.
	ProvisioningType *CloudPCProvisioningType `json:"provisioningType,omitempty"`

	// The size of the RAM in GB. Read-only.
	RamInGB *int64 `json:"ramInGB,omitempty"`

	// The size of the OS Disk in GB. Read-only.
	StorageInGB *int64 `json:"storageInGB,omitempty"`

	SupportedSolution *CloudPCManagementService `json:"supportedSolution,omitempty"`

	// The type of the service plan. Possible values are: enterprise, business, unknownFutureValue. Read-only.
	Type *CloudPCServicePlanType `json:"type,omitempty"`

	// The size of the user profile disk in GB. Read-only.
	UserProfileInGB *int64 `json:"userProfileInGB,omitempty"`

	// The number of vCPUs. Read-only.
	VCpuCount *int64 `json:"vCpuCount,omitempty"`

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

func (s CloudPCServicePlan) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCServicePlan{}

func (s CloudPCServicePlan) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCServicePlan
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCServicePlan: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCServicePlan: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "ramInGB")
	delete(decoded, "storageInGB")
	delete(decoded, "type")
	delete(decoded, "userProfileInGB")
	delete(decoded, "vCpuCount")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcServicePlan"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCServicePlan: %+v", err)
	}

	return encoded, nil
}
