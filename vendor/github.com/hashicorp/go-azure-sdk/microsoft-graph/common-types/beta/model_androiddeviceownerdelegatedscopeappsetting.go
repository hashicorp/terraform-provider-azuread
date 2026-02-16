package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerDelegatedScopeAppSetting struct {
	// Represents an app in the list of managed applications
	AppDetail AppListItem `json:"appDetail"`

	// List of scopes an app has been assigned.
	AppScopes *[]AndroidDeviceOwnerDelegatedAppScopeType `json:"appScopes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &AndroidDeviceOwnerDelegatedScopeAppSetting{}

func (s *AndroidDeviceOwnerDelegatedScopeAppSetting) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AppScopes *[]AndroidDeviceOwnerDelegatedAppScopeType `json:"appScopes,omitempty"`
		ODataId   *string                                    `json:"@odata.id,omitempty"`
		ODataType *string                                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppScopes = decoded.AppScopes
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AndroidDeviceOwnerDelegatedScopeAppSetting into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appDetail"]; ok {
		impl, err := UnmarshalAppListItemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AppDetail' for 'AndroidDeviceOwnerDelegatedScopeAppSetting': %+v", err)
		}
		s.AppDetail = impl
	}

	return nil
}
