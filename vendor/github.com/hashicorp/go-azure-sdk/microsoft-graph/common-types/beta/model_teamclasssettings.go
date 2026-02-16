package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamClassSettings struct {
	// If set to true, enables sending of weekly assignments digest emails to parents/guardians, provided the tenant admin
	// has enabled the setting globally.
	NotifyGuardiansAboutAssignments nullable.Type[bool] `json:"notifyGuardiansAboutAssignments,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
