package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Hashes struct {
	// The CRC32 value of the file (if available). Read-only.
	Crc32Hash nullable.Type[string] `json:"crc32Hash,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A proprietary hash of the file that can be used to determine if the contents of the file have changed (if available).
	// Read-only.
	QuickXorHash nullable.Type[string] `json:"quickXorHash,omitempty"`

	// SHA1 hash for the contents of the file (if available). Read-only.
	Sha1Hash nullable.Type[string] `json:"sha1Hash,omitempty"`

	// SHA256 hash for the contents of the file (if available). Read-only.
	Sha256Hash nullable.Type[string] `json:"sha256Hash,omitempty"`
}

var _ json.Marshaler = Hashes{}

func (s Hashes) MarshalJSON() ([]byte, error) {
	type wrapper Hashes
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Hashes: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Hashes: %+v", err)
	}

	delete(decoded, "crc32Hash")
	delete(decoded, "quickXorHash")
	delete(decoded, "sha1Hash")
	delete(decoded, "sha256Hash")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Hashes: %+v", err)
	}

	return encoded, nil
}
