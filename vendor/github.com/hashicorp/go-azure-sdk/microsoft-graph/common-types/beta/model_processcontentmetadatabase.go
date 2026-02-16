package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessContentMetadataBase interface {
	ProcessContentMetadataBase() BaseProcessContentMetadataBaseImpl
}

var _ ProcessContentMetadataBase = BaseProcessContentMetadataBaseImpl{}

type BaseProcessContentMetadataBaseImpl struct {
	// Represents the actual content, either as text (textContent) or binary data (binaryContent). Optional if metadata
	// alone is sufficient for policy evaluation. Do not use for contentActivities.
	Content ContentBase `json:"content"`

	// An GUID identifier used to group multiple related content entries (for example, different parts of the same file
	// upload, messages in a conversation).
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// Required. Timestamp indicating when the original content was created (for example, file creation time, message sent
	// time).
	CreatedDateTime string `json:"createdDateTime"`

	// Required. A unique identifier for this specific content entry within the context of the calling application or
	// enforcement plane (for example, message ID, file path/URL).
	Identifier string `json:"identifier"`

	// Required. Indicates if the provided content has been truncated from its original form (for example, due to size
	// limits).
	IsTruncated bool `json:"isTruncated"`

	// The length of the original content in bytes.
	Length nullable.Type[int64] `json:"length,omitempty"`

	// Required. Timestamp indicating when the original content was last modified. For ephemeral content like messages, this
	// might be the same as createdDateTime.
	ModifiedDateTime string `json:"modifiedDateTime"`

	// Required. A descriptive name for the content (for example, file name, web page title, 'Chat Message').
	Name string `json:"name"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A sequence number indicating the order in which content was generated or should be processed, required when
	// correlationId is used.
	SequenceNumber nullable.Type[int64] `json:"sequenceNumber,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseProcessContentMetadataBaseImpl) ProcessContentMetadataBase() BaseProcessContentMetadataBaseImpl {
	return s
}

var _ ProcessContentMetadataBase = RawProcessContentMetadataBaseImpl{}

// RawProcessContentMetadataBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawProcessContentMetadataBaseImpl struct {
	processContentMetadataBase BaseProcessContentMetadataBaseImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawProcessContentMetadataBaseImpl) ProcessContentMetadataBase() BaseProcessContentMetadataBaseImpl {
	return s.processContentMetadataBase
}

var _ json.Unmarshaler = &BaseProcessContentMetadataBaseImpl{}

func (s *BaseProcessContentMetadataBaseImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CorrelationId    nullable.Type[string] `json:"correlationId,omitempty"`
		CreatedDateTime  string                `json:"createdDateTime"`
		Identifier       string                `json:"identifier"`
		IsTruncated      bool                  `json:"isTruncated"`
		Length           nullable.Type[int64]  `json:"length,omitempty"`
		ModifiedDateTime string                `json:"modifiedDateTime"`
		Name             string                `json:"name"`
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
		SequenceNumber   nullable.Type[int64]  `json:"sequenceNumber,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CorrelationId = decoded.CorrelationId
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Identifier = decoded.Identifier
	s.IsTruncated = decoded.IsTruncated
	s.Length = decoded.Length
	s.ModifiedDateTime = decoded.ModifiedDateTime
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SequenceNumber = decoded.SequenceNumber

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseProcessContentMetadataBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["content"]; ok {
		impl, err := UnmarshalContentBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Content' for 'BaseProcessContentMetadataBaseImpl': %+v", err)
		}
		s.Content = impl
	}

	return nil
}

func UnmarshalProcessContentMetadataBaseImplementation(input []byte) (ProcessContentMetadataBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProcessContentMetadataBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.processConversationMetadata") {
		var out ProcessConversationMetadata
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProcessConversationMetadata: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.processFileMetadata") {
		var out ProcessFileMetadata
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProcessFileMetadata: %+v", err)
		}
		return out, nil
	}

	var parent BaseProcessContentMetadataBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseProcessContentMetadataBaseImpl: %+v", err)
	}

	return RawProcessContentMetadataBaseImpl{
		processContentMetadataBase: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
