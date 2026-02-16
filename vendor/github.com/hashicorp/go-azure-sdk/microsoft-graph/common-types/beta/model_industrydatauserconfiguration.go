package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataUserConfiguration struct {
	// The password settings for the users to be provisioned with.
	DefaultPasswordSettings IndustryDataPasswordSettings `json:"defaultPasswordSettings"`

	// The license skus for the users to be provisioned with.
	LicenseSkus *[]string `json:"licenseSkus,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RoleGroup *IndustryDataRoleGroup `json:"roleGroup,omitempty"`
}

var _ json.Unmarshaler = &IndustryDataUserConfiguration{}

func (s *IndustryDataUserConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		LicenseSkus *[]string              `json:"licenseSkus,omitempty"`
		ODataId     *string                `json:"@odata.id,omitempty"`
		ODataType   *string                `json:"@odata.type,omitempty"`
		RoleGroup   *IndustryDataRoleGroup `json:"roleGroup,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.LicenseSkus = decoded.LicenseSkus
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RoleGroup = decoded.RoleGroup

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IndustryDataUserConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["defaultPasswordSettings"]; ok {
		impl, err := UnmarshalIndustryDataPasswordSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DefaultPasswordSettings' for 'IndustryDataUserConfiguration': %+v", err)
		}
		s.DefaultPasswordSettings = impl
	}

	return nil
}
