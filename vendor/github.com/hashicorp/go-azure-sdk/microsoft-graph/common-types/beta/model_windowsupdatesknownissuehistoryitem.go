package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesKnownIssueHistoryItem struct {
	Body *WindowsUpdatesItemBody `json:"body,omitempty"`

	// The date and time when the post was created. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = WindowsUpdatesKnownIssueHistoryItem{}

func (s WindowsUpdatesKnownIssueHistoryItem) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesKnownIssueHistoryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesKnownIssueHistoryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesKnownIssueHistoryItem: %+v", err)
	}

	delete(decoded, "createdDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesKnownIssueHistoryItem: %+v", err)
	}

	return encoded, nil
}
