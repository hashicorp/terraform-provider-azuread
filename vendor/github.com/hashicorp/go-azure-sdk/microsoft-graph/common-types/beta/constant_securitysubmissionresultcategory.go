package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionResultCategory string

const (
	SecuritySubmissionResultCategory_AllowedByPolicy                  SecuritySubmissionResultCategory = "allowedByPolicy"
	SecuritySubmissionResultCategory_AllowedDueToOrganizationOverride SecuritySubmissionResultCategory = "allowedDueToOrganizationOverride"
	SecuritySubmissionResultCategory_AllowedDueToUserOverride         SecuritySubmissionResultCategory = "allowedDueToUserOverride"
	SecuritySubmissionResultCategory_AuthenticationFailure            SecuritySubmissionResultCategory = "authenticationFailure"
	SecuritySubmissionResultCategory_BeingAnalyzed                    SecuritySubmissionResultCategory = "beingAnalyzed"
	SecuritySubmissionResultCategory_BlockedByPolicy                  SecuritySubmissionResultCategory = "blockedByPolicy"
	SecuritySubmissionResultCategory_BlockedDueToOrganizationOverride SecuritySubmissionResultCategory = "blockedDueToOrganizationOverride"
	SecuritySubmissionResultCategory_BlockedDueToUserOverride         SecuritySubmissionResultCategory = "blockedDueToUserOverride"
	SecuritySubmissionResultCategory_BrandImpersonation               SecuritySubmissionResultCategory = "brandImpersonation"
	SecuritySubmissionResultCategory_Bulk                             SecuritySubmissionResultCategory = "bulk"
	SecuritySubmissionResultCategory_DomainImpersonation              SecuritySubmissionResultCategory = "domainImpersonation"
	SecuritySubmissionResultCategory_ItemNotfound                     SecuritySubmissionResultCategory = "itemNotfound"
	SecuritySubmissionResultCategory_Malware                          SecuritySubmissionResultCategory = "malware"
	SecuritySubmissionResultCategory_NoResultAvailable                SecuritySubmissionResultCategory = "noResultAvailable"
	SecuritySubmissionResultCategory_NoThreatsFound                   SecuritySubmissionResultCategory = "noThreatsFound"
	SecuritySubmissionResultCategory_NotJunk                          SecuritySubmissionResultCategory = "notJunk"
	SecuritySubmissionResultCategory_NotSubmittedToMicrosoft          SecuritySubmissionResultCategory = "notSubmittedToMicrosoft"
	SecuritySubmissionResultCategory_Phishing                         SecuritySubmissionResultCategory = "phishing"
	SecuritySubmissionResultCategory_PhishingSimulation               SecuritySubmissionResultCategory = "phishingSimulation"
	SecuritySubmissionResultCategory_ReasonLostInTransit              SecuritySubmissionResultCategory = "reasonLostInTransit"
	SecuritySubmissionResultCategory_Spam                             SecuritySubmissionResultCategory = "spam"
	SecuritySubmissionResultCategory_Spoof                            SecuritySubmissionResultCategory = "spoof"
	SecuritySubmissionResultCategory_SpoofedAllowed                   SecuritySubmissionResultCategory = "spoofedAllowed"
	SecuritySubmissionResultCategory_SpoofedBlocked                   SecuritySubmissionResultCategory = "spoofedBlocked"
	SecuritySubmissionResultCategory_ThreatsFound                     SecuritySubmissionResultCategory = "threatsFound"
	SecuritySubmissionResultCategory_Unknown                          SecuritySubmissionResultCategory = "unknown"
	SecuritySubmissionResultCategory_UserImpersonation                SecuritySubmissionResultCategory = "userImpersonation"
)

func PossibleValuesForSecuritySubmissionResultCategory() []string {
	return []string{
		string(SecuritySubmissionResultCategory_AllowedByPolicy),
		string(SecuritySubmissionResultCategory_AllowedDueToOrganizationOverride),
		string(SecuritySubmissionResultCategory_AllowedDueToUserOverride),
		string(SecuritySubmissionResultCategory_AuthenticationFailure),
		string(SecuritySubmissionResultCategory_BeingAnalyzed),
		string(SecuritySubmissionResultCategory_BlockedByPolicy),
		string(SecuritySubmissionResultCategory_BlockedDueToOrganizationOverride),
		string(SecuritySubmissionResultCategory_BlockedDueToUserOverride),
		string(SecuritySubmissionResultCategory_BrandImpersonation),
		string(SecuritySubmissionResultCategory_Bulk),
		string(SecuritySubmissionResultCategory_DomainImpersonation),
		string(SecuritySubmissionResultCategory_ItemNotfound),
		string(SecuritySubmissionResultCategory_Malware),
		string(SecuritySubmissionResultCategory_NoResultAvailable),
		string(SecuritySubmissionResultCategory_NoThreatsFound),
		string(SecuritySubmissionResultCategory_NotJunk),
		string(SecuritySubmissionResultCategory_NotSubmittedToMicrosoft),
		string(SecuritySubmissionResultCategory_Phishing),
		string(SecuritySubmissionResultCategory_PhishingSimulation),
		string(SecuritySubmissionResultCategory_ReasonLostInTransit),
		string(SecuritySubmissionResultCategory_Spam),
		string(SecuritySubmissionResultCategory_Spoof),
		string(SecuritySubmissionResultCategory_SpoofedAllowed),
		string(SecuritySubmissionResultCategory_SpoofedBlocked),
		string(SecuritySubmissionResultCategory_ThreatsFound),
		string(SecuritySubmissionResultCategory_Unknown),
		string(SecuritySubmissionResultCategory_UserImpersonation),
	}
}

func (s *SecuritySubmissionResultCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionResultCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionResultCategory(input string) (*SecuritySubmissionResultCategory, error) {
	vals := map[string]SecuritySubmissionResultCategory{
		"allowedbypolicy":                  SecuritySubmissionResultCategory_AllowedByPolicy,
		"allowedduetoorganizationoverride": SecuritySubmissionResultCategory_AllowedDueToOrganizationOverride,
		"allowedduetouseroverride":         SecuritySubmissionResultCategory_AllowedDueToUserOverride,
		"authenticationfailure":            SecuritySubmissionResultCategory_AuthenticationFailure,
		"beinganalyzed":                    SecuritySubmissionResultCategory_BeingAnalyzed,
		"blockedbypolicy":                  SecuritySubmissionResultCategory_BlockedByPolicy,
		"blockedduetoorganizationoverride": SecuritySubmissionResultCategory_BlockedDueToOrganizationOverride,
		"blockedduetouseroverride":         SecuritySubmissionResultCategory_BlockedDueToUserOverride,
		"brandimpersonation":               SecuritySubmissionResultCategory_BrandImpersonation,
		"bulk":                             SecuritySubmissionResultCategory_Bulk,
		"domainimpersonation":              SecuritySubmissionResultCategory_DomainImpersonation,
		"itemnotfound":                     SecuritySubmissionResultCategory_ItemNotfound,
		"malware":                          SecuritySubmissionResultCategory_Malware,
		"noresultavailable":                SecuritySubmissionResultCategory_NoResultAvailable,
		"nothreatsfound":                   SecuritySubmissionResultCategory_NoThreatsFound,
		"notjunk":                          SecuritySubmissionResultCategory_NotJunk,
		"notsubmittedtomicrosoft":          SecuritySubmissionResultCategory_NotSubmittedToMicrosoft,
		"phishing":                         SecuritySubmissionResultCategory_Phishing,
		"phishingsimulation":               SecuritySubmissionResultCategory_PhishingSimulation,
		"reasonlostintransit":              SecuritySubmissionResultCategory_ReasonLostInTransit,
		"spam":                             SecuritySubmissionResultCategory_Spam,
		"spoof":                            SecuritySubmissionResultCategory_Spoof,
		"spoofedallowed":                   SecuritySubmissionResultCategory_SpoofedAllowed,
		"spoofedblocked":                   SecuritySubmissionResultCategory_SpoofedBlocked,
		"threatsfound":                     SecuritySubmissionResultCategory_ThreatsFound,
		"unknown":                          SecuritySubmissionResultCategory_Unknown,
		"userimpersonation":                SecuritySubmissionResultCategory_UserImpersonation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionResultCategory(input)
	return &out, nil
}
