package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MentionsPreview struct {
	// True if the signed-in user is mentioned in the parent resource instance. Read-only. Supports filter.
	IsMentioned nullable.Type[bool] `json:"isMentioned,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = MentionsPreview{}

func (s MentionsPreview) MarshalJSON() ([]byte, error) {
	type wrapper MentionsPreview
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MentionsPreview: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MentionsPreview: %+v", err)
	}

	delete(decoded, "isMentioned")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MentionsPreview: %+v", err)
	}

	return encoded, nil
}
