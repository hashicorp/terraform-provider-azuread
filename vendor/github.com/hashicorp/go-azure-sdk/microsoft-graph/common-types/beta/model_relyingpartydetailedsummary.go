package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RelyingPartyDetailedSummary{}

type RelyingPartyDetailedSummary struct {
	// Number of failed sign in on Active Directory Federation Service in the period specified.
	FailedSignInCount *int64 `json:"failedSignInCount,omitempty"`

	MigrationStatus *MigrationStatus `json:"migrationStatus,omitempty"`

	// Specifies all the validations check done on applications configuration details to evaluate if the application is
	// ready to be moved to Microsoft Entra ID.
	MigrationValidationDetails *[]KeyValuePair `json:"migrationValidationDetails,omitempty"`

	// This identifier is used to identify the relying party to this Federation Service. It's used when issuing claims to
	// the relying party.
	RelyingPartyId *string `json:"relyingPartyId,omitempty"`

	// Name of application or other entity on the internet that uses an identity provider to authenticate a user who wants
	// to sign in.
	RelyingPartyName *string `json:"relyingPartyName,omitempty"`

	// Specifies where the relying party expects to receive the token.
	ReplyUrls *[]string `json:"replyUrls,omitempty"`

	// Uniquely identifies the Active Directory forest.
	ServiceId *string `json:"serviceId,omitempty"`

	// Number of successful sign ins on Active Directory Federation Service.
	SuccessfulSignInCount *int64 `json:"successfulSignInCount,omitempty"`

	// Number of successful + failed sign ins on Active Directory Federation Service in the period specified.
	TotalSignInCount *int64 `json:"totalSignInCount,omitempty"`

	// Number of unique users that have signed into the application.
	UniqueUserCount *int64 `json:"uniqueUserCount,omitempty"`

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

func (s RelyingPartyDetailedSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RelyingPartyDetailedSummary{}

func (s RelyingPartyDetailedSummary) MarshalJSON() ([]byte, error) {
	type wrapper RelyingPartyDetailedSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RelyingPartyDetailedSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RelyingPartyDetailedSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.relyingPartyDetailedSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RelyingPartyDetailedSummary: %+v", err)
	}

	return encoded, nil
}
