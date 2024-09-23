package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Teamwork{}

type Teamwork struct {
	// A collection of deleted chats.
	DeletedChats *[]DeletedChat `json:"deletedChats,omitempty"`

	// A collection of deleted teams.
	DeletedTeams *[]DeletedTeam `json:"deletedTeams,omitempty"`

	// The Teams devices provisioned for the tenant.
	Devices *[]TeamworkDevice `json:"devices,omitempty"`

	// Indicates whether Microsoft Teams is enabled for the organization.
	IsTeamsEnabled *bool `json:"isTeamsEnabled,omitempty"`

	// Represents the region of the organization or the tenant. The region value can be any region supported by the Teams
	// payload. The possible values are: Americas, Europe and MiddleEast, Asia Pacific, UAE, Australia, Brazil, Canada,
	// Switzerland, Germany, France, India, Japan, South Korea, Norway, Singapore, United Kingdom, South Africa, Sweden,
	// Qatar, Poland, Italy, Israel, USGov Community Cloud, USGov Community Cloud High, USGov Department of Defense, and
	// China.
	Region nullable.Type[string] `json:"region,omitempty"`

	// The templates associated with a team.
	TeamTemplates *[]TeamTemplate `json:"teamTemplates,omitempty"`

	// Represents tenant-wide settings for all Teams apps in the tenant.
	TeamsAppSettings *TeamsAppSettings `json:"teamsAppSettings,omitempty"`

	// A workforce integration with shifts.
	WorkforceIntegrations *[]WorkforceIntegration `json:"workforceIntegrations,omitempty"`

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

func (s Teamwork) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Teamwork{}

func (s Teamwork) MarshalJSON() ([]byte, error) {
	type wrapper Teamwork
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Teamwork: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Teamwork: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamwork"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Teamwork: %+v", err)
	}

	return encoded, nil
}
