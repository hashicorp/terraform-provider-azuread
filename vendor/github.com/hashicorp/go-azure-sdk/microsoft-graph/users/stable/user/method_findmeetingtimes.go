package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FindMeetingTimesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.MeetingTimeSuggestionsResult
}

type FindMeetingTimesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultFindMeetingTimesOperationOptions() FindMeetingTimesOperationOptions {
	return FindMeetingTimesOperationOptions{}
}

func (o FindMeetingTimesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o FindMeetingTimesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o FindMeetingTimesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// FindMeetingTimes - Invoke action findMeetingTimes. Suggest meeting times and locations based on organizer and
// attendee availability, and time or location constraints specified as parameters. If findMeetingTimes cannot return
// any meeting suggestions, the response would indicate a reason in the emptySuggestionsReason property. Based on this
// value, you can better adjust the parameters and call findMeetingTimes again. The algorithm used to suggest meeting
// times and locations undergoes fine-tuning from time to time. In scenarios like test environments where the input
// parameters and calendar data remain static, expect that the suggested results may differ over time.
func (c UserClient) FindMeetingTimes(ctx context.Context, id stable.UserId, input FindMeetingTimesRequest, options FindMeetingTimesOperationOptions) (result FindMeetingTimesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/findMeetingTimes", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model stable.MeetingTimeSuggestionsResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
