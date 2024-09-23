package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingQuestionAnswer struct {
	// The answer given by the user in case the answerInputType is text.
	Answer nullable.Type[string] `json:"answer,omitempty"`

	// The expected answer type. The possible values are: text, radioButton, unknownFutureValue.
	AnswerInputType *AnswerInputType `json:"answerInputType,omitempty"`

	// In case the answerInputType is radioButton, this will consists of a list of possible answer values.
	AnswerOptions *[]string `json:"answerOptions,omitempty"`

	// Indicates whether it is mandatory to answer the custom question.
	IsRequired nullable.Type[bool] `json:"isRequired,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The question.
	Question nullable.Type[string] `json:"question,omitempty"`

	// The ID of the custom question.
	QuestionId nullable.Type[string] `json:"questionId,omitempty"`

	// The answers selected by the user.
	SelectedOptions *[]string `json:"selectedOptions,omitempty"`
}
