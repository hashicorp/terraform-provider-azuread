// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package groups

const (
	groupResourceName        = "azuread_group"
	groupDuplicateValueError = "Request contains a property with duplicate values"
)

const (
	GroupTypeDynamicMembership = "DynamicMembership"
	GroupTypeUnified           = "Unified"
)

var possibleValuesForGroupType = []string{GroupTypeDynamicMembership, GroupTypeUnified}

const (
	GroupResourceBehaviorOptionAllowOnlyMembersToPost                   = "AllowOnlyMembersToPost"
	GroupResourceBehaviorOptionCalendarMemberReadOnly                   = "CalendarMemberReadOnly"
	GroupResourceBehaviorOptionConnectorsDisabled                       = "ConnectorsDisabled"
	GroupResourceBehaviorOptionHideGroupInOutlook                       = "HideGroupInOutlook"
	GroupResourceBehaviorOptionSkipExchangeInstantOn                    = "SkipExchangeInstantOn"
	GroupResourceBehaviorOptionSubscribeMembersToCalendarEventsDisabled = "SubscribeMembersToCalendarEventsDisabled"
	GroupResourceBehaviorOptionSubscribeNewGroupMembers                 = "SubscribeNewGroupMembers"
	GroupResourceBehaviorOptionWelcomeEmailDisabled                     = "WelcomeEmailDisabled"
)

var possibleValuesForGroupResourceBehaviorOptions = []string{
	GroupResourceBehaviorOptionAllowOnlyMembersToPost,
	GroupResourceBehaviorOptionCalendarMemberReadOnly,
	GroupResourceBehaviorOptionConnectorsDisabled,
	GroupResourceBehaviorOptionHideGroupInOutlook,
	GroupResourceBehaviorOptionSkipExchangeInstantOn,
	GroupResourceBehaviorOptionSubscribeMembersToCalendarEventsDisabled,
	GroupResourceBehaviorOptionSubscribeNewGroupMembers,
	GroupResourceBehaviorOptionWelcomeEmailDisabled,
}

const (
	GroupResourceProvisioningOptionTeam = "Team"
)

var possibleValuesForGroupResourceProvisioningOptions = []string{GroupResourceProvisioningOptionTeam}

const (
	GroupThemeNone   = ""
	GroupThemeBlue   = "Blue"
	GroupThemeGreen  = "Green"
	GroupThemeOrange = "Orange"
	GroupThemePink   = "Pink"
	GroupThemePurple = "Purple"
	GroupThemeRed    = "Red"
	GroupThemeTeal   = "Teal"
)

var possibleValuesForGroupTheme = []string{
	GroupThemeNone,
	GroupThemeBlue,
	GroupThemeGreen,
	GroupThemeOrange,
	GroupThemePink,
	GroupThemePurple,
	GroupThemeRed,
	GroupThemeTeal,
}

const (
	GroupVisibilityHiddenMembership = "Hiddenmembership"
	GroupVisibilityPrivate          = "Private"
	GroupVisibilityPublic           = "Public"
)

var possibleValuesForGroupVisibility = []string{GroupVisibilityHiddenMembership, GroupVisibilityPrivate, GroupVisibilityPublic}

const (
	OnPremisesGroupTypeUniversalDistributionGroup        = "UniversalDistributionGroup"
	OnPremisesGroupTypeUniversalMailEnabledSecurityGroup = "UniversalMailEnabledSecurityGroup"
	OnPremisesGroupTypeUniversalSecurityGroup            = "UniversalSecurityGroup"
)

var possibleValuesForOnPremisesGroupType = []string{OnPremisesGroupTypeUniversalDistributionGroup, OnPremisesGroupTypeUniversalMailEnabledSecurityGroup, OnPremisesGroupTypeUniversalSecurityGroup}
