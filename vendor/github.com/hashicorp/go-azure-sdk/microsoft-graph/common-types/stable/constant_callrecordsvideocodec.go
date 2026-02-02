package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsVideoCodec string

const (
	CallRecordsVideoCodec_Av1     CallRecordsVideoCodec = "av1"
	CallRecordsVideoCodec_H263    CallRecordsVideoCodec = "h263"
	CallRecordsVideoCodec_H264    CallRecordsVideoCodec = "h264"
	CallRecordsVideoCodec_H264s   CallRecordsVideoCodec = "h264s"
	CallRecordsVideoCodec_H264uc  CallRecordsVideoCodec = "h264uc"
	CallRecordsVideoCodec_H265    CallRecordsVideoCodec = "h265"
	CallRecordsVideoCodec_Invalid CallRecordsVideoCodec = "invalid"
	CallRecordsVideoCodec_RtVideo CallRecordsVideoCodec = "rtVideo"
	CallRecordsVideoCodec_Rtvc1   CallRecordsVideoCodec = "rtvc1"
	CallRecordsVideoCodec_Unknown CallRecordsVideoCodec = "unknown"
	CallRecordsVideoCodec_Xrtvc1  CallRecordsVideoCodec = "xrtvc1"
)

func PossibleValuesForCallRecordsVideoCodec() []string {
	return []string{
		string(CallRecordsVideoCodec_Av1),
		string(CallRecordsVideoCodec_H263),
		string(CallRecordsVideoCodec_H264),
		string(CallRecordsVideoCodec_H264s),
		string(CallRecordsVideoCodec_H264uc),
		string(CallRecordsVideoCodec_H265),
		string(CallRecordsVideoCodec_Invalid),
		string(CallRecordsVideoCodec_RtVideo),
		string(CallRecordsVideoCodec_Rtvc1),
		string(CallRecordsVideoCodec_Unknown),
		string(CallRecordsVideoCodec_Xrtvc1),
	}
}

func (s *CallRecordsVideoCodec) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsVideoCodec(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsVideoCodec(input string) (*CallRecordsVideoCodec, error) {
	vals := map[string]CallRecordsVideoCodec{
		"av1":     CallRecordsVideoCodec_Av1,
		"h263":    CallRecordsVideoCodec_H263,
		"h264":    CallRecordsVideoCodec_H264,
		"h264s":   CallRecordsVideoCodec_H264s,
		"h264uc":  CallRecordsVideoCodec_H264uc,
		"h265":    CallRecordsVideoCodec_H265,
		"invalid": CallRecordsVideoCodec_Invalid,
		"rtvideo": CallRecordsVideoCodec_RtVideo,
		"rtvc1":   CallRecordsVideoCodec_Rtvc1,
		"unknown": CallRecordsVideoCodec_Unknown,
		"xrtvc1":  CallRecordsVideoCodec_Xrtvc1,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsVideoCodec(input)
	return &out, nil
}
