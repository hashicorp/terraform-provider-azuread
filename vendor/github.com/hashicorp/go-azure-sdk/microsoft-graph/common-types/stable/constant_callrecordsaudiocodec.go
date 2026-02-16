package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsAudioCodec string

const (
	CallRecordsAudioCodec_AmrWide           CallRecordsAudioCodec = "amrWide"
	CallRecordsAudioCodec_Cn                CallRecordsAudioCodec = "cn"
	CallRecordsAudioCodec_G722              CallRecordsAudioCodec = "g722"
	CallRecordsAudioCodec_G7221             CallRecordsAudioCodec = "g7221"
	CallRecordsAudioCodec_G7221c            CallRecordsAudioCodec = "g7221c"
	CallRecordsAudioCodec_G729              CallRecordsAudioCodec = "g729"
	CallRecordsAudioCodec_Invalid           CallRecordsAudioCodec = "invalid"
	CallRecordsAudioCodec_Muchv2            CallRecordsAudioCodec = "muchv2"
	CallRecordsAudioCodec_MultiChannelAudio CallRecordsAudioCodec = "multiChannelAudio"
	CallRecordsAudioCodec_Opus              CallRecordsAudioCodec = "opus"
	CallRecordsAudioCodec_Pcma              CallRecordsAudioCodec = "pcma"
	CallRecordsAudioCodec_Pcmu              CallRecordsAudioCodec = "pcmu"
	CallRecordsAudioCodec_RtAudio16         CallRecordsAudioCodec = "rtAudio16"
	CallRecordsAudioCodec_RtAudio8          CallRecordsAudioCodec = "rtAudio8"
	CallRecordsAudioCodec_Satin             CallRecordsAudioCodec = "satin"
	CallRecordsAudioCodec_SatinFullband     CallRecordsAudioCodec = "satinFullband"
	CallRecordsAudioCodec_Silk              CallRecordsAudioCodec = "silk"
	CallRecordsAudioCodec_SilkNarrow        CallRecordsAudioCodec = "silkNarrow"
	CallRecordsAudioCodec_SilkWide          CallRecordsAudioCodec = "silkWide"
	CallRecordsAudioCodec_Siren             CallRecordsAudioCodec = "siren"
	CallRecordsAudioCodec_Unknown           CallRecordsAudioCodec = "unknown"
	CallRecordsAudioCodec_XmsRta            CallRecordsAudioCodec = "xmsRta"
)

func PossibleValuesForCallRecordsAudioCodec() []string {
	return []string{
		string(CallRecordsAudioCodec_AmrWide),
		string(CallRecordsAudioCodec_Cn),
		string(CallRecordsAudioCodec_G722),
		string(CallRecordsAudioCodec_G7221),
		string(CallRecordsAudioCodec_G7221c),
		string(CallRecordsAudioCodec_G729),
		string(CallRecordsAudioCodec_Invalid),
		string(CallRecordsAudioCodec_Muchv2),
		string(CallRecordsAudioCodec_MultiChannelAudio),
		string(CallRecordsAudioCodec_Opus),
		string(CallRecordsAudioCodec_Pcma),
		string(CallRecordsAudioCodec_Pcmu),
		string(CallRecordsAudioCodec_RtAudio16),
		string(CallRecordsAudioCodec_RtAudio8),
		string(CallRecordsAudioCodec_Satin),
		string(CallRecordsAudioCodec_SatinFullband),
		string(CallRecordsAudioCodec_Silk),
		string(CallRecordsAudioCodec_SilkNarrow),
		string(CallRecordsAudioCodec_SilkWide),
		string(CallRecordsAudioCodec_Siren),
		string(CallRecordsAudioCodec_Unknown),
		string(CallRecordsAudioCodec_XmsRta),
	}
}

func (s *CallRecordsAudioCodec) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsAudioCodec(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsAudioCodec(input string) (*CallRecordsAudioCodec, error) {
	vals := map[string]CallRecordsAudioCodec{
		"amrwide":           CallRecordsAudioCodec_AmrWide,
		"cn":                CallRecordsAudioCodec_Cn,
		"g722":              CallRecordsAudioCodec_G722,
		"g7221":             CallRecordsAudioCodec_G7221,
		"g7221c":            CallRecordsAudioCodec_G7221c,
		"g729":              CallRecordsAudioCodec_G729,
		"invalid":           CallRecordsAudioCodec_Invalid,
		"muchv2":            CallRecordsAudioCodec_Muchv2,
		"multichannelaudio": CallRecordsAudioCodec_MultiChannelAudio,
		"opus":              CallRecordsAudioCodec_Opus,
		"pcma":              CallRecordsAudioCodec_Pcma,
		"pcmu":              CallRecordsAudioCodec_Pcmu,
		"rtaudio16":         CallRecordsAudioCodec_RtAudio16,
		"rtaudio8":          CallRecordsAudioCodec_RtAudio8,
		"satin":             CallRecordsAudioCodec_Satin,
		"satinfullband":     CallRecordsAudioCodec_SatinFullband,
		"silk":              CallRecordsAudioCodec_Silk,
		"silknarrow":        CallRecordsAudioCodec_SilkNarrow,
		"silkwide":          CallRecordsAudioCodec_SilkWide,
		"siren":             CallRecordsAudioCodec_Siren,
		"unknown":           CallRecordsAudioCodec_Unknown,
		"xmsrta":            CallRecordsAudioCodec_XmsRta,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsAudioCodec(input)
	return &out, nil
}
