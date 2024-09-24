package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type File struct {
	// Hashes of the file's binary content, if available. Read-only.
	Hashes *Hashes `json:"hashes,omitempty"`

	// The MIME type for the file. This is determined by logic on the server and might not be the value provided when the
	// file was uploaded. Read-only.
	MimeType nullable.Type[string] `json:"mimeType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ProcessingMetadata nullable.Type[bool] `json:"processingMetadata,omitempty"`
}

var _ json.Marshaler = File{}

func (s File) MarshalJSON() ([]byte, error) {
	type wrapper File
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling File: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling File: %+v", err)
	}

	delete(decoded, "hashes")
	delete(decoded, "mimeType")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling File: %+v", err)
	}

	return encoded, nil
}
