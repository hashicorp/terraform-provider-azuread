package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccountTargetContent = AddressBookAccountTargetContent{}

type AddressBookAccountTargetContent struct {
	// List of user emails targeted for an attack simulation training campaign.
	AccountTargetEmails *[]string `json:"accountTargetEmails,omitempty"`

	// Fields inherited from AccountTargetContent

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of account target content. Possible values are: unknown,includeAll, addressBook, unknownFutureValue.
	Type *AccountTargetContentType `json:"type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AddressBookAccountTargetContent) AccountTargetContent() BaseAccountTargetContentImpl {
	return BaseAccountTargetContentImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Type:      s.Type,
	}
}

var _ json.Marshaler = AddressBookAccountTargetContent{}

func (s AddressBookAccountTargetContent) MarshalJSON() ([]byte, error) {
	type wrapper AddressBookAccountTargetContent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AddressBookAccountTargetContent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AddressBookAccountTargetContent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.addressBookAccountTargetContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AddressBookAccountTargetContent: %+v", err)
	}

	return encoded, nil
}
