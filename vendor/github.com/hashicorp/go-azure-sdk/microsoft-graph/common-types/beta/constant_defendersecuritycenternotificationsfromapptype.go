package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderSecurityCenterNotificationsFromAppType string

const (
	DefenderSecurityCenterNotificationsFromAppType_BlockAllNotifications         DefenderSecurityCenterNotificationsFromAppType = "blockAllNotifications"
	DefenderSecurityCenterNotificationsFromAppType_BlockNoncriticalNotifications DefenderSecurityCenterNotificationsFromAppType = "blockNoncriticalNotifications"
	DefenderSecurityCenterNotificationsFromAppType_NotConfigured                 DefenderSecurityCenterNotificationsFromAppType = "notConfigured"
)

func PossibleValuesForDefenderSecurityCenterNotificationsFromAppType() []string {
	return []string{
		string(DefenderSecurityCenterNotificationsFromAppType_BlockAllNotifications),
		string(DefenderSecurityCenterNotificationsFromAppType_BlockNoncriticalNotifications),
		string(DefenderSecurityCenterNotificationsFromAppType_NotConfigured),
	}
}

func (s *DefenderSecurityCenterNotificationsFromAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderSecurityCenterNotificationsFromAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderSecurityCenterNotificationsFromAppType(input string) (*DefenderSecurityCenterNotificationsFromAppType, error) {
	vals := map[string]DefenderSecurityCenterNotificationsFromAppType{
		"blockallnotifications":         DefenderSecurityCenterNotificationsFromAppType_BlockAllNotifications,
		"blocknoncriticalnotifications": DefenderSecurityCenterNotificationsFromAppType_BlockNoncriticalNotifications,
		"notconfigured":                 DefenderSecurityCenterNotificationsFromAppType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderSecurityCenterNotificationsFromAppType(input)
	return &out, nil
}
