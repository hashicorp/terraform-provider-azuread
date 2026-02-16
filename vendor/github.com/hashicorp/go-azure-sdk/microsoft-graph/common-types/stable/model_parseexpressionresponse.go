package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParseExpressionResponse struct {
	// Error details, if expression evaluation resulted in an error.
	Error *PublicError `json:"error,omitempty"`

	// A collection of values produced by the evaluation of the expression.
	EvaluationResult *[]string `json:"evaluationResult,omitempty"`

	// true if the evaluation was successful.
	EvaluationSucceeded *bool `json:"evaluationSucceeded,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// An attributeMappingSource object representing the parsed expression.
	ParsedExpression *AttributeMappingSource `json:"parsedExpression,omitempty"`

	// true if the expression was parsed successfully.
	ParsingSucceeded *bool `json:"parsingSucceeded,omitempty"`
}
