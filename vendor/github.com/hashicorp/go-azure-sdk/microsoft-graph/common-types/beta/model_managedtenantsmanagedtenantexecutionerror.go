package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedTenantsManagedTenantOperationError = ManagedTenantsManagedTenantExecutionError{}

type ManagedTenantsManagedTenantExecutionError struct {
	// Additional error information for the exception. Optional. Read-only.
	ErrorDetails nullable.Type[string] `json:"errorDetails,omitempty"`

	// The node identifier where the exception occurred. Required. Read-only.
	NodeId *int64 `json:"nodeId,omitempty"`

	// The token for the exception. Optional. Read-only.
	RawToken nullable.Type[string] `json:"rawToken,omitempty"`

	// The statement index for the exception. Required. Read-only.
	StatementIndex *int64 `json:"statementIndex,omitempty"`

	// Fields inherited from ManagedTenantsManagedTenantOperationError

	// The error message for the exception.
	Error nullable.Type[string] `json:"error,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant.
	TenantId *string `json:"tenantId,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ManagedTenantsManagedTenantExecutionError) ManagedTenantsManagedTenantOperationError() BaseManagedTenantsManagedTenantOperationErrorImpl {
	return BaseManagedTenantsManagedTenantOperationErrorImpl{
		Error:     s.Error,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		TenantId:  s.TenantId,
	}
}

var _ json.Marshaler = ManagedTenantsManagedTenantExecutionError{}

func (s ManagedTenantsManagedTenantExecutionError) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagedTenantExecutionError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagedTenantExecutionError: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedTenantExecutionError: %+v", err)
	}

	delete(decoded, "errorDetails")
	delete(decoded, "nodeId")
	delete(decoded, "rawToken")
	delete(decoded, "statementIndex")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managedTenantExecutionError"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagedTenantExecutionError: %+v", err)
	}

	return encoded, nil
}
