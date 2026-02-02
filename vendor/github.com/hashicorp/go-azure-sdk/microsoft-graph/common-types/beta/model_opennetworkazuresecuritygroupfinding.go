package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Finding = OpenNetworkAzureSecurityGroupFinding{}

type OpenNetworkAzureSecurityGroupFinding struct {
	InboundPorts  InboundPorts                 `json:"inboundPorts"`
	SecurityGroup *AuthorizationSystemResource `json:"securityGroup,omitempty"`

	// Represents a virtual machine in an authorization system.
	VirtualMachines *[]VirtualMachineDetails `json:"virtualMachines,omitempty"`

	// Fields inherited from Finding

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

func (s OpenNetworkAzureSecurityGroupFinding) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s OpenNetworkAzureSecurityGroupFinding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OpenNetworkAzureSecurityGroupFinding{}

func (s OpenNetworkAzureSecurityGroupFinding) MarshalJSON() ([]byte, error) {
	type wrapper OpenNetworkAzureSecurityGroupFinding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OpenNetworkAzureSecurityGroupFinding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OpenNetworkAzureSecurityGroupFinding: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.openNetworkAzureSecurityGroupFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OpenNetworkAzureSecurityGroupFinding: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OpenNetworkAzureSecurityGroupFinding{}

func (s *OpenNetworkAzureSecurityGroupFinding) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		VirtualMachines *[]VirtualMachineDetails `json:"virtualMachines,omitempty"`
		CreatedDateTime *string                  `json:"createdDateTime,omitempty"`
		Id              *string                  `json:"id,omitempty"`
		ODataId         *string                  `json:"@odata.id,omitempty"`
		ODataType       *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.VirtualMachines = decoded.VirtualMachines
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OpenNetworkAzureSecurityGroupFinding into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["inboundPorts"]; ok {
		impl, err := UnmarshalInboundPortsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'InboundPorts' for 'OpenNetworkAzureSecurityGroupFinding': %+v", err)
		}
		s.InboundPorts = impl
	}

	if v, ok := temp["securityGroup"]; ok {
		impl, err := UnmarshalAuthorizationSystemResourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SecurityGroup' for 'OpenNetworkAzureSecurityGroupFinding': %+v", err)
		}
		s.SecurityGroup = &impl
	}

	return nil
}
