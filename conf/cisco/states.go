package cisco

const (
	DEFAULT    = 0
	PRIVILEGED = 1
	CONF_T     = 2
	CONF_INT   = 3
	CONF_OSPF  = 4
	CONF_EIGRP = 5
	CONF_LINE  = 6
	CONF_BGP   = 7
)

var (
	state            = DEFAULT
	currentInterface = ""
)
