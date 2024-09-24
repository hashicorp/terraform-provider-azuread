package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsUserFeedback struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Rating *CallRecordsUserFeedbackRating `json:"rating,omitempty"`

	// The feedback text provided by the user of this endpoint for the session.
	Text nullable.Type[string] `json:"text,omitempty"`

	// The set of feedback tokens provided by the user of this endpoint for the session. This is a set of Boolean
	// properties. The property names should not be relied upon since they may change depending on what tokens are offered
	// to the user.
	Tokens *CallRecordsFeedbackTokenSet `json:"tokens,omitempty"`
}
