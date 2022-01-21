package gateway

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
)

const (
	requestID     = "x-request-id"
	traceID       = "x-b3-traceid"
	spanID        = "x-b3-spanid"
	parentspanID  = "x-b3-parentspanid"
	sampled       = "x-b3-sampled"
	flags         = "x-b3-flags"
	otSpanContext = "x-ot-span-context"
)

var otHeaders = []string{
	requestID,
	traceID,
	spanID,
	parentspanID,
	sampled,
	flags,
	otSpanContext,
}

func injectHeadersIntoMetadata(ctx context.Context, req *http.Request) metadata.MD {
	pairs := []string{}
	for _, h := range otHeaders {
		if v := req.Header.Get(h); len(v) > 0 {
			pairs = append(pairs, h, v)
		}
	}
	return metadata.Pairs(pairs...)
}

type annotator func(context.Context, *http.Request) metadata.MD

func chainGrpcAnnotators(annotators ...annotator) annotator {
	return func(c context.Context, r *http.Request) metadata.MD {
		mds := []metadata.MD{}
		for _, a := range annotators {
			mds = append(mds, a(c, r))
		}
		return metadata.Join(mds...)
	}
}

var PropagateTracingHeader = chainGrpcAnnotators([]annotator{injectHeadersIntoMetadata}...)
