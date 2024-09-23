package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSubject struct {
	// Email of the data subject.
	Email nullable.Type[string] `json:"email,omitempty"`

	// First name of the data subject.
	FirstName nullable.Type[string] `json:"firstName,omitempty"`

	// Last Name of the data subject.
	LastName nullable.Type[string] `json:"lastName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The country/region of residency. The residency information is uesed only for internal reporting but not for the
	// content search.
	Residency nullable.Type[string] `json:"residency,omitempty"`
}
