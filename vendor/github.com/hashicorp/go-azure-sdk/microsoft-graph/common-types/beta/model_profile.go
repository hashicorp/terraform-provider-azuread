package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Profile{}

type Profile struct {
	Account *[]UserAccountInformation `json:"account,omitempty"`

	// Represents details of addresses associated with the user.
	Addresses *[]ItemAddress `json:"addresses,omitempty"`

	// Represents the details of meaningful dates associated with a person.
	Anniversaries *[]PersonAnnualEvent `json:"anniversaries,omitempty"`

	// Represents the details of awards or honors associated with a person.
	Awards *[]PersonAward `json:"awards,omitempty"`

	// Represents the details of certifications associated with a person.
	Certifications *[]PersonCertification `json:"certifications,omitempty"`

	// Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational
	// activities.
	EducationalActivities *[]EducationalActivity `json:"educationalActivities,omitempty"`

	// Represents detailed information about email addresses associated with the user.
	Emails *[]ItemEmail `json:"emails,omitempty"`

	// Provides detailed information about interests the user has associated with themselves in various services.
	Interests *[]PersonInterest `json:"interests,omitempty"`

	// Represents detailed information about languages that a user has added to their profile.
	Languages *[]LanguageProficiency `json:"languages,omitempty"`

	// Represents the names a user has added to their profile.
	Names *[]PersonName `json:"names,omitempty"`

	// Represents notes that a user has added to their profile.
	Notes *[]PersonAnnotation `json:"notes,omitempty"`

	// Represents patents that a user has added to their profile.
	Patents *[]ItemPatent `json:"patents,omitempty"`

	// Represents detailed information about phone numbers associated with a user in various services.
	Phones *[]ItemPhone `json:"phones,omitempty"`

	// Represents detailed information about work positions associated with a user's profile.
	Positions *[]WorkPosition `json:"positions,omitempty"`

	// Represents detailed information about projects associated with a user.
	Projects *[]ProjectParticipation `json:"projects,omitempty"`

	// Represents details of any publications a user has added to their profile.
	Publications *[]ItemPublication `json:"publications,omitempty"`

	// Represents detailed information about skills associated with a user in various services.
	Skills *[]SkillProficiency `json:"skills,omitempty"`

	// Represents web accounts the user has indicated they use or has added to their user profile.
	WebAccounts *[]WebAccount `json:"webAccounts,omitempty"`

	// Represents detailed information about websites associated with a user in various services.
	Websites *[]PersonWebsite `json:"websites,omitempty"`

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

func (s Profile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Profile{}

func (s Profile) MarshalJSON() ([]byte, error) {
	type wrapper Profile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Profile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Profile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.profile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Profile: %+v", err)
	}

	return encoded, nil
}
