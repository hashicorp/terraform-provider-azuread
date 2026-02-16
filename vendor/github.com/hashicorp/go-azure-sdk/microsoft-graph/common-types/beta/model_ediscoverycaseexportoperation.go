package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EdiscoveryCaseOperation = EdiscoveryCaseExportOperation{}

type EdiscoveryCaseExportOperation struct {
	// The name of the Azure storage location where the export will be stored. This only applies to exports stored in your
	// own Azure storage location.
	AzureBlobContainer nullable.Type[string] `json:"azureBlobContainer,omitempty"`

	// The SAS token for the Azure storage location. This only applies to exports stored in your own Azure storage location.
	AzureBlobToken nullable.Type[string] `json:"azureBlobToken,omitempty"`

	// The description provided for the export.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The options provided for the export. For more information, see reviewSet: export. Possible values are: originalFiles,
	// text, pdfReplacement, fileInfo, tags.
	ExportOptions *EdiscoveryExportOptions `json:"exportOptions,omitempty"`

	// The options provided specify the structure of the export. For more information, see reviewSet: export. Possible
	// values are: none, directory, pst.
	ExportStructure *EdiscoveryExportFileStructure `json:"exportStructure,omitempty"`

	// The output folder ID.
	OutputFolderId nullable.Type[string] `json:"outputFolderId,omitempty"`

	// The name provided for the export.
	OutputName nullable.Type[string] `json:"outputName,omitempty"`

	// The review set the content is being exported from.
	ReviewSet *EdiscoveryReviewSet `json:"reviewSet,omitempty"`

	// Fields inherited from EdiscoveryCaseOperation

	// The type of action the operation represents. Possible values are:
	// addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
	Action *EdiscoveryCaseAction `json:"action,omitempty"`

	// The date and time the operation was completed.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The user that created the operation.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time the operation was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The progress of the operation.
	PercentProgress nullable.Type[int64] `json:"percentProgress,omitempty"`

	// Contains success and failure-specific result information.
	ResultInfo *ResultInfo `json:"resultInfo,omitempty"`

	// The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded,
	// partiallySucceeded, failed.
	Status *EdiscoveryCaseOperationStatus `json:"status,omitempty"`

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

func (s EdiscoveryCaseExportOperation) EdiscoveryCaseOperation() BaseEdiscoveryCaseOperationImpl {
	return BaseEdiscoveryCaseOperationImpl{
		Action:            s.Action,
		CompletedDateTime: s.CompletedDateTime,
		CreatedBy:         s.CreatedBy,
		CreatedDateTime:   s.CreatedDateTime,
		PercentProgress:   s.PercentProgress,
		ResultInfo:        s.ResultInfo,
		Status:            s.Status,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s EdiscoveryCaseExportOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdiscoveryCaseExportOperation{}

func (s EdiscoveryCaseExportOperation) MarshalJSON() ([]byte, error) {
	type wrapper EdiscoveryCaseExportOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdiscoveryCaseExportOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryCaseExportOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.caseExportOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdiscoveryCaseExportOperation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EdiscoveryCaseExportOperation{}

func (s *EdiscoveryCaseExportOperation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AzureBlobContainer nullable.Type[string]          `json:"azureBlobContainer,omitempty"`
		AzureBlobToken     nullable.Type[string]          `json:"azureBlobToken,omitempty"`
		Description        nullable.Type[string]          `json:"description,omitempty"`
		ExportOptions      *EdiscoveryExportOptions       `json:"exportOptions,omitempty"`
		ExportStructure    *EdiscoveryExportFileStructure `json:"exportStructure,omitempty"`
		OutputFolderId     nullable.Type[string]          `json:"outputFolderId,omitempty"`
		OutputName         nullable.Type[string]          `json:"outputName,omitempty"`
		ReviewSet          *EdiscoveryReviewSet           `json:"reviewSet,omitempty"`
		Action             *EdiscoveryCaseAction          `json:"action,omitempty"`
		CompletedDateTime  nullable.Type[string]          `json:"completedDateTime,omitempty"`
		CreatedDateTime    nullable.Type[string]          `json:"createdDateTime,omitempty"`
		PercentProgress    nullable.Type[int64]           `json:"percentProgress,omitempty"`
		ResultInfo         *ResultInfo                    `json:"resultInfo,omitempty"`
		Status             *EdiscoveryCaseOperationStatus `json:"status,omitempty"`
		Id                 *string                        `json:"id,omitempty"`
		ODataId            *string                        `json:"@odata.id,omitempty"`
		ODataType          *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AzureBlobContainer = decoded.AzureBlobContainer
	s.AzureBlobToken = decoded.AzureBlobToken
	s.Description = decoded.Description
	s.ExportOptions = decoded.ExportOptions
	s.ExportStructure = decoded.ExportStructure
	s.OutputFolderId = decoded.OutputFolderId
	s.OutputName = decoded.OutputName
	s.ReviewSet = decoded.ReviewSet
	s.Action = decoded.Action
	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PercentProgress = decoded.PercentProgress
	s.ResultInfo = decoded.ResultInfo
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EdiscoveryCaseExportOperation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EdiscoveryCaseExportOperation': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
