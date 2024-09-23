package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationContext struct {
	// Describes how the conditional access authentication context was triggered. A value of previouslySatisfied means the
	// auth context was because the user already satisfied the requirements for that authentication context in some previous
	// authentication event. A value of required means the user had to meet the authentication context requirement as part
	// of the sign-in flow. The possible values are: required, previouslySatisfied, notApplicable, unknownFutureValue.
	Detail *AuthenticationContextDetail `json:"detail,omitempty"`

	// The identifier of an authentication context in your tenant.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
