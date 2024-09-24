package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudRealtimeCommunicationInfo struct {
	// Indicates whether the user has a SIP-enabled client registered for them. Read-only.
	IsSipEnabled nullable.Type[bool] `json:"isSipEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = CloudRealtimeCommunicationInfo{}

func (s CloudRealtimeCommunicationInfo) MarshalJSON() ([]byte, error) {
	type wrapper CloudRealtimeCommunicationInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudRealtimeCommunicationInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudRealtimeCommunicationInfo: %+v", err)
	}

	delete(decoded, "isSipEnabled")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudRealtimeCommunicationInfo: %+v", err)
	}

	return encoded, nil
}
