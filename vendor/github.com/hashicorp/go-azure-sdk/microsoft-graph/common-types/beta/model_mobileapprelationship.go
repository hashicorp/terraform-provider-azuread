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
	// The target mobile app's display name. This property is read-only.
	TargetDisplayName nullable.Type[string] `json:"targetDisplayName,omitempty"`

	// The target mobile app's display version. This property is read-only.
	TargetDisplayVersion nullable.Type[string] `json:"targetDisplayVersion,omitempty"`

	// The target mobile app's app id.
	TargetId nullable.Type[string] `json:"targetId,omitempty"`

	// The target mobile app's publisher. This property is read-only.
	TargetPublisher nullable.Type[string] `json:"targetPublisher,omitempty"`

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

	delete(decoded, "targetDisplayName")
	delete(decoded, "targetDisplayVersion")
	delete(decoded, "targetPublisher")

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
