package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventQuery struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents unique identification for the query. 'Asset ID' for SharePoint Online and OneDrive for Business,
	// 'keywords' for Exchange Online.
	Query *string `json:"query,omitempty"`

	// Represents the type of query associated with an event. 'files' for SPO and ODB and 'messages' for EXO.The possible
	// values are: files, messages, unknownFutureValue.
	QueryType *SecurityQueryType `json:"queryType,omitempty"`
}
