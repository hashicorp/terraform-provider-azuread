package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsWifiBand string

const (
	CallRecordsWifiBand_Frequency24GHz CallRecordsWifiBand = "frequency24GHz"
	CallRecordsWifiBand_Frequency50GHz CallRecordsWifiBand = "frequency50GHz"
	CallRecordsWifiBand_Frequency60GHz CallRecordsWifiBand = "frequency60GHz"
	CallRecordsWifiBand_Unknown        CallRecordsWifiBand = "unknown"
)

func PossibleValuesForCallRecordsWifiBand() []string {
	return []string{
		string(CallRecordsWifiBand_Frequency24GHz),
		string(CallRecordsWifiBand_Frequency50GHz),
		string(CallRecordsWifiBand_Frequency60GHz),
		string(CallRecordsWifiBand_Unknown),
	}
}

func (s *CallRecordsWifiBand) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsWifiBand(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsWifiBand(input string) (*CallRecordsWifiBand, error) {
	vals := map[string]CallRecordsWifiBand{
		"frequency24ghz": CallRecordsWifiBand_Frequency24GHz,
		"frequency50ghz": CallRecordsWifiBand_Frequency50GHz,
		"frequency60ghz": CallRecordsWifiBand_Frequency60GHz,
		"unknown":        CallRecordsWifiBand_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsWifiBand(input)
	return &out, nil
}
