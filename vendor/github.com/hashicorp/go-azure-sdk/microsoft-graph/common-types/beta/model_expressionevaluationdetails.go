package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressionEvaluationDetails struct {
	// Represents expression which has been evaluated.
	Expression nullable.Type[string] `json:"expression,omitempty"`

	// Represents the details of the evaluation of the expression.
	ExpressionEvaluationDetails *[]ExpressionEvaluationDetails `json:"expressionEvaluationDetails,omitempty"`

	// Represents the value of the result of the current expression.
	ExpressionResult *bool `json:"expressionResult,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines the name of the property and the value of that property.
	PropertyToEvaluate *PropertyToEvaluate `json:"propertyToEvaluate,omitempty"`
}
