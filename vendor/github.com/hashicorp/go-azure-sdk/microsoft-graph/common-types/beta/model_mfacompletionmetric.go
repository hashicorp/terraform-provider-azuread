package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MfaCompletionMetric{}

type MfaCompletionMetric struct {
	// The ID of the Microsoft Entra application. Supports $filter (eq).
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// Number of users who attempted to sign up. Supports $filter (eq).
	AttemptsCount *int64 `json:"attemptsCount,omitempty"`

	Country nullable.Type[string] `json:"country,omitempty"`

	// The date of the user insight.
	FactDate nullable.Type[string] `json:"factDate,omitempty"`

	IdentityProvider nullable.Type[string] `json:"identityProvider,omitempty"`
	Language         nullable.Type[string] `json:"language,omitempty"`
	MfaFailures      *[]MfaFailure         `json:"mfaFailures,omitempty"`

	// The MFA authentication method used by the customers. Supports $filter (eq).
	MfaMethod nullable.Type[string] `json:"mfaMethod,omitempty"`

	// The platform of the device that the customers used. Supports $filter (eq).
	Os nullable.Type[string] `json:"os,omitempty"`

	// Number of users who signed up successfully. Supports $filter (eq).
	SuccessCount *int64 `json:"successCount,omitempty"`

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

func (s MfaCompletionMetric) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MfaCompletionMetric{}

func (s MfaCompletionMetric) MarshalJSON() ([]byte, error) {
	type wrapper MfaCompletionMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MfaCompletionMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MfaCompletionMetric: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mfaCompletionMetric"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MfaCompletionMetric: %+v", err)
	}

	return encoded, nil
}
