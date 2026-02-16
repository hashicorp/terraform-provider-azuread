package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacAddressRandomizationMode string

const (
	MacAddressRandomizationMode_Automatic MacAddressRandomizationMode = "automatic"
	MacAddressRandomizationMode_Hardware  MacAddressRandomizationMode = "hardware"
)

func PossibleValuesForMacAddressRandomizationMode() []string {
	return []string{
		string(MacAddressRandomizationMode_Automatic),
		string(MacAddressRandomizationMode_Hardware),
	}
}

func (s *MacAddressRandomizationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacAddressRandomizationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacAddressRandomizationMode(input string) (*MacAddressRandomizationMode, error) {
	vals := map[string]MacAddressRandomizationMode{
		"automatic": MacAddressRandomizationMode_Automatic,
		"hardware":  MacAddressRandomizationMode_Hardware,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacAddressRandomizationMode(input)
	return &out, nil
}
