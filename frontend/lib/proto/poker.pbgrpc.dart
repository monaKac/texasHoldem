// This is a generated file - do not edit.
//
// Generated from poker.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_relative_imports

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'poker.pb.dart' as $0;

export 'poker.pb.dart';

@$pb.GrpcServiceName('poker.PokerService')
class PokerServiceClient extends $grpc.Client {
  /// The hostname for this service.
  static const $core.String defaultHost = '';

  /// OAuth scopes needed for the client.
  static const $core.List<$core.String> oauthScopes = [
    '',
  ];

  PokerServiceClient(super.channel, {super.options, super.interceptors});

  $grpc.ResponseFuture<$0.EvaluateHandResponse> evaluateHand(
    $0.EvaluateHandRequest request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$evaluateHand, request, options: options);
  }

  $grpc.ResponseFuture<$0.CompareHandsResponse> compareHands(
    $0.CompareHandsRequest request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$compareHands, request, options: options);
  }

  $grpc.ResponseFuture<$0.WinProbabilityResponse> calculateWinProbability(
    $0.WinProbabilityRequest request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$calculateWinProbability, request,
        options: options);
  }

  // method descriptors

  static final _$evaluateHand =
      $grpc.ClientMethod<$0.EvaluateHandRequest, $0.EvaluateHandResponse>(
          '/poker.PokerService/EvaluateHand',
          ($0.EvaluateHandRequest value) => value.writeToBuffer(),
          $0.EvaluateHandResponse.fromBuffer);
  static final _$compareHands =
      $grpc.ClientMethod<$0.CompareHandsRequest, $0.CompareHandsResponse>(
          '/poker.PokerService/CompareHands',
          ($0.CompareHandsRequest value) => value.writeToBuffer(),
          $0.CompareHandsResponse.fromBuffer);
  static final _$calculateWinProbability =
      $grpc.ClientMethod<$0.WinProbabilityRequest, $0.WinProbabilityResponse>(
          '/poker.PokerService/CalculateWinProbability',
          ($0.WinProbabilityRequest value) => value.writeToBuffer(),
          $0.WinProbabilityResponse.fromBuffer);
}

@$pb.GrpcServiceName('poker.PokerService')
abstract class PokerServiceBase extends $grpc.Service {
  $core.String get $name => 'poker.PokerService';

  PokerServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$0.EvaluateHandRequest, $0.EvaluateHandResponse>(
            'EvaluateHand',
            evaluateHand_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.EvaluateHandRequest.fromBuffer(value),
            ($0.EvaluateHandResponse value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$0.CompareHandsRequest, $0.CompareHandsResponse>(
            'CompareHands',
            compareHands_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.CompareHandsRequest.fromBuffer(value),
            ($0.CompareHandsResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.WinProbabilityRequest,
            $0.WinProbabilityResponse>(
        'CalculateWinProbability',
        calculateWinProbability_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.WinProbabilityRequest.fromBuffer(value),
        ($0.WinProbabilityResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.EvaluateHandResponse> evaluateHand_Pre(
      $grpc.ServiceCall $call,
      $async.Future<$0.EvaluateHandRequest> $request) async {
    return evaluateHand($call, await $request);
  }

  $async.Future<$0.EvaluateHandResponse> evaluateHand(
      $grpc.ServiceCall call, $0.EvaluateHandRequest request);

  $async.Future<$0.CompareHandsResponse> compareHands_Pre(
      $grpc.ServiceCall $call,
      $async.Future<$0.CompareHandsRequest> $request) async {
    return compareHands($call, await $request);
  }

  $async.Future<$0.CompareHandsResponse> compareHands(
      $grpc.ServiceCall call, $0.CompareHandsRequest request);

  $async.Future<$0.WinProbabilityResponse> calculateWinProbability_Pre(
      $grpc.ServiceCall $call,
      $async.Future<$0.WinProbabilityRequest> $request) async {
    return calculateWinProbability($call, await $request);
  }

  $async.Future<$0.WinProbabilityResponse> calculateWinProbability(
      $grpc.ServiceCall call, $0.WinProbabilityRequest request);
}
