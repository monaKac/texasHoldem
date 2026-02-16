// This is a generated file - do not edit.
//
// Generated from poker.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_relative_imports

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class EvaluateHandRequest extends $pb.GeneratedMessage {
  factory EvaluateHandRequest({
    $core.Iterable<$core.String>? cards,
  }) {
    final result = create();
    if (cards != null) result.cards.addAll(cards);
    return result;
  }

  EvaluateHandRequest._();

  factory EvaluateHandRequest.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory EvaluateHandRequest.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'EvaluateHandRequest',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'poker'),
      createEmptyInstance: create)
    ..pPS(1, _omitFieldNames ? '' : 'cards')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  EvaluateHandRequest clone() => deepCopy();
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  EvaluateHandRequest copyWith(void Function(EvaluateHandRequest) updates) =>
      super.copyWith((message) => updates(message as EvaluateHandRequest))
          as EvaluateHandRequest;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static EvaluateHandRequest create() => EvaluateHandRequest._();
  @$core.override
  EvaluateHandRequest createEmptyInstance() => create();
  @$core.pragma('dart2js:noInline')
  static EvaluateHandRequest getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<EvaluateHandRequest>(create);
  static EvaluateHandRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<$core.String> get cards => $_getList(0);
}

class EvaluateHandResponse extends $pb.GeneratedMessage {
  factory EvaluateHandResponse({
    $core.String? handRank,
    $core.Iterable<$core.String>? bestFive,
    $core.int? rankValue,
  }) {
    final result = create();
    if (handRank != null) result.handRank = handRank;
    if (bestFive != null) result.bestFive.addAll(bestFive);
    if (rankValue != null) result.rankValue = rankValue;
    return result;
  }

  EvaluateHandResponse._();

  factory EvaluateHandResponse.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory EvaluateHandResponse.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'EvaluateHandResponse',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'poker'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'handRank')
    ..pPS(2, _omitFieldNames ? '' : 'bestFive')
    ..aI(3, _omitFieldNames ? '' : 'rankValue')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  EvaluateHandResponse clone() => deepCopy();
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  EvaluateHandResponse copyWith(void Function(EvaluateHandResponse) updates) =>
      super.copyWith((message) => updates(message as EvaluateHandResponse))
          as EvaluateHandResponse;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static EvaluateHandResponse create() => EvaluateHandResponse._();
  @$core.override
  EvaluateHandResponse createEmptyInstance() => create();
  @$core.pragma('dart2js:noInline')
  static EvaluateHandResponse getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<EvaluateHandResponse>(create);
  static EvaluateHandResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get handRank => $_getSZ(0);
  @$pb.TagNumber(1)
  set handRank($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasHandRank() => $_has(0);
  @$pb.TagNumber(1)
  void clearHandRank() => $_clearField(1);

  @$pb.TagNumber(2)
  $pb.PbList<$core.String> get bestFive => $_getList(1);

  @$pb.TagNumber(3)
  $core.int get rankValue => $_getIZ(2);
  @$pb.TagNumber(3)
  set rankValue($core.int value) => $_setSignedInt32(2, value);
  @$pb.TagNumber(3)
  $core.bool hasRankValue() => $_has(2);
  @$pb.TagNumber(3)
  void clearRankValue() => $_clearField(3);
}

class CompareHandsRequest extends $pb.GeneratedMessage {
  factory CompareHandsRequest({
    $core.Iterable<$core.String>? player1Cards,
    $core.Iterable<$core.String>? player2Cards,
  }) {
    final result = create();
    if (player1Cards != null) result.player1Cards.addAll(player1Cards);
    if (player2Cards != null) result.player2Cards.addAll(player2Cards);
    return result;
  }

  CompareHandsRequest._();

  factory CompareHandsRequest.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory CompareHandsRequest.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'CompareHandsRequest',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'poker'),
      createEmptyInstance: create)
    ..pPS(1, _omitFieldNames ? '' : 'player1Cards')
    ..pPS(2, _omitFieldNames ? '' : 'player2Cards')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CompareHandsRequest clone() => deepCopy();
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CompareHandsRequest copyWith(void Function(CompareHandsRequest) updates) =>
      super.copyWith((message) => updates(message as CompareHandsRequest))
          as CompareHandsRequest;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CompareHandsRequest create() => CompareHandsRequest._();
  @$core.override
  CompareHandsRequest createEmptyInstance() => create();
  @$core.pragma('dart2js:noInline')
  static CompareHandsRequest getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<CompareHandsRequest>(create);
  static CompareHandsRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<$core.String> get player1Cards => $_getList(0);

  @$pb.TagNumber(2)
  $pb.PbList<$core.String> get player2Cards => $_getList(1);
}

class CompareHandsResponse extends $pb.GeneratedMessage {
  factory CompareHandsResponse({
    $core.int? winner,
    EvaluateHandResponse? player1Hand,
    EvaluateHandResponse? player2Hand,
    $core.String? description,
  }) {
    final result = create();
    if (winner != null) result.winner = winner;
    if (player1Hand != null) result.player1Hand = player1Hand;
    if (player2Hand != null) result.player2Hand = player2Hand;
    if (description != null) result.description = description;
    return result;
  }

  CompareHandsResponse._();

  factory CompareHandsResponse.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory CompareHandsResponse.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'CompareHandsResponse',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'poker'),
      createEmptyInstance: create)
    ..aI(1, _omitFieldNames ? '' : 'winner')
    ..aOM<EvaluateHandResponse>(2, _omitFieldNames ? '' : 'player1Hand',
        subBuilder: EvaluateHandResponse.create)
    ..aOM<EvaluateHandResponse>(3, _omitFieldNames ? '' : 'player2Hand',
        subBuilder: EvaluateHandResponse.create)
    ..aOS(4, _omitFieldNames ? '' : 'description')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CompareHandsResponse clone() => deepCopy();
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CompareHandsResponse copyWith(void Function(CompareHandsResponse) updates) =>
      super.copyWith((message) => updates(message as CompareHandsResponse))
          as CompareHandsResponse;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CompareHandsResponse create() => CompareHandsResponse._();
  @$core.override
  CompareHandsResponse createEmptyInstance() => create();
  @$core.pragma('dart2js:noInline')
  static CompareHandsResponse getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<CompareHandsResponse>(create);
  static CompareHandsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get winner => $_getIZ(0);
  @$pb.TagNumber(1)
  set winner($core.int value) => $_setSignedInt32(0, value);
  @$pb.TagNumber(1)
  $core.bool hasWinner() => $_has(0);
  @$pb.TagNumber(1)
  void clearWinner() => $_clearField(1);

  @$pb.TagNumber(2)
  EvaluateHandResponse get player1Hand => $_getN(1);
  @$pb.TagNumber(2)
  set player1Hand(EvaluateHandResponse value) => $_setField(2, value);
  @$pb.TagNumber(2)
  $core.bool hasPlayer1Hand() => $_has(1);
  @$pb.TagNumber(2)
  void clearPlayer1Hand() => $_clearField(2);
  @$pb.TagNumber(2)
  EvaluateHandResponse ensurePlayer1Hand() => $_ensure(1);

  @$pb.TagNumber(3)
  EvaluateHandResponse get player2Hand => $_getN(2);
  @$pb.TagNumber(3)
  set player2Hand(EvaluateHandResponse value) => $_setField(3, value);
  @$pb.TagNumber(3)
  $core.bool hasPlayer2Hand() => $_has(2);
  @$pb.TagNumber(3)
  void clearPlayer2Hand() => $_clearField(3);
  @$pb.TagNumber(3)
  EvaluateHandResponse ensurePlayer2Hand() => $_ensure(2);

  @$pb.TagNumber(4)
  $core.String get description => $_getSZ(3);
  @$pb.TagNumber(4)
  set description($core.String value) => $_setString(3, value);
  @$pb.TagNumber(4)
  $core.bool hasDescription() => $_has(3);
  @$pb.TagNumber(4)
  void clearDescription() => $_clearField(4);
}

class WinProbabilityRequest extends $pb.GeneratedMessage {
  factory WinProbabilityRequest({
    $core.Iterable<$core.String>? holeCards,
    $core.Iterable<$core.String>? community,
    $core.int? numOpponents,
    $core.int? iterations,
  }) {
    final result = create();
    if (holeCards != null) result.holeCards.addAll(holeCards);
    if (community != null) result.community.addAll(community);
    if (numOpponents != null) result.numOpponents = numOpponents;
    if (iterations != null) result.iterations = iterations;
    return result;
  }

  WinProbabilityRequest._();

  factory WinProbabilityRequest.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory WinProbabilityRequest.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'WinProbabilityRequest',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'poker'),
      createEmptyInstance: create)
    ..pPS(1, _omitFieldNames ? '' : 'holeCards')
    ..pPS(2, _omitFieldNames ? '' : 'community')
    ..aI(3, _omitFieldNames ? '' : 'numOpponents')
    ..aI(4, _omitFieldNames ? '' : 'iterations')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  WinProbabilityRequest clone() => deepCopy();
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  WinProbabilityRequest copyWith(
          void Function(WinProbabilityRequest) updates) =>
      super.copyWith((message) => updates(message as WinProbabilityRequest))
          as WinProbabilityRequest;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static WinProbabilityRequest create() => WinProbabilityRequest._();
  @$core.override
  WinProbabilityRequest createEmptyInstance() => create();
  @$core.pragma('dart2js:noInline')
  static WinProbabilityRequest getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<WinProbabilityRequest>(create);
  static WinProbabilityRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<$core.String> get holeCards => $_getList(0);

  @$pb.TagNumber(2)
  $pb.PbList<$core.String> get community => $_getList(1);

  @$pb.TagNumber(3)
  $core.int get numOpponents => $_getIZ(2);
  @$pb.TagNumber(3)
  set numOpponents($core.int value) => $_setSignedInt32(2, value);
  @$pb.TagNumber(3)
  $core.bool hasNumOpponents() => $_has(2);
  @$pb.TagNumber(3)
  void clearNumOpponents() => $_clearField(3);

  @$pb.TagNumber(4)
  $core.int get iterations => $_getIZ(3);
  @$pb.TagNumber(4)
  set iterations($core.int value) => $_setSignedInt32(3, value);
  @$pb.TagNumber(4)
  $core.bool hasIterations() => $_has(3);
  @$pb.TagNumber(4)
  void clearIterations() => $_clearField(4);
}

class WinProbabilityResponse extends $pb.GeneratedMessage {
  factory WinProbabilityResponse({
    $core.double? winProbability,
    $core.double? tieProbability,
    $core.double? lossProbability,
    $core.int? iterationsRun,
  }) {
    final result = create();
    if (winProbability != null) result.winProbability = winProbability;
    if (tieProbability != null) result.tieProbability = tieProbability;
    if (lossProbability != null) result.lossProbability = lossProbability;
    if (iterationsRun != null) result.iterationsRun = iterationsRun;
    return result;
  }

  WinProbabilityResponse._();

  factory WinProbabilityResponse.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory WinProbabilityResponse.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'WinProbabilityResponse',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'poker'),
      createEmptyInstance: create)
    ..aD(1, _omitFieldNames ? '' : 'winProbability')
    ..aD(2, _omitFieldNames ? '' : 'tieProbability')
    ..aD(3, _omitFieldNames ? '' : 'lossProbability')
    ..aI(4, _omitFieldNames ? '' : 'iterationsRun')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  WinProbabilityResponse clone() => deepCopy();
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  WinProbabilityResponse copyWith(
          void Function(WinProbabilityResponse) updates) =>
      super.copyWith((message) => updates(message as WinProbabilityResponse))
          as WinProbabilityResponse;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static WinProbabilityResponse create() => WinProbabilityResponse._();
  @$core.override
  WinProbabilityResponse createEmptyInstance() => create();
  @$core.pragma('dart2js:noInline')
  static WinProbabilityResponse getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<WinProbabilityResponse>(create);
  static WinProbabilityResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.double get winProbability => $_getN(0);
  @$pb.TagNumber(1)
  set winProbability($core.double value) => $_setDouble(0, value);
  @$pb.TagNumber(1)
  $core.bool hasWinProbability() => $_has(0);
  @$pb.TagNumber(1)
  void clearWinProbability() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.double get tieProbability => $_getN(1);
  @$pb.TagNumber(2)
  set tieProbability($core.double value) => $_setDouble(1, value);
  @$pb.TagNumber(2)
  $core.bool hasTieProbability() => $_has(1);
  @$pb.TagNumber(2)
  void clearTieProbability() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.double get lossProbability => $_getN(2);
  @$pb.TagNumber(3)
  set lossProbability($core.double value) => $_setDouble(2, value);
  @$pb.TagNumber(3)
  $core.bool hasLossProbability() => $_has(2);
  @$pb.TagNumber(3)
  void clearLossProbability() => $_clearField(3);

  @$pb.TagNumber(4)
  $core.int get iterationsRun => $_getIZ(3);
  @$pb.TagNumber(4)
  set iterationsRun($core.int value) => $_setSignedInt32(3, value);
  @$pb.TagNumber(4)
  $core.bool hasIterationsRun() => $_has(3);
  @$pb.TagNumber(4)
  void clearIterationsRun() => $_clearField(4);
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
