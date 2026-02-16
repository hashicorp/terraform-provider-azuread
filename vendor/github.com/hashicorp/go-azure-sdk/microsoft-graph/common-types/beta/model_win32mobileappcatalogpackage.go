package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppCatalogPackage = Win32MobileAppCatalogPackage{}

type Win32MobileAppCatalogPackage struct {
	// Contains properties for Windows architecture.
	ApplicableArchitectures *WindowsArchitecture `json:"applicableArchitectures,omitempty"`

	// The product branch name, which is a specific subset of product functionality as defined by the publisher (example:
	// "Fabrikam for Business (x64)"). A specific product will have one or more branchDisplayNames. Read-only. Supports
	// $filter, $search, $select. This property is read-only.
	BranchDisplayName nullable.Type[string] `json:"branchDisplayName,omitempty"`

	// One or more locale(s) supported by the branch. Value is a two-letter ISO 639 language tags with optional two-letter
	// subtags (example: en-US, ko, de, de-DE), or mul to indicate multi-language. Read-only. This property is read-only.
	Locales *[]string `json:"locales,omitempty"`

	// Indicates whether the package is capable to auto-update to latest when software/application updates are available.
	// When TRUE, it indicates it is an auto-updating application. When FALSE, it indicates that it is not an auto-updating
	// application. This property is read-only.
	PackageAutoUpdateCapable *bool `json:"packageAutoUpdateCapable,omitempty"`

	// Fields inherited from MobileAppCatalogPackage

	// The name of the product (example: "Fabrikam for Business"). Returned by default. Read-only. Supports: $filter,
	// $search, $select. This property is read-only.
	ProductDisplayName nullable.Type[string] `json:"productDisplayName,omitempty"`

	// The identifier of a specific product irrespective of version, or other attributes. Read-only. Returned by default.
	// Supports: $filter, $select. This property is read-only.
	ProductId nullable.Type[string] `json:"productId,omitempty"`

	// The name of the application catalog package publisher (example: "Fabrikam"). Returned by default. Read-only. Supports
	// $filter, $search, $select. This property is read-only.
	PublisherDisplayName nullable.Type[string] `json:"publisherDisplayName,omitempty"`

	// The name of the product version (example: "1.2203.156"). Returned by default. Read-only. Supports: $filter, $search,
	// $select. This property is read-only.
	VersionDisplayName nullable.Type[string] `json:"versionDisplayName,omitempty"`

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

func (s Win32MobileAppCatalogPackage) MobileAppCatalogPackage() BaseMobileAppCatalogPackageImpl {
	return BaseMobileAppCatalogPackageImpl{
		ProductDisplayName:   s.ProductDisplayName,
		ProductId:            s.ProductId,
		PublisherDisplayName: s.PublisherDisplayName,
		VersionDisplayName:   s.VersionDisplayName,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s Win32MobileAppCatalogPackage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Win32MobileAppCatalogPackage{}

func (s Win32MobileAppCatalogPackage) MarshalJSON() ([]byte, error) {
	type wrapper Win32MobileAppCatalogPackage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32MobileAppCatalogPackage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32MobileAppCatalogPackage: %+v", err)
	}

	delete(decoded, "branchDisplayName")
	delete(decoded, "locales")
	delete(decoded, "packageAutoUpdateCapable")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32MobileAppCatalogPackage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32MobileAppCatalogPackage: %+v", err)
	}

	return encoded, nil
}
