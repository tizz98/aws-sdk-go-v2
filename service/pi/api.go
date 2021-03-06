// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pi

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

const opDescribeDimensionKeys = "DescribeDimensionKeys"

// DescribeDimensionKeysRequest is a API request type for the DescribeDimensionKeys API operation.
type DescribeDimensionKeysRequest struct {
	*aws.Request
	Input *DescribeDimensionKeysInput
	Copy  func(*DescribeDimensionKeysInput) DescribeDimensionKeysRequest
}

// Send marshals and sends the DescribeDimensionKeys API request.
func (r DescribeDimensionKeysRequest) Send() (*DescribeDimensionKeysOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*DescribeDimensionKeysOutput), nil
}

// DescribeDimensionKeysRequest returns a request value for making API operation for
// AWS Performance Insights.
//
// For a specific time period, retrieve the top N dimension keys for a metric.
//
//    // Example sending a request using the DescribeDimensionKeysRequest method.
//    req := client.DescribeDimensionKeysRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DescribeDimensionKeys
func (c *PI) DescribeDimensionKeysRequest(input *DescribeDimensionKeysInput) DescribeDimensionKeysRequest {
	op := &aws.Operation{
		Name:       opDescribeDimensionKeys,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDimensionKeysInput{}
	}

	output := &DescribeDimensionKeysOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return DescribeDimensionKeysRequest{Request: req, Input: input, Copy: c.DescribeDimensionKeysRequest}
}

const opGetResourceMetrics = "GetResourceMetrics"

// GetResourceMetricsRequest is a API request type for the GetResourceMetrics API operation.
type GetResourceMetricsRequest struct {
	*aws.Request
	Input *GetResourceMetricsInput
	Copy  func(*GetResourceMetricsInput) GetResourceMetricsRequest
}

// Send marshals and sends the GetResourceMetrics API request.
func (r GetResourceMetricsRequest) Send() (*GetResourceMetricsOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetResourceMetricsOutput), nil
}

// GetResourceMetricsRequest returns a request value for making API operation for
// AWS Performance Insights.
//
// Retrieve Performance Insights metrics for a set of data sources, over a time
// period. You can provide specific dimension groups and dimensions, and provide
// aggregation and filtering criteria for each group.
//
//    // Example sending a request using the GetResourceMetricsRequest method.
//    req := client.GetResourceMetricsRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/GetResourceMetrics
func (c *PI) GetResourceMetricsRequest(input *GetResourceMetricsInput) GetResourceMetricsRequest {
	op := &aws.Operation{
		Name:       opGetResourceMetrics,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetResourceMetricsInput{}
	}

	output := &GetResourceMetricsOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetResourceMetricsRequest{Request: req, Input: input, Copy: c.GetResourceMetricsRequest}
}

// A timestamp, and a single numerical value, which together represent a measurement
// at a particular point in time.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DataPoint
type DataPoint struct {
	_ struct{} `type:"structure"`

	// The time, in epoch format, associated with a particular Value.
	//
	// Timestamp is a required field
	Timestamp *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// The actual value associated with a particular Timestamp.
	//
	// Value is a required field
	Value *float64 `type:"double" required:"true"`
}

// String returns the string representation
func (s DataPoint) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DataPoint) GoString() string {
	return s.String()
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DescribeDimensionKeysRequest
type DescribeDimensionKeysInput struct {
	_ struct{} `type:"structure"`

	// The date and time specifying the end of the requested time series data. The
	// value specified is exclusive - data points less than (but not equal to) EndTime
	// will be returned.
	//
	// The value for EndTime must be later than the value for StartTime.
	//
	// EndTime is a required field
	EndTime *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// One or more filters to apply in the request. Restrictions:
	//
	//    * Any number of filters by the same dimension, as specified in the GroupBy
	//    or Partition parameters.
	//
	//    * A single filter for any other dimension in this dimension group.
	Filter map[string]string `type:"map"`

	// A specification for how to aggregate the data points from a query result.
	// You must specify a valid dimension group. Performance Insights will return
	// all of the dimensions within that group, unless you provide the names of
	// specific dimensions within that group. You can also request that Performance
	// Insights return a limited number of values for a dimension.
	//
	// GroupBy is a required field
	GroupBy *DimensionGroup `type:"structure" required:"true"`

	// An immutable, AWS Region-unique identifier for a data source. Performance
	// Insights gathers metrics from this data source.
	//
	// To use an Amazon RDS instance as a data source, you specify its DbiResourceId
	// value - for example: db-FAIHNTYBKTGAUSUZQYPDS2GW4A
	//
	// Identifier is a required field
	Identifier *string `type:"string" required:"true"`

	// The maximum number of items to return in the response. If more items exist
	// than the specified MaxRecords value, a pagination token is included in the
	// response so that the remaining results can be retrieved.
	MaxResults *int64 `type:"integer"`

	// The name of a Performance Insights metric to be measured.
	//
	// Valid values for Metric are:
	//
	//    * db.load.avg - a scaled representation of the number of active sessions
	//    for the database engine.
	//
	//    * db.sampledload.avg - the raw number of active sessions for the database
	//    engine.
	//
	// Metric is a required field
	Metric *string `type:"string" required:"true"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to
	// the value specified by MaxRecords.
	NextToken *string `type:"string"`

	// For each dimension specified in GroupBy, specify a secondary dimension to
	// further subdivide the partition keys in the response.
	PartitionBy *DimensionGroup `type:"structure"`

	// The granularity, in seconds, of the data points returned from Performance
	// Insights. A period can be as short as one second, or as long as one day (86400
	// seconds). Valid values are:
	//
	//    * 1 (one second)
	//
	//    * 60 (one minute)
	//
	//    * 300 (five minutes)
	//
	//    * 3600 (one hour)
	//
	//    * 86400 (twenty-four hours)
	//
	// If you don't specify PeriodInSeconds, then Performance Insights will choose
	// a value for you, with a goal of returning roughly 100-200 data points in
	// the response.
	PeriodInSeconds *int64 `type:"integer"`

	// The AWS service for which Performance Insights will return metrics. The only
	// valid value for ServiceType is: RDS
	//
	// ServiceType is a required field
	ServiceType ServiceType `type:"string" required:"true" enum:"true"`

	// The date and time specifying the beginning of the requested time series data.
	// You can't specify a StartTime that's earlier than 7 days ago. The value specified
	// is inclusive - data points equal to or greater than StartTime will be returned.
	//
	// The value for StartTime must be earlier than the value for EndTime.
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s DescribeDimensionKeysInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDimensionKeysInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDimensionKeysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDimensionKeysInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}

	if s.GroupBy == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupBy"))
	}

	if s.Identifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identifier"))
	}

	if s.Metric == nil {
		invalidParams.Add(aws.NewErrParamRequired("Metric"))
	}
	if len(s.ServiceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ServiceType"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}
	if s.GroupBy != nil {
		if err := s.GroupBy.Validate(); err != nil {
			invalidParams.AddNested("GroupBy", err.(aws.ErrInvalidParams))
		}
	}
	if s.PartitionBy != nil {
		if err := s.PartitionBy.Validate(); err != nil {
			invalidParams.AddNested("PartitionBy", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DescribeDimensionKeysResponse
type DescribeDimensionKeysOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The end time for the returned dimension keys, after alignment to a granular
	// boundary (as specified by PeriodInSeconds). AlignedEndTime will be greater
	// than or equal to the value of the user-specified Endtime.
	AlignedEndTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The start time for the returned dimension keys, after alignment to a granular
	// boundary (as specified by PeriodInSeconds). AlignedStartTime will be less
	// than or equal to the value of the user-specified StartTime.
	AlignedStartTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The dimension keys that were requested.
	Keys []DimensionKeyDescription `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to
	// the value specified by MaxRecords.
	NextToken *string `type:"string"`

	// If PartitionBy was present in the request, PartitionKeys contains the breakdown
	// of dimension keys by the specified partitions.
	PartitionKeys []ResponsePartitionKey `type:"list"`
}

// String returns the string representation
func (s DescribeDimensionKeysOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDimensionKeysOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s DescribeDimensionKeysOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// A logical grouping of Performance Insights metrics for a related subject
// area. For example, the db.sql dimension group consists of the following dimensions:
// db.sql.id, db.sql.db_id, db.sql.statement, and db.sql.tokenized_id.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DimensionGroup
type DimensionGroup struct {
	_ struct{} `type:"structure"`

	// A list of specific dimensions from a dimension group. If this parameter is
	// not present, then it signifies that all of the dimensions in the group were
	// requested, or are present in the response.
	//
	// Valid values for elements in the Dimensions array are:
	//
	//    * db.user.id
	//
	//    * db.user.name
	//
	//    * db.host.id
	//
	//    * db.host.name
	//
	//    * db.sql.id
	//
	//    * db.sql.db_id
	//
	//    * db.sql.statement
	//
	//    * db.sql.tokenized_id
	//
	//    * db.sql_tokenized.id
	//
	//    * db.sql_tokenized.db_id
	//
	//    * db.sql_tokenized.statement
	//
	//    * db.wait_event.name
	//
	//    * db.wait_event.type
	//
	//    * db.wait_event_type.name
	Dimensions []string `min:"1" type:"list"`

	// The name of the dimension group. Valid values are:
	//
	//    * db.user
	//
	//    * db.host
	//
	//    * db.sql
	//
	//    * db.sql_tokenized
	//
	//    * db.wait_event
	//
	//    * db.wait_event_type
	//
	// Group is a required field
	Group *string `type:"string" required:"true"`

	// The maximum number of items to fetch for this dimension group.
	Limit *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s DimensionGroup) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionGroup) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DimensionGroup) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DimensionGroup"}
	if s.Dimensions != nil && len(s.Dimensions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Dimensions", 1))
	}

	if s.Group == nil {
		invalidParams.Add(aws.NewErrParamRequired("Group"))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An array of descriptions and aggregated values for each dimension within
// a dimension group.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DimensionKeyDescription
type DimensionKeyDescription struct {
	_ struct{} `type:"structure"`

	// A map of name-value pairs for the dimensions in the group.
	Dimensions map[string]string `type:"map"`

	// If PartitionBy was specified, PartitionKeys contains the dimensions that
	// were.
	Partitions []float64 `type:"list"`

	// The aggregated metric value for the dimension(s), over the requested time
	// range.
	Total *float64 `type:"double"`
}

// String returns the string representation
func (s DimensionKeyDescription) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionKeyDescription) GoString() string {
	return s.String()
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/GetResourceMetricsRequest
type GetResourceMetricsInput struct {
	_ struct{} `type:"structure"`

	// The date and time specifiying the end of the requested time series data.
	// The value specified is exclusive - data points less than (but not equal to)
	// EndTime will be returned.
	//
	// The value for EndTime must be later than the value for StartTime.
	//
	// EndTime is a required field
	EndTime *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// An immutable, AWS Region-unique identifier for a data source. Performance
	// Insights gathers metrics from this data source.
	//
	// To use an Amazon RDS instance as a data source, you specify its DbiResourceId
	// value - for example: db-FAIHNTYBKTGAUSUZQYPDS2GW4A
	//
	// Identifier is a required field
	Identifier *string `type:"string" required:"true"`

	// The maximum number of items to return in the response. If more items exist
	// than the specified MaxRecords value, a pagination token is included in the
	// response so that the remaining results can be retrieved.
	MaxResults *int64 `type:"integer"`

	// An array of one or more queries to perform. Each query must specify a Performance
	// Insights metric, and can optionally specify aggregation and filtering criteria.
	//
	// MetricQueries is a required field
	MetricQueries []MetricQuery `min:"1" type:"list" required:"true"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to
	// the value specified by MaxRecords.
	NextToken *string `type:"string"`

	// The granularity, in seconds, of the data points returned from Performance
	// Insights. A period can be as short as one second, or as long as one day (86400
	// seconds). Valid values are:
	//
	//    * 1 (one second)
	//
	//    * 60 (one minute)
	//
	//    * 300 (five minutes)
	//
	//    * 3600 (one hour)
	//
	//    * 86400 (twenty-four hours)
	//
	// If you don't specify PeriodInSeconds, then Performance Insights will choose
	// a value for you, with a goal of returning roughly 100-200 data points in
	// the response.
	PeriodInSeconds *int64 `type:"integer"`

	// The AWS service for which Performance Insights will return metrics. The only
	// valid value for ServiceType is: RDS
	//
	// ServiceType is a required field
	ServiceType ServiceType `type:"string" required:"true" enum:"true"`

	// The date and time specifying the beginning of the requested time series data.
	// You can't specify a StartTime that's earlier than 7 days ago. The value specified
	// is inclusive - data points equal to or greater than StartTime will be returned.
	//
	// The value for StartTime must be earlier than the value for EndTime.
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s GetResourceMetricsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetResourceMetricsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourceMetricsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResourceMetricsInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}

	if s.Identifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identifier"))
	}

	if s.MetricQueries == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricQueries"))
	}
	if s.MetricQueries != nil && len(s.MetricQueries) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MetricQueries", 1))
	}
	if len(s.ServiceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ServiceType"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}
	if s.MetricQueries != nil {
		for i, v := range s.MetricQueries {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MetricQueries", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/GetResourceMetricsResponse
type GetResourceMetricsOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The end time for the returned metrics, after alignment to a granular boundary
	// (as specified by PeriodInSeconds). AlignedEndTime will be greater than or
	// equal to the value of the user-specified Endtime.
	AlignedEndTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The start time for the returned metrics, after alignment to a granular boundary
	// (as specified by PeriodInSeconds). AlignedStartTime will be less than or
	// equal to the value of the user-specified StartTime.
	AlignedStartTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// An immutable, AWS Region-unique identifier for a data source. Performance
	// Insights gathers metrics from this data source.
	//
	// To use an Amazon RDS instance as a data source, you specify its DbiResourceId
	// value - for example: db-FAIHNTYBKTGAUSUZQYPDS2GW4A
	Identifier *string `type:"string"`

	// An array of metric results,, where each array element contains all of the
	// data points for a particular dimension.
	MetricList []MetricKeyDataPoints `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to
	// the value specified by MaxRecords.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetResourceMetricsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetResourceMetricsOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetResourceMetricsOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// A time-ordered series of data points, correpsonding to a dimension of a Performance
// Insights metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/MetricKeyDataPoints
type MetricKeyDataPoints struct {
	_ struct{} `type:"structure"`

	// An array of timestamp-value pairs, representing measurements over a period
	// of time.
	DataPoints []DataPoint `type:"list"`

	// The dimension(s) to which the data points apply.
	Key *ResponseResourceMetricKey `type:"structure"`
}

// String returns the string representation
func (s MetricKeyDataPoints) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MetricKeyDataPoints) GoString() string {
	return s.String()
}

// A single query to be processed. You must provide the metric to query. If
// no other parameters are specified, Performance Insights returns all of the
// data points for that metric. You can optionally request that the data points
// be aggregated by dimension group ( GroupBy), and return only those data points
// that match your criteria (Filter).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/MetricQuery
type MetricQuery struct {
	_ struct{} `type:"structure"`

	// One or more filters to apply in the request. Restrictions:
	//
	//    * Any number of filters by the same dimension, as specified in the GroupBy
	//    parameter.
	//
	//    * A single filter for any other dimension in this dimension group.
	Filter map[string]string `type:"map"`

	// A specification for how to aggregate the data points from a query result.
	// You must specify a valid dimension group. Performance Insights will return
	// all of the dimensions within that group, unless you provide the names of
	// specific dimensions within that group. You can also request that Performance
	// Insights return a limited number of values for a dimension.
	GroupBy *DimensionGroup `type:"structure"`

	// The name of a Performance Insights metric to be measured.
	//
	// Valid values for Metric are:
	//
	//    * db.load.avg - a scaled representation of the number of active sessions
	//    for the database engine.
	//
	//    * db.sampledload.avg - the raw number of active sessions for the database
	//    engine.
	//
	// Metric is a required field
	Metric *string `type:"string" required:"true"`
}

// String returns the string representation
func (s MetricQuery) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MetricQuery) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricQuery) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MetricQuery"}

	if s.Metric == nil {
		invalidParams.Add(aws.NewErrParamRequired("Metric"))
	}
	if s.GroupBy != nil {
		if err := s.GroupBy.Validate(); err != nil {
			invalidParams.AddNested("GroupBy", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// If PartitionBy was specified in a DescribeDimensionKeys request, the dimensions
// are returned in an array. Each element in the array specifies one dimension.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/ResponsePartitionKey
type ResponsePartitionKey struct {
	_ struct{} `type:"structure"`

	// A dimension map that contains the dimension(s) for this partition.
	//
	// Dimensions is a required field
	Dimensions map[string]string `type:"map" required:"true"`
}

// String returns the string representation
func (s ResponsePartitionKey) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResponsePartitionKey) GoString() string {
	return s.String()
}

// An object describing a Performance Insights metric and one or more dimensions
// for that metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/ResponseResourceMetricKey
type ResponseResourceMetricKey struct {
	_ struct{} `type:"structure"`

	// The valid dimensions for the metric.
	Dimensions map[string]string `type:"map"`

	// The name of a Performance Insights metric to be measured.
	//
	// Valid values for Metric are:
	//
	//    * db.load.avg - a scaled representation of the number of active sessions
	//    for the database engine.
	//
	//    * db.sampledload.avg - the raw number of active sessions for the database
	//    engine.
	//
	// Metric is a required field
	Metric *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResponseResourceMetricKey) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResponseResourceMetricKey) GoString() string {
	return s.String()
}

type ServiceType string

// Enum values for ServiceType
const (
	ServiceTypeRds ServiceType = "RDS"
)

func (enum ServiceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ServiceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
