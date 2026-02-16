package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EdiscoveryLegalHold{}

type EdiscoveryLegalHold struct {
	// KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search
	// conditions for Content Search and eDiscovery. To hold all content in the specified locations, leave contentQuery
	// blank.
	ContentQuery nullable.Type[string] `json:"contentQuery,omitempty"`

	// The user who created the legal hold.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time the legal hold was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The legal hold description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the legal hold.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Lists any errors that happened while placing the hold.
	Errors *[]string `json:"errors,omitempty"`

	// Indicates whether the hold is enabled and actively holding content.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// the user who last modified the legal hold.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The date and time the legal hold was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Data source entity for SharePoint sites associated with the legal hold.
	SiteSources *[]EdiscoverySiteSource `json:"siteSources,omitempty"`

	// The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
	Status *EdiscoveryLegalHoldStatus `json:"status,omitempty"`

	UnifiedGroupSources *[]EdiscoveryUnifiedGroupSource `json:"unifiedGroupSources,omitempty"`

	// Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
	UserSources *[]EdiscoveryUserSource `json:"userSources,omitempty"`

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

func (s EdiscoveryLegalHold) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdiscoveryLegalHold{}

func (s EdiscoveryLegalHold) MarshalJSON() ([]byte, error) {
	type wrapper EdiscoveryLegalHold
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdiscoveryLegalHold: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryLegalHold: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.legalHold"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdiscoveryLegalHold: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EdiscoveryLegalHold{}

func (s *EdiscoveryLegalHold) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ContentQuery         nullable.Type[string]           `json:"contentQuery,omitempty"`
		CreatedDateTime      nullable.Type[string]           `json:"createdDateTime,omitempty"`
		Description          nullable.Type[string]           `json:"description,omitempty"`
		DisplayName          nullable.Type[string]           `json:"displayName,omitempty"`
		Errors               *[]string                       `json:"errors,omitempty"`
		IsEnabled            nullable.Type[bool]             `json:"isEnabled,omitempty"`
		LastModifiedDateTime nullable.Type[string]           `json:"lastModifiedDateTime,omitempty"`
		SiteSources          *[]EdiscoverySiteSource         `json:"siteSources,omitempty"`
		Status               *EdiscoveryLegalHoldStatus      `json:"status,omitempty"`
		UnifiedGroupSources  *[]EdiscoveryUnifiedGroupSource `json:"unifiedGroupSources,omitempty"`
		UserSources          *[]EdiscoveryUserSource         `json:"userSources,omitempty"`
		Id                   *string                         `json:"id,omitempty"`
		ODataId              *string                         `json:"@odata.id,omitempty"`
		ODataType            *string                         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ContentQuery = decoded.ContentQuery
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Errors = decoded.Errors
	s.IsEnabled = decoded.IsEnabled
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.SiteSources = decoded.SiteSources
	s.Status = decoded.Status
	s.UnifiedGroupSources = decoded.UnifiedGroupSources
	s.UserSources = decoded.UserSources
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EdiscoveryLegalHold into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EdiscoveryLegalHold': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'EdiscoveryLegalHold': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
