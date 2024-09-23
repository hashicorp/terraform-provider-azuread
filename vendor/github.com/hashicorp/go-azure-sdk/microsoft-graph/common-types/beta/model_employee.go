package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Employee struct {
	Address              *PostalAddressType    `json:"address,omitempty"`
	BirthDate            nullable.Type[string] `json:"birthDate,omitempty"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	Email                nullable.Type[string] `json:"email,omitempty"`
	EmploymentDate       nullable.Type[string] `json:"employmentDate,omitempty"`
	GivenName            nullable.Type[string] `json:"givenName,omitempty"`
	Id                   *string               `json:"id,omitempty"`
	JobTitle             nullable.Type[string] `json:"jobTitle,omitempty"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
	MiddleName           nullable.Type[string] `json:"middleName,omitempty"`
	MobilePhone          nullable.Type[string] `json:"mobilePhone,omitempty"`
	Number               nullable.Type[string] `json:"number,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PersonalEmail       nullable.Type[string] `json:"personalEmail,omitempty"`
	PhoneNumber         nullable.Type[string] `json:"phoneNumber,omitempty"`
	Picture             *[]Picture            `json:"picture,omitempty"`
	StatisticsGroupCode nullable.Type[string] `json:"statisticsGroupCode,omitempty"`
	Status              nullable.Type[string] `json:"status,omitempty"`
	Surname             nullable.Type[string] `json:"surname,omitempty"`
	TerminationDate     nullable.Type[string] `json:"terminationDate,omitempty"`
}
