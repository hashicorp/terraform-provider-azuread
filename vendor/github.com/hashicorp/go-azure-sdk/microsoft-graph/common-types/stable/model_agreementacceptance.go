package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AgreementAcceptance{}

type AgreementAcceptance struct {
	// The identifier of the agreement file accepted by the user.
	AgreementFileId nullable.Type[string] `json:"agreementFileId,omitempty"`

	// The identifier of the agreement.
	AgreementId nullable.Type[string] `json:"agreementId,omitempty"`

	// The display name of the device used for accepting the agreement.
	DeviceDisplayName nullable.Type[string] `json:"deviceDisplayName,omitempty"`

	// The unique identifier of the device used for accepting the agreement. Supports $filter (eq) and eq for null values.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The operating system used to accept the agreement.
	DeviceOSType nullable.Type[string] `json:"deviceOSType,omitempty"`

	// The operating system version of the device used to accept the agreement.
	DeviceOSVersion nullable.Type[string] `json:"deviceOSVersion,omitempty"`

	// The expiration date time of the acceptance. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter
	// (eq, ge, le) and eq for null values.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	RecordedDateTime nullable.Type[string] `json:"recordedDateTime,omitempty"`

	// The state of the agreement acceptance. Possible values are: accepted, declined. Supports $filter (eq).
	State *AgreementAcceptanceState `json:"state,omitempty"`

	// Display name of the user when the acceptance was recorded.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// Email of the user when the acceptance was recorded.
	UserEmail nullable.Type[string] `json:"userEmail,omitempty"`

	// The identifier of the user who accepted the agreement. Supports $filter (eq).
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// UPN of the user when the acceptance was recorded.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s AgreementAcceptance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AgreementAcceptance{}

func (s AgreementAcceptance) MarshalJSON() ([]byte, error) {
	type wrapper AgreementAcceptance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AgreementAcceptance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AgreementAcceptance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.agreementAcceptance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AgreementAcceptance: %+v", err)
	}

	return encoded, nil
}
