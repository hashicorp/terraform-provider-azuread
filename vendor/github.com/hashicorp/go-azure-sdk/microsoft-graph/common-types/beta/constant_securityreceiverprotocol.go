package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityReceiverProtocol string

const (
	SecurityReceiverProtocol_Ftp       SecurityReceiverProtocol = "ftp"
	SecurityReceiverProtocol_Ftps      SecurityReceiverProtocol = "ftps"
	SecurityReceiverProtocol_SyslogTcp SecurityReceiverProtocol = "syslogTcp"
	SecurityReceiverProtocol_SyslogTls SecurityReceiverProtocol = "syslogTls"
	SecurityReceiverProtocol_SyslogUdp SecurityReceiverProtocol = "syslogUdp"
)

func PossibleValuesForSecurityReceiverProtocol() []string {
	return []string{
		string(SecurityReceiverProtocol_Ftp),
		string(SecurityReceiverProtocol_Ftps),
		string(SecurityReceiverProtocol_SyslogTcp),
		string(SecurityReceiverProtocol_SyslogTls),
		string(SecurityReceiverProtocol_SyslogUdp),
	}
}

func (s *SecurityReceiverProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityReceiverProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityReceiverProtocol(input string) (*SecurityReceiverProtocol, error) {
	vals := map[string]SecurityReceiverProtocol{
		"ftp":       SecurityReceiverProtocol_Ftp,
		"ftps":      SecurityReceiverProtocol_Ftps,
		"syslogtcp": SecurityReceiverProtocol_SyslogTcp,
		"syslogtls": SecurityReceiverProtocol_SyslogTls,
		"syslogudp": SecurityReceiverProtocol_SyslogUdp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityReceiverProtocol(input)
	return &out, nil
}
