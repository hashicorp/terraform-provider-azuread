package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkActivityTopic struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of source. Possible values are: entityUrl, text. For supported Microsoft Graph URLs, use entityUrl. For custom
	// text, use text.
	Source *TeamworkActivityTopicSource `json:"source,omitempty"`

	// The topic value. If the value of the source property is entityUrl, this must be a Microsoft Graph URL. If the value
	// is text, this must be a plain text value.
	Value *string `json:"value,omitempty"`

	// The link the user clicks when they select the notification. Optional when source is entityUrl; required when source
	// is text.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`
}
