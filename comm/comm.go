package comm

import (
	"net"
	"syscall"
	"time"
)

func GetRlimitFile() uint64 {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		panic(err)
	}
	return rLimit.Cur
}

func ReadTimeout(c net.Conn, data []byte, timeout time.Duration) (int, error) {
	c.SetReadDeadline(time.Now().Add(timeout))
	return c.Read(data)
}

func WriteTimeout(c net.Conn, data []byte, timeout time.Duration) (int, error) {
	c.SetWriteDeadline(time.Now().Add(timeout))
	return c.Write(data)
}

// 请赋值成自己的根据addrType, addr返回ip:port的函数
var AddrFunc = func(addrType, addr string) (string, error) {
	return addr, nil
}