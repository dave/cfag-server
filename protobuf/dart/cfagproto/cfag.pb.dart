///
//  Generated code. Do not modify.
//  source: cfag.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

// ignore: UNUSED_SHOWN_NAME
import 'dart:core' show int, bool, double, String, List, Map, override;

import 'package:protobuf/protobuf.dart' as $pb;

class Person extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('Person', package: const $pb.PackageName('cfagproto'))
    ..aOS(1, 'firstName')
    ..aOS(2, 'lastName')
    ..aOS(3, 'email')
    ..hasRequiredFields = false
  ;

  Person() : super();
  Person.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Person.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Person clone() => new Person()..mergeFromMessage(this);
  Person copyWith(void Function(Person) updates) => super.copyWith((message) => updates(message as Person));
  $pb.BuilderInfo get info_ => _i;
  static Person create() => new Person();
  static $pb.PbList<Person> createRepeated() => new $pb.PbList<Person>();
  static Person getDefault() => _defaultInstance ??= create()..freeze();
  static Person _defaultInstance;
  static void $checkItem(Person v) {
    if (v is! Person) $pb.checkItemFailed(v, _i.qualifiedMessageName);
  }

  String get firstName => $_getS(0, '');
  set firstName(String v) { $_setString(0, v); }
  bool hasFirstName() => $_has(0);
  void clearFirstName() => clearField(1);

  String get lastName => $_getS(1, '');
  set lastName(String v) { $_setString(1, v); }
  bool hasLastName() => $_has(1);
  void clearLastName() => clearField(2);

  String get email => $_getS(2, '');
  set email(String v) { $_setString(2, v); }
  bool hasEmail() => $_has(2);
  void clearEmail() => clearField(3);
}

