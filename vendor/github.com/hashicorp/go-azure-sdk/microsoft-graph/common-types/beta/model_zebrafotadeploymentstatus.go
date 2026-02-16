package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentStatus struct {
	// A boolean that indicates if a cancellation was requested on the deployment. NOTE: A cancellation request does not
	// guarantee that the deployment was canceled.
	CancelRequested *bool `json:"cancelRequested,omitempty"`

	// The date and time when this deployment was completed or canceled. The actual date time is determined by the value of
	// state. If the state is canceled, this property holds the cancellation date/time. If the the state is completed, this
	// property holds the completion date/time. If the deployment is not completed before the deployment end date, then
	// completed date/time and end date/time are the same. This is always in the deployment timezone. Note: An installation
	// that is in progress can continue past the deployment end date.
	CompleteOrCanceledDateTime nullable.Type[string] `json:"completeOrCanceledDateTime,omitempty"`

	// An error code indicating the failure reason, when the deployment state is createFailed. Possible values: See
	// zebraFotaErrorCode enum.
	ErrorCode *ZebraFotaErrorCode `json:"errorCode,omitempty"`

	// Date and time when the deployment status was updated from Zebra
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the state of Zebra FOTA deployment.
	State *ZebraFotaDeploymentState `json:"state,omitempty"`

	// An integer that indicates the total number of devices where installation was successful.
	TotalAwaitingInstall *int64 `json:"totalAwaitingInstall,omitempty"`

	// An integer that indicates the total number of devices where installation was canceled.
	TotalCanceled *int64 `json:"totalCanceled,omitempty"`

	// An integer that indicates the total number of devices that have a job in the CREATED state. Typically indicates jobs
	// that did not reach the devices.
	TotalCreated *int64 `json:"totalCreated,omitempty"`

	// An integer that indicates the total number of devices in the deployment.
	TotalDevices *int64 `json:"totalDevices,omitempty"`

	// An integer that indicates the total number of devices where installation was successful.
	TotalDownloading *int64 `json:"totalDownloading,omitempty"`

	// An integer that indicates the total number of devices that have failed to download the new OS file.
	TotalFailedDownload *int64 `json:"totalFailedDownload,omitempty"`

	// An integer that indicates the total number of devices that have failed to install the new OS file.
	TotalFailedInstall *int64 `json:"totalFailedInstall,omitempty"`

	// An integer that indicates the total number of devices that received the json and are scheduled.
	TotalScheduled *int64 `json:"totalScheduled,omitempty"`

	// An integer that indicates the total number of devices where installation was successful.
	TotalSucceededInstall *int64 `json:"totalSucceededInstall,omitempty"`

	// An integer that indicates the total number of devices where no deployment status or end state has not received, even
	// after the scheduled end date was reached.
	TotalUnknown *int64 `json:"totalUnknown,omitempty"`
}
