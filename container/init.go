package container

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	//"path/filepath"
	"strings"
	"syscall"
)

func RunContainerInitProcess() error {
	cmdArray := readUserCommand()
	if cmdArray == nil || len(cmdArray) == 0 {
		return fmt.Errorf("Run container get user command error, cmdArray is nil")
	}

	setUpMount()

	path, err := exec.LookPath(cmdArray[0])
	if err != nil {
		log.Errorf("Exec loop path error %v", err)
		return err
	}
	log.Infof("Find path %s", path)
	if err := syscall.Exec(path, cmdArray[0:], os.Environ()); err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

func readUserCommand() []string {
	pipe := os.NewFile(uintptr(3), "pipe")
	defer pipe.Close()
	msg, err := ioutil.ReadAll(pipe)
	if err != nil {
		log.Errorf("init read pipe error %v", err)
		return nil
	}
	msgStr := string(msg)
	return strings.Split(msgStr, " ")
}

/**
Init 挂载点
*/
func setUpMount() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Errorf("Get current location error %v", err)
		return
	}
	log.Infof("Current location is %s", pwd)

	err = pivotRoot(pwd)
	if err != nil {
		log.Errorf("pivotRoot error:%s", err)
	}
	//mount proc
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	err = syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	if err != nil {
		log.Errorf("pivotRoot error:%s", err)
	}
	syscall.Mount("udev", "/dev", "devtmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755")
}

func pivotRoot(root string) error {
	/**
	  为了使当前root的老 root 和新 root 不在同一个文件系统下，我们把root重新mount了一次
	  bind mount是把相同的内容换了一个挂载点的挂载方法
	*/
	if err := syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf("Mount rootfs to itself error: %v", err)
	}
	/**
	  为了使当前root的老 root 和新 root 不在同一个文件系统下，我们把root重新mount了一次
	  bind mount是把相同的内容换了一个挂载点的挂载方法
	*/
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf("Mount rootfs to itself error: %v", err)
	}
	// 创建 rootfs/.pivot_root 存储 old_root
	pivotDir := filepath.Join(root, ".pivot_root")
	if err := os.Mkdir(pivotDir, 0777); err != nil {
		return err
	}
	// pivot_root 到新的rootfs, 现在老的 old_root 是挂载在rootfs/.pivot_root
	// 挂载点现在依然可以在mount命令中看到
	if err := syscall.PivotRoot(root, pivotDir); err != nil {
		return fmt.Errorf("pivot_root %v", err)
	}
	// 修改当前的工作目录到根目录
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("chdir / %v", err)
	}

	pivotDir = filepath.Join("/", ".pivot_root")
	// umount rootfs/.pivot_root
	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount pivot_root dir %v", err)
	}
	// 删除临时文件夹
	return os.Remove(pivotDir)
}

//func pivotRoot(rootfs string) error {
//	// While the documentation may claim otherwise, pivot_root(".", ".") is
//	// actually valid. What this results in is / being the new root but
//	// /proc/self/cwd being the old root. Since we can play around with the cwd
//	// with pivot_root this allows us to pivot without creating directories in
//	// the rootfs. Shout-outs to the LXC developers for giving us this idea.
//
//	oldroot, err := unix.Open("/", unix.O_DIRECTORY|unix.O_RDONLY, 0)
//	if err != nil {
//		return err
//	}
//	defer unix.Close(oldroot)
//
//	newroot, err := unix.Open(rootfs, unix.O_DIRECTORY|unix.O_RDONLY, 0)
//	if err != nil {
//		return err
//	}
//	defer unix.Close(newroot)
//
//	// Change to the new root so that the pivot_root actually acts on it.
//	if err := unix.Fchdir(newroot); err != nil {
//		return err
//	}
//
//	if err := unix.PivotRoot(".", "."); err != nil {
//		return fmt.Errorf("pivot_root %s", err)
//	}
//
//	// Currently our "." is oldroot (according to the current kernel code).
//	// However, purely for safety, we will fchdir(oldroot) since there isn't
//	// really any guarantee from the kernel what /proc/self/cwd will be after a
//	// pivot_root(2).
//
//	if err := unix.Fchdir(oldroot); err != nil {
//		return err
//	}
//
//	// Make oldroot rslave to make sure our unmounts don't propagate to the
//	// host (and thus bork the machine). We don't use rprivate because this is
//	// known to cause issues due to races where we still have a reference to a
//	// mount while a process in the host namespace are trying to operate on
//	// something they think has no mounts (devicemapper in particular).
//	if err := unix.Mount("", ".", "", unix.MS_SLAVE|unix.MS_REC, ""); err != nil {
//		return err
//	}
//	// Preform the unmount. MNT_DETACH allows us to unmount /proc/self/cwd.
//	if err := unix.Unmount(".", unix.MNT_DETACH); err != nil {
//		return err
//	}
//
//	// Switch back to our shiny new root.
//	if err := unix.Chdir("/"); err != nil {
//		return fmt.Errorf("chdir / %s", err)
//	}
//	return nil
//}
