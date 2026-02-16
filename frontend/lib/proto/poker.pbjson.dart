// This is a generated file - do not edit.
//
// Generated from poker.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_relative_imports
// ignore_for_file: unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use evaluateHandRequestDescriptor instead')
const EvaluateHandRequest$json = {
  '1': 'EvaluateHandRequest',
  '2': [
    {'1': 'cards', '3': 1, '4': 3, '5': 9, '10': 'cards'},
  ],
};

/// Descriptor for `EvaluateHandRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List evaluateHandRequestDescriptor =
    $convert.base64Decode(
        'ChNFdmFsdWF0ZUhhbmRSZXF1ZXN0EhQKBWNhcmRzGAEgAygJUgVjYXJkcw==');

@$core.Deprecated('Use evaluateHandResponseDescriptor instead')
const EvaluateHandResponse$json = {
  '1': 'EvaluateHandResponse',
  '2': [
    {'1': 'hand_rank', '3': 1, '4': 1, '5': 9, '10': 'handRank'},
    {'1': 'best_five', '3': 2, '4': 3, '5': 9, '10': 'bestFive'},
    {'1': 'rank_value', '3': 3, '4': 1, '5': 5, '10': 'rankValue'},
  ],
};

/// Descriptor for `EvaluateHandResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List evaluateHandResponseDescriptor = $convert.base64Decode(
    'ChRFdmFsdWF0ZUhhbmRSZXNwb25zZRIbCgloYW5kX3JhbmsYASABKAlSCGhhbmRSYW5rEhsKCW'
    'Jlc3RfZml2ZRgCIAMoCVIIYmVzdEZpdmUSHQoKcmFua192YWx1ZRgDIAEoBVIJcmFua1ZhbHVl');

@$core.Deprecated('Use compareHandsRequestDescriptor instead')
const CompareHandsRequest$json = {
  '1': 'CompareHandsRequest',
  '2': [
    {'1': 'player1_cards', '3': 1, '4': 3, '5': 9, '10': 'player1Cards'},
    {'1': 'player2_cards', '3': 2, '4': 3, '5': 9, '10': 'player2Cards'},
  ],
};

/// Descriptor for `CompareHandsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List compareHandsRequestDescriptor = $convert.base64Decode(
    'ChNDb21wYXJlSGFuZHNSZXF1ZXN0EiMKDXBsYXllcjFfY2FyZHMYASADKAlSDHBsYXllcjFDYX'
    'JkcxIjCg1wbGF5ZXIyX2NhcmRzGAIgAygJUgxwbGF5ZXIyQ2FyZHM=');

@$core.Deprecated('Use compareHandsResponseDescriptor instead')
const CompareHandsResponse$json = {
  '1': 'CompareHandsResponse',
  '2': [
    {'1': 'winner', '3': 1, '4': 1, '5': 5, '10': 'winner'},
    {
      '1': 'player1_hand',
      '3': 2,
      '4': 1,
      '5': 11,
      '6': '.poker.EvaluateHandResponse',
      '10': 'player1Hand'
    },
    {
      '1': 'player2_hand',
      '3': 3,
      '4': 1,
      '5': 11,
      '6': '.poker.EvaluateHandResponse',
      '10': 'player2Hand'
    },
    {'1': 'description', '3': 4, '4': 1, '5': 9, '10': 'description'},
  ],
};

/// Descriptor for `CompareHandsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List compareHandsResponseDescriptor = $convert.base64Decode(
    'ChRDb21wYXJlSGFuZHNSZXNwb25zZRIWCgZ3aW5uZXIYASABKAVSBndpbm5lchI+CgxwbGF5ZX'
    'IxX2hhbmQYAiABKAsyGy5wb2tlci5FdmFsdWF0ZUhhbmRSZXNwb25zZVILcGxheWVyMUhhbmQS'
    'PgoMcGxheWVyMl9oYW5kGAMgASgLMhsucG9rZXIuRXZhbHVhdGVIYW5kUmVzcG9uc2VSC3BsYX'
    'llcjJIYW5kEiAKC2Rlc2NyaXB0aW9uGAQgASgJUgtkZXNjcmlwdGlvbg==');

@$core.Deprecated('Use winProbabilityRequestDescriptor instead')
const WinProbabilityRequest$json = {
  '1': 'WinProbabilityRequest',
  '2': [
    {'1': 'hole_cards', '3': 1, '4': 3, '5': 9, '10': 'holeCards'},
    {'1': 'community', '3': 2, '4': 3, '5': 9, '10': 'community'},
    {'1': 'num_opponents', '3': 3, '4': 1, '5': 5, '10': 'numOpponents'},
    {'1': 'iterations', '3': 4, '4': 1, '5': 5, '10': 'iterations'},
  ],
};

/// Descriptor for `WinProbabilityRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List winProbabilityRequestDescriptor = $convert.base64Decode(
    'ChVXaW5Qcm9iYWJpbGl0eVJlcXVlc3QSHQoKaG9sZV9jYXJkcxgBIAMoCVIJaG9sZUNhcmRzEh'
    'wKCWNvbW11bml0eRgCIAMoCVIJY29tbXVuaXR5EiMKDW51bV9vcHBvbmVudHMYAyABKAVSDG51'
    'bU9wcG9uZW50cxIeCgppdGVyYXRpb25zGAQgASgFUgppdGVyYXRpb25z');

@$core.Deprecated('Use winProbabilityResponseDescriptor instead')
const WinProbabilityResponse$json = {
  '1': 'WinProbabilityResponse',
  '2': [
    {'1': 'win_probability', '3': 1, '4': 1, '5': 1, '10': 'winProbability'},
    {'1': 'tie_probability', '3': 2, '4': 1, '5': 1, '10': 'tieProbability'},
    {'1': 'loss_probability', '3': 3, '4': 1, '5': 1, '10': 'lossProbability'},
    {'1': 'iterations_run', '3': 4, '4': 1, '5': 5, '10': 'iterationsRun'},
  ],
};

/// Descriptor for `WinProbabilityResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List winProbabilityResponseDescriptor = $convert.base64Decode(
    'ChZXaW5Qcm9iYWJpbGl0eVJlc3BvbnNlEicKD3dpbl9wcm9iYWJpbGl0eRgBIAEoAVIOd2luUH'
    'JvYmFiaWxpdHkSJwoPdGllX3Byb2JhYmlsaXR5GAIgASgBUg50aWVQcm9iYWJpbGl0eRIpChBs'
    'b3NzX3Byb2JhYmlsaXR5GAMgASgBUg9sb3NzUHJvYmFiaWxpdHkSJQoOaXRlcmF0aW9uc19ydW'
    '4YBCABKAVSDWl0ZXJhdGlvbnNSdW4=');
