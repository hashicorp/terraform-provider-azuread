package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityCaseOperation = SecurityEdiscoveryEstimateOperation{}

type SecurityEdiscoveryEstimateOperation struct {
	// The estimated count of items for the search that matched the content query.
	IndexedItemCount nullable.Type[int64] `json:"indexedItemCount,omitempty"`

	// The estimated size of items for the search that matched the content query.
	IndexedItemsSize nullable.Type[int64] `json:"indexedItemsSize,omitempty"`

	// The number of mailboxes that had search hits.
	MailboxCount nullable.Type[int64] `json:"mailboxCount,omitempty"`

	// eDiscovery search.
	Search *SecurityEdiscoverySearch `json:"search,omitempty"`

	// The number of mailboxes that had search hits.
	SiteCount nullable.Type[int64] `json:"siteCount,omitempty"`

	// The estimated count of unindexed items for the collection.
	UnindexedItemCount nullable.Type[int64] `json:"unindexedItemCount,omitempty"`

	// The estimated size of unindexed items for the collection.
	UnindexedItemsSize nullable.Type[int64] `json:"unindexedItemsSize,omitempty"`

	// Fields inherited from SecurityCaseOperation

	// The type of action the operation represents. Possible values are:
	// addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
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

func (s SecurityEdiscoveryEstimateOperation) SecurityCaseOperation() BaseSecurityCaseOperationImpl {
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

func (s SecurityEdiscoveryEstimateOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryEstimateOperation{}

func (s SecurityEdiscoveryEstimateOperation) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryEstimateOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryEstimateOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryEstimateOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryEstimateOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryEstimateOperation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityEdiscoveryEstimateOperation{}

func (s *SecurityEdiscoveryEstimateOperation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IndexedItemCount   nullable.Type[int64]         `json:"indexedItemCount,omitempty"`
		IndexedItemsSize   nullable.Type[int64]         `json:"indexedItemsSize,omitempty"`
		MailboxCount       nullable.Type[int64]         `json:"mailboxCount,omitempty"`
		Search             *SecurityEdiscoverySearch    `json:"search,omitempty"`
		SiteCount          nullable.Type[int64]         `json:"siteCount,omitempty"`
		UnindexedItemCount nullable.Type[int64]         `json:"unindexedItemCount,omitempty"`
		UnindexedItemsSize nullable.Type[int64]         `json:"unindexedItemsSize,omitempty"`
		Action             *SecurityCaseAction          `json:"action,omitempty"`
		CompletedDateTime  nullable.Type[string]        `json:"completedDateTime,omitempty"`
		CreatedDateTime    nullable.Type[string]        `json:"createdDateTime,omitempty"`
		PercentProgress    nullable.Type[int64]         `json:"percentProgress,omitempty"`
		ResultInfo         *ResultInfo                  `json:"resultInfo,omitempty"`
		Status             *SecurityCaseOperationStatus `json:"status,omitempty"`
		Id                 *string                      `json:"id,omitempty"`
		ODataId            *string                      `json:"@odata.id,omitempty"`
		ODataType          *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IndexedItemCount = decoded.IndexedItemCount
	s.IndexedItemsSize = decoded.IndexedItemsSize
	s.MailboxCount = decoded.MailboxCount
	s.Search = decoded.Search
	s.SiteCount = decoded.SiteCount
	s.UnindexedItemCount = decoded.UnindexedItemCount
	s.UnindexedItemsSize = decoded.UnindexedItemsSize
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
		return fmt.Errorf("unmarshaling SecurityEdiscoveryEstimateOperation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityEdiscoveryEstimateOperation': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
