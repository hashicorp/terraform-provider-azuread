package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiableCredentialPresentationStatusCode string

const (
	VerifiableCredentialPresentationStatusCode_Presentationverified VerifiableCredentialPresentationStatusCode = "presentation_verified"
	VerifiableCredentialPresentationStatusCode_Requestretrieved     VerifiableCredentialPresentationStatusCode = "request_retrieved"
)

func PossibleValuesForVerifiableCredentialPresentationStatusCode() []string {
	return []string{
		string(VerifiableCredentialPresentationStatusCode_Presentationverified),
		string(VerifiableCredentialPresentationStatusCode_Requestretrieved),
	}
}

func (s *VerifiableCredentialPresentationStatusCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVerifiableCredentialPresentationStatusCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVerifiableCredentialPresentationStatusCode(input string) (*VerifiableCredentialPresentationStatusCode, error) {
	vals := map[string]VerifiableCredentialPresentationStatusCode{
		"presentation_verified": VerifiableCredentialPresentationStatusCode_Presentationverified,
		"request_retrieved":     VerifiableCredentialPresentationStatusCode_Requestretrieved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VerifiableCredentialPresentationStatusCode(input)
	return &out, nil
}
