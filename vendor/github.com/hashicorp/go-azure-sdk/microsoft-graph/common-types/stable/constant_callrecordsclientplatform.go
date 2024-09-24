package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsClientPlatform string

const (
	CallRecordsClientPlatform_Android    CallRecordsClientPlatform = "android"
	CallRecordsClientPlatform_HoloLens   CallRecordsClientPlatform = "holoLens"
	CallRecordsClientPlatform_IOS        CallRecordsClientPlatform = "iOS"
	CallRecordsClientPlatform_IPPhone    CallRecordsClientPlatform = "ipPhone"
	CallRecordsClientPlatform_MacOS      CallRecordsClientPlatform = "macOS"
	CallRecordsClientPlatform_RoomSystem CallRecordsClientPlatform = "roomSystem"
	CallRecordsClientPlatform_SurfaceHub CallRecordsClientPlatform = "surfaceHub"
	CallRecordsClientPlatform_Unknown    CallRecordsClientPlatform = "unknown"
	CallRecordsClientPlatform_Web        CallRecordsClientPlatform = "web"
	CallRecordsClientPlatform_Windows    CallRecordsClientPlatform = "windows"
)

func PossibleValuesForCallRecordsClientPlatform() []string {
	return []string{
		string(CallRecordsClientPlatform_Android),
		string(CallRecordsClientPlatform_HoloLens),
		string(CallRecordsClientPlatform_IOS),
		string(CallRecordsClientPlatform_IPPhone),
		string(CallRecordsClientPlatform_MacOS),
		string(CallRecordsClientPlatform_RoomSystem),
		string(CallRecordsClientPlatform_SurfaceHub),
		string(CallRecordsClientPlatform_Unknown),
		string(CallRecordsClientPlatform_Web),
		string(CallRecordsClientPlatform_Windows),
	}
}

func (s *CallRecordsClientPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsClientPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsClientPlatform(input string) (*CallRecordsClientPlatform, error) {
	vals := map[string]CallRecordsClientPlatform{
		"android":    CallRecordsClientPlatform_Android,
		"hololens":   CallRecordsClientPlatform_HoloLens,
		"ios":        CallRecordsClientPlatform_IOS,
		"ipphone":    CallRecordsClientPlatform_IPPhone,
		"macos":      CallRecordsClientPlatform_MacOS,
		"roomsystem": CallRecordsClientPlatform_RoomSystem,
		"surfacehub": CallRecordsClientPlatform_SurfaceHub,
		"unknown":    CallRecordsClientPlatform_Unknown,
		"web":        CallRecordsClientPlatform_Web,
		"windows":    CallRecordsClientPlatform_Windows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsClientPlatform(input)
	return &out, nil
}
