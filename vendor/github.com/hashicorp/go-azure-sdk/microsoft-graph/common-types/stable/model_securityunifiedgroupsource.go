package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityDataSource = SecurityUnifiedGroupSource{}

type SecurityUnifiedGroupSource struct {
	Group *Group `json:"group,omitempty"`

	// Specifies which sources are included in this group. Possible values are: mailbox, site.
	IncludedSources *SecuritySourceType `json:"includedSources,omitempty"`

	// Fields inherited from SecurityDataSource

	// The user who created the dataSource.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time the dataSource was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The display name of the dataSource and is the name of the SharePoint site.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The hold status of the dataSource.The possible values are: notApplied, applied, applying, removing, partial
	HoldStatus *SecurityDataSourceHoldStatus `json:"holdStatus,omitempty"`

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

func (s SecurityUnifiedGroupSource) SecurityDataSource() BaseSecurityDataSourceImpl {
	return BaseSecurityDataSourceImpl{
		CreatedBy:       s.CreatedBy,
		CreatedDateTime: s.CreatedDateTime,
		DisplayName:     s.DisplayName,
		HoldStatus:      s.HoldStatus,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s SecurityUnifiedGroupSource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityUnifiedGroupSource{}

func (s SecurityUnifiedGroupSource) MarshalJSON() ([]byte, error) {
	type wrapper SecurityUnifiedGroupSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityUnifiedGroupSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityUnifiedGroupSource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.unifiedGroupSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityUnifiedGroupSource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityUnifiedGroupSource{}

func (s *SecurityUnifiedGroupSource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Group           *Group                        `json:"group,omitempty"`
		IncludedSources *SecuritySourceType           `json:"includedSources,omitempty"`
		CreatedDateTime nullable.Type[string]         `json:"createdDateTime,omitempty"`
		DisplayName     nullable.Type[string]         `json:"displayName,omitempty"`
		HoldStatus      *SecurityDataSourceHoldStatus `json:"holdStatus,omitempty"`
		Id              *string                       `json:"id,omitempty"`
		ODataId         *string                       `json:"@odata.id,omitempty"`
		ODataType       *string                       `json:"@odata.type,omitempty"`
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
		return fmt.Errorf("unmarshaling SecurityUnifiedGroupSource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityUnifiedGroupSource': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
