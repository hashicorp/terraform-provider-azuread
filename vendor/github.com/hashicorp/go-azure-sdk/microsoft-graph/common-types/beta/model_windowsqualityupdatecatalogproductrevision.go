package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateCatalogProductRevision struct {
	// The display name of the windows quality update catalog product revision. For example, 'Windows 11, version 22H2,
	// build 22621.4112'. Read-only
	DisplayName *string `json:"displayName,omitempty"`

	// Represents a knowledge base (KB) article.
	KnowledgeBaseArticle *WindowsQualityUpdateProductKnowledgeBaseArticle `json:"knowledgeBaseArticle,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the build version details of a product revision that is associated with a quality update.
	OsBuild *WindowsQualityUpdateProductBuildVersionDetail `json:"osBuild,omitempty"`

	// The product name of the windows quality update catalog product revision. For example, 'Windows 11'. Read-only
	ProductName *string `json:"productName,omitempty"`

	// The date and time when the windows quality update catalog product revision was released. The Timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC time. Read-only
	ReleaseDateTime *string `json:"releaseDateTime,omitempty"`

	// The version name of the windows quality update catalog product revision. For example, '22H2'. Read-only
	VersionName *string `json:"versionName,omitempty"`
}
