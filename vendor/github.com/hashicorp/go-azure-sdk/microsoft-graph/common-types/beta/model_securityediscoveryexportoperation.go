package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityCaseOperation = SecurityEdiscoveryExportOperation{}

type SecurityEdiscoveryExportOperation struct {
	// The name of the Azure storage location where the export is stored. This name only applies to exports stored in your
	// own Azure storage location. The azureBlobContainer property is deprecated and stopped returning data on April 30,
	// 2023.
	AzureBlobContainer nullable.Type[string] `json:"azureBlobContainer,omitempty"`

	// The SAS token for the Azure storage location. This token only applies to exports stored in your own Azure storage
	// location. The azureBlobToken property is deprecated and stopped returning data on April 30, 2023.
	AzureBlobToken nullable.Type[string] `json:"azureBlobToken,omitempty"`

	// The description provided for the export.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Contains the properties for an export file metadata, including downloadUrl, fileName, and size. If you export to an
	// Azure storage location, this property returns empty.
	ExportFileMetadata *[]SecurityExportFileMetadata `json:"exportFileMetadata,omitempty"`

	// The options provided for the export. For more information, see reviewSet: export. The fileInfo member is deprecated
	// and stopped returning data on April 30, 2023. Going forward, the summary and load files are always included. Possible
	// values are: originalFiles, text, pdfReplacement, tags, unknownFutureValue, splitSource, includeFolderAndPath,
	// friendlyName, condensePaths, optimizedPartitionSize. Use the Prefer: include-unknown-enum-members request header to
	// get the following values from this evolvable enum: splitSource, includeFolderAndPath, friendlyName, condensePaths,
	// optimizedPartitionSize.
	ExportOptions *SecurityExportOptions `json:"exportOptions,omitempty"`

	// The options that specify the structure of the export. For more information, see reviewSet: export. Possible values
	// are: none, directory, pst, unknownFutureValue, msg. Use the Prefer: include-unknown-enum-members request header to
	// get the following value from this evolvable enum: msg.
	ExportStructure *SecurityExportFileStructure `json:"exportStructure,omitempty"`

	// The output folder ID. The outputFolderId property is deprecated and stopped returning data on April 30, 2023.
	OutputFolderId nullable.Type[string] `json:"outputFolderId,omitempty"`

	// The name provided for the export.
	OutputName nullable.Type[string] `json:"outputName,omitempty"`

	// Review set from where documents are exported.
	ReviewSet *SecurityEdiscoveryReviewSet `json:"reviewSet,omitempty"`

	// The review set query that is used to filter the documents for export.
	ReviewSetQuery *SecurityEdiscoveryReviewSetQuery `json:"reviewSetQuery,omitempty"`

	// Fields inherited from SecurityCaseOperation

	// The type of action the operation represents. Possible values are: contentExport, applyTags, convertToPdf, index,
	// estimateStatistics, addToReviewSet, holdUpdate, unknownFutureValue, purgeData, exportReport, exportResult. Use the
	// Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: purgeData,
	// exportReport, exportResult.
	Action *SecurityCaseAction `json:"action,omitempty"`

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
	Status *SecurityCaseOperationStatus `json:"status,omitempty"`

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

func (s SecurityEdiscoveryExportOperation) SecurityCaseOperation() BaseSecurityCaseOperationImpl {
	return BaseSecurityCaseOperationImpl{
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

func (s SecurityEdiscoveryExportOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryExportOperation{}

func (s SecurityEdiscoveryExportOperation) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryExportOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryExportOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryExportOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryExportOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryExportOperation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityEdiscoveryExportOperation{}

func (s *SecurityEdiscoveryExportOperation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AzureBlobContainer nullable.Type[string]             `json:"azureBlobContainer,omitempty"`
		AzureBlobToken     nullable.Type[string]             `json:"azureBlobToken,omitempty"`
		Description        nullable.Type[string]             `json:"description,omitempty"`
		ExportFileMetadata *[]SecurityExportFileMetadata     `json:"exportFileMetadata,omitempty"`
		ExportOptions      *SecurityExportOptions            `json:"exportOptions,omitempty"`
		ExportStructure    *SecurityExportFileStructure      `json:"exportStructure,omitempty"`
		OutputFolderId     nullable.Type[string]             `json:"outputFolderId,omitempty"`
		OutputName         nullable.Type[string]             `json:"outputName,omitempty"`
		ReviewSet          *SecurityEdiscoveryReviewSet      `json:"reviewSet,omitempty"`
		ReviewSetQuery     *SecurityEdiscoveryReviewSetQuery `json:"reviewSetQuery,omitempty"`
		Action             *SecurityCaseAction               `json:"action,omitempty"`
		CompletedDateTime  nullable.Type[string]             `json:"completedDateTime,omitempty"`
		CreatedDateTime    nullable.Type[string]             `json:"createdDateTime,omitempty"`
		PercentProgress    nullable.Type[int64]              `json:"percentProgress,omitempty"`
		ResultInfo         *ResultInfo                       `json:"resultInfo,omitempty"`
		Status             *SecurityCaseOperationStatus      `json:"status,omitempty"`
		Id                 *string                           `json:"id,omitempty"`
		ODataId            *string                           `json:"@odata.id,omitempty"`
		ODataType          *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AzureBlobContainer = decoded.AzureBlobContainer
	s.AzureBlobToken = decoded.AzureBlobToken
	s.Description = decoded.Description
	s.ExportFileMetadata = decoded.ExportFileMetadata
	s.ExportOptions = decoded.ExportOptions
	s.ExportStructure = decoded.ExportStructure
	s.OutputFolderId = decoded.OutputFolderId
	s.OutputName = decoded.OutputName
	s.ReviewSet = decoded.ReviewSet
	s.ReviewSetQuery = decoded.ReviewSetQuery
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
		return fmt.Errorf("unmarshaling SecurityEdiscoveryExportOperation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityEdiscoveryExportOperation': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
