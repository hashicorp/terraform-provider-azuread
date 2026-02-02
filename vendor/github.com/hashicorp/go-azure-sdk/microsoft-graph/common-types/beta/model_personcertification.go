package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ItemFacet = PersonCertification{}

type PersonCertification struct {
	// The referenceable identifier for the certification.
	CertificationId nullable.Type[string] `json:"certificationId,omitempty"`

	// Description of the certification.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Title of the certification.
	DisplayName *string `json:"displayName,omitempty"`

	// The date that the certification expires.
	EndDate nullable.Type[string] `json:"endDate,omitempty"`

	// The date that the certification was issued.
	IssuedDate nullable.Type[string] `json:"issuedDate,omitempty"`

	// Authority which granted the certification.
	IssuingAuthority nullable.Type[string] `json:"issuingAuthority,omitempty"`

	// Company which granted the certification.
	IssuingCompany nullable.Type[string] `json:"issuingCompany,omitempty"`

	// The date that the certification became valid.
	StartDate nullable.Type[string] `json:"startDate,omitempty"`

	// URL referencing a thumbnail of the certification.
	ThumbnailUrl nullable.Type[string] `json:"thumbnailUrl,omitempty"`

	// URL referencing the certification.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

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

func (s PersonCertification) ItemFacet() BaseItemFacetImpl {
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

func (s PersonCertification) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PersonCertification{}

func (s PersonCertification) MarshalJSON() ([]byte, error) {
	type wrapper PersonCertification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PersonCertification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PersonCertification: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.personCertification"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PersonCertification: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PersonCertification{}

func (s *PersonCertification) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CertificationId      nullable.Type[string]      `json:"certificationId,omitempty"`
		Description          nullable.Type[string]      `json:"description,omitempty"`
		DisplayName          *string                    `json:"displayName,omitempty"`
		EndDate              nullable.Type[string]      `json:"endDate,omitempty"`
		IssuedDate           nullable.Type[string]      `json:"issuedDate,omitempty"`
		IssuingAuthority     nullable.Type[string]      `json:"issuingAuthority,omitempty"`
		IssuingCompany       nullable.Type[string]      `json:"issuingCompany,omitempty"`
		StartDate            nullable.Type[string]      `json:"startDate,omitempty"`
		ThumbnailUrl         nullable.Type[string]      `json:"thumbnailUrl,omitempty"`
		WebUrl               nullable.Type[string]      `json:"webUrl,omitempty"`
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

	s.CertificationId = decoded.CertificationId
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.EndDate = decoded.EndDate
	s.IssuedDate = decoded.IssuedDate
	s.IssuingAuthority = decoded.IssuingAuthority
	s.IssuingCompany = decoded.IssuingCompany
	s.StartDate = decoded.StartDate
	s.ThumbnailUrl = decoded.ThumbnailUrl
	s.WebUrl = decoded.WebUrl
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
		return fmt.Errorf("unmarshaling PersonCertification into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'PersonCertification': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'PersonCertification': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
