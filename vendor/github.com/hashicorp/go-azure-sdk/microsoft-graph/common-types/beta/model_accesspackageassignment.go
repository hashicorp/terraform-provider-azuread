package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageAssignment{}

type AccessPackageAssignment struct {
	// Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
	AccessPackage *AccessPackage `json:"accessPackage,omitempty"`

	// Read-only. Nullable. Supports $filter (eq) on the id property
	AccessPackageAssignmentPolicy *AccessPackageAssignmentPolicy `json:"accessPackageAssignmentPolicy,omitempty"`

	AccessPackageAssignmentRequests *[]AccessPackageAssignmentRequest `json:"accessPackageAssignmentRequests,omitempty"`

	// The resource roles delivered to the target user for this assignment. Read-only. Nullable.
	AccessPackageAssignmentResourceRoles *[]AccessPackageAssignmentResourceRole `json:"accessPackageAssignmentResourceRoles,omitempty"`

	// The identifier of the access package. Read-only.
	AccessPackageId nullable.Type[string] `json:"accessPackageId,omitempty"`

	// The identifier of the access package assignment policy. Read-only.
	AssignmentPolicyId nullable.Type[string] `json:"assignmentPolicyId,omitempty"`

	// The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only.
	// Supports $filter (eq).
	AssignmentState nullable.Type[string] `json:"assignmentState,omitempty"`

	// More information about the assignment lifecycle. Possible values include Delivering, Delivered,
	// NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered. Read-only.
	AssignmentStatus nullable.Type[string] `json:"assignmentStatus,omitempty"`

	// The identifier of the catalog containing the access package. Read-only.
	CatalogId nullable.Type[string] `json:"catalogId,omitempty"`

	// Information about all the custom extension calls that were made during the access package assignment workflow.
	CustomExtensionCalloutInstances *[]CustomExtensionCalloutInstance `json:"customExtensionCalloutInstances,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ExpiredDateTime nullable.Type[string] `json:"expiredDateTime,omitempty"`

	// Indicates whether the access package assignment is extended. Read-only.
	IsExtended nullable.Type[bool] `json:"isExtended,omitempty"`

	// When the access assignment is to be in place. Read-only.
	Schedule *RequestSchedule `json:"schedule,omitempty"`

	// The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on
	// objectId.
	Target *AccessPackageSubject `json:"target,omitempty"`

	// This property should not be used as a dependency, as it may change without notice. Instead, expand the target
	// relationship and use the objectId property. Read-only.
	TargetId nullable.Type[string] `json:"targetId,omitempty"`

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

func (s AccessPackageAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageAssignment{}

func (s AccessPackageAssignment) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageAssignment: %+v", err)
	}

	delete(decoded, "accessPackage")
	delete(decoded, "accessPackageAssignmentPolicy")
	delete(decoded, "accessPackageAssignmentResourceRoles")
	delete(decoded, "accessPackageId")
	delete(decoded, "assignmentPolicyId")
	delete(decoded, "assignmentState")
	delete(decoded, "assignmentStatus")
	delete(decoded, "catalogId")
	delete(decoded, "isExtended")
	delete(decoded, "schedule")
	delete(decoded, "target")
	delete(decoded, "targetId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageAssignment: %+v", err)
	}

	return encoded, nil
}
