package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RelyingPartyDetailedSummary{}

type RelyingPartyDetailedSummary struct {
	// Number of failed sign ins on AD FS in the period specified. Supports $orderby, $filter (eq).
	FailedSignInCount *int64 `json:"failedSignInCount,omitempty"`

	MigrationStatus *MigrationStatus `json:"migrationStatus,omitempty"`

	// Specifies all the validations checks done on applications config details.
	MigrationValidationDetails *[]KeyValuePair `json:"migrationValidationDetails,omitempty"`

	// Identifies the relying party to this federation service. It's used when issuing claims to the relying party. Supports
	// $orderby, $filter (eq).
	RelyingPartyId *string `json:"relyingPartyId,omitempty"`

	// Name of the relying party's website or other entity on the Internet that uses an identity provider to authenticate a
	// user who wants to log in. Supports $orderby, $filter (eq).
	RelyingPartyName *string `json:"relyingPartyName,omitempty"`

	// Specifies where the relying party expects to receive the token.
	ReplyUrls *[]string `json:"replyUrls,omitempty"`

	// Uniquely identifies the Active Directory forest. Supports $orderby, $filter (eq).
	ServiceId *string `json:"serviceId,omitempty"`

	// Number of successful sign ins on AD FS. Supports $orderby, $filter (eq).
	SuccessfulSignInCount *int64 `json:"successfulSignInCount,omitempty"`

	// Number of successful + failed sign ins on AD FS in the period specified. Supports $orderby, $filter (eq).
	TotalSignInCount *int64 `json:"totalSignInCount,omitempty"`

	// Number of unique users that signed into the application. Supports $orderby, $filter (eq).
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
