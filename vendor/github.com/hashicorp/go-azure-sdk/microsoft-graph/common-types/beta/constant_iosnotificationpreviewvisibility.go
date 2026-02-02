package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNotificationPreviewVisibility string

const (
	IosNotificationPreviewVisibility_AlwaysShow     IosNotificationPreviewVisibility = "alwaysShow"
	IosNotificationPreviewVisibility_HideWhenLocked IosNotificationPreviewVisibility = "hideWhenLocked"
	IosNotificationPreviewVisibility_NeverShow      IosNotificationPreviewVisibility = "neverShow"
	IosNotificationPreviewVisibility_NotConfigured  IosNotificationPreviewVisibility = "notConfigured"
)

func PossibleValuesForIosNotificationPreviewVisibility() []string {
	return []string{
		string(IosNotificationPreviewVisibility_AlwaysShow),
		string(IosNotificationPreviewVisibility_HideWhenLocked),
		string(IosNotificationPreviewVisibility_NeverShow),
		string(IosNotificationPreviewVisibility_NotConfigured),
	}
}

func (s *IosNotificationPreviewVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosNotificationPreviewVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosNotificationPreviewVisibility(input string) (*IosNotificationPreviewVisibility, error) {
	vals := map[string]IosNotificationPreviewVisibility{
		"alwaysshow":     IosNotificationPreviewVisibility_AlwaysShow,
		"hidewhenlocked": IosNotificationPreviewVisibility_HideWhenLocked,
		"nevershow":      IosNotificationPreviewVisibility_NeverShow,
		"notconfigured":  IosNotificationPreviewVisibility_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosNotificationPreviewVisibility(input)
	return &out, nil
}
