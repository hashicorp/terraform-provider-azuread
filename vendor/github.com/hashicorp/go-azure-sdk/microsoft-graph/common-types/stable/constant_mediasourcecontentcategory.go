package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaSourceContentCategory string

const (
	MediaSourceContentCategory_Chat            MediaSourceContentCategory = "chat"
	MediaSourceContentCategory_Comment         MediaSourceContentCategory = "comment"
	MediaSourceContentCategory_LiveStream      MediaSourceContentCategory = "liveStream"
	MediaSourceContentCategory_Meeting         MediaSourceContentCategory = "meeting"
	MediaSourceContentCategory_Note            MediaSourceContentCategory = "note"
	MediaSourceContentCategory_Presentation    MediaSourceContentCategory = "presentation"
	MediaSourceContentCategory_Profile         MediaSourceContentCategory = "profile"
	MediaSourceContentCategory_ScreenRecording MediaSourceContentCategory = "screenRecording"
	MediaSourceContentCategory_Story           MediaSourceContentCategory = "story"
)

func PossibleValuesForMediaSourceContentCategory() []string {
	return []string{
		string(MediaSourceContentCategory_Chat),
		string(MediaSourceContentCategory_Comment),
		string(MediaSourceContentCategory_LiveStream),
		string(MediaSourceContentCategory_Meeting),
		string(MediaSourceContentCategory_Note),
		string(MediaSourceContentCategory_Presentation),
		string(MediaSourceContentCategory_Profile),
		string(MediaSourceContentCategory_ScreenRecording),
		string(MediaSourceContentCategory_Story),
	}
}

func (s *MediaSourceContentCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaSourceContentCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaSourceContentCategory(input string) (*MediaSourceContentCategory, error) {
	vals := map[string]MediaSourceContentCategory{
		"chat":            MediaSourceContentCategory_Chat,
		"comment":         MediaSourceContentCategory_Comment,
		"livestream":      MediaSourceContentCategory_LiveStream,
		"meeting":         MediaSourceContentCategory_Meeting,
		"note":            MediaSourceContentCategory_Note,
		"presentation":    MediaSourceContentCategory_Presentation,
		"profile":         MediaSourceContentCategory_Profile,
		"screenrecording": MediaSourceContentCategory_ScreenRecording,
		"story":           MediaSourceContentCategory_Story,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaSourceContentCategory(input)
	return &out, nil
}
