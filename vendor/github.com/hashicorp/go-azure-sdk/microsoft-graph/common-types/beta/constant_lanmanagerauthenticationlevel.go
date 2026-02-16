package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanManagerAuthenticationLevel string

const (
	LanManagerAuthenticationLevel_LmAndNltm             LanManagerAuthenticationLevel = "lmAndNltm"
	LanManagerAuthenticationLevel_LmAndNtlmOnly         LanManagerAuthenticationLevel = "lmAndNtlmOnly"
	LanManagerAuthenticationLevel_LmAndNtlmV2           LanManagerAuthenticationLevel = "lmAndNtlmV2"
	LanManagerAuthenticationLevel_LmNtlmAndNtlmV2       LanManagerAuthenticationLevel = "lmNtlmAndNtlmV2"
	LanManagerAuthenticationLevel_LmNtlmV2AndNotLm      LanManagerAuthenticationLevel = "lmNtlmV2AndNotLm"
	LanManagerAuthenticationLevel_LmNtlmV2AndNotLmOrNtm LanManagerAuthenticationLevel = "lmNtlmV2AndNotLmOrNtm"
)

func PossibleValuesForLanManagerAuthenticationLevel() []string {
	return []string{
		string(LanManagerAuthenticationLevel_LmAndNltm),
		string(LanManagerAuthenticationLevel_LmAndNtlmOnly),
		string(LanManagerAuthenticationLevel_LmAndNtlmV2),
		string(LanManagerAuthenticationLevel_LmNtlmAndNtlmV2),
		string(LanManagerAuthenticationLevel_LmNtlmV2AndNotLm),
		string(LanManagerAuthenticationLevel_LmNtlmV2AndNotLmOrNtm),
	}
}

func (s *LanManagerAuthenticationLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanManagerAuthenticationLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanManagerAuthenticationLevel(input string) (*LanManagerAuthenticationLevel, error) {
	vals := map[string]LanManagerAuthenticationLevel{
		"lmandnltm":             LanManagerAuthenticationLevel_LmAndNltm,
		"lmandntlmonly":         LanManagerAuthenticationLevel_LmAndNtlmOnly,
		"lmandntlmv2":           LanManagerAuthenticationLevel_LmAndNtlmV2,
		"lmntlmandntlmv2":       LanManagerAuthenticationLevel_LmNtlmAndNtlmV2,
		"lmntlmv2andnotlm":      LanManagerAuthenticationLevel_LmNtlmV2AndNotLm,
		"lmntlmv2andnotlmorntm": LanManagerAuthenticationLevel_LmNtlmV2AndNotLmOrNtm,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanManagerAuthenticationLevel(input)
	return &out, nil
}
