package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChangeTrackedEntity = WorkforceIntegration{}

type WorkforceIntegration struct {
	// API version for the call back URL. Start with 1.
	ApiVersion nullable.Type[int64] `json:"apiVersion,omitempty"`

	// Name of the workforce integration.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The workforce integration encryption resource.
	Encryption *WorkforceIntegrationEncryption `json:"encryption,omitempty"`

	// Indicates whether this workforce integration is currently active and available.
	IsActive nullable.Type[bool] `json:"isActive,omitempty"`

	// The Shifts entities supported for synchronous change notifications. Shifts will make a call back to the url provided
	// on client changes on those entities added here. By default, no entities are supported for change notifications.
	// Possible values are: none, shift, swapRequest, userShiftPreferences, openshift, openShiftRequest, offerShiftRequest,
	// unknownFutureValue.
	SupportedEntities *WorkforceIntegrationSupportedEntities `json:"supportedEntities,omitempty"`

	// Workforce Integration URL for callbacks from the Shifts service.
	Url nullable.Type[string] `json:"url,omitempty"`

	// Fields inherited from ChangeTrackedEntity

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Identity of the person who last modified the entity.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

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

func (s WorkforceIntegration) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
	return BaseChangeTrackedEntityImpl{
		CreatedDateTime:      s.CreatedDateTime,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s WorkforceIntegration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkforceIntegration{}

func (s WorkforceIntegration) MarshalJSON() ([]byte, error) {
	type wrapper WorkforceIntegration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkforceIntegration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkforceIntegration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workforceIntegration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkforceIntegration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WorkforceIntegration{}

func (s *WorkforceIntegration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApiVersion           nullable.Type[int64]                   `json:"apiVersion,omitempty"`
		DisplayName          nullable.Type[string]                  `json:"displayName,omitempty"`
		Encryption           *WorkforceIntegrationEncryption        `json:"encryption,omitempty"`
		IsActive             nullable.Type[bool]                    `json:"isActive,omitempty"`
		SupportedEntities    *WorkforceIntegrationSupportedEntities `json:"supportedEntities,omitempty"`
		Url                  nullable.Type[string]                  `json:"url,omitempty"`
		CreatedDateTime      nullable.Type[string]                  `json:"createdDateTime,omitempty"`
		LastModifiedDateTime nullable.Type[string]                  `json:"lastModifiedDateTime,omitempty"`
		Id                   *string                                `json:"id,omitempty"`
		ODataId              *string                                `json:"@odata.id,omitempty"`
		ODataType            *string                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApiVersion = decoded.ApiVersion
	s.DisplayName = decoded.DisplayName
	s.Encryption = decoded.Encryption
	s.IsActive = decoded.IsActive
	s.SupportedEntities = decoded.SupportedEntities
	s.Url = decoded.Url
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WorkforceIntegration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'WorkforceIntegration': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
