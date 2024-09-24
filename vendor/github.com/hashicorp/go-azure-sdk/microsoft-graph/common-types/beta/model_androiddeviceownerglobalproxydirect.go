package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AndroidDeviceOwnerGlobalProxy = AndroidDeviceOwnerGlobalProxyDirect{}

type AndroidDeviceOwnerGlobalProxyDirect struct {
	// The excluded hosts
	ExcludedHosts *[]string `json:"excludedHosts,omitempty"`

	// The host name
	Host *string `json:"host,omitempty"`

	// The port
	Port *int64 `json:"port,omitempty"`

	// Fields inherited from AndroidDeviceOwnerGlobalProxy

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AndroidDeviceOwnerGlobalProxyDirect) AndroidDeviceOwnerGlobalProxy() BaseAndroidDeviceOwnerGlobalProxyImpl {
	return BaseAndroidDeviceOwnerGlobalProxyImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidDeviceOwnerGlobalProxyDirect{}

func (s AndroidDeviceOwnerGlobalProxyDirect) MarshalJSON() ([]byte, error) {
	type wrapper AndroidDeviceOwnerGlobalProxyDirect
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidDeviceOwnerGlobalProxyDirect: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerGlobalProxyDirect: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerGlobalProxyDirect"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidDeviceOwnerGlobalProxyDirect: %+v", err)
	}

	return encoded, nil
}
