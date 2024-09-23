package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutboundSharedUserProfile struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The collection of external Microsoft Entra tenants that the user shared profile data with. Read-only.
	Tenants *[]TenantReference `json:"tenants,omitempty"`

	// The object id of the external user. Read-only.
	UserId *string `json:"userId,omitempty"`
}

var _ json.Marshaler = OutboundSharedUserProfile{}

func (s OutboundSharedUserProfile) MarshalJSON() ([]byte, error) {
	type wrapper OutboundSharedUserProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OutboundSharedUserProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OutboundSharedUserProfile: %+v", err)
	}

	delete(decoded, "tenants")
	delete(decoded, "userId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OutboundSharedUserProfile: %+v", err)
	}

	return encoded, nil
}
