package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCDeviceImage{}

type CloudPCDeviceImage struct {
	// The display name of the associated device image. The device image display name and the version are used to uniquely
	// identify the Cloud PC device image. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The error code of the status of the image that indicates why the upload failed, if applicable. Possible values are:
	// internalServerError, sourceImageNotFound, osVersionNotSupported, sourceImageInvalid, sourceImageNotGeneralized,
	// unknownFutureValue, vmAlreadyAzureAdJoined, paidSourceImageNotSupport, sourceImageNotSupportCustomizeVMName,
	// sourceImageSizeExceedsLimitation. Note that you must use the Prefer: include-unknown-enum-members request header to
	// get the following values from this evolvable enum: vmAlreadyAzureAdJoined, paidSourceImageNotSupport,
	// sourceImageNotSupportCustomizeVMName, sourceImageSizeExceedsLimitation. Read-only.
	ErrorCode *CloudPCDeviceImageErrorCode `json:"errorCode,omitempty"`

	// The date when the image became unavailable. Read-only.
	ExpirationDate nullable.Type[string] `json:"expirationDate,omitempty"`

	// The data and time when the image was last modified. The timestamp represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The operating system of the image. For example, Windows 10 Enterprise. Read-only.
	OperatingSystem nullable.Type[string] `json:"operatingSystem,omitempty"`

	// The OS build version of the image. For example, 1909. Read-only.
	OsBuildNumber nullable.Type[string] `json:"osBuildNumber,omitempty"`

	// The OS status of this image. Possible values are: supported, supportedWithWarning, unknown, unknownFutureValue. The
	// default value is unknown. Read-only.
	OsStatus *CloudPCDeviceImageOsStatus `json:"osStatus,omitempty"`

	ScopeIds *[]string `json:"scopeIds,omitempty"`

	// The unique identifier (ID) of the source image resource on Azure. The required ID format is:
	// '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}'.
	// Read-only.
	SourceImageResourceId nullable.Type[string] `json:"sourceImageResourceId,omitempty"`

	// The status of the image on the Cloud PC. Possible values are: pending, ready, failed, unknownFutureValue. Read-only.
	Status *CloudPCDeviceImageStatus `json:"status,omitempty"`

	// The details of the status of the image that indicates why the upload failed, if applicable. Possible values are:
	// internalServerError, sourceImageNotFound, osVersionNotSupported, sourceImageInvalid, sourceImageNotGeneralized,
	// unknownFutureValue, vmAlreadyAzureAdJoined, paidSourceImageNotSupport, sourceImageNotSupportCustomizeVMName,
	// sourceImageSizeExceedsLimitation. Note that you must use the Prefer: include-unknown-enum-members request header to
	// get the following values from this evolvable enum: vmAlreadyAzureAdJoined, paidSourceImageNotSupport,
	// sourceImageNotSupportCustomizeVMName, sourceImageSizeExceedsLimitation. Read-only. The statusDetails property is
	// deprecated and will stop returning data on January 31, 2024. Going forward, use the errorCode property.
	StatusDetails *CloudPCDeviceImageStatusDetails `json:"statusDetails,omitempty"`

	// The image version. For example, 0.0.1 and 1.5.13. Read-only.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s CloudPCDeviceImage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCDeviceImage{}

func (s CloudPCDeviceImage) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCDeviceImage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCDeviceImage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCDeviceImage: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "errorCode")
	delete(decoded, "expirationDate")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "operatingSystem")
	delete(decoded, "osBuildNumber")
	delete(decoded, "osStatus")
	delete(decoded, "sourceImageResourceId")
	delete(decoded, "status")
	delete(decoded, "statusDetails")
	delete(decoded, "version")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcDeviceImage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCDeviceImage: %+v", err)
	}

	return encoded, nil
}
