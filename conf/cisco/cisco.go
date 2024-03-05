package cisco

import (
	"fmt"
)

func enable() (string, bool) {
	if state == DEFAULT {
		state = PRIVILEGED
		return "enable", true
	}
	return "", false
}

func confT() (string, bool) {
	if state == PRIVILEGED {
		state = CONF_T
		return "configure terminal", true
	}
	return "", false
}

func inter(str string) (string, bool) {
	if state != PRIVILEGED && state != DEFAULT {
		state = CONF_INT
		currentInterface = str
		return fmt.Sprintf("interface %s", str), true
	}
	return "", false
}

type ACL struct {
}

func (acl *ACL) Configure() (string, error) {
	return "", nil
}
