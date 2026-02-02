package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaNetworkType string

const (
	ZebraFotaNetworkType_Any             ZebraFotaNetworkType = "any"
	ZebraFotaNetworkType_Cellular        ZebraFotaNetworkType = "cellular"
	ZebraFotaNetworkType_Wifi            ZebraFotaNetworkType = "wifi"
	ZebraFotaNetworkType_WifiAndCellular ZebraFotaNetworkType = "wifiAndCellular"
)

func PossibleValuesForZebraFotaNetworkType() []string {
	return []string{
		string(ZebraFotaNetworkType_Any),
		string(ZebraFotaNetworkType_Cellular),
		string(ZebraFotaNetworkType_Wifi),
		string(ZebraFotaNetworkType_WifiAndCellular),
	}
}

func (s *ZebraFotaNetworkType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZebraFotaNetworkType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZebraFotaNetworkType(input string) (*ZebraFotaNetworkType, error) {
	vals := map[string]ZebraFotaNetworkType{
		"any":             ZebraFotaNetworkType_Any,
		"cellular":        ZebraFotaNetworkType_Cellular,
		"wifi":            ZebraFotaNetworkType_Wifi,
		"wifiandcellular": ZebraFotaNetworkType_WifiAndCellular,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaNetworkType(input)
	return &out, nil
}
