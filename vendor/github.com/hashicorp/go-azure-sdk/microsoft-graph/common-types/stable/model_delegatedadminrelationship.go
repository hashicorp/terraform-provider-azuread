package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationship interface {
	Entity
	DelegatedAdminRelationship() BaseDelegatedAdminRelationshipImpl
}

var _ DelegatedAdminRelationship = BaseDelegatedAdminRelationshipImpl{}

type BaseDelegatedAdminRelationshipImpl struct {
	// The access assignments associated with the delegated admin relationship.
	AccessAssignments *[]DelegatedAdminAccessAssignment `json:"accessAssignments,omitempty"`

	AccessDetails *DelegatedAdminAccessDetails `json:"accessDetails,omitempty"`

	// The date and time in ISO 8601 format and in UTC time when the relationship became active. Read-only.
	ActivatedDateTime nullable.Type[string] `json:"activatedDateTime,omitempty"`

	// The duration by which the validity of the relationship is automatically extended, denoted in ISO 8601 format.
	// Supported values are: P0D, PT0S, P180D. The default value is PT0S. PT0S indicates that the relationship expires when
	// the endDateTime is reached and it isn't automatically extended.
	AutoExtendDuration nullable.Type[string] `json:"autoExtendDuration,omitempty"`

	// The date and time in ISO 8601 format and in UTC time when the relationship was created. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The display name and unique identifier of the customer of the relationship. This is configured either by the partner
	// at the time the relationship is created or by the system after the customer approves the relationship. Can't be
	// changed by the customer.
	Customer *DelegatedAdminRelationshipCustomerParticipant `json:"customer,omitempty"`

	// The display name of the relationship used for ease of identification. Must be unique across all delegated admin
	// relationships of the partner and is set by the partner only when the relationship is in the created status and can't
	// be changed by the customer. Maximum length is 50 characters.
	DisplayName *string `json:"displayName,omitempty"`

	// The duration of the relationship in ISO 8601 format. Must be a value between P1D and P2Y inclusive. This is set by
	// the partner only when the relationship is in the created status and can't be changed by the customer.
	Duration *string `json:"duration,omitempty"`

	// The date and time in ISO 8601 format and in UTC time when the status of relationship changes to either terminated or
	// expired. Calculated as endDateTime = activatedDateTime + duration. Read-only.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The date and time in ISO 8601 format and in UTC time when the relationship was last modified. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The long running operations associated with the delegated admin relationship.
	Operations *[]DelegatedAdminRelationshipOperation `json:"operations,omitempty"`

	// The requests associated with the delegated admin relationship.
	Requests *[]DelegatedAdminRelationshipRequest `json:"requests,omitempty"`

	// The status of the relationship. Read Only. The possible values are: activating, active, approvalPending, approved,
	// created, expired, expiring, terminated, terminating, terminationRequested, unknownFutureValue. Supports $orderby.
	Status *DelegatedAdminRelationshipStatus `json:"status,omitempty"`

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

func (s BaseDelegatedAdminRelationshipImpl) DelegatedAdminRelationship() BaseDelegatedAdminRelationshipImpl {
	return s
}

func (s BaseDelegatedAdminRelationshipImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DelegatedAdminRelationship = RawDelegatedAdminRelationshipImpl{}

// RawDelegatedAdminRelationshipImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDelegatedAdminRelationshipImpl struct {
	delegatedAdminRelationship BaseDelegatedAdminRelationshipImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawDelegatedAdminRelationshipImpl) DelegatedAdminRelationship() BaseDelegatedAdminRelationshipImpl {
	return s.delegatedAdminRelationship
}

func (s RawDelegatedAdminRelationshipImpl) Entity() BaseEntityImpl {
	return s.delegatedAdminRelationship.Entity()
}

var _ json.Marshaler = BaseDelegatedAdminRelationshipImpl{}

func (s BaseDelegatedAdminRelationshipImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDelegatedAdminRelationshipImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDelegatedAdminRelationshipImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDelegatedAdminRelationshipImpl: %+v", err)
	}

	delete(decoded, "activatedDateTime")
	delete(decoded, "createdDateTime")
	delete(decoded, "endDateTime")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.delegatedAdminRelationship"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDelegatedAdminRelationshipImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDelegatedAdminRelationshipImplementation(input []byte) (DelegatedAdminRelationship, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DelegatedAdminRelationship into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.resellerDelegatedAdminRelationship") {
		var out ResellerDelegatedAdminRelationship
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResellerDelegatedAdminRelationship: %+v", err)
		}
		return out, nil
	}

	var parent BaseDelegatedAdminRelationshipImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDelegatedAdminRelationshipImpl: %+v", err)
	}

	return RawDelegatedAdminRelationshipImpl{
		delegatedAdminRelationship: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
