package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportRoot struct {
	// Container for navigation properties for Microsoft Entra authentication methods resources.
	AuthenticationMethods *AuthenticationMethodsRoot `json:"authenticationMethods,omitempty"`

	// Retrieve a list of daily print usage summaries, grouped by printer.
	DailyPrintUsageByPrinter *[]PrintUsageByPrinter `json:"dailyPrintUsageByPrinter,omitempty"`

	// Retrieve a list of daily print usage summaries, grouped by user.
	DailyPrintUsageByUser *[]PrintUsageByUser `json:"dailyPrintUsageByUser,omitempty"`

	// Retrieve a list of monthly print usage summaries, grouped by printer.
	MonthlyPrintUsageByPrinter *[]PrintUsageByPrinter `json:"monthlyPrintUsageByPrinter,omitempty"`

	// Retrieve a list of monthly print usage summaries, grouped by user.
	MonthlyPrintUsageByUser *[]PrintUsageByUser `json:"monthlyPrintUsageByUser,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents billing details for a Microsoft direct partner.
	Partners *Partners `json:"partners,omitempty"`

	// Represents an abstract type that contains resources for attack simulation and training reports.
	Security *SecurityReportsRoot `json:"security,omitempty"`
}
