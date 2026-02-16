package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataAdditionalUserOptions struct {
	// Indicates whether student contact association should be allowed.
	AllowStudentContactAssociation nullable.Type[bool] `json:"allowStudentContactAssociation,omitempty"`

	// Indicates whether all students should be marked as minors.
	MarkAllStudentsAsMinors *bool `json:"markAllStudentsAsMinors,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
