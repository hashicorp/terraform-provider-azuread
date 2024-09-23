package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeNotification struct {
	ChangeType *ChangeType `json:"changeType,omitempty"`

	// Value of the clientState property sent in the subscription request (if any). The maximum length is 255 characters.
	// The client can check whether the change notification came from the service by comparing the values of the clientState
	// property. The value of the clientState property sent with the subscription is compared with the value of the
	// clientState property received with each change notification. Optional.
	ClientState nullable.Type[string] `json:"clientState,omitempty"`

	// (Preview) Encrypted content attached with the change notification. Only provided if encryptionCertificate and
	// includeResourceData were defined during the subscription request and if the resource supports it. Optional.
	EncryptedContent *ChangeNotificationEncryptedContent `json:"encryptedContent,omitempty"`

	// Unique ID for the notification. Optional.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The type of lifecycle notification if the current notification is a lifecycle notification. Optional. Supported
	// values are missed, subscriptionRemoved, reauthorizationRequired. Optional.
	LifecycleEvent *LifecycleEventType `json:"lifecycleEvent,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The URI of the resource that emitted the change notification relative to https://graph.microsoft.com. Required.
	Resource string `json:"resource"`

	// The content of this property depends on the type of resource being subscribed to. Optional.
	ResourceData *ResourceData `json:"resourceData,omitempty"`

	// The expiration time for the subscription. Required.
	SubscriptionExpirationDateTime string `json:"subscriptionExpirationDateTime"`

	// The unique identifier of the subscription that generated the notification.Required.
	SubscriptionId *string `json:"subscriptionId,omitempty"`

	// The unique identifier of the tenant from which the change notification originated. Required.
	TenantId string `json:"tenantId"`
}
