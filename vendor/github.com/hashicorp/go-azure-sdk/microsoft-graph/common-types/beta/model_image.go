package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Image struct {
	// Optional. Height of the image, in pixels. Read-only.
	Height nullable.Type[int64] `json:"height,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Optional. Width of the image, in pixels. Read-only.
	Width nullable.Type[int64] `json:"width,omitempty"`
}

var _ json.Marshaler = Image{}

func (s Image) MarshalJSON() ([]byte, error) {
	type wrapper Image
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Image: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Image: %+v", err)
	}

	delete(decoded, "height")
	delete(decoded, "width")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Image: %+v", err)
	}

	return encoded, nil
}
