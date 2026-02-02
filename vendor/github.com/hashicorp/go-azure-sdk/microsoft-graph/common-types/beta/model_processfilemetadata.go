package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProcessContentMetadataBase = ProcessFileMetadata{}

type ProcessFileMetadata struct {
	// A dictionary containing custom metadata associated with the file, potentially extracted by the calling application.
	CustomProperties *CustomMetadataDictionary `json:"customProperties,omitempty"`

	// The unique identifier (for example, Object ID or UPN) of the owner of the file.
	OwnerId nullable.Type[string] `json:"ownerId,omitempty"`

	// Fields inherited from ProcessContentMetadataBase

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

func (s ProcessFileMetadata) ProcessContentMetadataBase() BaseProcessContentMetadataBaseImpl {
	return BaseProcessContentMetadataBaseImpl{
		Content:          s.Content,
		CorrelationId:    s.CorrelationId,
		CreatedDateTime:  s.CreatedDateTime,
		Identifier:       s.Identifier,
		IsTruncated:      s.IsTruncated,
		Length:           s.Length,
		ModifiedDateTime: s.ModifiedDateTime,
		Name:             s.Name,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
		SequenceNumber:   s.SequenceNumber,
	}
}

var _ json.Marshaler = ProcessFileMetadata{}

func (s ProcessFileMetadata) MarshalJSON() ([]byte, error) {
	type wrapper ProcessFileMetadata
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProcessFileMetadata: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProcessFileMetadata: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.processFileMetadata"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProcessFileMetadata: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ProcessFileMetadata{}

func (s *ProcessFileMetadata) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CustomProperties *CustomMetadataDictionary `json:"customProperties,omitempty"`
		OwnerId          nullable.Type[string]     `json:"ownerId,omitempty"`
		CorrelationId    nullable.Type[string]     `json:"correlationId,omitempty"`
		CreatedDateTime  string                    `json:"createdDateTime"`
		Identifier       string                    `json:"identifier"`
		IsTruncated      bool                      `json:"isTruncated"`
		Length           nullable.Type[int64]      `json:"length,omitempty"`
		ModifiedDateTime string                    `json:"modifiedDateTime"`
		Name             string                    `json:"name"`
		ODataId          *string                   `json:"@odata.id,omitempty"`
		ODataType        *string                   `json:"@odata.type,omitempty"`
		SequenceNumber   nullable.Type[int64]      `json:"sequenceNumber,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CustomProperties = decoded.CustomProperties
	s.OwnerId = decoded.OwnerId
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
		return fmt.Errorf("unmarshaling ProcessFileMetadata into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["content"]; ok {
		impl, err := UnmarshalContentBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Content' for 'ProcessFileMetadata': %+v", err)
		}
		s.Content = impl
	}

	return nil
}
