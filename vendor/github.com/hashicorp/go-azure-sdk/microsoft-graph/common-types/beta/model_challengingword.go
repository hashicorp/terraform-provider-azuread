package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChallengingWord struct {
	// Number of times the word was found challenging by the student during the reading session.
	Count *int64 `json:"count,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The specific word that the student found challenging during the reading session.
	Word *string `json:"word,omitempty"`
}
