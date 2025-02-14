package cgolmnfct

/*
#cgo CFLAGS: -I./include
#include <libnetfilter_conntrack/libnetfilter_conntrack.h>
*/
import "C"

const (
	CONNTRACK = C.CONNTRACK
	EXPECT    = C.EXPECT
)

const NFCT_ALL_CT_GROUPS = C.NFCT_ALL_CT_GROUPS

type NfctHandle C.struct_nfct_handle

// conntrack attributes
type ConntrackAttr C.enum_nf_conntrack_attr

const (
	ATTR_ORIG_IPV4_SRC               ConntrackAttr = C.ATTR_ORIG_IPV4_SRC
	ATTR_IPV4_SRC                    ConntrackAttr = C.ATTR_IPV4_SRC
	ATTR_ORIG_IPV4_DST               ConntrackAttr = C.ATTR_ORIG_IPV4_DST
	ATTR_IPV4_DST                    ConntrackAttr = C.ATTR_IPV4_DST
	ATTR_REPL_IPV4_SRC               ConntrackAttr = C.ATTR_REPL_IPV4_SRC
	ATTR_REPL_IPV4_DST               ConntrackAttr = C.ATTR_REPL_IPV4_DST
	ATTR_ORIG_IPV6_SRC               ConntrackAttr = C.ATTR_ORIG_IPV6_SRC
	ATTR_IPV6_SRC                    ConntrackAttr = C.ATTR_IPV6_SRC
	ATTR_ORIG_IPV6_DST               ConntrackAttr = C.ATTR_ORIG_IPV6_DST
	ATTR_IPV6_DST                    ConntrackAttr = C.ATTR_IPV6_DST
	ATTR_REPL_IPV6_SRC               ConntrackAttr = C.ATTR_REPL_IPV6_SRC
	ATTR_REPL_IPV6_DST               ConntrackAttr = C.ATTR_REPL_IPV6_DST
	ATTR_ORIG_PORT_SRC               ConntrackAttr = C.ATTR_ORIG_PORT_SRC
	ATTR_PORT_SRC                    ConntrackAttr = C.ATTR_PORT_SRC
	ATTR_ORIG_PORT_DST               ConntrackAttr = C.ATTR_ORIG_PORT_DST
	ATTR_PORT_DST                    ConntrackAttr = C.ATTR_PORT_DST
	ATTR_REPL_PORT_SRC               ConntrackAttr = C.ATTR_REPL_PORT_SRC
	ATTR_REPL_PORT_DST               ConntrackAttr = C.ATTR_REPL_PORT_DST
	ATTR_ICMP_TYPE                   ConntrackAttr = C.ATTR_ICMP_TYPE
	ATTR_ICMP_CODE                   ConntrackAttr = C.ATTR_ICMP_CODE
	ATTR_ICMP_ID                     ConntrackAttr = C.ATTR_ICMP_ID
	ATTR_ORIG_L3PROTO                ConntrackAttr = C.ATTR_ORIG_L3PROTO
	ATTR_L3PROTO                     ConntrackAttr = C.ATTR_L3PROTO
	ATTR_REPL_L3PROTO                ConntrackAttr = C.ATTR_REPL_L3PROTO
	ATTR_ORIG_L4PROTO                ConntrackAttr = C.ATTR_ORIG_L4PROTO
	ATTR_L4PROTO                     ConntrackAttr = C.ATTR_L4PROTO
	ATTR_REPL_L4PROTO                ConntrackAttr = C.ATTR_REPL_L4PROTO
	ATTR_TCP_STATE                   ConntrackAttr = C.ATTR_TCP_STATE
	ATTR_SNAT_IPV4                   ConntrackAttr = C.ATTR_SNAT_IPV4
	ATTR_DNAT_IPV4                   ConntrackAttr = C.ATTR_DNAT_IPV4
	ATTR_SNAT_PORT                   ConntrackAttr = C.ATTR_SNAT_PORT
	ATTR_DNAT_PORT                   ConntrackAttr = C.ATTR_DNAT_PORT
	ATTR_TIMEOUT                     ConntrackAttr = C.ATTR_TIMEOUT
	ATTR_MARK                        ConntrackAttr = C.ATTR_MARK
	ATTR_ORIG_COUNTER_PACKETS        ConntrackAttr = C.ATTR_ORIG_COUNTER_PACKETS
	ATTR_REPL_COUNTER_PACKETS        ConntrackAttr = C.ATTR_REPL_COUNTER_PACKETS
	ATTR_ORIG_COUNTER_BYTES          ConntrackAttr = C.ATTR_ORIG_COUNTER_BYTES
	ATTR_REPL_COUNTER_BYTES          ConntrackAttr = C.ATTR_REPL_COUNTER_BYTES
	ATTR_USE                         ConntrackAttr = C.ATTR_USE
	ATTR_ID                          ConntrackAttr = C.ATTR_ID
	ATTR_STATUS                      ConntrackAttr = C.ATTR_STATUS
	ATTR_TCP_FLAGS_ORIG              ConntrackAttr = C.ATTR_TCP_FLAGS_ORIG
	ATTR_TCP_FLAGS_REPL              ConntrackAttr = C.ATTR_TCP_FLAGS_REPL
	ATTR_TCP_MASK_ORIG               ConntrackAttr = C.ATTR_TCP_MASK_ORIG
	ATTR_TCP_MASK_REPL               ConntrackAttr = C.ATTR_TCP_MASK_REPL
	ATTR_MASTER_IPV4_SRC             ConntrackAttr = C.ATTR_MASTER_IPV4_SRC
	ATTR_MASTER_IPV4_DST             ConntrackAttr = C.ATTR_MASTER_IPV4_DST
	ATTR_MASTER_IPV6_SRC             ConntrackAttr = C.ATTR_MASTER_IPV6_SRC
	ATTR_MASTER_IPV6_DST             ConntrackAttr = C.ATTR_MASTER_IPV6_DST
	ATTR_MASTER_PORT_SRC             ConntrackAttr = C.ATTR_MASTER_PORT_SRC
	ATTR_MASTER_PORT_DST             ConntrackAttr = C.ATTR_MASTER_PORT_DST
	ATTR_MASTER_L3PROTO              ConntrackAttr = C.ATTR_MASTER_L3PROTO
	ATTR_MASTER_L4PROTO              ConntrackAttr = C.ATTR_MASTER_L4PROTO
	ATTR_SECMARK                     ConntrackAttr = C.ATTR_SECMARK
	ATTR_ORIG_NAT_SEQ_CORRECTION_POS ConntrackAttr = C.ATTR_ORIG_NAT_SEQ_CORRECTION_POS
	ATTR_ORIG_NAT_SEQ_OFFSET_BEFORE  ConntrackAttr = C.ATTR_ORIG_NAT_SEQ_OFFSET_BEFORE
	ATTR_ORIG_NAT_SEQ_OFFSET_AFTER   ConntrackAttr = C.ATTR_ORIG_NAT_SEQ_OFFSET_AFTER
	ATTR_REPL_NAT_SEQ_CORRECTION_POS ConntrackAttr = C.ATTR_REPL_NAT_SEQ_CORRECTION_POS
	ATTR_REPL_NAT_SEQ_OFFSET_BEFORE  ConntrackAttr = C.ATTR_REPL_NAT_SEQ_OFFSET_BEFORE
	ATTR_REPL_NAT_SEQ_OFFSET_AFTER   ConntrackAttr = C.ATTR_REPL_NAT_SEQ_OFFSET_AFTER
	ATTR_SCTP_STATE                  ConntrackAttr = C.ATTR_SCTP_STATE
	ATTR_SCTP_VTAG_ORIG              ConntrackAttr = C.ATTR_SCTP_VTAG_ORIG
	ATTR_SCTP_VTAG_REPL              ConntrackAttr = C.ATTR_SCTP_VTAG_REPL
	ATTR_HELPER_NAME                 ConntrackAttr = C.ATTR_HELPER_NAME
	ATTR_DCCP_STATE                  ConntrackAttr = C.ATTR_DCCP_STATE
	ATTR_DCCP_ROLE                   ConntrackAttr = C.ATTR_DCCP_ROLE
	ATTR_DCCP_HANDSHAKE_SEQ          ConntrackAttr = C.ATTR_DCCP_HANDSHAKE_SEQ
	ATTR_TCP_WSCALE_ORIG             ConntrackAttr = C.ATTR_TCP_WSCALE_ORIG
	ATTR_TCP_WSCALE_REPL             ConntrackAttr = C.ATTR_TCP_WSCALE_REPL
	ATTR_ZONE                        ConntrackAttr = C.ATTR_ZONE
	ATTR_SECCTX                      ConntrackAttr = C.ATTR_SECCTX
	ATTR_TIMESTAMP_START             ConntrackAttr = C.ATTR_TIMESTAMP_START
	ATTR_TIMESTAMP_STOP              ConntrackAttr = C.ATTR_TIMESTAMP_STOP
	ATTR_HELPER_INFO                 ConntrackAttr = C.ATTR_HELPER_INFO
	ATTR_CONNLABELS                  ConntrackAttr = C.ATTR_CONNLABELS
	ATTR_CONNLABELS_MASK             ConntrackAttr = C.ATTR_CONNLABELS_MASK
	ATTR_MAX                         ConntrackAttr = C.ATTR_MAX
)

// conntrack attribute groups
type ConntrackAttrGrp C.enum_nf_conntrack_attr_grp

const (
	ATTR_GRP_ORIG_IPV4     ConntrackAttrGrp = C.ATTR_GRP_ORIG_IPV4
	ATTR_GRP_REPL_IPV4     ConntrackAttrGrp = C.ATTR_GRP_REPL_IPV4
	ATTR_GRP_ORIG_IPV6     ConntrackAttrGrp = C.ATTR_GRP_ORIG_IPV6
	ATTR_GRP_REPL_IPV6     ConntrackAttrGrp = C.ATTR_GRP_REPL_IPV6
	ATTR_GRP_ORIG_PORT     ConntrackAttrGrp = C.ATTR_GRP_ORIG_PORT
	ATTR_GRP_REPL_PORT     ConntrackAttrGrp = C.ATTR_GRP_REPL_PORT
	ATTR_GRP_ICMP          ConntrackAttrGrp = C.ATTR_GRP_ICMP
	ATTR_GRP_MASTER_IPV4   ConntrackAttrGrp = C.ATTR_GRP_MASTER_IPV4
	ATTR_GRP_MASTER_IPV6   ConntrackAttrGrp = C.ATTR_GRP_MASTER_IPV6
	ATTR_GRP_MASTER_PORT   ConntrackAttrGrp = C.ATTR_GRP_MASTER_PORT
	ATTR_GRP_ORIG_COUNTERS ConntrackAttrGrp = C.ATTR_GRP_ORIG_COUNTERS
	ATTR_GRP_REPL_COUNTERS ConntrackAttrGrp = C.ATTR_GRP_REPL_COUNTERS
	ATTR_GRP_ORIG_ADDR_SRC ConntrackAttrGrp = C.ATTR_GRP_ORIG_ADDR_SRC
	ATTR_GRP_ORIG_ADDR_DST ConntrackAttrGrp = C.ATTR_GRP_ORIG_ADDR_DST
	ATTR_GRP_REPL_ADDR_SRC ConntrackAttrGrp = C.ATTR_GRP_REPL_ADDR_SRC
	ATTR_GRP_REPL_ADDR_DST ConntrackAttrGrp = C.ATTR_GRP_REPL_ADDR_DST
	ATTR_GRP_MAX           ConntrackAttrGrp = C.ATTR_GRP_MAX
)

// message type
type ConntrackMsgType C.enum_nf_conntrack_msg_type

const (
	NFCT_T_UNKNOWN     ConntrackMsgType = C.NFCT_T_UNKNOWN
	NFCT_T_NEW_BIT     ConntrackMsgType = C.NFCT_T_NEW_BIT
	NFCT_T_NEW         ConntrackMsgType = C.NFCT_T_NEW
	NFCT_T_UPDATE_BIT  ConntrackMsgType = C.NFCT_T_UPDATE_BIT
	NFCT_T_UPDATE      ConntrackMsgType = C.NFCT_T_UPDATE
	NFCT_T_DESTROY_BIT ConntrackMsgType = C.NFCT_T_DESTROY_BIT
	NFCT_T_DESTROY     ConntrackMsgType = C.NFCT_T_DESTROY
	NFCT_T_ALL         ConntrackMsgType = C.NFCT_T_ALL
	NFCT_T_ERROR_BIT   ConntrackMsgType = C.NFCT_T_ERROR_BIT
	NFCT_T_ERROR       ConntrackMsgType = C.NFCT_T_ERROR
)

// set option
const (
	NFCT_SOPT_UNDO_SNAT      = C.NFCT_SOPT_UNDO_SNAT
	NFCT_SOPT_UNDO_DNAT      = C.NFCT_SOPT_UNDO_DNAT
	NFCT_SOPT_UNDO_SPAT      = C.NFCT_SOPT_UNDO_SPAT
	NFCT_SOPT_UNDO_DPAT      = C.NFCT_SOPT_UNDO_DPAT
	NFCT_SOPT_SETUP_ORIGINAL = C.NFCT_SOPT_SETUP_ORIGINAL
	NFCT_SOPT_SETUP_REPLY    = C.NFCT_SOPT_SETUP_REPLY
	__NFCT_SOPT_MAX          = C.__NFCT_SOPT_MAX
	NFCT_SOPT_MAX            = C.NFCT_SOPT_MAX
)

// get option
const (
	NFCT_GOPT_IS_SNAT = C.NFCT_GOPT_IS_SNAT
	NFCT_GOPT_IS_DNAT = C.NFCT_GOPT_IS_DNAT
	NFCT_GOPT_IS_SPAT = C.NFCT_GOPT_IS_SPAT
	NFCT_GOPT_IS_DPAT = C.NFCT_GOPT_IS_DPAT
	__NFCT_GOPT_MAX   = C.__NFCT_GOPT_MAX
	NFCT_GOPT_MAX     = C.NFCT_GOPT_MAX
)

// print - output type
const (
	NFCT_O_PLAIN   = C.NFCT_O_PLAIN
	NFCT_O_DEFAULT = C.NFCT_O_DEFAULT
	NFCT_O_XML     = C.NFCT_O_XML
	NFCT_O_MAX     = C.NFCT_O_MAX
)

// print - output flags
const (
	NFCT_OF_SHOW_LAYER3_BIT = C.NFCT_OF_SHOW_LAYER3_BIT
	NFCT_OF_SHOW_LAYER3     = C.NFCT_OF_SHOW_LAYER3
	NFCT_OF_TIME_BIT        = C.NFCT_OF_TIME_BIT
	NFCT_OF_TIME            = C.NFCT_OF_TIME
	NFCT_OF_ID_BIT          = C.NFCT_OF_ID_BIT
	NFCT_OF_ID              = C.NFCT_OF_ID
	NFCT_OF_TIMESTAMP_BIT   = C.NFCT_OF_TIMESTAMP_BIT
	NFCT_OF_TIMESTAMP       = C.NFCT_OF_TIMESTAMP
)

// comparison
const (
	NFCT_CMP_ALL        = C.NFCT_CMP_ALL
	NFCT_CMP_ORIG       = C.NFCT_CMP_ORIG
	NFCT_CMP_REPL       = C.NFCT_CMP_REPL
	NFCT_CMP_TIMEOUT_EQ = C.NFCT_CMP_TIMEOUT_EQ
	NFCT_CMP_TIMEOUT_GT = C.NFCT_CMP_TIMEOUT_GT
	NFCT_CMP_TIMEOUT_GE = C.NFCT_CMP_TIMEOUT_GE
	NFCT_CMP_TIMEOUT_LT = C.NFCT_CMP_TIMEOUT_LT
	NFCT_CMP_TIMEOUT_LE = C.NFCT_CMP_TIMEOUT_LE
	NFCT_CMP_MASK       = C.NFCT_CMP_MASK
	NFCT_CMP_STRICT     = C.NFCT_CMP_STRICT
)

// copy
const (
	NFCT_CP_ALL      = C.NFCT_CP_ALL
	NFCT_CP_ORIG     = C.NFCT_CP_ORIG
	NFCT_CP_REPL     = C.NFCT_CP_REPL
	NFCT_CP_META     = C.NFCT_CP_META
	NFCT_CP_OVERRIDE = C.NFCT_CP_OVERRIDE
)

// event filtering
type FilterAttr C.enum_nfct_filter_attr

const (
	NFCT_FILTER_L4PROTO       FilterAttr = C.NFCT_FILTER_L4PROTO
	NFCT_FILTER_L4PROTO_STATE FilterAttr = C.NFCT_FILTER_L4PROTO_STATE
	NFCT_FILTER_SRC_IPV4      FilterAttr = C.NFCT_FILTER_SRC_IPV4
	NFCT_FILTER_DST_IPV4      FilterAttr = C.NFCT_FILTER_DST_IPV4
	NFCT_FILTER_SRC_IPV6      FilterAttr = C.NFCT_FILTER_SRC_IPV6
	NFCT_FILTER_DST_IPV6      FilterAttr = C.NFCT_FILTER_DST_IPV6
	NFCT_FILTER_MARK          FilterAttr = C.NFCT_FILTER_MARK
	NFCT_FILTER_MAX           FilterAttr = C.NFCT_FILTER_MAX
)

type FilterLogic C.enum_nfct_filter_logic

const (
	NFCT_FILTER_LOGIC_POSITIVE FilterLogic = C.NFCT_FILTER_LOGIC_POSITIVE
	NFCT_FILTER_LOGIC_NEGATIVE FilterLogic = C.NFCT_FILTER_LOGIC_NEGATIVE
	NFCT_FILTER_LOGIC_MAX      FilterLogic = C.NFCT_FILTER_LOGIC_MAX
)

// expect attributes
type ExpectAttr C.enum_nf_expect_attr

const (
	ATTR_EXP_MASTER      ExpectAttr = C.ATTR_EXP_MASTER
	ATTR_EXP_EXPECTED    ExpectAttr = C.ATTR_EXP_EXPECTED
	ATTR_EXP_MASK        ExpectAttr = C.ATTR_EXP_MASK
	ATTR_EXP_TIMEOUT     ExpectAttr = C.ATTR_EXP_TIMEOUT
	ATTR_EXP_ZONE        ExpectAttr = C.ATTR_EXP_ZONE
	ATTR_EXP_FLAGS       ExpectAttr = C.ATTR_EXP_FLAGS
	ATTR_EXP_HELPER_NAME ExpectAttr = C.ATTR_EXP_HELPER_NAME
	ATTR_EXP_CLASS       ExpectAttr = C.ATTR_EXP_CLASS
	ATTR_EXP_NAT_TUPLE   ExpectAttr = C.ATTR_EXP_NAT_TUPLE
	ATTR_EXP_NAT_DIR     ExpectAttr = C.ATTR_EXP_NAT_DIR
	ATTR_EXP_FN          ExpectAttr = C.ATTR_EXP_FN
	ATTR_EXP_MAX         ExpectAttr = C.ATTR_EXP_MAX
)

// Bitset representing status of connection. Taken from linux/netfilter/nf_conntrack_common.h
//
// Note: For backward compatibility this shouldn't ever change
// 	 in kernel space.
type IpConntrackStatus C.enum_ip_conntrack_status

const (
	IPS_EXPECTED_BIT      IpConntrackStatus = C.IPS_EXPECTED_BIT
	IPS_EXPECTED          IpConntrackStatus = C.IPS_EXPECTED
	IPS_SEEN_REPLY_BIT    IpConntrackStatus = C.IPS_SEEN_REPLY_BIT
	IPS_SEEN_REPLY        IpConntrackStatus = C.IPS_SEEN_REPLY
	IPS_ASSURED_BIT       IpConntrackStatus = C.IPS_ASSURED_BIT
	IPS_ASSURED           IpConntrackStatus = C.IPS_ASSURED
	IPS_CONFIRMED_BIT     IpConntrackStatus = C.IPS_CONFIRMED_BIT
	IPS_CONFIRMED         IpConntrackStatus = C.IPS_CONFIRMED
	IPS_SRC_NAT_BIT       IpConntrackStatus = C.IPS_SRC_NAT_BIT
	IPS_SRC_NAT           IpConntrackStatus = C.IPS_SRC_NAT
	IPS_DST_NAT_BIT       IpConntrackStatus = C.IPS_DST_NAT_BIT
	IPS_DST_NAT           IpConntrackStatus = C.IPS_DST_NAT
	IPS_NAT_MASK          IpConntrackStatus = C.IPS_NAT_MASK
	IPS_SEQ_ADJUST_BIT    IpConntrackStatus = C.IPS_SEQ_ADJUST_BIT
	IPS_SEQ_ADJUST        IpConntrackStatus = C.IPS_SEQ_ADJUST
	IPS_SRC_NAT_DONE_BIT  IpConntrackStatus = C.IPS_SRC_NAT_DONE_BIT
	IPS_SRC_NAT_DONE      IpConntrackStatus = C.IPS_SRC_NAT_DONE
	IPS_DST_NAT_DONE_BIT  IpConntrackStatus = C.IPS_DST_NAT_DONE_BIT
	IPS_DST_NAT_DONE      IpConntrackStatus = C.IPS_DST_NAT_DONE
	IPS_NAT_DONE_MASK     IpConntrackStatus = C.IPS_NAT_DONE_MASK
	IPS_DYING_BIT         IpConntrackStatus = C.IPS_DYING_BIT
	IPS_DYING             IpConntrackStatus = C.IPS_DYING
	IPS_FIXED_TIMEOUT_BIT IpConntrackStatus = C.IPS_FIXED_TIMEOUT_BIT
	IPS_FIXED_TIMEOUT     IpConntrackStatus = C.IPS_FIXED_TIMEOUT
	IPS_TEMPLATE_BIT      IpConntrackStatus = C.IPS_TEMPLATE_BIT
	IPS_TEMPLATE          IpConntrackStatus = C.IPS_TEMPLATE
	IPS_UNTRACKED_BIT     IpConntrackStatus = C.IPS_UNTRACKED_BIT
	IPS_UNTRACKED         IpConntrackStatus = C.IPS_UNTRACKED
)

// expectation flags
// defined in linux/netfilter/nf_conntrack_common.h too
const (
	NF_CT_EXPECT_PERMANENT = C.NF_CT_EXPECT_PERMANENT
	NF_CT_EXPECT_INACTIVE  = C.NF_CT_EXPECT_INACTIVE
	NF_CT_EXPECT_USERSPACE = C.NF_CT_EXPECT_USERSPACE
)

// TCP flags
// defined in linux/netfilter/nf_conntrack_tcp.h
const (
	IP_CT_TCP_FLAG_WINDOW_SCALE = C.IP_CT_TCP_FLAG_WINDOW_SCALE
	IP_CT_TCP_FLAG_SACK_PERM    = C.IP_CT_TCP_FLAG_SACK_PERM
	IP_CT_TCP_FLAG_CLOSE_INIT   = C.IP_CT_TCP_FLAG_CLOSE_INIT
	IP_CT_TCP_FLAG_BE_LIBERAL   = C.IP_CT_TCP_FLAG_BE_LIBERAL
)

const NFCT_HELPER_NAME_MAX = C.NFCT_HELPER_NAME_MAX
