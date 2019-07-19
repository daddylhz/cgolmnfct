// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs types_linux.go

package cgolmnfct

type (
	Size_t  uint64
	Ssize_t int64
)

const SizeofAttrGrpIpv4 = 0x8

type AttrGrpIpv4 struct {
	Src uint32
	Dst uint32
}

const SizeofAttrGrpIpv6 = 0x20

type AttrGrpIpv6 struct {
	Src [4]uint32
	Dst [4]uint32
}

const SizeofAttrGrpPort = 0x4

type AttrGrpPort struct {
	Sport uint16
	Dport uint16
}

const SizeofAttrGrpIcmp = 0x4

type AttrGrpIcmp struct {
	Id   uint16
	Code uint8
	Type uint8
}

const SizeofAttrGrpCtrs = 0x10

type AttrGrpCtrs struct {
	Packets uint64
	Bytes   uint64
}

const SizeofAttrGrpAddr = 0x10

type AttrGrpAddr [16]byte

const SizeofFilterProto = 0x4

type FilterProto struct {
	Proto uint16
	State uint16
}

const SizeofFilterIpv4 = 0x8

type FilterIpv4 struct {
	Addr uint32
	Mask uint32
}

const SizeofFilterIpv6 = 0x20

type FilterIpv6 struct {
	Addr [4]uint32
	Mask [4]uint32
}

const SizeofFilterDumpMark = 0x8

type FilterDumpMark struct {
	Val  uint32
	Mask uint32
}
