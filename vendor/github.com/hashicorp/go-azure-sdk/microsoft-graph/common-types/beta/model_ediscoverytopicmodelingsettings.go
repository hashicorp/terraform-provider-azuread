package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryTopicModelingSettings struct {
	// To learn more, see Adjust maximum number of themes dynamically.
	DynamicallyAdjustTopicCount nullable.Type[bool] `json:"dynamicallyAdjustTopicCount,omitempty"`

	// To learn more, see Include numbers in themes.
	IgnoreNumbers nullable.Type[bool] `json:"ignoreNumbers,omitempty"`

	// Indicates whether themes are enabled for the case.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// To learn more, see Maximum number of themes.
	TopicCount nullable.Type[int64] `json:"topicCount,omitempty"`
}
