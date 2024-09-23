package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInStatus struct {
	// Provides additional details on the sign-in activity
	AdditionalDetails nullable.Type[string] `json:"additionalDetails,omitempty"`

	// Provides the 5-6 digit error code that's generated during a sign-in failure. Check out the list of error codes and
	// messages.
	ErrorCode nullable.Type[int64] `json:"errorCode,omitempty"`

	// Provides the error message or the reason for failure for the corresponding sign-in activity. Check out the list of
	// error codes and messages.
	FailureReason nullable.Type[string] `json:"failureReason,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
