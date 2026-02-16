package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId{}

// MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId is a struct representing the Resource ID for a Me Online Meeting Id Attendance Report Id Attendance Record
type MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId struct {
	OnlineMeetingId           string
	MeetingAttendanceReportId string
	AttendanceRecordId        string
}

// NewMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID returns a new MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId struct
func NewMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID(onlineMeetingId string, meetingAttendanceReportId string, attendanceRecordId string) MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId {
	return MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId{
		OnlineMeetingId:           onlineMeetingId,
		MeetingAttendanceReportId: meetingAttendanceReportId,
		AttendanceRecordId:        attendanceRecordId,
	}
}

// ParseMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID parses 'input' into a MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId
func ParseMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID(input string) (*MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnlineMeetingIdAttendanceReportIdAttendanceRecordIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingIdAttendanceReportIdAttendanceRecordIDInsensitively(input string) (*MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.MeetingAttendanceReportId, ok = input.Parsed["meetingAttendanceReportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "meetingAttendanceReportId", input)
	}

	if id.AttendanceRecordId, ok = input.Parsed["attendanceRecordId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attendanceRecordId", input)
	}

	return nil
}

// ValidateMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID checks that 'input' can be parsed as a Me Online Meeting Id Attendance Report Id Attendance Record ID
func ValidateMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Id Attendance Report Id Attendance Record ID
func (id MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId) ID() string {
	fmtString := "/me/onlineMeetings/%s/attendanceReports/%s/attendanceRecords/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.MeetingAttendanceReportId, id.AttendanceRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Id Attendance Report Id Attendance Record ID
func (id MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("attendanceReports", "attendanceReports", "attendanceReports"),
		resourceids.UserSpecifiedSegment("meetingAttendanceReportId", "meetingAttendanceReportId"),
		resourceids.StaticSegment("attendanceRecords", "attendanceRecords", "attendanceRecords"),
		resourceids.UserSpecifiedSegment("attendanceRecordId", "attendanceRecordId"),
	}
}

// String returns a human-readable description of this Me Online Meeting Id Attendance Report Id Attendance Record ID
func (id MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Attendance Report: %q", id.MeetingAttendanceReportId),
		fmt.Sprintf("Attendance Record: %q", id.AttendanceRecordId),
	}
	return fmt.Sprintf("Me Online Meeting Id Attendance Report Id Attendance Record (%s)", strings.Join(components, "\n"))
}
