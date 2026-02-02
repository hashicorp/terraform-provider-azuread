package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserTeamwork{}

type UserTeamwork struct {
	// The list of associatedTeamInfo objects that a user is associated with.
	AssociatedTeams *[]AssociatedTeamInfo `json:"associatedTeams,omitempty"`

	// The apps installed in the personal scope of this user.
	InstalledApps *[]UserScopeTeamsAppInstallation `json:"installedApps,omitempty"`

	// Represents the location that a user selected in Microsoft Teams and doesn't follow the Office's locale setting. A
	// user's locale is represented by their preferred language and country or region. For example, en-us. The language
	// component follows two-letter codes as defined in ISO 639-1, and the country component follows two-letter codes as
	// defined in ISO 3166-1 alpha-2.
	Locale nullable.Type[string] `json:"locale,omitempty"`

	// Represents the region of the organization or the user. For users with multigeo licenses, the property contains the
	// user's region (if available). For users without multigeo licenses, the property contains the organization's
	// region.The region value can be any region supported by the Teams payload. The possible values are: Americas, Europe
	// and MiddleEast, Asia Pacific, UAE, Australia, Brazil, Canada, Switzerland, Germany, France, India, Japan, South
	// Korea, Norway, Singapore, United Kingdom, South Africa, Sweden, Qatar, Poland, Italy, Israel, Spain, Mexico, USGov
	// Community Cloud, USGov Community Cloud High, USGov Department of Defense, and China.
	Region nullable.Type[string] `json:"region,omitempty"`

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

func (s UserTeamwork) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserTeamwork{}

func (s UserTeamwork) MarshalJSON() ([]byte, error) {
	type wrapper UserTeamwork
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserTeamwork: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserTeamwork: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userTeamwork"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserTeamwork: %+v", err)
	}

	return encoded, nil
}
