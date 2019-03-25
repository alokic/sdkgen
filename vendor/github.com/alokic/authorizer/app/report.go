package app

import (
	"context"

	"github.com/alokic/authorizer/app/proto"
)

// Report is domain object.
type Report proto.Report

// Reports is list of domain objects.
type Reports []*Report

// ReportRepository is interface for services which can be pinged for availibility
type ReportRepository interface {
	Ping(context.Context, ...interface{}) (*Report, error)
}

// ReportHealthRequest is definition of health request.
type ReportHealthRequest proto.ReportHealthRequest

// ReportHealthReply is definition of health reply.
type ReportHealthReply proto.ReportHealthReply

// ReportService provides interface for report service.
type ReportService interface {
	Health(context.Context, *ReportHealthRequest) (*ReportHealthReply, error)
}

// ToProto converts report domain object to protobuf definition.
func (s *Report) ToProto() *proto.Report {
	report := proto.Report(*s)
	return &report
}

// ToProto converts list of report domain object to protobuf definition.
func (roles Reports) ToProto() []*proto.Report {
	var pr []*proto.Report
	for _, r := range roles {
		pr = append(pr, r.ToProto())
	}
	return pr
}
