package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftManagedDesktop struct {
	// Indicates the provisioning policy associated with Microsoft Managed Desktop settings. Possible values are:
	// notManaged, premiumManaged, standardManaged, starterManaged, unknownFutureValue. The default is notManaged.
	ManagedType *MicrosoftManagedDesktopType `json:"managedType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The name of the Microsoft Managed Desktop profile that the Windows 365 Cloud PC is associated with.
	Profile nullable.Type[string] `json:"profile,omitempty"`
}
