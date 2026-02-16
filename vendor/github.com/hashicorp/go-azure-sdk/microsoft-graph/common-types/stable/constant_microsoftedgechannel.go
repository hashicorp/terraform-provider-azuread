package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftEdgeChannel string

const (
	MicrosoftEdgeChannel_Beta   MicrosoftEdgeChannel = "beta"
	MicrosoftEdgeChannel_Dev    MicrosoftEdgeChannel = "dev"
	MicrosoftEdgeChannel_Stable MicrosoftEdgeChannel = "stable"
)

func PossibleValuesForMicrosoftEdgeChannel() []string {
	return []string{
		string(MicrosoftEdgeChannel_Beta),
		string(MicrosoftEdgeChannel_Dev),
		string(MicrosoftEdgeChannel_Stable),
	}
}

func (s *MicrosoftEdgeChannel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftEdgeChannel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftEdgeChannel(input string) (*MicrosoftEdgeChannel, error) {
	vals := map[string]MicrosoftEdgeChannel{
		"beta":   MicrosoftEdgeChannel_Beta,
		"dev":    MicrosoftEdgeChannel_Dev,
		"stable": MicrosoftEdgeChannel_Stable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftEdgeChannel(input)
	return &out, nil
}
