package pb

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *Span) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 12
	// string "service"
	o = append(o, 0x8c, 0xa7, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65)
	o = msgp.AppendString(o, z.Service)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "resource"
	o = append(o, 0xa8, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65)
	o = msgp.AppendString(o, z.Resource)
	// string "trace_id"
	o = append(o, 0xa8, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64)
	o = msgp.AppendUint64(o, z.TraceID)
	// string "span_id"
	o = append(o, 0xa7, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64)
	o = msgp.AppendUint64(o, z.SpanID)
	// string "parent_id"
	o = append(o, 0xa9, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64)
	o = msgp.AppendUint64(o, z.ParentID)
	// string "start"
	o = append(o, 0xa5, 0x73, 0x74, 0x61, 0x72, 0x74)
	o = msgp.AppendInt64(o, z.Start)
	// string "duration"
	o = append(o, 0xa8, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendInt64(o, z.Duration)
	// string "error"
	o = append(o, 0xa5, 0x65, 0x72, 0x72, 0x6f, 0x72)
	o = msgp.AppendInt32(o, z.Error)
	// string "meta"
	o = append(o, 0xa4, 0x6d, 0x65, 0x74, 0x61)
	o = msgp.AppendMapHeader(o, uint32(len(z.Meta)))
	for za0001, za0002 := range z.Meta {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendString(o, za0002)
	}
	// string "metrics"
	o = append(o, 0xa7, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Metrics)))
	for za0003, za0004 := range z.Metrics {
		o = msgp.AppendString(o, za0003)
		o = msgp.AppendFloat64(o, za0004)
	}
	// string "type"
	o = append(o, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Span) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "service":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.Service = ""
				break
			}
			z.Service, bts, err = parseStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.Name = ""
				break
			}
			z.Name, bts, err = parseStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "resource":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.Resource = ""
				break
			}
			z.Resource, bts, err = parseStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "trace_id":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.TraceID = 0
				break
			}
			z.TraceID, bts, err = parseUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TraceID")
				return
			}
		case "span_id":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.SpanID = 0
				break
			}
			z.SpanID, bts, err = parseUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SpanID")
				return
			}
		case "parent_id":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.ParentID = 0
				break
			}
			z.ParentID, bts, err = parseUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ParentID")
				return
			}
		case "start":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.Start = 0
				break
			}
			z.Start, bts, err = parseInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Start")
				return
			}
		case "duration":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.Duration = 0
				break
			}
			z.Duration, bts, err = parseInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Duration")
				return
			}
		case "error":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.Error = 0
				break
			}
			z.Error, bts, err = parseInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Error")
				return
			}
		case "meta":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.Meta = nil
				break
			}
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Meta")
				return
			}
			if z.Meta == nil && zb0002 > 0 {
				z.Meta = make(map[string]string, zb0002)
			} else if len(z.Meta) > 0 {
				for key := range z.Meta {
					delete(z.Meta, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 string
				zb0002--
				za0001, bts, err = parseStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Meta")
					return
				}
				za0002, bts, err = parseStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Meta", za0001)
					return
				}
				z.Meta[za0001] = za0002
			}
		case "metrics":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.Metrics = nil
				break
			}
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Metrics")
				return
			}
			if z.Metrics == nil && zb0003 > 0 {
				z.Metrics = make(map[string]float64, zb0003)
			} else if len(z.Metrics) > 0 {
				for key := range z.Metrics {
					delete(z.Metrics, key)
				}
			}
			for zb0003 > 0 {
				var za0003 string
				var za0004 float64
				zb0003--
				za0003, bts, err = parseStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Metrics")
					return
				}
				za0004, bts, err = parseFloat64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Metrics", za0003)
					return
				}
				z.Metrics[za0003] = za0004
			}
		case "type":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				z.Type = ""
				break
			}
			z.Type, bts, err = parseStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Type")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Span) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Service) + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.StringPrefixSize + len(z.Resource) + 9 + msgp.Uint64Size + 8 + msgp.Uint64Size + 10 + msgp.Uint64Size + 6 + msgp.Int64Size + 9 + msgp.Int64Size + 6 + msgp.Int32Size + 5 + msgp.MapHeaderSize
	if z.Meta != nil {
		for za0001, za0002 := range z.Meta {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += 8 + msgp.MapHeaderSize
	if z.Metrics != nil {
		for za0003, za0004 := range z.Metrics {
			_ = za0004
			s += msgp.StringPrefixSize + len(za0003) + msgp.Float64Size
		}
	}
	s += 5 + msgp.StringPrefixSize + len(z.Type)
	return
}
