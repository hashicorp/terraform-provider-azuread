package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppTrack struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Friendly name for track. This property is read-only.
	TrackAlias nullable.Type[string] `json:"trackAlias,omitempty"`

	// Unique track identifier. This property is read-only.
	TrackId nullable.Type[string] `json:"trackId,omitempty"`
}

var _ json.Marshaler = AndroidManagedStoreAppTrack{}

func (s AndroidManagedStoreAppTrack) MarshalJSON() ([]byte, error) {
	type wrapper AndroidManagedStoreAppTrack
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidManagedStoreAppTrack: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidManagedStoreAppTrack: %+v", err)
	}

	delete(decoded, "trackAlias")
	delete(decoded, "trackId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidManagedStoreAppTrack: %+v", err)
	}

	return encoded, nil
}
