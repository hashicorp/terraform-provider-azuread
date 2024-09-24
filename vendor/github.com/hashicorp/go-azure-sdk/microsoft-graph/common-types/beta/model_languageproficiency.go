package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ItemFacet = LanguageProficiency{}

type LanguageProficiency struct {
	// Contains the long-form name for the language.
	DisplayName *string `json:"displayName,omitempty"`

	Proficiency *LanguageProficiencyLevel `json:"proficiency,omitempty"`

	// Represents the users reading comprehension for the language represented by the object. Possible values are:
	// elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual,
	// unknownFutureValue.
	Reading *LanguageProficiencyLevel `json:"reading,omitempty"`

	// Represents the users spoken proficiency for the language represented by the object. Possible values are: elementary,
	// conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
	Spoken *LanguageProficiencyLevel `json:"spoken,omitempty"`

	// Contains the four-character BCP47 name for the language (en-US, no-NB, en-AU).
	Tag nullable.Type[string] `json:"tag,omitempty"`

	ThumbnailUrl nullable.Type[string] `json:"thumbnailUrl,omitempty"`

	// Represents the users written proficiency for the language represented by the object. Possible values are: elementary,
	// conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
	Written *LanguageProficiencyLevel `json:"written,omitempty"`

	// Fields inherited from ItemFacet

	// The audiences that are able to see the values contained within the associated entity. Possible values are: me,
	// family, contacts, groupMembers, organization, federatedOrganizations, everyone, unknownFutureValue.
	AllowedAudiences *AllowedAudiences `json:"allowedAudiences,omitempty"`

	CreatedBy IdentitySet `json:"createdBy"`

	// Provides the dateTimeOffset for when the entity was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Contains inference detail if the entity is inferred by the creating or modifying application.
	Inference *InferenceData `json:"inference,omitempty"`

	IsSearchable   nullable.Type[bool] `json:"isSearchable,omitempty"`
	LastModifiedBy IdentitySet         `json:"lastModifiedBy"`

	// Provides the dateTimeOffset for when the entity was created.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Where the values within an entity originated if synced from another service.
	Source *PersonDataSources `json:"source,omitempty"`

	// Where the values within an entity originated if synced from another source.
	Sources *[]ProfileSourceAnnotation `json:"sources,omitempty"`

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

func (s LanguageProficiency) ItemFacet() BaseItemFacetImpl {
	return BaseItemFacetImpl{
		AllowedAudiences:     s.AllowedAudiences,
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		Inference:            s.Inference,
		IsSearchable:         s.IsSearchable,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Source:               s.Source,
		Sources:              s.Sources,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s LanguageProficiency) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = LanguageProficiency{}

func (s LanguageProficiency) MarshalJSON() ([]byte, error) {
	type wrapper LanguageProficiency
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LanguageProficiency: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LanguageProficiency: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.languageProficiency"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LanguageProficiency: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &LanguageProficiency{}

func (s *LanguageProficiency) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName          *string                    `json:"displayName,omitempty"`
		Proficiency          *LanguageProficiencyLevel  `json:"proficiency,omitempty"`
		Reading              *LanguageProficiencyLevel  `json:"reading,omitempty"`
		Spoken               *LanguageProficiencyLevel  `json:"spoken,omitempty"`
		Tag                  nullable.Type[string]      `json:"tag,omitempty"`
		ThumbnailUrl         nullable.Type[string]      `json:"thumbnailUrl,omitempty"`
		Written              *LanguageProficiencyLevel  `json:"written,omitempty"`
		AllowedAudiences     *AllowedAudiences          `json:"allowedAudiences,omitempty"`
		CreatedDateTime      *string                    `json:"createdDateTime,omitempty"`
		Inference            *InferenceData             `json:"inference,omitempty"`
		IsSearchable         nullable.Type[bool]        `json:"isSearchable,omitempty"`
		LastModifiedDateTime *string                    `json:"lastModifiedDateTime,omitempty"`
		Source               *PersonDataSources         `json:"source,omitempty"`
		Sources              *[]ProfileSourceAnnotation `json:"sources,omitempty"`
		Id                   *string                    `json:"id,omitempty"`
		ODataId              *string                    `json:"@odata.id,omitempty"`
		ODataType            *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.Proficiency = decoded.Proficiency
	s.Reading = decoded.Reading
	s.Spoken = decoded.Spoken
	s.Tag = decoded.Tag
	s.ThumbnailUrl = decoded.ThumbnailUrl
	s.Written = decoded.Written
	s.AllowedAudiences = decoded.AllowedAudiences
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.Inference = decoded.Inference
	s.IsSearchable = decoded.IsSearchable
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Source = decoded.Source
	s.Sources = decoded.Sources

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling LanguageProficiency into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'LanguageProficiency': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'LanguageProficiency': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
