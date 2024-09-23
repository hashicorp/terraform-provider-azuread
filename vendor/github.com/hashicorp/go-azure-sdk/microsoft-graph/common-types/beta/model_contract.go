package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = Contract{}

type Contract struct {
	// Type of contract. Possible values are: SyndicationPartner, BreadthPartner, ResellerPartner. See more in the table
	// below.
	ContractType nullable.Type[string] `json:"contractType,omitempty"`

	// The unique identifier for the customer tenant referenced by this partnership. Corresponds to the id property of the
	// customer tenant's organization resource.
	CustomerId nullable.Type[string] `json:"customerId,omitempty"`

	// A copy of the customer tenant's default domain name. The copy is made when the partnership with the customer is
	// established. It isn't automatically updated if the customer tenant's default domain name changes.
	DefaultDomainName nullable.Type[string] `json:"defaultDomainName,omitempty"`

	// A copy of the customer tenant's display name. The copy is made when the partnership with the customer is established.
	// It isn't automatically updated if the customer tenant's display name changes.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s Contract) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s Contract) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Contract{}

func (s Contract) MarshalJSON() ([]byte, error) {
	type wrapper Contract
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Contract: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Contract: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.contract"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Contract: %+v", err)
	}

	return encoded, nil
}
