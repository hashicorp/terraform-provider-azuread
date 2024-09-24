package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TlpLevel string

const (
	TlpLevel_Amber   TlpLevel = "amber"
	TlpLevel_Green   TlpLevel = "green"
	TlpLevel_Red     TlpLevel = "red"
	TlpLevel_Unknown TlpLevel = "unknown"
	TlpLevel_White   TlpLevel = "white"
)

func PossibleValuesForTlpLevel() []string {
	return []string{
		string(TlpLevel_Amber),
		string(TlpLevel_Green),
		string(TlpLevel_Red),
		string(TlpLevel_Unknown),
		string(TlpLevel_White),
	}
}

func (s *TlpLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTlpLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTlpLevel(input string) (*TlpLevel, error) {
	vals := map[string]TlpLevel{
		"amber":   TlpLevel_Amber,
		"green":   TlpLevel_Green,
		"red":     TlpLevel_Red,
		"unknown": TlpLevel_Unknown,
		"white":   TlpLevel_White,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TlpLevel(input)
	return &out, nil
}
