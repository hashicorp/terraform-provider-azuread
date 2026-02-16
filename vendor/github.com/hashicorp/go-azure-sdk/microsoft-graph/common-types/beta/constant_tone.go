package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Tone string

const (
	Tone_A     Tone = "a"
	Tone_B     Tone = "b"
	Tone_C     Tone = "c"
	Tone_D     Tone = "d"
	Tone_Flash Tone = "flash"
	Tone_Pound Tone = "pound"
	Tone_Star  Tone = "star"
	Tone_Tone0 Tone = "tone0"
	Tone_Tone1 Tone = "tone1"
	Tone_Tone2 Tone = "tone2"
	Tone_Tone3 Tone = "tone3"
	Tone_Tone4 Tone = "tone4"
	Tone_Tone5 Tone = "tone5"
	Tone_Tone6 Tone = "tone6"
	Tone_Tone7 Tone = "tone7"
	Tone_Tone8 Tone = "tone8"
	Tone_Tone9 Tone = "tone9"
)

func PossibleValuesForTone() []string {
	return []string{
		string(Tone_A),
		string(Tone_B),
		string(Tone_C),
		string(Tone_D),
		string(Tone_Flash),
		string(Tone_Pound),
		string(Tone_Star),
		string(Tone_Tone0),
		string(Tone_Tone1),
		string(Tone_Tone2),
		string(Tone_Tone3),
		string(Tone_Tone4),
		string(Tone_Tone5),
		string(Tone_Tone6),
		string(Tone_Tone7),
		string(Tone_Tone8),
		string(Tone_Tone9),
	}
}

func (s *Tone) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTone(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTone(input string) (*Tone, error) {
	vals := map[string]Tone{
		"a":     Tone_A,
		"b":     Tone_B,
		"c":     Tone_C,
		"d":     Tone_D,
		"flash": Tone_Flash,
		"pound": Tone_Pound,
		"star":  Tone_Star,
		"tone0": Tone_Tone0,
		"tone1": Tone_Tone1,
		"tone2": Tone_Tone2,
		"tone3": Tone_Tone3,
		"tone4": Tone_Tone4,
		"tone5": Tone_Tone5,
		"tone6": Tone_Tone6,
		"tone7": Tone_Tone7,
		"tone8": Tone_Tone8,
		"tone9": Tone_Tone9,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Tone(input)
	return &out, nil
}
