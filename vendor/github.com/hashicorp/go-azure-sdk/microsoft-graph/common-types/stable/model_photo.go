package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Photo struct {
	// Camera manufacturer. Read-only.
	CameraMake nullable.Type[string] `json:"cameraMake,omitempty"`

	// Camera model. Read-only.
	CameraModel nullable.Type[string] `json:"cameraModel,omitempty"`

	// The ISO value from the camera. Read-only.
	Iso nullable.Type[int64] `json:"iso,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The orientation value from the camera. Writable on OneDrive Personal.
	Orientation nullable.Type[int64] `json:"orientation,omitempty"`

	// Represents the date and time the photo was taken. Read-only.
	TakenDateTime nullable.Type[string] `json:"takenDateTime,omitempty"`
}

var _ json.Marshaler = Photo{}

func (s Photo) MarshalJSON() ([]byte, error) {
	type wrapper Photo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Photo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Photo: %+v", err)
	}

	delete(decoded, "cameraMake")
	delete(decoded, "cameraModel")
	delete(decoded, "iso")
	delete(decoded, "takenDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Photo: %+v", err)
	}

	return encoded, nil
}
