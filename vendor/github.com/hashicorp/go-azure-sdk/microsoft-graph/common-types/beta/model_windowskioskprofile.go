package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskProfile struct {
	// The app base class used to identify the application info for the kiosk configuration
	AppConfiguration WindowsKioskAppConfiguration `json:"appConfiguration"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Key of the entity.
	ProfileId *string `json:"profileId,omitempty"`

	// This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the
	// users to whom this kiosk configuration is assigned.
	ProfileName *string `json:"profileName,omitempty"`

	// The user accounts that will be locked to this kiosk configuration. This collection can contain a maximum of 100
	// elements.
	UserAccountsConfiguration *[]WindowsKioskUser `json:"userAccountsConfiguration,omitempty"`
}

var _ json.Unmarshaler = &WindowsKioskProfile{}

func (s *WindowsKioskProfile) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId     *string `json:"@odata.id,omitempty"`
		ODataType   *string `json:"@odata.type,omitempty"`
		ProfileId   *string `json:"profileId,omitempty"`
		ProfileName *string `json:"profileName,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ProfileId = decoded.ProfileId
	s.ProfileName = decoded.ProfileName

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsKioskProfile into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appConfiguration"]; ok {
		impl, err := UnmarshalWindowsKioskAppConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AppConfiguration' for 'WindowsKioskProfile': %+v", err)
		}
		s.AppConfiguration = impl
	}

	if v, ok := temp["userAccountsConfiguration"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling UserAccountsConfiguration into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsKioskUser, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsKioskUserImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'UserAccountsConfiguration' for 'WindowsKioskProfile': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.UserAccountsConfiguration = &output
	}

	return nil
}
