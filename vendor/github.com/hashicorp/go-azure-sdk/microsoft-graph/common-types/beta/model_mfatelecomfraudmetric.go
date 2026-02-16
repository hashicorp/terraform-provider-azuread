package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MfaTelecomFraudMetric{}

type MfaTelecomFraudMetric struct {
	CaptchaFailureCount          nullable.Type[int64]  `json:"captchaFailureCount,omitempty"`
	CaptchaNotTriggeredUserCount nullable.Type[int64]  `json:"captchaNotTriggeredUserCount,omitempty"`
	CaptchaShownUserCount        nullable.Type[int64]  `json:"captchaShownUserCount,omitempty"`
	CaptchaSuccessCount          nullable.Type[int64]  `json:"captchaSuccessCount,omitempty"`
	FactDate                     nullable.Type[string] `json:"factDate,omitempty"`
	TelecomBlockedUserCount      nullable.Type[int64]  `json:"telecomBlockedUserCount,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MfaTelecomFraudMetric) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MfaTelecomFraudMetric{}

func (s MfaTelecomFraudMetric) MarshalJSON() ([]byte, error) {
	type wrapper MfaTelecomFraudMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MfaTelecomFraudMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MfaTelecomFraudMetric: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mfaTelecomFraudMetric"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MfaTelecomFraudMetric: %+v", err)
	}

	return encoded, nil
}
