package php

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
)

// Chgrp - Changes file group
func Chgrp(name string, uid, gid int) error {
	return Chown(name, uid, gid)
}

// Chmod - Changes file mode
func Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

// Chown - Chown changes the numeric uid and gid of the named file.
func Chown(name string, uid int, gid int) error {
	return os.Chown(name, uid, gid)
}

func FileInfo(filePath string) (*os.FileInfo, error) {
	if info, err := os.Stat(filePath); os.IsNotExist(err) {
		return &info, nil
	} else {
		return nil, err
	}
}

func CopyDir(src string, dest string) (bool, error) {
	FormatPath := func(s string) string {
		switch runtime.GOOS {
		case "windows":
			return strings.Replace(s, "/", "\\", -1)
		case "darwin", "linux":
			return strings.Replace(s, "\\", "/", -1)
		default:
			panic("only support linux,windows,darwin, but os is " + runtime.GOOS)
			return s
		}
	}
	src = FormatPath(src)
	dest = FormatPath(dest)
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("xcopy", src, dest, "/I", "/E")
	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
	default:
		panic("not support os")
	}

	outPut, e := cmd.Output()
	if e != nil {
		return false, e
	}
	fmt.Println(outPut)
	return true, nil
}

// CopySymLink - copy a link
func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}

// Copy - Copies file
func Copy(dstName string, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// Delete - Deletes a file
func Delete(name string) error {
	return Unlink(name)
}

// Dirname - Returns a parent directory's path
func Dirname(dirPth string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirPth)
}

// Fclose - Closes an open file pointer
func Fclose(file *os.File) error {
	return file.Close()
}

// FileExists - Checks whether a file or directory exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Filemtime - Gets file modification time
func Filemtime(file string) time.Time {
	fi, err := os.Stat(file)
	if err != nil {
		return time.Time{}
	}
	return fi.ModTime()
}

// Glob - Find pathnames matching a pattern
func Glob(pattern string) (matches []string, err error) {
	return filepath.Glob(pattern)
}

// IsDir - Tells whether the filename is a directory
func IsDir(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.IsDir()
}

// IsReadable - Tells whether a file exists and is readable
func IsReadable(name string) bool {
	_, err := syscall.Open(name, syscall.O_RDONLY, 0)
	return err == nil
}

// IsWritable - Tells whether the filename is writable
func IsWritable(name string) bool {
	_, err := syscall.Open(name, syscall.O_WRONLY, 0)
	return err == nil
}

// IsWriteable - Alias of IsWritable()
func IsWriteable(name string) bool {
	return IsWritable(name)
}

// Mkdir - Makes directory
//func Mkdir(name string, mode os.FileMode) error {
//	return os.Mkdir(name, mode)
//}

func Mkdir(dir string) error {
	if r := FileExists(dir); !r {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			//log.Println("mkdir err:==", dir)
			return errors.New("upload fail,mkdir:==")
		}
	} else {
		//log.Println("path exist:==", dir)
	}
	return nil
}

// Realpath - Returns canonicalized absolute pathname
func Realpath(path string) (string, error) {
	return filepath.Abs(path)
}

// Rename - Renames a file or directory
func Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

// Rmdir — Removes directory
func Rmdir(path string) error {
	return os.RemoveAll(path)
}

// Stat - Gives information about a file
func Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// Unlink - Deletes a file
func Unlink(name string) error {
	return os.Remove(name)
}
