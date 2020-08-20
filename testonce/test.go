package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

const(
	SocketPath = "/host/etc/csi-tool/connector.sock"
)

func main() {
	mount()
	i := 0
	for {
		time.Sleep(3 * time.Second)
		if i == 5 {
			umount("/tmp/mountp") // 模拟程序中杀服务进程
		}
		i ++
	}
}

func umount(mountpoint string) error {
	umountCmd := fmt.Sprintf("umount -l %v", mountpoint)
	res, err := connectorRun(umountCmd)
	fmt.Printf("umount response:%v, result: %v \n", res, err)
	return err
}

func mount() error {
	mntCmd := fmt.Sprintf("systemd-run --scope -- /search/odin/commands/binary/hdfs-mount/hdfs-mount.latest -confPath=/search/odin/hadoopconf/conf.polaris  -logLevel=2 rsync.master003.polaris.hadoop.js.ted:8020 /tmp/mountp &>/tmp/1.log &")
	res, err := connectorRun(mntCmd)
	fmt.Printf("mount response:%v, result: %v \n", res, err)
	return err
}

func connectorRun(cmd string) (string, error) {
	c, err := net.Dial("unix", SocketPath)
	if err != nil {
		fmt.Errorf("Oss connector Dial error: %s", err.Error())
		return err.Error(), err
	}
	defer c.Close()

	_, err = c.Write([]byte(cmd))
	if err != nil {
		fmt.Errorf("Oss connector write error: %s", err.Error())
		return err.Error(), err
	}

	buf := make([]byte, 2048)
	n, err := c.Read(buf[:])
	response := string(buf[0:n])
	if strings.HasPrefix(response, "Success") {
		respstr := response[8:]
		return respstr, nil
	}
	return response, errors.New("oss connector exec command error:" + response)
}