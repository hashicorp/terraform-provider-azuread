package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingNote struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A collection of subpoints of the meeting note.
	Subpoints *[]MeetingNoteSubpoint `json:"subpoints,omitempty"`

	// The text of the meeting note.
	Text nullable.Type[string] `json:"text,omitempty"`

	// The title of the meeting note.
	Title nullable.Type[string] `json:"title,omitempty"`
}
