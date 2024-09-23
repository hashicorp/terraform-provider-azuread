package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppRelationship = MobileAppDependency{}

type MobileAppDependency struct {
	// Indicates the dependency type associated with a relationship between two mobile apps.
	DependencyType *MobileAppDependencyType `json:"dependencyType,omitempty"`

	// The total number of apps that directly or indirectly depend on the parent app. This property is read-only.
	DependentAppCount *int64 `json:"dependentAppCount,omitempty"`

	// The total number of apps the child app directly or indirectly depends on. This property is read-only.
	DependsOnAppCount *int64 `json:"dependsOnAppCount,omitempty"`

	// Fields inherited from MobileAppRelationship

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

func (s MobileAppDependency) MobileAppRelationship() BaseMobileAppRelationshipImpl {
	return BaseMobileAppRelationshipImpl{
		TargetDisplayName:    s.TargetDisplayName,
		TargetDisplayVersion: s.TargetDisplayVersion,
		TargetId:             s.TargetId,
		TargetPublisher:      s.TargetPublisher,
		TargetType:           s.TargetType,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s MobileAppDependency) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobileAppDependency{}

func (s MobileAppDependency) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppDependency
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppDependency: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppDependency: %+v", err)
	}

	delete(decoded, "dependentAppCount")
	delete(decoded, "dependsOnAppCount")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppDependency"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppDependency: %+v", err)
	}

	return encoded, nil
}
