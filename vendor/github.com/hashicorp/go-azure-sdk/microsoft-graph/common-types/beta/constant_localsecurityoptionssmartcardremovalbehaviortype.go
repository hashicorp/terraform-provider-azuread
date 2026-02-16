package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocalSecurityOptionsSmartCardRemovalBehaviorType string

const (
	LocalSecurityOptionsSmartCardRemovalBehaviorType_DisconnectRemoteDesktopSession LocalSecurityOptionsSmartCardRemovalBehaviorType = "disconnectRemoteDesktopSession"
	LocalSecurityOptionsSmartCardRemovalBehaviorType_ForceLogoff                    LocalSecurityOptionsSmartCardRemovalBehaviorType = "forceLogoff"
	LocalSecurityOptionsSmartCardRemovalBehaviorType_LockWorkstation                LocalSecurityOptionsSmartCardRemovalBehaviorType = "lockWorkstation"
	LocalSecurityOptionsSmartCardRemovalBehaviorType_NoAction                       LocalSecurityOptionsSmartCardRemovalBehaviorType = "noAction"
)

func PossibleValuesForLocalSecurityOptionsSmartCardRemovalBehaviorType() []string {
	return []string{
		string(LocalSecurityOptionsSmartCardRemovalBehaviorType_DisconnectRemoteDesktopSession),
		string(LocalSecurityOptionsSmartCardRemovalBehaviorType_ForceLogoff),
		string(LocalSecurityOptionsSmartCardRemovalBehaviorType_LockWorkstation),
		string(LocalSecurityOptionsSmartCardRemovalBehaviorType_NoAction),
	}
}

func (s *LocalSecurityOptionsSmartCardRemovalBehaviorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocalSecurityOptionsSmartCardRemovalBehaviorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocalSecurityOptionsSmartCardRemovalBehaviorType(input string) (*LocalSecurityOptionsSmartCardRemovalBehaviorType, error) {
	vals := map[string]LocalSecurityOptionsSmartCardRemovalBehaviorType{
		"disconnectremotedesktopsession": LocalSecurityOptionsSmartCardRemovalBehaviorType_DisconnectRemoteDesktopSession,
		"forcelogoff":                    LocalSecurityOptionsSmartCardRemovalBehaviorType_ForceLogoff,
		"lockworkstation":                LocalSecurityOptionsSmartCardRemovalBehaviorType_LockWorkstation,
		"noaction":                       LocalSecurityOptionsSmartCardRemovalBehaviorType_NoAction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocalSecurityOptionsSmartCardRemovalBehaviorType(input)
	return &out, nil
}
