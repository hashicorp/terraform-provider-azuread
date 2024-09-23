package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResponsibleSensitiveType struct {
	Description nullable.Type[string] `json:"description,omitempty"`
	Id          nullable.Type[string] `json:"id,omitempty"`
	Name        nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PublisherName   nullable.Type[string] `json:"publisherName,omitempty"`
	RulePackageId   nullable.Type[string] `json:"rulePackageId,omitempty"`
	RulePackageType nullable.Type[string] `json:"rulePackageType,omitempty"`
}
