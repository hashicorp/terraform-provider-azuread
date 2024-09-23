package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InboundSharedUserProfile struct {
	// The name displayed in the address book for the user at the time when the sharing record was created. Read-only.
	DisplayName *string `json:"displayName,omitempty"`

	// The home tenant id of the external user. Read-only.
	HomeTenantId *string `json:"homeTenantId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The object id of the external user. Read-only.
	UserId *string `json:"userId,omitempty"`

	// The user principal name (UPN) of the external user. Read-only.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

var _ json.Marshaler = InboundSharedUserProfile{}

func (s InboundSharedUserProfile) MarshalJSON() ([]byte, error) {
	type wrapper InboundSharedUserProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InboundSharedUserProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InboundSharedUserProfile: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "homeTenantId")
	delete(decoded, "userId")
	delete(decoded, "userPrincipalName")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InboundSharedUserProfile: %+v", err)
	}

	return encoded, nil
}
