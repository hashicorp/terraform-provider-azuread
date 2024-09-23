package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceLevelAgreementAttainment struct {
	// The end date for the calendar month for which SLA attainment is measured.
	EndDate *string `json:"endDate,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The start date for the calendar month for which SLA attainment is measured.
	StartDate *string `json:"startDate,omitempty"`
}
