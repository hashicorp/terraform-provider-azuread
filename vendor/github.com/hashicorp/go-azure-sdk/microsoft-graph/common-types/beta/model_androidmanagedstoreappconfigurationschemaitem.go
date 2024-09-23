package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppConfigurationSchemaItem struct {
	// Data type for a configuration item inside an Android application's custom configuration schema
	DataType *AndroidManagedStoreAppConfigurationSchemaItemDataType `json:"dataType,omitempty"`

	// Default value for boolean type items, if specified by the app developer
	DefaultBoolValue nullable.Type[bool] `json:"defaultBoolValue,omitempty"`

	// Default value for integer type items, if specified by the app developer
	DefaultIntValue nullable.Type[int64] `json:"defaultIntValue,omitempty"`

	// Default value for string array type items, if specified by the app developer
	DefaultStringArrayValue *[]string `json:"defaultStringArrayValue,omitempty"`

	// Default value for string type items, if specified by the app developer
	DefaultStringValue nullable.Type[string] `json:"defaultStringValue,omitempty"`

	// Description of what the item controls within the application
	Description nullable.Type[string] `json:"description,omitempty"`

	// Human readable name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Unique index the application uses to maintain nested schema items
	Index *int64 `json:"index,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Index of parent schema item to track nested schema items
	ParentIndex nullable.Type[int64] `json:"parentIndex,omitempty"`

	// Unique key the application uses to identify the item
	SchemaItemKey nullable.Type[string] `json:"schemaItemKey,omitempty"`

	// List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect
	// items only)
	Selections *[]KeyValuePair `json:"selections,omitempty"`
}
