package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationQuestionAnswer struct {
	// Boolean answer of the virtual event registration question. Only appears when answerInputType is boolean.
	BooleanValue nullable.Type[bool] `json:"booleanValue,omitempty"`

	// Display name of the registration question.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Collection of text answer of the virtual event registration question. Only appears when answerInputType is
	// multiChoice.
	MultiChoiceValues *[]string `json:"multiChoiceValues,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// id of the virtual event registration question.
	QuestionId nullable.Type[string] `json:"questionId,omitempty"`

	// Text answer of the virtual event registration question. Appears when answerInputType is text, multilineText or
	// singleChoice.
	Value nullable.Type[string] `json:"value,omitempty"`
}
