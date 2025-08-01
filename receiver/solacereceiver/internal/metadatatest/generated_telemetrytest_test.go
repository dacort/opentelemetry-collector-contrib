// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"go.opentelemetry.io/collector/component/componenttest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver/internal/metadata"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := componenttest.NewTelemetry()
	tb, err := metadata.NewTelemetryBuilder(testTel.NewTelemetrySettings())
	require.NoError(t, err)
	defer tb.Shutdown()
	tb.SolacereceiverDroppedEgressSpans.Add(context.Background(), 1)
	tb.SolacereceiverDroppedSpanMessages.Add(context.Background(), 1)
	tb.SolacereceiverFailedReconnections.Add(context.Background(), 1)
	tb.SolacereceiverFatalUnmarshallingErrors.Add(context.Background(), 1)
	tb.SolacereceiverNeedUpgrade.Record(context.Background(), 1)
	tb.SolacereceiverReceivedSpanMessages.Add(context.Background(), 1)
	tb.SolacereceiverReceiverFlowControlRecentRetries.Record(context.Background(), 1)
	tb.SolacereceiverReceiverFlowControlStatus.Record(context.Background(), 1)
	tb.SolacereceiverReceiverFlowControlTotal.Add(context.Background(), 1)
	tb.SolacereceiverReceiverFlowControlWithSingleSuccessfulRetry.Add(context.Background(), 1)
	tb.SolacereceiverReceiverStatus.Record(context.Background(), 1)
	tb.SolacereceiverRecoverableUnmarshallingErrors.Add(context.Background(), 1)
	tb.SolacereceiverReportedSpans.Add(context.Background(), 1)
	AssertEqualSolacereceiverDroppedEgressSpans(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverDroppedSpanMessages(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverFailedReconnections(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverFatalUnmarshallingErrors(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverNeedUpgrade(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverReceivedSpanMessages(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverReceiverFlowControlRecentRetries(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverReceiverFlowControlStatus(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverReceiverFlowControlTotal(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverReceiverFlowControlWithSingleSuccessfulRetry(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverReceiverStatus(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverRecoverableUnmarshallingErrors(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualSolacereceiverReportedSpans(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())

	require.NoError(t, testTel.Shutdown(context.Background()))
}
