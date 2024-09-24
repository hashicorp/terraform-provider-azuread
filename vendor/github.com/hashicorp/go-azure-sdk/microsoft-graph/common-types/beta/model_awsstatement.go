package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsStatement struct {
	// The AWS actions.
	Actions *[]string `json:"actions,omitempty"`

	// The AWS conditions associated with the statement.
	Condition *AwsCondition `json:"condition,omitempty"`

	Effect *AwsStatementEffect `json:"effect,omitempty"`

	// AWS Not Actions
	NotActions *[]string `json:"notActions,omitempty"`

	// AWS Not Resources
	NotResources *[]string `json:"notResources,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The AWS resources associated with the statement.
	Resources *[]string `json:"resources,omitempty"`

	// The ID of the AWS statement.
	StatementId *string `json:"statementId,omitempty"`
}
