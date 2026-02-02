package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DocumentProcessingJob{}

type DocumentProcessingJob struct {
	// Date and time of item creation. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The document processing job type. The possible values are: file, folder
	JobType *DocumentProcessingJobType `json:"jobType,omitempty"`

	// The listItemUniqueId of the file, or folder to process. Use GET driveItem resource operation and read sharepointIds
	// property to get listItemUniqueId.
	ListItemUniqueId *string `json:"listItemUniqueId,omitempty"`

	// The document processing Job status. The possible values are: inProgress, completed, failed, unknownFutureValue.
	Status *DocumentProcessingJobStatus `json:"status,omitempty"`

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

func (s DocumentProcessingJob) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DocumentProcessingJob{}

func (s DocumentProcessingJob) MarshalJSON() ([]byte, error) {
	type wrapper DocumentProcessingJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DocumentProcessingJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DocumentProcessingJob: %+v", err)
	}

	delete(decoded, "createdDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.documentProcessingJob"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DocumentProcessingJob: %+v", err)
	}

	return encoded, nil
}
