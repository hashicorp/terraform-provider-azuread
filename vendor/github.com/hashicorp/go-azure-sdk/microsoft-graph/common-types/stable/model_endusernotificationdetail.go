package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EndUserNotificationDetail{}

type EndUserNotificationDetail struct {
	// Email HTML content.
	EmailContent *string `json:"emailContent,omitempty"`

	// Indicates whether this language is default.
	IsDefaultLangauge nullable.Type[bool] `json:"isDefaultLangauge,omitempty"`

	// Notification language.
	Language *string `json:"language,omitempty"`

	// Notification locale.
	Locale *string `json:"locale,omitempty"`

	SentFrom *EmailIdentity `json:"sentFrom,omitempty"`

	// Mail subject.
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

func (s EndUserNotificationDetail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EndUserNotificationDetail{}

func (s EndUserNotificationDetail) MarshalJSON() ([]byte, error) {
	type wrapper EndUserNotificationDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EndUserNotificationDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EndUserNotificationDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.endUserNotificationDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EndUserNotificationDetail: %+v", err)
	}

	return encoded, nil
}
