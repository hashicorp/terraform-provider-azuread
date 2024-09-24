package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AdminConsentRequestPolicy{}

type AdminConsentRequestPolicy struct {
	// Specifies whether the admin consent request feature is enabled or disabled. Required.
	IsEnabled bool `json:"isEnabled"`

	// Specifies whether reviewers will receive notifications. Required.
	NotifyReviewers bool `json:"notifyReviewers"`

	// Specifies whether reviewers will receive reminder emails. Required.
	RemindersEnabled bool `json:"remindersEnabled"`

	// Specifies the duration the request is active before it automatically expires if no decision is applied.
	RequestDurationInDays *int64 `json:"requestDurationInDays,omitempty"`

	// The list of reviewers for the admin consent. Required.
	Reviewers []AccessReviewReviewerScope `json:"reviewers"`

	// Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
	Version *int64 `json:"version,omitempty"`

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

func (s AdminConsentRequestPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AdminConsentRequestPolicy{}

func (s AdminConsentRequestPolicy) MarshalJSON() ([]byte, error) {
	type wrapper AdminConsentRequestPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AdminConsentRequestPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AdminConsentRequestPolicy: %+v", err)
	}

	delete(decoded, "version")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.adminConsentRequestPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AdminConsentRequestPolicy: %+v", err)
	}

	return encoded, nil
}
