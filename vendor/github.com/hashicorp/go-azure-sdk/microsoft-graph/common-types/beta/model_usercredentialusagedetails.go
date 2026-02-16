package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserCredentialUsageDetails{}

type UserCredentialUsageDetails struct {
	AuthMethod *UsageAuthMethod `json:"authMethod,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	EventDateTime *string `json:"eventDateTime,omitempty"`

	// Provides the failure reason for the corresponding reset or registration workflow.
	FailureReason *string `json:"failureReason,omitempty"`

	Feature *FeatureType `json:"feature,omitempty"`

	// Indicates success or failure of the workflow.
	IsSuccess *bool `json:"isSuccess,omitempty"`

	// User name of the user performing the reset or registration workflow.
	UserDisplayName *string `json:"userDisplayName,omitempty"`

	// User principal name of the user performing the reset or registration workflow.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`

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

func (s UserCredentialUsageDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserCredentialUsageDetails{}

func (s UserCredentialUsageDetails) MarshalJSON() ([]byte, error) {
	type wrapper UserCredentialUsageDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserCredentialUsageDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserCredentialUsageDetails: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userCredentialUsageDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserCredentialUsageDetails: %+v", err)
	}

	return encoded, nil
}
