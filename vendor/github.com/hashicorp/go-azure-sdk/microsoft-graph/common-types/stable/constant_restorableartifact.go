package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableArtifact string

const (
	RestorableArtifact_Message RestorableArtifact = "message"
)

func PossibleValuesForRestorableArtifact() []string {
	return []string{
		string(RestorableArtifact_Message),
	}
}

func (s *RestorableArtifact) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestorableArtifact(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestorableArtifact(input string) (*RestorableArtifact, error) {
	vals := map[string]RestorableArtifact{
		"message": RestorableArtifact_Message,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestorableArtifact(input)
	return &out, nil
}
