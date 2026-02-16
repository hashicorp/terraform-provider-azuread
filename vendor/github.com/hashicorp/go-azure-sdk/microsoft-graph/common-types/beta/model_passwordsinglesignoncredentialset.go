package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordSingleSignOnCredentialSet struct {
	// A list of credential objects that define the complete sign in flow.
	Credentials *[]Credential `json:"credentials,omitempty"`

	// The ID of the user or group this credential set belongs to.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
