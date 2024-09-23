package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaScheduleMode string

const (
	ZebraFotaScheduleMode_InstallNow ZebraFotaScheduleMode = "installNow"
	ZebraFotaScheduleMode_Scheduled  ZebraFotaScheduleMode = "scheduled"
)

func PossibleValuesForZebraFotaScheduleMode() []string {
	return []string{
		string(ZebraFotaScheduleMode_InstallNow),
		string(ZebraFotaScheduleMode_Scheduled),
	}
}

func (s *ZebraFotaScheduleMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZebraFotaScheduleMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZebraFotaScheduleMode(input string) (*ZebraFotaScheduleMode, error) {
	vals := map[string]ZebraFotaScheduleMode{
		"installnow": ZebraFotaScheduleMode_InstallNow,
		"scheduled":  ZebraFotaScheduleMode_Scheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaScheduleMode(input)
	return &out, nil
}
