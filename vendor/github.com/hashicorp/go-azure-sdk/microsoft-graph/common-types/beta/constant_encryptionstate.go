package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionState string

const (
	EncryptionState_Encrypted    EncryptionState = "encrypted"
	EncryptionState_NotEncrypted EncryptionState = "notEncrypted"
)

func PossibleValuesForEncryptionState() []string {
	return []string{
		string(EncryptionState_Encrypted),
		string(EncryptionState_NotEncrypted),
	}
}

func (s *EncryptionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptionState(input string) (*EncryptionState, error) {
	vals := map[string]EncryptionState{
		"encrypted":    EncryptionState_Encrypted,
		"notencrypted": EncryptionState_NotEncrypted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionState(input)
	return &out, nil
}
