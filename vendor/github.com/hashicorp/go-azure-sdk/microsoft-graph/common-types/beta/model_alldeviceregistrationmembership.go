package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceRegistrationMembership = AllDeviceRegistrationMembership{}

type AllDeviceRegistrationMembership struct {

	// Fields inherited from DeviceRegistrationMembership

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AllDeviceRegistrationMembership) DeviceRegistrationMembership() BaseDeviceRegistrationMembershipImpl {
	return BaseDeviceRegistrationMembershipImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AllDeviceRegistrationMembership{}

func (s AllDeviceRegistrationMembership) MarshalJSON() ([]byte, error) {
	type wrapper AllDeviceRegistrationMembership
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AllDeviceRegistrationMembership: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AllDeviceRegistrationMembership: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.allDeviceRegistrationMembership"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AllDeviceRegistrationMembership: %+v", err)
	}

	return encoded, nil
}
