package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosSingleSignOnSettings struct {
	// List of app identifiers that are allowed to use this login. If this field is omitted, the login applies to all
	// applications on the device. This collection can contain a maximum of 500 elements.
	AllowedAppsList *[]AppListItem `json:"allowedAppsList,omitempty"`

	// List of HTTP URLs that must be matched in order to use this login. With iOS 9.0 or later, a wildcard characters may
	// be used.
	AllowedUrls *[]string `json:"allowedUrls,omitempty"`

	// The display name of login settings shown on the receiving device.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// A Kerberos principal name. If not provided, the user is prompted for one during profile installation.
	KerberosPrincipalName nullable.Type[string] `json:"kerberosPrincipalName,omitempty"`

	// A Kerberos realm name. Case sensitive.
	KerberosRealm nullable.Type[string] `json:"kerberosRealm,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &IosSingleSignOnSettings{}

func (s *IosSingleSignOnSettings) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowedUrls           *[]string             `json:"allowedUrls,omitempty"`
		DisplayName           nullable.Type[string] `json:"displayName,omitempty"`
		KerberosPrincipalName nullable.Type[string] `json:"kerberosPrincipalName,omitempty"`
		KerberosRealm         nullable.Type[string] `json:"kerberosRealm,omitempty"`
		ODataId               *string               `json:"@odata.id,omitempty"`
		ODataType             *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowedUrls = decoded.AllowedUrls
	s.DisplayName = decoded.DisplayName
	s.KerberosPrincipalName = decoded.KerberosPrincipalName
	s.KerberosRealm = decoded.KerberosRealm
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IosSingleSignOnSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["allowedAppsList"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AllowedAppsList into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AllowedAppsList' for 'IosSingleSignOnSettings': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AllowedAppsList = &output
	}

	return nil
}
