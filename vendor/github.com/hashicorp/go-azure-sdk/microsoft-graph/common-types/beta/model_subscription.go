package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Subscription{}

type Subscription struct {
	// Optional. Identifier of the application used to create the subscription. Read-only.
	ApplicationId nullable.Type[string] `json:"applicationId,omitempty"`

	// Required. Indicates the type of change in the subscribed resource that raises a change notification. The supported
	// values are: created, updated, deleted. Multiple values can be combined using a comma-separated list. Note: Drive root
	// item and list change notifications support only the updated changeType. User and group change notifications support
	// updated and deleted changeType. Use updated to receive notifications when user or group is created, updated, or soft
	// deleted. Use deleted to receive notifications when user or group is permanently deleted.
	ChangeType string `json:"changeType"`

	// Optional. Specifies the value of the clientState property sent by the service in each change notification. The
	// maximum length is 255 characters. The client can check that the change notification came from the service by
	// comparing the value of the clientState property sent with the subscription with the value of the clientState property
	// received with each change notification.
	ClientState nullable.Type[string] `json:"clientState,omitempty"`

	// Optional. Identifier of the user or service principal that created the subscription. If the app used delegated
	// permissions to create the subscription, this field contains the ID of the signed-in user the app called on behalf of.
	// If the app used application permissions, this field contains the ID of the service principal corresponding to the
	// app. Read-only.
	CreatorId nullable.Type[string] `json:"creatorId,omitempty"`

	// Optional. A base64-encoded representation of a certificate with a public key used to encrypt resource data in change
	// notifications. Optional but required when includeResourceData is true.
	EncryptionCertificate nullable.Type[string] `json:"encryptionCertificate,omitempty"`

	// Optional. A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Required
	// when includeResourceData is true.
	EncryptionCertificateId nullable.Type[string] `json:"encryptionCertificateId,omitempty"`

	// Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount
	// of time from subscription creation that varies for the resource subscribed to. For the maximum supported subscription
	// length of time, see Subscription lifetime.
	ExpirationDateTime string `json:"expirationDateTime"`

	// Optional. When set to true, change notifications include resource data (such as content of a chat message).
	IncludeResourceData nullable.Type[bool] `json:"includeResourceData,omitempty"`

	// Optional. Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by
	// notificationUrl, supports. The possible values are: v10, v11, v12, v13. For subscribers whose notification endpoint
	// supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline
	// allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these
	// subscribers, not setting this property per the timeline would result in subscription operations failing. For
	// subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases,
	// Microsoft Graph defaults the property to v1_2.
	LatestSupportedTlsVersion nullable.Type[string] `json:"latestSupportedTlsVersion,omitempty"`

	// Required for Teams resources if the expirationDateTime value is more than 1 hour from now; optional otherwise. The
	// URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved, reauthorizationRequired,
	// and missed notifications. This URL must make use of the HTTPS protocol. For more information, see Reduce missing
	// subscriptions and change notifications.
	LifecycleNotificationUrl nullable.Type[string] `json:"lifecycleNotificationUrl,omitempty"`

	// Optional. Desired content-type for Microsoft Graph change notifications for supported resource types. The default
	// content-type is application/json.
	NotificationContentType nullable.Type[string] `json:"notificationContentType,omitempty"`

	// Optional. OData query options for specifying the value for the targeting resource. Clients receive notifications when
	// the resource reaches the state matching the query options provided here. With this new property in the subscription
	// creation payload along with all existing properties, Webhooks deliver notifications whenever a resource reaches the
	// desired state mentioned in the notificationQueryOptions property. For example, when the print job is completed or
	// when a print job resource isFetchable property value becomes true etc. Supported only for Universal Print Service.
	// For more information, see Subscribe to change notifications from cloud printing APIs using Microsoft Graph.
	NotificationQueryOptions nullable.Type[string] `json:"notificationQueryOptions,omitempty"`

	// Required. The URL of the endpoint that receives the change notifications. This URL must make use of the HTTPS
	// protocol. Any query string parameter included in the notificationUrl property is included in the HTTP POST request
	// when Microsoft Graph sends the change notifications.
	NotificationUrl string `json:"notificationUrl"`

	// Optional. The app ID that the subscription service can use to generate the validation token. The value allows the
	// client to validate the authenticity of the notification received.
	NotificationUrlAppId nullable.Type[string] `json:"notificationUrlAppId,omitempty"`

	// Required. Specifies the resource that is monitored for changes. Don't include the base URL
	// (https://graph.microsoft.com/beta/). See the possible resource path values for each supported resource.
	Resource string `json:"resource"`

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

func (s Subscription) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Subscription{}

func (s Subscription) MarshalJSON() ([]byte, error) {
	type wrapper Subscription
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Subscription: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Subscription: %+v", err)
	}

	delete(decoded, "applicationId")
	delete(decoded, "creatorId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.subscription"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Subscription: %+v", err)
	}

	return encoded, nil
}
