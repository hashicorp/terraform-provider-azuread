package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Notification{}

type Notification struct {
	// Sets how long (in seconds) this notification content stays in each platform's notification viewer. For example, when
	// the notification is delivered to a Windows device, the value of this property is passed on to
	// ToastNotification.ExpirationTime, which determines how long the toast notification stays in the user's Windows Action
	// Center.
	DisplayTimeToLive nullable.Type[int64] `json:"displayTimeToLive,omitempty"`

	// Sets a UTC expiration date and time on a user notification using ISO 8601 format (for example, midnight UTC on Jan 1,
	// 2019 would look like this: '2019-01-01T00:00:00Z'). When time is up, the notification is removed from the Microsoft
	// Graph notification feed store completely and is no longer part of notification history. Max value is 30 days.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// The name of the group that this notification belongs to. It is set by the developer for grouping notifications
	// together.
	GroupName nullable.Type[string] `json:"groupName,omitempty"`

	Payload *PayloadTypes `json:"payload,omitempty"`

	// Indicates the priority of a raw user notification. Visual notifications are sent with high priority by default. Valid
	// values are None, High and Low.
	Priority *Priority `json:"priority,omitempty"`

	// Represents the host name of the app to which the calling service wants to post the notification, for the given user.
	// If targeting web endpoints (see targetPolicy.platformTypes), ensure that targetHostName is the same as the name used
	// when creating a subscription on the client side within the application JSON property.
	TargetHostName *string `json:"targetHostName,omitempty"`

	// Target policy object handles notification delivery policy for endpoint types that should be targeted (Windows, iOS,
	// Android and WebPush) for the given user.
	TargetPolicy *TargetPolicyEndpoints `json:"targetPolicy,omitempty"`

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

func (s Notification) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Notification{}

func (s Notification) MarshalJSON() ([]byte, error) {
	type wrapper Notification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Notification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Notification: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.notification"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Notification: %+v", err)
	}

	return encoded, nil
}
