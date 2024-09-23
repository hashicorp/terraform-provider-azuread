package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewNotificationRecipientItem struct {
	// Determines the recipient of the notification email.
	NotificationRecipientScope AccessReviewNotificationRecipientScope `json:"notificationRecipientScope"`

	// Indicates the type of access review email to be sent. Supported template type is CompletedAdditionalRecipients which
	// sends review completion notifications to the recipients.
	NotificationTemplateType nullable.Type[string] `json:"notificationTemplateType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &AccessReviewNotificationRecipientItem{}

func (s *AccessReviewNotificationRecipientItem) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		NotificationTemplateType nullable.Type[string] `json:"notificationTemplateType,omitempty"`
		ODataId                  *string               `json:"@odata.id,omitempty"`
		ODataType                *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.NotificationTemplateType = decoded.NotificationTemplateType
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessReviewNotificationRecipientItem into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["notificationRecipientScope"]; ok {
		impl, err := UnmarshalAccessReviewNotificationRecipientScopeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'NotificationRecipientScope' for 'AccessReviewNotificationRecipientItem': %+v", err)
		}
		s.NotificationRecipientScope = impl
	}

	return nil
}
