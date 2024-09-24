package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindowsDeviceAccountActionParameter struct {
	CalendarSyncEnabled nullable.Type[bool]   `json:"calendarSyncEnabled,omitempty"`
	DeviceAccount       WindowsDeviceAccount  `json:"deviceAccount"`
	DeviceAccountEmail  nullable.Type[string] `json:"deviceAccountEmail,omitempty"`
	ExchangeServer      nullable.Type[string] `json:"exchangeServer,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PasswordRotationEnabled          nullable.Type[bool]   `json:"passwordRotationEnabled,omitempty"`
	SessionInitiationProtocalAddress nullable.Type[string] `json:"sessionInitiationProtocalAddress,omitempty"`
}

var _ json.Unmarshaler = &UpdateWindowsDeviceAccountActionParameter{}

func (s *UpdateWindowsDeviceAccountActionParameter) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CalendarSyncEnabled              nullable.Type[bool]   `json:"calendarSyncEnabled,omitempty"`
		DeviceAccountEmail               nullable.Type[string] `json:"deviceAccountEmail,omitempty"`
		ExchangeServer                   nullable.Type[string] `json:"exchangeServer,omitempty"`
		ODataId                          *string               `json:"@odata.id,omitempty"`
		ODataType                        *string               `json:"@odata.type,omitempty"`
		PasswordRotationEnabled          nullable.Type[bool]   `json:"passwordRotationEnabled,omitempty"`
		SessionInitiationProtocalAddress nullable.Type[string] `json:"sessionInitiationProtocalAddress,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CalendarSyncEnabled = decoded.CalendarSyncEnabled
	s.DeviceAccountEmail = decoded.DeviceAccountEmail
	s.ExchangeServer = decoded.ExchangeServer
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PasswordRotationEnabled = decoded.PasswordRotationEnabled
	s.SessionInitiationProtocalAddress = decoded.SessionInitiationProtocalAddress

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UpdateWindowsDeviceAccountActionParameter into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["deviceAccount"]; ok {
		impl, err := UnmarshalWindowsDeviceAccountImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DeviceAccount' for 'UpdateWindowsDeviceAccountActionParameter': %+v", err)
		}
		s.DeviceAccount = impl
	}

	return nil
}
