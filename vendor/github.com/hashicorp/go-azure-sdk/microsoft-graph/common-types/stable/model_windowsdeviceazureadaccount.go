package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsDeviceAccount = WindowsDeviceAzureADAccount{}

type WindowsDeviceAzureADAccount struct {
	// Not yet documented
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// Fields inherited from WindowsDeviceAccount

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Not yet documented
	Password nullable.Type[string] `json:"password,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsDeviceAzureADAccount) WindowsDeviceAccount() BaseWindowsDeviceAccountImpl {
	return BaseWindowsDeviceAccountImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Password:  s.Password,
	}
}

var _ json.Marshaler = WindowsDeviceAzureADAccount{}

func (s WindowsDeviceAzureADAccount) MarshalJSON() ([]byte, error) {
	type wrapper WindowsDeviceAzureADAccount
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsDeviceAzureADAccount: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsDeviceAzureADAccount: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsDeviceAzureADAccount"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsDeviceAzureADAccount: %+v", err)
	}

	return encoded, nil
}
