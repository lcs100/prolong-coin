package cpu

import (
	"fmt"
	"os/exec"
	"strings"
	"strconv"
)
var CpuUsage float64 = 0.0

func GetCpuUsage() float64{

	cmd := exec.Command("/bin/sh", "-c", "sar -u 1 1")
	res, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	data := strings.Split(string(res),"     ")
	cpuUsage,_ := strconv.ParseFloat(data[len(data)-1][:len(data[len(data)-1])-1],64)
	return cpuUsage
}