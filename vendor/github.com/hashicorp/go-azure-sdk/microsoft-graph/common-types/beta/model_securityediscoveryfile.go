package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityFile = SecurityEdiscoveryFile{}

type SecurityEdiscoveryFile struct {
	// Custodians associated with the file.
	Custodian *SecurityEdiscoveryCustodian `json:"custodian,omitempty"`

	// Tags associated with the file.
	Tags *[]SecurityEdiscoveryReviewTag `json:"tags,omitempty"`

	// Fields inherited from SecurityFile

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

func (s SecurityEdiscoveryFile) SecurityFile() BaseSecurityFileImpl {
	return BaseSecurityFileImpl{
		Content:              s.Content,
		DateTime:             s.DateTime,
		Extension:            s.Extension,
		ExtractedTextContent: s.ExtractedTextContent,
		MediaType:            s.MediaType,
		Name:                 s.Name,
		OtherProperties:      s.OtherProperties,
		ProcessingStatus:     s.ProcessingStatus,
		SenderOrAuthors:      s.SenderOrAuthors,
		Size:                 s.Size,
		SourceType:           s.SourceType,
		SubjectTitle:         s.SubjectTitle,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SecurityEdiscoveryFile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryFile{}

func (s SecurityEdiscoveryFile) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryFile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryFile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryFile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryFile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryFile: %+v", err)
	}

	return encoded, nil
}
