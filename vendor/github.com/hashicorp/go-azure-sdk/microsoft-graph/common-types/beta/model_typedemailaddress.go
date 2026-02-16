package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EmailAddress = TypedEmailAddress{}

type TypedEmailAddress struct {
	// To specify a custom type of email address, set type to other, and assign otherLabel to a custom string. For example,
	// you may use a specific email address for your volunteer activities. Set type to other, and set otherLabel to a custom
	// string such as Volunteer work.
	OtherLabel nullable.Type[string] `json:"otherLabel,omitempty"`

	// The type of email address. Possible values are: unknown, work, personal, main, other. The default value is unknown,
	// which means address has not been set as a specific type.
	Type *EmailType `json:"type,omitempty"`

	// Fields inherited from EmailAddress

	// The email address of an entity instance.
	Address nullable.Type[string] `json:"address,omitempty"`

	// The display name of an entity instance.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TypedEmailAddress) EmailAddress() BaseEmailAddressImpl {
	return BaseEmailAddressImpl{
		Address:   s.Address,
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TypedEmailAddress{}

func (s TypedEmailAddress) MarshalJSON() ([]byte, error) {
	type wrapper TypedEmailAddress
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TypedEmailAddress: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TypedEmailAddress: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.typedEmailAddress"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TypedEmailAddress: %+v", err)
	}

	return encoded, nil
}
