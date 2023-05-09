package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// 1	  2	   3	  4...
// docker run <image> /bin/bash
// docker child <cmd> <arguments>

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func run() {
	fmt.Println("RUN Pid:", os.Getgid())
	fmt.Printf("Running %v \n", os.Args[2:])

	args := append([]string{"child"}, os.Args[2:]...)

	cmd := exec.Command("/proc/self/exe", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())

}

func child() {
	fmt.Println("CHILD Pid:", os.Getgid())
	fmt.Printf("Child Running %v \n", os.Args[2:])

	cg()

	// need sudo
	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot("./ubuntu-fs"))
	must(syscall.Chdir("/"))

	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")

	os.Mkdir(filepath.Join(pids, "container"), 0755)

	must(ioutil.WriteFile(filepath.Join(pids, "container/pids.max"), []byte("20"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "container/notify_on_release"), []byte("1"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "container/cgroup.procs"), []byte(strconv.Itoa(os.Getgid())), 0700))

}

func main() {
	if len(os.Args) <= 1 {
		panic("no enough arguments")
	}

	if len(os.Args) <= 2 {
		panic("please specify the command to run")
	}

	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("unrecognized command" + os.Args[1])
	}
}
