package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NotificationMessageTemplate{}

type NotificationMessageTemplate struct {
	// Branding Options for the Message Template. Branding is defined in the Intune Admin Console.
	BrandingOptions *NotificationTemplateBrandingOptions `json:"brandingOptions,omitempty"`

	// The default locale to fallback onto when the requested locale is not available.
	DefaultLocale nullable.Type[string] `json:"defaultLocale,omitempty"`

	// Display name for the Notification Message Template.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name for the Notification Message Template.
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The list of localized messages for this Notification Message Template.
	LocalizedNotificationMessages *[]LocalizedNotificationMessage `json:"localizedNotificationMessages,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

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

func (s NotificationMessageTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NotificationMessageTemplate{}

func (s NotificationMessageTemplate) MarshalJSON() ([]byte, error) {
	type wrapper NotificationMessageTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NotificationMessageTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NotificationMessageTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.notificationMessageTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NotificationMessageTemplate: %+v", err)
	}

	return encoded, nil
}
