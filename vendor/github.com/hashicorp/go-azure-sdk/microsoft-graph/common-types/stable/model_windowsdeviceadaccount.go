package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsDeviceAccount = WindowsDeviceADAccount{}

type WindowsDeviceADAccount struct {
	// Not yet documented
	DomainName nullable.Type[string] `json:"domainName,omitempty"`

	// Not yet documented
	UserName nullable.Type[string] `json:"userName,omitempty"`

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

func (s WindowsDeviceADAccount) WindowsDeviceAccount() BaseWindowsDeviceAccountImpl {
	return BaseWindowsDeviceAccountImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Password:  s.Password,
	}
}

var _ json.Marshaler = WindowsDeviceADAccount{}

func (s WindowsDeviceADAccount) MarshalJSON() ([]byte, error) {
	type wrapper WindowsDeviceADAccount
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsDeviceADAccount: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsDeviceADAccount: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsDeviceADAccount"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsDeviceADAccount: %+v", err)
	}

	return encoded, nil
}
