/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/divinerapier/learn-kubernetes/vendor/k8s.io/api/events/v1beta1/generated.proto

package v1beta1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	v11 "k8s.io/api/core/v1"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *Event) Reset()      { *m = Event{} }
func (*Event) ProtoMessage() {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f97f691c32a5ac8, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *EventList) Reset()      { *m = EventList{} }
func (*EventList) ProtoMessage() {}
func (*EventList) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f97f691c32a5ac8, []int{1}
}
func (m *EventList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EventList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventList.Merge(m, src)
}
func (m *EventList) XXX_Size() int {
	return m.Size()
}
func (m *EventList) XXX_DiscardUnknown() {
	xxx_messageInfo_EventList.DiscardUnknown(m)
}

var xxx_messageInfo_EventList proto.InternalMessageInfo

func (m *EventSeries) Reset()      { *m = EventSeries{} }
func (*EventSeries) ProtoMessage() {}
func (*EventSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f97f691c32a5ac8, []int{2}
}
func (m *EventSeries) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EventSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSeries.Merge(m, src)
}
func (m *EventSeries) XXX_Size() int {
	return m.Size()
}
func (m *EventSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSeries.DiscardUnknown(m)
}

var xxx_messageInfo_EventSeries proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Event)(nil), "k8s.io.api.events.v1beta1.Event")
	proto.RegisterType((*EventList)(nil), "k8s.io.api.events.v1beta1.EventList")
	proto.RegisterType((*EventSeries)(nil), "k8s.io.api.events.v1beta1.EventSeries")
}

func init() {
	proto.RegisterFile("github.com/divinerapier/learn-kubernetes/vendor/k8s.io/api/events/v1beta1/generated.proto", fileDescriptor_4f97f691c32a5ac8)
}

var fileDescriptor_4f97f691c32a5ac8 = []byte{
	// 801 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x16, 0x13, 0x4b, 0xb2, 0x56, 0x49, 0x2c, 0x6f, 0x0e, 0xde, 0xb8, 0x00, 0xa5, 0x2a, 0x40,
	0x20, 0x14, 0x08, 0x59, 0x07, 0x45, 0xdb, 0x6b, 0x18, 0xb9, 0x45, 0x02, 0xbb, 0x01, 0xd6, 0x3e,
	0x15, 0x3d, 0x64, 0x45, 0x4d, 0x68, 0x56, 0xe2, 0x2e, 0xb1, 0xbb, 0x12, 0xe0, 0x5b, 0x2f, 0x05,
	0x7a, 0xec, 0x33, 0xf4, 0x09, 0xfa, 0x18, 0x3e, 0xe6, 0x98, 0x93, 0x50, 0xb3, 0x6f, 0xd1, 0x53,
	0xc1, 0xe5, 0x4a, 0x94, 0xf5, 0x83, 0xa8, 0xe8, 0x4d, 0x9c, 0xf9, 0x7e, 0x66, 0x66, 0x47, 0x83,
	0x82, 0xd1, 0xb7, 0xca, 0x8b, 0x85, 0x3f, 0x9a, 0x0c, 0x40, 0x72, 0xd0, 0xa0, 0xfc, 0x29, 0xf0,
	0xa1, 0x90, 0xbe, 0x4d, 0xb0, 0x34, 0xf6, 0x61, 0x0a, 0x5c, 0x2b, 0x7f, 0x7a, 0x32, 0x00, 0xcd,
	0x4e, 0xfc, 0x08, 0x38, 0x48, 0xa6, 0x61, 0xe8, 0xa5, 0x52, 0x68, 0x81, 0x9f, 0x14, 0x50, 0x8f,
	0xa5, 0xb1, 0x57, 0x40, 0x3d, 0x0b, 0x3d, 0x7e, 0x1e, 0xc5, 0xfa, 0x6a, 0x32, 0xf0, 0x42, 0x91,
	0xf8, 0x91, 0x88, 0x84, 0x6f, 0x18, 0x83, 0xc9, 0x7b, 0xf3, 0x65, 0x3e, 0xcc, 0xaf, 0x42, 0xe9,
	0xb8, 0xbb, 0x64, 0x1a, 0x0a, 0x09, 0xfe, 0x74, 0xcd, 0xed, 0xf8, 0xab, 0x12, 0x93, 0xb0, 0xf0,
	0x2a, 0xe6, 0x20, 0xaf, 0xfd, 0x74, 0x14, 0xe5, 0x01, 0xe5, 0x27, 0xa0, 0xd9, 0x26, 0x96, 0xbf,
	0x8d, 0x25, 0x27, 0x5c, 0xc7, 0x09, 0xac, 0x11, 0xbe, 0xfe, 0x14, 0x41, 0x85, 0x57, 0x90, 0xb0,
	0x55, 0x5e, 0xf7, 0x8f, 0x06, 0xaa, 0x9e, 0xe6, 0x43, 0xc0, 0xef, 0xd0, 0x7e, 0x5e, 0xcd, 0x90,
	0x69, 0x46, 0x9c, 0x8e, 0xd3, 0x6b, 0xbe, 0xf8, 0xd2, 0x2b, 0x27, 0xb5, 0x10, 0xf5, 0xd2, 0x51,
	0x94, 0x07, 0x94, 0x97, 0xa3, 0xbd, 0xe9, 0x89, 0xf7, 0x76, 0xf0, 0x33, 0x84, 0xfa, 0x1c, 0x34,
	0x0b, 0xf0, 0xcd, 0xac, 0x5d, 0xc9, 0x66, 0x6d, 0x54, 0xc6, 0xe8, 0x42, 0x15, 0xbf, 0x43, 0x0d,
	0x33, 0xef, 0xcb, 0x38, 0x01, 0x72, 0xcf, 0x58, 0xf8, 0xbb, 0x59, 0x9c, 0xc7, 0xa1, 0x14, 0x39,
	0x2d, 0x38, 0xb4, 0x0e, 0x8d, 0xd3, 0xb9, 0x12, 0x2d, 0x45, 0xf1, 0x1b, 0x54, 0x53, 0x20, 0x63,
	0x50, 0xe4, 0xbe, 0x91, 0x7f, 0xe6, 0x6d, 0x7d, 0x6b, 0xcf, 0x08, 0x5c, 0x18, 0x74, 0x80, 0xb2,
	0x59, 0xbb, 0x56, 0xfc, 0xa6, 0x56, 0x01, 0x9f, 0xa3, 0xc7, 0x12, 0x52, 0x21, 0x75, 0xcc, 0xa3,
	0x57, 0x82, 0x6b, 0x29, 0xc6, 0x63, 0x90, 0x64, 0xaf, 0xe3, 0xf4, 0x1a, 0xc1, 0x67, 0xb6, 0x8c,
	0xc7, 0x74, 0x1d, 0x42, 0x37, 0xf1, 0xf0, 0xf7, 0xe8, 0x70, 0x11, 0x7e, 0xcd, 0x95, 0x66, 0x3c,
	0x04, 0x52, 0x35, 0x62, 0x4f, 0xac, 0xd8, 0x21, 0x5d, 0x05, 0xd0, 0x75, 0x0e, 0x7e, 0x86, 0x6a,
	0x2c, 0xd4, 0xb1, 0xe0, 0xa4, 0x66, 0xd8, 0x8f, 0x2c, 0xbb, 0xf6, 0xd2, 0x44, 0xa9, 0xcd, 0xe6,
	0x38, 0x09, 0x4c, 0x09, 0x4e, 0xea, 0x77, 0x71, 0xd4, 0x44, 0xa9, 0xcd, 0xe2, 0x4b, 0xd4, 0x90,
	0x10, 0x31, 0x39, 0x8c, 0x79, 0x44, 0xf6, 0xcd, 0xd8, 0x9e, 0x2e, 0x8f, 0x2d, 0x5f, 0xec, 0xf2,
	0x99, 0x29, 0xbc, 0x07, 0x09, 0x3c, 0x5c, 0x7a, 0x09, 0x3a, 0x67, 0xd3, 0x52, 0x08, 0xbf, 0x41,
	0x75, 0x09, 0xe3, 0x7c, 0xd1, 0x48, 0x63, 0x77, 0xcd, 0x66, 0x36, 0x6b, 0xd7, 0x69, 0xc1, 0xa3,
	0x73, 0x01, 0xdc, 0x41, 0x7b, 0x5c, 0x68, 0x20, 0xc8, 0xf4, 0xf1, 0xc0, 0xfa, 0xee, 0xfd, 0x20,
	0x34, 0x50, 0x93, 0xc9, 0x11, 0xfa, 0x3a, 0x05, 0xd2, 0xbc, 0x8b, 0xb8, 0xbc, 0x4e, 0x81, 0x9a,
	0x0c, 0x06, 0xd4, 0x1a, 0x42, 0x2a, 0x21, 0xcc, 0x15, 0x2f, 0xc4, 0x44, 0x86, 0x40, 0x1e, 0x98,
	0xc2, 0xda, 0x9b, 0x0a, 0x2b, 0x96, 0xc3, 0xc0, 0x02, 0x62, 0xe5, 0x5a, 0xfd, 0x15, 0x01, 0xba,
	0x26, 0x89, 0x7f, 0x73, 0x10, 0x29, 0x83, 0xdf, 0xc5, 0x52, 0x99, 0xc5, 0x54, 0x9a, 0x25, 0x29,
	0x79, 0x68, 0xfc, 0xbe, 0xd8, 0x6d, 0xe5, 0xcd, 0xb6, 0x77, 0xac, 0x35, 0xe9, 0x6f, 0xd1, 0xa4,
	0x5b, 0xdd, 0xf0, 0xaf, 0x0e, 0x3a, 0x2a, 0x93, 0x67, 0x6c, 0xb9, 0x92, 0x47, 0xff, 0xb9, 0x92,
	0xb6, 0xad, 0xe4, 0xa8, 0xbf, 0x59, 0x92, 0x6e, 0xf3, 0xc2, 0x2f, 0xd1, 0x41, 0x99, 0x7a, 0x25,
	0x26, 0x5c, 0x93, 0x83, 0x8e, 0xd3, 0xab, 0x06, 0x47, 0x56, 0xf2, 0xa0, 0x7f, 0x37, 0x4d, 0x57,
	0xf1, 0xdd, 0x3f, 0x1d, 0x54, 0xfc, 0xdf, 0xcf, 0x62, 0xa5, 0xf1, 0x4f, 0x6b, 0x87, 0xca, 0xdb,
	0xad, 0x91, 0x9c, 0x6d, 0xce, 0x54, 0xcb, 0x3a, 0xef, 0xcf, 0x23, 0x4b, 0x47, 0xea, 0x14, 0x55,
	0x63, 0x0d, 0x89, 0x22, 0xf7, 0x3a, 0xf7, 0x7b, 0xcd, 0x17, 0x9d, 0x4f, 0x5d, 0x90, 0xe0, 0xa1,
	0x15, 0xab, 0xbe, 0xce, 0x69, 0xb4, 0x60, 0x77, 0x33, 0x07, 0x35, 0x97, 0x2e, 0x0c, 0x7e, 0x8a,
	0xaa, 0xa1, 0xe9, 0xdd, 0x31, 0xbd, 0x2f, 0x48, 0x45, 0xc7, 0x45, 0x0e, 0x4f, 0x50, 0x6b, 0xcc,
	0x94, 0x7e, 0x3b, 0x50, 0x20, 0xa7, 0x30, 0xfc, 0x3f, 0x77, 0x72, 0xb1, 0xb4, 0x67, 0x2b, 0x82,
	0x74, 0xcd, 0x02, 0x7f, 0x83, 0xaa, 0x4a, 0x33, 0x0d, 0xe6, 0x68, 0x36, 0x82, 0xcf, 0xe7, 0xb5,
	0x5d, 0xe4, 0xc1, 0x7f, 0x66, 0xed, 0xd6, 0x52, 0x23, 0x26, 0x46, 0x0b, 0x7c, 0xf0, 0xfc, 0xe6,
	0xd6, 0xad, 0x7c, 0xb8, 0x75, 0x2b, 0x1f, 0x6f, 0xdd, 0xca, 0x2f, 0x99, 0xeb, 0xdc, 0x64, 0xae,
	0xf3, 0x21, 0x73, 0x9d, 0x8f, 0x99, 0xeb, 0xfc, 0x95, 0xb9, 0xce, 0xef, 0x7f, 0xbb, 0x95, 0x1f,
	0xeb, 0x76, 0x5e, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x25, 0x9b, 0x14, 0x4d, 0xbd, 0x07, 0x00,
	0x00,
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintGenerated(dAtA, i, uint64(m.DeprecatedCount))
	i--
	dAtA[i] = 0x78
	{
		size, err := m.DeprecatedLastTimestamp.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size, err := m.DeprecatedFirstTimestamp.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size, err := m.DeprecatedSource.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	i -= len(m.Type)
	copy(dAtA[i:], m.Type)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Type)))
	i--
	dAtA[i] = 0x5a
	i -= len(m.Note)
	copy(dAtA[i:], m.Note)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Note)))
	i--
	dAtA[i] = 0x52
	if m.Related != nil {
		{
			size, err := m.Related.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	{
		size, err := m.Regarding.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	i -= len(m.Reason)
	copy(dAtA[i:], m.Reason)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Reason)))
	i--
	dAtA[i] = 0x3a
	i -= len(m.Action)
	copy(dAtA[i:], m.Action)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Action)))
	i--
	dAtA[i] = 0x32
	i -= len(m.ReportingInstance)
	copy(dAtA[i:], m.ReportingInstance)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ReportingInstance)))
	i--
	dAtA[i] = 0x2a
	i -= len(m.ReportingController)
	copy(dAtA[i:], m.ReportingController)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ReportingController)))
	i--
	dAtA[i] = 0x22
	if m.Series != nil {
		{
			size, err := m.Series.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.EventTime.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EventList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EventSeries) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSeries) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSeries) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.State)
	copy(dAtA[i:], m.State)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.State)))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.LastObservedTime.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	i = encodeVarintGenerated(dAtA, i, uint64(m.Count))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.EventTime.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.Series != nil {
		l = m.Series.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.ReportingController)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.ReportingInstance)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Action)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Reason)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Regarding.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.Related != nil {
		l = m.Related.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Note)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Type)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.DeprecatedSource.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.DeprecatedFirstTimestamp.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.DeprecatedLastTimestamp.Size()
	n += 1 + l + sovGenerated(uint64(l))
	n += 1 + sovGenerated(uint64(m.DeprecatedCount))
	return n
}

func (m *EventList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *EventSeries) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Count))
	l = m.LastObservedTime.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.State)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Event) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Event{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`EventTime:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.EventTime), "MicroTime", "v1.MicroTime", 1), `&`, ``, 1) + `,`,
		`Series:` + strings.Replace(this.Series.String(), "EventSeries", "EventSeries", 1) + `,`,
		`ReportingController:` + fmt.Sprintf("%v", this.ReportingController) + `,`,
		`ReportingInstance:` + fmt.Sprintf("%v", this.ReportingInstance) + `,`,
		`Action:` + fmt.Sprintf("%v", this.Action) + `,`,
		`Reason:` + fmt.Sprintf("%v", this.Reason) + `,`,
		`Regarding:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Regarding), "ObjectReference", "v11.ObjectReference", 1), `&`, ``, 1) + `,`,
		`Related:` + strings.Replace(fmt.Sprintf("%v", this.Related), "ObjectReference", "v11.ObjectReference", 1) + `,`,
		`Note:` + fmt.Sprintf("%v", this.Note) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`DeprecatedSource:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.DeprecatedSource), "EventSource", "v11.EventSource", 1), `&`, ``, 1) + `,`,
		`DeprecatedFirstTimestamp:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.DeprecatedFirstTimestamp), "Time", "v1.Time", 1), `&`, ``, 1) + `,`,
		`DeprecatedLastTimestamp:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.DeprecatedLastTimestamp), "Time", "v1.Time", 1), `&`, ``, 1) + `,`,
		`DeprecatedCount:` + fmt.Sprintf("%v", this.DeprecatedCount) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EventList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]Event{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "Event", "Event", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&EventList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *EventSeries) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EventSeries{`,
		`Count:` + fmt.Sprintf("%v", this.Count) + `,`,
		`LastObservedTime:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.LastObservedTime), "MicroTime", "v1.MicroTime", 1), `&`, ``, 1) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EventTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Series", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Series == nil {
				m.Series = &EventSeries{}
			}
			if err := m.Series.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportingController", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReportingController = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportingInstance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReportingInstance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regarding", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Regarding.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Related", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Related == nil {
				m.Related = &v11.ObjectReference{}
			}
			if err := m.Related.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Note", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Note = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedSource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeprecatedSource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedFirstTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeprecatedFirstTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedLastTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeprecatedLastTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedCount", wireType)
			}
			m.DeprecatedCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeprecatedCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, Event{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventSeries) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventSeries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSeries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastObservedTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastObservedTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = EventSeriesState(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGenerated
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)
