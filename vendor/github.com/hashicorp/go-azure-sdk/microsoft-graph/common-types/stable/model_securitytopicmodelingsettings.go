package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTopicModelingSettings struct {
	// Indicates whether the themes model should dynamically optimize the number of generated topics. To learn more, see
	// Adjust maximum number of themes dynamically.
	DynamicallyAdjustTopicCount nullable.Type[bool] `json:"dynamicallyAdjustTopicCount,omitempty"`

	// Indicates whether the themes model should exclude numbers while parsing document texts. To learn more, see Include
	// numbers in themes.
	IgnoreNumbers nullable.Type[bool] `json:"ignoreNumbers,omitempty"`

	// Indicates whether themes model is enabled for the case.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The total number of topics that the themes model will generate for a review set. To learn more, see Maximum number of
	// themes.
	TopicCount nullable.Type[int64] `json:"topicCount,omitempty"`
}
