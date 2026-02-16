package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestStageDetail struct {
	// Describes the error, if any, for the current stage.
	Error *PublicError `json:"error,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The stage of the subject rights request. Possible values are: contentRetrieval, contentReview, generateReport,
	// contentDeletion, caseResolved, unknownFutureValue, approval. Use the Prefer: include-unknown-enum-members request
	// header to get the following value in this evolvable enum: approval.
	Stage *SubjectRightsRequestStage `json:"stage,omitempty"`

	// Status of the current stage. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
	Status *SubjectRightsRequestStageStatus `json:"status,omitempty"`
}
