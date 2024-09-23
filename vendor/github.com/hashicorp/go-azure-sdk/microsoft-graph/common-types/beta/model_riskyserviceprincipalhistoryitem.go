package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RiskyServicePrincipal = RiskyServicePrincipalHistoryItem{}

type RiskyServicePrincipalHistoryItem struct {
	// The activity related to service principal risk level change.
	Activity *RiskServicePrincipalActivity `json:"activity,omitempty"`

	// The identifier of the actor of the operation.
	InitiatedBy nullable.Type[string] `json:"initiatedBy,omitempty"`

	// The identifier of the service principal.
	ServicePrincipalId nullable.Type[string] `json:"servicePrincipalId,omitempty"`

	// Fields inherited from RiskyServicePrincipal

	// true if the service principal account is enabled; otherwise, false.
	AccountEnabled nullable.Type[bool] `json:"accountEnabled,omitempty"`

	// The globally unique identifier for the associated application (its appId property), if any.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// The display name for the service principal.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Represents the risk history of Microsoft Entra service principals.
	History *[]RiskyServicePrincipalHistoryItem `json:"history,omitempty"`

	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// Indicates whether Microsoft Entra ID Protection is currently processing the service principal's risky state.
	IsProcessing nullable.Type[bool] `json:"isProcessing,omitempty"`

	// Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium
	// customers. Events in tenants without this license will be returned hidden. The possible values are: none, hidden,
	// unknownFutureValue, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that
	// you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable
	// enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
	RiskDetail *RiskDetail `json:"riskDetail,omitempty"`

	// The date and time that the risk state was last updated. The DateTimeOffset type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2021 is 2021-01-01T00:00:00Z.
	// Supports $filter (eq).
	RiskLastUpdatedDateTime nullable.Type[string] `json:"riskLastUpdatedDateTime,omitempty"`

	// Level of the detected risky workload identity. The possible values are: low, medium, high, hidden, none,
	// unknownFutureValue. Supports $filter (eq).
	RiskLevel *RiskLevel `json:"riskLevel,omitempty"`

	// State of the service principal's risk. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk,
	// confirmedCompromised, unknownFutureValue.
	RiskState *RiskState `json:"riskState,omitempty"`

	// Identifies whether the service principal represents an Application, a ManagedIdentity, or a legacy application
	// (socialIdp). This is set by Microsoft Entra ID internally and is inherited from servicePrincipal.
	ServicePrincipalType nullable.Type[string] `json:"servicePrincipalType,omitempty"`

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

func (s RiskyServicePrincipalHistoryItem) RiskyServicePrincipal() BaseRiskyServicePrincipalImpl {
	return BaseRiskyServicePrincipalImpl{
		AccountEnabled:          s.AccountEnabled,
		AppId:                   s.AppId,
		DisplayName:             s.DisplayName,
		History:                 s.History,
		IsEnabled:               s.IsEnabled,
		IsProcessing:            s.IsProcessing,
		RiskDetail:              s.RiskDetail,
		RiskLastUpdatedDateTime: s.RiskLastUpdatedDateTime,
		RiskLevel:               s.RiskLevel,
		RiskState:               s.RiskState,
		ServicePrincipalType:    s.ServicePrincipalType,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s RiskyServicePrincipalHistoryItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RiskyServicePrincipalHistoryItem{}

func (s RiskyServicePrincipalHistoryItem) MarshalJSON() ([]byte, error) {
	type wrapper RiskyServicePrincipalHistoryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RiskyServicePrincipalHistoryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RiskyServicePrincipalHistoryItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.riskyServicePrincipalHistoryItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RiskyServicePrincipalHistoryItem: %+v", err)
	}

	return encoded, nil
}
