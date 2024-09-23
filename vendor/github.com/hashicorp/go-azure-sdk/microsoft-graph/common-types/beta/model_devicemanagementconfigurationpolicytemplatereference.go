package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplateReference struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Template Display Name of the referenced template. This property is read-only.
	TemplateDisplayName nullable.Type[string] `json:"templateDisplayName,omitempty"`

	// Template Display Version of the referenced Template. This property is read-only.
	TemplateDisplayVersion nullable.Type[string] `json:"templateDisplayVersion,omitempty"`

	// Describes the TemplateFamily for the Template entity
	TemplateFamily *DeviceManagementConfigurationTemplateFamily `json:"templateFamily,omitempty"`

	// Template id
	TemplateId nullable.Type[string] `json:"templateId,omitempty"`
}

var _ json.Marshaler = DeviceManagementConfigurationPolicyTemplateReference{}

func (s DeviceManagementConfigurationPolicyTemplateReference) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationPolicyTemplateReference
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationPolicyTemplateReference: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationPolicyTemplateReference: %+v", err)
	}

	delete(decoded, "templateDisplayName")
	delete(decoded, "templateDisplayVersion")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationPolicyTemplateReference: %+v", err)
	}

	return encoded, nil
}
