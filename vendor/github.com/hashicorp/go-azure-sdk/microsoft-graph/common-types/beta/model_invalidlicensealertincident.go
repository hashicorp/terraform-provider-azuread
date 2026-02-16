package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnifiedRoleManagementAlertIncident = InvalidLicenseAlertIncident{}

type InvalidLicenseAlertIncident struct {
	// Status of the tenant's Microsoft Entra ID P2 license.
	TenantLicenseStatus nullable.Type[string] `json:"tenantLicenseStatus,omitempty"`

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

func (s InvalidLicenseAlertIncident) UnifiedRoleManagementAlertIncident() BaseUnifiedRoleManagementAlertIncidentImpl {
	return BaseUnifiedRoleManagementAlertIncidentImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s InvalidLicenseAlertIncident) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InvalidLicenseAlertIncident{}

func (s InvalidLicenseAlertIncident) MarshalJSON() ([]byte, error) {
	type wrapper InvalidLicenseAlertIncident
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InvalidLicenseAlertIncident: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InvalidLicenseAlertIncident: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.invalidLicenseAlertIncident"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InvalidLicenseAlertIncident: %+v", err)
	}

	return encoded, nil
}
