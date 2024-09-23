package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LicenseUnitsDetail struct {
	// The number of units that are enabled for the active subscription of the service SKU.
	Enabled nullable.Type[int64] `json:"enabled,omitempty"`

	// The number of units that are locked out because the customer canceled their subscription of the service SKU.
	LockedOut nullable.Type[int64] `json:"lockedOut,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of units that are suspended because the subscription of the service SKU has been canceled. The units can't
	// be assigned but can still be reactivated before they're deleted.
	Suspended nullable.Type[int64] `json:"suspended,omitempty"`

	// The number of units that are in warning status. When the subscription of the service SKU has expired, the customer
	// has a grace period to renew their subscription before it's canceled (moved to a suspended state).
	Warning nullable.Type[int64] `json:"warning,omitempty"`
}
