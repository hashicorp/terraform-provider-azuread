package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersistentBrowserSessionMode string

const (
	PersistentBrowserSessionMode_Always PersistentBrowserSessionMode = "always"
	PersistentBrowserSessionMode_Never  PersistentBrowserSessionMode = "never"
)

func PossibleValuesForPersistentBrowserSessionMode() []string {
	return []string{
		string(PersistentBrowserSessionMode_Always),
		string(PersistentBrowserSessionMode_Never),
	}
}

func (s *PersistentBrowserSessionMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersistentBrowserSessionMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersistentBrowserSessionMode(input string) (*PersistentBrowserSessionMode, error) {
	vals := map[string]PersistentBrowserSessionMode{
		"always": PersistentBrowserSessionMode_Always,
		"never":  PersistentBrowserSessionMode_Never,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersistentBrowserSessionMode(input)
	return &out, nil
}
