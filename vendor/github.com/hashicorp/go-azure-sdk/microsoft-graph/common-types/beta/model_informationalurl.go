package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationalUrl struct {
	// CDN URL to the application's logo, Read-only.
	LogoUrl nullable.Type[string] `json:"logoUrl,omitempty"`

	// Link to the application's marketing page. For example, https://www.contoso.com/app/marketing
	MarketingUrl nullable.Type[string] `json:"marketingUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Link to the application's privacy statement. For example, https://www.contoso.com/app/privacy
	PrivacyStatementUrl nullable.Type[string] `json:"privacyStatementUrl,omitempty"`

	// Link to the application's support page. For example, https://www.contoso.com/app/support
	SupportUrl nullable.Type[string] `json:"supportUrl,omitempty"`

	// Link to the application's terms of service statement. For example, https://www.contoso.com/app/termsofservice
	TermsOfServiceUrl nullable.Type[string] `json:"termsOfServiceUrl,omitempty"`
}

var _ json.Marshaler = InformationalUrl{}

func (s InformationalUrl) MarshalJSON() ([]byte, error) {
	type wrapper InformationalUrl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InformationalUrl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InformationalUrl: %+v", err)
	}

	delete(decoded, "logoUrl")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InformationalUrl: %+v", err)
	}

	return encoded, nil
}
