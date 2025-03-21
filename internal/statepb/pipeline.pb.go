// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.6
// source: pipeline.proto

package statepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The degree of automation to use when generating/releasing.
type AutomationLevel int32

const (
	// Not used.
	AutomationLevel_AUTOMATION_LEVEL_NONE AutomationLevel = 0
	// Automation is blocked: this API/library should be skipped
	AutomationLevel_AUTOMATION_LEVEL_BLOCKED AutomationLevel = 1
	// Automation can generate changes/releases,
	// but they need to be reviewed.
	AutomationLevel_AUTOMATION_LEVEL_MANUAL_REVIEW AutomationLevel = 2
	// Automation can generated changes/releases which can
	// proceed without further review if all tests pass.
	AutomationLevel_AUTOMATION_LEVEL_AUTOMATIC AutomationLevel = 3
)

// Enum value maps for AutomationLevel.
var (
	AutomationLevel_name = map[int32]string{
		0: "AUTOMATION_LEVEL_NONE",
		1: "AUTOMATION_LEVEL_BLOCKED",
		2: "AUTOMATION_LEVEL_MANUAL_REVIEW",
		3: "AUTOMATION_LEVEL_AUTOMATIC",
	}
	AutomationLevel_value = map[string]int32{
		"AUTOMATION_LEVEL_NONE":          0,
		"AUTOMATION_LEVEL_BLOCKED":       1,
		"AUTOMATION_LEVEL_MANUAL_REVIEW": 2,
		"AUTOMATION_LEVEL_AUTOMATIC":     3,
	}
)

func (x AutomationLevel) Enum() *AutomationLevel {
	p := new(AutomationLevel)
	*p = x
	return p
}

func (x AutomationLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AutomationLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_pipeline_proto_enumTypes[0].Descriptor()
}

func (AutomationLevel) Type() protoreflect.EnumType {
	return &file_pipeline_proto_enumTypes[0]
}

func (x AutomationLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AutomationLevel.Descriptor instead.
func (AutomationLevel) EnumDescriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{0}
}

// Overall state of the generation and release pipeline. This is expected
// to be stored in each language repo as generator-input/pipeline-state.json.
type PipelineState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The image tag that the CLI should use when invoking the
	// language-specific tooling. The image name is assumed by convention.
	ImageTag string `protobuf:"bytes,1,opt,name=image_tag,json=imageTag,proto3" json:"image_tag,omitempty"`
	// The state of each API which has code generated within this repository.
	ApiGenerationStates []*ApiGenerationState `protobuf:"bytes,2,rep,name=api_generation_states,json=apiGenerationStates,proto3" json:"api_generation_states,omitempty"`
	// The state of each library which is released within this repository.
	LibraryReleaseStates []*LibraryReleaseState `protobuf:"bytes,3,rep,name=library_release_states,json=libraryReleaseStates,proto3" json:"library_release_states,omitempty"`
	// Paths to files/directories which can trigger
	// a release in all libraries.
	CommonLibrarySourcePaths []string `protobuf:"bytes,4,rep,name=common_library_source_paths,json=commonLibrarySourcePaths,proto3" json:"common_library_source_paths,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *PipelineState) Reset() {
	*x = PipelineState{}
	mi := &file_pipeline_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PipelineState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineState) ProtoMessage() {}

func (x *PipelineState) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineState.ProtoReflect.Descriptor instead.
func (*PipelineState) Descriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineState) GetImageTag() string {
	if x != nil {
		return x.ImageTag
	}
	return ""
}

func (x *PipelineState) GetApiGenerationStates() []*ApiGenerationState {
	if x != nil {
		return x.ApiGenerationStates
	}
	return nil
}

func (x *PipelineState) GetLibraryReleaseStates() []*LibraryReleaseState {
	if x != nil {
		return x.LibraryReleaseStates
	}
	return nil
}

func (x *PipelineState) GetCommonLibrarySourcePaths() []string {
	if x != nil {
		return x.CommonLibrarySourcePaths
	}
	return nil
}

// Generation state of a single API.
type ApiGenerationState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The API identifier (currently expected to be a path)
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The commit hash (in googleapis) from which the API was
	// last generated.
	LastGeneratedCommit string `protobuf:"bytes,2,opt,name=last_generated_commit,json=lastGeneratedCommit,proto3" json:"last_generated_commit,omitempty"`
	// The automation level for regeneration of this API.
	AutomationLevel AutomationLevel `protobuf:"varint,3,opt,name=automation_level,json=automationLevel,proto3,enum=google.cloud.sdk.pipeline.AutomationLevel" json:"automation_level,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ApiGenerationState) Reset() {
	*x = ApiGenerationState{}
	mi := &file_pipeline_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiGenerationState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiGenerationState) ProtoMessage() {}

func (x *ApiGenerationState) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiGenerationState.ProtoReflect.Descriptor instead.
func (*ApiGenerationState) Descriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *ApiGenerationState) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApiGenerationState) GetLastGeneratedCommit() string {
	if x != nil {
		return x.LastGeneratedCommit
	}
	return ""
}

func (x *ApiGenerationState) GetAutomationLevel() AutomationLevel {
	if x != nil {
		return x.AutomationLevel
	}
	return AutomationLevel_AUTOMATION_LEVEL_NONE
}

// Generation state of a single library.
type LibraryReleaseState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The library identifier (language-specific format)
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The last version that was released, if any.
	CurrentVersion string `protobuf:"bytes,2,opt,name=current_version,json=currentVersion,proto3" json:"current_version,omitempty"`
	// The next version to release (to force a specific version number).
	// This should usually be unset.
	NextVersion string `protobuf:"bytes,3,opt,name=next_version,json=nextVersion,proto3" json:"next_version,omitempty"`
	// The automation level for releases for this library.
	AutomationLevel AutomationLevel `protobuf:"varint,4,opt,name=automation_level,json=automationLevel,proto3,enum=google.cloud.sdk.pipeline.AutomationLevel" json:"automation_level,omitempty"`
	// The timestamp of the latest release. (This is typically
	// some timestamp within the process of generating the release
	// PR for the library. Importantly, it's not just a date as
	// there may be reasons to release a library multiple times
	// within a day.)
	ReleaseTimestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=release_timestamp,json=releaseTimestamp,proto3" json:"release_timestamp,omitempty"`
	// The APIs included in this library.
	Apis []*LibraryApiMapping `protobuf:"bytes,6,rep,name=apis,proto3" json:"apis,omitempty"`
	// Paths to files or directories contributing to this library.
	SourcePaths   []string `protobuf:"bytes,7,rep,name=source_paths,json=sourcePaths,proto3" json:"source_paths,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LibraryReleaseState) Reset() {
	*x = LibraryReleaseState{}
	mi := &file_pipeline_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LibraryReleaseState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LibraryReleaseState) ProtoMessage() {}

func (x *LibraryReleaseState) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LibraryReleaseState.ProtoReflect.Descriptor instead.
func (*LibraryReleaseState) Descriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{2}
}

func (x *LibraryReleaseState) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LibraryReleaseState) GetCurrentVersion() string {
	if x != nil {
		return x.CurrentVersion
	}
	return ""
}

func (x *LibraryReleaseState) GetNextVersion() string {
	if x != nil {
		return x.NextVersion
	}
	return ""
}

func (x *LibraryReleaseState) GetAutomationLevel() AutomationLevel {
	if x != nil {
		return x.AutomationLevel
	}
	return AutomationLevel_AUTOMATION_LEVEL_NONE
}

func (x *LibraryReleaseState) GetReleaseTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ReleaseTimestamp
	}
	return nil
}

func (x *LibraryReleaseState) GetApis() []*LibraryApiMapping {
	if x != nil {
		return x.Apis
	}
	return nil
}

func (x *LibraryReleaseState) GetSourcePaths() []string {
	if x != nil {
		return x.SourcePaths
	}
	return nil
}

// Information about the state of an API which is included in a library.
type LibraryApiMapping struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The API identifier as per ApiGenerationState.
	ApiId string `protobuf:"bytes,1,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"`
	// The last-generated commit hash (within the API definition repo)
	// at the point of the last/in-progress release. (This is taken
	// from the generation state at the time of the release.)
	LastReleasedCommit string `protobuf:"bytes,2,opt,name=last_released_commit,json=lastReleasedCommit,proto3" json:"last_released_commit,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *LibraryApiMapping) Reset() {
	*x = LibraryApiMapping{}
	mi := &file_pipeline_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LibraryApiMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LibraryApiMapping) ProtoMessage() {}

func (x *LibraryApiMapping) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LibraryApiMapping.ProtoReflect.Descriptor instead.
func (*LibraryApiMapping) Descriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{3}
}

func (x *LibraryApiMapping) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

func (x *LibraryApiMapping) GetLastReleasedCommit() string {
	if x != nil {
		return x.LastReleasedCommit
	}
	return ""
}

var File_pipeline_proto protoreflect.FileDescriptor

var file_pipeline_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a,
	0x0d, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x61, 0x67, 0x12, 0x61, 0x0a, 0x15, 0x61,
	0x70, 0x69, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x13, 0x61, 0x70, 0x69, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x64,
	0x0a, 0x16, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x14,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x12, 0x41, 0x70, 0x69, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x55,
	0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0xf6, 0x02, 0x0a, 0x13, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65,
	0x78, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x10, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x0f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x47, 0x0a, 0x11, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x40, 0x0a, 0x04, 0x61, 0x70, 0x69,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x70, 0x69, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x61, 0x70, 0x69, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x22, 0x5c,
	0x0a, 0x11, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x70, 0x69, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2a, 0x8e, 0x01, 0x0a,
	0x0f, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x19, 0x0a, 0x15, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x41,
	0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x41, 0x55, 0x54,
	0x4f, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x41,
	0x4e, 0x55, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x02, 0x12, 0x1e, 0x0a,
	0x1a, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x10, 0x03, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x70,
	0x62, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_pipeline_proto_rawDescOnce sync.Once
	file_pipeline_proto_rawDescData []byte
)

func file_pipeline_proto_rawDescGZIP() []byte {
	file_pipeline_proto_rawDescOnce.Do(func() {
		file_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pipeline_proto_rawDesc), len(file_pipeline_proto_rawDesc)))
	})
	return file_pipeline_proto_rawDescData
}

var file_pipeline_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pipeline_proto_goTypes = []any{
	(AutomationLevel)(0),          // 0: google.cloud.sdk.pipeline.AutomationLevel
	(*PipelineState)(nil),         // 1: google.cloud.sdk.pipeline.PipelineState
	(*ApiGenerationState)(nil),    // 2: google.cloud.sdk.pipeline.ApiGenerationState
	(*LibraryReleaseState)(nil),   // 3: google.cloud.sdk.pipeline.LibraryReleaseState
	(*LibraryApiMapping)(nil),     // 4: google.cloud.sdk.pipeline.LibraryApiMapping
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_pipeline_proto_depIdxs = []int32{
	2, // 0: google.cloud.sdk.pipeline.PipelineState.api_generation_states:type_name -> google.cloud.sdk.pipeline.ApiGenerationState
	3, // 1: google.cloud.sdk.pipeline.PipelineState.library_release_states:type_name -> google.cloud.sdk.pipeline.LibraryReleaseState
	0, // 2: google.cloud.sdk.pipeline.ApiGenerationState.automation_level:type_name -> google.cloud.sdk.pipeline.AutomationLevel
	0, // 3: google.cloud.sdk.pipeline.LibraryReleaseState.automation_level:type_name -> google.cloud.sdk.pipeline.AutomationLevel
	5, // 4: google.cloud.sdk.pipeline.LibraryReleaseState.release_timestamp:type_name -> google.protobuf.Timestamp
	4, // 5: google.cloud.sdk.pipeline.LibraryReleaseState.apis:type_name -> google.cloud.sdk.pipeline.LibraryApiMapping
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pipeline_proto_init() }
func file_pipeline_proto_init() {
	if File_pipeline_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pipeline_proto_rawDesc), len(file_pipeline_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pipeline_proto_goTypes,
		DependencyIndexes: file_pipeline_proto_depIdxs,
		EnumInfos:         file_pipeline_proto_enumTypes,
		MessageInfos:      file_pipeline_proto_msgTypes,
	}.Build()
	File_pipeline_proto = out.File
	file_pipeline_proto_goTypes = nil
	file_pipeline_proto_depIdxs = nil
}
