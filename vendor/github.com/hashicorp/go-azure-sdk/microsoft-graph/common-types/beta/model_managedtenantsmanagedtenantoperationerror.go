package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantOperationError interface {
	ManagedTenantsManagedTenantOperationError() BaseManagedTenantsManagedTenantOperationErrorImpl
}

var _ ManagedTenantsManagedTenantOperationError = BaseManagedTenantsManagedTenantOperationErrorImpl{}

type BaseManagedTenantsManagedTenantOperationErrorImpl struct {
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

func (s BaseManagedTenantsManagedTenantOperationErrorImpl) ManagedTenantsManagedTenantOperationError() BaseManagedTenantsManagedTenantOperationErrorImpl {
	return s
}

var _ ManagedTenantsManagedTenantOperationError = RawManagedTenantsManagedTenantOperationErrorImpl{}

// RawManagedTenantsManagedTenantOperationErrorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawManagedTenantsManagedTenantOperationErrorImpl struct {
	managedTenantsManagedTenantOperationError BaseManagedTenantsManagedTenantOperationErrorImpl
	Type                                      string
	Values                                    map[string]interface{}
}

func (s RawManagedTenantsManagedTenantOperationErrorImpl) ManagedTenantsManagedTenantOperationError() BaseManagedTenantsManagedTenantOperationErrorImpl {
	return s.managedTenantsManagedTenantOperationError
}

func UnmarshalManagedTenantsManagedTenantOperationErrorImplementation(input []byte) (ManagedTenantsManagedTenantOperationError, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedTenantOperationError into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managedTenantExecutionError") {
		var out ManagedTenantsManagedTenantExecutionError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagedTenantExecutionError: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedTenants.managedTenantGenericError") {
		var out ManagedTenantsManagedTenantGenericError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedTenantsManagedTenantGenericError: %+v", err)
		}
		return out, nil
	}

	var parent BaseManagedTenantsManagedTenantOperationErrorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseManagedTenantsManagedTenantOperationErrorImpl: %+v", err)
	}

	return RawManagedTenantsManagedTenantOperationErrorImpl{
		managedTenantsManagedTenantOperationError: parent,
		Type:   value,
		Values: temp,
	}, nil

}
