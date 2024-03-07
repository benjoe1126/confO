package cisco

import (
	"fmt"
)

func enable() (string, bool) {
	if State == DEFAULT {
		State = PRIVILEGED
		return "enable", true
	}
	return "", false
}

func confT() (string, bool) {
	if State == PRIVILEGED {
		State = CONF_T
		return "configure terminal", true
	}
	return "", false
}

func inter(str string) (string, bool) {
	if State != PRIVILEGED && State != DEFAULT {
		State = CONF_INT
		currentInterface = str
		return fmt.Sprintf("interface %s", str), true
	}
	return "", false
}
