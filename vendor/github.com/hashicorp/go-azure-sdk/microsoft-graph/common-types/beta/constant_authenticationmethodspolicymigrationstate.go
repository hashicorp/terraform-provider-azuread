package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsPolicyMigrationState string

const (
	AuthenticationMethodsPolicyMigrationState_MigrationComplete   AuthenticationMethodsPolicyMigrationState = "migrationComplete"
	AuthenticationMethodsPolicyMigrationState_MigrationInProgress AuthenticationMethodsPolicyMigrationState = "migrationInProgress"
	AuthenticationMethodsPolicyMigrationState_PreMigration        AuthenticationMethodsPolicyMigrationState = "preMigration"
)

func PossibleValuesForAuthenticationMethodsPolicyMigrationState() []string {
	return []string{
		string(AuthenticationMethodsPolicyMigrationState_MigrationComplete),
		string(AuthenticationMethodsPolicyMigrationState_MigrationInProgress),
		string(AuthenticationMethodsPolicyMigrationState_PreMigration),
	}
}

func (s *AuthenticationMethodsPolicyMigrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodsPolicyMigrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodsPolicyMigrationState(input string) (*AuthenticationMethodsPolicyMigrationState, error) {
	vals := map[string]AuthenticationMethodsPolicyMigrationState{
		"migrationcomplete":   AuthenticationMethodsPolicyMigrationState_MigrationComplete,
		"migrationinprogress": AuthenticationMethodsPolicyMigrationState_MigrationInProgress,
		"premigration":        AuthenticationMethodsPolicyMigrationState_PreMigration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodsPolicyMigrationState(input)
	return &out, nil
}
