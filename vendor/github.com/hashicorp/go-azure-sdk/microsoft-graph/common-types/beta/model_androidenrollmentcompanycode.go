package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnrollmentCompanyCode struct {
	// Enrollment Token used by the User to enroll their device.
	EnrollmentToken nullable.Type[string] `json:"enrollmentToken,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// String used to generate a QR code for the token.
	QrCodeContent nullable.Type[string] `json:"qrCodeContent,omitempty"`

	// Generated QR code for the token.
	QrCodeImage *MimeContent `json:"qrCodeImage,omitempty"`
}
