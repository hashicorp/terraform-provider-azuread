package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudCommunications struct {
	CallRecords *[]CallRecordsCallRecord `json:"callRecords,omitempty"`
	Calls       *[]Call                  `json:"calls,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	OnlineMeetings *[]OnlineMeeting `json:"onlineMeetings,omitempty"`
	Presences      *[]Presence      `json:"presences,omitempty"`
}
