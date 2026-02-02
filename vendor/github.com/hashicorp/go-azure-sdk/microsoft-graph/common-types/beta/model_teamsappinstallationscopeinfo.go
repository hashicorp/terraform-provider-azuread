package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppInstallationScopeInfo interface {
	TeamsAppInstallationScopeInfo() BaseTeamsAppInstallationScopeInfoImpl
}

var _ TeamsAppInstallationScopeInfo = BaseTeamsAppInstallationScopeInfoImpl{}

type BaseTeamsAppInstallationScopeInfoImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Scope *TeamsAppInstallationScopes `json:"scope,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseTeamsAppInstallationScopeInfoImpl) TeamsAppInstallationScopeInfo() BaseTeamsAppInstallationScopeInfoImpl {
	return s
}

var _ TeamsAppInstallationScopeInfo = RawTeamsAppInstallationScopeInfoImpl{}

// RawTeamsAppInstallationScopeInfoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTeamsAppInstallationScopeInfoImpl struct {
	teamsAppInstallationScopeInfo BaseTeamsAppInstallationScopeInfoImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawTeamsAppInstallationScopeInfoImpl) TeamsAppInstallationScopeInfo() BaseTeamsAppInstallationScopeInfoImpl {
	return s.teamsAppInstallationScopeInfo
}

func UnmarshalTeamsAppInstallationScopeInfoImplementation(input []byte) (TeamsAppInstallationScopeInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAppInstallationScopeInfo into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.groupChatTeamsAppInstallationScopeInfo") {
		var out GroupChatTeamsAppInstallationScopeInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupChatTeamsAppInstallationScopeInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personalTeamsAppInstallationScopeInfo") {
		var out PersonalTeamsAppInstallationScopeInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonalTeamsAppInstallationScopeInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamTeamsAppInstallationScopeInfo") {
		var out TeamTeamsAppInstallationScopeInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamTeamsAppInstallationScopeInfo: %+v", err)
		}
		return out, nil
	}

	var parent BaseTeamsAppInstallationScopeInfoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTeamsAppInstallationScopeInfoImpl: %+v", err)
	}

	return RawTeamsAppInstallationScopeInfoImpl{
		teamsAppInstallationScopeInfo: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
