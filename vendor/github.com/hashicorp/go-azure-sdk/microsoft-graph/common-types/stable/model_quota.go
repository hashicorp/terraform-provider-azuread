package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Quota struct {
	// Total space consumed by files in the recycle bin, in bytes. Read-only.
	Deleted nullable.Type[int64] `json:"deleted,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Total space remaining before reaching the capacity limit, in bytes. Read-only.
	Remaining nullable.Type[int64] `json:"remaining,omitempty"`

	// Enumeration value that indicates the state of the storage space. Read-only.
	State nullable.Type[string] `json:"state,omitempty"`

	// Information about the drive's storage quota plans. Only in Personal OneDrive.
	StoragePlanInformation *StoragePlanInformation `json:"storagePlanInformation,omitempty"`

	// Total allowed storage space, in bytes. Read-only.
	Total nullable.Type[int64] `json:"total,omitempty"`

	// Total space used, in bytes. Read-only.
	Used nullable.Type[int64] `json:"used,omitempty"`
}

var _ json.Marshaler = Quota{}

func (s Quota) MarshalJSON() ([]byte, error) {
	type wrapper Quota
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Quota: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Quota: %+v", err)
	}

	delete(decoded, "deleted")
	delete(decoded, "remaining")
	delete(decoded, "state")
	delete(decoded, "total")
	delete(decoded, "used")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Quota: %+v", err)
	}

	return encoded, nil
}
