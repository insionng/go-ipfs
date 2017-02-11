// Code generated by protoc-gen-gogo.
// source: diagnostics.proto
// DO NOT EDIT!

/*
Package diagnostics_pb is a generated protocol buffer package.

It is generated from these files:
	diagnostics.proto

It has these top-level messages:
	Message
*/
package diagnostics_pb

import proto "gx/ipfs/QmZ4Qi3GaRbjcx28Sme5eMH7RQjGkt8wHxt2a65oLaeFEV/gogo-protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Message struct {
	DiagID           *string `protobuf:"bytes,1,req" json:"DiagID,omitempty"`
	Data             []byte  `protobuf:"bytes,2,opt" json:"Data,omitempty"`
	Timeout          *int64  `protobuf:"varint,3,opt" json:"Timeout,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

func (m *Message) GetDiagID() string {
	if m != nil && m.DiagID != nil {
		return *m.DiagID
	}
	return ""
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetTimeout() int64 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func init() {
}
