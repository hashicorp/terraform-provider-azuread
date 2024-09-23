package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EntitlementManagement{}

type EntitlementManagement struct {
	// Approval stages for decisions associated with access package assignment requests.
	AccessPackageAssignmentApprovals *[]Approval `json:"accessPackageAssignmentApprovals,omitempty"`

	// Access packages define the collection of resource roles and the policies for which subjects can request or be
	// assigned access to those resources.
	AccessPackages *[]AccessPackage `json:"accessPackages,omitempty"`

	// Access package assignment policies govern which subjects can request or be assigned an access package via an access
	// package assignment.
	AssignmentPolicies *[]AccessPackageAssignmentPolicy `json:"assignmentPolicies,omitempty"`

	// Access package assignment requests created by or on behalf of a subject.
	AssignmentRequests *[]AccessPackageAssignmentRequest `json:"assignmentRequests,omitempty"`

	// The assignment of an access package to a subject for a period of time.
	Assignments *[]AccessPackageAssignment `json:"assignments,omitempty"`

	// A container for access packages.
	Catalogs *[]AccessPackageCatalog `json:"catalogs,omitempty"`

	// References to a directory or domain of another organization whose users can request access.
	ConnectedOrganizations *[]ConnectedOrganization `json:"connectedOrganizations,omitempty"`

	// A reference to the geolocation environments in which a resource is located.
	ResourceEnvironments *[]AccessPackageResourceEnvironment `json:"resourceEnvironments,omitempty"`

	// Represents a request to add or remove a resource to or from a catalog respectively.
	ResourceRequests *[]AccessPackageResourceRequest `json:"resourceRequests,omitempty"`

	ResourceRoleScopes *[]AccessPackageResourceRoleScope `json:"resourceRoleScopes,omitempty"`

	// The resources associated with the catalogs.
	Resources *[]AccessPackageResource `json:"resources,omitempty"`

	// The settings that control the behavior of Microsoft Entra entitlement management.
	Settings *EntitlementManagementSettings `json:"settings,omitempty"`

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

func (s EntitlementManagement) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EntitlementManagement{}

func (s EntitlementManagement) MarshalJSON() ([]byte, error) {
	type wrapper EntitlementManagement
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EntitlementManagement: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EntitlementManagement: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.entitlementManagement"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EntitlementManagement: %+v", err)
	}

	return encoded, nil
}
