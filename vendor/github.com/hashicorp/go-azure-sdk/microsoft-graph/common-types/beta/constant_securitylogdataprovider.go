package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityLogDataProvider string

const (
	SecurityLogDataProvider_Barracuda                     SecurityLogDataProvider = "barracuda"
	SecurityLogDataProvider_BarracudaNextGenFw            SecurityLogDataProvider = "barracudaNextGenFw"
	SecurityLogDataProvider_BarracudaNextGenFwWeblog      SecurityLogDataProvider = "barracudaNextGenFwWeblog"
	SecurityLogDataProvider_Bluecoat                      SecurityLogDataProvider = "bluecoat"
	SecurityLogDataProvider_Checkpoint                    SecurityLogDataProvider = "checkpoint"
	SecurityLogDataProvider_CheckpointCef                 SecurityLogDataProvider = "checkpointCef"
	SecurityLogDataProvider_CheckpointSmartViewTracker    SecurityLogDataProvider = "checkpointSmartViewTracker"
	SecurityLogDataProvider_CheckpointXml                 SecurityLogDataProvider = "checkpointXml"
	SecurityLogDataProvider_CiscoAsa                      SecurityLogDataProvider = "ciscoAsa"
	SecurityLogDataProvider_CiscoAsaFirepower             SecurityLogDataProvider = "ciscoAsaFirepower"
	SecurityLogDataProvider_CiscoFirepowerV6              SecurityLogDataProvider = "ciscoFirepowerV6"
	SecurityLogDataProvider_CiscoFwsm                     SecurityLogDataProvider = "ciscoFwsm"
	SecurityLogDataProvider_CiscoIronportProxy            SecurityLogDataProvider = "ciscoIronportProxy"
	SecurityLogDataProvider_CiscoIronportWsaIi            SecurityLogDataProvider = "ciscoIronportWsaIi"
	SecurityLogDataProvider_CiscoIronportWsaIii           SecurityLogDataProvider = "ciscoIronportWsaIii"
	SecurityLogDataProvider_CiscoScanSafe                 SecurityLogDataProvider = "ciscoScanSafe"
	SecurityLogDataProvider_Clavister                     SecurityLogDataProvider = "clavister"
	SecurityLogDataProvider_Contentkeeper                 SecurityLogDataProvider = "contentkeeper"
	SecurityLogDataProvider_Corrata                       SecurityLogDataProvider = "corrata"
	SecurityLogDataProvider_CustomParser                  SecurityLogDataProvider = "customParser"
	SecurityLogDataProvider_Forcepoint                    SecurityLogDataProvider = "forcepoint"
	SecurityLogDataProvider_ForcepointLeef                SecurityLogDataProvider = "forcepointLeef"
	SecurityLogDataProvider_Fortigate                     SecurityLogDataProvider = "fortigate"
	SecurityLogDataProvider_Fortios                       SecurityLogDataProvider = "fortios"
	SecurityLogDataProvider_GenericCef                    SecurityLogDataProvider = "genericCef"
	SecurityLogDataProvider_GenericLeef                   SecurityLogDataProvider = "genericLeef"
	SecurityLogDataProvider_GenericW3C                    SecurityLogDataProvider = "genericW3C"
	SecurityLogDataProvider_IFilter                       SecurityLogDataProvider = "iFilter"
	SecurityLogDataProvider_Iboss                         SecurityLogDataProvider = "iboss"
	SecurityLogDataProvider_JuniperSrx                    SecurityLogDataProvider = "juniperSrx"
	SecurityLogDataProvider_JuniperSrxSd                  SecurityLogDataProvider = "juniperSrxSd"
	SecurityLogDataProvider_JuniperSrxWelf                SecurityLogDataProvider = "juniperSrxWelf"
	SecurityLogDataProvider_JuniperSsg                    SecurityLogDataProvider = "juniperSsg"
	SecurityLogDataProvider_MachineZoneMeraki             SecurityLogDataProvider = "machineZoneMeraki"
	SecurityLogDataProvider_McafeeSwg                     SecurityLogDataProvider = "mcafeeSwg"
	SecurityLogDataProvider_MenloSecurityCef              SecurityLogDataProvider = "menloSecurityCef"
	SecurityLogDataProvider_MicrosoftConditionalAppAccess SecurityLogDataProvider = "microsoftConditionalAppAccess"
	SecurityLogDataProvider_MicrosoftDefenderForEndpoint  SecurityLogDataProvider = "microsoftDefenderForEndpoint"
	SecurityLogDataProvider_MicrosoftIsaW3C               SecurityLogDataProvider = "microsoftIsaW3C"
	SecurityLogDataProvider_OpenSystemsSecureWebGateway   SecurityLogDataProvider = "openSystemsSecureWebGateway"
	SecurityLogDataProvider_PaloAlto                      SecurityLogDataProvider = "paloAlto"
	SecurityLogDataProvider_PaloAltoLeef                  SecurityLogDataProvider = "paloAltoLeef"
	SecurityLogDataProvider_Sonicwall                     SecurityLogDataProvider = "sonicwall"
	SecurityLogDataProvider_SophosCyberoam                SecurityLogDataProvider = "sophosCyberoam"
	SecurityLogDataProvider_SophosSg                      SecurityLogDataProvider = "sophosSg"
	SecurityLogDataProvider_SophosXg                      SecurityLogDataProvider = "sophosXg"
	SecurityLogDataProvider_Squid                         SecurityLogDataProvider = "squid"
	SecurityLogDataProvider_SquidNative                   SecurityLogDataProvider = "squidNative"
	SecurityLogDataProvider_Stormshield                   SecurityLogDataProvider = "stormshield"
	SecurityLogDataProvider_Wandera                       SecurityLogDataProvider = "wandera"
	SecurityLogDataProvider_WatchguardXtm                 SecurityLogDataProvider = "watchguardXtm"
	SecurityLogDataProvider_WebsenseSiemCef               SecurityLogDataProvider = "websenseSiemCef"
	SecurityLogDataProvider_WebsenseV75                   SecurityLogDataProvider = "websenseV75"
	SecurityLogDataProvider_Zscaler                       SecurityLogDataProvider = "zscaler"
	SecurityLogDataProvider_ZscalerCef                    SecurityLogDataProvider = "zscalerCef"
	SecurityLogDataProvider_ZscalerQradar                 SecurityLogDataProvider = "zscalerQradar"
)

func PossibleValuesForSecurityLogDataProvider() []string {
	return []string{
		string(SecurityLogDataProvider_Barracuda),
		string(SecurityLogDataProvider_BarracudaNextGenFw),
		string(SecurityLogDataProvider_BarracudaNextGenFwWeblog),
		string(SecurityLogDataProvider_Bluecoat),
		string(SecurityLogDataProvider_Checkpoint),
		string(SecurityLogDataProvider_CheckpointCef),
		string(SecurityLogDataProvider_CheckpointSmartViewTracker),
		string(SecurityLogDataProvider_CheckpointXml),
		string(SecurityLogDataProvider_CiscoAsa),
		string(SecurityLogDataProvider_CiscoAsaFirepower),
		string(SecurityLogDataProvider_CiscoFirepowerV6),
		string(SecurityLogDataProvider_CiscoFwsm),
		string(SecurityLogDataProvider_CiscoIronportProxy),
		string(SecurityLogDataProvider_CiscoIronportWsaIi),
		string(SecurityLogDataProvider_CiscoIronportWsaIii),
		string(SecurityLogDataProvider_CiscoScanSafe),
		string(SecurityLogDataProvider_Clavister),
		string(SecurityLogDataProvider_Contentkeeper),
		string(SecurityLogDataProvider_Corrata),
		string(SecurityLogDataProvider_CustomParser),
		string(SecurityLogDataProvider_Forcepoint),
		string(SecurityLogDataProvider_ForcepointLeef),
		string(SecurityLogDataProvider_Fortigate),
		string(SecurityLogDataProvider_Fortios),
		string(SecurityLogDataProvider_GenericCef),
		string(SecurityLogDataProvider_GenericLeef),
		string(SecurityLogDataProvider_GenericW3C),
		string(SecurityLogDataProvider_IFilter),
		string(SecurityLogDataProvider_Iboss),
		string(SecurityLogDataProvider_JuniperSrx),
		string(SecurityLogDataProvider_JuniperSrxSd),
		string(SecurityLogDataProvider_JuniperSrxWelf),
		string(SecurityLogDataProvider_JuniperSsg),
		string(SecurityLogDataProvider_MachineZoneMeraki),
		string(SecurityLogDataProvider_McafeeSwg),
		string(SecurityLogDataProvider_MenloSecurityCef),
		string(SecurityLogDataProvider_MicrosoftConditionalAppAccess),
		string(SecurityLogDataProvider_MicrosoftDefenderForEndpoint),
		string(SecurityLogDataProvider_MicrosoftIsaW3C),
		string(SecurityLogDataProvider_OpenSystemsSecureWebGateway),
		string(SecurityLogDataProvider_PaloAlto),
		string(SecurityLogDataProvider_PaloAltoLeef),
		string(SecurityLogDataProvider_Sonicwall),
		string(SecurityLogDataProvider_SophosCyberoam),
		string(SecurityLogDataProvider_SophosSg),
		string(SecurityLogDataProvider_SophosXg),
		string(SecurityLogDataProvider_Squid),
		string(SecurityLogDataProvider_SquidNative),
		string(SecurityLogDataProvider_Stormshield),
		string(SecurityLogDataProvider_Wandera),
		string(SecurityLogDataProvider_WatchguardXtm),
		string(SecurityLogDataProvider_WebsenseSiemCef),
		string(SecurityLogDataProvider_WebsenseV75),
		string(SecurityLogDataProvider_Zscaler),
		string(SecurityLogDataProvider_ZscalerCef),
		string(SecurityLogDataProvider_ZscalerQradar),
	}
}

func (s *SecurityLogDataProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityLogDataProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityLogDataProvider(input string) (*SecurityLogDataProvider, error) {
	vals := map[string]SecurityLogDataProvider{
		"barracuda":                     SecurityLogDataProvider_Barracuda,
		"barracudanextgenfw":            SecurityLogDataProvider_BarracudaNextGenFw,
		"barracudanextgenfwweblog":      SecurityLogDataProvider_BarracudaNextGenFwWeblog,
		"bluecoat":                      SecurityLogDataProvider_Bluecoat,
		"checkpoint":                    SecurityLogDataProvider_Checkpoint,
		"checkpointcef":                 SecurityLogDataProvider_CheckpointCef,
		"checkpointsmartviewtracker":    SecurityLogDataProvider_CheckpointSmartViewTracker,
		"checkpointxml":                 SecurityLogDataProvider_CheckpointXml,
		"ciscoasa":                      SecurityLogDataProvider_CiscoAsa,
		"ciscoasafirepower":             SecurityLogDataProvider_CiscoAsaFirepower,
		"ciscofirepowerv6":              SecurityLogDataProvider_CiscoFirepowerV6,
		"ciscofwsm":                     SecurityLogDataProvider_CiscoFwsm,
		"ciscoironportproxy":            SecurityLogDataProvider_CiscoIronportProxy,
		"ciscoironportwsaii":            SecurityLogDataProvider_CiscoIronportWsaIi,
		"ciscoironportwsaiii":           SecurityLogDataProvider_CiscoIronportWsaIii,
		"ciscoscansafe":                 SecurityLogDataProvider_CiscoScanSafe,
		"clavister":                     SecurityLogDataProvider_Clavister,
		"contentkeeper":                 SecurityLogDataProvider_Contentkeeper,
		"corrata":                       SecurityLogDataProvider_Corrata,
		"customparser":                  SecurityLogDataProvider_CustomParser,
		"forcepoint":                    SecurityLogDataProvider_Forcepoint,
		"forcepointleef":                SecurityLogDataProvider_ForcepointLeef,
		"fortigate":                     SecurityLogDataProvider_Fortigate,
		"fortios":                       SecurityLogDataProvider_Fortios,
		"genericcef":                    SecurityLogDataProvider_GenericCef,
		"genericleef":                   SecurityLogDataProvider_GenericLeef,
		"genericw3c":                    SecurityLogDataProvider_GenericW3C,
		"ifilter":                       SecurityLogDataProvider_IFilter,
		"iboss":                         SecurityLogDataProvider_Iboss,
		"junipersrx":                    SecurityLogDataProvider_JuniperSrx,
		"junipersrxsd":                  SecurityLogDataProvider_JuniperSrxSd,
		"junipersrxwelf":                SecurityLogDataProvider_JuniperSrxWelf,
		"juniperssg":                    SecurityLogDataProvider_JuniperSsg,
		"machinezonemeraki":             SecurityLogDataProvider_MachineZoneMeraki,
		"mcafeeswg":                     SecurityLogDataProvider_McafeeSwg,
		"menlosecuritycef":              SecurityLogDataProvider_MenloSecurityCef,
		"microsoftconditionalappaccess": SecurityLogDataProvider_MicrosoftConditionalAppAccess,
		"microsoftdefenderforendpoint":  SecurityLogDataProvider_MicrosoftDefenderForEndpoint,
		"microsoftisaw3c":               SecurityLogDataProvider_MicrosoftIsaW3C,
		"opensystemssecurewebgateway":   SecurityLogDataProvider_OpenSystemsSecureWebGateway,
		"paloalto":                      SecurityLogDataProvider_PaloAlto,
		"paloaltoleef":                  SecurityLogDataProvider_PaloAltoLeef,
		"sonicwall":                     SecurityLogDataProvider_Sonicwall,
		"sophoscyberoam":                SecurityLogDataProvider_SophosCyberoam,
		"sophossg":                      SecurityLogDataProvider_SophosSg,
		"sophosxg":                      SecurityLogDataProvider_SophosXg,
		"squid":                         SecurityLogDataProvider_Squid,
		"squidnative":                   SecurityLogDataProvider_SquidNative,
		"stormshield":                   SecurityLogDataProvider_Stormshield,
		"wandera":                       SecurityLogDataProvider_Wandera,
		"watchguardxtm":                 SecurityLogDataProvider_WatchguardXtm,
		"websensesiemcef":               SecurityLogDataProvider_WebsenseSiemCef,
		"websensev75":                   SecurityLogDataProvider_WebsenseV75,
		"zscaler":                       SecurityLogDataProvider_Zscaler,
		"zscalercef":                    SecurityLogDataProvider_ZscalerCef,
		"zscalerqradar":                 SecurityLogDataProvider_ZscalerQradar,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityLogDataProvider(input)
	return &out, nil
}
