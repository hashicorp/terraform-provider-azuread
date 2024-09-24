package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CredentialUsageSummary{}

type CredentialUsageSummary struct {
	AuthMethod *UsageAuthMethod `json:"authMethod,omitempty"`

	// Provides the count of failed resets or registration data.
	FailureActivityCount *int64 `json:"failureActivityCount,omitempty"`

	Feature *FeatureType `json:"feature,omitempty"`

	// Provides the count of successful registrations or resets.
	SuccessfulActivityCount *int64 `json:"successfulActivityCount,omitempty"`

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

func (s CredentialUsageSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CredentialUsageSummary{}

func (s CredentialUsageSummary) MarshalJSON() ([]byte, error) {
	type wrapper CredentialUsageSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CredentialUsageSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CredentialUsageSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.credentialUsageSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CredentialUsageSummary: %+v", err)
	}

	return encoded, nil
}
