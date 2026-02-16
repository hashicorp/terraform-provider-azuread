package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUser interface {
	Entity
	RiskyUser() BaseRiskyUserImpl
}

var _ RiskyUser = BaseRiskyUserImpl{}

type BaseRiskyUserImpl struct {
	History *[]RiskyUserHistoryItem `json:"history,omitempty"`

	// Indicates whether the user is deleted. Possible values are: true, false.
	IsDeleted nullable.Type[bool] `json:"isDeleted,omitempty"`

	// Indicates whether a user's risky state is being processed by the backend.
	IsProcessing nullable.Type[bool] `json:"isProcessing,omitempty"`

	// The possible values are none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange,
	// userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe,
	// userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden,
	// adminConfirmedUserCompromised, unknownFutureValue, adminConfirmedServicePrincipalCompromised,
	// adminDismissedAllRiskForServicePrincipal, m365DAdminDismissedDetection, userChangedPasswordOnPremises,
	// adminDismissedRiskForSignIn, adminConfirmedAccountSafe. Use the Prefer: include-unknown-enum-members request header
	// to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised,
	// adminDismissedAllRiskForServicePrincipal, m365DAdminDismissedDetection, userChangedPasswordOnPremises,
	// adminDismissedRiskForSignIn, adminConfirmedAccountSafe.
	RiskDetail *RiskDetail `json:"riskDetail,omitempty"`

	// The date and time that the risky user was last updated. The DateTimeOffset type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	RiskLastUpdatedDateTime nullable.Type[string] `json:"riskLastUpdatedDateTime,omitempty"`

	// Level of the detected risky user. The possible values are low, medium, high, hidden, none, unknownFutureValue.
	RiskLevel *RiskLevel `json:"riskLevel,omitempty"`

	// State of the user's risk. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk,
	// confirmedCompromised, unknownFutureValue.
	RiskState *RiskState `json:"riskState,omitempty"`

	// Risky user display name.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// Risky user principal name.
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

func (s BaseRiskyUserImpl) RiskyUser() BaseRiskyUserImpl {
	return s
}

func (s BaseRiskyUserImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ RiskyUser = RawRiskyUserImpl{}

// RawRiskyUserImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRiskyUserImpl struct {
	riskyUser BaseRiskyUserImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawRiskyUserImpl) RiskyUser() BaseRiskyUserImpl {
	return s.riskyUser
}

func (s RawRiskyUserImpl) Entity() BaseEntityImpl {
	return s.riskyUser.Entity()
}

var _ json.Marshaler = BaseRiskyUserImpl{}

func (s BaseRiskyUserImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseRiskyUserImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseRiskyUserImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseRiskyUserImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.riskyUser"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseRiskyUserImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalRiskyUserImplementation(input []byte) (RiskyUser, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RiskyUser into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.riskyUserHistoryItem") {
		var out RiskyUserHistoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RiskyUserHistoryItem: %+v", err)
		}
		return out, nil
	}

	var parent BaseRiskyUserImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRiskyUserImpl: %+v", err)
	}

	return RawRiskyUserImpl{
		riskyUser: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
