package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRemoteNetworkStatus string

const (
	NetworkaccessRemoteNetworkStatus_BgpConnected       NetworkaccessRemoteNetworkStatus = "bgpConnected"
	NetworkaccessRemoteNetworkStatus_BgpDisconnected    NetworkaccessRemoteNetworkStatus = "bgpDisconnected"
	NetworkaccessRemoteNetworkStatus_RemoteNetworkAlive NetworkaccessRemoteNetworkStatus = "remoteNetworkAlive"
	NetworkaccessRemoteNetworkStatus_TunnelConnected    NetworkaccessRemoteNetworkStatus = "tunnelConnected"
	NetworkaccessRemoteNetworkStatus_TunnelDisconnected NetworkaccessRemoteNetworkStatus = "tunnelDisconnected"
)

func PossibleValuesForNetworkaccessRemoteNetworkStatus() []string {
	return []string{
		string(NetworkaccessRemoteNetworkStatus_BgpConnected),
		string(NetworkaccessRemoteNetworkStatus_BgpDisconnected),
		string(NetworkaccessRemoteNetworkStatus_RemoteNetworkAlive),
		string(NetworkaccessRemoteNetworkStatus_TunnelConnected),
		string(NetworkaccessRemoteNetworkStatus_TunnelDisconnected),
	}
}

func (s *NetworkaccessRemoteNetworkStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessRemoteNetworkStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessRemoteNetworkStatus(input string) (*NetworkaccessRemoteNetworkStatus, error) {
	vals := map[string]NetworkaccessRemoteNetworkStatus{
		"bgpconnected":       NetworkaccessRemoteNetworkStatus_BgpConnected,
		"bgpdisconnected":    NetworkaccessRemoteNetworkStatus_BgpDisconnected,
		"remotenetworkalive": NetworkaccessRemoteNetworkStatus_RemoteNetworkAlive,
		"tunnelconnected":    NetworkaccessRemoteNetworkStatus_TunnelConnected,
		"tunneldisconnected": NetworkaccessRemoteNetworkStatus_TunnelDisconnected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessRemoteNetworkStatus(input)
	return &out, nil
}
