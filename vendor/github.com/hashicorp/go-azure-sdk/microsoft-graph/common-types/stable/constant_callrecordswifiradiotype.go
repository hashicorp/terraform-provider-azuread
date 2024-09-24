package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsWifiRadioType string

const (
	CallRecordsWifiRadioType_Unknown     CallRecordsWifiRadioType = "unknown"
	CallRecordsWifiRadioType_Wifi80211a  CallRecordsWifiRadioType = "wifi80211a"
	CallRecordsWifiRadioType_Wifi80211ac CallRecordsWifiRadioType = "wifi80211ac"
	CallRecordsWifiRadioType_Wifi80211ax CallRecordsWifiRadioType = "wifi80211ax"
	CallRecordsWifiRadioType_Wifi80211b  CallRecordsWifiRadioType = "wifi80211b"
	CallRecordsWifiRadioType_Wifi80211g  CallRecordsWifiRadioType = "wifi80211g"
	CallRecordsWifiRadioType_Wifi80211n  CallRecordsWifiRadioType = "wifi80211n"
)

func PossibleValuesForCallRecordsWifiRadioType() []string {
	return []string{
		string(CallRecordsWifiRadioType_Unknown),
		string(CallRecordsWifiRadioType_Wifi80211a),
		string(CallRecordsWifiRadioType_Wifi80211ac),
		string(CallRecordsWifiRadioType_Wifi80211ax),
		string(CallRecordsWifiRadioType_Wifi80211b),
		string(CallRecordsWifiRadioType_Wifi80211g),
		string(CallRecordsWifiRadioType_Wifi80211n),
	}
}

func (s *CallRecordsWifiRadioType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsWifiRadioType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsWifiRadioType(input string) (*CallRecordsWifiRadioType, error) {
	vals := map[string]CallRecordsWifiRadioType{
		"unknown":     CallRecordsWifiRadioType_Unknown,
		"wifi80211a":  CallRecordsWifiRadioType_Wifi80211a,
		"wifi80211ac": CallRecordsWifiRadioType_Wifi80211ac,
		"wifi80211ax": CallRecordsWifiRadioType_Wifi80211ax,
		"wifi80211b":  CallRecordsWifiRadioType_Wifi80211b,
		"wifi80211g":  CallRecordsWifiRadioType_Wifi80211g,
		"wifi80211n":  CallRecordsWifiRadioType_Wifi80211n,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsWifiRadioType(input)
	return &out, nil
}
