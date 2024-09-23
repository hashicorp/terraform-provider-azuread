package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedTenantsManagedTenantOperationError = ManagedTenantsManagedTenantGenericError{}

type ManagedTenantsManagedTenantGenericError struct {

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

func (s ManagedTenantsManagedTenantGenericError) ManagedTenantsManagedTenantOperationError() BaseManagedTenantsManagedTenantOperationErrorImpl {
	return BaseManagedTenantsManagedTenantOperationErrorImpl{
		Error:     s.Error,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		TenantId:  s.TenantId,
	}
}

var _ json.Marshaler = ManagedTenantsManagedTenantGenericError{}

func (s ManagedTenantsManagedTenantGenericError) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagedTenantGenericError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagedTenantGenericError: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedTenantGenericError: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managedTenantGenericError"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagedTenantGenericError: %+v", err)
	}

	return encoded, nil
}
