package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerUserFacingMessage struct {
	// The default message displayed if the user's locale doesn't match with any of the localized messages
	DefaultMessage *string `json:"defaultMessage,omitempty"`

	// The list of <locale, message> pairs. This collection can contain a maximum of 500 elements.
	LocalizedMessages *[]KeyValuePair `json:"localizedMessages,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
