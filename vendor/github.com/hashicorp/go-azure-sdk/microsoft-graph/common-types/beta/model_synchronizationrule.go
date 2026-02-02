package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationRule struct {
	// The names and identifiers of organizational units that are in scope for a synchronization rule. containerFilter and
	// groupFilter are mutually exclusive properties that cannot be configured in the same request. Currently only supported
	// for Azure AD Connect cloud sync scenarios.
	ContainerFilter *ContainerFilter `json:"containerFilter,omitempty"`

	// true if the synchronization rule can be customized; false if this rule is read-only and shouldn't be changed.
	Editable *bool `json:"editable,omitempty"`

	// The names and identifiers of groups that are in scope for a synchronization rule. containerFilter and groupFilter are
	// mutually exclusive properties that cannot be configured in the same request. Currently only supported for Azure AD
	// Connect cloud sync scenarios.
	GroupFilter *GroupFilter `json:"groupFilter,omitempty"`

	// Synchronization rule identifier. Must be one of the identifiers recognized by the synchronization engine. Supported
	// rule identifiers can be found in the synchronization template returned by the API.
	Id nullable.Type[string] `json:"id,omitempty"`

	// Additional extension properties. Unless instructed explicitly by the support team, metadata values shouldn't be
	// changed.
	Metadata *[]StringKeyStringValuePair `json:"metadata,omitempty"`

	// Human-readable name of the synchronization rule. Not nullable.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Collection of object mappings supported by the rule. Tells the synchronization engine which objects should be
	// synchronized.
	ObjectMappings *[]ObjectMapping `json:"objectMappings,omitempty"`

	// Priority relative to other rules in the synchronizationSchema. Rules with the lowest priority number will be
	// processed first.
	Priority *int64 `json:"priority,omitempty"`

	// Name of the source directory. Must match one of the directory definitions in synchronizationSchema.
	SourceDirectoryName nullable.Type[string] `json:"sourceDirectoryName,omitempty"`

	// Name of the target directory. Must match one of the directory definitions in synchronizationSchema.
	TargetDirectoryName nullable.Type[string] `json:"targetDirectoryName,omitempty"`
}
