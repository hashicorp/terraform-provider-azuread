package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Finding interface {
	Entity
	Finding() BaseFindingImpl
}

var _ Finding = BaseFindingImpl{}

type BaseFindingImpl struct {
	// Defines when the finding was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseFindingImpl) Finding() BaseFindingImpl {
	return s
}

func (s BaseFindingImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ Finding = RawFindingImpl{}

// RawFindingImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawFindingImpl struct {
	finding BaseFindingImpl
	Type    string
	Values  map[string]interface{}
}

func (s RawFindingImpl) Finding() BaseFindingImpl {
	return s.finding
}

func (s RawFindingImpl) Entity() BaseEntityImpl {
	return s.finding.Entity()
}

var _ json.Marshaler = BaseFindingImpl{}

func (s BaseFindingImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseFindingImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseFindingImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseFindingImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.finding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseFindingImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalFindingImplementation(input []byte) (Finding, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Finding into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.awsExternalSystemAccessFinding") {
		var out AwsExternalSystemAccessFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsExternalSystemAccessFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsExternalSystemAccessRoleFinding") {
		var out AwsExternalSystemAccessRoleFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsExternalSystemAccessRoleFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsIdentityAccessManagementKeyAgeFinding") {
		var out AwsIdentityAccessManagementKeyAgeFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsIdentityAccessManagementKeyAgeFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsIdentityAccessManagementKeyUsageFinding") {
		var out AwsIdentityAccessManagementKeyUsageFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsIdentityAccessManagementKeyUsageFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsSecretInformationAccessFinding") {
		var out AwsSecretInformationAccessFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsSecretInformationAccessFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsSecurityToolAdministrationFinding") {
		var out AwsSecurityToolAdministrationFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsSecurityToolAdministrationFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.encryptedAwsStorageBucketFinding") {
		var out EncryptedAwsStorageBucketFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EncryptedAwsStorageBucketFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.encryptedAzureStorageAccountFinding") {
		var out EncryptedAzureStorageAccountFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EncryptedAzureStorageAccountFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.encryptedGcpStorageBucketFinding") {
		var out EncryptedGcpStorageBucketFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EncryptedGcpStorageBucketFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externallyAccessibleAwsStorageBucketFinding") {
		var out ExternallyAccessibleAwsStorageBucketFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternallyAccessibleAwsStorageBucketFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externallyAccessibleAzureBlobContainerFinding") {
		var out ExternallyAccessibleAzureBlobContainerFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternallyAccessibleAzureBlobContainerFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externallyAccessibleGcpStorageBucketFinding") {
		var out ExternallyAccessibleGcpStorageBucketFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternallyAccessibleGcpStorageBucketFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityFinding") {
		var out IdentityFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveGroupFinding") {
		var out InactiveGroupFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveGroupFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.openAwsSecurityGroupFinding") {
		var out OpenAwsSecurityGroupFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenAwsSecurityGroupFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.openNetworkAzureSecurityGroupFinding") {
		var out OpenNetworkAzureSecurityGroupFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenNetworkAzureSecurityGroupFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeEscalationFinding") {
		var out PrivilegeEscalationFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeEscalationFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualMachineWithAwsStorageBucketAccessFinding") {
		var out VirtualMachineWithAwsStorageBucketAccessFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualMachineWithAwsStorageBucketAccessFinding: %+v", err)
		}
		return out, nil
	}

	var parent BaseFindingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseFindingImpl: %+v", err)
	}

	return RawFindingImpl{
		finding: parent,
		Type:    value,
		Values:  temp,
	}, nil

}
