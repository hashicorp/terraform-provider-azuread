package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationOrganization = EducationSchool{}

type EducationSchool struct {
	// Address of the school.
	Address *PhysicalAddress `json:"address,omitempty"`

	AdministrativeUnit *AdministrativeUnit `json:"administrativeUnit,omitempty"`

	// Classes taught at the school. Nullable.
	Classes *[]EducationClass `json:"classes,omitempty"`

	// Entity who created the school.
	CreatedBy IdentitySet `json:"createdBy"`

	// ID of school in syncing system.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// ID of principal in syncing system.
	ExternalPrincipalId nullable.Type[string] `json:"externalPrincipalId,omitempty"`

	Fax nullable.Type[string] `json:"fax,omitempty"`

	// Highest grade taught.
	HighestGrade nullable.Type[string] `json:"highestGrade,omitempty"`

	// Lowest grade taught.
	LowestGrade nullable.Type[string] `json:"lowestGrade,omitempty"`

	// Phone number of school.
	Phone nullable.Type[string] `json:"phone,omitempty"`

	// Email address of the principal.
	PrincipalEmail nullable.Type[string] `json:"principalEmail,omitempty"`

	// Name of the principal.
	PrincipalName nullable.Type[string] `json:"principalName,omitempty"`

	// School Number.
	SchoolNumber nullable.Type[string] `json:"schoolNumber,omitempty"`

	// Users in the school. Nullable.
	Users *[]EducationUser `json:"users,omitempty"`

	// Fields inherited from EducationOrganization

	// Organization description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Organization display name.
	DisplayName *string `json:"displayName,omitempty"`

	// Where this user was created from. Possible values are: sis, lms, or manual.
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`

	ExternalSourceDetail nullable.Type[string] `json:"externalSourceDetail,omitempty"`

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

func (s EducationSchool) EducationOrganization() BaseEducationOrganizationImpl {
	return BaseEducationOrganizationImpl{
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		ExternalSource:       s.ExternalSource,
		ExternalSourceDetail: s.ExternalSourceDetail,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s EducationSchool) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationSchool{}

func (s EducationSchool) MarshalJSON() ([]byte, error) {
	type wrapper EducationSchool
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationSchool: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSchool: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationSchool"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationSchool: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationSchool{}

func (s *EducationSchool) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Address              *PhysicalAddress         `json:"address,omitempty"`
		AdministrativeUnit   *AdministrativeUnit      `json:"administrativeUnit,omitempty"`
		Classes              *[]EducationClass        `json:"classes,omitempty"`
		ExternalId           nullable.Type[string]    `json:"externalId,omitempty"`
		ExternalPrincipalId  nullable.Type[string]    `json:"externalPrincipalId,omitempty"`
		Fax                  nullable.Type[string]    `json:"fax,omitempty"`
		HighestGrade         nullable.Type[string]    `json:"highestGrade,omitempty"`
		LowestGrade          nullable.Type[string]    `json:"lowestGrade,omitempty"`
		Phone                nullable.Type[string]    `json:"phone,omitempty"`
		PrincipalEmail       nullable.Type[string]    `json:"principalEmail,omitempty"`
		PrincipalName        nullable.Type[string]    `json:"principalName,omitempty"`
		SchoolNumber         nullable.Type[string]    `json:"schoolNumber,omitempty"`
		Users                *[]EducationUser         `json:"users,omitempty"`
		Description          nullable.Type[string]    `json:"description,omitempty"`
		DisplayName          *string                  `json:"displayName,omitempty"`
		ExternalSource       *EducationExternalSource `json:"externalSource,omitempty"`
		ExternalSourceDetail nullable.Type[string]    `json:"externalSourceDetail,omitempty"`
		Id                   *string                  `json:"id,omitempty"`
		ODataId              *string                  `json:"@odata.id,omitempty"`
		ODataType            *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Address = decoded.Address
	s.AdministrativeUnit = decoded.AdministrativeUnit
	s.Classes = decoded.Classes
	s.ExternalId = decoded.ExternalId
	s.ExternalPrincipalId = decoded.ExternalPrincipalId
	s.Fax = decoded.Fax
	s.HighestGrade = decoded.HighestGrade
	s.LowestGrade = decoded.LowestGrade
	s.Phone = decoded.Phone
	s.PrincipalEmail = decoded.PrincipalEmail
	s.PrincipalName = decoded.PrincipalName
	s.SchoolNumber = decoded.SchoolNumber
	s.Users = decoded.Users
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ExternalSource = decoded.ExternalSource
	s.ExternalSourceDetail = decoded.ExternalSourceDetail
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationSchool into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EducationSchool': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
