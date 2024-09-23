package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BroadcastMeetingCaptionSettings struct {
	// Indicates whether captions are enabled for this Teams live event.
	IsCaptionEnabled nullable.Type[bool] `json:"isCaptionEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The spoken language.
	SpokenLanguage nullable.Type[string] `json:"spokenLanguage,omitempty"`

	// The translation languages (choose up to 6).
	TranslationLanguages *[]string `json:"translationLanguages,omitempty"`
}
