package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecentNotebook struct {
	// The name of the notebook.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time when the notebook was last modified. The timestamp represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	LastAccessedTime nullable.Type[string] `json:"lastAccessedTime,omitempty"`

	// Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote client, if it's
	// installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
	Links *RecentNotebookLinks `json:"links,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The backend store where the Notebook resides, either OneDriveForBusiness or OneDrive.
	SourceService *OnenoteSourceService `json:"sourceService,omitempty"`
}

var _ json.Marshaler = RecentNotebook{}

func (s RecentNotebook) MarshalJSON() ([]byte, error) {
	type wrapper RecentNotebook
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RecentNotebook: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RecentNotebook: %+v", err)
	}

	delete(decoded, "lastAccessedTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RecentNotebook: %+v", err)
	}

	return encoded, nil
}
