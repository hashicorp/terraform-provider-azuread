package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterClause struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Name of the operator to be applied to the source and target operands. Must be one of the supported operators.
	// Supported operators can be discovered.
	OperatorName nullable.Type[string] `json:"operatorName,omitempty"`

	// Name of source operand (the operand being tested). The source operand name must match one of the attribute names on
	// the source object.
	SourceOperandName nullable.Type[string] `json:"sourceOperandName,omitempty"`

	// Values that the source operand will be tested against.
	TargetOperand *FilterOperand `json:"targetOperand,omitempty"`
}
