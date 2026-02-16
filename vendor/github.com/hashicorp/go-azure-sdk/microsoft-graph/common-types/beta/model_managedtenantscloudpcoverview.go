package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsCloudPCOverview{}

type ManagedTenantsCloudPCOverview struct {
	// The total number of cloud PC devices that have the Frontline SKU. Optional. Read-only.
	FrontlineLicensesCount nullable.Type[int64] `json:"frontlineLicensesCount,omitempty"`

	// Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
	LastRefreshedDateTime nullable.Type[string] `json:"lastRefreshedDateTime,omitempty"`

	// The number of cloud PC connections that have a status of failed. Optional. Read-only.
	NumberOfCloudPCConnectionStatusFailed nullable.Type[int64] `json:"numberOfCloudPcConnectionStatusFailed,omitempty"`

	// The number of cloud PC connections that have a status of passed. Optional. Read-only.
	NumberOfCloudPCConnectionStatusPassed nullable.Type[int64] `json:"numberOfCloudPcConnectionStatusPassed,omitempty"`

	// The number of cloud PC connections that have a status of pending. Optional. Read-only.
	NumberOfCloudPCConnectionStatusPending nullable.Type[int64] `json:"numberOfCloudPcConnectionStatusPending,omitempty"`

	// The number of cloud PC connections that have a status of running. Optional. Read-only.
	NumberOfCloudPCConnectionStatusRunning nullable.Type[int64] `json:"numberOfCloudPcConnectionStatusRunning,omitempty"`

	// The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
	NumberOfCloudPCConnectionStatusUnkownFutureValue nullable.Type[int64] `json:"numberOfCloudPcConnectionStatusUnkownFutureValue,omitempty"`

	// The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
	NumberOfCloudPCStatusDeprovisioning nullable.Type[int64] `json:"numberOfCloudPcStatusDeprovisioning,omitempty"`

	// The number of cloud PCs that have a status of failed. Optional. Read-only.
	NumberOfCloudPCStatusFailed nullable.Type[int64] `json:"numberOfCloudPcStatusFailed,omitempty"`

	// The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
	NumberOfCloudPCStatusInGracePeriod nullable.Type[int64] `json:"numberOfCloudPcStatusInGracePeriod,omitempty"`

	// The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
	NumberOfCloudPCStatusNotProvisioned nullable.Type[int64] `json:"numberOfCloudPcStatusNotProvisioned,omitempty"`

	// The number of cloud PCs that have a status of provisioned. Optional. Read-only.
	NumberOfCloudPCStatusProvisioned nullable.Type[int64] `json:"numberOfCloudPcStatusProvisioned,omitempty"`

	// The number of cloud PCs that have a status of provisioning. Optional. Read-only.
	NumberOfCloudPCStatusProvisioning nullable.Type[int64] `json:"numberOfCloudPcStatusProvisioning,omitempty"`

	// The number of cloud PCs that have a status of unknown. Optional. Read-only.
	NumberOfCloudPCStatusUnknown nullable.Type[int64] `json:"numberOfCloudPcStatusUnknown,omitempty"`

	// The number of cloud PCs that have a status of upgrading. Optional. Read-only.
	NumberOfCloudPCStatusUpgrading nullable.Type[int64] `json:"numberOfCloudPcStatusUpgrading,omitempty"`

	// The display name for the managed tenant. Optional. Read-only.
	TenantDisplayName nullable.Type[string] `json:"tenantDisplayName,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	// The total number of cloud PC devices that have the Business SKU. Optional. Read-only.
	TotalBusinessLicenses nullable.Type[int64] `json:"totalBusinessLicenses,omitempty"`

	// The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
	TotalCloudPCConnectionStatus nullable.Type[int64] `json:"totalCloudPcConnectionStatus,omitempty"`

	// The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
	TotalCloudPCStatus nullable.Type[int64] `json:"totalCloudPcStatus,omitempty"`

	// The total number of cloud PC devices that have the Enterprise SKU. Optional. Read-only.
	TotalEnterpriseLicenses nullable.Type[int64] `json:"totalEnterpriseLicenses,omitempty"`

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

func (s ManagedTenantsCloudPCOverview) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsCloudPCOverview{}

func (s ManagedTenantsCloudPCOverview) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsCloudPCOverview
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsCloudPCOverview: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsCloudPCOverview: %+v", err)
	}

	delete(decoded, "frontlineLicensesCount")
	delete(decoded, "lastRefreshedDateTime")
	delete(decoded, "numberOfCloudPcConnectionStatusFailed")
	delete(decoded, "numberOfCloudPcConnectionStatusPassed")
	delete(decoded, "numberOfCloudPcConnectionStatusPending")
	delete(decoded, "numberOfCloudPcConnectionStatusRunning")
	delete(decoded, "numberOfCloudPcConnectionStatusUnkownFutureValue")
	delete(decoded, "numberOfCloudPcStatusDeprovisioning")
	delete(decoded, "numberOfCloudPcStatusFailed")
	delete(decoded, "numberOfCloudPcStatusInGracePeriod")
	delete(decoded, "numberOfCloudPcStatusNotProvisioned")
	delete(decoded, "numberOfCloudPcStatusProvisioned")
	delete(decoded, "numberOfCloudPcStatusProvisioning")
	delete(decoded, "numberOfCloudPcStatusUnknown")
	delete(decoded, "numberOfCloudPcStatusUpgrading")
	delete(decoded, "tenantDisplayName")
	delete(decoded, "totalBusinessLicenses")
	delete(decoded, "totalCloudPcConnectionStatus")
	delete(decoded, "totalCloudPcStatus")
	delete(decoded, "totalEnterpriseLicenses")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.cloudPcOverview"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsCloudPCOverview: %+v", err)
	}

	return encoded, nil
}
