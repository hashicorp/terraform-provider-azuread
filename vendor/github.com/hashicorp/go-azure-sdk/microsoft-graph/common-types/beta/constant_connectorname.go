package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorName string

const (
	ConnectorName_AppleDepExpirationDateTime                        ConnectorName = "appleDepExpirationDateTime"
	ConnectorName_AppleDepLastSyncDateTime                          ConnectorName = "appleDepLastSyncDateTime"
	ConnectorName_ApplePushNotificationServiceExpirationDateTime    ConnectorName = "applePushNotificationServiceExpirationDateTime"
	ConnectorName_ChromebookLastDirectorySyncDateTime               ConnectorName = "chromebookLastDirectorySyncDateTime"
	ConnectorName_FutureValue                                       ConnectorName = "futureValue"
	ConnectorName_GooglePlayAppLastSyncDateTime                     ConnectorName = "googlePlayAppLastSyncDateTime"
	ConnectorName_GooglePlayConnectorLastModifiedDateTime           ConnectorName = "googlePlayConnectorLastModifiedDateTime"
	ConnectorName_JamfLastSyncDateTime                              ConnectorName = "jamfLastSyncDateTime"
	ConnectorName_MobileThreatDefenceConnectorLastHeartbeatDateTime ConnectorName = "mobileThreatDefenceConnectorLastHeartbeatDateTime"
	ConnectorName_NdesConnectorLastConnectionDateTime               ConnectorName = "ndesConnectorLastConnectionDateTime"
	ConnectorName_OnPremConnectorLastSyncDateTime                   ConnectorName = "onPremConnectorLastSyncDateTime"
	ConnectorName_VppTokenExpirationDateTime                        ConnectorName = "vppTokenExpirationDateTime"
	ConnectorName_VppTokenLastSyncDateTime                          ConnectorName = "vppTokenLastSyncDateTime"
	ConnectorName_WindowsAutopilotLastSyncDateTime                  ConnectorName = "windowsAutopilotLastSyncDateTime"
	ConnectorName_WindowsDefenderATPConnectorLastHeartbeatDateTime  ConnectorName = "windowsDefenderATPConnectorLastHeartbeatDateTime"
	ConnectorName_WindowsStoreForBusinessLastSyncDateTime           ConnectorName = "windowsStoreForBusinessLastSyncDateTime"
)

func PossibleValuesForConnectorName() []string {
	return []string{
		string(ConnectorName_AppleDepExpirationDateTime),
		string(ConnectorName_AppleDepLastSyncDateTime),
		string(ConnectorName_ApplePushNotificationServiceExpirationDateTime),
		string(ConnectorName_ChromebookLastDirectorySyncDateTime),
		string(ConnectorName_FutureValue),
		string(ConnectorName_GooglePlayAppLastSyncDateTime),
		string(ConnectorName_GooglePlayConnectorLastModifiedDateTime),
		string(ConnectorName_JamfLastSyncDateTime),
		string(ConnectorName_MobileThreatDefenceConnectorLastHeartbeatDateTime),
		string(ConnectorName_NdesConnectorLastConnectionDateTime),
		string(ConnectorName_OnPremConnectorLastSyncDateTime),
		string(ConnectorName_VppTokenExpirationDateTime),
		string(ConnectorName_VppTokenLastSyncDateTime),
		string(ConnectorName_WindowsAutopilotLastSyncDateTime),
		string(ConnectorName_WindowsDefenderATPConnectorLastHeartbeatDateTime),
		string(ConnectorName_WindowsStoreForBusinessLastSyncDateTime),
	}
}

func (s *ConnectorName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectorName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectorName(input string) (*ConnectorName, error) {
	vals := map[string]ConnectorName{
		"appledepexpirationdatetime":                        ConnectorName_AppleDepExpirationDateTime,
		"appledeplastsyncdatetime":                          ConnectorName_AppleDepLastSyncDateTime,
		"applepushnotificationserviceexpirationdatetime":    ConnectorName_ApplePushNotificationServiceExpirationDateTime,
		"chromebooklastdirectorysyncdatetime":               ConnectorName_ChromebookLastDirectorySyncDateTime,
		"futurevalue":                                       ConnectorName_FutureValue,
		"googleplayapplastsyncdatetime":                     ConnectorName_GooglePlayAppLastSyncDateTime,
		"googleplayconnectorlastmodifieddatetime":           ConnectorName_GooglePlayConnectorLastModifiedDateTime,
		"jamflastsyncdatetime":                              ConnectorName_JamfLastSyncDateTime,
		"mobilethreatdefenceconnectorlastheartbeatdatetime": ConnectorName_MobileThreatDefenceConnectorLastHeartbeatDateTime,
		"ndesconnectorlastconnectiondatetime":               ConnectorName_NdesConnectorLastConnectionDateTime,
		"onpremconnectorlastsyncdatetime":                   ConnectorName_OnPremConnectorLastSyncDateTime,
		"vpptokenexpirationdatetime":                        ConnectorName_VppTokenExpirationDateTime,
		"vpptokenlastsyncdatetime":                          ConnectorName_VppTokenLastSyncDateTime,
		"windowsautopilotlastsyncdatetime":                  ConnectorName_WindowsAutopilotLastSyncDateTime,
		"windowsdefenderatpconnectorlastheartbeatdatetime":  ConnectorName_WindowsDefenderATPConnectorLastHeartbeatDateTime,
		"windowsstoreforbusinesslastsyncdatetime":           ConnectorName_WindowsStoreForBusinessLastSyncDateTime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorName(input)
	return &out, nil
}
