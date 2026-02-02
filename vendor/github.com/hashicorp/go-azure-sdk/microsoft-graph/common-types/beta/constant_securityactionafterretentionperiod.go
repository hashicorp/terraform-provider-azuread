package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityActionAfterRetentionPeriod string

const (
	SecurityActionAfterRetentionPeriod_Delete                 SecurityActionAfterRetentionPeriod = "delete"
	SecurityActionAfterRetentionPeriod_None                   SecurityActionAfterRetentionPeriod = "none"
	SecurityActionAfterRetentionPeriod_Relabel                SecurityActionAfterRetentionPeriod = "relabel"
	SecurityActionAfterRetentionPeriod_StartDispositionReview SecurityActionAfterRetentionPeriod = "startDispositionReview"
)

func PossibleValuesForSecurityActionAfterRetentionPeriod() []string {
	return []string{
		string(SecurityActionAfterRetentionPeriod_Delete),
		string(SecurityActionAfterRetentionPeriod_None),
		string(SecurityActionAfterRetentionPeriod_Relabel),
		string(SecurityActionAfterRetentionPeriod_StartDispositionReview),
	}
}

func (s *SecurityActionAfterRetentionPeriod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityActionAfterRetentionPeriod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityActionAfterRetentionPeriod(input string) (*SecurityActionAfterRetentionPeriod, error) {
	vals := map[string]SecurityActionAfterRetentionPeriod{
		"delete":                 SecurityActionAfterRetentionPeriod_Delete,
		"none":                   SecurityActionAfterRetentionPeriod_None,
		"relabel":                SecurityActionAfterRetentionPeriod_Relabel,
		"startdispositionreview": SecurityActionAfterRetentionPeriod_StartDispositionReview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityActionAfterRetentionPeriod(input)
	return &out, nil
}
