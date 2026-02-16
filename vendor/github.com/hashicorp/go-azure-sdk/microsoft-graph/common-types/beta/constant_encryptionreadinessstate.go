package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionReadinessState string

const (
	EncryptionReadinessState_NotReady EncryptionReadinessState = "notReady"
	EncryptionReadinessState_Ready    EncryptionReadinessState = "ready"
)

func PossibleValuesForEncryptionReadinessState() []string {
	return []string{
		string(EncryptionReadinessState_NotReady),
		string(EncryptionReadinessState_Ready),
	}
}

func (s *EncryptionReadinessState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptionReadinessState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptionReadinessState(input string) (*EncryptionReadinessState, error) {
	vals := map[string]EncryptionReadinessState{
		"notready": EncryptionReadinessState_NotReady,
		"ready":    EncryptionReadinessState_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionReadinessState(input)
	return &out, nil
}
