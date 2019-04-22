// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/bbq/video/model/grpc/bvc.proto

/*
	Package grpc is a generated protocol buffer package.

	It is generated from these files:
		app/service/bbq/video/model/grpc/bvc.proto

	It has these top-level messages:
		RequestMsg
		VideoKeyItem
		ResponseMsg
*/
package grpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc1 "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"
import sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RequestMsg struct {
	Keys     []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Platform string   `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	UIP      string   `protobuf:"bytes,3,opt,name=uip,proto3" json:"uip,omitempty"`
	Uiplong  uint32   `protobuf:"varint,4,opt,name=uiplong,proto3" json:"uiplong,omitempty"`
	Backup   uint32   `protobuf:"varint,5,opt,name=backup,proto3" json:"backup,omitempty"`
	UUID     string   `protobuf:"bytes,6,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (m *RequestMsg) Reset()                    { *m = RequestMsg{} }
func (*RequestMsg) ProtoMessage()               {}
func (*RequestMsg) Descriptor() ([]byte, []int) { return fileDescriptorBvc, []int{0} }

type VideoKeyItem struct {
	Etime uint32   `protobuf:"varint,1,opt,name=etime,proto3" json:"etime,omitempty"`
	URL   []string `protobuf:"bytes,2,rep,name=url" json:"url,omitempty"`
}

func (m *VideoKeyItem) Reset()                    { *m = VideoKeyItem{} }
func (*VideoKeyItem) ProtoMessage()               {}
func (*VideoKeyItem) Descriptor() ([]byte, []int) { return fileDescriptorBvc, []int{1} }

type ResponseMsg struct {
	Code uint32                   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data map[string]*VideoKeyItem `protobuf:"bytes,2,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ResponseMsg) Reset()                    { *m = ResponseMsg{} }
func (*ResponseMsg) ProtoMessage()               {}
func (*ResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptorBvc, []int{2} }

func init() {
	proto.RegisterType((*RequestMsg)(nil), "video.vod.playurlbbq.RequestMsg")
	proto.RegisterType((*VideoKeyItem)(nil), "video.vod.playurlbbq.VideoKeyItem")
	proto.RegisterType((*ResponseMsg)(nil), "video.vod.playurlbbq.ResponseMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for PlayurlService service

type PlayurlServiceClient interface {
	ProtobufPlayurl(ctx context.Context, in *RequestMsg, opts ...grpc1.CallOption) (*ResponseMsg, error)
}

type playurlServiceClient struct {
	cc *grpc1.ClientConn
}

func NewPlayurlServiceClient(cc *grpc1.ClientConn) PlayurlServiceClient {
	return &playurlServiceClient{cc}
}

func (c *playurlServiceClient) ProtobufPlayurl(ctx context.Context, in *RequestMsg, opts ...grpc1.CallOption) (*ResponseMsg, error) {
	out := new(ResponseMsg)
	err := grpc1.Invoke(ctx, "/video.vod.playurlbbq.PlayurlService/ProtobufPlayurl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PlayurlService service

type PlayurlServiceServer interface {
	ProtobufPlayurl(context.Context, *RequestMsg) (*ResponseMsg, error)
}

func RegisterPlayurlServiceServer(s *grpc1.Server, srv PlayurlServiceServer) {
	s.RegisterService(&_PlayurlService_serviceDesc, srv)
}

func _PlayurlService_ProtobufPlayurl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayurlServiceServer).ProtobufPlayurl(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.vod.playurlbbq.PlayurlService/ProtobufPlayurl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayurlServiceServer).ProtobufPlayurl(ctx, req.(*RequestMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlayurlService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "video.vod.playurlbbq.PlayurlService",
	HandlerType: (*PlayurlServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "ProtobufPlayurl",
			Handler:    _PlayurlService_ProtobufPlayurl_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "app/service/bbq/video/model/grpc/bvc.proto",
}

func (m *RequestMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Keys) > 0 {
		for _, s := range m.Keys {
			dAtA[i] = 0xa
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
	if len(m.Platform) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBvc(dAtA, i, uint64(len(m.Platform)))
		i += copy(dAtA[i:], m.Platform)
	}
	if len(m.UIP) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBvc(dAtA, i, uint64(len(m.UIP)))
		i += copy(dAtA[i:], m.UIP)
	}
	if m.Uiplong != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintBvc(dAtA, i, uint64(m.Uiplong))
	}
	if m.Backup != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintBvc(dAtA, i, uint64(m.Backup))
	}
	if len(m.UUID) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintBvc(dAtA, i, uint64(len(m.UUID)))
		i += copy(dAtA[i:], m.UUID)
	}
	return i, nil
}

func (m *VideoKeyItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VideoKeyItem) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Etime != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBvc(dAtA, i, uint64(m.Etime))
	}
	if len(m.URL) > 0 {
		for _, s := range m.URL {
			dAtA[i] = 0x12
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
	return i, nil
}

func (m *ResponseMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBvc(dAtA, i, uint64(m.Code))
	}
	if len(m.Data) > 0 {
		for k, _ := range m.Data {
			dAtA[i] = 0x12
			i++
			v := m.Data[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovBvc(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovBvc(uint64(len(k))) + msgSize
			i = encodeVarintBvc(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintBvc(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintBvc(dAtA, i, uint64(v.Size()))
				n1, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	return i, nil
}

func encodeVarintBvc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RequestMsg) Size() (n int) {
	var l int
	_ = l
	if len(m.Keys) > 0 {
		for _, s := range m.Keys {
			l = len(s)
			n += 1 + l + sovBvc(uint64(l))
		}
	}
	l = len(m.Platform)
	if l > 0 {
		n += 1 + l + sovBvc(uint64(l))
	}
	l = len(m.UIP)
	if l > 0 {
		n += 1 + l + sovBvc(uint64(l))
	}
	if m.Uiplong != 0 {
		n += 1 + sovBvc(uint64(m.Uiplong))
	}
	if m.Backup != 0 {
		n += 1 + sovBvc(uint64(m.Backup))
	}
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovBvc(uint64(l))
	}
	return n
}

func (m *VideoKeyItem) Size() (n int) {
	var l int
	_ = l
	if m.Etime != 0 {
		n += 1 + sovBvc(uint64(m.Etime))
	}
	if len(m.URL) > 0 {
		for _, s := range m.URL {
			l = len(s)
			n += 1 + l + sovBvc(uint64(l))
		}
	}
	return n
}

func (m *ResponseMsg) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovBvc(uint64(m.Code))
	}
	if len(m.Data) > 0 {
		for k, v := range m.Data {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovBvc(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovBvc(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovBvc(uint64(mapEntrySize))
		}
	}
	return n
}

func sovBvc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBvc(x uint64) (n int) {
	return sovBvc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RequestMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RequestMsg{`,
		`Keys:` + fmt.Sprintf("%v", this.Keys) + `,`,
		`Platform:` + fmt.Sprintf("%v", this.Platform) + `,`,
		`UIP:` + fmt.Sprintf("%v", this.UIP) + `,`,
		`Uiplong:` + fmt.Sprintf("%v", this.Uiplong) + `,`,
		`Backup:` + fmt.Sprintf("%v", this.Backup) + `,`,
		`UUID:` + fmt.Sprintf("%v", this.UUID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VideoKeyItem) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VideoKeyItem{`,
		`Etime:` + fmt.Sprintf("%v", this.Etime) + `,`,
		`URL:` + fmt.Sprintf("%v", this.URL) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ResponseMsg) String() string {
	if this == nil {
		return "nil"
	}
	keysForData := make([]string, 0, len(this.Data))
	for k, _ := range this.Data {
		keysForData = append(keysForData, k)
	}
	sortkeys.Strings(keysForData)
	mapStringForData := "map[string]*VideoKeyItem{"
	for _, k := range keysForData {
		mapStringForData += fmt.Sprintf("%v: %v,", k, this.Data[k])
	}
	mapStringForData += "}"
	s := strings.Join([]string{`&ResponseMsg{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Data:` + mapStringForData + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBvc(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RequestMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBvc
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
			return fmt.Errorf("proto: RequestMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBvc
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
				return ErrInvalidLengthBvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platform", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBvc
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
				return ErrInvalidLengthBvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Platform = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UIP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBvc
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
				return ErrInvalidLengthBvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UIP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uiplong", wireType)
			}
			m.Uiplong = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uiplong |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Backup", wireType)
			}
			m.Backup = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Backup |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBvc
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
				return ErrInvalidLengthBvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBvc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBvc
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
func (m *VideoKeyItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBvc
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
			return fmt.Errorf("proto: VideoKeyItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VideoKeyItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Etime", wireType)
			}
			m.Etime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Etime |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBvc
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
				return ErrInvalidLengthBvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URL = append(m.URL, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBvc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBvc
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
func (m *ResponseMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBvc
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
			return fmt.Errorf("proto: ResponseMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBvc
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
				return ErrInvalidLengthBvc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = make(map[string]*VideoKeyItem)
			}
			var mapkey string
			var mapvalue *VideoKeyItem
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBvc
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBvc
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthBvc
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBvc
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthBvc
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthBvc
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &VideoKeyItem{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipBvc(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthBvc
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Data[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBvc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBvc
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
func skipBvc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBvc
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
					return 0, ErrIntOverflowBvc
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
					return 0, ErrIntOverflowBvc
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
				return 0, ErrInvalidLengthBvc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBvc
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
				next, err := skipBvc(dAtA[start:])
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
	ErrInvalidLengthBvc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBvc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("app/service/bbq/video/model/grpc/bvc.proto", fileDescriptorBvc) }

var fileDescriptorBvc = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3d, 0x6f, 0xdb, 0x30,
	0x10, 0x0d, 0x2d, 0xd9, 0x89, 0xcf, 0xfd, 0x02, 0x11, 0x14, 0xac, 0x51, 0x28, 0xae, 0x27, 0xa3,
	0x45, 0x25, 0xc0, 0x5d, 0x82, 0x2e, 0x01, 0x82, 0x74, 0x30, 0xda, 0x02, 0x06, 0x0b, 0x67, 0x68,
	0x27, 0x52, 0x62, 0x14, 0xc1, 0x92, 0x49, 0x4b, 0xa4, 0x00, 0x6d, 0xfd, 0x37, 0xfd, 0x19, 0x5d,
	0x33, 0x76, 0xec, 0x14, 0x34, 0xfa, 0x25, 0x85, 0x28, 0x39, 0xc8, 0xe0, 0xa2, 0xdb, 0xbd, 0x77,
	0xbc, 0xc7, 0xf7, 0xc8, 0x83, 0xd7, 0x4c, 0xa9, 0xa0, 0x10, 0x79, 0x99, 0x84, 0x22, 0xe0, 0x7c,
	0x1b, 0x94, 0x49, 0x24, 0x64, 0x90, 0xc9, 0x48, 0xa4, 0x41, 0x9c, 0xab, 0x30, 0xe0, 0x65, 0xe8,
	0xab, 0x5c, 0x6a, 0x89, 0x8f, 0x6d, 0xcf, 0x2f, 0x65, 0xe4, 0xab, 0x94, 0x55, 0x26, 0x4f, 0x39,
	0xdf, 0x8e, 0xdf, 0xc6, 0x89, 0xbe, 0x36, 0xdc, 0x0f, 0x65, 0x16, 0xc4, 0x32, 0x96, 0x81, 0x3d,
	0xcc, 0xcd, 0x95, 0x45, 0x16, 0xd8, 0xaa, 0x15, 0x99, 0xfe, 0x40, 0x00, 0x54, 0x6c, 0x8d, 0x28,
	0xf4, 0xe7, 0x22, 0xc6, 0x18, 0xdc, 0xb5, 0xa8, 0x0a, 0x82, 0x26, 0xce, 0x6c, 0x48, 0x6d, 0x8d,
	0xc7, 0x70, 0xa4, 0x52, 0xa6, 0xaf, 0x64, 0x9e, 0x91, 0xde, 0x04, 0xcd, 0x86, 0xf4, 0x1e, 0xe3,
	0x17, 0xe0, 0x98, 0x44, 0x11, 0xa7, 0xa1, 0xcf, 0x0f, 0xeb, 0xdb, 0x13, 0x67, 0xb5, 0x58, 0xd2,
	0x86, 0xc3, 0x04, 0x0e, 0x4d, 0xa2, 0x52, 0xb9, 0x89, 0x89, 0x3b, 0x41, 0xb3, 0xc7, 0x74, 0x07,
	0xf1, 0x73, 0x18, 0x70, 0x16, 0xae, 0x8d, 0x22, 0x7d, 0xdb, 0xe8, 0x10, 0x7e, 0x09, 0xae, 0x31,
	0x49, 0x44, 0x06, 0x56, 0xed, 0xa8, 0xbe, 0x3d, 0x71, 0x57, 0xab, 0xc5, 0x05, 0xb5, 0xec, 0xf4,
	0x0c, 0x1e, 0x5d, 0x36, 0x81, 0x3f, 0x8a, 0x6a, 0xa1, 0x45, 0x86, 0x8f, 0xa1, 0x2f, 0x74, 0x92,
	0x09, 0x82, 0xac, 0x48, 0x0b, 0xac, 0xa1, 0x3c, 0x25, 0xbd, 0xc6, 0x7f, 0x67, 0x88, 0x7e, 0xa2,
	0x0d, 0x37, 0xfd, 0x89, 0x60, 0x44, 0x45, 0xa1, 0xe4, 0xa6, 0x10, 0x5d, 0xd6, 0x50, 0x46, 0xbb,
	0x79, 0x5b, 0xe3, 0x33, 0x70, 0x23, 0xa6, 0x99, 0x9d, 0x1f, 0xcd, 0xdf, 0xf8, 0xfb, 0x9e, 0xd8,
	0x7f, 0x20, 0xe2, 0x5f, 0x30, 0xcd, 0x3e, 0x6c, 0x74, 0x5e, 0x51, 0x3b, 0x38, 0xfe, 0x06, 0xc3,
	0x7b, 0x0a, 0x3f, 0x03, 0x67, 0x2d, 0x2a, 0x7b, 0xc1, 0x90, 0x36, 0x25, 0x3e, 0x85, 0x7e, 0xc9,
	0x52, 0x23, 0xec, 0x43, 0x8e, 0xe6, 0xd3, 0xfd, 0x17, 0x3c, 0xcc, 0x49, 0xdb, 0x81, 0xf7, 0xbd,
	0x53, 0x34, 0xbf, 0x86, 0x27, 0xcb, 0xf6, 0xd4, 0x97, 0x76, 0x45, 0xf0, 0x25, 0x3c, 0x5d, 0x76,
	0xff, 0xdb, 0x75, 0xf0, 0xe4, 0x5f, 0xa6, 0x77, 0x9f, 0x3c, 0x7e, 0xf5, 0xdf, 0x58, 0xe7, 0xde,
	0xcd, 0x9d, 0x77, 0xf0, 0xfb, 0xce, 0x3b, 0xf8, 0x5e, 0x7b, 0xe8, 0xa6, 0xf6, 0xd0, 0xaf, 0xda,
	0x43, 0x7f, 0x6a, 0x0f, 0x7d, 0x75, 0x9b, 0x2d, 0xe4, 0x03, 0xbb, 0x3d, 0xef, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xc1, 0x6e, 0x03, 0xe2, 0xb0, 0x02, 0x00, 0x00,
}