package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicySettingType string

const (
	GroupPolicySettingType_Account                 GroupPolicySettingType = "account"
	GroupPolicySettingType_AppLockerRuleCollection GroupPolicySettingType = "appLockerRuleCollection"
	GroupPolicySettingType_AuditSetting            GroupPolicySettingType = "auditSetting"
	GroupPolicySettingType_DataSourcesSettings     GroupPolicySettingType = "dataSourcesSettings"
	GroupPolicySettingType_DevicesSettings         GroupPolicySettingType = "devicesSettings"
	GroupPolicySettingType_DriveMapSettings        GroupPolicySettingType = "driveMapSettings"
	GroupPolicySettingType_EnvironmentVariables    GroupPolicySettingType = "environmentVariables"
	GroupPolicySettingType_FilesSettings           GroupPolicySettingType = "filesSettings"
	GroupPolicySettingType_FolderOptions           GroupPolicySettingType = "folderOptions"
	GroupPolicySettingType_Folders                 GroupPolicySettingType = "folders"
	GroupPolicySettingType_IniFiles                GroupPolicySettingType = "iniFiles"
	GroupPolicySettingType_InternetOptions         GroupPolicySettingType = "internetOptions"
	GroupPolicySettingType_LocalUsersAndGroups     GroupPolicySettingType = "localUsersAndGroups"
	GroupPolicySettingType_NetworkOptions          GroupPolicySettingType = "networkOptions"
	GroupPolicySettingType_NetworkShares           GroupPolicySettingType = "networkShares"
	GroupPolicySettingType_NtServices              GroupPolicySettingType = "ntServices"
	GroupPolicySettingType_Policy                  GroupPolicySettingType = "policy"
	GroupPolicySettingType_PowerOptions            GroupPolicySettingType = "powerOptions"
	GroupPolicySettingType_Printers                GroupPolicySettingType = "printers"
	GroupPolicySettingType_RegionalOptionsSettings GroupPolicySettingType = "regionalOptionsSettings"
	GroupPolicySettingType_RegistrySettings        GroupPolicySettingType = "registrySettings"
	GroupPolicySettingType_ScheduledTasks          GroupPolicySettingType = "scheduledTasks"
	GroupPolicySettingType_SecurityOptions         GroupPolicySettingType = "securityOptions"
	GroupPolicySettingType_ShortcutSettings        GroupPolicySettingType = "shortcutSettings"
	GroupPolicySettingType_StartMenuSettings       GroupPolicySettingType = "startMenuSettings"
	GroupPolicySettingType_Unknown                 GroupPolicySettingType = "unknown"
	GroupPolicySettingType_UserRightsAssignment    GroupPolicySettingType = "userRightsAssignment"
	GroupPolicySettingType_WindowsFirewallSettings GroupPolicySettingType = "windowsFirewallSettings"
)

func PossibleValuesForGroupPolicySettingType() []string {
	return []string{
		string(GroupPolicySettingType_Account),
		string(GroupPolicySettingType_AppLockerRuleCollection),
		string(GroupPolicySettingType_AuditSetting),
		string(GroupPolicySettingType_DataSourcesSettings),
		string(GroupPolicySettingType_DevicesSettings),
		string(GroupPolicySettingType_DriveMapSettings),
		string(GroupPolicySettingType_EnvironmentVariables),
		string(GroupPolicySettingType_FilesSettings),
		string(GroupPolicySettingType_FolderOptions),
		string(GroupPolicySettingType_Folders),
		string(GroupPolicySettingType_IniFiles),
		string(GroupPolicySettingType_InternetOptions),
		string(GroupPolicySettingType_LocalUsersAndGroups),
		string(GroupPolicySettingType_NetworkOptions),
		string(GroupPolicySettingType_NetworkShares),
		string(GroupPolicySettingType_NtServices),
		string(GroupPolicySettingType_Policy),
		string(GroupPolicySettingType_PowerOptions),
		string(GroupPolicySettingType_Printers),
		string(GroupPolicySettingType_RegionalOptionsSettings),
		string(GroupPolicySettingType_RegistrySettings),
		string(GroupPolicySettingType_ScheduledTasks),
		string(GroupPolicySettingType_SecurityOptions),
		string(GroupPolicySettingType_ShortcutSettings),
		string(GroupPolicySettingType_StartMenuSettings),
		string(GroupPolicySettingType_Unknown),
		string(GroupPolicySettingType_UserRightsAssignment),
		string(GroupPolicySettingType_WindowsFirewallSettings),
	}
}

func (s *GroupPolicySettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicySettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicySettingType(input string) (*GroupPolicySettingType, error) {
	vals := map[string]GroupPolicySettingType{
		"account":                 GroupPolicySettingType_Account,
		"applockerrulecollection": GroupPolicySettingType_AppLockerRuleCollection,
		"auditsetting":            GroupPolicySettingType_AuditSetting,
		"datasourcessettings":     GroupPolicySettingType_DataSourcesSettings,
		"devicessettings":         GroupPolicySettingType_DevicesSettings,
		"drivemapsettings":        GroupPolicySettingType_DriveMapSettings,
		"environmentvariables":    GroupPolicySettingType_EnvironmentVariables,
		"filessettings":           GroupPolicySettingType_FilesSettings,
		"folderoptions":           GroupPolicySettingType_FolderOptions,
		"folders":                 GroupPolicySettingType_Folders,
		"inifiles":                GroupPolicySettingType_IniFiles,
		"internetoptions":         GroupPolicySettingType_InternetOptions,
		"localusersandgroups":     GroupPolicySettingType_LocalUsersAndGroups,
		"networkoptions":          GroupPolicySettingType_NetworkOptions,
		"networkshares":           GroupPolicySettingType_NetworkShares,
		"ntservices":              GroupPolicySettingType_NtServices,
		"policy":                  GroupPolicySettingType_Policy,
		"poweroptions":            GroupPolicySettingType_PowerOptions,
		"printers":                GroupPolicySettingType_Printers,
		"regionaloptionssettings": GroupPolicySettingType_RegionalOptionsSettings,
		"registrysettings":        GroupPolicySettingType_RegistrySettings,
		"scheduledtasks":          GroupPolicySettingType_ScheduledTasks,
		"securityoptions":         GroupPolicySettingType_SecurityOptions,
		"shortcutsettings":        GroupPolicySettingType_ShortcutSettings,
		"startmenusettings":       GroupPolicySettingType_StartMenuSettings,
		"unknown":                 GroupPolicySettingType_Unknown,
		"userrightsassignment":    GroupPolicySettingType_UserRightsAssignment,
		"windowsfirewallsettings": GroupPolicySettingType_WindowsFirewallSettings,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicySettingType(input)
	return &out, nil
}
