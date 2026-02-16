package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SubscribedSku{}

type SubscribedSku struct {
	// The unique ID of the account this SKU belongs to.
	AccountId nullable.Type[string] `json:"accountId,omitempty"`

	// The name of the account this SKU belongs to.
	AccountName nullable.Type[string] `json:"accountName,omitempty"`

	// The target class for this SKU. Only SKUs with target class User are assignable. Possible values are: User, Company.
	AppliesTo nullable.Type[string] `json:"appliesTo,omitempty"`

	// Enabled indicates that the prepaidUnits property has at least one unit that is enabled. LockedOut indicates that the
	// customer canceled their subscription. Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut.
	CapabilityStatus nullable.Type[string] `json:"capabilityStatus,omitempty"`

	// The number of licenses that have been assigned.
	ConsumedUnits nullable.Type[int64] `json:"consumedUnits,omitempty"`

	// Information about the number and status of prepaid licenses.
	PrepaidUnits *LicenseUnitsDetail `json:"prepaidUnits,omitempty"`

	// Information about the service plans that are available with the SKU. Not nullable
	ServicePlans *[]ServicePlanInfo `json:"servicePlans,omitempty"`

	// The unique identifier (GUID) for the service SKU.
	SkuId nullable.Type[string] `json:"skuId,omitempty"`

	// The SKU part number; for example, AAD_PREMIUM or RMSBASIC. To get a list of commercial subscriptions that an
	// organization has acquired, see List subscribedSkus.
	SkuPartNumber nullable.Type[string] `json:"skuPartNumber,omitempty"`

	// A list of all subscription IDs associated with this SKU.
	SubscriptionIds *[]string `json:"subscriptionIds,omitempty"`

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

func (s SubscribedSku) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SubscribedSku{}

func (s SubscribedSku) MarshalJSON() ([]byte, error) {
	type wrapper SubscribedSku
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SubscribedSku: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SubscribedSku: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.subscribedSku"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SubscribedSku: %+v", err)
	}

	return encoded, nil
}
