package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = InsightSummary{}

type InsightSummary struct {
	// Daily active users.
	ActiveUsers *int64 `json:"activeUsers,omitempty"`

	// The ID of the Microsoft Entra application.
	AppId *string `json:"appId,omitempty"`

	// Daily authentication completions.
	AuthenticationCompletions *int64 `json:"authenticationCompletions,omitempty"`

	// Daily authentication requests.
	AuthenticationRequests *int64 `json:"authenticationRequests,omitempty"`

	// The date of the insight.
	FactDate *string `json:"factDate,omitempty"`

	// The platform for the device that the customers used. Supports $filter (eq).
	Os *string `json:"os,omitempty"`

	// Daily MFA SMS completions.
	SecurityTextCompletions *int64 `json:"securityTextCompletions,omitempty"`

	// Daily MFA SMS requests.
	SecurityTextRequests *int64 `json:"securityTextRequests,omitempty"`

	// Daily MFA Voice completions.
	SecurityVoiceCompletions *int64 `json:"securityVoiceCompletions,omitempty"`

	// Daily MFA Voice requests.
	SecurityVoiceRequests *int64 `json:"securityVoiceRequests,omitempty"`

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

func (s InsightSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InsightSummary{}

func (s InsightSummary) MarshalJSON() ([]byte, error) {
	type wrapper InsightSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InsightSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InsightSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.insightSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InsightSummary: %+v", err)
	}

	return encoded, nil
}
