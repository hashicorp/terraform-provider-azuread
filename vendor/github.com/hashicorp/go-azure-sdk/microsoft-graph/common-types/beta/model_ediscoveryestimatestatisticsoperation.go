package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EdiscoveryCaseOperation = EdiscoveryEstimateStatisticsOperation{}

type EdiscoveryEstimateStatisticsOperation struct {
	// The estimated count of items for the sourceCollection that matched the content query.
	IndexedItemCount nullable.Type[int64] `json:"indexedItemCount,omitempty"`

	// The estimated size of items for the sourceCollection that matched the content query.
	IndexedItemsSize nullable.Type[int64] `json:"indexedItemsSize,omitempty"`

	// The number of mailboxes that had search hits.
	MailboxCount nullable.Type[int64] `json:"mailboxCount,omitempty"`

	// The number of mailboxes that had search hits.
	SiteCount nullable.Type[int64] `json:"siteCount,omitempty"`

	// eDiscovery collection, commonly known as a search.
	SourceCollection *EdiscoverySourceCollection `json:"sourceCollection,omitempty"`

	// The estimated count of unindexed items for the collection.
	UnindexedItemCount nullable.Type[int64] `json:"unindexedItemCount,omitempty"`

	// The estimated size of unindexed items for the collection.
	UnindexedItemsSize nullable.Type[int64] `json:"unindexedItemsSize,omitempty"`

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

func (s EdiscoveryEstimateStatisticsOperation) EdiscoveryCaseOperation() BaseEdiscoveryCaseOperationImpl {
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

func (s EdiscoveryEstimateStatisticsOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdiscoveryEstimateStatisticsOperation{}

func (s EdiscoveryEstimateStatisticsOperation) MarshalJSON() ([]byte, error) {
	type wrapper EdiscoveryEstimateStatisticsOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdiscoveryEstimateStatisticsOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryEstimateStatisticsOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.estimateStatisticsOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdiscoveryEstimateStatisticsOperation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EdiscoveryEstimateStatisticsOperation{}

func (s *EdiscoveryEstimateStatisticsOperation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IndexedItemCount   nullable.Type[int64]           `json:"indexedItemCount,omitempty"`
		IndexedItemsSize   nullable.Type[int64]           `json:"indexedItemsSize,omitempty"`
		MailboxCount       nullable.Type[int64]           `json:"mailboxCount,omitempty"`
		SiteCount          nullable.Type[int64]           `json:"siteCount,omitempty"`
		SourceCollection   *EdiscoverySourceCollection    `json:"sourceCollection,omitempty"`
		UnindexedItemCount nullable.Type[int64]           `json:"unindexedItemCount,omitempty"`
		UnindexedItemsSize nullable.Type[int64]           `json:"unindexedItemsSize,omitempty"`
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

	s.IndexedItemCount = decoded.IndexedItemCount
	s.IndexedItemsSize = decoded.IndexedItemsSize
	s.MailboxCount = decoded.MailboxCount
	s.SiteCount = decoded.SiteCount
	s.SourceCollection = decoded.SourceCollection
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
		return fmt.Errorf("unmarshaling EdiscoveryEstimateStatisticsOperation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EdiscoveryEstimateStatisticsOperation': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
