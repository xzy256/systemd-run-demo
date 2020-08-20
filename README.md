## Introduce
This project aim to startup third part binary server in your program and kill it another gracefully. If your entry program(or main) not exit, kill child thread leads `zombie process `

## build
make


## run
./server
.test

## Compare
### exec.Command
this example shows `exec.Command` to exec a binary, but when your main thread not exit, you kill this thread, it will be `zombie` thread, until your main thread exit.
```
func mount() int {
	cmd := exec.Command("/search/odin/commands/binary/hdfs-mount/hdfs-mount.latest", "-confPath=/search/odin/hadoopconf/conf.polaris",
		"-logLevel=2", "rsync.master003.polaris.hadoop.js.ted:8020", "/tmp/mountp")

	fmt.Println(cmd)
	f, err := os.OpenFile("/tmp/1.log", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		fmt.Printf("fail: %v\n", err)
		return -1
	}
	cmd.Stdout = f
	cmd.Stderr = f
	user, err := user.Lookup("root")
	if err == nil {
		fmt.Printf("uid=%s,gid=%s \n", user.Uid, user.Gid)

		uid, _ := strconv.Atoi(user.Uid)
		gid, _ := strconv.Atoi(user.Gid)

		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	}

	err = cmd.Start()
	if err != nil {
		fmt.Printf("fail: %v, %v\n", err)
		return -1
	}

	fmt.Println("start hdfs-mount success")
	return cmd.Process.Pid
}
```
### systemd-run
this way exec binary thread that is managed systemd, you can kill it gracefully, not be `zombie` thread
