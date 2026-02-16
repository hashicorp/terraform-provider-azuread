package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookWorksheetProtectionOptions struct {
	// Represents the worksheet protection option of allowing using auto filter feature.
	AllowAutoFilter *bool `json:"allowAutoFilter,omitempty"`

	// Represents the worksheet protection option of allowing deleting columns.
	AllowDeleteColumns *bool `json:"allowDeleteColumns,omitempty"`

	// Represents the worksheet protection option of allowing deleting rows.
	AllowDeleteRows *bool `json:"allowDeleteRows,omitempty"`

	// Represents the worksheet protection option of allowing formatting cells.
	AllowFormatCells *bool `json:"allowFormatCells,omitempty"`

	// Represents the worksheet protection option of allowing formatting columns.
	AllowFormatColumns *bool `json:"allowFormatColumns,omitempty"`

	// Represents the worksheet protection option of allowing formatting rows.
	AllowFormatRows *bool `json:"allowFormatRows,omitempty"`

	// Represents the worksheet protection option of allowing inserting columns.
	AllowInsertColumns *bool `json:"allowInsertColumns,omitempty"`

	// Represents the worksheet protection option of allowing inserting hyperlinks.
	AllowInsertHyperlinks *bool `json:"allowInsertHyperlinks,omitempty"`

	// Represents the worksheet protection option of allowing inserting rows.
	AllowInsertRows *bool `json:"allowInsertRows,omitempty"`

	// Represents the worksheet protection option of allowing using pivot table feature.
	AllowPivotTables *bool `json:"allowPivotTables,omitempty"`

	// Represents the worksheet protection option of allowing using sort feature.
	AllowSort *bool `json:"allowSort,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
