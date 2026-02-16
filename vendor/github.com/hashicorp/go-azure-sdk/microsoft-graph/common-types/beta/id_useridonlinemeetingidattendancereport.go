package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdAttendanceReportId{}

// UserIdOnlineMeetingIdAttendanceReportId is a struct representing the Resource ID for a User Id Online Meeting Id Attendance Report
type UserIdOnlineMeetingIdAttendanceReportId struct {
	UserId                    string
	OnlineMeetingId           string
	MeetingAttendanceReportId string
}

// NewUserIdOnlineMeetingIdAttendanceReportID returns a new UserIdOnlineMeetingIdAttendanceReportId struct
func NewUserIdOnlineMeetingIdAttendanceReportID(userId string, onlineMeetingId string, meetingAttendanceReportId string) UserIdOnlineMeetingIdAttendanceReportId {
	return UserIdOnlineMeetingIdAttendanceReportId{
		UserId:                    userId,
		OnlineMeetingId:           onlineMeetingId,
		MeetingAttendanceReportId: meetingAttendanceReportId,
	}
}

// ParseUserIdOnlineMeetingIdAttendanceReportID parses 'input' into a UserIdOnlineMeetingIdAttendanceReportId
func ParseUserIdOnlineMeetingIdAttendanceReportID(input string) (*UserIdOnlineMeetingIdAttendanceReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdAttendanceReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdAttendanceReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnlineMeetingIdAttendanceReportIDInsensitively parses 'input' case-insensitively into a UserIdOnlineMeetingIdAttendanceReportId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnlineMeetingIdAttendanceReportIDInsensitively(input string) (*UserIdOnlineMeetingIdAttendanceReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdAttendanceReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdAttendanceReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnlineMeetingIdAttendanceReportId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdOnlineMeetingIdAttendanceReportID checks that 'input' can be parsed as a User Id Online Meeting Id Attendance Report ID
func ValidateUserIdOnlineMeetingIdAttendanceReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnlineMeetingIdAttendanceReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Online Meeting Id Attendance Report ID
func (id UserIdOnlineMeetingIdAttendanceReportId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/attendanceReports/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.MeetingAttendanceReportId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Online Meeting Id Attendance Report ID
func (id UserIdOnlineMeetingIdAttendanceReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("attendanceReports", "attendanceReports", "attendanceReports"),
		resourceids.UserSpecifiedSegment("meetingAttendanceReportId", "meetingAttendanceReportId"),
	}
}

// String returns a human-readable description of this User Id Online Meeting Id Attendance Report ID
func (id UserIdOnlineMeetingIdAttendanceReportId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Attendance Report: %q", id.MeetingAttendanceReportId),
	}
	return fmt.Sprintf("User Id Online Meeting Id Attendance Report (%s)", strings.Join(components, "\n"))
}
