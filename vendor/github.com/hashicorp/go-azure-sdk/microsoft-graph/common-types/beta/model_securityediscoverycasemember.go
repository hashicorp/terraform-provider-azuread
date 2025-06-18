package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityEdiscoveryCaseMember{}

type SecurityEdiscoveryCaseMember struct {
	// The display name of the eDiscovery case member. Allowed only for case members of type roleGroup.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Specifies the recipient type of the eDiscovery case member. The possible values are: user, roleGroup,
	// unknownFutureValue.
	RecipientType *SecurityRecipientType `json:"recipientType,omitempty"`

	// The smtp address of the eDiscovery case member. Allowed only for case members of type user.
	SmtpAddress nullable.Type[string] `json:"smtpAddress,omitempty"`

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

func (s SecurityEdiscoveryCaseMember) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryCaseMember{}

func (s SecurityEdiscoveryCaseMember) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryCaseMember
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryCaseMember: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryCaseMember: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryCaseMember"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryCaseMember: %+v", err)
	}

	return encoded, nil
}
