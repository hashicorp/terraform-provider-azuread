package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdAttendanceReportId{}

// MeOnlineMeetingIdAttendanceReportId is a struct representing the Resource ID for a Me Online Meeting Id Attendance Report
type MeOnlineMeetingIdAttendanceReportId struct {
	OnlineMeetingId           string
	MeetingAttendanceReportId string
}

// NewMeOnlineMeetingIdAttendanceReportID returns a new MeOnlineMeetingIdAttendanceReportId struct
func NewMeOnlineMeetingIdAttendanceReportID(onlineMeetingId string, meetingAttendanceReportId string) MeOnlineMeetingIdAttendanceReportId {
	return MeOnlineMeetingIdAttendanceReportId{
		OnlineMeetingId:           onlineMeetingId,
		MeetingAttendanceReportId: meetingAttendanceReportId,
	}
}

// ParseMeOnlineMeetingIdAttendanceReportID parses 'input' into a MeOnlineMeetingIdAttendanceReportId
func ParseMeOnlineMeetingIdAttendanceReportID(input string) (*MeOnlineMeetingIdAttendanceReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdAttendanceReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdAttendanceReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnlineMeetingIdAttendanceReportIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingIdAttendanceReportId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingIdAttendanceReportIDInsensitively(input string) (*MeOnlineMeetingIdAttendanceReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdAttendanceReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdAttendanceReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnlineMeetingIdAttendanceReportId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.MeetingAttendanceReportId, ok = input.Parsed["meetingAttendanceReportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "meetingAttendanceReportId", input)
	}

	return nil
}

// ValidateMeOnlineMeetingIdAttendanceReportID checks that 'input' can be parsed as a Me Online Meeting Id Attendance Report ID
func ValidateMeOnlineMeetingIdAttendanceReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingIdAttendanceReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Id Attendance Report ID
func (id MeOnlineMeetingIdAttendanceReportId) ID() string {
	fmtString := "/me/onlineMeetings/%s/attendanceReports/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.MeetingAttendanceReportId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Id Attendance Report ID
func (id MeOnlineMeetingIdAttendanceReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("attendanceReports", "attendanceReports", "attendanceReports"),
		resourceids.UserSpecifiedSegment("meetingAttendanceReportId", "meetingAttendanceReportId"),
	}
}

// String returns a human-readable description of this Me Online Meeting Id Attendance Report ID
func (id MeOnlineMeetingIdAttendanceReportId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Attendance Report: %q", id.MeetingAttendanceReportId),
	}
	return fmt.Sprintf("Me Online Meeting Id Attendance Report (%s)", strings.Join(components, "\n"))
}
