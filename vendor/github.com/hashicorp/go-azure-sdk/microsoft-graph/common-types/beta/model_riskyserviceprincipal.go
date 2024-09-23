package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipal interface {
	Entity
	RiskyServicePrincipal() BaseRiskyServicePrincipalImpl
}

var _ RiskyServicePrincipal = BaseRiskyServicePrincipalImpl{}

type BaseRiskyServicePrincipalImpl struct {
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

func (s BaseRiskyServicePrincipalImpl) RiskyServicePrincipal() BaseRiskyServicePrincipalImpl {
	return s
}

func (s BaseRiskyServicePrincipalImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ RiskyServicePrincipal = RawRiskyServicePrincipalImpl{}

// RawRiskyServicePrincipalImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRiskyServicePrincipalImpl struct {
	riskyServicePrincipal BaseRiskyServicePrincipalImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawRiskyServicePrincipalImpl) RiskyServicePrincipal() BaseRiskyServicePrincipalImpl {
	return s.riskyServicePrincipal
}

func (s RawRiskyServicePrincipalImpl) Entity() BaseEntityImpl {
	return s.riskyServicePrincipal.Entity()
}

var _ json.Marshaler = BaseRiskyServicePrincipalImpl{}

func (s BaseRiskyServicePrincipalImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseRiskyServicePrincipalImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseRiskyServicePrincipalImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseRiskyServicePrincipalImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.riskyServicePrincipal"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseRiskyServicePrincipalImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalRiskyServicePrincipalImplementation(input []byte) (RiskyServicePrincipal, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RiskyServicePrincipal into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.riskyServicePrincipalHistoryItem") {
		var out RiskyServicePrincipalHistoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RiskyServicePrincipalHistoryItem: %+v", err)
		}
		return out, nil
	}

	var parent BaseRiskyServicePrincipalImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRiskyServicePrincipalImpl: %+v", err)
	}

	return RawRiskyServicePrincipalImpl{
		riskyServicePrincipal: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
