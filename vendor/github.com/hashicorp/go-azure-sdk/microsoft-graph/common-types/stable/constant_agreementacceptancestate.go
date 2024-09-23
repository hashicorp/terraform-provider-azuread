package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementAcceptanceState string

const (
	AgreementAcceptanceState_Accepted AgreementAcceptanceState = "accepted"
	AgreementAcceptanceState_Declined AgreementAcceptanceState = "declined"
)

func PossibleValuesForAgreementAcceptanceState() []string {
	return []string{
		string(AgreementAcceptanceState_Accepted),
		string(AgreementAcceptanceState_Declined),
	}
}

func (s *AgreementAcceptanceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgreementAcceptanceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgreementAcceptanceState(input string) (*AgreementAcceptanceState, error) {
	vals := map[string]AgreementAcceptanceState{
		"accepted": AgreementAcceptanceState_Accepted,
		"declined": AgreementAcceptanceState_Declined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgreementAcceptanceState(input)
	return &out, nil
}
