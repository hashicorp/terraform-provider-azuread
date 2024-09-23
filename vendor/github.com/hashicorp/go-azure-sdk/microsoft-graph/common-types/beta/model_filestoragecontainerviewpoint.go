package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileStorageContainerViewpoint struct {
	// The current user's effective role. Read-only.
	EffectiveRole *string `json:"effectiveRole,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = FileStorageContainerViewpoint{}

func (s FileStorageContainerViewpoint) MarshalJSON() ([]byte, error) {
	type wrapper FileStorageContainerViewpoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FileStorageContainerViewpoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FileStorageContainerViewpoint: %+v", err)
	}

	delete(decoded, "effectiveRole")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FileStorageContainerViewpoint: %+v", err)
	}

	return encoded, nil
}
