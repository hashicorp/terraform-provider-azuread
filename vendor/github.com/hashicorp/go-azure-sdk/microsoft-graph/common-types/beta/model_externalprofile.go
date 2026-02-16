package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalProfile interface {
	Entity
	DirectoryObject
	ExternalProfile() BaseExternalProfileImpl
}

var _ ExternalProfile = BaseExternalProfileImpl{}

type BaseExternalProfileImpl struct {
	// The office address of the external user profile.
	Address *PhysicalOfficeAddress `json:"address,omitempty"`

	// The company name of the external user profile. Supports $filter (eq, startswith).
	CompanyName nullable.Type[string] `json:"companyName,omitempty"`

	// The object ID of the user who created the external user profile. Read-only. Not nullable.
	CreatedBy nullable.Type[string] `json:"createdBy,omitempty"`

	// Date and time when this external user was created. Not nullable. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The department of the external user profile.
	Department nullable.Type[string] `json:"department,omitempty"`

	// The display name of the external user profile.
	DisplayName *string `json:"displayName,omitempty"`

	// Represents whether the external user profile is discoverable in the directory. When true, this external profile shows
	// up in Teams search.
	IsDiscoverable nullable.Type[bool] `json:"isDiscoverable,omitempty"`

	// Represents whether the external user profile is enabled in the directory. This property is peer to the accountEnabled
	// property on the user object.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The job title of the external user profile.
	JobTitle nullable.Type[string] `json:"jobTitle,omitempty"`

	// The phone number of the external user profile. Must be in E164 format.
	PhoneNumber nullable.Type[string] `json:"phoneNumber,omitempty"`

	// The object ID of the supervisor of the external user profile. Supports $filter (eq, startswith).
	SupervisorId nullable.Type[string] `json:"supervisorId,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s BaseExternalProfileImpl) ExternalProfile() BaseExternalProfileImpl {
	return s
}

func (s BaseExternalProfileImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseExternalProfileImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ExternalProfile = RawExternalProfileImpl{}

// RawExternalProfileImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawExternalProfileImpl struct {
	externalProfile BaseExternalProfileImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawExternalProfileImpl) ExternalProfile() BaseExternalProfileImpl {
	return s.externalProfile
}

func (s RawExternalProfileImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return s.externalProfile.DirectoryObject()
}

func (s RawExternalProfileImpl) Entity() BaseEntityImpl {
	return s.externalProfile.Entity()
}

var _ json.Marshaler = BaseExternalProfileImpl{}

func (s BaseExternalProfileImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseExternalProfileImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseExternalProfileImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseExternalProfileImpl: %+v", err)
	}

	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseExternalProfileImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalExternalProfileImplementation(input []byte) (ExternalProfile, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalProfile into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.externalUserProfile") {
		var out ExternalUserProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalUserProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.pendingExternalUserProfile") {
		var out PendingExternalUserProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PendingExternalUserProfile: %+v", err)
		}
		return out, nil
	}

	var parent BaseExternalProfileImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseExternalProfileImpl: %+v", err)
	}

	return RawExternalProfileImpl{
		externalProfile: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
