package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningResult string

const (
	ProvisioningResult_Failure ProvisioningResult = "failure"
	ProvisioningResult_Skipped ProvisioningResult = "skipped"
	ProvisioningResult_Success ProvisioningResult = "success"
	ProvisioningResult_Warning ProvisioningResult = "warning"
)

func PossibleValuesForProvisioningResult() []string {
	return []string{
		string(ProvisioningResult_Failure),
		string(ProvisioningResult_Skipped),
		string(ProvisioningResult_Success),
		string(ProvisioningResult_Warning),
	}
}

func (s *ProvisioningResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningResult(input string) (*ProvisioningResult, error) {
	vals := map[string]ProvisioningResult{
		"failure": ProvisioningResult_Failure,
		"skipped": ProvisioningResult_Skipped,
		"success": ProvisioningResult_Success,
		"warning": ProvisioningResult_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningResult(input)
	return &out, nil
}
