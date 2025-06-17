package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SolutionsRoot struct {
	Approval      *ApprovalSolution  `json:"approval,omitempty"`
	BackupRestore *BackupRestoreRoot `json:"backupRestore,omitempty"`

	// A collection of businesses in Microsoft Bookings. Read-only. Nullable.
	BookingBusinesses *[]BookingBusiness `json:"bookingBusinesses,omitempty"`

	// A collection of monetary currencies supported by a bookingBusiness. Read-only. Nullable.
	BookingCurrencies *[]BookingCurrency `json:"bookingCurrencies,omitempty"`

	// A collection of scenarios that contain relevant data and configuration information for a specific problem domain.
	BusinessScenarios *[]BusinessScenario `json:"businessScenarios,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A collection of virtual events.
	VirtualEvents *VirtualEventsRoot `json:"virtualEvents,omitempty"`
}

var _ json.Marshaler = SolutionsRoot{}

func (s SolutionsRoot) MarshalJSON() ([]byte, error) {
	type wrapper SolutionsRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SolutionsRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SolutionsRoot: %+v", err)
	}

	delete(decoded, "bookingBusinesses")
	delete(decoded, "bookingCurrencies")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SolutionsRoot: %+v", err)
	}

	return encoded, nil
}
