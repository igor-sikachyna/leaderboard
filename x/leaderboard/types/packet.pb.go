// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: leaderboard/leaderboard/packet.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LeaderboardPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*LeaderboardPacketData_NoData
	//	*LeaderboardPacketData_IbcTopRankPacket
	Packet isLeaderboardPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *LeaderboardPacketData) Reset()         { *m = LeaderboardPacketData{} }
func (m *LeaderboardPacketData) String() string { return proto.CompactTextString(m) }
func (*LeaderboardPacketData) ProtoMessage()    {}
func (*LeaderboardPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_384b6f5a4b1b8232, []int{0}
}
func (m *LeaderboardPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaderboardPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaderboardPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaderboardPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaderboardPacketData.Merge(m, src)
}
func (m *LeaderboardPacketData) XXX_Size() int {
	return m.Size()
}
func (m *LeaderboardPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaderboardPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_LeaderboardPacketData proto.InternalMessageInfo

type isLeaderboardPacketData_Packet interface {
	isLeaderboardPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type LeaderboardPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type LeaderboardPacketData_IbcTopRankPacket struct {
	IbcTopRankPacket *IbcTopRankPacketData `protobuf:"bytes,2,opt,name=ibcTopRankPacket,proto3,oneof" json:"ibcTopRankPacket,omitempty"`
}

func (*LeaderboardPacketData_NoData) isLeaderboardPacketData_Packet()           {}
func (*LeaderboardPacketData_IbcTopRankPacket) isLeaderboardPacketData_Packet() {}

func (m *LeaderboardPacketData) GetPacket() isLeaderboardPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *LeaderboardPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*LeaderboardPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *LeaderboardPacketData) GetIbcTopRankPacket() *IbcTopRankPacketData {
	if x, ok := m.GetPacket().(*LeaderboardPacketData_IbcTopRankPacket); ok {
		return x.IbcTopRankPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LeaderboardPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LeaderboardPacketData_NoData)(nil),
		(*LeaderboardPacketData_IbcTopRankPacket)(nil),
	}
}

type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_384b6f5a4b1b8232, []int{1}
}
func (m *NoData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoData.Merge(m, src)
}
func (m *NoData) XXX_Size() int {
	return m.Size()
}
func (m *NoData) XXX_DiscardUnknown() {
	xxx_messageInfo_NoData.DiscardUnknown(m)
}

var xxx_messageInfo_NoData proto.InternalMessageInfo

// IbcTopRankPacketData defines a struct for the packet payload
type IbcTopRankPacketData struct {
	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Rank     uint64 `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Score    uint64 `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (m *IbcTopRankPacketData) Reset()         { *m = IbcTopRankPacketData{} }
func (m *IbcTopRankPacketData) String() string { return proto.CompactTextString(m) }
func (*IbcTopRankPacketData) ProtoMessage()    {}
func (*IbcTopRankPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_384b6f5a4b1b8232, []int{2}
}
func (m *IbcTopRankPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IbcTopRankPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IbcTopRankPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IbcTopRankPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IbcTopRankPacketData.Merge(m, src)
}
func (m *IbcTopRankPacketData) XXX_Size() int {
	return m.Size()
}
func (m *IbcTopRankPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_IbcTopRankPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_IbcTopRankPacketData proto.InternalMessageInfo

func (m *IbcTopRankPacketData) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *IbcTopRankPacketData) GetRank() uint64 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *IbcTopRankPacketData) GetScore() uint64 {
	if m != nil {
		return m.Score
	}
	return 0
}

// IbcTopRankPacketAck defines a struct for the packet acknowledgment
type IbcTopRankPacketAck struct {
	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (m *IbcTopRankPacketAck) Reset()         { *m = IbcTopRankPacketAck{} }
func (m *IbcTopRankPacketAck) String() string { return proto.CompactTextString(m) }
func (*IbcTopRankPacketAck) ProtoMessage()    {}
func (*IbcTopRankPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_384b6f5a4b1b8232, []int{3}
}
func (m *IbcTopRankPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IbcTopRankPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IbcTopRankPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IbcTopRankPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IbcTopRankPacketAck.Merge(m, src)
}
func (m *IbcTopRankPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *IbcTopRankPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_IbcTopRankPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_IbcTopRankPacketAck proto.InternalMessageInfo

func (m *IbcTopRankPacketAck) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func init() {
	proto.RegisterType((*LeaderboardPacketData)(nil), "leaderboard.leaderboard.LeaderboardPacketData")
	proto.RegisterType((*NoData)(nil), "leaderboard.leaderboard.NoData")
	proto.RegisterType((*IbcTopRankPacketData)(nil), "leaderboard.leaderboard.IbcTopRankPacketData")
	proto.RegisterType((*IbcTopRankPacketAck)(nil), "leaderboard.leaderboard.IbcTopRankPacketAck")
}

func init() {
	proto.RegisterFile("leaderboard/leaderboard/packet.proto", fileDescriptor_384b6f5a4b1b8232)
}

var fileDescriptor_384b6f5a4b1b8232 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0x49, 0x4d, 0x4c,
	0x49, 0x2d, 0x4a, 0xca, 0x4f, 0x2c, 0x4a, 0xd1, 0x47, 0x66, 0x17, 0x24, 0x26, 0x67, 0xa7, 0x96,
	0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x23, 0xc9, 0xe8, 0x21, 0xb1, 0x95, 0x76, 0x32,
	0x72, 0x89, 0xfa, 0x20, 0xf8, 0x01, 0x60, 0x4d, 0x2e, 0x89, 0x25, 0x89, 0x42, 0x96, 0x5c, 0x6c,
	0x79, 0xf9, 0x20, 0x96, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xbc, 0x1e, 0x0e, 0x33, 0xf4,
	0xfc, 0xc0, 0xca, 0x3c, 0x18, 0x82, 0xa0, 0x1a, 0x84, 0xa2, 0xb9, 0x04, 0x32, 0x93, 0x92, 0x43,
	0xf2, 0x0b, 0x82, 0x12, 0xf3, 0xb2, 0x21, 0x46, 0x4a, 0x30, 0x81, 0x0d, 0xd1, 0xc5, 0x69, 0x88,
	0x27, 0x9a, 0x06, 0xa8, 0x91, 0x18, 0x06, 0x39, 0x71, 0x70, 0xb1, 0x41, 0xbc, 0xa6, 0xc4, 0xc1,
	0xc5, 0x06, 0xb1, 0x5a, 0x29, 0x86, 0x4b, 0x04, 0x9b, 0x7e, 0x21, 0x29, 0x2e, 0x8e, 0x82, 0x9c,
	0xc4, 0xca, 0xd4, 0x22, 0xcf, 0x14, 0xb0, 0x2f, 0x38, 0x83, 0xe0, 0x7c, 0x21, 0x21, 0x2e, 0x96,
	0xa2, 0xc4, 0xbc, 0x6c, 0xb0, 0xc3, 0x58, 0x82, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xe2, 0xe4,
	0xfc, 0xa2, 0x54, 0x09, 0x66, 0xb0, 0x20, 0x84, 0xa3, 0x64, 0xc8, 0x25, 0x8c, 0x6e, 0xba, 0x63,
	0x72, 0x36, 0x3e, 0xc3, 0x9d, 0x82, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1,
	0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21,
	0xca, 0x32, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x33, 0x3d, 0xbf,
	0x48, 0xb7, 0x38, 0x33, 0x3b, 0x31, 0x39, 0xa3, 0x32, 0x2f, 0x11, 0x25, 0xf6, 0x2a, 0x50, 0x78,
	0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xb8, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x04, 0x16, 0xcd, 0x6a, 0xf3, 0x01, 0x00, 0x00,
}

func (m *LeaderboardPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaderboardPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaderboardPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *LeaderboardPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaderboardPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *LeaderboardPacketData_IbcTopRankPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaderboardPacketData_IbcTopRankPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IbcTopRankPacket != nil {
		{
			size, err := m.IbcTopRankPacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *IbcTopRankPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IbcTopRankPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IbcTopRankPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Score != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Score))
		i--
		dAtA[i] = 0x18
	}
	if m.Rank != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Rank))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PlayerId) > 0 {
		i -= len(m.PlayerId)
		copy(dAtA[i:], m.PlayerId)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.PlayerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IbcTopRankPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IbcTopRankPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IbcTopRankPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PlayerId) > 0 {
		i -= len(m.PlayerId)
		copy(dAtA[i:], m.PlayerId)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.PlayerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LeaderboardPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *LeaderboardPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *LeaderboardPacketData_IbcTopRankPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IbcTopRankPacket != nil {
		l = m.IbcTopRankPacket.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *IbcTopRankPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PlayerId)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.Rank != 0 {
		n += 1 + sovPacket(uint64(m.Rank))
	}
	if m.Score != 0 {
		n += 1 + sovPacket(uint64(m.Score))
	}
	return n
}

func (m *IbcTopRankPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PlayerId)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LeaderboardPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: LeaderboardPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaderboardPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &LeaderboardPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcTopRankPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &IbcTopRankPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &LeaderboardPacketData_IbcTopRankPacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *NoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: NoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *IbcTopRankPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: IbcTopRankPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IbcTopRankPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			m.Rank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rank |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *IbcTopRankPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: IbcTopRankPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IbcTopRankPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
