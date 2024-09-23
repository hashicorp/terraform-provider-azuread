package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticRepliesMailTips struct {
	// The automatic reply message.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The language that the automatic reply message is in.
	MessageLanguage *LocaleInfo `json:"messageLanguage,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date and time that automatic replies are set to end.
	ScheduledEndTime *DateTimeTimeZone `json:"scheduledEndTime,omitempty"`

	// The date and time that automatic replies are set to begin.
	ScheduledStartTime *DateTimeTimeZone `json:"scheduledStartTime,omitempty"`
}
