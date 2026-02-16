package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiedPublisher struct {
	// The timestamp when the verified publisher was first added or most recently updated.
	AddedDateTime nullable.Type[string] `json:"addedDateTime,omitempty"`

	// The verified publisher name from the app publisher's Partner Center account.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The ID of the verified publisher from the app publisher's Partner Center account.
	VerifiedPublisherId nullable.Type[string] `json:"verifiedPublisherId,omitempty"`
}
