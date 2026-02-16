package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagedTenantAlert{}

type ManagedTenantsManagedTenantAlert struct {
	AlertData                 *ManagedTenantsAlertData                        `json:"alertData,omitempty"`
	AlertDataReferenceStrings *[]ManagedTenantsAlertDataReferenceString       `json:"alertDataReferenceStrings,omitempty"`
	AlertLogs                 *[]ManagedTenantsManagedTenantAlertLog          `json:"alertLogs,omitempty"`
	AlertRule                 *ManagedTenantsManagedTenantAlertRule           `json:"alertRule,omitempty"`
	AlertRuleDisplayName      nullable.Type[string]                           `json:"alertRuleDisplayName,omitempty"`
	ApiNotifications          *[]ManagedTenantsManagedTenantApiNotification   `json:"apiNotifications,omitempty"`
	AssignedToUserId          nullable.Type[string]                           `json:"assignedToUserId,omitempty"`
	CorrelationCount          nullable.Type[int64]                            `json:"correlationCount,omitempty"`
	CorrelationId             nullable.Type[string]                           `json:"correlationId,omitempty"`
	CreatedByUserId           nullable.Type[string]                           `json:"createdByUserId,omitempty"`
	CreatedDateTime           nullable.Type[string]                           `json:"createdDateTime,omitempty"`
	EmailNotifications        *[]ManagedTenantsManagedTenantEmailNotification `json:"emailNotifications,omitempty"`
	LastActionByUserId        nullable.Type[string]                           `json:"lastActionByUserId,omitempty"`
	LastActionDateTime        nullable.Type[string]                           `json:"lastActionDateTime,omitempty"`
	Message                   nullable.Type[string]                           `json:"message,omitempty"`
	Severity                  *ManagedTenantsAlertSeverity                    `json:"severity,omitempty"`
	Status                    *ManagedTenantsAlertStatus                      `json:"status,omitempty"`
	TenantId                  nullable.Type[string]                           `json:"tenantId,omitempty"`
	Title                     nullable.Type[string]                           `json:"title,omitempty"`

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

func (s ManagedTenantsManagedTenantAlert) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagedTenantAlert{}

func (s ManagedTenantsManagedTenantAlert) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagedTenantAlert
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagedTenantAlert: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedTenantAlert: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managedTenantAlert"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagedTenantAlert: %+v", err)
	}

	return encoded, nil
}
