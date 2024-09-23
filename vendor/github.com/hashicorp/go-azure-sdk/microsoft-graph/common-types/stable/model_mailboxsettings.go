package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxSettings struct {
	// Folder ID of an archive folder for the user.
	ArchiveFolder nullable.Type[string] `json:"archiveFolder,omitempty"`

	// Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in
	// user.
	AutomaticRepliesSetting *AutomaticRepliesSetting `json:"automaticRepliesSetting,omitempty"`

	// The date format for the user's mailbox.
	DateFormat nullable.Type[string] `json:"dateFormat,omitempty"`

	// If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting
	// messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal,
	// sendToDelegateAndPrincipal, sendToDelegateOnly.
	DelegateMeetingMessageDeliveryOptions *DelegateMeetingMessageDeliveryOptions `json:"delegateMeetingMessageDeliveryOptions,omitempty"`

	// The locale information for the user, including the preferred language and country/region.
	Language *LocaleInfo `json:"language,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The time format for the user's mailbox.
	TimeFormat nullable.Type[string] `json:"timeFormat,omitempty"`

	// The default time zone for the user's mailbox.
	TimeZone nullable.Type[string] `json:"timeZone,omitempty"`

	// The purpose of the mailbox. Differentiates a mailbox for a single user from a shared mailbox and equipment mailbox in
	// Exchange Online. Possible values are: user, linked, shared, room, equipment, others, unknownFutureValue. Read-only.
	UserPurpose *UserPurpose `json:"userPurpose,omitempty"`

	// The days of the week and hours in a specific time zone that the user works.
	WorkingHours *WorkingHours `json:"workingHours,omitempty"`
}

var _ json.Marshaler = MailboxSettings{}

func (s MailboxSettings) MarshalJSON() ([]byte, error) {
	type wrapper MailboxSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MailboxSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MailboxSettings: %+v", err)
	}

	delete(decoded, "userPurpose")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MailboxSettings: %+v", err)
	}

	return encoded, nil
}
