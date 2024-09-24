package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ApplicationTemplate{}

type ApplicationTemplate struct {
	// The list of categories for the application. Supported values can be: Collaboration, Business Management, Consumer,
	// Content management, CRM, Data services, Developer services, E-commerce, Education, ERP, Finance, Health, Human
	// resources, IT infrastructure, Mail, Management, Marketing, Media, Productivity, Project management,
	// Telecommunications, Tools, Travel, and Web design & hosting.
	Categories *[]string `json:"categories,omitempty"`

	// The URIs required for the single sign-on configuration of a preintegrated application.
	ConfigurationUris *[]ConfigurationUri `json:"configurationUris,omitempty"`

	// A description of the application.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of the application.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The home page URL of the application.
	HomePageUrl nullable.Type[string] `json:"homePageUrl,omitempty"`

	InformationalUrls *InformationalUrls `json:"informationalUrls,omitempty"`

	// The URL to get the logo for this application.
	LogoUrl nullable.Type[string] `json:"logoUrl,omitempty"`

	// The name of the publisher for this application.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	SupportedClaimConfiguration *SupportedClaimConfiguration `json:"supportedClaimConfiguration,omitempty"`

	// The list of provisioning modes supported by this application. The only valid value is sync.
	SupportedProvisioningTypes *[]string `json:"supportedProvisioningTypes,omitempty"`

	// The list of single sign-on modes supported by this application. The supported values are oidc, password, saml, and
	// notSupported.
	SupportedSingleSignOnModes *[]string `json:"supportedSingleSignOnModes,omitempty"`

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

func (s ApplicationTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ApplicationTemplate{}

func (s ApplicationTemplate) MarshalJSON() ([]byte, error) {
	type wrapper ApplicationTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApplicationTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApplicationTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.applicationTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApplicationTemplate: %+v", err)
	}

	return encoded, nil
}
