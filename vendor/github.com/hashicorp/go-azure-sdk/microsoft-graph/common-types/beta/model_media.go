package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Media struct {
	// If a file has a transcript, this setting controls if the closed captions / transcription for the media file should be
	// shown to people during viewing. Read-Write.
	IsTranscriptionShown nullable.Type[bool] `json:"isTranscriptionShown,omitempty"`

	// Information about the source of media. Read-only.
	MediaSource *MediaSource `json:"mediaSource,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = Media{}

func (s Media) MarshalJSON() ([]byte, error) {
	type wrapper Media
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Media: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Media: %+v", err)
	}

	delete(decoded, "mediaSource")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Media: %+v", err)
	}

	return encoded, nil
}
