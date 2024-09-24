package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMapping struct {
	// Default value to be used in case the source property was evaluated to null. Optional.
	DefaultValue nullable.Type[string] `json:"defaultValue,omitempty"`

	// For internal use only.
	ExportMissingReferences *bool `json:"exportMissingReferences,omitempty"`

	FlowBehavior *AttributeFlowBehavior `json:"flowBehavior,omitempty"`
	FlowType     *AttributeFlowType     `json:"flowType,omitempty"`

	// If higher than 0, this attribute will be used to perform an initial match of the objects between source and target
	// directories. The synchronization engine will try to find the matching object using attribute with lowest value of
	// matching priority first. If not found, the attribute with the next matching priority will be used, and so on a until
	// match is found or no more matching attributes are left. Only attributes that are expected to have unique values, such
	// as email, should be used as matching attributes.
	MatchingPriority *int64 `json:"matchingPriority,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines how a value should be extracted (or transformed) from the source object.
	Source *AttributeMappingSource `json:"source,omitempty"`

	// Name of the attribute on the target object.
	TargetAttributeName nullable.Type[string] `json:"targetAttributeName,omitempty"`
}
