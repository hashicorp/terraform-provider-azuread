package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringEmailNotificationConfiguration struct {
	// The identifier of the group to send an email to. All group types with configured email addresses are supported.
	GroupId *string `json:"groupId,omitempty"`

	// Indicates whether email notifications are enabled on the alert type.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
