package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId{}

// UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId is a struct representing the Resource ID for a User Id Online Meeting Id Attendance Report Id Attendance Record
type UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId struct {
	UserId                    string
	OnlineMeetingId           string
	MeetingAttendanceReportId string
	AttendanceRecordId        string
}

// NewUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID returns a new UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId struct
func NewUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID(userId string, onlineMeetingId string, meetingAttendanceReportId string, attendanceRecordId string) UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId {
	return UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId{
		UserId:                    userId,
		OnlineMeetingId:           onlineMeetingId,
		MeetingAttendanceReportId: meetingAttendanceReportId,
		AttendanceRecordId:        attendanceRecordId,
	}
}

// ParseUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID parses 'input' into a UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId
func ParseUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID(input string) (*UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordIDInsensitively parses 'input' case-insensitively into a UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordIDInsensitively(input string) (*UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

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

// ValidateUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID checks that 'input' can be parsed as a User Id Online Meeting Id Attendance Report Id Attendance Record ID
func ValidateUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Online Meeting Id Attendance Report Id Attendance Record ID
func (id UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/attendanceReports/%s/attendanceRecords/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.MeetingAttendanceReportId, id.AttendanceRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Online Meeting Id Attendance Report Id Attendance Record ID
func (id UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("attendanceReports", "attendanceReports", "attendanceReports"),
		resourceids.UserSpecifiedSegment("meetingAttendanceReportId", "meetingAttendanceReportId"),
		resourceids.StaticSegment("attendanceRecords", "attendanceRecords", "attendanceRecords"),
		resourceids.UserSpecifiedSegment("attendanceRecordId", "attendanceRecordId"),
	}
}

// String returns a human-readable description of this User Id Online Meeting Id Attendance Report Id Attendance Record ID
func (id UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Attendance Report: %q", id.MeetingAttendanceReportId),
		fmt.Sprintf("Attendance Record: %q", id.AttendanceRecordId),
	}
	return fmt.Sprintf("User Id Online Meeting Id Attendance Report Id Attendance Record (%s)", strings.Join(components, "\n"))
}
