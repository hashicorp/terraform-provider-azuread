package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PeopleAdminSettings{}

type PeopleAdminSettings struct {
	// Administrator settings that manage the support for item insights in an organization.
	ItemInsights *InsightsSettings `json:"itemInsights,omitempty"`

	// Administrator settings that manage the support of name pronunciation in an organization.
	NamePronunciation *NamePronunciationSettings `json:"namePronunciation,omitempty"`

	// A collection of the properties an administrator defined as visible on the Microsoft 365 profile card.
	ProfileCardProperties *[]ProfileCardProperty `json:"profileCardProperties,omitempty"`

	// A collection of profile property configuration settings defined by an administrator for an organization.
	ProfilePropertySettings *[]ProfilePropertySetting `json:"profilePropertySettings,omitempty"`

	// A collection of profile source settings configured by an administrator in an organization.
	ProfileSources *[]ProfileSource `json:"profileSources,omitempty"`

	// Administrator settings that manage the support of pronouns in an organization.
	Pronouns *PronounsSettings `json:"pronouns,omitempty"`

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

func (s PeopleAdminSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PeopleAdminSettings{}

func (s PeopleAdminSettings) MarshalJSON() ([]byte, error) {
	type wrapper PeopleAdminSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PeopleAdminSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PeopleAdminSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.peopleAdminSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PeopleAdminSettings: %+v", err)
	}

	return encoded, nil
}
