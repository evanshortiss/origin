// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package pinpointiface provides an interface to enable mocking the Amazon Pinpoint service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package pinpointiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/pinpoint"
)

// PinpointAPI provides an interface to enable mocking the
// pinpoint.Pinpoint service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Pinpoint.
//    func myFunc(svc pinpointiface.PinpointAPI) bool {
//        // Make svc.CreateCampaign request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := pinpoint.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPinpointClient struct {
//        pinpointiface.PinpointAPI
//    }
//    func (m *mockPinpointClient) CreateCampaign(input *pinpoint.CreateCampaignInput) (*pinpoint.CreateCampaignOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPinpointClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type PinpointAPI interface {
	CreateCampaignRequest(*pinpoint.CreateCampaignInput) (*request.Request, *pinpoint.CreateCampaignOutput)

	CreateCampaign(*pinpoint.CreateCampaignInput) (*pinpoint.CreateCampaignOutput, error)

	CreateImportJobRequest(*pinpoint.CreateImportJobInput) (*request.Request, *pinpoint.CreateImportJobOutput)

	CreateImportJob(*pinpoint.CreateImportJobInput) (*pinpoint.CreateImportJobOutput, error)

	CreateSegmentRequest(*pinpoint.CreateSegmentInput) (*request.Request, *pinpoint.CreateSegmentOutput)

	CreateSegment(*pinpoint.CreateSegmentInput) (*pinpoint.CreateSegmentOutput, error)

	DeleteApnsChannelRequest(*pinpoint.DeleteApnsChannelInput) (*request.Request, *pinpoint.DeleteApnsChannelOutput)

	DeleteApnsChannel(*pinpoint.DeleteApnsChannelInput) (*pinpoint.DeleteApnsChannelOutput, error)

	DeleteCampaignRequest(*pinpoint.DeleteCampaignInput) (*request.Request, *pinpoint.DeleteCampaignOutput)

	DeleteCampaign(*pinpoint.DeleteCampaignInput) (*pinpoint.DeleteCampaignOutput, error)

	DeleteGcmChannelRequest(*pinpoint.DeleteGcmChannelInput) (*request.Request, *pinpoint.DeleteGcmChannelOutput)

	DeleteGcmChannel(*pinpoint.DeleteGcmChannelInput) (*pinpoint.DeleteGcmChannelOutput, error)

	DeleteSegmentRequest(*pinpoint.DeleteSegmentInput) (*request.Request, *pinpoint.DeleteSegmentOutput)

	DeleteSegment(*pinpoint.DeleteSegmentInput) (*pinpoint.DeleteSegmentOutput, error)

	GetApnsChannelRequest(*pinpoint.GetApnsChannelInput) (*request.Request, *pinpoint.GetApnsChannelOutput)

	GetApnsChannel(*pinpoint.GetApnsChannelInput) (*pinpoint.GetApnsChannelOutput, error)

	GetApplicationSettingsRequest(*pinpoint.GetApplicationSettingsInput) (*request.Request, *pinpoint.GetApplicationSettingsOutput)

	GetApplicationSettings(*pinpoint.GetApplicationSettingsInput) (*pinpoint.GetApplicationSettingsOutput, error)

	GetCampaignRequest(*pinpoint.GetCampaignInput) (*request.Request, *pinpoint.GetCampaignOutput)

	GetCampaign(*pinpoint.GetCampaignInput) (*pinpoint.GetCampaignOutput, error)

	GetCampaignActivitiesRequest(*pinpoint.GetCampaignActivitiesInput) (*request.Request, *pinpoint.GetCampaignActivitiesOutput)

	GetCampaignActivities(*pinpoint.GetCampaignActivitiesInput) (*pinpoint.GetCampaignActivitiesOutput, error)

	GetCampaignVersionRequest(*pinpoint.GetCampaignVersionInput) (*request.Request, *pinpoint.GetCampaignVersionOutput)

	GetCampaignVersion(*pinpoint.GetCampaignVersionInput) (*pinpoint.GetCampaignVersionOutput, error)

	GetCampaignVersionsRequest(*pinpoint.GetCampaignVersionsInput) (*request.Request, *pinpoint.GetCampaignVersionsOutput)

	GetCampaignVersions(*pinpoint.GetCampaignVersionsInput) (*pinpoint.GetCampaignVersionsOutput, error)

	GetCampaignsRequest(*pinpoint.GetCampaignsInput) (*request.Request, *pinpoint.GetCampaignsOutput)

	GetCampaigns(*pinpoint.GetCampaignsInput) (*pinpoint.GetCampaignsOutput, error)

	GetEndpointRequest(*pinpoint.GetEndpointInput) (*request.Request, *pinpoint.GetEndpointOutput)

	GetEndpoint(*pinpoint.GetEndpointInput) (*pinpoint.GetEndpointOutput, error)

	GetGcmChannelRequest(*pinpoint.GetGcmChannelInput) (*request.Request, *pinpoint.GetGcmChannelOutput)

	GetGcmChannel(*pinpoint.GetGcmChannelInput) (*pinpoint.GetGcmChannelOutput, error)

	GetImportJobRequest(*pinpoint.GetImportJobInput) (*request.Request, *pinpoint.GetImportJobOutput)

	GetImportJob(*pinpoint.GetImportJobInput) (*pinpoint.GetImportJobOutput, error)

	GetImportJobsRequest(*pinpoint.GetImportJobsInput) (*request.Request, *pinpoint.GetImportJobsOutput)

	GetImportJobs(*pinpoint.GetImportJobsInput) (*pinpoint.GetImportJobsOutput, error)

	GetSegmentRequest(*pinpoint.GetSegmentInput) (*request.Request, *pinpoint.GetSegmentOutput)

	GetSegment(*pinpoint.GetSegmentInput) (*pinpoint.GetSegmentOutput, error)

	GetSegmentImportJobsRequest(*pinpoint.GetSegmentImportJobsInput) (*request.Request, *pinpoint.GetSegmentImportJobsOutput)

	GetSegmentImportJobs(*pinpoint.GetSegmentImportJobsInput) (*pinpoint.GetSegmentImportJobsOutput, error)

	GetSegmentVersionRequest(*pinpoint.GetSegmentVersionInput) (*request.Request, *pinpoint.GetSegmentVersionOutput)

	GetSegmentVersion(*pinpoint.GetSegmentVersionInput) (*pinpoint.GetSegmentVersionOutput, error)

	GetSegmentVersionsRequest(*pinpoint.GetSegmentVersionsInput) (*request.Request, *pinpoint.GetSegmentVersionsOutput)

	GetSegmentVersions(*pinpoint.GetSegmentVersionsInput) (*pinpoint.GetSegmentVersionsOutput, error)

	GetSegmentsRequest(*pinpoint.GetSegmentsInput) (*request.Request, *pinpoint.GetSegmentsOutput)

	GetSegments(*pinpoint.GetSegmentsInput) (*pinpoint.GetSegmentsOutput, error)

	UpdateApnsChannelRequest(*pinpoint.UpdateApnsChannelInput) (*request.Request, *pinpoint.UpdateApnsChannelOutput)

	UpdateApnsChannel(*pinpoint.UpdateApnsChannelInput) (*pinpoint.UpdateApnsChannelOutput, error)

	UpdateApplicationSettingsRequest(*pinpoint.UpdateApplicationSettingsInput) (*request.Request, *pinpoint.UpdateApplicationSettingsOutput)

	UpdateApplicationSettings(*pinpoint.UpdateApplicationSettingsInput) (*pinpoint.UpdateApplicationSettingsOutput, error)

	UpdateCampaignRequest(*pinpoint.UpdateCampaignInput) (*request.Request, *pinpoint.UpdateCampaignOutput)

	UpdateCampaign(*pinpoint.UpdateCampaignInput) (*pinpoint.UpdateCampaignOutput, error)

	UpdateEndpointRequest(*pinpoint.UpdateEndpointInput) (*request.Request, *pinpoint.UpdateEndpointOutput)

	UpdateEndpoint(*pinpoint.UpdateEndpointInput) (*pinpoint.UpdateEndpointOutput, error)

	UpdateEndpointsBatchRequest(*pinpoint.UpdateEndpointsBatchInput) (*request.Request, *pinpoint.UpdateEndpointsBatchOutput)

	UpdateEndpointsBatch(*pinpoint.UpdateEndpointsBatchInput) (*pinpoint.UpdateEndpointsBatchOutput, error)

	UpdateGcmChannelRequest(*pinpoint.UpdateGcmChannelInput) (*request.Request, *pinpoint.UpdateGcmChannelOutput)

	UpdateGcmChannel(*pinpoint.UpdateGcmChannelInput) (*pinpoint.UpdateGcmChannelOutput, error)

	UpdateSegmentRequest(*pinpoint.UpdateSegmentInput) (*request.Request, *pinpoint.UpdateSegmentOutput)

	UpdateSegment(*pinpoint.UpdateSegmentInput) (*pinpoint.UpdateSegmentOutput, error)
}

var _ PinpointAPI = (*pinpoint.Pinpoint)(nil)