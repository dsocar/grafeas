// Copyright 2019 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: proto/v1/common.proto

package grafeas_go_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Kind represents the kinds of notes supported.
type NoteKind int32

const (
	// Default value. This value is unused.
	NoteKind_NOTE_KIND_UNSPECIFIED NoteKind = 0
	// The note and occurrence represent a package vulnerability.
	NoteKind_VULNERABILITY NoteKind = 1
	// The note and occurrence assert build provenance.
	NoteKind_BUILD NoteKind = 2
	// This represents an image basis relationship.
	NoteKind_IMAGE NoteKind = 3
	// This represents a package installed via a package manager.
	NoteKind_PACKAGE NoteKind = 4
	// The note and occurrence track deployment events.
	NoteKind_DEPLOYMENT NoteKind = 5
	// The note and occurrence track the initial discovery status of a resource.
	NoteKind_DISCOVERY NoteKind = 6
	// This represents a logical "role" that can attest to artifacts.
	NoteKind_ATTESTATION NoteKind = 7
	// This represents an available package upgrade.
	NoteKind_UPGRADE NoteKind = 8
	// This represents a Compliance Note
	NoteKind_COMPLIANCE NoteKind = 9
	// This represents a DSSE attestation Note
	NoteKind_DSSE_ATTESTATION NoteKind = 10
)

// Enum value maps for NoteKind.
var (
	NoteKind_name = map[int32]string{
		0:  "NOTE_KIND_UNSPECIFIED",
		1:  "VULNERABILITY",
		2:  "BUILD",
		3:  "IMAGE",
		4:  "PACKAGE",
		5:  "DEPLOYMENT",
		6:  "DISCOVERY",
		7:  "ATTESTATION",
		8:  "UPGRADE",
		9:  "COMPLIANCE",
		10: "DSSE_ATTESTATION",
	}
	NoteKind_value = map[string]int32{
		"NOTE_KIND_UNSPECIFIED": 0,
		"VULNERABILITY":         1,
		"BUILD":                 2,
		"IMAGE":                 3,
		"PACKAGE":               4,
		"DEPLOYMENT":            5,
		"DISCOVERY":             6,
		"ATTESTATION":           7,
		"UPGRADE":               8,
		"COMPLIANCE":            9,
		"DSSE_ATTESTATION":      10,
	}
)

func (x NoteKind) Enum() *NoteKind {
	p := new(NoteKind)
	*p = x
	return p
}

func (x NoteKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NoteKind) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1_common_proto_enumTypes[0].Descriptor()
}

func (NoteKind) Type() protoreflect.EnumType {
	return &file_proto_v1_common_proto_enumTypes[0]
}

func (x NoteKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NoteKind.Descriptor instead.
func (NoteKind) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1_common_proto_rawDescGZIP(), []int{0}
}

// Metadata for any related URL information.
type RelatedUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specific URL associated with the resource.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Label to describe usage of the URL.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *RelatedUrl) Reset() {
	*x = RelatedUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelatedUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelatedUrl) ProtoMessage() {}

func (x *RelatedUrl) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelatedUrl.ProtoReflect.Descriptor instead.
func (*RelatedUrl) Descriptor() ([]byte, []int) {
	return file_proto_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *RelatedUrl) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RelatedUrl) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// Verifiers (e.g. Kritis implementations) MUST verify signatures
// with respect to the trust anchors defined in policy (e.g. a Kritis policy).
// Typically this means that the verifier has been configured with a map from
// `public_key_id` to public key material (and any required parameters, e.g.
// signing algorithm).
//
// In particular, verification implementations MUST NOT treat the signature
// `public_key_id` as anything more than a key lookup hint. The `public_key_id`
// DOES NOT validate or authenticate a public key; it only provides a mechanism
// for quickly selecting a public key ALREADY CONFIGURED on the verifier through
// a trusted channel. Verification implementations MUST reject signatures in any
// of the following circumstances:
//   * The `public_key_id` is not recognized by the verifier.
//   * The public key that `public_key_id` refers to does not verify the
//     signature with respect to the payload.
//
// The `signature` contents SHOULD NOT be "attached" (where the payload is
// included with the serialized `signature` bytes). Verifiers MUST ignore any
// "attached" payload and only verify signatures with respect to explicitly
// provided payload (e.g. a `payload` field on the proto message that holds
// this Signature, or the canonical serialization of the proto message that
// holds this signature).
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The content of the signature, an opaque bytestring.
	// The payload that this signature verifies MUST be unambiguously provided
	// with the Signature during verification. A wrapper message might provide
	// the payload explicitly. Alternatively, a message might have a canonical
	// serialization that can always be unambiguously computed to derive the
	// payload.
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// The identifier for the public key that verifies this signature.
	//   * The `public_key_id` is required.
	//   * The `public_key_id` SHOULD be an RFC3986 conformant URI.
	//   * When possible, the `public_key_id` SHOULD be an immutable reference,
	//     such as a cryptographic digest.
	//
	// Examples of valid `public_key_id`s:
	//
	// OpenPGP V4 public key fingerprint:
	//   * "openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA"
	// See https://www.iana.org/assignments/uri-schemes/prov/openpgp4fpr for more
	// details on this scheme.
	//
	// RFC6920 digest-named SubjectPublicKeyInfo (digest of the DER
	// serialization):
	//   * "ni:///sha-256;cD9o9Cq6LG3jD0iKXqEi_vdjJGecm_iXkbqVoScViaU"
	//   * "nih:///sha-256;703f68f42aba2c6de30f488a5ea122fef76324679c9bf89791ba95a1271589a5"
	PublicKeyId string `protobuf:"bytes,2,opt,name=public_key_id,json=publicKeyId,proto3" json:"public_key_id,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_proto_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *Signature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Signature) GetPublicKeyId() string {
	if x != nil {
		return x.PublicKeyId
	}
	return ""
}

// MUST match https://github.com/secure-systems-lab/dsse/blob/master/envelope.proto.
// An authenticated message of arbitrary type.
type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload     []byte               `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	PayloadType string               `protobuf:"bytes,2,opt,name=payload_type,json=payloadType,proto3" json:"payload_type,omitempty"`
	Signatures  []*EnvelopeSignature `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_proto_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *Envelope) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Envelope) GetPayloadType() string {
	if x != nil {
		return x.PayloadType
	}
	return ""
}

func (x *Envelope) GetSignatures() []*EnvelopeSignature {
	if x != nil {
		return x.Signatures
	}
	return nil
}

type EnvelopeSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sig   []byte `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
	Keyid string `protobuf:"bytes,2,opt,name=keyid,proto3" json:"keyid,omitempty"`
}

func (x *EnvelopeSignature) Reset() {
	*x = EnvelopeSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvelopeSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvelopeSignature) ProtoMessage() {}

func (x *EnvelopeSignature) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvelopeSignature.ProtoReflect.Descriptor instead.
func (*EnvelopeSignature) Descriptor() ([]byte, []int) {
	return file_proto_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *EnvelopeSignature) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *EnvelopeSignature) GetKeyid() string {
	if x != nil {
		return x.Keyid
	}
	return ""
}

var File_proto_v1_common_proto protoreflect.FileDescriptor

var file_proto_v1_common_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x22, 0x34, 0x0a, 0x0a, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x55, 0x72,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x4d, 0x0a, 0x09, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x22, 0x3b, 0x0a, 0x11, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x65, 0x79, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x69, 0x64, 0x2a, 0xbe,
	0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x15, 0x4e,
	0x4f, 0x54, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x55, 0x4c, 0x4e, 0x45, 0x52,
	0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x55, 0x49,
	0x4c, 0x44, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x03, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a,
	0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09,
	0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x41,
	0x54, 0x54, 0x45, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x53, 0x53,
	0x45, 0x5f, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x42,
	0x4d, 0x0a, 0x0d, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x5f,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_common_proto_rawDescOnce sync.Once
	file_proto_v1_common_proto_rawDescData = file_proto_v1_common_proto_rawDesc
)

func file_proto_v1_common_proto_rawDescGZIP() []byte {
	file_proto_v1_common_proto_rawDescOnce.Do(func() {
		file_proto_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_common_proto_rawDescData)
	})
	return file_proto_v1_common_proto_rawDescData
}

var file_proto_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_v1_common_proto_goTypes = []interface{}{
	(NoteKind)(0),             // 0: grafeas.v1.NoteKind
	(*RelatedUrl)(nil),        // 1: grafeas.v1.RelatedUrl
	(*Signature)(nil),         // 2: grafeas.v1.Signature
	(*Envelope)(nil),          // 3: grafeas.v1.Envelope
	(*EnvelopeSignature)(nil), // 4: grafeas.v1.EnvelopeSignature
}
var file_proto_v1_common_proto_depIdxs = []int32{
	4, // 0: grafeas.v1.Envelope.signatures:type_name -> grafeas.v1.EnvelopeSignature
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_v1_common_proto_init() }
func file_proto_v1_common_proto_init() {
	if File_proto_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelatedUrl); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvelopeSignature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_v1_common_proto_goTypes,
		DependencyIndexes: file_proto_v1_common_proto_depIdxs,
		EnumInfos:         file_proto_v1_common_proto_enumTypes,
		MessageInfos:      file_proto_v1_common_proto_msgTypes,
	}.Build()
	File_proto_v1_common_proto = out.File
	file_proto_v1_common_proto_rawDesc = nil
	file_proto_v1_common_proto_goTypes = nil
	file_proto_v1_common_proto_depIdxs = nil
}
