package utils

import (
	"math/bits"
	"strconv"
	"strings"
)

// Calculate prefix (i.e /24) from a netmask given in dotted decimal format (i.e 255.255.255.0)
func CalcPrefix(netmask string) (int8, error) {
	prefArr := strings.Split(netmask, ".")
	var prefix int8 = 0
	for _, val := range prefArr {
		temp, err := strconv.ParseUint(val, 10, 8)
		if err != nil {
			return -1, err
		}
		prefix += int8(bits.OnesCount(uint(temp)))
		if temp != 255 {
			break
		}
	}
	return prefix, nil
}
func CalcIpToNumeric(ip string) (uint32, error) {
	prefArr := strings.Split(ip, ".")
	var numeric uint32 = 0
	for i, val := range prefArr {
		temp, err := strconv.ParseUint(val, 10, 8)
		if err != nil {
			return 0, err
		}
		numeric += uint32(temp) << uint32(24-i*8)

	}
	return numeric, nil
}

func CalcIpv6ToNumeric(ipv6 string) ([]uint64, error) {
	addrSplt := strings.Split(ipv6, ":")
	var num1 uint64 = 0
	var num2 uint64 = 0
	for i := 0; 4 > i; i++ {
		num, err := strconv.ParseUint(addrSplt[i], 10, 16)
		if err != nil {
			return nil, err
		}
		num_2, err := strconv.ParseUint(addrSplt[i+4], 10, 16)
		if err != nil {
			return nil, err
		}
		num1 += num
		num1 = num1 << 16
		num2 += num_2
		num2 = num2 << 16
	}
	return []uint64{num1, num2}, nil
}
