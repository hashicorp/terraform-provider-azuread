package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileThreatDefensePartnerPriority string

const (
	MobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner MobileThreatDefensePartnerPriority = "defenderOverThirdPartyPartner"
	MobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender MobileThreatDefensePartnerPriority = "thirdPartyPartnerOverDefender"
)

func PossibleValuesForMobileThreatDefensePartnerPriority() []string {
	return []string{
		string(MobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner),
		string(MobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender),
	}
}

func (s *MobileThreatDefensePartnerPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileThreatDefensePartnerPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileThreatDefensePartnerPriority(input string) (*MobileThreatDefensePartnerPriority, error) {
	vals := map[string]MobileThreatDefensePartnerPriority{
		"defenderoverthirdpartypartner": MobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner,
		"thirdpartypartneroverdefender": MobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileThreatDefensePartnerPriority(input)
	return &out, nil
}
