package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppRelationship = MobileAppSupersedence{}

type MobileAppSupersedence struct {
	// The total number of apps directly or indirectly superseded by the child app. Read-Only. This property is read-only.
	SupersededAppCount *int64 `json:"supersededAppCount,omitempty"`

	// Indicates the supersedence type associated with a relationship between two mobile apps.
	SupersedenceType *MobileAppSupersedenceType `json:"supersedenceType,omitempty"`

	// The total number of apps directly or indirectly superseding the parent app. Read-Only. This property is read-only.
	SupersedingAppCount *int64 `json:"supersedingAppCount,omitempty"`

	// Fields inherited from MobileAppRelationship

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

func (s MobileAppSupersedence) MobileAppRelationship() BaseMobileAppRelationshipImpl {
	return BaseMobileAppRelationshipImpl{
		SourceDisplayName:          s.SourceDisplayName,
		SourceDisplayVersion:       s.SourceDisplayVersion,
		SourceId:                   s.SourceId,
		SourcePublisherDisplayName: s.SourcePublisherDisplayName,
		TargetDisplayName:          s.TargetDisplayName,
		TargetDisplayVersion:       s.TargetDisplayVersion,
		TargetId:                   s.TargetId,
		TargetPublisher:            s.TargetPublisher,
		TargetPublisherDisplayName: s.TargetPublisherDisplayName,
		TargetType:                 s.TargetType,
		Id:                         s.Id,
		ODataId:                    s.ODataId,
		ODataType:                  s.ODataType,
	}
}

func (s MobileAppSupersedence) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobileAppSupersedence{}

func (s MobileAppSupersedence) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppSupersedence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppSupersedence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppSupersedence: %+v", err)
	}

	delete(decoded, "supersededAppCount")
	delete(decoded, "supersedingAppCount")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppSupersedence"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppSupersedence: %+v", err)
	}

	return encoded, nil
}
