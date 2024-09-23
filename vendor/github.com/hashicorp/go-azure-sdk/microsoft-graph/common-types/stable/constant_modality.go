package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Modality string

const (
	Modality_Audio                   Modality = "audio"
	Modality_Data                    Modality = "data"
	Modality_Video                   Modality = "video"
	Modality_VideoBasedScreenSharing Modality = "videoBasedScreenSharing"
)

func PossibleValuesForModality() []string {
	return []string{
		string(Modality_Audio),
		string(Modality_Data),
		string(Modality_Video),
		string(Modality_VideoBasedScreenSharing),
	}
}

func (s *Modality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseModality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseModality(input string) (*Modality, error) {
	vals := map[string]Modality{
		"audio":                   Modality_Audio,
		"data":                    Modality_Data,
		"video":                   Modality_Video,
		"videobasedscreensharing": Modality_VideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Modality(input)
	return &out, nil
}
