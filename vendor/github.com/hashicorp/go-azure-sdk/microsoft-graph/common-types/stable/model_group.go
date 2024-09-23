package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = Group{}

type Group struct {
	// The list of users or groups allowed to create posts or calendar events in this group. If this list is nonempty, then
	// only users or groups listed here are allowed to post.
	AcceptedSenders *[]DirectoryObject `json:"acceptedSenders,omitempty"`

	// List of OData IDs for `AcceptedSenders` to bind to this entity
	AcceptedSenders_ODataBind *[]string `json:"acceptedSenders@odata.bind,omitempty"`

	// Indicates if people external to the organization can send messages to the group. The default value is false. Returned
	// only on $select. Supported only on the Get group API (GET /groups/{ID}).
	AllowExternalSenders nullable.Type[bool] `json:"allowExternalSenders,omitempty"`

	// Represents the app roles granted to a group for an application. Supports $expand.
	AppRoleAssignments *[]AppRoleAssignment `json:"appRoleAssignments,omitempty"`

	// The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on
	// $select.
	AssignedLabels *[]AssignedLabel `json:"assignedLabels,omitempty"`

	// The licenses that are assigned to the group. Returned only on $select. Supports $filter (eq).Read-only.
	AssignedLicenses *[]AssignedLicense `json:"assignedLicenses,omitempty"`

	// Indicates if new members added to the group are autosubscribed to receive email notifications. You can set this
	// property in a PATCH request for the group; don't set it in the initial POST request that creates the group. Default
	// value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
	AutoSubscribeNewMembers nullable.Type[bool] `json:"autoSubscribeNewMembers,omitempty"`

	// The group's calendar. Read-only.
	Calendar *Calendar `json:"calendar,omitempty"`

	// The calendar view for the calendar. Read-only.
	CalendarView *[]Event `json:"calendarView,omitempty"`

	// Describes a classification for the group (such as low, medium, or high business impact). Valid values for this
	// property are defined by creating a ClassificationList setting value, based on the template definition.Returned by
	// default. Supports $filter (eq, ne, not, ge, le, startsWith).
	Classification nullable.Type[string] `json:"classification,omitempty"`

	// The group's conversations.
	Conversations *[]Conversation `json:"conversations,omitempty"`

	// Timestamp of when the group was created. The value can't be modified and is automatically populated when the group is
	// created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For
	// example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The user (or application) that created the group. NOTE: This property isn't set if the user is an administrator.
	// Read-only.
	CreatedOnBehalfOf *DirectoryObject `json:"createdOnBehalfOf,omitempty"`

	// OData ID for `CreatedOnBehalfOf` to bind to this entity
	CreatedOnBehalfOf_ODataBind *string `json:"createdOnBehalfOf@odata.bind,omitempty"`

	// An optional description for the group. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and
	// $search.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the group. This property is required when a group is created and can't be cleared during
	// updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in,
	// startsWith, and eq on null values), $search, and $orderby.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The group's default drive. Read-only.
	Drive *Drive `json:"drive,omitempty"`

	// The group's drives. Read-only.
	Drives *[]Drive `json:"drives,omitempty"`

	// The group's calendar events.
	Events *[]Event `json:"events,omitempty"`

	// Timestamp of when the group is set to expire. It's null for security groups, but for Microsoft 365 groups, it
	// represents when the group is set to expire as defined in the groupLifecyclePolicy. The Timestamp type represents date
	// and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on January 1, 2014 is
	// 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// The collection of open extensions defined for the group. Read-only. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`

	// The collection of lifecycle policies for this group. Read-only. Nullable.
	GroupLifecyclePolicies *[]GroupLifecyclePolicy `json:"groupLifecyclePolicies,omitempty"`

	// Specifies the group type and its membership. If the collection contains Unified, the group is a Microsoft 365 group;
	// otherwise, it's either a security group or a distribution group. For details, see groups overview.If the collection
	// includes DynamicMembership, the group has dynamic membership; otherwise, membership is static. Returned by default.
	// Supports $filter (eq, not).
	GroupTypes *[]string `json:"groupTypes,omitempty"`

	// Indicates whether there are members in this group that have license errors from its group-based license assignment.
	// This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have
	// members with license errors (that is, filter for this property being true). See an example. Supports $filter (eq).
	HasMembersWithLicenseErrors nullable.Type[bool] `json:"hasMembersWithLicenseErrors,omitempty"`

	// True if the group isn't displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting
	// message recipients, and the Browse Groups dialog for searching groups; otherwise, false. Default value is false.
	// Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
	HideFromAddressLists nullable.Type[bool] `json:"hideFromAddressLists,omitempty"`

	// True if the group isn't displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise,
	// false. The default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
	HideFromOutlookClients nullable.Type[bool] `json:"hideFromOutlookClients,omitempty"`

	// When a group is associated with a team, this property determines whether the team is in read-only mode.To read this
	// property, use the /group/{groupId}/team endpoint or the Get team API. To update this property, use the archiveTeam
	// and unarchiveTeam APIs.
	IsArchived nullable.Type[bool] `json:"isArchived,omitempty"`

	// Indicates whether this group can be assigned to a Microsoft Entra role. Optional. This property can only be set while
	// creating the group and is immutable. If set to true, the securityEnabled property must also be set to true,
	// visibility must be Hidden, and the group can't be a dynamic group (that is, groupTypes can't contain
	// DynamicMembership). Only callers with at least the Privileged Role Administrator role can set this property. The
	// caller must also be assigned the RoleManagement.ReadWrite.Directory permission to set this property or update the
	// membership of such groups. For more, see Using a group to manage Microsoft Entra role assignmentsUsing this feature
	// requires a Microsoft Entra ID P1 license. Returned by default. Supports $filter (eq, ne, not).
	IsAssignableToRole nullable.Type[bool] `json:"isAssignableToRole,omitempty"`

	// Indicates whether the signed-in user is subscribed to receive email conversations. The default value is true.
	// Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
	IsSubscribedByMail nullable.Type[bool] `json:"isSubscribedByMail,omitempty"`

	// Indicates the status of the group license assignment to all group members. The default value is false. Read-only.
	// Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select.
	// Read-only.
	LicenseProcessingState *LicenseProcessingState `json:"licenseProcessingState,omitempty"`

	// The SMTP address for the group, for example, 'serviceadmins@contoso.com'. Returned by default. Read-only. Supports
	// $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	Mail nullable.Type[string] `json:"mail,omitempty"`

	// Specifies whether the group is mail-enabled. Required. Returned by default. Supports $filter (eq, ne, not).
	MailEnabled nullable.Type[bool] `json:"mailEnabled,omitempty"`

	// The mail alias for the group, unique for Microsoft 365 groups in the organization. Maximum length is 64 characters.
	// This property can contain only characters in the ASCII character set 0 - 127 except the following characters: @ () /
	// [] ' ; : <> , SPACE. Required. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on
	// null values).
	MailNickname nullable.Type[string] `json:"mailNickname,omitempty"`

	// Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports
	// $expand.
	MemberOf *[]DirectoryObject `json:"memberOf,omitempty"`

	// List of OData IDs for `MemberOf` to bind to this entity
	MemberOf_ODataBind *[]string `json:"memberOf@odata.bind,omitempty"`

	// The members of this group, who can be users, devices, other groups, or service principals. Supports the List members,
	// Add member, and Remove member operations. Nullable. Supports $expand including nested $select. For example,
	// /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=members($select=id,userPrincipalName,displayName).
	Members *[]DirectoryObject `json:"members,omitempty"`

	// A list of group members with license errors from this group-based license assignment. Read-only.
	MembersWithLicenseErrors *[]DirectoryObject `json:"membersWithLicenseErrors,omitempty"`

	// List of OData IDs for `MembersWithLicenseErrors` to bind to this entity
	MembersWithLicenseErrors_ODataBind *[]string `json:"membersWithLicenseErrors@odata.bind,omitempty"`

	// List of OData IDs for `Members` to bind to this entity
	Members_ODataBind *[]string `json:"members@odata.bind,omitempty"`

	// The rule that determines members for this group if the group is a dynamic group (groupTypes contains
	// DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax.
	// Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith).
	MembershipRule nullable.Type[string] `json:"membershipRule,omitempty"`

	// Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by
	// default. Supports $filter (eq, ne, not, in).
	MembershipRuleProcessingState nullable.Type[string] `json:"membershipRuleProcessingState,omitempty"`

	OnPremisesDomainName nullable.Type[string] `json:"onPremisesDomainName,omitempty"`

	// Indicates the last time at which the group was synced with the on-premises directory. The Timestamp type represents
	// date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on January 1,
	// 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in).
	OnPremisesLastSyncDateTime nullable.Type[string] `json:"onPremisesLastSyncDateTime,omitempty"`

	OnPremisesNetBiosName nullable.Type[string] `json:"onPremisesNetBiosName,omitempty"`

	// Errors when using Microsoft synchronization product during provisioning. Returned by default. Supports $filter (eq,
	// not).
	OnPremisesProvisioningErrors *[]OnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`

	// Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated
	// for customers synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect.Returned by
	// default. Supports $filter (eq, ne, not, ge, le, in, startsWith). Read-only.
	OnPremisesSamAccountName nullable.Type[string] `json:"onPremisesSamAccountName,omitempty"`

	// Contains the on-premises security identifier (SID) for the group synchronized from on-premises to the cloud.
	// Read-only. Returned by default. Supports $filter (eq including on null values).
	OnPremisesSecurityIdentifier nullable.Type[string] `json:"onPremisesSecurityIdentifier,omitempty"`

	// true if this group is synced from an on-premises directory; false if this group was originally synced from an
	// on-premises directory but is no longer synced; null if this object has never synced from an on-premises directory
	// (default). Returned by default. Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
	OnPremisesSyncEnabled nullable.Type[bool] `json:"onPremisesSyncEnabled,omitempty"`

	Onenote *Onenote `json:"onenote,omitempty"`

	// The owners of the group. Limited to 100 owners. Nullable. If this property isn't specified when creating a Microsoft
	// 365 group, the calling user is automatically assigned as the group owner. Supports $filter (/$count eq 0, /$count ne
	// 0, /$count eq 1, /$count ne 1). Supports $expand including nested $select. For example,
	// /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
	Owners *[]DirectoryObject `json:"owners,omitempty"`

	// List of OData IDs for `Owners` to bind to this entity
	Owners_ODataBind *[]string `json:"owners@odata.bind,omitempty"`

	PermissionGrants *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`

	// The group's profile photo
	Photo *ProfilePhoto `json:"photo,omitempty"`

	// The profile photos owned by the group. Read-only. Nullable.
	Photos *[]ProfilePhoto `json:"photos,omitempty"`

	// Entry-point to Planner resource that might exist for a Unified Group.
	Planner *PlannerGroup `json:"planner,omitempty"`

	// The preferred data location for the Microsoft 365 group. By default, the group inherits the group creator's preferred
	// data location. To set this property, the calling app must be granted the Directory.ReadWrite.All permission and the
	// user be assigned at least one of the following Microsoft Entra roles: User Account Administrator Directory Writer
	// Exchange Administrator SharePoint Administrator For more information about this property, see OneDrive Online
	// Multi-Geo. Nullable. Returned by default.
	PreferredDataLocation nullable.Type[string] `json:"preferredDataLocation,omitempty"`

	// The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example, en-US. Returned by
	// default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	PreferredLanguage nullable.Type[string] `json:"preferredLanguage,omitempty"`

	// Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp:
	// bob@sales.contoso.com']. The any operator is required to filter expressions on multi-valued properties. Returned by
	// default. Read-only. Not nullable. Supports $filter (eq, not, ge, le, startsWith, endsWith, /$count eq 0, /$count ne
	// 0).
	ProxyAddresses *[]string `json:"proxyAddresses,omitempty"`

	// The list of users or groups not allowed to create posts or calendar events in this group. Nullable
	RejectedSenders *[]DirectoryObject `json:"rejectedSenders,omitempty"`

	// List of OData IDs for `RejectedSenders` to bind to this entity
	RejectedSenders_ODataBind *[]string `json:"rejectedSenders@odata.bind,omitempty"`

	// Timestamp of when the group was last renewed. This value can't be modified directly and is only updated via the renew
	// service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC.
	// For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne,
	// not, ge, le, in). Read-only.
	RenewedDateTime nullable.Type[string] `json:"renewedDateTime,omitempty"`

	// Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, not, in).
	SecurityEnabled nullable.Type[bool] `json:"securityEnabled,omitempty"`

	// Security identifier of the group, used in Windows scenarios. Read-only. Returned by default.
	SecurityIdentifier nullable.Type[string] `json:"securityIdentifier,omitempty"`

	// Errors published by a federated service describing a nontransient, service-specific error regarding the properties or
	// link from a group object. Supports $filter (eq, not, for isResolved and serviceInstance).
	ServiceProvisioningErrors *[]ServiceProvisioningError `json:"serviceProvisioningErrors,omitempty"`

	// Settings that can govern this group's behavior, like whether members can invite guests to the group. Nullable.
	Settings *[]GroupSetting `json:"settings,omitempty"`

	// The list of SharePoint sites in this group. Access the default site with /sites/root.
	Sites *[]Site `json:"sites,omitempty"`

	// The team associated with this group.
	Team *Team `json:"team,omitempty"`

	// Specifies a Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange, or Red.
	// Returned by default.
	Theme nullable.Type[string] `json:"theme,omitempty"`

	// The group's conversation threads. Nullable.
	Threads *[]ConversationThread `json:"threads,omitempty"`

	// The groups that a group is a member of, either directly or through nested membership. Nullable.
	TransitiveMemberOf *[]DirectoryObject `json:"transitiveMemberOf,omitempty"`

	// List of OData IDs for `TransitiveMemberOf` to bind to this entity
	TransitiveMemberOf_ODataBind *[]string `json:"transitiveMemberOf@odata.bind,omitempty"`

	// The direct and transitive members of a group. Nullable.
	TransitiveMembers *[]DirectoryObject `json:"transitiveMembers,omitempty"`

	// List of OData IDs for `TransitiveMembers` to bind to this entity
	TransitiveMembers_ODataBind *[]string `json:"transitiveMembers@odata.bind,omitempty"`

	// The unique identifier that can be assigned to a group and used as an alternate key. Immutable. Read-only.
	UniqueName nullable.Type[string] `json:"uniqueName,omitempty"`

	// Count of conversations that received new posts since the signed-in user last visited the group. Returned only on
	// $select. Supported only on the Get group API (GET /groups/{ID}).
	UnseenCount nullable.Type[int64] `json:"unseenCount,omitempty"`

	// Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or
	// HiddenMembership. HiddenMembership can be set only for Microsoft 365 groups when the groups are created. It can't be
	// updated later. Other values of visibility can be updated after group creation. If visibility value isn't specified
	// during group creation on Microsoft Graph, a security group is created as Private by default, and the Microsoft 365
	// group is Public. Groups assignable to roles are always Private. To learn more, see group visibility options. Returned
	// by default. Nullable.
	Visibility nullable.Type[string] `json:"visibility,omitempty"`

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

func (s Group) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s Group) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Group{}

func (s Group) MarshalJSON() ([]byte, error) {
	type wrapper Group
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Group: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Group: %+v", err)
	}

	delete(decoded, "calendar")
	delete(decoded, "calendarView")
	delete(decoded, "createdDateTime")
	delete(decoded, "createdOnBehalfOf")
	delete(decoded, "drive")
	delete(decoded, "drives")
	delete(decoded, "expirationDateTime")
	delete(decoded, "extensions")
	delete(decoded, "groupLifecyclePolicies")
	delete(decoded, "licenseProcessingState")
	delete(decoded, "mail")
	delete(decoded, "memberOf")
	delete(decoded, "membersWithLicenseErrors")
	delete(decoded, "onPremisesLastSyncDateTime")
	delete(decoded, "onPremisesSamAccountName")
	delete(decoded, "onPremisesSecurityIdentifier")
	delete(decoded, "onPremisesSyncEnabled")
	delete(decoded, "photos")
	delete(decoded, "proxyAddresses")
	delete(decoded, "renewedDateTime")
	delete(decoded, "securityIdentifier")
	delete(decoded, "uniqueName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.group"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Group: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Group{}

func (s *Group) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AcceptedSenders_ODataBind          *[]string                          `json:"acceptedSenders@odata.bind,omitempty"`
		AllowExternalSenders               nullable.Type[bool]                `json:"allowExternalSenders,omitempty"`
		AppRoleAssignments                 *[]AppRoleAssignment               `json:"appRoleAssignments,omitempty"`
		AssignedLabels                     *[]AssignedLabel                   `json:"assignedLabels,omitempty"`
		AssignedLicenses                   *[]AssignedLicense                 `json:"assignedLicenses,omitempty"`
		AutoSubscribeNewMembers            nullable.Type[bool]                `json:"autoSubscribeNewMembers,omitempty"`
		Calendar                           *Calendar                          `json:"calendar,omitempty"`
		CalendarView                       *[]Event                           `json:"calendarView,omitempty"`
		Classification                     nullable.Type[string]              `json:"classification,omitempty"`
		Conversations                      *[]Conversation                    `json:"conversations,omitempty"`
		CreatedDateTime                    nullable.Type[string]              `json:"createdDateTime,omitempty"`
		CreatedOnBehalfOf_ODataBind        *string                            `json:"createdOnBehalfOf@odata.bind,omitempty"`
		Description                        nullable.Type[string]              `json:"description,omitempty"`
		DisplayName                        nullable.Type[string]              `json:"displayName,omitempty"`
		Drive                              *Drive                             `json:"drive,omitempty"`
		Drives                             *[]Drive                           `json:"drives,omitempty"`
		Events                             *[]Event                           `json:"events,omitempty"`
		ExpirationDateTime                 nullable.Type[string]              `json:"expirationDateTime,omitempty"`
		GroupLifecyclePolicies             *[]GroupLifecyclePolicy            `json:"groupLifecyclePolicies,omitempty"`
		GroupTypes                         *[]string                          `json:"groupTypes,omitempty"`
		HasMembersWithLicenseErrors        nullable.Type[bool]                `json:"hasMembersWithLicenseErrors,omitempty"`
		HideFromAddressLists               nullable.Type[bool]                `json:"hideFromAddressLists,omitempty"`
		HideFromOutlookClients             nullable.Type[bool]                `json:"hideFromOutlookClients,omitempty"`
		IsArchived                         nullable.Type[bool]                `json:"isArchived,omitempty"`
		IsAssignableToRole                 nullable.Type[bool]                `json:"isAssignableToRole,omitempty"`
		IsSubscribedByMail                 nullable.Type[bool]                `json:"isSubscribedByMail,omitempty"`
		LicenseProcessingState             *LicenseProcessingState            `json:"licenseProcessingState,omitempty"`
		Mail                               nullable.Type[string]              `json:"mail,omitempty"`
		MailEnabled                        nullable.Type[bool]                `json:"mailEnabled,omitempty"`
		MailNickname                       nullable.Type[string]              `json:"mailNickname,omitempty"`
		MemberOf_ODataBind                 *[]string                          `json:"memberOf@odata.bind,omitempty"`
		MembersWithLicenseErrors_ODataBind *[]string                          `json:"membersWithLicenseErrors@odata.bind,omitempty"`
		Members_ODataBind                  *[]string                          `json:"members@odata.bind,omitempty"`
		MembershipRule                     nullable.Type[string]              `json:"membershipRule,omitempty"`
		MembershipRuleProcessingState      nullable.Type[string]              `json:"membershipRuleProcessingState,omitempty"`
		OnPremisesDomainName               nullable.Type[string]              `json:"onPremisesDomainName,omitempty"`
		OnPremisesLastSyncDateTime         nullable.Type[string]              `json:"onPremisesLastSyncDateTime,omitempty"`
		OnPremisesNetBiosName              nullable.Type[string]              `json:"onPremisesNetBiosName,omitempty"`
		OnPremisesProvisioningErrors       *[]OnPremisesProvisioningError     `json:"onPremisesProvisioningErrors,omitempty"`
		OnPremisesSamAccountName           nullable.Type[string]              `json:"onPremisesSamAccountName,omitempty"`
		OnPremisesSecurityIdentifier       nullable.Type[string]              `json:"onPremisesSecurityIdentifier,omitempty"`
		OnPremisesSyncEnabled              nullable.Type[bool]                `json:"onPremisesSyncEnabled,omitempty"`
		Onenote                            *Onenote                           `json:"onenote,omitempty"`
		Owners_ODataBind                   *[]string                          `json:"owners@odata.bind,omitempty"`
		PermissionGrants                   *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`
		Photo                              *ProfilePhoto                      `json:"photo,omitempty"`
		Photos                             *[]ProfilePhoto                    `json:"photos,omitempty"`
		Planner                            *PlannerGroup                      `json:"planner,omitempty"`
		PreferredDataLocation              nullable.Type[string]              `json:"preferredDataLocation,omitempty"`
		PreferredLanguage                  nullable.Type[string]              `json:"preferredLanguage,omitempty"`
		ProxyAddresses                     *[]string                          `json:"proxyAddresses,omitempty"`
		RejectedSenders_ODataBind          *[]string                          `json:"rejectedSenders@odata.bind,omitempty"`
		RenewedDateTime                    nullable.Type[string]              `json:"renewedDateTime,omitempty"`
		SecurityEnabled                    nullable.Type[bool]                `json:"securityEnabled,omitempty"`
		SecurityIdentifier                 nullable.Type[string]              `json:"securityIdentifier,omitempty"`
		Settings                           *[]GroupSetting                    `json:"settings,omitempty"`
		Sites                              *[]Site                            `json:"sites,omitempty"`
		Team                               *Team                              `json:"team,omitempty"`
		Theme                              nullable.Type[string]              `json:"theme,omitempty"`
		Threads                            *[]ConversationThread              `json:"threads,omitempty"`
		TransitiveMemberOf_ODataBind       *[]string                          `json:"transitiveMemberOf@odata.bind,omitempty"`
		TransitiveMembers_ODataBind        *[]string                          `json:"transitiveMembers@odata.bind,omitempty"`
		UniqueName                         nullable.Type[string]              `json:"uniqueName,omitempty"`
		UnseenCount                        nullable.Type[int64]               `json:"unseenCount,omitempty"`
		Visibility                         nullable.Type[string]              `json:"visibility,omitempty"`
		DeletedDateTime                    nullable.Type[string]              `json:"deletedDateTime,omitempty"`
		Id                                 *string                            `json:"id,omitempty"`
		ODataId                            *string                            `json:"@odata.id,omitempty"`
		ODataType                          *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AcceptedSenders_ODataBind = decoded.AcceptedSenders_ODataBind
	s.AllowExternalSenders = decoded.AllowExternalSenders
	s.AppRoleAssignments = decoded.AppRoleAssignments
	s.AssignedLabels = decoded.AssignedLabels
	s.AssignedLicenses = decoded.AssignedLicenses
	s.AutoSubscribeNewMembers = decoded.AutoSubscribeNewMembers
	s.Calendar = decoded.Calendar
	s.CalendarView = decoded.CalendarView
	s.Classification = decoded.Classification
	s.Conversations = decoded.Conversations
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CreatedOnBehalfOf_ODataBind = decoded.CreatedOnBehalfOf_ODataBind
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Drive = decoded.Drive
	s.Drives = decoded.Drives
	s.Events = decoded.Events
	s.ExpirationDateTime = decoded.ExpirationDateTime
	s.GroupLifecyclePolicies = decoded.GroupLifecyclePolicies
	s.GroupTypes = decoded.GroupTypes
	s.HasMembersWithLicenseErrors = decoded.HasMembersWithLicenseErrors
	s.HideFromAddressLists = decoded.HideFromAddressLists
	s.HideFromOutlookClients = decoded.HideFromOutlookClients
	s.IsArchived = decoded.IsArchived
	s.IsAssignableToRole = decoded.IsAssignableToRole
	s.IsSubscribedByMail = decoded.IsSubscribedByMail
	s.LicenseProcessingState = decoded.LicenseProcessingState
	s.Mail = decoded.Mail
	s.MailEnabled = decoded.MailEnabled
	s.MailNickname = decoded.MailNickname
	s.MemberOf_ODataBind = decoded.MemberOf_ODataBind
	s.MembersWithLicenseErrors_ODataBind = decoded.MembersWithLicenseErrors_ODataBind
	s.Members_ODataBind = decoded.Members_ODataBind
	s.MembershipRule = decoded.MembershipRule
	s.MembershipRuleProcessingState = decoded.MembershipRuleProcessingState
	s.OnPremisesDomainName = decoded.OnPremisesDomainName
	s.OnPremisesLastSyncDateTime = decoded.OnPremisesLastSyncDateTime
	s.OnPremisesNetBiosName = decoded.OnPremisesNetBiosName
	s.OnPremisesProvisioningErrors = decoded.OnPremisesProvisioningErrors
	s.OnPremisesSamAccountName = decoded.OnPremisesSamAccountName
	s.OnPremisesSecurityIdentifier = decoded.OnPremisesSecurityIdentifier
	s.OnPremisesSyncEnabled = decoded.OnPremisesSyncEnabled
	s.Onenote = decoded.Onenote
	s.Owners_ODataBind = decoded.Owners_ODataBind
	s.PermissionGrants = decoded.PermissionGrants
	s.Photo = decoded.Photo
	s.Photos = decoded.Photos
	s.Planner = decoded.Planner
	s.PreferredDataLocation = decoded.PreferredDataLocation
	s.PreferredLanguage = decoded.PreferredLanguage
	s.ProxyAddresses = decoded.ProxyAddresses
	s.RejectedSenders_ODataBind = decoded.RejectedSenders_ODataBind
	s.RenewedDateTime = decoded.RenewedDateTime
	s.SecurityEnabled = decoded.SecurityEnabled
	s.SecurityIdentifier = decoded.SecurityIdentifier
	s.Settings = decoded.Settings
	s.Sites = decoded.Sites
	s.Team = decoded.Team
	s.Theme = decoded.Theme
	s.Threads = decoded.Threads
	s.TransitiveMemberOf_ODataBind = decoded.TransitiveMemberOf_ODataBind
	s.TransitiveMembers_ODataBind = decoded.TransitiveMembers_ODataBind
	s.UniqueName = decoded.UniqueName
	s.UnseenCount = decoded.UnseenCount
	s.Visibility = decoded.Visibility
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Group into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["acceptedSenders"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AcceptedSenders into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AcceptedSenders' for 'Group': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AcceptedSenders = &output
	}

	if v, ok := temp["createdOnBehalfOf"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedOnBehalfOf' for 'Group': %+v", err)
		}
		s.CreatedOnBehalfOf = &impl
	}

	if v, ok := temp["extensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Extensions into list []json.RawMessage: %+v", err)
		}

		output := make([]Extension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Extensions' for 'Group': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Extensions = &output
	}

	if v, ok := temp["memberOf"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MemberOf into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MemberOf' for 'Group': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MemberOf = &output
	}

	if v, ok := temp["members"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Members into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Members' for 'Group': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Members = &output
	}

	if v, ok := temp["membersWithLicenseErrors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MembersWithLicenseErrors into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MembersWithLicenseErrors' for 'Group': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MembersWithLicenseErrors = &output
	}

	if v, ok := temp["owners"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Owners into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Owners' for 'Group': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Owners = &output
	}

	if v, ok := temp["rejectedSenders"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RejectedSenders into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RejectedSenders' for 'Group': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RejectedSenders = &output
	}

	if v, ok := temp["serviceProvisioningErrors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ServiceProvisioningErrors into list []json.RawMessage: %+v", err)
		}

		output := make([]ServiceProvisioningError, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalServiceProvisioningErrorImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ServiceProvisioningErrors' for 'Group': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ServiceProvisioningErrors = &output
	}

	if v, ok := temp["transitiveMemberOf"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling TransitiveMemberOf into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'TransitiveMemberOf' for 'Group': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TransitiveMemberOf = &output
	}

	if v, ok := temp["transitiveMembers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling TransitiveMembers into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'TransitiveMembers' for 'Group': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TransitiveMembers = &output
	}

	return nil
}
