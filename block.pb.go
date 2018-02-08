// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/geoah/go-nimona-identity/generated.proto

/*
	Package identity is a generated protocol buffer package.

	It is generated from these files:
		github.com/geoah/go-nimona-identity/generated.proto

	It has these top-level messages:
		Block
		BlockEvent
		BlockEventPayload
*/
package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// EventType for the BlockEvents
var EventType_name = map[int32]string{
	0: "EVENT_TYPE_GRAPH_CREATE",
}
var EventType_value = map[string]int32{
	"EVENT_TYPE_GRAPH_CREATE": 0,
}

func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *BlockEvent) Reset()                    { *m = BlockEvent{} }
func (m *BlockEvent) String() string            { return proto.CompactTextString(m) }
func (*BlockEvent) ProtoMessage()               {}
func (*BlockEvent) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *BlockEventPayload) Reset()                    { *m = BlockEventPayload{} }
func (m *BlockEventPayload) String() string            { return proto.CompactTextString(m) }
func (*BlockEventPayload) ProtoMessage()               {}
func (*BlockEventPayload) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*Block)(nil), "github.com.geoah.gonimonaidentity.Block")
	proto.RegisterType((*BlockEvent)(nil), "github.com.geoah.gonimonaidentity.BlockEvent")
	proto.RegisterType((*BlockEventPayload)(nil), "github.com.geoah.gonimonaidentity.BlockEventPayload")
	proto.RegisterEnum("github.com.geoah.gonimonaidentity.EventType", EventType_name, EventType_value)
}
func (m *Block) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Event.ProtoSize()))
	n1, err := m.Event.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Signature) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	if len(m.Hash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	return i, nil
}

func (m *BlockEvent) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Author) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Author)))
		i += copy(dAtA[i:], m.Author)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Payload.ProtoSize()))
	n2, err := m.Payload.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Nonce) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Nonce)))
		i += copy(dAtA[i:], m.Nonce)
	}
	if len(m.Parents) > 0 {
		for _, s := range m.Parents {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	return i, nil
}

func (m *BlockEventPayload) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockEventPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Content) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Content)))
		i += copy(dAtA[i:], m.Content)
	}
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Block) ProtoSize() (n int) {
	var l int
	_ = l
	l = m.Event.ProtoSize()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *BlockEvent) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.Author)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = m.Payload.ProtoSize()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Parents) > 0 {
		for _, s := range m.Parents {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *BlockEventPayload) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Block) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
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
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
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
func (m *BlockEvent) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Author = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
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
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parents", wireType)
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parents = append(m.Parents, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
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
func (m *BlockEventPayload) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockEventPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockEventPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
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
			iNdEx += length
			if length < 0 {
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

func init() {
	proto.RegisterFile("github.com/geoah/go-nimona-identity/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x8a, 0xd3, 0x40,
	0x1c, 0xc6, 0x67, 0xdc, 0x66, 0xdd, 0xcc, 0xee, 0x41, 0x87, 0x45, 0x43, 0x91, 0x69, 0xdc, 0x53,
	0x11, 0x92, 0xc0, 0xae, 0x5e, 0xbc, 0xb5, 0x25, 0x54, 0x11, 0x24, 0x84, 0x20, 0xe8, 0xa5, 0x4c,
	0xd3, 0x31, 0x09, 0xb6, 0x33, 0x21, 0x9d, 0x08, 0x7d, 0x83, 0xe2, 0xc1, 0xb3, 0x17, 0xa1, 0x62,
	0x0f, 0x3e, 0x86, 0xc7, 0x1e, 0x7d, 0x02, 0x91, 0xf6, 0x05, 0x7c, 0x04, 0x99, 0x49, 0xd3, 0x14,
	0x3c, 0x88, 0xb7, 0xff, 0x97, 0xe4, 0xf7, 0x7d, 0xff, 0xef, 0x4f, 0xd0, 0x4d, 0x92, 0xc9, 0xb4,
	0x1c, 0xbb, 0xb1, 0x98, 0x79, 0x09, 0x13, 0x34, 0xf5, 0x12, 0xe1, 0xf0, 0x6c, 0x26, 0x38, 0x75,
	0xb2, 0x09, 0xe3, 0x32, 0x93, 0x0b, 0x2f, 0x61, 0x9c, 0x15, 0x54, 0xb2, 0x89, 0x9b, 0x17, 0x42,
	0x0a, 0xfc, 0xb0, 0x81, 0x5c, 0x0d, 0xb9, 0x89, 0xa8, 0x98, 0x1a, 0x69, 0x3b, 0xc7, 0xbe, 0x22,
	0x11, 0x9e, 0x26, 0xc7, 0xe5, 0x5b, 0xad, 0xb4, 0xd0, 0x53, 0xe5, 0x78, 0xf5, 0x11, 0x22, 0xa3,
	0x3f, 0x15, 0xf1, 0x3b, 0xfc, 0x1c, 0x19, 0xec, 0x3d, 0xe3, 0xd2, 0x82, 0x36, 0xec, 0x9e, 0x5f,
	0x3b, 0xee, 0x3f, 0xb3, 0x5c, 0x0d, 0xfa, 0x0a, 0xea, 0xb7, 0x36, 0x3f, 0x3b, 0x20, 0xac, 0x1c,
	0xf0, 0x03, 0x64, 0xce, 0xb3, 0x84, 0x53, 0x59, 0x16, 0xcc, 0xba, 0x65, 0xc3, 0xae, 0x19, 0x36,
	0x0f, 0x30, 0x46, 0xad, 0x94, 0xce, 0x53, 0xeb, 0x44, 0xbf, 0xd0, 0xf3, 0xd3, 0xb3, 0xe5, 0xaa,
	0x03, 0x7e, 0x7f, 0xe9, 0x80, 0xab, 0x0d, 0x44, 0xa8, 0xf1, 0xc5, 0xf7, 0xd0, 0x29, 0x2d, 0x65,
	0x2a, 0x0a, 0xbd, 0x96, 0x19, 0xee, 0x15, 0x8e, 0xd0, 0xed, 0x9c, 0x2e, 0xa6, 0x82, 0x4e, 0x74,
	0xc0, 0xf9, 0xf5, 0xe3, 0xff, 0xda, 0x37, 0xa8, 0xd8, 0xfd, 0xda, 0xb5, 0x15, 0xbe, 0x44, 0x06,
	0x17, 0x3c, 0x66, 0xfb, 0xdd, 0x2a, 0x81, 0x2d, 0x95, 0x55, 0x30, 0x2e, 0xe7, 0x56, 0xcb, 0x3e,
	0xe9, 0x9a, 0x61, 0x2d, 0x55, 0x15, 0xb9, 0xc8, 0x99, 0x65, 0x54, 0x55, 0xd4, 0x7c, 0x54, 0xe5,
	0x05, 0xba, 0xfb, 0x57, 0xe2, 0x01, 0x81, 0x0d, 0xa2, 0x02, 0x62, 0xc1, 0xa5, 0x3a, 0xbe, 0x2a,
	0x73, 0x11, 0xd6, 0xb2, 0x31, 0x7b, 0x14, 0x20, 0x53, 0xfb, 0x44, 0x0a, 0x78, 0x82, 0xee, 0xfb,
	0xaf, 0xfc, 0x97, 0xd1, 0x28, 0x7a, 0x1d, 0xf8, 0xa3, 0x61, 0xd8, 0x0b, 0x9e, 0x8d, 0x06, 0xa1,
	0xdf, 0x8b, 0xfc, 0x3b, 0xa0, 0x6d, 0x7d, 0xf8, 0x6c, 0x5f, 0x1e, 0xbe, 0x1d, 0x16, 0x34, 0x4f,
	0x07, 0x05, 0xa3, 0x92, 0xb5, 0x2f, 0x96, 0x5f, 0x09, 0xf8, 0xb6, 0x26, 0xe0, 0xfb, 0x9a, 0x80,
	0x3e, 0xd9, 0x6c, 0x09, 0xfc, 0xb1, 0x25, 0xf0, 0xd7, 0x96, 0x80, 0x4f, 0x3b, 0x02, 0x56, 0x3b,
	0x02, 0xdf, 0x9c, 0xd5, 0xd7, 0x1a, 0x9f, 0xea, 0x3f, 0xe4, 0xe6, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8a, 0x74, 0xd3, 0xaf, 0xaa, 0x02, 0x00, 0x00,
}
