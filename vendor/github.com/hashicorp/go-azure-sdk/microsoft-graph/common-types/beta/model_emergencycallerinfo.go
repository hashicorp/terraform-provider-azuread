package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmergencyCallerInfo struct {
	// The display name of the emergency caller.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The location of the emergency caller.
	Location Location `json:"location"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The phone number of the emergency caller.
	PhoneNumber nullable.Type[string] `json:"phoneNumber,omitempty"`

	// The tenant ID of the emergency caller.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// The user principal name of the emergency caller.
	Upn nullable.Type[string] `json:"upn,omitempty"`
}

var _ json.Unmarshaler = &EmergencyCallerInfo{}

func (s *EmergencyCallerInfo) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName nullable.Type[string] `json:"displayName,omitempty"`
		ODataId     *string               `json:"@odata.id,omitempty"`
		ODataType   *string               `json:"@odata.type,omitempty"`
		PhoneNumber nullable.Type[string] `json:"phoneNumber,omitempty"`
		TenantId    nullable.Type[string] `json:"tenantId,omitempty"`
		Upn         nullable.Type[string] `json:"upn,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PhoneNumber = decoded.PhoneNumber
	s.TenantId = decoded.TenantId
	s.Upn = decoded.Upn

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EmergencyCallerInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["location"]; ok {
		impl, err := UnmarshalLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Location' for 'EmergencyCallerInfo': %+v", err)
		}
		s.Location = impl
	}

	return nil
}
