package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionResultDetail string

const (
	SecuritySubmissionResultDetail_AllowedByAdvancedDelivery              SecuritySubmissionResultDetail = "allowedByAdvancedDelivery"
	SecuritySubmissionResultDetail_AllowedByConnection                    SecuritySubmissionResultDetail = "allowedByConnection"
	SecuritySubmissionResultDetail_AllowedByEnhancedFiltering             SecuritySubmissionResultDetail = "allowedByEnhancedFiltering"
	SecuritySubmissionResultDetail_AllowedByExchangeTransportRule         SecuritySubmissionResultDetail = "allowedByExchangeTransportRule"
	SecuritySubmissionResultDetail_AllowedBySecOps                        SecuritySubmissionResultDetail = "allowedBySecOps"
	SecuritySubmissionResultDetail_AllowedByTenant                        SecuritySubmissionResultDetail = "allowedByTenant"
	SecuritySubmissionResultDetail_AllowedByTenantAllowBlockList          SecuritySubmissionResultDetail = "allowedByTenantAllowBlockList"
	SecuritySubmissionResultDetail_AllowedByThirdPartyFilters             SecuritySubmissionResultDetail = "allowedByThirdPartyFilters"
	SecuritySubmissionResultDetail_AllowedByUserSetting                   SecuritySubmissionResultDetail = "allowedByUserSetting"
	SecuritySubmissionResultDetail_AllowedFileByTenantAllowBlockList      SecuritySubmissionResultDetail = "allowedFileByTenantAllowBlockList"
	SecuritySubmissionResultDetail_AllowedRecipientByTenantAllowBlockList SecuritySubmissionResultDetail = "allowedRecipientByTenantAllowBlockList"
	SecuritySubmissionResultDetail_AllowedSenderByTenantAllowBlockList    SecuritySubmissionResultDetail = "allowedSenderByTenantAllowBlockList"
	SecuritySubmissionResultDetail_AllowedUrlByTenantAllowBlockList       SecuritySubmissionResultDetail = "allowedUrlByTenantAllowBlockList"
	SecuritySubmissionResultDetail_AssociatedWithBrand                    SecuritySubmissionResultDetail = "associatedWithBrand"
	SecuritySubmissionResultDetail_BadReclassifiedAsBad                   SecuritySubmissionResultDetail = "badReclassifiedAsBad"
	SecuritySubmissionResultDetail_BadReclassifiedAsBulk                  SecuritySubmissionResultDetail = "badReclassifiedAsBulk"
	SecuritySubmissionResultDetail_BadReclassifiedAsCannotMakeDecision    SecuritySubmissionResultDetail = "badReclassifiedAsCannotMakeDecision"
	SecuritySubmissionResultDetail_BadReclassifiedAsGood                  SecuritySubmissionResultDetail = "badReclassifiedAsGood"
	SecuritySubmissionResultDetail_BlockedByConnection                    SecuritySubmissionResultDetail = "blockedByConnection"
	SecuritySubmissionResultDetail_BlockedByExchangeTransportRule         SecuritySubmissionResultDetail = "blockedByExchangeTransportRule"
	SecuritySubmissionResultDetail_BlockedByTenant                        SecuritySubmissionResultDetail = "blockedByTenant"
	SecuritySubmissionResultDetail_BlockedByTenantAllowBlockList          SecuritySubmissionResultDetail = "blockedByTenantAllowBlockList"
	SecuritySubmissionResultDetail_BlockedByUserSetting                   SecuritySubmissionResultDetail = "blockedByUserSetting"
	SecuritySubmissionResultDetail_BlockedFileByTenantAllowBlockList      SecuritySubmissionResultDetail = "blockedFileByTenantAllowBlockList"
	SecuritySubmissionResultDetail_BlockedRecipientByTenantAllowBlockList SecuritySubmissionResultDetail = "blockedRecipientByTenantAllowBlockList"
	SecuritySubmissionResultDetail_BlockedSenderByTenantAllowBlockList    SecuritySubmissionResultDetail = "blockedSenderByTenantAllowBlockList"
	SecuritySubmissionResultDetail_BlockedUrlByTenantAllowBlockList       SecuritySubmissionResultDetail = "blockedUrlByTenantAllowBlockList"
	SecuritySubmissionResultDetail_BrandImpersonation                     SecuritySubmissionResultDetail = "brandImpersonation"
	SecuritySubmissionResultDetail_CheckUserReportedSettings              SecuritySubmissionResultDetail = "checkUserReportedSettings"
	SecuritySubmissionResultDetail_DomainImpersonation                    SecuritySubmissionResultDetail = "domainImpersonation"
	SecuritySubmissionResultDetail_DomainResembledYourOrganization        SecuritySubmissionResultDetail = "domainResembledYourOrganization"
	SecuritySubmissionResultDetail_EndUserBeingImpersonated               SecuritySubmissionResultDetail = "endUserBeingImpersonated"
	SecuritySubmissionResultDetail_EndUserBeingSpoofed                    SecuritySubmissionResultDetail = "endUserBeingSpoofed"
	SecuritySubmissionResultDetail_GoodReclassifiedAsBad                  SecuritySubmissionResultDetail = "goodReclassifiedAsBad"
	SecuritySubmissionResultDetail_GoodReclassifiedAsBulk                 SecuritySubmissionResultDetail = "goodReclassifiedAsBulk"
	SecuritySubmissionResultDetail_GoodReclassifiedAsCannotMakeDecision   SecuritySubmissionResultDetail = "goodReclassifiedAsCannotMakeDecision"
	SecuritySubmissionResultDetail_GoodReclassifiedAsGood                 SecuritySubmissionResultDetail = "goodReclassifiedAsGood"
	SecuritySubmissionResultDetail_InvalidFalseNegative                   SecuritySubmissionResultDetail = "invalidFalseNegative"
	SecuritySubmissionResultDetail_InvalidFalsePositive                   SecuritySubmissionResultDetail = "invalidFalsePositive"
	SecuritySubmissionResultDetail_ItemDeleted                            SecuritySubmissionResultDetail = "itemDeleted"
	SecuritySubmissionResultDetail_ItemFoundBulk                          SecuritySubmissionResultDetail = "itemFoundBulk"
	SecuritySubmissionResultDetail_ItemFoundClean                         SecuritySubmissionResultDetail = "itemFoundClean"
	SecuritySubmissionResultDetail_ItemFoundMalicious                     SecuritySubmissionResultDetail = "itemFoundMalicious"
	SecuritySubmissionResultDetail_ItemFoundSpam                          SecuritySubmissionResultDetail = "itemFoundSpam"
	SecuritySubmissionResultDetail_ItemNotReceivedByService               SecuritySubmissionResultDetail = "itemNotReceivedByService"
	SecuritySubmissionResultDetail_JunkMailRuleDisabled                   SecuritySubmissionResultDetail = "junkMailRuleDisabled"
	SecuritySubmissionResultDetail_MessageNotFound                        SecuritySubmissionResultDetail = "messageNotFound"
	SecuritySubmissionResultDetail_None                                   SecuritySubmissionResultDetail = "none"
	SecuritySubmissionResultDetail_OnPremisesSkip                         SecuritySubmissionResultDetail = "onPremisesSkip"
	SecuritySubmissionResultDetail_OutboundBulk                           SecuritySubmissionResultDetail = "outboundBulk"
	SecuritySubmissionResultDetail_OutboundCannotMakeDecision             SecuritySubmissionResultDetail = "outboundCannotMakeDecision"
	SecuritySubmissionResultDetail_OutboundNotRescanned                   SecuritySubmissionResultDetail = "outboundNotRescanned"
	SecuritySubmissionResultDetail_OutboundShouldBeBlocked                SecuritySubmissionResultDetail = "outboundShouldBeBlocked"
	SecuritySubmissionResultDetail_OutboundShouldNotBeBlocked             SecuritySubmissionResultDetail = "outboundShouldNotBeBlocked"
	SecuritySubmissionResultDetail_PartOfEducationCampaign                SecuritySubmissionResultDetail = "partOfEducationCampaign"
	SecuritySubmissionResultDetail_QuarantineReleased                     SecuritySubmissionResultDetail = "quarantineReleased"
	SecuritySubmissionResultDetail_QuarantineReleasedThenBlocked          SecuritySubmissionResultDetail = "quarantineReleasedThenBlocked"
	SecuritySubmissionResultDetail_SenderFailedAuthentication             SecuritySubmissionResultDetail = "senderFailedAuthentication"
	SecuritySubmissionResultDetail_SimulatedThreat                        SecuritySubmissionResultDetail = "simulatedThreat"
	SecuritySubmissionResultDetail_SpoofBlocked                           SecuritySubmissionResultDetail = "spoofBlocked"
	SecuritySubmissionResultDetail_UnableToMakeDecision                   SecuritySubmissionResultDetail = "unableToMakeDecision"
	SecuritySubmissionResultDetail_UnderInvestigation                     SecuritySubmissionResultDetail = "underInvestigation"
	SecuritySubmissionResultDetail_UrlFileCannotMakeDecision              SecuritySubmissionResultDetail = "urlFileCannotMakeDecision"
	SecuritySubmissionResultDetail_UrlFileShouldBeBlocked                 SecuritySubmissionResultDetail = "urlFileShouldBeBlocked"
	SecuritySubmissionResultDetail_UrlFileShouldNotBeBlocked              SecuritySubmissionResultDetail = "urlFileShouldNotBeBlocked"
	SecuritySubmissionResultDetail_UserImpersonation                      SecuritySubmissionResultDetail = "userImpersonation"
	SecuritySubmissionResultDetail_WillNotifyOnceDone                     SecuritySubmissionResultDetail = "willNotifyOnceDone"
	SecuritySubmissionResultDetail_ZeroHourAutoPurgeAllowed               SecuritySubmissionResultDetail = "zeroHourAutoPurgeAllowed"
	SecuritySubmissionResultDetail_ZeroHourAutoPurgeBlocked               SecuritySubmissionResultDetail = "zeroHourAutoPurgeBlocked"
	SecuritySubmissionResultDetail_ZeroHourAutoPurgeQuarantineReleased    SecuritySubmissionResultDetail = "zeroHourAutoPurgeQuarantineReleased"
)

func PossibleValuesForSecuritySubmissionResultDetail() []string {
	return []string{
		string(SecuritySubmissionResultDetail_AllowedByAdvancedDelivery),
		string(SecuritySubmissionResultDetail_AllowedByConnection),
		string(SecuritySubmissionResultDetail_AllowedByEnhancedFiltering),
		string(SecuritySubmissionResultDetail_AllowedByExchangeTransportRule),
		string(SecuritySubmissionResultDetail_AllowedBySecOps),
		string(SecuritySubmissionResultDetail_AllowedByTenant),
		string(SecuritySubmissionResultDetail_AllowedByTenantAllowBlockList),
		string(SecuritySubmissionResultDetail_AllowedByThirdPartyFilters),
		string(SecuritySubmissionResultDetail_AllowedByUserSetting),
		string(SecuritySubmissionResultDetail_AllowedFileByTenantAllowBlockList),
		string(SecuritySubmissionResultDetail_AllowedRecipientByTenantAllowBlockList),
		string(SecuritySubmissionResultDetail_AllowedSenderByTenantAllowBlockList),
		string(SecuritySubmissionResultDetail_AllowedUrlByTenantAllowBlockList),
		string(SecuritySubmissionResultDetail_AssociatedWithBrand),
		string(SecuritySubmissionResultDetail_BadReclassifiedAsBad),
		string(SecuritySubmissionResultDetail_BadReclassifiedAsBulk),
		string(SecuritySubmissionResultDetail_BadReclassifiedAsCannotMakeDecision),
		string(SecuritySubmissionResultDetail_BadReclassifiedAsGood),
		string(SecuritySubmissionResultDetail_BlockedByConnection),
		string(SecuritySubmissionResultDetail_BlockedByExchangeTransportRule),
		string(SecuritySubmissionResultDetail_BlockedByTenant),
		string(SecuritySubmissionResultDetail_BlockedByTenantAllowBlockList),
		string(SecuritySubmissionResultDetail_BlockedByUserSetting),
		string(SecuritySubmissionResultDetail_BlockedFileByTenantAllowBlockList),
		string(SecuritySubmissionResultDetail_BlockedRecipientByTenantAllowBlockList),
		string(SecuritySubmissionResultDetail_BlockedSenderByTenantAllowBlockList),
		string(SecuritySubmissionResultDetail_BlockedUrlByTenantAllowBlockList),
		string(SecuritySubmissionResultDetail_BrandImpersonation),
		string(SecuritySubmissionResultDetail_CheckUserReportedSettings),
		string(SecuritySubmissionResultDetail_DomainImpersonation),
		string(SecuritySubmissionResultDetail_DomainResembledYourOrganization),
		string(SecuritySubmissionResultDetail_EndUserBeingImpersonated),
		string(SecuritySubmissionResultDetail_EndUserBeingSpoofed),
		string(SecuritySubmissionResultDetail_GoodReclassifiedAsBad),
		string(SecuritySubmissionResultDetail_GoodReclassifiedAsBulk),
		string(SecuritySubmissionResultDetail_GoodReclassifiedAsCannotMakeDecision),
		string(SecuritySubmissionResultDetail_GoodReclassifiedAsGood),
		string(SecuritySubmissionResultDetail_InvalidFalseNegative),
		string(SecuritySubmissionResultDetail_InvalidFalsePositive),
		string(SecuritySubmissionResultDetail_ItemDeleted),
		string(SecuritySubmissionResultDetail_ItemFoundBulk),
		string(SecuritySubmissionResultDetail_ItemFoundClean),
		string(SecuritySubmissionResultDetail_ItemFoundMalicious),
		string(SecuritySubmissionResultDetail_ItemFoundSpam),
		string(SecuritySubmissionResultDetail_ItemNotReceivedByService),
		string(SecuritySubmissionResultDetail_JunkMailRuleDisabled),
		string(SecuritySubmissionResultDetail_MessageNotFound),
		string(SecuritySubmissionResultDetail_None),
		string(SecuritySubmissionResultDetail_OnPremisesSkip),
		string(SecuritySubmissionResultDetail_OutboundBulk),
		string(SecuritySubmissionResultDetail_OutboundCannotMakeDecision),
		string(SecuritySubmissionResultDetail_OutboundNotRescanned),
		string(SecuritySubmissionResultDetail_OutboundShouldBeBlocked),
		string(SecuritySubmissionResultDetail_OutboundShouldNotBeBlocked),
		string(SecuritySubmissionResultDetail_PartOfEducationCampaign),
		string(SecuritySubmissionResultDetail_QuarantineReleased),
		string(SecuritySubmissionResultDetail_QuarantineReleasedThenBlocked),
		string(SecuritySubmissionResultDetail_SenderFailedAuthentication),
		string(SecuritySubmissionResultDetail_SimulatedThreat),
		string(SecuritySubmissionResultDetail_SpoofBlocked),
		string(SecuritySubmissionResultDetail_UnableToMakeDecision),
		string(SecuritySubmissionResultDetail_UnderInvestigation),
		string(SecuritySubmissionResultDetail_UrlFileCannotMakeDecision),
		string(SecuritySubmissionResultDetail_UrlFileShouldBeBlocked),
		string(SecuritySubmissionResultDetail_UrlFileShouldNotBeBlocked),
		string(SecuritySubmissionResultDetail_UserImpersonation),
		string(SecuritySubmissionResultDetail_WillNotifyOnceDone),
		string(SecuritySubmissionResultDetail_ZeroHourAutoPurgeAllowed),
		string(SecuritySubmissionResultDetail_ZeroHourAutoPurgeBlocked),
		string(SecuritySubmissionResultDetail_ZeroHourAutoPurgeQuarantineReleased),
	}
}

func (s *SecuritySubmissionResultDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionResultDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionResultDetail(input string) (*SecuritySubmissionResultDetail, error) {
	vals := map[string]SecuritySubmissionResultDetail{
		"allowedbyadvanceddelivery":              SecuritySubmissionResultDetail_AllowedByAdvancedDelivery,
		"allowedbyconnection":                    SecuritySubmissionResultDetail_AllowedByConnection,
		"allowedbyenhancedfiltering":             SecuritySubmissionResultDetail_AllowedByEnhancedFiltering,
		"allowedbyexchangetransportrule":         SecuritySubmissionResultDetail_AllowedByExchangeTransportRule,
		"allowedbysecops":                        SecuritySubmissionResultDetail_AllowedBySecOps,
		"allowedbytenant":                        SecuritySubmissionResultDetail_AllowedByTenant,
		"allowedbytenantallowblocklist":          SecuritySubmissionResultDetail_AllowedByTenantAllowBlockList,
		"allowedbythirdpartyfilters":             SecuritySubmissionResultDetail_AllowedByThirdPartyFilters,
		"allowedbyusersetting":                   SecuritySubmissionResultDetail_AllowedByUserSetting,
		"allowedfilebytenantallowblocklist":      SecuritySubmissionResultDetail_AllowedFileByTenantAllowBlockList,
		"allowedrecipientbytenantallowblocklist": SecuritySubmissionResultDetail_AllowedRecipientByTenantAllowBlockList,
		"allowedsenderbytenantallowblocklist":    SecuritySubmissionResultDetail_AllowedSenderByTenantAllowBlockList,
		"allowedurlbytenantallowblocklist":       SecuritySubmissionResultDetail_AllowedUrlByTenantAllowBlockList,
		"associatedwithbrand":                    SecuritySubmissionResultDetail_AssociatedWithBrand,
		"badreclassifiedasbad":                   SecuritySubmissionResultDetail_BadReclassifiedAsBad,
		"badreclassifiedasbulk":                  SecuritySubmissionResultDetail_BadReclassifiedAsBulk,
		"badreclassifiedascannotmakedecision":    SecuritySubmissionResultDetail_BadReclassifiedAsCannotMakeDecision,
		"badreclassifiedasgood":                  SecuritySubmissionResultDetail_BadReclassifiedAsGood,
		"blockedbyconnection":                    SecuritySubmissionResultDetail_BlockedByConnection,
		"blockedbyexchangetransportrule":         SecuritySubmissionResultDetail_BlockedByExchangeTransportRule,
		"blockedbytenant":                        SecuritySubmissionResultDetail_BlockedByTenant,
		"blockedbytenantallowblocklist":          SecuritySubmissionResultDetail_BlockedByTenantAllowBlockList,
		"blockedbyusersetting":                   SecuritySubmissionResultDetail_BlockedByUserSetting,
		"blockedfilebytenantallowblocklist":      SecuritySubmissionResultDetail_BlockedFileByTenantAllowBlockList,
		"blockedrecipientbytenantallowblocklist": SecuritySubmissionResultDetail_BlockedRecipientByTenantAllowBlockList,
		"blockedsenderbytenantallowblocklist":    SecuritySubmissionResultDetail_BlockedSenderByTenantAllowBlockList,
		"blockedurlbytenantallowblocklist":       SecuritySubmissionResultDetail_BlockedUrlByTenantAllowBlockList,
		"brandimpersonation":                     SecuritySubmissionResultDetail_BrandImpersonation,
		"checkuserreportedsettings":              SecuritySubmissionResultDetail_CheckUserReportedSettings,
		"domainimpersonation":                    SecuritySubmissionResultDetail_DomainImpersonation,
		"domainresembledyourorganization":        SecuritySubmissionResultDetail_DomainResembledYourOrganization,
		"enduserbeingimpersonated":               SecuritySubmissionResultDetail_EndUserBeingImpersonated,
		"enduserbeingspoofed":                    SecuritySubmissionResultDetail_EndUserBeingSpoofed,
		"goodreclassifiedasbad":                  SecuritySubmissionResultDetail_GoodReclassifiedAsBad,
		"goodreclassifiedasbulk":                 SecuritySubmissionResultDetail_GoodReclassifiedAsBulk,
		"goodreclassifiedascannotmakedecision":   SecuritySubmissionResultDetail_GoodReclassifiedAsCannotMakeDecision,
		"goodreclassifiedasgood":                 SecuritySubmissionResultDetail_GoodReclassifiedAsGood,
		"invalidfalsenegative":                   SecuritySubmissionResultDetail_InvalidFalseNegative,
		"invalidfalsepositive":                   SecuritySubmissionResultDetail_InvalidFalsePositive,
		"itemdeleted":                            SecuritySubmissionResultDetail_ItemDeleted,
		"itemfoundbulk":                          SecuritySubmissionResultDetail_ItemFoundBulk,
		"itemfoundclean":                         SecuritySubmissionResultDetail_ItemFoundClean,
		"itemfoundmalicious":                     SecuritySubmissionResultDetail_ItemFoundMalicious,
		"itemfoundspam":                          SecuritySubmissionResultDetail_ItemFoundSpam,
		"itemnotreceivedbyservice":               SecuritySubmissionResultDetail_ItemNotReceivedByService,
		"junkmailruledisabled":                   SecuritySubmissionResultDetail_JunkMailRuleDisabled,
		"messagenotfound":                        SecuritySubmissionResultDetail_MessageNotFound,
		"none":                                   SecuritySubmissionResultDetail_None,
		"onpremisesskip":                         SecuritySubmissionResultDetail_OnPremisesSkip,
		"outboundbulk":                           SecuritySubmissionResultDetail_OutboundBulk,
		"outboundcannotmakedecision":             SecuritySubmissionResultDetail_OutboundCannotMakeDecision,
		"outboundnotrescanned":                   SecuritySubmissionResultDetail_OutboundNotRescanned,
		"outboundshouldbeblocked":                SecuritySubmissionResultDetail_OutboundShouldBeBlocked,
		"outboundshouldnotbeblocked":             SecuritySubmissionResultDetail_OutboundShouldNotBeBlocked,
		"partofeducationcampaign":                SecuritySubmissionResultDetail_PartOfEducationCampaign,
		"quarantinereleased":                     SecuritySubmissionResultDetail_QuarantineReleased,
		"quarantinereleasedthenblocked":          SecuritySubmissionResultDetail_QuarantineReleasedThenBlocked,
		"senderfailedauthentication":             SecuritySubmissionResultDetail_SenderFailedAuthentication,
		"simulatedthreat":                        SecuritySubmissionResultDetail_SimulatedThreat,
		"spoofblocked":                           SecuritySubmissionResultDetail_SpoofBlocked,
		"unabletomakedecision":                   SecuritySubmissionResultDetail_UnableToMakeDecision,
		"underinvestigation":                     SecuritySubmissionResultDetail_UnderInvestigation,
		"urlfilecannotmakedecision":              SecuritySubmissionResultDetail_UrlFileCannotMakeDecision,
		"urlfileshouldbeblocked":                 SecuritySubmissionResultDetail_UrlFileShouldBeBlocked,
		"urlfileshouldnotbeblocked":              SecuritySubmissionResultDetail_UrlFileShouldNotBeBlocked,
		"userimpersonation":                      SecuritySubmissionResultDetail_UserImpersonation,
		"willnotifyoncedone":                     SecuritySubmissionResultDetail_WillNotifyOnceDone,
		"zerohourautopurgeallowed":               SecuritySubmissionResultDetail_ZeroHourAutoPurgeAllowed,
		"zerohourautopurgeblocked":               SecuritySubmissionResultDetail_ZeroHourAutoPurgeBlocked,
		"zerohourautopurgequarantinereleased":    SecuritySubmissionResultDetail_ZeroHourAutoPurgeQuarantineReleased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionResultDetail(input)
	return &out, nil
}
