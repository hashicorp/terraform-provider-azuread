package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagedTenantAlertRule{}

type ManagedTenantsManagedTenantAlertRule struct {
	AlertDisplayName              nullable.Type[string]                           `json:"alertDisplayName,omitempty"`
	AlertTTL                      nullable.Type[int64]                            `json:"alertTTL,omitempty"`
	Alerts                        *[]ManagedTenantsManagedTenantAlert             `json:"alerts,omitempty"`
	CreatedByUserId               nullable.Type[string]                           `json:"createdByUserId,omitempty"`
	CreatedDateTime               nullable.Type[string]                           `json:"createdDateTime,omitempty"`
	Description                   nullable.Type[string]                           `json:"description,omitempty"`
	DisplayName                   nullable.Type[string]                           `json:"displayName,omitempty"`
	LastActionByUserId            nullable.Type[string]                           `json:"lastActionByUserId,omitempty"`
	LastActionDateTime            nullable.Type[string]                           `json:"lastActionDateTime,omitempty"`
	LastRunDateTime               nullable.Type[string]                           `json:"lastRunDateTime,omitempty"`
	NotificationFinalDestinations *ManagedTenantsNotificationDestination          `json:"notificationFinalDestinations,omitempty"`
	RuleDefinition                *ManagedTenantsManagedTenantAlertRuleDefinition `json:"ruleDefinition,omitempty"`
	Severity                      *ManagedTenantsAlertSeverity                    `json:"severity,omitempty"`
	Targets                       *[]ManagedTenantsNotificationTarget             `json:"targets,omitempty"`
	TenantIds                     *[]ManagedTenantsTenantInfo                     `json:"tenantIds,omitempty"`

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

func (s ManagedTenantsManagedTenantAlertRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagedTenantAlertRule{}

func (s ManagedTenantsManagedTenantAlertRule) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagedTenantAlertRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagedTenantAlertRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedTenantAlertRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managedTenantAlertRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagedTenantAlertRule: %+v", err)
	}

	return encoded, nil
}
