package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsCloudPCDevice{}

type ManagedTenantsCloudPCDevice struct {
	// The status of the cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod,
	// deprovisioning, failed. Required. Read-only.
	CloudPCStatus nullable.Type[string] `json:"cloudPcStatus,omitempty"`

	// The specification of the cloud PC device. Required. Read-only.
	DeviceSpecification nullable.Type[string] `json:"deviceSpecification,omitempty"`

	// The display name of the cloud PC device. Required. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Date and time the entity was last updated in the multi-tenant management platform. Required. Read-only.
	LastRefreshedDateTime nullable.Type[string] `json:"lastRefreshedDateTime,omitempty"`

	// The managed device identifier of the cloud PC device. Optional. Read-only.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// The managed device display name of the cloud PC device. Optional. Read-only.
	ManagedDeviceName nullable.Type[string] `json:"managedDeviceName,omitempty"`

	// The provisioning policy identifier for the cloud PC device. Required. Read-only.
	ProvisioningPolicyId nullable.Type[string] `json:"provisioningPolicyId,omitempty"`

	// The service plan name of the cloud PC device. Required. Read-only.
	ServicePlanName nullable.Type[string] `json:"servicePlanName,omitempty"`

	// The service plan type of the cloud PC device. Required. Read-only.
	ServicePlanType nullable.Type[string] `json:"servicePlanType,omitempty"`

	// The display name for the managed tenant. Required. Read-only.
	TenantDisplayName nullable.Type[string] `json:"tenantDisplayName,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant. Required. Read-only.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// The user principal name (UPN) of the user assigned to the cloud PC device. Required. Read-only.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s ManagedTenantsCloudPCDevice) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsCloudPCDevice{}

func (s ManagedTenantsCloudPCDevice) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsCloudPCDevice
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsCloudPCDevice: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsCloudPCDevice: %+v", err)
	}

	delete(decoded, "cloudPcStatus")
	delete(decoded, "deviceSpecification")
	delete(decoded, "displayName")
	delete(decoded, "lastRefreshedDateTime")
	delete(decoded, "managedDeviceId")
	delete(decoded, "managedDeviceName")
	delete(decoded, "provisioningPolicyId")
	delete(decoded, "servicePlanName")
	delete(decoded, "servicePlanType")
	delete(decoded, "tenantDisplayName")
	delete(decoded, "tenantId")
	delete(decoded, "userPrincipalName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.cloudPcDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsCloudPCDevice: %+v", err)
	}

	return encoded, nil
}
