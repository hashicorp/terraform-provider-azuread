package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityCaseOperation = SecurityEdiscoverySearchExportOperation{}

type SecurityEdiscoverySearchExportOperation struct {
	// The additional items to include in the export. The possible values are: none, teamsAndYammerConversations,
	// cloudAttachments, allDocumentVersions, subfolderContents, listAttachments, unknownFutureValue, htmlTranscripts,
	// advancedIndexing, allItemsInFolder, includeFolderAndPath, condensePaths, friendlyName, splitSource,
	// optimizedPartitionSize, includeReport. Use the Prefer: include-unknown-enum-members request header to get the
	// following values from this evolvable enum: htmlTranscripts, advancedIndexing, allItemsInFolder, includeFolderAndPath,
	// condensePaths, friendlyName, splitSource, optimizedPartitionSize, includeReport.
	AdditionalOptions *SecurityAdditionalOptions `json:"additionalOptions,omitempty"`

	// The versions of cloud attachments to include in messages. Possible values are: latest, recent10, recent100, all,
	// unknownFutureValue.
	CloudAttachmentVersion *SecurityCloudAttachmentVersion `json:"cloudAttachmentVersion,omitempty"`

	// The description of the export by the user.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of export provided by the user.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The versions of files in SharePoint to include. Possible values are: latest, recent10, recent100, all,
	// unknownFutureValue.
	DocumentVersion *SecurityDocumentVersion `json:"documentVersion,omitempty"`

	// Items to be included in the export. The possible values are: searchHits, partiallyIndexed, unknownFutureValue.
	ExportCriteria *SecurityExportCriteria `json:"exportCriteria,omitempty"`

	// Contains the properties for an export file metadata, including downloadUrl, fileName, and size.
	ExportFileMetadata *[]SecurityExportFileMetadata `json:"exportFileMetadata,omitempty"`

	// Format of the emails of the export. The possible values are: pst, msg, eml, unknownFutureValue.
	ExportFormat *SecurityExportFormat `json:"exportFormat,omitempty"`

	// Location scope for partially indexed items. You can choose to include partially indexed items only in responsive
	// locations with search hits or in all targeted locations. The possible values are: responsiveLocations,
	// nonresponsiveLocations, unknownFutureValue.
	ExportLocation *SecurityExportLocation `json:"exportLocation,omitempty"`

	// Indicates whether to export single items.
	ExportSingleItems nullable.Type[bool] `json:"exportSingleItems,omitempty"`

	// The eDiscovery searches under each case.
	Search *SecurityEdiscoverySearch `json:"search,omitempty"`

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

func (s SecurityEdiscoverySearchExportOperation) SecurityCaseOperation() BaseSecurityCaseOperationImpl {
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

func (s SecurityEdiscoverySearchExportOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoverySearchExportOperation{}

func (s SecurityEdiscoverySearchExportOperation) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoverySearchExportOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoverySearchExportOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoverySearchExportOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoverySearchExportOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoverySearchExportOperation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityEdiscoverySearchExportOperation{}

func (s *SecurityEdiscoverySearchExportOperation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AdditionalOptions      *SecurityAdditionalOptions      `json:"additionalOptions,omitempty"`
		CloudAttachmentVersion *SecurityCloudAttachmentVersion `json:"cloudAttachmentVersion,omitempty"`
		Description            nullable.Type[string]           `json:"description,omitempty"`
		DisplayName            nullable.Type[string]           `json:"displayName,omitempty"`
		DocumentVersion        *SecurityDocumentVersion        `json:"documentVersion,omitempty"`
		ExportCriteria         *SecurityExportCriteria         `json:"exportCriteria,omitempty"`
		ExportFileMetadata     *[]SecurityExportFileMetadata   `json:"exportFileMetadata,omitempty"`
		ExportFormat           *SecurityExportFormat           `json:"exportFormat,omitempty"`
		ExportLocation         *SecurityExportLocation         `json:"exportLocation,omitempty"`
		ExportSingleItems      nullable.Type[bool]             `json:"exportSingleItems,omitempty"`
		Search                 *SecurityEdiscoverySearch       `json:"search,omitempty"`
		Action                 *SecurityCaseAction             `json:"action,omitempty"`
		CompletedDateTime      nullable.Type[string]           `json:"completedDateTime,omitempty"`
		CreatedDateTime        nullable.Type[string]           `json:"createdDateTime,omitempty"`
		PercentProgress        nullable.Type[int64]            `json:"percentProgress,omitempty"`
		ResultInfo             *ResultInfo                     `json:"resultInfo,omitempty"`
		Status                 *SecurityCaseOperationStatus    `json:"status,omitempty"`
		Id                     *string                         `json:"id,omitempty"`
		ODataId                *string                         `json:"@odata.id,omitempty"`
		ODataType              *string                         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AdditionalOptions = decoded.AdditionalOptions
	s.CloudAttachmentVersion = decoded.CloudAttachmentVersion
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.DocumentVersion = decoded.DocumentVersion
	s.ExportCriteria = decoded.ExportCriteria
	s.ExportFileMetadata = decoded.ExportFileMetadata
	s.ExportFormat = decoded.ExportFormat
	s.ExportLocation = decoded.ExportLocation
	s.ExportSingleItems = decoded.ExportSingleItems
	s.Search = decoded.Search
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
		return fmt.Errorf("unmarshaling SecurityEdiscoverySearchExportOperation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityEdiscoverySearchExportOperation': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
