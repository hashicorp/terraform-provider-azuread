package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamsAdministrationTeamsUserConfiguration{}

type TeamsAdministrationTeamsUserConfiguration struct {
	// The type of the account in the Teams context. The possible values are: user, resourceAccount, guest, sfbOnPremUser,
	// unknown, unknownFutureValue, ineligibleUser. Use the Prefer: include-unknown-enum-members request header to get the
	// following value from this enum evolvable enum: ineligibleUser.
	AccountType *TeamsAdministrationAccountType `json:"accountType,omitempty"`

	// The date and time when the user was created. The timestamp represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Contains the user's effective policy assignments, with each assignment including policyType and policyAssignment
	// details.
	EffectivePolicyAssignments *[]TeamsAdministrationEffectivePolicyAssignment `json:"effectivePolicyAssignments,omitempty"`

	// The Teams features enabled for a given user based on licensing or service plan.
	FeatureTypes *[]string `json:"featureTypes,omitempty"`

	// Indicates whether voice capability is enabled.
	IsEnterpriseVoiceEnabled nullable.Type[bool] `json:"isEnterpriseVoiceEnabled,omitempty"`

	// The date and time when the user's details were last modified. The system updates this value each time the user's
	// details are changed. The timestamp represents date and time information using ISO 8601 format and is always in UTC
	// time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ModifiedDateTime *string `json:"modifiedDateTime,omitempty"`

	// Includes both the phone number and its corresponding assignment category. The assignment category can include values
	// such as primary, private, and alternate.
	TelephoneNumbers *[]TeamsAdministrationAssignedTelephoneNumber `json:"telephoneNumbers,omitempty"`

	// The unique identifier of the tenant in Entra to which this user is assigned.
	TenantId *string `json:"tenantId,omitempty"`

	// Represents an Entra user account.
	User *User `json:"user,omitempty"`

	// The sign-in address of the user.
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

func (s TeamsAdministrationTeamsUserConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsAdministrationTeamsUserConfiguration{}

func (s TeamsAdministrationTeamsUserConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper TeamsAdministrationTeamsUserConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsAdministrationTeamsUserConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAdministrationTeamsUserConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamsAdministration.teamsUserConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsAdministrationTeamsUserConfiguration: %+v", err)
	}

	return encoded, nil
}
