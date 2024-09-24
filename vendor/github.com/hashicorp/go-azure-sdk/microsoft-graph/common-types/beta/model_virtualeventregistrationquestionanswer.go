package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationQuestionAnswer struct {
	// Boolean answer to the virtualEventRegistrationCustomQuestion. Only appears when answerInputType is boolean.
	BooleanValue nullable.Type[bool] `json:"booleanValue,omitempty"`

	// Display name of the registration question.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// A collection of text answers to the virtualEventRegistrationCustomQuestion. Only appears when answerInputType is
	// multiChoice.
	MultiChoiceValues *[]string `json:"multiChoiceValues,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identifier of either a virtualEventRegistrationCustomQuestion or a virtualEventRegistrationPredefinedQuestion.
	QuestionId nullable.Type[string] `json:"questionId,omitempty"`

	// Text answer to the virtualEventRegistrationCustomQuestion or the virtualEventRegistrationPredefinedQuestion. Appears
	// when answerInputType is text, multilineText or singleChoice.
	Value nullable.Type[string] `json:"value,omitempty"`
}
