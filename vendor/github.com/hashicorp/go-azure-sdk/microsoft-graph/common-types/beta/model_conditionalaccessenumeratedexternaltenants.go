package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConditionalAccessExternalTenants = ConditionalAccessEnumeratedExternalTenants{}

type ConditionalAccessEnumeratedExternalTenants struct {

	// Fields inherited from ConditionalAccessExternalTenants

	Members *[]string `json:"members,omitempty"`

	// The membership kind. Possible values are: all, enumerated, unknownFutureValue. The enumerated member references an
	// conditionalAccessEnumeratedExternalTenants object.
	MembershipKind *ConditionalAccessExternalTenantsMembershipKind `json:"membershipKind,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ConditionalAccessEnumeratedExternalTenants) ConditionalAccessExternalTenants() BaseConditionalAccessExternalTenantsImpl {
	return BaseConditionalAccessExternalTenantsImpl{
		Members:        s.Members,
		MembershipKind: s.MembershipKind,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

var _ json.Marshaler = ConditionalAccessEnumeratedExternalTenants{}

func (s ConditionalAccessEnumeratedExternalTenants) MarshalJSON() ([]byte, error) {
	type wrapper ConditionalAccessEnumeratedExternalTenants
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConditionalAccessEnumeratedExternalTenants: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConditionalAccessEnumeratedExternalTenants: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.conditionalAccessEnumeratedExternalTenants"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConditionalAccessEnumeratedExternalTenants: %+v", err)
	}

	return encoded, nil
}
