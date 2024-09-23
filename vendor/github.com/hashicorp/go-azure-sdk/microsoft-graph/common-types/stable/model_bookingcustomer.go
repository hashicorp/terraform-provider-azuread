package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BookingCustomerBase = BookingCustomer{}

type BookingCustomer struct {
	// Addresses associated with the customer. The attribute type of physicalAddress is not supported in v1.0. Internally we
	// map the addresses to the type others.
	Addresses *[]PhysicalAddress `json:"addresses,omitempty"`

	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The name of the customer.
	DisplayName *string `json:"displayName,omitempty"`

	// The SMTP address of the customer.
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// Phone numbers associated with the customer, including home, business and mobile numbers.
	Phones *[]Phone `json:"phones,omitempty"`

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

func (s BookingCustomer) BookingCustomerBase() BaseBookingCustomerBaseImpl {
	return BaseBookingCustomerBaseImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s BookingCustomer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BookingCustomer{}

func (s BookingCustomer) MarshalJSON() ([]byte, error) {
	type wrapper BookingCustomer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BookingCustomer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingCustomer: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.bookingCustomer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BookingCustomer: %+v", err)
	}

	return encoded, nil
}
