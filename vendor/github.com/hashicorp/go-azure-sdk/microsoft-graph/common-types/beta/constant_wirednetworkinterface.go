package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WiredNetworkInterface string

const (
	WiredNetworkInterface_AnyEthernet          WiredNetworkInterface = "anyEthernet"
	WiredNetworkInterface_FirstActiveEthernet  WiredNetworkInterface = "firstActiveEthernet"
	WiredNetworkInterface_FirstEthernet        WiredNetworkInterface = "firstEthernet"
	WiredNetworkInterface_SecondActiveEthernet WiredNetworkInterface = "secondActiveEthernet"
	WiredNetworkInterface_SecondEthernet       WiredNetworkInterface = "secondEthernet"
	WiredNetworkInterface_ThirdActiveEthernet  WiredNetworkInterface = "thirdActiveEthernet"
	WiredNetworkInterface_ThirdEthernet        WiredNetworkInterface = "thirdEthernet"
)

func PossibleValuesForWiredNetworkInterface() []string {
	return []string{
		string(WiredNetworkInterface_AnyEthernet),
		string(WiredNetworkInterface_FirstActiveEthernet),
		string(WiredNetworkInterface_FirstEthernet),
		string(WiredNetworkInterface_SecondActiveEthernet),
		string(WiredNetworkInterface_SecondEthernet),
		string(WiredNetworkInterface_ThirdActiveEthernet),
		string(WiredNetworkInterface_ThirdEthernet),
	}
}

func (s *WiredNetworkInterface) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWiredNetworkInterface(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWiredNetworkInterface(input string) (*WiredNetworkInterface, error) {
	vals := map[string]WiredNetworkInterface{
		"anyethernet":          WiredNetworkInterface_AnyEthernet,
		"firstactiveethernet":  WiredNetworkInterface_FirstActiveEthernet,
		"firstethernet":        WiredNetworkInterface_FirstEthernet,
		"secondactiveethernet": WiredNetworkInterface_SecondActiveEthernet,
		"secondethernet":       WiredNetworkInterface_SecondEthernet,
		"thirdactiveethernet":  WiredNetworkInterface_ThirdActiveEthernet,
		"thirdethernet":        WiredNetworkInterface_ThirdEthernet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WiredNetworkInterface(input)
	return &out, nil
}
