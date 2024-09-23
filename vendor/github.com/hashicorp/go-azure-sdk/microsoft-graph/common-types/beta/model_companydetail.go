package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyDetail struct {
	// Address of the company.
	Address *PhysicalAddress `json:"address,omitempty"`

	// Legal entity number of the company or its subdivision. For information on how to set the value for the companyCode,
	// see profileSourceAnnotation.
	CompanyCode nullable.Type[string] `json:"companyCode,omitempty"`

	// Department Name within a company.
	Department nullable.Type[string] `json:"department,omitempty"`

	// Company name.
	DisplayName *string `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Office Location of the person referred to.
	OfficeLocation nullable.Type[string] `json:"officeLocation,omitempty"`

	// Pronunciation guide for the company name.
	Pronunciation nullable.Type[string] `json:"pronunciation,omitempty"`

	SecondaryDepartment nullable.Type[string] `json:"secondaryDepartment,omitempty"`

	// Link to the company home page.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`
}
