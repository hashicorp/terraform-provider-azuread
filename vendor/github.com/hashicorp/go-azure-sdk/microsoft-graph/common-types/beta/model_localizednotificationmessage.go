package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = LocalizedNotificationMessage{}

type LocalizedNotificationMessage struct {
	// Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To
	// unset, set this property to true on another Localized Notification Message.
	IsDefault *bool `json:"isDefault,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The Locale for which this message is destined.
	Locale *string `json:"locale,omitempty"`

	// The Message Template content.
	MessageTemplate *string `json:"messageTemplate,omitempty"`

	// The Message Template Subject.
	Subject *string `json:"subject,omitempty"`

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

func (s LocalizedNotificationMessage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = LocalizedNotificationMessage{}

func (s LocalizedNotificationMessage) MarshalJSON() ([]byte, error) {
	type wrapper LocalizedNotificationMessage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LocalizedNotificationMessage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LocalizedNotificationMessage: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.localizedNotificationMessage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LocalizedNotificationMessage: %+v", err)
	}

	return encoded, nil
}
