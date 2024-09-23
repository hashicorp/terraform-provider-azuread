package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AwsPolicy{}

type AwsPolicy struct {
	AwsPolicyType *AwsPolicyType `json:"awsPolicyType,omitempty"`

	// The display name for the AWS policy. Read-only. Supports $filter and (eq,contains).
	DisplayName *string `json:"displayName,omitempty"`

	// The base64 encoded identifier for the AWS policy as defined by AWS. Read-only. Alternate key. Supports $filter and
	// eq.
	ExternalId *string `json:"externalId,omitempty"`

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

func (s AwsPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AwsPolicy{}

func (s AwsPolicy) MarshalJSON() ([]byte, error) {
	type wrapper AwsPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsPolicy: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "externalId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsPolicy: %+v", err)
	}

	return encoded, nil
}
