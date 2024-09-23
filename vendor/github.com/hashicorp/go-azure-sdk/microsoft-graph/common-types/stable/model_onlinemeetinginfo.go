package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingInfo struct {
	// The ID of the conference.
	ConferenceId nullable.Type[string] `json:"conferenceId,omitempty"`

	// The external link that launches the online meeting. This is a URL that clients launch into a browser and will
	// redirect the user to join the meeting.
	JoinUrl nullable.Type[string] `json:"joinUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// All of the phone numbers associated with this conference.
	Phones *[]Phone `json:"phones,omitempty"`

	// The preformatted quick dial for this call.
	QuickDial nullable.Type[string] `json:"quickDial,omitempty"`

	// The toll free numbers that can be used to join the conference.
	TollFreeNumbers *[]string `json:"tollFreeNumbers,omitempty"`

	// The toll number that can be used to join the conference.
	TollNumber nullable.Type[string] `json:"tollNumber,omitempty"`
}
