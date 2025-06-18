package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppRelationship interface {
	Entity
	MobileAppRelationship() BaseMobileAppRelationshipImpl
}

var _ MobileAppRelationship = BaseMobileAppRelationshipImpl{}

type BaseMobileAppRelationshipImpl struct {
	// The display name of the app that is the source of the mobile app relationship entity. For example: Orca. Maximum
	// length is 500 characters. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is
	// read-only.
	SourceDisplayName nullable.Type[string] `json:"sourceDisplayName,omitempty"`

	// The display version of the app that is the source of the mobile app relationship entity. For example 1.0.12 or
	// 1.2203.156 or 3. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is
	// read-only.
	SourceDisplayVersion nullable.Type[string] `json:"sourceDisplayVersion,omitempty"`

	// The unique app identifier of the source of the mobile app relationship entity. For example:
	// 2dbc75b9-e993-4e4d-a071-91ac5a218672. If null during relationship creation, then it will be populated with parent Id.
	// Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is read-only.
	SourceId nullable.Type[string] `json:"sourceId,omitempty"`

	// The publisher display name of the app that is the source of the mobile app relationship entity. For example:
	// Fabrikam. Maximum length is 500 characters. Read-Only. Supports: $select. Does not support $search, $filter,
	// $orderBy. This property is read-only.
	SourcePublisherDisplayName nullable.Type[string] `json:"sourcePublisherDisplayName,omitempty"`

	// The display name of the app that is the target of the mobile app relationship entity. For example: Firefox Setup
	// 52.0.2 32bit.intunewin. Maximum length is 500 characters. Read-Only. Returned by default. Supports: $select. Does not
	// support $search, $filter, $orderBy. This property is read-only.
	TargetDisplayName nullable.Type[string] `json:"targetDisplayName,omitempty"`

	// The display version of the app that is the target of the mobile app relationship entity. For example 1.0 or
	// 1.2203.156. Read-Only. Returned by default. Supports: $select. Does not support $search, $filter, $orderBy. This
	// property is read-only.
	TargetDisplayVersion nullable.Type[string] `json:"targetDisplayVersion,omitempty"`

	// The unique app identifier of the target of the mobile app relationship entity. For example:
	// 2dbc75b9-e993-4e4d-a071-91ac5a218672. Read-Only. Returned by default. Supports: $select. Does not support $search,
	// $filter, $orderBy.
	TargetId nullable.Type[string] `json:"targetId,omitempty"`

	// The publisher of the app that is the target of the mobile app relationship entity. For example: Fabrikam. Maximum
	// length is 500 characters. Read-Only. Returned by default. Supports: $select. Does not support $search, $filter,
	// $orderBy. This property is read-only.
	TargetPublisher nullable.Type[string] `json:"targetPublisher,omitempty"`

	// The publisher display name of the app that is the target of the mobile app relationship entity. For example:
	// Fabrikam. Maximum length is 500 characters. Read-Only. Supports: $select. Does not support $search, $filter,
	// $orderBy. This property is read-only.
	TargetPublisherDisplayName nullable.Type[string] `json:"targetPublisherDisplayName,omitempty"`

	// Indicates whether the target of a relationship is the parent or the child in the relationship.
	TargetType *MobileAppRelationshipType `json:"targetType,omitempty"`

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

func (s BaseMobileAppRelationshipImpl) MobileAppRelationship() BaseMobileAppRelationshipImpl {
	return s
}

func (s BaseMobileAppRelationshipImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ MobileAppRelationship = RawMobileAppRelationshipImpl{}

// RawMobileAppRelationshipImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMobileAppRelationshipImpl struct {
	mobileAppRelationship BaseMobileAppRelationshipImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawMobileAppRelationshipImpl) MobileAppRelationship() BaseMobileAppRelationshipImpl {
	return s.mobileAppRelationship
}

func (s RawMobileAppRelationshipImpl) Entity() BaseEntityImpl {
	return s.mobileAppRelationship.Entity()
}

var _ json.Marshaler = BaseMobileAppRelationshipImpl{}

func (s BaseMobileAppRelationshipImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseMobileAppRelationshipImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseMobileAppRelationshipImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseMobileAppRelationshipImpl: %+v", err)
	}

	delete(decoded, "sourceDisplayName")
	delete(decoded, "sourceDisplayVersion")
	delete(decoded, "sourceId")
	delete(decoded, "sourcePublisherDisplayName")
	delete(decoded, "targetDisplayName")
	delete(decoded, "targetDisplayVersion")
	delete(decoded, "targetPublisher")
	delete(decoded, "targetPublisherDisplayName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppRelationship"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseMobileAppRelationshipImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalMobileAppRelationshipImplementation(input []byte) (MobileAppRelationship, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppRelationship into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppDependency") {
		var out MobileAppDependency
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppDependency: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppSupersedence") {
		var out MobileAppSupersedence
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppSupersedence: %+v", err)
		}
		return out, nil
	}

	var parent BaseMobileAppRelationshipImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMobileAppRelationshipImpl: %+v", err)
	}

	return RawMobileAppRelationshipImpl{
		mobileAppRelationship: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
