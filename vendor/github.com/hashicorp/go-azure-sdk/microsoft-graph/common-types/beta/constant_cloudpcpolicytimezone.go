package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPolicyTimezone string

const (
	CloudPCPolicyTimezone_Acst CloudPCPolicyTimezone = "acst"
	CloudPCPolicyTimezone_Akst CloudPCPolicyTimezone = "akst"
	CloudPCPolicyTimezone_Art  CloudPCPolicyTimezone = "art"
	CloudPCPolicyTimezone_Ast  CloudPCPolicyTimezone = "ast"
	CloudPCPolicyTimezone_Azot CloudPCPolicyTimezone = "azot"
	CloudPCPolicyTimezone_Bit  CloudPCPolicyTimezone = "bit"
	CloudPCPolicyTimezone_Bst  CloudPCPolicyTimezone = "bst"
	CloudPCPolicyTimezone_Cat  CloudPCPolicyTimezone = "cat"
	CloudPCPolicyTimezone_Cet  CloudPCPolicyTimezone = "cet"
	CloudPCPolicyTimezone_Cst  CloudPCPolicyTimezone = "cst"
	CloudPCPolicyTimezone_East CloudPCPolicyTimezone = "east"
	CloudPCPolicyTimezone_Eat  CloudPCPolicyTimezone = "eat"
	CloudPCPolicyTimezone_Est  CloudPCPolicyTimezone = "est"
	CloudPCPolicyTimezone_Fjt  CloudPCPolicyTimezone = "fjt"
	CloudPCPolicyTimezone_Get  CloudPCPolicyTimezone = "get"
	CloudPCPolicyTimezone_Gmt  CloudPCPolicyTimezone = "gmt"
	CloudPCPolicyTimezone_Gst  CloudPCPolicyTimezone = "gst"
	CloudPCPolicyTimezone_Hst  CloudPCPolicyTimezone = "hst"
	CloudPCPolicyTimezone_Ist  CloudPCPolicyTimezone = "ist"
	CloudPCPolicyTimezone_Jst  CloudPCPolicyTimezone = "jst"
	CloudPCPolicyTimezone_Lint CloudPCPolicyTimezone = "lint"
	CloudPCPolicyTimezone_Mit  CloudPCPolicyTimezone = "mit"
	CloudPCPolicyTimezone_Mst  CloudPCPolicyTimezone = "mst"
	CloudPCPolicyTimezone_Nst  CloudPCPolicyTimezone = "nst"
	CloudPCPolicyTimezone_Nut  CloudPCPolicyTimezone = "nut"
	CloudPCPolicyTimezone_Pgt  CloudPCPolicyTimezone = "pgt"
	CloudPCPolicyTimezone_Pkt  CloudPCPolicyTimezone = "pkt"
	CloudPCPolicyTimezone_Pst  CloudPCPolicyTimezone = "pst"
	CloudPCPolicyTimezone_Sbt  CloudPCPolicyTimezone = "sbt"
	CloudPCPolicyTimezone_Tha  CloudPCPolicyTimezone = "tha"
	CloudPCPolicyTimezone_Tot  CloudPCPolicyTimezone = "tot"
)

func PossibleValuesForCloudPCPolicyTimezone() []string {
	return []string{
		string(CloudPCPolicyTimezone_Acst),
		string(CloudPCPolicyTimezone_Akst),
		string(CloudPCPolicyTimezone_Art),
		string(CloudPCPolicyTimezone_Ast),
		string(CloudPCPolicyTimezone_Azot),
		string(CloudPCPolicyTimezone_Bit),
		string(CloudPCPolicyTimezone_Bst),
		string(CloudPCPolicyTimezone_Cat),
		string(CloudPCPolicyTimezone_Cet),
		string(CloudPCPolicyTimezone_Cst),
		string(CloudPCPolicyTimezone_East),
		string(CloudPCPolicyTimezone_Eat),
		string(CloudPCPolicyTimezone_Est),
		string(CloudPCPolicyTimezone_Fjt),
		string(CloudPCPolicyTimezone_Get),
		string(CloudPCPolicyTimezone_Gmt),
		string(CloudPCPolicyTimezone_Gst),
		string(CloudPCPolicyTimezone_Hst),
		string(CloudPCPolicyTimezone_Ist),
		string(CloudPCPolicyTimezone_Jst),
		string(CloudPCPolicyTimezone_Lint),
		string(CloudPCPolicyTimezone_Mit),
		string(CloudPCPolicyTimezone_Mst),
		string(CloudPCPolicyTimezone_Nst),
		string(CloudPCPolicyTimezone_Nut),
		string(CloudPCPolicyTimezone_Pgt),
		string(CloudPCPolicyTimezone_Pkt),
		string(CloudPCPolicyTimezone_Pst),
		string(CloudPCPolicyTimezone_Sbt),
		string(CloudPCPolicyTimezone_Tha),
		string(CloudPCPolicyTimezone_Tot),
	}
}

func (s *CloudPCPolicyTimezone) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPolicyTimezone(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPolicyTimezone(input string) (*CloudPCPolicyTimezone, error) {
	vals := map[string]CloudPCPolicyTimezone{
		"acst": CloudPCPolicyTimezone_Acst,
		"akst": CloudPCPolicyTimezone_Akst,
		"art":  CloudPCPolicyTimezone_Art,
		"ast":  CloudPCPolicyTimezone_Ast,
		"azot": CloudPCPolicyTimezone_Azot,
		"bit":  CloudPCPolicyTimezone_Bit,
		"bst":  CloudPCPolicyTimezone_Bst,
		"cat":  CloudPCPolicyTimezone_Cat,
		"cet":  CloudPCPolicyTimezone_Cet,
		"cst":  CloudPCPolicyTimezone_Cst,
		"east": CloudPCPolicyTimezone_East,
		"eat":  CloudPCPolicyTimezone_Eat,
		"est":  CloudPCPolicyTimezone_Est,
		"fjt":  CloudPCPolicyTimezone_Fjt,
		"get":  CloudPCPolicyTimezone_Get,
		"gmt":  CloudPCPolicyTimezone_Gmt,
		"gst":  CloudPCPolicyTimezone_Gst,
		"hst":  CloudPCPolicyTimezone_Hst,
		"ist":  CloudPCPolicyTimezone_Ist,
		"jst":  CloudPCPolicyTimezone_Jst,
		"lint": CloudPCPolicyTimezone_Lint,
		"mit":  CloudPCPolicyTimezone_Mit,
		"mst":  CloudPCPolicyTimezone_Mst,
		"nst":  CloudPCPolicyTimezone_Nst,
		"nut":  CloudPCPolicyTimezone_Nut,
		"pgt":  CloudPCPolicyTimezone_Pgt,
		"pkt":  CloudPCPolicyTimezone_Pkt,
		"pst":  CloudPCPolicyTimezone_Pst,
		"sbt":  CloudPCPolicyTimezone_Sbt,
		"tha":  CloudPCPolicyTimezone_Tha,
		"tot":  CloudPCPolicyTimezone_Tot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPolicyTimezone(input)
	return &out, nil
}
