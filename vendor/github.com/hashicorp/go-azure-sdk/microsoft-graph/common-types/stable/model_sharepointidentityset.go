package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentitySet = SharePointIdentitySet{}

type SharePointIdentitySet struct {
	// The group associated with this action. Optional.
	Group Identity `json:"group"`

	// The SharePoint group associated with this action. Optional.
	SiteGroup *SharePointIdentity `json:"siteGroup,omitempty"`

	// The SharePoint user associated with this action. Optional.
	SiteUser *SharePointIdentity `json:"siteUser,omitempty"`

	// Fields inherited from IdentitySet

	// Optional. The application associated with this action.
	Application Identity `json:"application"`

	// Optional. The device associated with this action.
	Device Identity `json:"device"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Optional. The user associated with this action.
	User Identity `json:"user"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SharePointIdentitySet) IdentitySet() BaseIdentitySetImpl {
	return BaseIdentitySetImpl{
		Application: s.Application,
		Device:      s.Device,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		User:        s.User,
	}
}

var _ json.Marshaler = SharePointIdentitySet{}

func (s SharePointIdentitySet) MarshalJSON() ([]byte, error) {
	type wrapper SharePointIdentitySet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SharePointIdentitySet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SharePointIdentitySet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sharePointIdentitySet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SharePointIdentitySet: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SharePointIdentitySet{}

func (s *SharePointIdentitySet) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		SiteGroup *SharePointIdentity `json:"siteGroup,omitempty"`
		SiteUser  *SharePointIdentity `json:"siteUser,omitempty"`
		ODataId   *string             `json:"@odata.id,omitempty"`
		ODataType *string             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.SiteGroup = decoded.SiteGroup
	s.SiteUser = decoded.SiteUser
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SharePointIdentitySet into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["application"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Application' for 'SharePointIdentitySet': %+v", err)
		}
		s.Application = impl
	}

	if v, ok := temp["device"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Device' for 'SharePointIdentitySet': %+v", err)
		}
		s.Device = impl
	}

	if v, ok := temp["group"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Group' for 'SharePointIdentitySet': %+v", err)
		}
		s.Group = impl
	}

	if v, ok := temp["user"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'User' for 'SharePointIdentitySet': %+v", err)
		}
		s.User = impl
	}

	return nil
}
