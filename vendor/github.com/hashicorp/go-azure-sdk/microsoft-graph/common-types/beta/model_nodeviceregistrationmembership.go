package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceRegistrationMembership = NoDeviceRegistrationMembership{}

type NoDeviceRegistrationMembership struct {

	// Fields inherited from DeviceRegistrationMembership

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NoDeviceRegistrationMembership) DeviceRegistrationMembership() BaseDeviceRegistrationMembershipImpl {
	return BaseDeviceRegistrationMembershipImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NoDeviceRegistrationMembership{}

func (s NoDeviceRegistrationMembership) MarshalJSON() ([]byte, error) {
	type wrapper NoDeviceRegistrationMembership
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NoDeviceRegistrationMembership: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NoDeviceRegistrationMembership: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.noDeviceRegistrationMembership"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NoDeviceRegistrationMembership: %+v", err)
	}

	return encoded, nil
}
