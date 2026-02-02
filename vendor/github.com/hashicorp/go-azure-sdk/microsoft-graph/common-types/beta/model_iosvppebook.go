package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedEBook = IosVppEBook{}

type IosVppEBook struct {
	// The Apple ID associated with Vpp token.
	AppleId nullable.Type[string] `json:"appleId,omitempty"`

	// Genres.
	Genres *[]string `json:"genres,omitempty"`

	// Language.
	Language nullable.Type[string] `json:"language,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Seller.
	Seller nullable.Type[string] `json:"seller,omitempty"`

	// Total license count.
	TotalLicenseCount *int64 `json:"totalLicenseCount,omitempty"`

	// Used license count.
	UsedLicenseCount *int64 `json:"usedLicenseCount,omitempty"`

	// The Vpp token's organization name.
	VppOrganizationName nullable.Type[string] `json:"vppOrganizationName,omitempty"`

	// The Vpp token ID.
	VppTokenId *string `json:"vppTokenId,omitempty"`

	// Fields inherited from ManagedEBook

	// The list of assignments for this eBook.
	Assignments *[]ManagedEBookAssignment `json:"assignments,omitempty"`

	// The list of categories for this eBook.
	Categories *[]ManagedEBookCategory `json:"categories,omitempty"`

	// The date and time when the eBook file was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The list of installation states for this eBook.
	DeviceStates *[]DeviceInstallState `json:"deviceStates,omitempty"`

	// Name of the eBook.
	DisplayName *string `json:"displayName,omitempty"`

	// The more information Url.
	InformationUrl nullable.Type[string] `json:"informationUrl,omitempty"`

	// Mobile App Install Summary.
	InstallSummary *EBookInstallSummary `json:"installSummary,omitempty"`

	// Book cover.
	LargeCover *MimeContent `json:"largeCover,omitempty"`

	// The date and time when the eBook was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The privacy statement Url.
	PrivacyInformationUrl nullable.Type[string] `json:"privacyInformationUrl,omitempty"`

	// The date and time when the eBook was published.
	PublishedDateTime *string `json:"publishedDateTime,omitempty"`

	// Publisher.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// The list of installation states for this eBook.
	UserStateSummary *[]UserInstallStateSummary `json:"userStateSummary,omitempty"`

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

func (s IosVppEBook) ManagedEBook() BaseManagedEBookImpl {
	return BaseManagedEBookImpl{
		Assignments:           s.Assignments,
		Categories:            s.Categories,
		CreatedDateTime:       s.CreatedDateTime,
		Description:           s.Description,
		DeviceStates:          s.DeviceStates,
		DisplayName:           s.DisplayName,
		InformationUrl:        s.InformationUrl,
		InstallSummary:        s.InstallSummary,
		LargeCover:            s.LargeCover,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		PrivacyInformationUrl: s.PrivacyInformationUrl,
		PublishedDateTime:     s.PublishedDateTime,
		Publisher:             s.Publisher,
		UserStateSummary:      s.UserStateSummary,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s IosVppEBook) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosVppEBook{}

func (s IosVppEBook) MarshalJSON() ([]byte, error) {
	type wrapper IosVppEBook
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosVppEBook: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosVppEBook: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosVppEBook"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosVppEBook: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IosVppEBook{}

func (s *IosVppEBook) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AppleId               nullable.Type[string]      `json:"appleId,omitempty"`
		Genres                *[]string                  `json:"genres,omitempty"`
		Language              nullable.Type[string]      `json:"language,omitempty"`
		RoleScopeTagIds       *[]string                  `json:"roleScopeTagIds,omitempty"`
		Seller                nullable.Type[string]      `json:"seller,omitempty"`
		TotalLicenseCount     *int64                     `json:"totalLicenseCount,omitempty"`
		UsedLicenseCount      *int64                     `json:"usedLicenseCount,omitempty"`
		VppOrganizationName   nullable.Type[string]      `json:"vppOrganizationName,omitempty"`
		VppTokenId            *string                    `json:"vppTokenId,omitempty"`
		Categories            *[]ManagedEBookCategory    `json:"categories,omitempty"`
		CreatedDateTime       *string                    `json:"createdDateTime,omitempty"`
		Description           nullable.Type[string]      `json:"description,omitempty"`
		DeviceStates          *[]DeviceInstallState      `json:"deviceStates,omitempty"`
		DisplayName           *string                    `json:"displayName,omitempty"`
		InformationUrl        nullable.Type[string]      `json:"informationUrl,omitempty"`
		InstallSummary        *EBookInstallSummary       `json:"installSummary,omitempty"`
		LargeCover            *MimeContent               `json:"largeCover,omitempty"`
		LastModifiedDateTime  *string                    `json:"lastModifiedDateTime,omitempty"`
		PrivacyInformationUrl nullable.Type[string]      `json:"privacyInformationUrl,omitempty"`
		PublishedDateTime     *string                    `json:"publishedDateTime,omitempty"`
		Publisher             nullable.Type[string]      `json:"publisher,omitempty"`
		UserStateSummary      *[]UserInstallStateSummary `json:"userStateSummary,omitempty"`
		Id                    *string                    `json:"id,omitempty"`
		ODataId               *string                    `json:"@odata.id,omitempty"`
		ODataType             *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppleId = decoded.AppleId
	s.Genres = decoded.Genres
	s.Language = decoded.Language
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.Seller = decoded.Seller
	s.TotalLicenseCount = decoded.TotalLicenseCount
	s.UsedLicenseCount = decoded.UsedLicenseCount
	s.VppOrganizationName = decoded.VppOrganizationName
	s.VppTokenId = decoded.VppTokenId
	s.Categories = decoded.Categories
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DeviceStates = decoded.DeviceStates
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.InformationUrl = decoded.InformationUrl
	s.InstallSummary = decoded.InstallSummary
	s.LargeCover = decoded.LargeCover
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PrivacyInformationUrl = decoded.PrivacyInformationUrl
	s.PublishedDateTime = decoded.PublishedDateTime
	s.Publisher = decoded.Publisher
	s.UserStateSummary = decoded.UserStateSummary

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IosVppEBook into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["assignments"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Assignments into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedEBookAssignment, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedEBookAssignmentImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Assignments' for 'IosVppEBook': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Assignments = &output
	}

	return nil
}
