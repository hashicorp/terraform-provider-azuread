package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemFacet interface {
	Entity
	ItemFacet() BaseItemFacetImpl
}

var _ ItemFacet = BaseItemFacetImpl{}

type BaseItemFacetImpl struct {
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

func (s BaseItemFacetImpl) ItemFacet() BaseItemFacetImpl {
	return s
}

func (s BaseItemFacetImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ItemFacet = RawItemFacetImpl{}

// RawItemFacetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawItemFacetImpl struct {
	itemFacet BaseItemFacetImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawItemFacetImpl) ItemFacet() BaseItemFacetImpl {
	return s.itemFacet
}

func (s RawItemFacetImpl) Entity() BaseEntityImpl {
	return s.itemFacet.Entity()
}

var _ json.Marshaler = BaseItemFacetImpl{}

func (s BaseItemFacetImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseItemFacetImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseItemFacetImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseItemFacetImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.itemFacet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseItemFacetImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseItemFacetImpl{}

func (s *BaseItemFacetImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
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

	s.AllowedAudiences = decoded.AllowedAudiences
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Inference = decoded.Inference
	s.IsSearchable = decoded.IsSearchable
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Source = decoded.Source
	s.Sources = decoded.Sources
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseItemFacetImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseItemFacetImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseItemFacetImpl': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}

func UnmarshalItemFacetImplementation(input []byte) (ItemFacet, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ItemFacet into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.educationalActivity") {
		var out EducationalActivity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationalActivity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemAddress") {
		var out ItemAddress
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemAddress: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemEmail") {
		var out ItemEmail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemEmail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemPatent") {
		var out ItemPatent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemPatent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemPhone") {
		var out ItemPhone
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemPhone: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemPublication") {
		var out ItemPublication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemPublication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.languageProficiency") {
		var out LanguageProficiency
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LanguageProficiency: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personAnnotation") {
		var out PersonAnnotation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonAnnotation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personAnnualEvent") {
		var out PersonAnnualEvent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonAnnualEvent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personAward") {
		var out PersonAward
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonAward: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personCertification") {
		var out PersonCertification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonCertification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personInterest") {
		var out PersonInterest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonInterest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personName") {
		var out PersonName
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonName: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personResponsibility") {
		var out PersonResponsibility
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonResponsibility: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.personWebsite") {
		var out PersonWebsite
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersonWebsite: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.projectParticipation") {
		var out ProjectParticipation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProjectParticipation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.skillProficiency") {
		var out SkillProficiency
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SkillProficiency: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userAccountInformation") {
		var out UserAccountInformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserAccountInformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.webAccount") {
		var out WebAccount
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebAccount: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workPosition") {
		var out WorkPosition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkPosition: %+v", err)
		}
		return out, nil
	}

	var parent BaseItemFacetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseItemFacetImpl: %+v", err)
	}

	return RawItemFacetImpl{
		itemFacet: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
