package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AndroidDeviceOwnerGlobalProxy = AndroidDeviceOwnerGlobalProxyAutoConfig{}

type AndroidDeviceOwnerGlobalProxyAutoConfig struct {
	// The proxy auto-config URL
	ProxyAutoConfigURL *string `json:"proxyAutoConfigURL,omitempty"`

	// Fields inherited from AndroidDeviceOwnerGlobalProxy

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AndroidDeviceOwnerGlobalProxyAutoConfig) AndroidDeviceOwnerGlobalProxy() BaseAndroidDeviceOwnerGlobalProxyImpl {
	return BaseAndroidDeviceOwnerGlobalProxyImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidDeviceOwnerGlobalProxyAutoConfig{}

func (s AndroidDeviceOwnerGlobalProxyAutoConfig) MarshalJSON() ([]byte, error) {
	type wrapper AndroidDeviceOwnerGlobalProxyAutoConfig
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidDeviceOwnerGlobalProxyAutoConfig: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerGlobalProxyAutoConfig: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerGlobalProxyAutoConfig"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidDeviceOwnerGlobalProxyAutoConfig: %+v", err)
	}

	return encoded, nil
}
