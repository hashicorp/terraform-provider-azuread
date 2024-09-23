package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageAssignmentResourceRole{}

type AccessPackageAssignmentResourceRole struct {
	// The access package assignments resulting in this role assignment. Read-only. Nullable.
	AccessPackageAssignments *[]AccessPackageAssignment `json:"accessPackageAssignments,omitempty"`

	AccessPackageResourceRole  *AccessPackageResourceRole  `json:"accessPackageResourceRole,omitempty"`
	AccessPackageResourceScope *AccessPackageResourceScope `json:"accessPackageResourceScope,omitempty"`

	// Read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
	AccessPackageSubject *AccessPackageSubject `json:"accessPackageSubject,omitempty"`

	// A unique identifier relative to the origin system, corresponding to the originId property of the
	// accessPackageResourceRole.
	OriginId nullable.Type[string] `json:"originId,omitempty"`

	// The system where the role assignment is to be created or has been created for an access package assignment, such as
	// SharePointOnline, AadGroup, or AadApplication, corresponding to the originSystem property of the
	// accessPackageResourceRole.
	OriginSystem nullable.Type[string] `json:"originSystem,omitempty"`

	// The value is PendingFulfillment before the access package assignment is delivered to the origin system, and Fulfilled
	// after the access package assignment is delivered to the origin system.
	Status nullable.Type[string] `json:"status,omitempty"`

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

func (s AccessPackageAssignmentResourceRole) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageAssignmentResourceRole{}

func (s AccessPackageAssignmentResourceRole) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageAssignmentResourceRole
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageAssignmentResourceRole: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageAssignmentResourceRole: %+v", err)
	}

	delete(decoded, "accessPackageAssignments")
	delete(decoded, "accessPackageSubject")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageAssignmentResourceRole"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageAssignmentResourceRole: %+v", err)
	}

	return encoded, nil
}
