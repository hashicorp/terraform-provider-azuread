package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectMapping struct {
	// Attribute mappings define which attributes to map from the source object into the target object and how they should
	// flow. A number of functions are available to support the transformation of the original source values.
	AttributeMappings *[]AttributeMapping `json:"attributeMappings,omitempty"`

	// When true, this object mapping will be processed during synchronization. When false, this object mapping will be
	// skipped.
	Enabled *bool `json:"enabled,omitempty"`

	FlowTypes *ObjectFlowTypes `json:"flowTypes,omitempty"`

	// Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
	Metadata *[]ObjectMappingMetadataEntry `json:"metadata,omitempty"`

	// Human-friendly name of the object mapping.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines a filter to be used when deciding whether a given object should be provisioned. For example, you might want
	// to only provision users that are located in the US.
	Scope *Filter `json:"scope,omitempty"`

	// Name of the object in the source directory. Must match the object name from the source directory definition.
	SourceObjectName nullable.Type[string] `json:"sourceObjectName,omitempty"`

	// Name of the object in target directory. Must match the object name from the target directory definition.
	TargetObjectName nullable.Type[string] `json:"targetObjectName,omitempty"`
}
