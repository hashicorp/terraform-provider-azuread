package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFile interface {
	Entity
	SecurityFile() BaseSecurityFileImpl
}

var _ SecurityFile = BaseSecurityFileImpl{}

type BaseSecurityFileImpl struct {
	Content              nullable.Type[string]          `json:"content,omitempty"`
	DateTime             nullable.Type[string]          `json:"dateTime,omitempty"`
	Extension            nullable.Type[string]          `json:"extension,omitempty"`
	ExtractedTextContent nullable.Type[string]          `json:"extractedTextContent,omitempty"`
	MediaType            nullable.Type[string]          `json:"mediaType,omitempty"`
	Name                 nullable.Type[string]          `json:"name,omitempty"`
	OtherProperties      *SecurityStringValueDictionary `json:"otherProperties,omitempty"`
	ProcessingStatus     *SecurityFileProcessingStatus  `json:"processingStatus,omitempty"`
	SenderOrAuthors      *[]string                      `json:"senderOrAuthors,omitempty"`
	Size                 nullable.Type[int64]           `json:"size,omitempty"`
	SourceType           *SecuritySourceType            `json:"sourceType,omitempty"`
	SubjectTitle         nullable.Type[string]          `json:"subjectTitle,omitempty"`

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

func (s BaseSecurityFileImpl) SecurityFile() BaseSecurityFileImpl {
	return s
}

func (s BaseSecurityFileImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityFile = RawSecurityFileImpl{}

// RawSecurityFileImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityFileImpl struct {
	securityFile BaseSecurityFileImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawSecurityFileImpl) SecurityFile() BaseSecurityFileImpl {
	return s.securityFile
}

func (s RawSecurityFileImpl) Entity() BaseEntityImpl {
	return s.securityFile.Entity()
}

var _ json.Marshaler = BaseSecurityFileImpl{}

func (s BaseSecurityFileImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityFileImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityFileImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityFileImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.file"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityFileImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSecurityFileImplementation(input []byte) (SecurityFile, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityFile into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryFile") {
		var out SecurityEdiscoveryFile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryFile: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityFileImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityFileImpl: %+v", err)
	}

	return RawSecurityFileImpl{
		securityFile: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
