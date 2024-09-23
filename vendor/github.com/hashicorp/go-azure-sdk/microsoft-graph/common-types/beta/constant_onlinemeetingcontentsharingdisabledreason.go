package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingContentSharingDisabledReason string

const (
	OnlineMeetingContentSharingDisabledReason_WatermarkProtection OnlineMeetingContentSharingDisabledReason = "watermarkProtection"
)

func PossibleValuesForOnlineMeetingContentSharingDisabledReason() []string {
	return []string{
		string(OnlineMeetingContentSharingDisabledReason_WatermarkProtection),
	}
}

func (s *OnlineMeetingContentSharingDisabledReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingContentSharingDisabledReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingContentSharingDisabledReason(input string) (*OnlineMeetingContentSharingDisabledReason, error) {
	vals := map[string]OnlineMeetingContentSharingDisabledReason{
		"watermarkprotection": OnlineMeetingContentSharingDisabledReason_WatermarkProtection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingContentSharingDisabledReason(input)
	return &out, nil
}
