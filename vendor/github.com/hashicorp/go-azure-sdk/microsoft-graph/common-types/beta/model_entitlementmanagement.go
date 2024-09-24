package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EntitlementManagement{}

type EntitlementManagement struct {
	AccessPackageAssignmentApprovals *[]Approval `json:"accessPackageAssignmentApprovals,omitempty"`

	// Represents the policy that governs which subjects can request or be assigned an access package via an access package
	// assignment.
	AccessPackageAssignmentPolicies *[]AccessPackageAssignmentPolicy `json:"accessPackageAssignmentPolicies,omitempty"`

	// Represents access package assignment requests created by or on behalf of a user. DO NOT USE. TO BE RETIRED SOON. Use
	// the assignmentRequests relationship instead.
	AccessPackageAssignmentRequests *[]AccessPackageAssignmentRequest `json:"accessPackageAssignmentRequests,omitempty"`

	// Represents the resource-specific role which a subject has been assigned through an access package assignment.
	AccessPackageAssignmentResourceRoles *[]AccessPackageAssignmentResourceRole `json:"accessPackageAssignmentResourceRoles,omitempty"`

	// The assignment of an access package to a subject for a period of time.
	AccessPackageAssignments *[]AccessPackageAssignment `json:"accessPackageAssignments,omitempty"`

	// A container of access packages.
	AccessPackageCatalogs *[]AccessPackageCatalog `json:"accessPackageCatalogs,omitempty"`

	// A reference to the geolocation environment in which a resource is located.
	AccessPackageResourceEnvironments *[]AccessPackageResourceEnvironment `json:"accessPackageResourceEnvironments,omitempty"`

	// Represents a request to add or remove a resource to or from a catalog respectively.
	AccessPackageResourceRequests *[]AccessPackageResourceRequest `json:"accessPackageResourceRequests,omitempty"`

	// A reference to both a scope within a resource, and a role in that resource for that scope.
	AccessPackageResourceRoleScopes *[]AccessPackageResourceRoleScope `json:"accessPackageResourceRoleScopes,omitempty"`

	// A reference to a resource associated with an access package catalog.
	AccessPackageResources *[]AccessPackageResource `json:"accessPackageResources,omitempty"`

	// Represents access package objects.
	AccessPackages *[]AccessPackage `json:"accessPackages,omitempty"`

	// Represents access package assignment requests created by or on behalf of a user.
	AssignmentRequests *[]AccessPackageAssignmentRequest `json:"assignmentRequests,omitempty"`

	// Represents references to a directory or domain of another organization whose users can request access.
	ConnectedOrganizations *[]ConnectedOrganization `json:"connectedOrganizations,omitempty"`

	// Represents the settings that control the behavior of Microsoft Entra entitlement management.
	Settings *EntitlementManagementSettings `json:"settings,omitempty"`

	// Represents the subjects within entitlement management.
	Subjects *[]AccessPackageSubject `json:"subjects,omitempty"`

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
