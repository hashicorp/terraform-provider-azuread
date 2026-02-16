package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventExternalRegistrationInformation struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A URL or string that represents the location from which the registrant registered. Optional.
	Referrer nullable.Type[string] `json:"referrer,omitempty"`

	// The identifier for a virtualEventExternalRegistrationInformation object. Optional. If set, the maximum supported
	// length is 256 characters.
	RegistrationId nullable.Type[string] `json:"registrationId,omitempty"`
}
