package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationDisposition string

const (
	SynchronizationDisposition_Discard SynchronizationDisposition = "Discard"
	SynchronizationDisposition_Escrow  SynchronizationDisposition = "Escrow"
	SynchronizationDisposition_Normal  SynchronizationDisposition = "Normal"
)

func PossibleValuesForSynchronizationDisposition() []string {
	return []string{
		string(SynchronizationDisposition_Discard),
		string(SynchronizationDisposition_Escrow),
		string(SynchronizationDisposition_Normal),
	}
}

func (s *SynchronizationDisposition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationDisposition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationDisposition(input string) (*SynchronizationDisposition, error) {
	vals := map[string]SynchronizationDisposition{
		"discard": SynchronizationDisposition_Discard,
		"escrow":  SynchronizationDisposition_Escrow,
		"normal":  SynchronizationDisposition_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationDisposition(input)
	return &out, nil
}
