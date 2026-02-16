package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppPublishingState string

const (
	MobileAppPublishingState_NotPublished MobileAppPublishingState = "notPublished"
	MobileAppPublishingState_Processing   MobileAppPublishingState = "processing"
	MobileAppPublishingState_Published    MobileAppPublishingState = "published"
)

func PossibleValuesForMobileAppPublishingState() []string {
	return []string{
		string(MobileAppPublishingState_NotPublished),
		string(MobileAppPublishingState_Processing),
		string(MobileAppPublishingState_Published),
	}
}

func (s *MobileAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppPublishingState(input string) (*MobileAppPublishingState, error) {
	vals := map[string]MobileAppPublishingState{
		"notpublished": MobileAppPublishingState_NotPublished,
		"processing":   MobileAppPublishingState_Processing,
		"published":    MobileAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppPublishingState(input)
	return &out, nil
}
