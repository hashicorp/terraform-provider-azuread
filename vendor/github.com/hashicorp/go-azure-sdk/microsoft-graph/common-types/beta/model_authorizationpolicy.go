package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PolicyBase = AuthorizationPolicy{}

type AuthorizationPolicy struct {
	// Indicates whether a user can join the tenant by email validation.
	AllowEmailVerifiedUsersToJoinOrganization *bool `json:"allowEmailVerifiedUsersToJoinOrganization,omitempty"`

	// Indicates who can invite guests to the organization. Possible values are: none, adminsAndGuestInviters,
	// adminsGuestInvitersAndAllMembers, everyone. everyone is the default setting for all cloud environments except US
	// Government. For more information, see allowInvitesFrom values.
	AllowInvitesFrom *AllowInvitesFrom `json:"allowInvitesFrom,omitempty"`

	// Indicates whether user consent for risky apps is allowed. Default value is false. We recommend that you keep the
	// value set to false.
	AllowUserConsentForRiskyApps nullable.Type[bool] `json:"allowUserConsentForRiskyApps,omitempty"`

	// Indicates whether users can sign up for email based subscriptions.
	AllowedToSignUpEmailBasedSubscriptions *bool `json:"allowedToSignUpEmailBasedSubscriptions,omitempty"`

	// Indicates whether administrators of the tenant can use the Self-Service Password Reset (SSPR). For more information,
	// see Self-service password reset for administrators.
	AllowedToUseSSPR *bool `json:"allowedToUseSSPR,omitempty"`

	// To disable the use of the MSOnline PowerShell module set this property to true. This also disables user-based access
	// to the legacy service endpoint used by the MSOnline PowerShell module. This doesn't affect Microsoft Entra Connect or
	// Microsoft Graph.
	BlockMsolPowerShell nullable.Type[bool] `json:"blockMsolPowerShell,omitempty"`

	DefaultUserRoleOverrides   *[]DefaultUserRoleOverride  `json:"defaultUserRoleOverrides,omitempty"`
	DefaultUserRolePermissions *DefaultUserRolePermissions `json:"defaultUserRolePermissions,omitempty"`

	// List of features enabled for private preview on the tenant.
	EnabledPreviewFeatures *[]string `json:"enabledPreviewFeatures,omitempty"`

	// Represents role templateId for the role that should be granted to guests. Refer to List unifiedRoleDefinitions to
	// find the list of available role templates. Currently following roles are supported: User
	// (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User
	// (2af84b1e-32c8-42b7-82bc-daa82404023b).
	GuestUserRoleId nullable.Type[string] `json:"guestUserRoleId,omitempty"`

	// Indicates if user consent to apps is allowed, and if it is, the app consent policy that governs the permission for
	// users to grant consent. Values should be in the format managePermissionGrantsForSelf.{id} for user consent policies
	// or managePermissionGrantsForOwnedResource.{id} for resource-specific consent policies, where {id} is the id of a
	// built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
	PermissionGrantPolicyIdsAssignedToDefaultUserRole *[]string `json:"permissionGrantPolicyIdsAssignedToDefaultUserRole,omitempty"`

	// Fields inherited from PolicyBase

	// Description for this policy. Required.
	Description string `json:"description"`

	// Display name for this policy. Required.
	DisplayName string `json:"displayName"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s AuthorizationPolicy) PolicyBase() BasePolicyBaseImpl {
	return BasePolicyBaseImpl{
		Description:     s.Description,
		DisplayName:     s.DisplayName,
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s AuthorizationPolicy) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s AuthorizationPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AuthorizationPolicy{}

func (s AuthorizationPolicy) MarshalJSON() ([]byte, error) {
	type wrapper AuthorizationPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuthorizationPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthorizationPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authorizationPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuthorizationPolicy: %+v", err)
	}

	return encoded, nil
}
