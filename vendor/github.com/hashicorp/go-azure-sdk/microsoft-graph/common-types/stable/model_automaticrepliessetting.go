package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticRepliesSetting struct {
	// The set of audience external to the signed-in user's organization who will receive the ExternalReplyMessage, if
	// Status is AlwaysEnabled or Scheduled. The possible values are: none, contactsOnly, all.
	ExternalAudience *ExternalAudienceScope `json:"externalAudience,omitempty"`

	// The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled.
	ExternalReplyMessage nullable.Type[string] `json:"externalReplyMessage,omitempty"`

	// The automatic reply to send to the audience internal to the signed-in user's organization, if Status is AlwaysEnabled
	// or Scheduled.
	InternalReplyMessage nullable.Type[string] `json:"internalReplyMessage,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date and time that automatic replies are set to end, if Status is set to Scheduled.
	ScheduledEndDateTime *DateTimeTimeZone `json:"scheduledEndDateTime,omitempty"`

	// The date and time that automatic replies are set to begin, if Status is set to Scheduled.
	ScheduledStartDateTime *DateTimeTimeZone `json:"scheduledStartDateTime,omitempty"`

	// Configurations status for automatic replies. The possible values are: disabled, alwaysEnabled, scheduled.
	Status *AutomaticRepliesStatus `json:"status,omitempty"`
}
