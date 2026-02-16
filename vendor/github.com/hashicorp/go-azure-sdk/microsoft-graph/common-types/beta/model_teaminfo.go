package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamInfo interface {
	Entity
	TeamInfo() BaseTeamInfoImpl
}

var _ TeamInfo = BaseTeamInfoImpl{}

type BaseTeamInfoImpl struct {
	// The name of the team.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	Team *Team `json:"team,omitempty"`

	// The ID of the Microsoft Entra tenant.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseTeamInfoImpl) TeamInfo() BaseTeamInfoImpl {
	return s
}

func (s BaseTeamInfoImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ TeamInfo = RawTeamInfoImpl{}

// RawTeamInfoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTeamInfoImpl struct {
	teamInfo BaseTeamInfoImpl
	Type     string
	Values   map[string]interface{}
}

func (s RawTeamInfoImpl) TeamInfo() BaseTeamInfoImpl {
	return s.teamInfo
}

func (s RawTeamInfoImpl) Entity() BaseEntityImpl {
	return s.teamInfo.Entity()
}

var _ json.Marshaler = BaseTeamInfoImpl{}

func (s BaseTeamInfoImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseTeamInfoImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseTeamInfoImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseTeamInfoImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseTeamInfoImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalTeamInfoImplementation(input []byte) (TeamInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamInfo into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.associatedTeamInfo") {
		var out AssociatedTeamInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AssociatedTeamInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharedWithChannelTeamInfo") {
		var out SharedWithChannelTeamInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedWithChannelTeamInfo: %+v", err)
		}
		return out, nil
	}

	var parent BaseTeamInfoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTeamInfoImpl: %+v", err)
	}

	return RawTeamInfoImpl{
		teamInfo: parent,
		Type:     value,
		Values:   temp,
	}, nil

}
