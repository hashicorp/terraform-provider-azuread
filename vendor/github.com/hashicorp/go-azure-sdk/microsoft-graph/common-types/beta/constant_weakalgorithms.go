package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WeakAlgorithms string

const (
	WeakAlgorithms_RsaSha1 WeakAlgorithms = "rsaSha1"
)

func PossibleValuesForWeakAlgorithms() []string {
	return []string{
		string(WeakAlgorithms_RsaSha1),
	}
}

func (s *WeakAlgorithms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWeakAlgorithms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWeakAlgorithms(input string) (*WeakAlgorithms, error) {
	vals := map[string]WeakAlgorithms{
		"rsasha1": WeakAlgorithms_RsaSha1,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WeakAlgorithms(input)
	return &out, nil
}
