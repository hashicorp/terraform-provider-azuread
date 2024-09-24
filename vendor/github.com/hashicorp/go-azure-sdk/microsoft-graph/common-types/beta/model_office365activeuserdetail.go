package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Office365ActiveUserDetail{}

type Office365ActiveUserDetail struct {
	// All the products assigned for the user.
	AssignedProducts *[]string `json:"assignedProducts,omitempty"`

	// The date when the delete operation happened. Default value is 'null' when the user hasn't been deleted.
	DeletedDate nullable.Type[string] `json:"deletedDate,omitempty"`

	// The name displayed in the address book for the user. This is usually the combination of the user's first name, middle
	// initial, and last name. This property is required when a user is created and it can't be cleared during updates.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date when user last read or sent email.
	ExchangeLastActivityDate nullable.Type[string] `json:"exchangeLastActivityDate,omitempty"`

	// The last date when the user was assigned an Exchange license.
	ExchangeLicenseAssignDate nullable.Type[string] `json:"exchangeLicenseAssignDate,omitempty"`

	// Whether the user has been assigned an Exchange license.
	HasExchangeLicense nullable.Type[bool] `json:"hasExchangeLicense,omitempty"`

	// Whether the user has been assigned a OneDrive license.
	HasOneDriveLicense nullable.Type[bool] `json:"hasOneDriveLicense,omitempty"`

	// Whether the user has been assigned a SharePoint license.
	HasSharePointLicense nullable.Type[bool] `json:"hasSharePointLicense,omitempty"`

	// Whether the user has been assigned a Skype For Business license.
	HasSkypeForBusinessLicense nullable.Type[bool] `json:"hasSkypeForBusinessLicense,omitempty"`

	// Whether the user has been assigned a Teams license.
	HasTeamsLicense nullable.Type[bool] `json:"hasTeamsLicense,omitempty"`

	// Whether the user has been assigned a Yammer license.
	HasYammerLicense nullable.Type[bool] `json:"hasYammerLicense,omitempty"`

	// Whether this user has been deleted or soft deleted.
	IsDeleted nullable.Type[bool] `json:"isDeleted,omitempty"`

	// The date when user last viewed or edited files, shared files internally or externally, or synced files.
	OneDriveLastActivityDate nullable.Type[string] `json:"oneDriveLastActivityDate,omitempty"`

	// The last date when the user was assigned a OneDrive license.
	OneDriveLicenseAssignDate nullable.Type[string] `json:"oneDriveLicenseAssignDate,omitempty"`

	// The latest date of the content.
	ReportRefreshDate nullable.Type[string] `json:"reportRefreshDate,omitempty"`

	// The date when user last viewed or edited files, shared files internally or externally, synced files, or viewed
	// SharePoint pages.
	SharePointLastActivityDate nullable.Type[string] `json:"sharePointLastActivityDate,omitempty"`

	// The last date when the user was assigned a SharePoint license.
	SharePointLicenseAssignDate nullable.Type[string] `json:"sharePointLicenseAssignDate,omitempty"`

	// The date when user last organized or participated in conferences, or joined peer-to-peer sessions.
	SkypeForBusinessLastActivityDate nullable.Type[string] `json:"skypeForBusinessLastActivityDate,omitempty"`

	// The last date when the user was assigned a Skype For Business license.
	SkypeForBusinessLicenseAssignDate nullable.Type[string] `json:"skypeForBusinessLicenseAssignDate,omitempty"`

	// The date when user last posted messages in team channels, sent messages in private chat sessions, or participated in
	// meetings or calls.
	TeamsLastActivityDate nullable.Type[string] `json:"teamsLastActivityDate,omitempty"`

	// The last date when the user was assigned a Teams license.
	TeamsLicenseAssignDate nullable.Type[string] `json:"teamsLicenseAssignDate,omitempty"`

	// The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet
	// standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where
	// domain must be present in the tenant’s collection of verified domains. This property is required when a user is
	// created.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// The date when user last posted, read, or liked message.
	YammerLastActivityDate nullable.Type[string] `json:"yammerLastActivityDate,omitempty"`

	// The last date when the user was assigned a Yammer license.
	YammerLicenseAssignDate nullable.Type[string] `json:"yammerLicenseAssignDate,omitempty"`

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

func (s Office365ActiveUserDetail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Office365ActiveUserDetail{}

func (s Office365ActiveUserDetail) MarshalJSON() ([]byte, error) {
	type wrapper Office365ActiveUserDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Office365ActiveUserDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Office365ActiveUserDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.office365ActiveUserDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Office365ActiveUserDetail: %+v", err)
	}

	return encoded, nil
}
