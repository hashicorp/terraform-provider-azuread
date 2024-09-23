package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppCatalogPackage interface {
	Entity
	MobileAppCatalogPackage() BaseMobileAppCatalogPackageImpl
}

var _ MobileAppCatalogPackage = BaseMobileAppCatalogPackageImpl{}

type BaseMobileAppCatalogPackageImpl struct {
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

func (s BaseMobileAppCatalogPackageImpl) MobileAppCatalogPackage() BaseMobileAppCatalogPackageImpl {
	return s
}

func (s BaseMobileAppCatalogPackageImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ MobileAppCatalogPackage = RawMobileAppCatalogPackageImpl{}

// RawMobileAppCatalogPackageImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMobileAppCatalogPackageImpl struct {
	mobileAppCatalogPackage BaseMobileAppCatalogPackageImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawMobileAppCatalogPackageImpl) MobileAppCatalogPackage() BaseMobileAppCatalogPackageImpl {
	return s.mobileAppCatalogPackage
}

func (s RawMobileAppCatalogPackageImpl) Entity() BaseEntityImpl {
	return s.mobileAppCatalogPackage.Entity()
}

var _ json.Marshaler = BaseMobileAppCatalogPackageImpl{}

func (s BaseMobileAppCatalogPackageImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseMobileAppCatalogPackageImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseMobileAppCatalogPackageImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseMobileAppCatalogPackageImpl: %+v", err)
	}

	delete(decoded, "productDisplayName")
	delete(decoded, "productId")
	delete(decoded, "publisherDisplayName")
	delete(decoded, "versionDisplayName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppCatalogPackage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseMobileAppCatalogPackageImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalMobileAppCatalogPackageImplementation(input []byte) (MobileAppCatalogPackage, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppCatalogPackage into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.win32MobileAppCatalogPackage") {
		var out Win32MobileAppCatalogPackage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32MobileAppCatalogPackage: %+v", err)
		}
		return out, nil
	}

	var parent BaseMobileAppCatalogPackageImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMobileAppCatalogPackageImpl: %+v", err)
	}

	return RawMobileAppCatalogPackageImpl{
		mobileAppCatalogPackage: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
