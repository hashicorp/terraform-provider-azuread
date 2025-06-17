package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = HardwareOathTokenAuthenticationMethodDeviceCollectionResponse{}

type HardwareOathTokenAuthenticationMethodDeviceCollectionResponse struct {
	Value *[]HardwareOathTokenAuthenticationMethodDevice `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s HardwareOathTokenAuthenticationMethodDeviceCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = HardwareOathTokenAuthenticationMethodDeviceCollectionResponse{}

func (s HardwareOathTokenAuthenticationMethodDeviceCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper HardwareOathTokenAuthenticationMethodDeviceCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HardwareOathTokenAuthenticationMethodDeviceCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HardwareOathTokenAuthenticationMethodDeviceCollectionResponse: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.hardwareOathTokenAuthenticationMethodDeviceCollectionResponse"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HardwareOathTokenAuthenticationMethodDeviceCollectionResponse: %+v", err)
	}

	return encoded, nil
}
