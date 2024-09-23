package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CompanySubscription{}

type CompanySubscription struct {
	// The ID of this subscription in the commerce system. Alternate key.
	CommerceSubscriptionId nullable.Type[string] `json:"commerceSubscriptionId,omitempty"`

	// The date and time when this subscription was created. The DateTimeOffset type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Whether the subscription is a free trial or purchased.
	IsTrial nullable.Type[bool] `json:"isTrial,omitempty"`

	// The date and time when the subscription will move to the next state (as defined by the status property) if not
	// renewed by the tenant. The DateTimeOffset type represents date and time information using ISO 8601 format and is
	// always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	NextLifecycleDateTime nullable.Type[string] `json:"nextLifecycleDateTime,omitempty"`

	// The object ID of the account admin.
	OwnerId nullable.Type[string] `json:"ownerId,omitempty"`

	// The unique identifier for the Microsoft partner tenant that created the subscription on a customer tenant.
	OwnerTenantId nullable.Type[string] `json:"ownerTenantId,omitempty"`

	// Indicates the entity that ownerId belongs to, for example, 'User'.
	OwnerType nullable.Type[string] `json:"ownerType,omitempty"`

	// The provisioning status of each service included in this subscription.
	ServiceStatus *[]ServicePlanInfo `json:"serviceStatus,omitempty"`

	// The object ID of the SKU associated with this subscription.
	SkuId nullable.Type[string] `json:"skuId,omitempty"`

	// The SKU associated with this subscription.
	SkuPartNumber nullable.Type[string] `json:"skuPartNumber,omitempty"`

	// The status of this subscription. Possible values are: Enabled, Deleted, Suspended, Warning, LockedOut.
	Status nullable.Type[string] `json:"status,omitempty"`

	// The number of licenses included in this subscription.
	TotalLicenses nullable.Type[int64] `json:"totalLicenses,omitempty"`

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

func (s CompanySubscription) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CompanySubscription{}

func (s CompanySubscription) MarshalJSON() ([]byte, error) {
	type wrapper CompanySubscription
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CompanySubscription: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CompanySubscription: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.companySubscription"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CompanySubscription: %+v", err)
	}

	return encoded, nil
}
