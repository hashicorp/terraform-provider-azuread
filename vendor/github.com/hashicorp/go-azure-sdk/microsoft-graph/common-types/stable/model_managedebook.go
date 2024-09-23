package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedEBook interface {
	Entity
	ManagedEBook() BaseManagedEBookImpl
}

var _ ManagedEBook = BaseManagedEBookImpl{}

type BaseManagedEBookImpl struct {
	// The list of assignments for this eBook.
	Assignments *[]ManagedEBookAssignment `json:"assignments,omitempty"`

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

func (s BaseManagedEBookImpl) ManagedEBook() BaseManagedEBookImpl {
	return s
}

func (s BaseManagedEBookImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ManagedEBook = RawManagedEBookImpl{}

// RawManagedEBookImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawManagedEBookImpl struct {
	managedEBook BaseManagedEBookImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawManagedEBookImpl) ManagedEBook() BaseManagedEBookImpl {
	return s.managedEBook
}

func (s RawManagedEBookImpl) Entity() BaseEntityImpl {
	return s.managedEBook.Entity()
}

var _ json.Marshaler = BaseManagedEBookImpl{}

func (s BaseManagedEBookImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseManagedEBookImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseManagedEBookImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseManagedEBookImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedEBook"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseManagedEBookImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseManagedEBookImpl{}

func (s *BaseManagedEBookImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
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

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DeviceStates = decoded.DeviceStates
	s.DisplayName = decoded.DisplayName
	s.InformationUrl = decoded.InformationUrl
	s.InstallSummary = decoded.InstallSummary
	s.LargeCover = decoded.LargeCover
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.PrivacyInformationUrl = decoded.PrivacyInformationUrl
	s.PublishedDateTime = decoded.PublishedDateTime
	s.Publisher = decoded.Publisher
	s.UserStateSummary = decoded.UserStateSummary
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseManagedEBookImpl into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Assignments' for 'BaseManagedEBookImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Assignments = &output
	}

	return nil
}

func UnmarshalManagedEBookImplementation(input []byte) (ManagedEBook, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedEBook into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.iosVppEBook") {
		var out IosVppEBook
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppEBook: %+v", err)
		}
		return out, nil
	}

	var parent BaseManagedEBookImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseManagedEBookImpl: %+v", err)
	}

	return RawManagedEBookImpl{
		managedEBook: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
