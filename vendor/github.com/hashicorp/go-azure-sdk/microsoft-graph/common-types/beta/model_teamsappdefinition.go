package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamsAppDefinition{}

type TeamsAppDefinition struct {
	// A collection of scopes where the Teams app can be installed. Possible values are:team—Indicates that the Teams app
	// can be installed within a team and is authorized to access that team's data. groupChat—Indicates that the Teams app
	// can be installed within a group chat and is authorized to access that group chat's data. personal—Indicates that
	// the Teams app can be installed in the personal scope of a user and is authorized to access that user's data.
	AllowedInstallationScopes *TeamsAppInstallationScopes `json:"allowedInstallationScopes,omitempty"`

	// Authorization requirements specified in the Teams app manifest.
	Authorization *TeamsAppAuthorization `json:"authorization,omitempty"`

	// The WebApplicationInfo.Id from the Teams app manifest.
	AzureADAppId nullable.Type[string] `json:"azureADAppId,omitempty"`

	// The details of the bot specified in the Teams app manifest.
	Bot *TeamworkBot `json:"bot,omitempty"`

	// The color version of the Teams app's icon.
	ColorIcon *TeamsAppIcon `json:"colorIcon,omitempty"`

	CreatedBy IdentitySet `json:"createdBy"`

	// Dashboard cards specified in the Teams app manifest.
	DashboardCards *[]TeamsAppDashboardCardDefinition `json:"dashboardCards,omitempty"`

	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of the app provided by the app developer.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The outline version of the Teams app's icon.
	OutlineIcon *TeamsAppIcon `json:"outlineIcon,omitempty"`

	// The published status of a specific version of a Teams app. Possible values are:submitted—The specific version of
	// the Teams app has been submitted and is under review. published - The request to publish the specific version of the
	// Teams app has been approved by the admin and the app is published. rejected - The request to publish the specific
	// version of the Teams app was rejected by the admin.
	PublishingState *TeamsAppPublishingState `json:"publishingState,omitempty"`

	Shortdescription nullable.Type[string] `json:"shortdescription,omitempty"`

	// The ID from the Teams app manifest.
	TeamsAppId nullable.Type[string] `json:"teamsAppId,omitempty"`

	// The version number of the application.
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

func (s TeamsAppDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsAppDefinition{}

func (s TeamsAppDefinition) MarshalJSON() ([]byte, error) {
	type wrapper TeamsAppDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsAppDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAppDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamsAppDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsAppDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamsAppDefinition{}

func (s *TeamsAppDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowedInstallationScopes *TeamsAppInstallationScopes        `json:"allowedInstallationScopes,omitempty"`
		Authorization             *TeamsAppAuthorization             `json:"authorization,omitempty"`
		AzureADAppId              nullable.Type[string]              `json:"azureADAppId,omitempty"`
		Bot                       *TeamworkBot                       `json:"bot,omitempty"`
		ColorIcon                 *TeamsAppIcon                      `json:"colorIcon,omitempty"`
		DashboardCards            *[]TeamsAppDashboardCardDefinition `json:"dashboardCards,omitempty"`
		Description               nullable.Type[string]              `json:"description,omitempty"`
		DisplayName               nullable.Type[string]              `json:"displayName,omitempty"`
		LastModifiedDateTime      nullable.Type[string]              `json:"lastModifiedDateTime,omitempty"`
		OutlineIcon               *TeamsAppIcon                      `json:"outlineIcon,omitempty"`
		PublishingState           *TeamsAppPublishingState           `json:"publishingState,omitempty"`
		Shortdescription          nullable.Type[string]              `json:"shortdescription,omitempty"`
		TeamsAppId                nullable.Type[string]              `json:"teamsAppId,omitempty"`
		Version                   nullable.Type[string]              `json:"version,omitempty"`
		Id                        *string                            `json:"id,omitempty"`
		ODataId                   *string                            `json:"@odata.id,omitempty"`
		ODataType                 *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowedInstallationScopes = decoded.AllowedInstallationScopes
	s.Authorization = decoded.Authorization
	s.AzureADAppId = decoded.AzureADAppId
	s.Bot = decoded.Bot
	s.ColorIcon = decoded.ColorIcon
	s.DashboardCards = decoded.DashboardCards
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.OutlineIcon = decoded.OutlineIcon
	s.PublishingState = decoded.PublishingState
	s.Shortdescription = decoded.Shortdescription
	s.TeamsAppId = decoded.TeamsAppId
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamsAppDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'TeamsAppDefinition': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
