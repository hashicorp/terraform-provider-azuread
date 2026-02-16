package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingVideoDisabledReason string

const (
	OnlineMeetingVideoDisabledReason_WatermarkProtection OnlineMeetingVideoDisabledReason = "watermarkProtection"
)

func PossibleValuesForOnlineMeetingVideoDisabledReason() []string {
	return []string{
		string(OnlineMeetingVideoDisabledReason_WatermarkProtection),
	}
}

func (s *OnlineMeetingVideoDisabledReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingVideoDisabledReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingVideoDisabledReason(input string) (*OnlineMeetingVideoDisabledReason, error) {
	vals := map[string]OnlineMeetingVideoDisabledReason{
		"watermarkprotection": OnlineMeetingVideoDisabledReason_WatermarkProtection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingVideoDisabledReason(input)
	return &out, nil
}
