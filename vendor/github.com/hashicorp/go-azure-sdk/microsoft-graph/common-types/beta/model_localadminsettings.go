package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocalAdminSettings struct {
	// Indicates whether global administrators are local administrators on all Microsoft Entra-joined devices. This setting
	// only applies to future registrations. Default is true.
	EnableGlobalAdmins nullable.Type[bool] `json:"enableGlobalAdmins,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Determines the users and groups that become local administrators on Microsoft Entra joined devices that they
	// register.
	RegisteringUsers DeviceRegistrationMembership `json:"registeringUsers"`
}

var _ json.Unmarshaler = &LocalAdminSettings{}

func (s *LocalAdminSettings) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EnableGlobalAdmins nullable.Type[bool] `json:"enableGlobalAdmins,omitempty"`
		ODataId            *string             `json:"@odata.id,omitempty"`
		ODataType          *string             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EnableGlobalAdmins = decoded.EnableGlobalAdmins
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling LocalAdminSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["registeringUsers"]; ok {
		impl, err := UnmarshalDeviceRegistrationMembershipImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RegisteringUsers' for 'LocalAdminSettings': %+v", err)
		}
		s.RegisteringUsers = impl
	}

	return nil
}
