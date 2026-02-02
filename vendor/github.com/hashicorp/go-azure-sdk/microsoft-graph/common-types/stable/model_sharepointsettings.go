package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SharepointSettings{}

type SharepointSettings struct {
	// Collection of trusted domain GUIDs for the OneDrive sync app.
	AllowedDomainGuidsForSyncApp *[]string `json:"allowedDomainGuidsForSyncApp,omitempty"`

	// Collection of managed paths available for site creation. Read-only.
	AvailableManagedPathsForSiteCreation *[]string `json:"availableManagedPathsForSiteCreation,omitempty"`

	// The number of days for preserving a deleted user's OneDrive.
	DeletedUserPersonalSiteRetentionPeriodInDays nullable.Type[int64] `json:"deletedUserPersonalSiteRetentionPeriodInDays,omitempty"`

	// Collection of file extensions not uploaded by the OneDrive sync app.
	ExcludedFileExtensionsForSyncApp *[]string `json:"excludedFileExtensionsForSyncApp,omitempty"`

	// Specifies the idle session sign-out policies for the tenant.
	IdleSessionSignOut *IdleSessionSignOut `json:"idleSessionSignOut,omitempty"`

	// Specifies the image tagging option for the tenant. Possible values are: disabled, basic, enhanced.
	ImageTaggingOption *ImageTaggingChoice `json:"imageTaggingOption,omitempty"`

	// Indicates whether comments are allowed on modern site pages in SharePoint.
	IsCommentingOnSitePagesEnabled nullable.Type[bool] `json:"isCommentingOnSitePagesEnabled,omitempty"`

	// Indicates whether push notifications are enabled for OneDrive events.
	IsFileActivityNotificationEnabled nullable.Type[bool] `json:"isFileActivityNotificationEnabled,omitempty"`

	// Indicates whether legacy authentication protocols are enabled for the tenant.
	IsLegacyAuthProtocolsEnabled nullable.Type[bool] `json:"isLegacyAuthProtocolsEnabled,omitempty"`

	// Indicates whether if Fluid Framework is allowed on SharePoint sites.
	IsLoopEnabled nullable.Type[bool] `json:"isLoopEnabled,omitempty"`

	// Indicates whether files can be synced using the OneDrive sync app for Mac.
	IsMacSyncAppEnabled nullable.Type[bool] `json:"isMacSyncAppEnabled,omitempty"`

	// Indicates whether guests must sign in using the same account to which sharing invitations are sent.
	IsRequireAcceptingUserToMatchInvitedUserEnabled nullable.Type[bool] `json:"isRequireAcceptingUserToMatchInvitedUserEnabled,omitempty"`

	// Indicates whether guests are allowed to reshare files, folders, and sites they don't own.
	IsResharingByExternalUsersEnabled nullable.Type[bool] `json:"isResharingByExternalUsersEnabled,omitempty"`

	// Indicates whether mobile push notifications are enabled for SharePoint.
	IsSharePointMobileNotificationEnabled nullable.Type[bool] `json:"isSharePointMobileNotificationEnabled,omitempty"`

	// Indicates whether the newsfeed is allowed on the modern site pages in SharePoint.
	IsSharePointNewsfeedEnabled nullable.Type[bool] `json:"isSharePointNewsfeedEnabled,omitempty"`

	// Indicates whether users are allowed to create sites.
	IsSiteCreationEnabled nullable.Type[bool] `json:"isSiteCreationEnabled,omitempty"`

	// Indicates whether the UI commands for creating sites are shown.
	IsSiteCreationUIEnabled nullable.Type[bool] `json:"isSiteCreationUIEnabled,omitempty"`

	// Indicates whether creating new modern pages is allowed on SharePoint sites.
	IsSitePagesCreationEnabled nullable.Type[bool] `json:"isSitePagesCreationEnabled,omitempty"`

	// Indicates whether site storage space is automatically managed or if specific storage limits are set per site.
	IsSitesStorageLimitAutomatic nullable.Type[bool] `json:"isSitesStorageLimitAutomatic,omitempty"`

	// Indicates whether the sync button in OneDrive is hidden.
	IsSyncButtonHiddenOnPersonalSite nullable.Type[bool] `json:"isSyncButtonHiddenOnPersonalSite,omitempty"`

	// Indicates whether users are allowed to sync files only on PCs joined to specific domains.
	IsUnmanagedSyncAppForTenantRestricted nullable.Type[bool] `json:"isUnmanagedSyncAppForTenantRestricted,omitempty"`

	// The default OneDrive storage limit for all new and existing users who are assigned a qualifying license. Measured in
	// megabytes (MB).
	PersonalSiteDefaultStorageLimitInMB nullable.Type[int64] `json:"personalSiteDefaultStorageLimitInMB,omitempty"`

	// Collection of email domains that are allowed for sharing outside the organization.
	SharingAllowedDomainList *[]string `json:"sharingAllowedDomainList,omitempty"`

	// Collection of email domains that are blocked for sharing outside the organization.
	SharingBlockedDomainList *[]string `json:"sharingBlockedDomainList,omitempty"`

	// Sharing capability for the tenant. Possible values are: disabled, externalUserSharingOnly,
	// externalUserAndGuestSharing, existingExternalUserSharingOnly.
	SharingCapability *SharingCapabilities `json:"sharingCapability,omitempty"`

	// Specifies the external sharing mode for domains. Possible values are: none, allowList, blockList.
	SharingDomainRestrictionMode *SharingDomainRestrictionMode `json:"sharingDomainRestrictionMode,omitempty"`

	// The value of the team site managed path. This is the path under which new team sites will be created.
	SiteCreationDefaultManagedPath *string `json:"siteCreationDefaultManagedPath,omitempty"`

	// The default storage quota for a new site upon creation. Measured in megabytes (MB).
	SiteCreationDefaultStorageLimitInMB nullable.Type[int64] `json:"siteCreationDefaultStorageLimitInMB,omitempty"`

	// The default timezone of a tenant for newly created sites. For a list of possible values, see
	// SPRegionalSettings.TimeZones property.
	TenantDefaultTimezone nullable.Type[string] `json:"tenantDefaultTimezone,omitempty"`

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

func (s SharepointSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SharepointSettings{}

func (s SharepointSettings) MarshalJSON() ([]byte, error) {
	type wrapper SharepointSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SharepointSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SharepointSettings: %+v", err)
	}

	delete(decoded, "availableManagedPathsForSiteCreation")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sharepointSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SharepointSettings: %+v", err)
	}

	return encoded, nil
}
