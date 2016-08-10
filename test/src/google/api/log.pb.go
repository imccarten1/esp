// Code generated by protoc-gen-go.
// source: google/api/log.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A description of a log type. Example in YAML format:
//
//     - name: library.googleapis.com/activity_history
//       description: The history of borrowing and returning library items.
//       display_name: Activity
//       labels:
//       - key: /customer_id
//         description: Identifier of a library customer
type LogDescriptor struct {
	// The name of the log. It must be less than 512 characters long and can
	// include the following characters: upper- and lower-case alphanumeric
	// characters [A-Za-z0-9], and punctuation characters including
	// slash, underscore, hyphen, period [/_-.].
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The set of labels that are available to describe a specific log entry.
	// Runtime requests that contain labels not specified here are
	// considered invalid.
	Labels []*LabelDescriptor `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
	// A human-readable description of this log. This information appears in
	// the documentation and can contain details.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// The human-readable name for this log. This information appears on
	// the user interface and should be concise.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
}

func (m *LogDescriptor) Reset()                    { *m = LogDescriptor{} }
func (m *LogDescriptor) String() string            { return proto.CompactTextString(m) }
func (*LogDescriptor) ProtoMessage()               {}
func (*LogDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *LogDescriptor) GetLabels() []*LabelDescriptor {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*LogDescriptor)(nil), "google.api.LogDescriptor")
}

func init() { proto.RegisterFile("google/api/log.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0xc9, 0x4f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x82, 0x88, 0xea, 0x25, 0x16, 0x64, 0x4a, 0x89, 0x21, 0xab, 0x48, 0x4c, 0x4a, 0xcd,
	0x81, 0xa8, 0x51, 0x9a, 0xcb, 0xc8, 0xc5, 0xeb, 0x93, 0x9f, 0xee, 0x92, 0x5a, 0x9c, 0x5c, 0x94,
	0x59, 0x50, 0x92, 0x5f, 0x24, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x19, 0x73, 0xb1, 0x81, 0x35, 0x15, 0x4b, 0x30, 0x29, 0x30,
	0x6b, 0x70, 0x1b, 0x49, 0xeb, 0x21, 0x8c, 0xd6, 0xf3, 0x01, 0xc9, 0x20, 0x0c, 0x08, 0x82, 0x2a,
	0x15, 0x52, 0xe0, 0xe2, 0x4e, 0x81, 0x8a, 0x66, 0xe6, 0xe7, 0x49, 0x30, 0x83, 0xcd, 0x43, 0x16,
	0x12, 0x52, 0xe4, 0xe2, 0x49, 0xc9, 0x2c, 0x2e, 0xc8, 0x49, 0xac, 0x8c, 0x07, 0x5b, 0xc9, 0x02,
	0x55, 0x02, 0x11, 0xf3, 0x4b, 0xcc, 0x4d, 0x75, 0x92, 0xe1, 0xe2, 0x4b, 0xce, 0xcf, 0x45, 0xb2,
	0xce, 0x89, 0xc3, 0x27, 0x3f, 0x3d, 0x00, 0xe4, 0xf6, 0x00, 0xc6, 0x24, 0x36, 0xb0, 0x27, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xfa, 0xf4, 0xae, 0x00, 0x01, 0x00, 0x00,
}
