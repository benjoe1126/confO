package cisco

type state int

const (
	DEFAULT state = iota
	PRIVILEGED
	CONF_T
	CONF_INT
	CONF_OSPF
	CONF_EIGRP
	CONF_LINE
	CONF_BGP
)

var (
	State            = DEFAULT
	currentInterface = ""
)
