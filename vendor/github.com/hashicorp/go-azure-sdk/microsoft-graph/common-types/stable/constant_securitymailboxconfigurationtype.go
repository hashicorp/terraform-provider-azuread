package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailboxConfigurationType string

const (
	SecurityMailboxConfigurationType_EwsSettings        SecurityMailboxConfigurationType = "ewsSettings"
	SecurityMailboxConfigurationType_MailDelegation     SecurityMailboxConfigurationType = "mailDelegation"
	SecurityMailboxConfigurationType_MailForwardingRule SecurityMailboxConfigurationType = "mailForwardingRule"
	SecurityMailboxConfigurationType_OwaSettings        SecurityMailboxConfigurationType = "owaSettings"
	SecurityMailboxConfigurationType_UserInboxRule      SecurityMailboxConfigurationType = "userInboxRule"
)

func PossibleValuesForSecurityMailboxConfigurationType() []string {
	return []string{
		string(SecurityMailboxConfigurationType_EwsSettings),
		string(SecurityMailboxConfigurationType_MailDelegation),
		string(SecurityMailboxConfigurationType_MailForwardingRule),
		string(SecurityMailboxConfigurationType_OwaSettings),
		string(SecurityMailboxConfigurationType_UserInboxRule),
	}
}

func (s *SecurityMailboxConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMailboxConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMailboxConfigurationType(input string) (*SecurityMailboxConfigurationType, error) {
	vals := map[string]SecurityMailboxConfigurationType{
		"ewssettings":        SecurityMailboxConfigurationType_EwsSettings,
		"maildelegation":     SecurityMailboxConfigurationType_MailDelegation,
		"mailforwardingrule": SecurityMailboxConfigurationType_MailForwardingRule,
		"owasettings":        SecurityMailboxConfigurationType_OwaSettings,
		"userinboxrule":      SecurityMailboxConfigurationType_UserInboxRule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailboxConfigurationType(input)
	return &out, nil
}
