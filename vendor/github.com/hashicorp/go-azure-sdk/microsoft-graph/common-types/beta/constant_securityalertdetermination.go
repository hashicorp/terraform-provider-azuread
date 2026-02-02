package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertDetermination string

const (
	SecurityAlertDetermination_Apt                       SecurityAlertDetermination = "apt"
	SecurityAlertDetermination_CompromisedAccount        SecurityAlertDetermination = "compromisedAccount"
	SecurityAlertDetermination_ConfirmedActivity         SecurityAlertDetermination = "confirmedActivity"
	SecurityAlertDetermination_LineOfBusinessApplication SecurityAlertDetermination = "lineOfBusinessApplication"
	SecurityAlertDetermination_MaliciousUserActivity     SecurityAlertDetermination = "maliciousUserActivity"
	SecurityAlertDetermination_Malware                   SecurityAlertDetermination = "malware"
	SecurityAlertDetermination_MultiStagedAttack         SecurityAlertDetermination = "multiStagedAttack"
	SecurityAlertDetermination_NotEnoughDataToValidate   SecurityAlertDetermination = "notEnoughDataToValidate"
	SecurityAlertDetermination_NotMalicious              SecurityAlertDetermination = "notMalicious"
	SecurityAlertDetermination_Other                     SecurityAlertDetermination = "other"
	SecurityAlertDetermination_Phishing                  SecurityAlertDetermination = "phishing"
	SecurityAlertDetermination_SecurityPersonnel         SecurityAlertDetermination = "securityPersonnel"
	SecurityAlertDetermination_SecurityTesting           SecurityAlertDetermination = "securityTesting"
	SecurityAlertDetermination_Unknown                   SecurityAlertDetermination = "unknown"
	SecurityAlertDetermination_UnwantedSoftware          SecurityAlertDetermination = "unwantedSoftware"
)

func PossibleValuesForSecurityAlertDetermination() []string {
	return []string{
		string(SecurityAlertDetermination_Apt),
		string(SecurityAlertDetermination_CompromisedAccount),
		string(SecurityAlertDetermination_ConfirmedActivity),
		string(SecurityAlertDetermination_LineOfBusinessApplication),
		string(SecurityAlertDetermination_MaliciousUserActivity),
		string(SecurityAlertDetermination_Malware),
		string(SecurityAlertDetermination_MultiStagedAttack),
		string(SecurityAlertDetermination_NotEnoughDataToValidate),
		string(SecurityAlertDetermination_NotMalicious),
		string(SecurityAlertDetermination_Other),
		string(SecurityAlertDetermination_Phishing),
		string(SecurityAlertDetermination_SecurityPersonnel),
		string(SecurityAlertDetermination_SecurityTesting),
		string(SecurityAlertDetermination_Unknown),
		string(SecurityAlertDetermination_UnwantedSoftware),
	}
}

func (s *SecurityAlertDetermination) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAlertDetermination(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAlertDetermination(input string) (*SecurityAlertDetermination, error) {
	vals := map[string]SecurityAlertDetermination{
		"apt":                       SecurityAlertDetermination_Apt,
		"compromisedaccount":        SecurityAlertDetermination_CompromisedAccount,
		"confirmedactivity":         SecurityAlertDetermination_ConfirmedActivity,
		"lineofbusinessapplication": SecurityAlertDetermination_LineOfBusinessApplication,
		"malicioususeractivity":     SecurityAlertDetermination_MaliciousUserActivity,
		"malware":                   SecurityAlertDetermination_Malware,
		"multistagedattack":         SecurityAlertDetermination_MultiStagedAttack,
		"notenoughdatatovalidate":   SecurityAlertDetermination_NotEnoughDataToValidate,
		"notmalicious":              SecurityAlertDetermination_NotMalicious,
		"other":                     SecurityAlertDetermination_Other,
		"phishing":                  SecurityAlertDetermination_Phishing,
		"securitypersonnel":         SecurityAlertDetermination_SecurityPersonnel,
		"securitytesting":           SecurityAlertDetermination_SecurityTesting,
		"unknown":                   SecurityAlertDetermination_Unknown,
		"unwantedsoftware":          SecurityAlertDetermination_UnwantedSoftware,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertDetermination(input)
	return &out, nil
}
