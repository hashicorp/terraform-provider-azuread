package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCLoginResult struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The time of the Cloud PC sign in action. The timestamp is shown in ISO 8601 format and Coordinated Universal Time
	// (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'. Read-only.
	Time nullable.Type[string] `json:"time,omitempty"`
}

var _ json.Marshaler = CloudPCLoginResult{}

func (s CloudPCLoginResult) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCLoginResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCLoginResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCLoginResult: %+v", err)
	}

	delete(decoded, "time")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCLoginResult: %+v", err)
	}

	return encoded, nil
}
