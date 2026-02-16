package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EdiscoveryDataSource = EdiscoveryUnifiedGroupSource{}

type EdiscoveryUnifiedGroupSource struct {
	Group *Group `json:"group,omitempty"`

	// Specifies which sources are included in this group. Possible values are: mailbox, site.
	IncludedSources *EdiscoverySourceType `json:"includedSources,omitempty"`

	// Fields inherited from EdiscoveryDataSource

	// The user who created the dataSource.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time the dataSource was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The display name of the dataSource, and is the name of the SharePoint site.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	HoldStatus *EdiscoveryDataSourceHoldStatus `json:"holdStatus,omitempty"`

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

func (s EdiscoveryUnifiedGroupSource) EdiscoveryDataSource() BaseEdiscoveryDataSourceImpl {
	return BaseEdiscoveryDataSourceImpl{
		CreatedBy:       s.CreatedBy,
		CreatedDateTime: s.CreatedDateTime,
		DisplayName:     s.DisplayName,
		HoldStatus:      s.HoldStatus,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s EdiscoveryUnifiedGroupSource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdiscoveryUnifiedGroupSource{}

func (s EdiscoveryUnifiedGroupSource) MarshalJSON() ([]byte, error) {
	type wrapper EdiscoveryUnifiedGroupSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdiscoveryUnifiedGroupSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryUnifiedGroupSource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.unifiedGroupSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdiscoveryUnifiedGroupSource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EdiscoveryUnifiedGroupSource{}

func (s *EdiscoveryUnifiedGroupSource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Group           *Group                          `json:"group,omitempty"`
		IncludedSources *EdiscoverySourceType           `json:"includedSources,omitempty"`
		CreatedDateTime nullable.Type[string]           `json:"createdDateTime,omitempty"`
		DisplayName     nullable.Type[string]           `json:"displayName,omitempty"`
		HoldStatus      *EdiscoveryDataSourceHoldStatus `json:"holdStatus,omitempty"`
		Id              *string                         `json:"id,omitempty"`
		ODataId         *string                         `json:"@odata.id,omitempty"`
		ODataType       *string                         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Group = decoded.Group
	s.IncludedSources = decoded.IncludedSources
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.HoldStatus = decoded.HoldStatus
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EdiscoveryUnifiedGroupSource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EdiscoveryUnifiedGroupSource': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
