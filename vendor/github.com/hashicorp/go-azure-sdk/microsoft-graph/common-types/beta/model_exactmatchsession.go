package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExactMatchSessionBase = ExactMatchSession{}

type ExactMatchSession struct {
	Checksum      nullable.Type[string]  `json:"checksum,omitempty"`
	DataUploadURI nullable.Type[string]  `json:"dataUploadURI,omitempty"`
	Fields        *[]string              `json:"fields,omitempty"`
	FileName      nullable.Type[string]  `json:"fileName,omitempty"`
	RowsPerBlock  nullable.Type[int64]   `json:"rowsPerBlock,omitempty"`
	Salt          nullable.Type[string]  `json:"salt,omitempty"`
	UploadAgent   *ExactMatchUploadAgent `json:"uploadAgent,omitempty"`
	UploadAgentId nullable.Type[string]  `json:"uploadAgentId,omitempty"`

	// Fields inherited from ExactMatchSessionBase

	DataStoreId                  nullable.Type[string] `json:"dataStoreId,omitempty"`
	ProcessingCompletionDateTime nullable.Type[string] `json:"processingCompletionDateTime,omitempty"`
	RemainingBlockCount          nullable.Type[int64]  `json:"remainingBlockCount,omitempty"`
	RemainingJobCount            nullable.Type[int64]  `json:"remainingJobCount,omitempty"`
	State                        nullable.Type[string] `json:"state,omitempty"`
	TotalBlockCount              nullable.Type[int64]  `json:"totalBlockCount,omitempty"`
	TotalJobCount                nullable.Type[int64]  `json:"totalJobCount,omitempty"`
	UploadCompletionDateTime     nullable.Type[string] `json:"uploadCompletionDateTime,omitempty"`

	// Fields inherited from ExactMatchJobBase

	CompletionDateTime  nullable.Type[string] `json:"completionDateTime,omitempty"`
	CreationDateTime    nullable.Type[string] `json:"creationDateTime,omitempty"`
	Error               *ClassificationError  `json:"error,omitempty"`
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`
	StartDateTime       nullable.Type[string] `json:"startDateTime,omitempty"`

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

func (s ExactMatchSession) ExactMatchSessionBase() BaseExactMatchSessionBaseImpl {
	return BaseExactMatchSessionBaseImpl{
		DataStoreId:                  s.DataStoreId,
		ProcessingCompletionDateTime: s.ProcessingCompletionDateTime,
		RemainingBlockCount:          s.RemainingBlockCount,
		RemainingJobCount:            s.RemainingJobCount,
		State:                        s.State,
		TotalBlockCount:              s.TotalBlockCount,
		TotalJobCount:                s.TotalJobCount,
		UploadCompletionDateTime:     s.UploadCompletionDateTime,
		CompletionDateTime:           s.CompletionDateTime,
		CreationDateTime:             s.CreationDateTime,
		Error:                        s.Error,
		LastUpdatedDateTime:          s.LastUpdatedDateTime,
		StartDateTime:                s.StartDateTime,
		Id:                           s.Id,
		ODataId:                      s.ODataId,
		ODataType:                    s.ODataType,
	}
}

func (s ExactMatchSession) ExactMatchJobBase() BaseExactMatchJobBaseImpl {
	return BaseExactMatchJobBaseImpl{
		CompletionDateTime:  s.CompletionDateTime,
		CreationDateTime:    s.CreationDateTime,
		Error:               s.Error,
		LastUpdatedDateTime: s.LastUpdatedDateTime,
		StartDateTime:       s.StartDateTime,
		Id:                  s.Id,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

func (s ExactMatchSession) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExactMatchSession{}

func (s ExactMatchSession) MarshalJSON() ([]byte, error) {
	type wrapper ExactMatchSession
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExactMatchSession: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExactMatchSession: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.exactMatchSession"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExactMatchSession: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ExactMatchSession{}

func (s *ExactMatchSession) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Checksum                     nullable.Type[string]  `json:"checksum,omitempty"`
		DataUploadURI                nullable.Type[string]  `json:"dataUploadURI,omitempty"`
		Fields                       *[]string              `json:"fields,omitempty"`
		FileName                     nullable.Type[string]  `json:"fileName,omitempty"`
		RowsPerBlock                 nullable.Type[int64]   `json:"rowsPerBlock,omitempty"`
		Salt                         nullable.Type[string]  `json:"salt,omitempty"`
		UploadAgent                  *ExactMatchUploadAgent `json:"uploadAgent,omitempty"`
		UploadAgentId                nullable.Type[string]  `json:"uploadAgentId,omitempty"`
		DataStoreId                  nullable.Type[string]  `json:"dataStoreId,omitempty"`
		ProcessingCompletionDateTime nullable.Type[string]  `json:"processingCompletionDateTime,omitempty"`
		RemainingBlockCount          nullable.Type[int64]   `json:"remainingBlockCount,omitempty"`
		RemainingJobCount            nullable.Type[int64]   `json:"remainingJobCount,omitempty"`
		State                        nullable.Type[string]  `json:"state,omitempty"`
		TotalBlockCount              nullable.Type[int64]   `json:"totalBlockCount,omitempty"`
		TotalJobCount                nullable.Type[int64]   `json:"totalJobCount,omitempty"`
		UploadCompletionDateTime     nullable.Type[string]  `json:"uploadCompletionDateTime,omitempty"`
		CompletionDateTime           nullable.Type[string]  `json:"completionDateTime,omitempty"`
		CreationDateTime             nullable.Type[string]  `json:"creationDateTime,omitempty"`
		LastUpdatedDateTime          nullable.Type[string]  `json:"lastUpdatedDateTime,omitempty"`
		StartDateTime                nullable.Type[string]  `json:"startDateTime,omitempty"`
		Id                           *string                `json:"id,omitempty"`
		ODataId                      *string                `json:"@odata.id,omitempty"`
		ODataType                    *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Checksum = decoded.Checksum
	s.DataUploadURI = decoded.DataUploadURI
	s.Fields = decoded.Fields
	s.FileName = decoded.FileName
	s.RowsPerBlock = decoded.RowsPerBlock
	s.Salt = decoded.Salt
	s.UploadAgent = decoded.UploadAgent
	s.UploadAgentId = decoded.UploadAgentId
	s.CompletionDateTime = decoded.CompletionDateTime
	s.CreationDateTime = decoded.CreationDateTime
	s.DataStoreId = decoded.DataStoreId
	s.Id = decoded.Id
	s.LastUpdatedDateTime = decoded.LastUpdatedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ProcessingCompletionDateTime = decoded.ProcessingCompletionDateTime
	s.RemainingBlockCount = decoded.RemainingBlockCount
	s.RemainingJobCount = decoded.RemainingJobCount
	s.StartDateTime = decoded.StartDateTime
	s.State = decoded.State
	s.TotalBlockCount = decoded.TotalBlockCount
	s.TotalJobCount = decoded.TotalJobCount
	s.UploadCompletionDateTime = decoded.UploadCompletionDateTime

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ExactMatchSession into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["error"]; ok {
		impl, err := UnmarshalClassificationErrorImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Error' for 'ExactMatchSession': %+v", err)
		}
		s.Error = &impl
	}

	return nil
}
