package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChangeTrackedEntity = WorkforceIntegration{}

type WorkforceIntegration struct {
	// API version for the callback URL. Start with 1.
	ApiVersion nullable.Type[int64] `json:"apiVersion,omitempty"`

	// Name of the workforce integration.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	EligibilityFilteringEnabledEntities *EligibilityFilteringEnabledEntities `json:"eligibilityFilteringEnabledEntities,omitempty"`

	// The workforce integration encryption resource.
	Encryption *WorkforceIntegrationEncryption `json:"encryption,omitempty"`

	// Indicates whether this workforce integration is currently active and available.
	IsActive nullable.Type[bool] `json:"isActive,omitempty"`

	// This property has replaced supports in v1.0. We recommend that you use this property instead of supports. The
	// supports property is still supported in beta for the time being. The possible values are: none, shift, swapRequest,
	// openshift, openShiftRequest, userShiftPreferences, offerShiftRequest, unknownFutureValue, timeCard, timeOffReason,
	// timeOff, timeOffRequest. You must use the Prefer: include-unknown-enum-members request header to get the following
	// values in this evolvable enum: timeCard, timeOffReason, timeOff, timeOffRequest. If selecting more than one value,
	// all values must start with the first letter in uppercase.
	SupportedEntities *WorkforceIntegrationSupportedEntities `json:"supportedEntities,omitempty"`

	// The Shifts entities supported for synchronous change notifications. Shifts make a callback to the url provided on
	// client changes on those entities added here. By default, no entities are supported for change notifications. The
	// possible values are: none, shift, swapRequest, openshift, openShiftRequest, userShiftPreferences, offerShiftRequest,
	// unknownFutureValue, timeCard, timeOffReason, timeOff, timeOffRequest. You must use the Prefer:
	// include-unknown-enum-members request header to get the following values in this evolvable enum: timeCard,
	// timeOffReason, timeOff, timeOffRequest. If selecting more than one value, all values must start with the first letter
	// in uppercase.
	Supports *WorkforceIntegrationSupportedEntities `json:"supports,omitempty"`

	// Workforce Integration URL for callbacks from the Shifts service.
	Url nullable.Type[string] `json:"url,omitempty"`

	// Fields inherited from ChangeTrackedEntity

	// Identity of the user who created the entity.
	CreatedBy IdentitySet `json:"createdBy"`

	// The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Identity of the user who last modified the entity.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
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
		CreatedBy:            s.CreatedBy,
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
		ApiVersion                          nullable.Type[int64]                   `json:"apiVersion,omitempty"`
		DisplayName                         nullable.Type[string]                  `json:"displayName,omitempty"`
		EligibilityFilteringEnabledEntities *EligibilityFilteringEnabledEntities   `json:"eligibilityFilteringEnabledEntities,omitempty"`
		Encryption                          *WorkforceIntegrationEncryption        `json:"encryption,omitempty"`
		IsActive                            nullable.Type[bool]                    `json:"isActive,omitempty"`
		SupportedEntities                   *WorkforceIntegrationSupportedEntities `json:"supportedEntities,omitempty"`
		Supports                            *WorkforceIntegrationSupportedEntities `json:"supports,omitempty"`
		Url                                 nullable.Type[string]                  `json:"url,omitempty"`
		CreatedDateTime                     nullable.Type[string]                  `json:"createdDateTime,omitempty"`
		LastModifiedDateTime                nullable.Type[string]                  `json:"lastModifiedDateTime,omitempty"`
		Id                                  *string                                `json:"id,omitempty"`
		ODataId                             *string                                `json:"@odata.id,omitempty"`
		ODataType                           *string                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApiVersion = decoded.ApiVersion
	s.DisplayName = decoded.DisplayName
	s.EligibilityFilteringEnabledEntities = decoded.EligibilityFilteringEnabledEntities
	s.Encryption = decoded.Encryption
	s.IsActive = decoded.IsActive
	s.SupportedEntities = decoded.SupportedEntities
	s.Supports = decoded.Supports
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

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'WorkforceIntegration': %+v", err)
		}
		s.CreatedBy = impl
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
