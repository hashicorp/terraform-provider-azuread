package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCReviewStatus struct {
	// The blob access tier of the Azure Storage account in which the Cloud PC snapshot is saved with. Possible values are
	// hot, cool, cold, and archive, default value is hot.
	AccessTier *CloudPCBlobAccessTier `json:"accessTier,omitempty"`

	// The resource ID of the Azure Storage account in which the Cloud PC snapshot is being saved.
	AzureStorageAccountId nullable.Type[string] `json:"azureStorageAccountId,omitempty"`

	// The name of the Azure Storage account in which the Cloud PC snapshot is being saved.
	AzureStorageAccountName nullable.Type[string] `json:"azureStorageAccountName,omitempty"`

	// The name of the container in an Azure Storage account in which the Cloud PC snapshot is being saved.
	AzureStorageContainerName nullable.Type[string] `json:"azureStorageContainerName,omitempty"`

	// True if the Cloud PC is set to in review by the administrator.
	InReview *bool `json:"inReview,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The specific date and time of the Cloud PC snapshot that was taken and saved automatically, when the Cloud PC is set
	// to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight
	// UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
	RestorePointDateTime nullable.Type[string] `json:"restorePointDateTime,omitempty"`

	// The specific date and time when the Cloud PC was set to in review. The timestamp is shown in ISO 8601 format and
	// Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
	ReviewStartDateTime nullable.Type[string] `json:"reviewStartDateTime,omitempty"`

	// The ID of the Azure subscription in which the Cloud PC snapshot is being saved, in GUID format.
	SubscriptionId nullable.Type[string] `json:"subscriptionId,omitempty"`

	// The name of the Azure subscription in which the Cloud PC snapshot is being saved.
	SubscriptionName nullable.Type[string] `json:"subscriptionName,omitempty"`

	UserAccessLevel *CloudPCUserAccessLevel `json:"userAccessLevel,omitempty"`
}
