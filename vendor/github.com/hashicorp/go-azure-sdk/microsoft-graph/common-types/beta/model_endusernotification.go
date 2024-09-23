package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EndUserNotification{}

type EndUserNotification struct {
	// Identity of the user who created the notification.
	CreatedBy *EmailIdentity `json:"createdBy,omitempty"`

	// Date and time when the notification was created. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description of the notification as defined by the user.
	Description nullable.Type[string] `json:"description,omitempty"`

	Details *[]EndUserNotificationDetail `json:"details,omitempty"`

	// Name of the notification as defined by the user.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Identity of the user who last modified the notification.
	LastModifiedBy *EmailIdentity `json:"lastModifiedBy,omitempty"`

	// Date and time when the notification was last modified. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Type of notification. Possible values are: unknown, positiveReinforcement, noTraining, trainingAssignment,
	// trainingReminder, unknownFutureValue.
	NotificationType *EndUserNotificationType `json:"notificationType,omitempty"`

	// The source of the content. Possible values are: unknown, global, tenant, unknownFutureValue.
	Source *SimulationContentSource `json:"source,omitempty"`

	// The status of the notification. Possible values are: unknown, draft, ready, archive, delete, unknownFutureValue.
	Status *SimulationContentStatus `json:"status,omitempty"`

	// Supported locales for endUserNotification content.
	SupportedLocales *[]string `json:"supportedLocales,omitempty"`

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

func (s EndUserNotification) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EndUserNotification{}

func (s EndUserNotification) MarshalJSON() ([]byte, error) {
	type wrapper EndUserNotification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EndUserNotification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EndUserNotification: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.endUserNotification"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EndUserNotification: %+v", err)
	}

	return encoded, nil
}
