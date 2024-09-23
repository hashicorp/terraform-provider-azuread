package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedRoleManagementAlertDefinition{}

type UnifiedRoleManagementAlertDefinition struct {
	// The description of the alert.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The friendly display name that renders in Privileged Identity Management (PIM) alerts in the Microsoft Entra admin
	// center.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Long-form text that indicates the ways to prevent the alert from being triggered in your tenant.
	HowToPrevent nullable.Type[string] `json:"howToPrevent,omitempty"`

	// true if the alert configuration can be customized in the tenant, and false otherwise. For example, the number and
	// percentage thresholds of the 'There are too many global administrators' alert can be configured by users, while the
	// 'This organization doesn't have Microsoft Entra ID P2' can't be configured, because the criteria are restricted.
	IsConfigurable nullable.Type[bool] `json:"isConfigurable,omitempty"`

	// true if the alert can be remediated, and false otherwise.
	IsRemediatable nullable.Type[bool] `json:"isRemediatable,omitempty"`

	// The methods to mitigate the alert when it's triggered in the tenant. For example, to mitigate the 'There are too many
	// global administrators', you could remove redundant privileged role assignments.
	MitigationSteps nullable.Type[string] `json:"mitigationSteps,omitempty"`

	// The identifier of the scope where the alert is related. / is the only supported one for the tenant. Supports $filter
	// (eq, ne).
	ScopeId nullable.Type[string] `json:"scopeId,omitempty"`

	// The type of scope where the alert is created. DirectoryRole is the only currently supported scope type for Microsoft
	// Entra roles.
	ScopeType nullable.Type[string] `json:"scopeType,omitempty"`

	// Security impact of the alert. For example, it could be information leaks or unauthorized access.
	SecurityImpact nullable.Type[string] `json:"securityImpact,omitempty"`

	// Severity level of the alert. The possible values are: unknown, informational, low, medium, high, unknownFutureValue.
	SeverityLevel *AlertSeverity `json:"severityLevel,omitempty"`

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

func (s UnifiedRoleManagementAlertDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleManagementAlertDefinition{}

func (s UnifiedRoleManagementAlertDefinition) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleManagementAlertDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleManagementAlertDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleManagementAlertDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleManagementAlertDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleManagementAlertDefinition: %+v", err)
	}

	return encoded, nil
}
