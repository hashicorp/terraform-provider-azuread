package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAnswerVariant struct {
	// The answer variation description that is shown on the search results page.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The answer variation name that is displayed in search results.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The country or region that can view this answer variation.
	LanguageTag nullable.Type[string] `json:"languageTag,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The device or operating system that can view this answer variation. Possible values are: android, androidForWork,
	// ios, macOS, windowsPhone81, windowsPhone81AndLater, windows10AndLater, androidWorkProfile, unknown, androidASOP,
	// androidMobileApplicationManagement, iOSMobileApplicationManagement, unknownFutureValue.
	Platform *DevicePlatformType `json:"platform,omitempty"`

	// The URL link for the answer variation. When users select this answer variation from the search results, they're
	// directed to the specified URL.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`
}
