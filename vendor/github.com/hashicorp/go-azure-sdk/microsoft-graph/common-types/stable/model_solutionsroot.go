package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SolutionsRoot struct {
	BackupRestore     *BackupRestoreRoot `json:"backupRestore,omitempty"`
	BookingBusinesses *[]BookingBusiness `json:"bookingBusinesses,omitempty"`
	BookingCurrencies *[]BookingCurrency `json:"bookingCurrencies,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	VirtualEvents *VirtualEventsRoot `json:"virtualEvents,omitempty"`
}
