package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnifiedRoleManagementPolicyRule = UnifiedRoleManagementPolicyNotificationRule{}

type UnifiedRoleManagementPolicyNotificationRule struct {
	// Indicates whether a default recipient will receive the notification email.
	IsDefaultRecipientsEnabled nullable.Type[bool] `json:"isDefaultRecipientsEnabled,omitempty"`

	// The level of notification. The possible values are None, Critical, All.
	NotificationLevel nullable.Type[string] `json:"notificationLevel,omitempty"`

	// The list of recipients of the email notifications.
	NotificationRecipients *[]string `json:"notificationRecipients,omitempty"`

	// The type of notification. Only Email is supported.
	NotificationType nullable.Type[string] `json:"notificationType,omitempty"`

	// The type of recipient of the notification. The possible values are Requestor, Approver, Admin.
	RecipientType nullable.Type[string] `json:"recipientType,omitempty"`

	// Fields inherited from UnifiedRoleManagementPolicyRule

	// Not implemented. Defines details of scope that's targeted by role management policy rule. The details can include the
	// principal type, the role assignment type, and actions affecting a role. Supports $filter (eq, ne).
	Target *UnifiedRoleManagementPolicyRuleTarget `json:"target,omitempty"`

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

func (s UnifiedRoleManagementPolicyNotificationRule) UnifiedRoleManagementPolicyRule() BaseUnifiedRoleManagementPolicyRuleImpl {
	return BaseUnifiedRoleManagementPolicyRuleImpl{
		Target:    s.Target,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s UnifiedRoleManagementPolicyNotificationRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleManagementPolicyNotificationRule{}

func (s UnifiedRoleManagementPolicyNotificationRule) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleManagementPolicyNotificationRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleManagementPolicyNotificationRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleManagementPolicyNotificationRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleManagementPolicyNotificationRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleManagementPolicyNotificationRule: %+v", err)
	}

	return encoded, nil
}
