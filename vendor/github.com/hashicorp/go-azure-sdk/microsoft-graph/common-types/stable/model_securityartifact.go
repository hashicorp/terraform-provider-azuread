package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityArtifact interface {
	Entity
	SecurityArtifact() BaseSecurityArtifactImpl
}

var _ SecurityArtifact = BaseSecurityArtifactImpl{}

type BaseSecurityArtifactImpl struct {

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

func (s BaseSecurityArtifactImpl) SecurityArtifact() BaseSecurityArtifactImpl {
	return s
}

func (s BaseSecurityArtifactImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityArtifact = RawSecurityArtifactImpl{}

// RawSecurityArtifactImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityArtifactImpl struct {
	securityArtifact BaseSecurityArtifactImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawSecurityArtifactImpl) SecurityArtifact() BaseSecurityArtifactImpl {
	return s.securityArtifact
}

func (s RawSecurityArtifactImpl) Entity() BaseEntityImpl {
	return s.securityArtifact.Entity()
}

var _ json.Marshaler = BaseSecurityArtifactImpl{}

func (s BaseSecurityArtifactImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityArtifactImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityArtifactImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityArtifactImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.artifact"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityArtifactImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSecurityArtifactImplementation(input []byte) (SecurityArtifact, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityArtifact into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.host") {
		var out SecurityHost
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHost: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostComponent") {
		var out SecurityHostComponent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostComponent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostCookie") {
		var out SecurityHostCookie
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostCookie: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostSslCertificate") {
		var out SecurityHostSslCertificate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostSslCertificate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostTracker") {
		var out SecurityHostTracker
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostTracker: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.passiveDnsRecord") {
		var out SecurityPassiveDnsRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPassiveDnsRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sslCertificate") {
		var out SecuritySslCertificate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySslCertificate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.unclassifiedArtifact") {
		var out SecurityUnclassifiedArtifact
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUnclassifiedArtifact: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityArtifactImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityArtifactImpl: %+v", err)
	}

	return RawSecurityArtifactImpl{
		securityArtifact: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
