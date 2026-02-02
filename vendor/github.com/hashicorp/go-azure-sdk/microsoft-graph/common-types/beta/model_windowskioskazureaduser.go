package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsKioskUser = WindowsKioskAzureADUser{}

type WindowsKioskAzureADUser struct {
	// The ID of the AzureAD user that will be locked to this kiosk configuration
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The user accounts that will be locked to this kiosk configuration
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`

	// Fields inherited from WindowsKioskUser

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsKioskAzureADUser) WindowsKioskUser() BaseWindowsKioskUserImpl {
	return BaseWindowsKioskUserImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsKioskAzureADUser{}

func (s WindowsKioskAzureADUser) MarshalJSON() ([]byte, error) {
	type wrapper WindowsKioskAzureADUser
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsKioskAzureADUser: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskAzureADUser: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsKioskAzureADUser"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsKioskAzureADUser: %+v", err)
	}

	return encoded, nil
}
