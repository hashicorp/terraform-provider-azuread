package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityRetentionDuration = SecurityRetentionDurationInDays{}

type SecurityRetentionDurationInDays struct {
	// Specifies the time period in days for which an item with the applied retention label will be retained for.
	Days nullable.Type[int64] `json:"days,omitempty"`

	// Fields inherited from SecurityRetentionDuration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityRetentionDurationInDays) SecurityRetentionDuration() BaseSecurityRetentionDurationImpl {
	return BaseSecurityRetentionDurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityRetentionDurationInDays{}

func (s SecurityRetentionDurationInDays) MarshalJSON() ([]byte, error) {
	type wrapper SecurityRetentionDurationInDays
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityRetentionDurationInDays: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityRetentionDurationInDays: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.retentionDurationInDays"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityRetentionDurationInDays: %+v", err)
	}

	return encoded, nil
}
