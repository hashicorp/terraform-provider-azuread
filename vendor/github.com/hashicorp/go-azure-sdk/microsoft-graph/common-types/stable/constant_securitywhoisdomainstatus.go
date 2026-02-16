package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityWhoisDomainStatus string

const (
	SecurityWhoisDomainStatus_ClientDeleteProhibited   SecurityWhoisDomainStatus = "clientDeleteProhibited"
	SecurityWhoisDomainStatus_ClientHold               SecurityWhoisDomainStatus = "clientHold"
	SecurityWhoisDomainStatus_ClientRenewProhibited    SecurityWhoisDomainStatus = "clientRenewProhibited"
	SecurityWhoisDomainStatus_ClientTransferProhibited SecurityWhoisDomainStatus = "clientTransferProhibited"
	SecurityWhoisDomainStatus_ClientUpdateProhibited   SecurityWhoisDomainStatus = "clientUpdateProhibited"
)

func PossibleValuesForSecurityWhoisDomainStatus() []string {
	return []string{
		string(SecurityWhoisDomainStatus_ClientDeleteProhibited),
		string(SecurityWhoisDomainStatus_ClientHold),
		string(SecurityWhoisDomainStatus_ClientRenewProhibited),
		string(SecurityWhoisDomainStatus_ClientTransferProhibited),
		string(SecurityWhoisDomainStatus_ClientUpdateProhibited),
	}
}

func (s *SecurityWhoisDomainStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityWhoisDomainStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityWhoisDomainStatus(input string) (*SecurityWhoisDomainStatus, error) {
	vals := map[string]SecurityWhoisDomainStatus{
		"clientdeleteprohibited":   SecurityWhoisDomainStatus_ClientDeleteProhibited,
		"clienthold":               SecurityWhoisDomainStatus_ClientHold,
		"clientrenewprohibited":    SecurityWhoisDomainStatus_ClientRenewProhibited,
		"clienttransferprohibited": SecurityWhoisDomainStatus_ClientTransferProhibited,
		"clientupdateprohibited":   SecurityWhoisDomainStatus_ClientUpdateProhibited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityWhoisDomainStatus(input)
	return &out, nil
}
