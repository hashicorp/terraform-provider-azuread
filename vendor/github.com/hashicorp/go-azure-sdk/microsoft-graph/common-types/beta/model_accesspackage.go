package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackage{}

type AccessPackage struct {
	// Read-only. Nullable. Supports $expand.
	AccessPackageAssignmentPolicies *[]AccessPackageAssignmentPolicy `json:"accessPackageAssignmentPolicies,omitempty"`

	AccessPackageCatalog            *AccessPackageCatalog             `json:"accessPackageCatalog,omitempty"`
	AccessPackageResourceRoleScopes *[]AccessPackageResourceRoleScope `json:"accessPackageResourceRoleScopes,omitempty"`

	// The access packages that are incompatible with this package. Read-only.
	AccessPackagesIncompatibleWith *[]AccessPackage `json:"accessPackagesIncompatibleWith,omitempty"`

	// Identifier of the access package catalog referencing this access package. Read-only.
	CatalogId nullable.Type[string] `json:"catalogId,omitempty"`

	// The userPrincipalName of the user or identity of the subject who created this resource. Read-only.
	CreatedBy nullable.Type[string] `json:"createdBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The description of the access package.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the access package. Supports $filter (eq, contains).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The access packages whose assigned users are ineligible to be assigned this access package.
	IncompatibleAccessPackages *[]AccessPackage `json:"incompatibleAccessPackages,omitempty"`

	// The groups whose members are ineligible to be assigned this access package.
	IncompatibleGroups *[]Group `json:"incompatibleGroups,omitempty"`

	// Indicates whether the access package is hidden from the requestor.
	IsHidden nullable.Type[bool] `json:"isHidden,omitempty"`

	// Indicates whether role scopes are visible.
	IsRoleScopesVisible nullable.Type[bool] `json:"isRoleScopesVisible,omitempty"`

	// The userPrincipalName of the user who last modified this resource. Read-only.
	ModifiedBy nullable.Type[string] `json:"modifiedBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

	UniqueName nullable.Type[string] `json:"uniqueName,omitempty"`

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

func (s AccessPackage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackage{}

func (s AccessPackage) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackage: %+v", err)
	}

	delete(decoded, "accessPackageAssignmentPolicies")
	delete(decoded, "accessPackagesIncompatibleWith")
	delete(decoded, "catalogId")
	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "modifiedBy")
	delete(decoded, "modifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackage: %+v", err)
	}

	return encoded, nil
}
