package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = BusinessScenario{}

type BusinessScenario struct {
	// The identity of the user who created the scenario.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time when the scenario was created. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Display name of the scenario.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The identity of the user who last modified the scenario.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The date and time when the scenario was last modified. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Identifiers of applications that are authorized to work with this scenario.
	OwnerAppIds *[]string `json:"ownerAppIds,omitempty"`

	// Planner content related to the scenario.
	Planner *BusinessScenarioPlanner `json:"planner,omitempty"`

	// Unique name of the scenario. To avoid conflicts, the recommended value for the unique name is a reverse domain name
	// format, owned by the author of the scenario. For example, a scenario authored by Contoso.com would have a unique name
	// that starts with com.contoso.
	UniqueName nullable.Type[string] `json:"uniqueName,omitempty"`

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

func (s BusinessScenario) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BusinessScenario{}

func (s BusinessScenario) MarshalJSON() ([]byte, error) {
	type wrapper BusinessScenario
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BusinessScenario: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BusinessScenario: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.businessScenario"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BusinessScenario: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BusinessScenario{}

func (s *BusinessScenario) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      *string                  `json:"createdDateTime,omitempty"`
		DisplayName          nullable.Type[string]    `json:"displayName,omitempty"`
		LastModifiedDateTime *string                  `json:"lastModifiedDateTime,omitempty"`
		OwnerAppIds          *[]string                `json:"ownerAppIds,omitempty"`
		Planner              *BusinessScenarioPlanner `json:"planner,omitempty"`
		UniqueName           nullable.Type[string]    `json:"uniqueName,omitempty"`
		Id                   *string                  `json:"id,omitempty"`
		ODataId              *string                  `json:"@odata.id,omitempty"`
		ODataType            *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.OwnerAppIds = decoded.OwnerAppIds
	s.Planner = decoded.Planner
	s.UniqueName = decoded.UniqueName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BusinessScenario into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BusinessScenario': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BusinessScenario': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
