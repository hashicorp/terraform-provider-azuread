package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageCatalog{}

type AccessPackageCatalog struct {
	// The attributes of a logic app, which can be called at various stages of an access package request and assignment
	// cycle.
	AccessPackageCustomWorkflowExtensions *[]CustomCalloutExtension `json:"accessPackageCustomWorkflowExtensions,omitempty"`

	// The roles in each resource in a catalog. Read-only.
	AccessPackageResourceRoles *[]AccessPackageResourceRole `json:"accessPackageResourceRoles,omitempty"`

	AccessPackageResourceScopes *[]AccessPackageResourceScope `json:"accessPackageResourceScopes,omitempty"`
	AccessPackageResources      *[]AccessPackageResource      `json:"accessPackageResources,omitempty"`

	// The access packages in this catalog. Read-only. Nullable. Supports $expand.
	AccessPackages *[]AccessPackage `json:"accessPackages,omitempty"`

	// Has the value Published if the access packages are available for management.
	CatalogStatus nullable.Type[string] `json:"catalogStatus,omitempty"`

	// One of UserManaged or ServiceDefault.
	CatalogType nullable.Type[string] `json:"catalogType,omitempty"`

	// UPN of the user who created this resource. Read-only.
	CreatedBy nullable.Type[string] `json:"createdBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	CustomAccessPackageWorkflowExtensions *[]CustomAccessPackageWorkflowExtension `json:"customAccessPackageWorkflowExtensions,omitempty"`

	// The description of the access package catalog.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the access package catalog. Supports $filter (eq, contains).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Whether the access packages in this catalog can be requested by users outside of the tenant.
	IsExternallyVisible nullable.Type[bool] `json:"isExternallyVisible,omitempty"`

	// The UPN of the user who last modified this resource. Read-only.
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

func (s AccessPackageCatalog) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageCatalog{}

func (s AccessPackageCatalog) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageCatalog
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageCatalog: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageCatalog: %+v", err)
	}

	delete(decoded, "accessPackageResourceRoles")
	delete(decoded, "accessPackages")
	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "modifiedBy")
	delete(decoded, "modifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageCatalog"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageCatalog: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessPackageCatalog{}

func (s *AccessPackageCatalog) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessPackageResourceRoles            *[]AccessPackageResourceRole            `json:"accessPackageResourceRoles,omitempty"`
		AccessPackageResourceScopes           *[]AccessPackageResourceScope           `json:"accessPackageResourceScopes,omitempty"`
		AccessPackageResources                *[]AccessPackageResource                `json:"accessPackageResources,omitempty"`
		AccessPackages                        *[]AccessPackage                        `json:"accessPackages,omitempty"`
		CatalogStatus                         nullable.Type[string]                   `json:"catalogStatus,omitempty"`
		CatalogType                           nullable.Type[string]                   `json:"catalogType,omitempty"`
		CreatedBy                             nullable.Type[string]                   `json:"createdBy,omitempty"`
		CreatedDateTime                       nullable.Type[string]                   `json:"createdDateTime,omitempty"`
		CustomAccessPackageWorkflowExtensions *[]CustomAccessPackageWorkflowExtension `json:"customAccessPackageWorkflowExtensions,omitempty"`
		Description                           nullable.Type[string]                   `json:"description,omitempty"`
		DisplayName                           nullable.Type[string]                   `json:"displayName,omitempty"`
		IsExternallyVisible                   nullable.Type[bool]                     `json:"isExternallyVisible,omitempty"`
		ModifiedBy                            nullable.Type[string]                   `json:"modifiedBy,omitempty"`
		ModifiedDateTime                      nullable.Type[string]                   `json:"modifiedDateTime,omitempty"`
		UniqueName                            nullable.Type[string]                   `json:"uniqueName,omitempty"`
		Id                                    *string                                 `json:"id,omitempty"`
		ODataId                               *string                                 `json:"@odata.id,omitempty"`
		ODataType                             *string                                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessPackageResourceRoles = decoded.AccessPackageResourceRoles
	s.AccessPackageResourceScopes = decoded.AccessPackageResourceScopes
	s.AccessPackageResources = decoded.AccessPackageResources
	s.AccessPackages = decoded.AccessPackages
	s.CatalogStatus = decoded.CatalogStatus
	s.CatalogType = decoded.CatalogType
	s.CreatedBy = decoded.CreatedBy
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomAccessPackageWorkflowExtensions = decoded.CustomAccessPackageWorkflowExtensions
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.IsExternallyVisible = decoded.IsExternallyVisible
	s.ModifiedBy = decoded.ModifiedBy
	s.ModifiedDateTime = decoded.ModifiedDateTime
	s.UniqueName = decoded.UniqueName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageCatalog into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["accessPackageCustomWorkflowExtensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AccessPackageCustomWorkflowExtensions into list []json.RawMessage: %+v", err)
		}

		output := make([]CustomCalloutExtension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalCustomCalloutExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AccessPackageCustomWorkflowExtensions' for 'AccessPackageCatalog': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AccessPackageCustomWorkflowExtensions = &output
	}

	return nil
}
