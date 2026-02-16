package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessGuestsOrExternalUsers struct {
	// The tenant IDs of the selected types of external users. Either all B2B tenant or a collection of tenant IDs. External
	// tenants can be specified only when the property guestOrExternalUserTypes isn't null or an empty String.
	ExternalTenants ConditionalAccessExternalTenants `json:"externalTenants"`

	GuestOrExternalUserTypes *ConditionalAccessGuestOrExternalUserTypes `json:"guestOrExternalUserTypes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &ConditionalAccessGuestsOrExternalUsers{}

func (s *ConditionalAccessGuestsOrExternalUsers) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		GuestOrExternalUserTypes *ConditionalAccessGuestOrExternalUserTypes `json:"guestOrExternalUserTypes,omitempty"`
		ODataId                  *string                                    `json:"@odata.id,omitempty"`
		ODataType                *string                                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.GuestOrExternalUserTypes = decoded.GuestOrExternalUserTypes
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ConditionalAccessGuestsOrExternalUsers into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["externalTenants"]; ok {
		impl, err := UnmarshalConditionalAccessExternalTenantsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ExternalTenants' for 'ConditionalAccessGuestsOrExternalUsers': %+v", err)
		}
		s.ExternalTenants = impl
	}

	return nil
}
