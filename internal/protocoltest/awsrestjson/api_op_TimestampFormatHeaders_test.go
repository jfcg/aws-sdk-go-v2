// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithyrand "github.com/awslabs/smithy-go/rand"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_TimestampFormatHeaders_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *TimestampFormatHeadersInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Tests how timestamp request headers are serialized
		"RestJsonTimestampFormatHeaders": {
			Params: &TimestampFormatHeadersInput{
				MemberEpochSeconds: ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				MemberHttpDate:     ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				MemberDateTime:     ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				DefaultFormat:      ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				TargetEpochSeconds: ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				TargetHttpDate:     ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				TargetDateTime:     ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/TimestampFormatHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-defaultFormat":      []string{"Mon, 16 Dec 2019 23:48:18 GMT"},
				"X-memberDateTime":     []string{"2019-12-16T23:48:18Z"},
				"X-memberEpochSeconds": []string{"1576540098"},
				"X-memberHttpDate":     []string{"Mon, 16 Dec 2019 23:48:18 GMT"},
				"X-targetDateTime":     []string{"2019-12-16T23:48:18Z"},
				"X-targetEpochSeconds": []string{"1576540098"},
				"X-targetHttpDate":     []string{"Mon, 16 Dec 2019 23:48:18 GMT"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			url := server.URL
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               aws.NewBuildableHTTPClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.TimestampFormatHeaders(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_TimestampFormatHeaders_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *TimestampFormatHeadersOutput
	}{
		// Tests how timestamp response headers are serialized
		"RestJsonTimestampFormatHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-defaultFormat":      []string{"Mon, 16 Dec 2019 23:48:18 GMT"},
				"X-memberDateTime":     []string{"2019-12-16T23:48:18Z"},
				"X-memberEpochSeconds": []string{"1576540098"},
				"X-memberHttpDate":     []string{"Mon, 16 Dec 2019 23:48:18 GMT"},
				"X-targetDateTime":     []string{"2019-12-16T23:48:18Z"},
				"X-targetEpochSeconds": []string{"1576540098"},
				"X-targetHttpDate":     []string{"Mon, 16 Dec 2019 23:48:18 GMT"},
			},
			Body: []byte(``),
			ExpectResult: &TimestampFormatHeadersOutput{
				MemberEpochSeconds: ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				MemberHttpDate:     ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				MemberDateTime:     ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				DefaultFormat:      ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				TargetEpochSeconds: ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				TargetHttpDate:     ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				TargetDateTime:     ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params TimestampFormatHeadersInput
			result, err := client.TimestampFormatHeaders(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
