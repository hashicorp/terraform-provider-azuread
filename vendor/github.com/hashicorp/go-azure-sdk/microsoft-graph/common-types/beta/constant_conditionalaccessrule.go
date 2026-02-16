package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessRule string

const (
	ConditionalAccessRule_Acr                               ConditionalAccessRule = "acr"
	ConditionalAccessRule_AllApps                           ConditionalAccessRule = "allApps"
	ConditionalAccessRule_AllDevicePlatforms                ConditionalAccessRule = "allDevicePlatforms"
	ConditionalAccessRule_AllDeviceStates                   ConditionalAccessRule = "allDeviceStates"
	ConditionalAccessRule_AllDevices                        ConditionalAccessRule = "allDevices"
	ConditionalAccessRule_AllLocations                      ConditionalAccessRule = "allLocations"
	ConditionalAccessRule_AllTrustedLocations               ConditionalAccessRule = "allTrustedLocations"
	ConditionalAccessRule_AllUsers                          ConditionalAccessRule = "allUsers"
	ConditionalAccessRule_AnonymizedIPAddress               ConditionalAccessRule = "anonymizedIPAddress"
	ConditionalAccessRule_AppFilter                         ConditionalAccessRule = "appFilter"
	ConditionalAccessRule_AppId                             ConditionalAccessRule = "appId"
	ConditionalAccessRule_AuthenticationTransfer            ConditionalAccessRule = "authenticationTransfer"
	ConditionalAccessRule_B2bCollaborationGuest             ConditionalAccessRule = "b2bCollaborationGuest"
	ConditionalAccessRule_B2bCollaborationMember            ConditionalAccessRule = "b2bCollaborationMember"
	ConditionalAccessRule_B2bDirectConnectUser              ConditionalAccessRule = "b2bDirectConnectUser"
	ConditionalAccessRule_DeviceCodeFlow                    ConditionalAccessRule = "deviceCodeFlow"
	ConditionalAccessRule_DeviceFilter                      ConditionalAccessRule = "deviceFilter"
	ConditionalAccessRule_DeviceFilterIncludeRuleNotMatched ConditionalAccessRule = "deviceFilterIncludeRuleNotMatched"
	ConditionalAccessRule_DevicePlatform                    ConditionalAccessRule = "devicePlatform"
	ConditionalAccessRule_DeviceState                       ConditionalAccessRule = "deviceState"
	ConditionalAccessRule_FirstPartyApps                    ConditionalAccessRule = "firstPartyApps"
	ConditionalAccessRule_GroupId                           ConditionalAccessRule = "groupId"
	ConditionalAccessRule_Guest                             ConditionalAccessRule = "guest"
	ConditionalAccessRule_InsideCorpnet                     ConditionalAccessRule = "insideCorpnet"
	ConditionalAccessRule_InsiderRisk                       ConditionalAccessRule = "insiderRisk"
	ConditionalAccessRule_InternalGuest                     ConditionalAccessRule = "internalGuest"
	ConditionalAccessRule_LocationId                        ConditionalAccessRule = "locationId"
	ConditionalAccessRule_MicrosoftAdminPortals             ConditionalAccessRule = "microsoftAdminPortals"
	ConditionalAccessRule_NationStateIPAddress              ConditionalAccessRule = "nationStateIPAddress"
	ConditionalAccessRule_Office365                         ConditionalAccessRule = "office365"
	ConditionalAccessRule_OtherExternalUser                 ConditionalAccessRule = "otherExternalUser"
	ConditionalAccessRule_RealTimeThreatIntelligence        ConditionalAccessRule = "realTimeThreatIntelligence"
	ConditionalAccessRule_RoleId                            ConditionalAccessRule = "roleId"
	ConditionalAccessRule_ServiceProvider                   ConditionalAccessRule = "serviceProvider"
	ConditionalAccessRule_UnfamiliarFeatures                ConditionalAccessRule = "unfamiliarFeatures"
	ConditionalAccessRule_UserId                            ConditionalAccessRule = "userId"
)

func PossibleValuesForConditionalAccessRule() []string {
	return []string{
		string(ConditionalAccessRule_Acr),
		string(ConditionalAccessRule_AllApps),
		string(ConditionalAccessRule_AllDevicePlatforms),
		string(ConditionalAccessRule_AllDeviceStates),
		string(ConditionalAccessRule_AllDevices),
		string(ConditionalAccessRule_AllLocations),
		string(ConditionalAccessRule_AllTrustedLocations),
		string(ConditionalAccessRule_AllUsers),
		string(ConditionalAccessRule_AnonymizedIPAddress),
		string(ConditionalAccessRule_AppFilter),
		string(ConditionalAccessRule_AppId),
		string(ConditionalAccessRule_AuthenticationTransfer),
		string(ConditionalAccessRule_B2bCollaborationGuest),
		string(ConditionalAccessRule_B2bCollaborationMember),
		string(ConditionalAccessRule_B2bDirectConnectUser),
		string(ConditionalAccessRule_DeviceCodeFlow),
		string(ConditionalAccessRule_DeviceFilter),
		string(ConditionalAccessRule_DeviceFilterIncludeRuleNotMatched),
		string(ConditionalAccessRule_DevicePlatform),
		string(ConditionalAccessRule_DeviceState),
		string(ConditionalAccessRule_FirstPartyApps),
		string(ConditionalAccessRule_GroupId),
		string(ConditionalAccessRule_Guest),
		string(ConditionalAccessRule_InsideCorpnet),
		string(ConditionalAccessRule_InsiderRisk),
		string(ConditionalAccessRule_InternalGuest),
		string(ConditionalAccessRule_LocationId),
		string(ConditionalAccessRule_MicrosoftAdminPortals),
		string(ConditionalAccessRule_NationStateIPAddress),
		string(ConditionalAccessRule_Office365),
		string(ConditionalAccessRule_OtherExternalUser),
		string(ConditionalAccessRule_RealTimeThreatIntelligence),
		string(ConditionalAccessRule_RoleId),
		string(ConditionalAccessRule_ServiceProvider),
		string(ConditionalAccessRule_UnfamiliarFeatures),
		string(ConditionalAccessRule_UserId),
	}
}

func (s *ConditionalAccessRule) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessRule(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessRule(input string) (*ConditionalAccessRule, error) {
	vals := map[string]ConditionalAccessRule{
		"acr":                               ConditionalAccessRule_Acr,
		"allapps":                           ConditionalAccessRule_AllApps,
		"alldeviceplatforms":                ConditionalAccessRule_AllDevicePlatforms,
		"alldevicestates":                   ConditionalAccessRule_AllDeviceStates,
		"alldevices":                        ConditionalAccessRule_AllDevices,
		"alllocations":                      ConditionalAccessRule_AllLocations,
		"alltrustedlocations":               ConditionalAccessRule_AllTrustedLocations,
		"allusers":                          ConditionalAccessRule_AllUsers,
		"anonymizedipaddress":               ConditionalAccessRule_AnonymizedIPAddress,
		"appfilter":                         ConditionalAccessRule_AppFilter,
		"appid":                             ConditionalAccessRule_AppId,
		"authenticationtransfer":            ConditionalAccessRule_AuthenticationTransfer,
		"b2bcollaborationguest":             ConditionalAccessRule_B2bCollaborationGuest,
		"b2bcollaborationmember":            ConditionalAccessRule_B2bCollaborationMember,
		"b2bdirectconnectuser":              ConditionalAccessRule_B2bDirectConnectUser,
		"devicecodeflow":                    ConditionalAccessRule_DeviceCodeFlow,
		"devicefilter":                      ConditionalAccessRule_DeviceFilter,
		"devicefilterincluderulenotmatched": ConditionalAccessRule_DeviceFilterIncludeRuleNotMatched,
		"deviceplatform":                    ConditionalAccessRule_DevicePlatform,
		"devicestate":                       ConditionalAccessRule_DeviceState,
		"firstpartyapps":                    ConditionalAccessRule_FirstPartyApps,
		"groupid":                           ConditionalAccessRule_GroupId,
		"guest":                             ConditionalAccessRule_Guest,
		"insidecorpnet":                     ConditionalAccessRule_InsideCorpnet,
		"insiderrisk":                       ConditionalAccessRule_InsiderRisk,
		"internalguest":                     ConditionalAccessRule_InternalGuest,
		"locationid":                        ConditionalAccessRule_LocationId,
		"microsoftadminportals":             ConditionalAccessRule_MicrosoftAdminPortals,
		"nationstateipaddress":              ConditionalAccessRule_NationStateIPAddress,
		"office365":                         ConditionalAccessRule_Office365,
		"otherexternaluser":                 ConditionalAccessRule_OtherExternalUser,
		"realtimethreatintelligence":        ConditionalAccessRule_RealTimeThreatIntelligence,
		"roleid":                            ConditionalAccessRule_RoleId,
		"serviceprovider":                   ConditionalAccessRule_ServiceProvider,
		"unfamiliarfeatures":                ConditionalAccessRule_UnfamiliarFeatures,
		"userid":                            ConditionalAccessRule_UserId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessRule(input)
	return &out, nil
}
