package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsApplication struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The package security identifier that Microsoft has assigned the application. Optional. Read-only.
	PackageSid nullable.Type[string] `json:"packageSid,omitempty"`

	// Specifies the URLs where user tokens are sent for sign-in or the redirect URIs where OAuth 2.0 authorization codes
	// and access tokens are sent. Only available for applications that support the PersonalMicrosoftAccount signInAudience.
	RedirectUris *[]string `json:"redirectUris,omitempty"`
}

var _ json.Marshaler = WindowsApplication{}

func (s WindowsApplication) MarshalJSON() ([]byte, error) {
	type wrapper WindowsApplication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsApplication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsApplication: %+v", err)
	}

	delete(decoded, "packageSid")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsApplication: %+v", err)
	}

	return encoded, nil
}
