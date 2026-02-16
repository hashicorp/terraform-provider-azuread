package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AudioConferencing struct {
	// The conference id of the online meeting.
	ConferenceId nullable.Type[string] `json:"conferenceId,omitempty"`

	// A URL to the externally-accessible web page that contains dial-in information.
	DialinUrl nullable.Type[string] `json:"dialinUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The toll-free number that connects to the Audio Conference Provider.
	TollFreeNumber nullable.Type[string] `json:"tollFreeNumber,omitempty"`

	// List of toll-free numbers that are displayed in the meeting invite.
	TollFreeNumbers *[]string `json:"tollFreeNumbers,omitempty"`

	// The toll number that connects to the Audio Conference Provider.
	TollNumber nullable.Type[string] `json:"tollNumber,omitempty"`

	// List of toll numbers that are displayed in the meeting invite.
	TollNumbers *[]string `json:"tollNumbers,omitempty"`
}
