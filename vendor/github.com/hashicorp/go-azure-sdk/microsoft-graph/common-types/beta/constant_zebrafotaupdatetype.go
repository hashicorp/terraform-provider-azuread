package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaUpdateType string

const (
	ZebraFotaUpdateType_Auto   ZebraFotaUpdateType = "auto"
	ZebraFotaUpdateType_Custom ZebraFotaUpdateType = "custom"
	ZebraFotaUpdateType_Latest ZebraFotaUpdateType = "latest"
)

func PossibleValuesForZebraFotaUpdateType() []string {
	return []string{
		string(ZebraFotaUpdateType_Auto),
		string(ZebraFotaUpdateType_Custom),
		string(ZebraFotaUpdateType_Latest),
	}
}

func (s *ZebraFotaUpdateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZebraFotaUpdateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZebraFotaUpdateType(input string) (*ZebraFotaUpdateType, error) {
	vals := map[string]ZebraFotaUpdateType{
		"auto":   ZebraFotaUpdateType_Auto,
		"custom": ZebraFotaUpdateType_Custom,
		"latest": ZebraFotaUpdateType_Latest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaUpdateType(input)
	return &out, nil
}
