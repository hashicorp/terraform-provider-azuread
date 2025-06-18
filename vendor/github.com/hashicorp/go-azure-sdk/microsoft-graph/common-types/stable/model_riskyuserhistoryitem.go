package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RiskyUser = RiskyUserHistoryItem{}

type RiskyUserHistoryItem struct {
	// The activity related to user risk level change.
	Activity *RiskUserActivity `json:"activity,omitempty"`

	// The ID of actor that does the operation.
	InitiatedBy nullable.Type[string] `json:"initiatedBy,omitempty"`

	// The ID of the user.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Fields inherited from RiskyUser

	// The activity related to user risk level change
	History *[]RiskyUserHistoryItem `json:"history,omitempty"`

	// Indicates whether the user is deleted. Possible values are: true, false.
	IsDeleted nullable.Type[bool] `json:"isDeleted,omitempty"`

	// Indicates whether the backend is processing a user's risky state.
	IsProcessing nullable.Type[bool] `json:"isProcessing,omitempty"`

	// The possible values are none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange,
	// userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe,
	// userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden,
	// adminConfirmedUserCompromised, unknownFutureValue, adminConfirmedServicePrincipalCompromised,
	// adminDismissedAllRiskForServicePrincipal, m365DAdminDismissedDetection, userChangedPasswordOnPremises,
	// adminDismissedRiskForSignIn, adminConfirmedAccountSafe. Use the Prefer: include-unknown-enum-members request header
	// to get the following value or values in this evolvable enum: adminConfirmedServicePrincipalCompromised,
	// adminDismissedAllRiskForServicePrincipal, m365DAdminDismissedDetection, userChangedPasswordOnPremises,
	// adminDismissedRiskForSignIn, adminConfirmedAccountSafe.
	RiskDetail *RiskDetail `json:"riskDetail,omitempty"`

	// The date and time that the risky user was last updated. The DateTimeOffset type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	RiskLastUpdatedDateTime nullable.Type[string] `json:"riskLastUpdatedDateTime,omitempty"`

	// Level of the detected risky user. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
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

func (s RiskyUserHistoryItem) RiskyUser() BaseRiskyUserImpl {
	return BaseRiskyUserImpl{
		History:                 s.History,
		IsDeleted:               s.IsDeleted,
		IsProcessing:            s.IsProcessing,
		RiskDetail:              s.RiskDetail,
		RiskLastUpdatedDateTime: s.RiskLastUpdatedDateTime,
		RiskLevel:               s.RiskLevel,
		RiskState:               s.RiskState,
		UserDisplayName:         s.UserDisplayName,
		UserPrincipalName:       s.UserPrincipalName,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s RiskyUserHistoryItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RiskyUserHistoryItem{}

func (s RiskyUserHistoryItem) MarshalJSON() ([]byte, error) {
	type wrapper RiskyUserHistoryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RiskyUserHistoryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RiskyUserHistoryItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.riskyUserHistoryItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RiskyUserHistoryItem: %+v", err)
	}

	return encoded, nil
}
