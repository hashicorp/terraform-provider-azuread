package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionHandlerInstance struct {
	// Identifier of the customAccessPackageWorkflowExtension triggered at this instance.
	CustomExtensionId nullable.Type[string] `json:"customExtensionId,omitempty"`

	// The unique run ID for the logic app.
	ExternalCorrelationId nullable.Type[string] `json:"externalCorrelationId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the stage of the request workflow when the access package custom extension runs. The possible values are:
	// assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved,
	// assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
	Stage *AccessPackageCustomExtensionStage `json:"stage,omitempty"`

	// Status of the request to run the access package custom extension workflow that is associated with the logic app. The
	// possible values are: requestSent, requestReceived, unknownFutureValue.
	Status *AccessPackageCustomExtensionHandlerStatus `json:"status,omitempty"`
}
