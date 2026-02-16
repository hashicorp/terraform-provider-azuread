package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingDetail struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The user who shared the document.
	SharedBy *InsightIdentity `json:"sharedBy,omitempty"`

	// The date and time the file was last shared. The timestamp represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	SharedDateTime nullable.Type[string] `json:"sharedDateTime,omitempty"`

	// Reference properties of the document, such as the URL and type of the document. Read-only
	SharingReference *ResourceReference `json:"sharingReference,omitempty"`

	// The subject with which the document was shared.
	SharingSubject nullable.Type[string] `json:"sharingSubject,omitempty"`

	// Determines the way the document was shared, can be by a 'Link', 'Attachment', 'Group', 'Site'.
	SharingType nullable.Type[string] `json:"sharingType,omitempty"`
}

var _ json.Marshaler = SharingDetail{}

func (s SharingDetail) MarshalJSON() ([]byte, error) {
	type wrapper SharingDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SharingDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SharingDetail: %+v", err)
	}

	delete(decoded, "sharedDateTime")
	delete(decoded, "sharingReference")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SharingDetail: %+v", err)
	}

	return encoded, nil
}
