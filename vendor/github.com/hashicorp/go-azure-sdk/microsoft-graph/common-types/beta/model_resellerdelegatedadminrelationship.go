package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DelegatedAdminRelationship = ResellerDelegatedAdminRelationship{}

type ResellerDelegatedAdminRelationship struct {
	// The tenant ID of the indirect provider partner who created the relationship for the indirect reseller partner.
	IndirectProviderTenantId *string `json:"indirectProviderTenantId,omitempty"`

	// Indicates the indirect reseller partner consent status. true indicates that the partner has yet to review the
	// relationship; false indicates that the partner has already provided consent by approving or rejecting the
	// relationship.
	IsPartnerConsentPending *bool `json:"isPartnerConsentPending,omitempty"`

	// Fields inherited from DelegatedAdminRelationship

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
	// relationships of the partner. This is set by the partner only when the relationship is in the created status and
	// can't be changed by the customer. Maximum length is 50 characters.
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

func (s ResellerDelegatedAdminRelationship) DelegatedAdminRelationship() BaseDelegatedAdminRelationshipImpl {
	return BaseDelegatedAdminRelationshipImpl{
		AccessAssignments:    s.AccessAssignments,
		AccessDetails:        s.AccessDetails,
		ActivatedDateTime:    s.ActivatedDateTime,
		AutoExtendDuration:   s.AutoExtendDuration,
		CreatedDateTime:      s.CreatedDateTime,
		Customer:             s.Customer,
		DisplayName:          s.DisplayName,
		Duration:             s.Duration,
		EndDateTime:          s.EndDateTime,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Operations:           s.Operations,
		Requests:             s.Requests,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s ResellerDelegatedAdminRelationship) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ResellerDelegatedAdminRelationship{}

func (s ResellerDelegatedAdminRelationship) MarshalJSON() ([]byte, error) {
	type wrapper ResellerDelegatedAdminRelationship
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ResellerDelegatedAdminRelationship: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ResellerDelegatedAdminRelationship: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.resellerDelegatedAdminRelationship"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ResellerDelegatedAdminRelationship: %+v", err)
	}

	return encoded, nil
}
