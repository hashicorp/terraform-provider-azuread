package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADJoinPolicy struct {
	AllowedToJoin       DeviceRegistrationMembership `json:"allowedToJoin"`
	IsAdminConfigurable nullable.Type[bool]          `json:"isAdminConfigurable,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &AzureADJoinPolicy{}

func (s *AzureADJoinPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsAdminConfigurable nullable.Type[bool] `json:"isAdminConfigurable,omitempty"`
		ODataId             *string             `json:"@odata.id,omitempty"`
		ODataType           *string             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsAdminConfigurable = decoded.IsAdminConfigurable
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AzureADJoinPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["allowedToJoin"]; ok {
		impl, err := UnmarshalDeviceRegistrationMembershipImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AllowedToJoin' for 'AzureADJoinPolicy': %+v", err)
		}
		s.AllowedToJoin = impl
	}

	return nil
}
