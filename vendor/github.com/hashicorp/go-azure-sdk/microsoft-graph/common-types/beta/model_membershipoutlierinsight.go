package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GovernanceInsight = MembershipOutlierInsight{}

type MembershipOutlierInsight struct {
	// Navigation link to the container directory object. For example, to a group.
	Container *DirectoryObject `json:"container,omitempty"`

	// Indicates the identifier of the container, for example, a group ID.
	ContainerId *string `json:"containerId,omitempty"`

	// OData ID for `Container` to bind to this entity
	Container_ODataBind *string `json:"container@odata.bind,omitempty"`

	// Navigation link to a member object who modified the record. For example, to a user.
	LastModifiedBy *User `json:"lastModifiedBy,omitempty"`

	// Navigation link to a member object. For example, to a user.
	Member *DirectoryObject `json:"member,omitempty"`

	// Indicates the identifier of the user.
	MemberId *string `json:"memberId,omitempty"`

	// OData ID for `Member` to bind to this entity
	Member_ODataBind *string `json:"member@odata.bind,omitempty"`

	OutlierContainerType *OutlierContainerType `json:"outlierContainerType,omitempty"`
	OutlierMemberType    *OutlierMemberType    `json:"outlierMemberType,omitempty"`

	// Fields inherited from GovernanceInsight

	// Indicates when the insight was created.
	InsightCreatedDateTime nullable.Type[string] `json:"insightCreatedDateTime,omitempty"`

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

func (s MembershipOutlierInsight) GovernanceInsight() BaseGovernanceInsightImpl {
	return BaseGovernanceInsightImpl{
		InsightCreatedDateTime: s.InsightCreatedDateTime,
		Id:                     s.Id,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
	}
}

func (s MembershipOutlierInsight) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MembershipOutlierInsight{}

func (s MembershipOutlierInsight) MarshalJSON() ([]byte, error) {
	type wrapper MembershipOutlierInsight
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MembershipOutlierInsight: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MembershipOutlierInsight: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.membershipOutlierInsight"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MembershipOutlierInsight: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MembershipOutlierInsight{}

func (s *MembershipOutlierInsight) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ContainerId            *string               `json:"containerId,omitempty"`
		Container_ODataBind    *string               `json:"container@odata.bind,omitempty"`
		LastModifiedBy         *User                 `json:"lastModifiedBy,omitempty"`
		MemberId               *string               `json:"memberId,omitempty"`
		Member_ODataBind       *string               `json:"member@odata.bind,omitempty"`
		OutlierContainerType   *OutlierContainerType `json:"outlierContainerType,omitempty"`
		OutlierMemberType      *OutlierMemberType    `json:"outlierMemberType,omitempty"`
		InsightCreatedDateTime nullable.Type[string] `json:"insightCreatedDateTime,omitempty"`
		Id                     *string               `json:"id,omitempty"`
		ODataId                *string               `json:"@odata.id,omitempty"`
		ODataType              *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ContainerId = decoded.ContainerId
	s.Container_ODataBind = decoded.Container_ODataBind
	s.LastModifiedBy = decoded.LastModifiedBy
	s.MemberId = decoded.MemberId
	s.Member_ODataBind = decoded.Member_ODataBind
	s.OutlierContainerType = decoded.OutlierContainerType
	s.OutlierMemberType = decoded.OutlierMemberType
	s.Id = decoded.Id
	s.InsightCreatedDateTime = decoded.InsightCreatedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MembershipOutlierInsight into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["container"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Container' for 'MembershipOutlierInsight': %+v", err)
		}
		s.Container = &impl
	}

	if v, ok := temp["member"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Member' for 'MembershipOutlierInsight': %+v", err)
		}
		s.Member = &impl
	}

	return nil
}
