package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadDetail interface {
	PayloadDetail() BasePayloadDetailImpl
}

var _ PayloadDetail = BasePayloadDetailImpl{}

type BasePayloadDetailImpl struct {
	Coachmarks *[]PayloadCoachmark `json:"coachmarks,omitempty"`

	// Payload content details.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The phishing URL used to target a user.
	PhishingUrl nullable.Type[string] `json:"phishingUrl,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePayloadDetailImpl) PayloadDetail() BasePayloadDetailImpl {
	return s
}

var _ PayloadDetail = RawPayloadDetailImpl{}

// RawPayloadDetailImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPayloadDetailImpl struct {
	payloadDetail BasePayloadDetailImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawPayloadDetailImpl) PayloadDetail() BasePayloadDetailImpl {
	return s.payloadDetail
}

func UnmarshalPayloadDetailImplementation(input []byte) (PayloadDetail, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PayloadDetail into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.emailPayloadDetail") {
		var out EmailPayloadDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailPayloadDetail: %+v", err)
		}
		return out, nil
	}

	var parent BasePayloadDetailImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePayloadDetailImpl: %+v", err)
	}

	return RawPayloadDetailImpl{
		payloadDetail: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
