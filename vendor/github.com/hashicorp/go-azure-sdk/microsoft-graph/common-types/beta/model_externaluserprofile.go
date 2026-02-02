package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExternalProfile = ExternalUserProfile{}

type ExternalUserProfile struct {

	// Fields inherited from ExternalProfile

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

func (s ExternalUserProfile) ExternalProfile() BaseExternalProfileImpl {
	return BaseExternalProfileImpl{
		Address:         s.Address,
		CompanyName:     s.CompanyName,
		CreatedBy:       s.CreatedBy,
		CreatedDateTime: s.CreatedDateTime,
		Department:      s.Department,
		DisplayName:     s.DisplayName,
		IsDiscoverable:  s.IsDiscoverable,
		IsEnabled:       s.IsEnabled,
		JobTitle:        s.JobTitle,
		PhoneNumber:     s.PhoneNumber,
		SupervisorId:    s.SupervisorId,
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s ExternalUserProfile) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s ExternalUserProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalUserProfile{}

func (s ExternalUserProfile) MarshalJSON() ([]byte, error) {
	type wrapper ExternalUserProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalUserProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalUserProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalUserProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalUserProfile: %+v", err)
	}

	return encoded, nil
}
