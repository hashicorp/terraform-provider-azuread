package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TeamInfo = AssociatedTeamInfo{}

type AssociatedTeamInfo struct {

	// Fields inherited from TeamInfo

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

func (s AssociatedTeamInfo) TeamInfo() BaseTeamInfoImpl {
	return BaseTeamInfoImpl{
		DisplayName: s.DisplayName,
		Team:        s.Team,
		TenantId:    s.TenantId,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s AssociatedTeamInfo) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AssociatedTeamInfo{}

func (s AssociatedTeamInfo) MarshalJSON() ([]byte, error) {
	type wrapper AssociatedTeamInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AssociatedTeamInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AssociatedTeamInfo: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.associatedTeamInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AssociatedTeamInfo: %+v", err)
	}

	return encoded, nil
}
