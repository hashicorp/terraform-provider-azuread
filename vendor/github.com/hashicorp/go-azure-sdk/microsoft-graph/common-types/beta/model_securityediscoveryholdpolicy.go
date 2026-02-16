package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityPolicyBase = SecurityEdiscoveryHoldPolicy{}

type SecurityEdiscoveryHoldPolicy struct {
	// KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search
	// conditions for Content Search and eDiscovery. To hold all content in the specified locations, leave contentQuery
	// blank.
	ContentQuery nullable.Type[string] `json:"contentQuery,omitempty"`

	// Lists any errors that happened while placing the hold.
	Errors *[]string `json:"errors,omitempty"`

	// Indicates whether the hold is enabled and actively holding content.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// Data sources that represent SharePoint sites.
	SiteSources *[]SecuritySiteSource `json:"siteSources,omitempty"`

	// Data sources that represent Exchange mailboxes.
	UserSources *[]SecurityUserSource `json:"userSources,omitempty"`

	// Fields inherited from SecurityPolicyBase

	CreatedBy            IdentitySet           `json:"createdBy"`
	CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
	Description          nullable.Type[string] `json:"description,omitempty"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	LastModifiedBy       IdentitySet           `json:"lastModifiedBy"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
	Status               *SecurityPolicyStatus `json:"status,omitempty"`

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

func (s SecurityEdiscoveryHoldPolicy) SecurityPolicyBase() BaseSecurityPolicyBaseImpl {
	return BaseSecurityPolicyBaseImpl{
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SecurityEdiscoveryHoldPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryHoldPolicy{}

func (s SecurityEdiscoveryHoldPolicy) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryHoldPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryHoldPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryHoldPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryHoldPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryHoldPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityEdiscoveryHoldPolicy{}

func (s *SecurityEdiscoveryHoldPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ContentQuery         nullable.Type[string] `json:"contentQuery,omitempty"`
		Errors               *[]string             `json:"errors,omitempty"`
		IsEnabled            nullable.Type[bool]   `json:"isEnabled,omitempty"`
		SiteSources          *[]SecuritySiteSource `json:"siteSources,omitempty"`
		UserSources          *[]SecurityUserSource `json:"userSources,omitempty"`
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		Description          nullable.Type[string] `json:"description,omitempty"`
		DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Status               *SecurityPolicyStatus `json:"status,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ContentQuery = decoded.ContentQuery
	s.Errors = decoded.Errors
	s.IsEnabled = decoded.IsEnabled
	s.SiteSources = decoded.SiteSources
	s.UserSources = decoded.UserSources
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityEdiscoveryHoldPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityEdiscoveryHoldPolicy': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SecurityEdiscoveryHoldPolicy': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
