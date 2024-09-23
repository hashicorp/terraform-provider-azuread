package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaDirection string

const (
	MediaDirection_Inactive    MediaDirection = "inactive"
	MediaDirection_ReceiveOnly MediaDirection = "receiveOnly"
	MediaDirection_SendOnly    MediaDirection = "sendOnly"
	MediaDirection_SendReceive MediaDirection = "sendReceive"
)

func PossibleValuesForMediaDirection() []string {
	return []string{
		string(MediaDirection_Inactive),
		string(MediaDirection_ReceiveOnly),
		string(MediaDirection_SendOnly),
		string(MediaDirection_SendReceive),
	}
}

func (s *MediaDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaDirection(input string) (*MediaDirection, error) {
	vals := map[string]MediaDirection{
		"inactive":    MediaDirection_Inactive,
		"receiveonly": MediaDirection_ReceiveOnly,
		"sendonly":    MediaDirection_SendOnly,
		"sendreceive": MediaDirection_SendReceive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaDirection(input)
	return &out, nil
}
