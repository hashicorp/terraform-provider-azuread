package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalEmailOtpState string

const (
	ExternalEmailOtpState_Default  ExternalEmailOtpState = "default"
	ExternalEmailOtpState_Disabled ExternalEmailOtpState = "disabled"
	ExternalEmailOtpState_Enabled  ExternalEmailOtpState = "enabled"
)

func PossibleValuesForExternalEmailOtpState() []string {
	return []string{
		string(ExternalEmailOtpState_Default),
		string(ExternalEmailOtpState_Disabled),
		string(ExternalEmailOtpState_Enabled),
	}
}

func (s *ExternalEmailOtpState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalEmailOtpState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalEmailOtpState(input string) (*ExternalEmailOtpState, error) {
	vals := map[string]ExternalEmailOtpState{
		"default":  ExternalEmailOtpState_Default,
		"disabled": ExternalEmailOtpState_Disabled,
		"enabled":  ExternalEmailOtpState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalEmailOtpState(input)
	return &out, nil
}
