package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PayloadDetail = EmailPayloadDetail{}

type EmailPayloadDetail struct {
	// Email address of the user.
	FromEmail nullable.Type[string] `json:"fromEmail,omitempty"`

	// Display name of the user.
	FromName nullable.Type[string] `json:"fromName,omitempty"`

	// Indicates whether the sender isn't from the user's organization.
	IsExternalSender nullable.Type[bool] `json:"isExternalSender,omitempty"`

	// The subject of the email address sent to the user.
	Subject nullable.Type[string] `json:"subject,omitempty"`

	// Fields inherited from PayloadDetail

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

func (s EmailPayloadDetail) PayloadDetail() BasePayloadDetailImpl {
	return BasePayloadDetailImpl{
		Coachmarks:  s.Coachmarks,
		Content:     s.Content,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		PhishingUrl: s.PhishingUrl,
	}
}

var _ json.Marshaler = EmailPayloadDetail{}

func (s EmailPayloadDetail) MarshalJSON() ([]byte, error) {
	type wrapper EmailPayloadDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EmailPayloadDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EmailPayloadDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.emailPayloadDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EmailPayloadDetail: %+v", err)
	}

	return encoded, nil
}
