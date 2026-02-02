package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageCatalog{}

type AccessPackageCatalog struct {
	// The access packages in this catalog. Read-only. Nullable.
	AccessPackages *[]AccessPackage `json:"accessPackages,omitempty"`

	// Whether the catalog is created by a user or entitlement management. The possible values are: userManaged,
	// serviceDefault, serviceManaged, unknownFutureValue.
	CatalogType *AccessPackageCatalogType `json:"catalogType,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	CustomWorkflowExtensions *[]CustomCalloutExtension `json:"customWorkflowExtensions,omitempty"`

	// The description of the access package catalog.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the access package catalog.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Whether the access packages in this catalog can be requested by users outside of the tenant.
	IsExternallyVisible nullable.Type[bool] `json:"isExternallyVisible,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

	ResourceRoles  *[]AccessPackageResourceRole  `json:"resourceRoles,omitempty"`
	ResourceScopes *[]AccessPackageResourceScope `json:"resourceScopes,omitempty"`

	// Access package resources in this catalog.
	Resources *[]AccessPackageResource `json:"resources,omitempty"`

	// Has the value published if the access packages are available for management. The possible values are: unpublished,
	// published, unknownFutureValue.
	State *AccessPackageCatalogState `json:"state,omitempty"`

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

	delete(decoded, "accessPackages")
	delete(decoded, "createdDateTime")
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
		AccessPackages      *[]AccessPackage              `json:"accessPackages,omitempty"`
		CatalogType         *AccessPackageCatalogType     `json:"catalogType,omitempty"`
		CreatedDateTime     nullable.Type[string]         `json:"createdDateTime,omitempty"`
		Description         nullable.Type[string]         `json:"description,omitempty"`
		DisplayName         nullable.Type[string]         `json:"displayName,omitempty"`
		IsExternallyVisible nullable.Type[bool]           `json:"isExternallyVisible,omitempty"`
		ModifiedDateTime    nullable.Type[string]         `json:"modifiedDateTime,omitempty"`
		ResourceRoles       *[]AccessPackageResourceRole  `json:"resourceRoles,omitempty"`
		ResourceScopes      *[]AccessPackageResourceScope `json:"resourceScopes,omitempty"`
		Resources           *[]AccessPackageResource      `json:"resources,omitempty"`
		State               *AccessPackageCatalogState    `json:"state,omitempty"`
		Id                  *string                       `json:"id,omitempty"`
		ODataId             *string                       `json:"@odata.id,omitempty"`
		ODataType           *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessPackages = decoded.AccessPackages
	s.CatalogType = decoded.CatalogType
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.IsExternallyVisible = decoded.IsExternallyVisible
	s.ModifiedDateTime = decoded.ModifiedDateTime
	s.ResourceRoles = decoded.ResourceRoles
	s.ResourceScopes = decoded.ResourceScopes
	s.Resources = decoded.Resources
	s.State = decoded.State
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageCatalog into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["customWorkflowExtensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CustomWorkflowExtensions into list []json.RawMessage: %+v", err)
		}

		output := make([]CustomCalloutExtension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalCustomCalloutExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CustomWorkflowExtensions' for 'AccessPackageCatalog': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CustomWorkflowExtensions = &output
	}

	return nil
}
