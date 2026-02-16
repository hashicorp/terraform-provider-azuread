package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthResult string

const (
	AuthenticationStrengthResult_CannotSatisfy                              AuthenticationStrengthResult = "cannotSatisfy"
	AuthenticationStrengthResult_CannotSatisfyDueToCombinationConfiguration AuthenticationStrengthResult = "cannotSatisfyDueToCombinationConfiguration"
	AuthenticationStrengthResult_MultipleChallengesRequired                 AuthenticationStrengthResult = "multipleChallengesRequired"
	AuthenticationStrengthResult_MultipleRegistrationsRequired              AuthenticationStrengthResult = "multipleRegistrationsRequired"
	AuthenticationStrengthResult_NotSet                                     AuthenticationStrengthResult = "notSet"
	AuthenticationStrengthResult_Satisfied                                  AuthenticationStrengthResult = "satisfied"
	AuthenticationStrengthResult_SingleChallengeRequired                    AuthenticationStrengthResult = "singleChallengeRequired"
	AuthenticationStrengthResult_SingleRegistrationRequired                 AuthenticationStrengthResult = "singleRegistrationRequired"
	AuthenticationStrengthResult_SkippedForProofUp                          AuthenticationStrengthResult = "skippedForProofUp"
)

func PossibleValuesForAuthenticationStrengthResult() []string {
	return []string{
		string(AuthenticationStrengthResult_CannotSatisfy),
		string(AuthenticationStrengthResult_CannotSatisfyDueToCombinationConfiguration),
		string(AuthenticationStrengthResult_MultipleChallengesRequired),
		string(AuthenticationStrengthResult_MultipleRegistrationsRequired),
		string(AuthenticationStrengthResult_NotSet),
		string(AuthenticationStrengthResult_Satisfied),
		string(AuthenticationStrengthResult_SingleChallengeRequired),
		string(AuthenticationStrengthResult_SingleRegistrationRequired),
		string(AuthenticationStrengthResult_SkippedForProofUp),
	}
}

func (s *AuthenticationStrengthResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationStrengthResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationStrengthResult(input string) (*AuthenticationStrengthResult, error) {
	vals := map[string]AuthenticationStrengthResult{
		"cannotsatisfy": AuthenticationStrengthResult_CannotSatisfy,
		"cannotsatisfyduetocombinationconfiguration": AuthenticationStrengthResult_CannotSatisfyDueToCombinationConfiguration,
		"multiplechallengesrequired":                 AuthenticationStrengthResult_MultipleChallengesRequired,
		"multipleregistrationsrequired":              AuthenticationStrengthResult_MultipleRegistrationsRequired,
		"notset":                                     AuthenticationStrengthResult_NotSet,
		"satisfied":                                  AuthenticationStrengthResult_Satisfied,
		"singlechallengerequired":                    AuthenticationStrengthResult_SingleChallengeRequired,
		"singleregistrationrequired":                 AuthenticationStrengthResult_SingleRegistrationRequired,
		"skippedforproofup":                          AuthenticationStrengthResult_SkippedForProofUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthResult(input)
	return &out, nil
}
