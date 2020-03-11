// Copyright (c) 2019 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by generate-adapter. DO NOT EDIT.

package adapter

import (
	"github.com/gogo/protobuf/types"
	"go.temporal.io/temporal-proto/common"
	"go.temporal.io/temporal-proto/enums"

	"github.com/temporalio/temporal/.gen/go/shared"
)

func ToProtoBool(in *bool) *types.BoolValue {
	if in == nil {
		return nil
	}

	return &types.BoolValue{Value: *in}
}

func ToProtoDouble(in *float64) *types.DoubleValue {
	if in == nil {
		return nil
	}

	return &types.DoubleValue{Value: *in}
}

func ToThriftBool(in *types.BoolValue) *bool {
	if in == nil {
		return nil
	}

	return &in.Value
}

func ToThriftDouble(in *types.DoubleValue) *float64 {
	if in == nil {
		return nil
	}

	return &in.Value
}

func ToProtoDomainInfo(in *shared.DomainInfo) *common.DomainInfo {
	if in == nil {
		return nil
	}
	return &common.DomainInfo{
		Name:        in.GetName(),
		Status:      enums.DomainStatus(in.GetStatus()),
		Description: in.GetDescription(),
		OwnerEmail:  in.GetOwnerEmail(),
		Data:        in.GetData(),
		Uuid:        in.GetUUID(),
	}
}

func ToProtoDomainReplicationConfiguration(in *shared.DomainReplicationConfiguration) *common.DomainReplicationConfiguration {
	if in == nil {
		return nil
	}
	return &common.DomainReplicationConfiguration{
		ActiveClusterName: in.GetActiveClusterName(),
		Clusters:          ToProtoClusterReplicationConfigurations(in.GetClusters()),
	}
}

func ToProtoDomainConfiguration(in *shared.DomainConfiguration) *common.DomainConfiguration {
	if in == nil {
		return nil
	}
	return &common.DomainConfiguration{
		WorkflowExecutionRetentionPeriodInDays: in.GetWorkflowExecutionRetentionPeriodInDays(),
		EmitMetric:                             ToProtoBool(in.EmitMetric),
		BadBinaries:                            ToProtoBadBinariesPtr(in.GetBadBinaries()),
		HistoryArchivalStatus:                  ToProtoArchivalStatus(in.HistoryArchivalStatus),
		HistoryArchivalURI:                     in.GetHistoryArchivalURI(),
		VisibilityArchivalStatus:               ToProtoArchivalStatus(in.VisibilityArchivalStatus),
		VisibilityArchivalURI:                  in.GetVisibilityArchivalURI(),
	}
}

func ToProtoBadBinariesPtr(in *shared.BadBinaries) *common.BadBinaries {
	if in == nil {
		return nil
	}

	ret := ToProtoBadBinaries(*in)
	return &ret
}

func ToProtoBadBinaries(in shared.BadBinaries) common.BadBinaries {
	ret := make(map[string]*common.BadBinaryInfo, len(in.GetBinaries()))

	for key, value := range in.GetBinaries() {
		ret[key] = ToProtoBadBinaryInfo(value)
	}

	return common.BadBinaries{
		Binaries: ret,
	}
}

func ToProtoBadBinaryInfo(in *shared.BadBinaryInfo) *common.BadBinaryInfo {
	if in == nil {
		return nil
	}
	return &common.BadBinaryInfo{
		Reason:          in.GetReason(),
		Operator:        in.GetOperator(),
		CreatedTimeNano: in.GetCreatedTimeNano(),
	}
}

func ToThriftClusterReplicationConfigurations(in []*common.ClusterReplicationConfiguration) []*shared.ClusterReplicationConfiguration {
	var ret []*shared.ClusterReplicationConfiguration
	for _, cluster := range in {
		ret = append(ret, &shared.ClusterReplicationConfiguration{ClusterName: &cluster.ClusterName})
	}

	return ret
}

func ToProtoClusterReplicationConfigurations(in []*shared.ClusterReplicationConfiguration) []*common.ClusterReplicationConfiguration {
	var ret []*common.ClusterReplicationConfiguration
	for _, cluster := range in {
		ret = append(ret, &common.ClusterReplicationConfiguration{ClusterName: *cluster.ClusterName})
	}

	return ret
}

func ToThriftUpdateDomainInfo(in *common.UpdateDomainInfo) *shared.UpdateDomainInfo {
	if in == nil {
		return nil
	}
	return &shared.UpdateDomainInfo{
		Description: &in.Description,
		OwnerEmail:  &in.OwnerEmail,
		Data:        in.Data,
	}
}
func ToThriftDomainConfiguration(in *common.DomainConfiguration) *shared.DomainConfiguration {
	if in == nil {
		return nil
	}
	return &shared.DomainConfiguration{
		WorkflowExecutionRetentionPeriodInDays: &in.WorkflowExecutionRetentionPeriodInDays,
		EmitMetric:                             ToThriftBool(in.EmitMetric),
		BadBinaries:                            ToThriftBadBinariesPtr(in.BadBinaries),
		HistoryArchivalStatus:                  ToThriftArchivalStatus(in.HistoryArchivalStatus),
		HistoryArchivalURI:                     &in.HistoryArchivalURI,
		VisibilityArchivalStatus:               ToThriftArchivalStatus(in.VisibilityArchivalStatus),
		VisibilityArchivalURI:                  &in.VisibilityArchivalURI,
	}
}
func ToThriftDomainReplicationConfiguration(in *common.DomainReplicationConfiguration) *shared.DomainReplicationConfiguration {
	if in == nil {
		return nil
	}
	return &shared.DomainReplicationConfiguration{
		ActiveClusterName: &in.ActiveClusterName,
		Clusters:          ToThriftClusterReplicationConfigurations(in.Clusters),
	}
}

func ToThriftBadBinariesPtr(in *common.BadBinaries) *shared.BadBinaries {
	if in == nil {
		return nil
	}

	ret := ToThriftBadBinaries(*in)
	return &ret
}

// ToThriftBadBinaries ...
func ToThriftBadBinaries(in common.BadBinaries) shared.BadBinaries {
	ret := make(map[string]*shared.BadBinaryInfo, len(in.Binaries))

	for key, value := range in.Binaries {
		ret[key] = ToThriftBadBinaryInfo(value)
	}

	return shared.BadBinaries{
		Binaries: ret,
	}
}

func ToThriftBadBinaryInfo(in *common.BadBinaryInfo) *shared.BadBinaryInfo {
	if in == nil {
		return nil
	}
	return &shared.BadBinaryInfo{
		Reason:          &in.Reason,
		Operator:        &in.Operator,
		CreatedTimeNano: &in.CreatedTimeNano,
	}
}

func ToThriftWorkflowType(in *common.WorkflowType) *shared.WorkflowType {
	if in == nil {
		return nil
	}
	return &shared.WorkflowType{
		Name: &in.Name,
	}
}

func ToProtoTaskList(in *shared.TaskList) *common.TaskList {
	if in == nil {
		return nil
	}
	return &common.TaskList{
		Name: in.GetName(),
		Kind: enums.TaskListKind(in.GetKind()),
	}
}

func ToProtoRetryPolicy(in *shared.RetryPolicy) *common.RetryPolicy {
	if in == nil {
		return nil
	}
	return &common.RetryPolicy{
		InitialIntervalInSeconds:    in.GetInitialIntervalInSeconds(),
		BackoffCoefficient:          in.GetBackoffCoefficient(),
		MaximumIntervalInSeconds:    in.GetMaximumIntervalInSeconds(),
		MaximumAttempts:             in.GetMaximumAttempts(),
		NonRetriableErrorReasons:    in.GetNonRetriableErrorReasons(),
		ExpirationIntervalInSeconds: in.GetExpirationIntervalInSeconds(),
	}
}

func ToProtoMemo(in *shared.Memo) *common.Memo {
	if in == nil {
		return nil
	}
	return &common.Memo{
		Fields: in.GetFields(),
	}
}

func ToProtoResetPoints(in *shared.ResetPoints) *common.ResetPoints {
	if in == nil {
		return nil
	}

	return &common.ResetPoints{
		Points: ToProtoResetPointInfos(in.GetPoints()),
	}
}

func ToProtoResetPointInfos(in []*shared.ResetPointInfo) []*common.ResetPointInfo {
	if in == nil {
		return nil
	}
	var points []*common.ResetPointInfo
	for _, point := range in {
		points = append(points, ToProtoResetPointInfo(point))
	}

	return points
}

func ToProtoHeader(in *shared.Header) *common.Header {
	if in == nil {
		return nil
	}
	return &common.Header{
		Fields: in.GetFields(),
	}
}

func ToProtoActivityType(in *shared.ActivityType) *common.ActivityType {
	if in == nil {
		return nil
	}
	return &common.ActivityType{
		Name: in.GetName(),
	}
}

func ToProtoResetPointInfo(in *shared.ResetPointInfo) *common.ResetPointInfo {
	if in == nil {
		return nil
	}
	return &common.ResetPointInfo{
		BinaryChecksum:           in.GetBinaryChecksum(),
		RunId:                    in.GetRunId(),
		FirstDecisionCompletedId: in.GetFirstDecisionCompletedId(),
		CreatedTimeNano:          in.GetCreatedTimeNano(),
		ExpiringTimeNano:         in.GetExpiringTimeNano(),
		Resettable:               in.GetResettable(),
	}
}

func ToProtoWorkflowQuery(in *shared.WorkflowQuery) *common.WorkflowQuery {
	if in == nil {
		return nil
	}
	return &common.WorkflowQuery{
		QueryType: in.GetQueryType(),
		QueryArgs: in.GetQueryArgs(),
	}
}

// ToThriftResetPoints ...
func ToThriftResetPoints(in *common.ResetPoints) *shared.ResetPoints {
	if in == nil {
		return nil
	}

	return &shared.ResetPoints{
		Points: ToThriftResetPointInfos(in.Points),
	}
}

func ToThriftResetPointInfos(in []*common.ResetPointInfo) []*shared.ResetPointInfo {
	if in == nil {
		return nil
	}

	var ret []*shared.ResetPointInfo
	for _, item := range in {
		ret = append(ret, ToThriftResetPointInfo(item))
	}
	return ret
}

func ToThriftResetPointInfo(in *common.ResetPointInfo) *shared.ResetPointInfo {
	if in == nil {
		return nil
	}

	return &shared.ResetPointInfo{
		BinaryChecksum:           &in.BinaryChecksum,
		RunId:                    &in.RunId,
		FirstDecisionCompletedId: &in.FirstDecisionCompletedId,
		CreatedTimeNano:          &in.CreatedTimeNano,
		ExpiringTimeNano:         &in.ExpiringTimeNano,
		Resettable:               &in.Resettable,
	}
}
