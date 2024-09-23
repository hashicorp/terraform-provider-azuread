package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationCustomization struct {
	// Indicates whether the display name of the resource can be overwritten by the sync.
	AllowDisplayNameUpdate nullable.Type[bool] `json:"allowDisplayNameUpdate,omitempty"`

	// Indicates whether synchronization of the parent entity is deferred to a later date.
	IsSyncDeferred nullable.Type[bool] `json:"isSyncDeferred,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The collection of property names to sync. If set to null, all properties will be synchronized. Does not apply to
	// Student Enrollments or Teacher Rosters
	OptionalPropertiesToSync *[]string `json:"optionalPropertiesToSync,omitempty"`

	// The date that the synchronization should start. This value should be set to a future date. If set to null, the
	// resource will be synchronized when the profile setup completes. Only applies to Student Enrollments
	SynchronizationStartDate nullable.Type[string] `json:"synchronizationStartDate,omitempty"`
}
