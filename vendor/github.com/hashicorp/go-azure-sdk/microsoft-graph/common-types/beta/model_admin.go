package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Admin struct {
	AppsAndServices *AdminAppsAndServices `json:"appsAndServices,omitempty"`
	Dynamics        *AdminDynamics        `json:"dynamics,omitempty"`

	// A container for Microsoft Edge resources. Read-only.
	Edge *Edge `json:"edge,omitempty"`

	Forms *AdminForms `json:"forms,omitempty"`

	// A container for the Microsoft 365 apps admin functionality.
	Microsoft365Apps *AdminMicrosoft365Apps `json:"microsoft365Apps,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents a setting to control people-related admin settings in the tenant.
	People *PeopleAdminSettings `json:"people,omitempty"`

	// A container for administrative resources to manage reports.
	ReportSettings *AdminReportSettings `json:"reportSettings,omitempty"`

	// A container for service communications resources. Read-only.
	ServiceAnnouncement *ServiceAnnouncement `json:"serviceAnnouncement,omitempty"`

	Sharepoint *Sharepoint `json:"sharepoint,omitempty"`
	Todo       *AdminTodo  `json:"todo,omitempty"`

	// A container for all Windows administrator functionalities. Read-only.
	Windows *AdminWindows `json:"windows,omitempty"`
}

var _ json.Marshaler = Admin{}

func (s Admin) MarshalJSON() ([]byte, error) {
	type wrapper Admin
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Admin: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Admin: %+v", err)
	}

	delete(decoded, "edge")
	delete(decoded, "serviceAnnouncement")
	delete(decoded, "windows")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Admin: %+v", err)
	}

	return encoded, nil
}
