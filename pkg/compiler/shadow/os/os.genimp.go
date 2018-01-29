package shadow_os

import "os"

var Pkg = make(map[string]interface{})

func init() {
	Pkg["Args"] = os.Args
	Pkg["Chdir"] = os.Chdir
	Pkg["Chmod"] = os.Chmod
	Pkg["Chown"] = os.Chown
	Pkg["Chtimes"] = os.Chtimes
	Pkg["Clearenv"] = os.Clearenv
	Pkg["Create"] = os.Create
	Pkg["DevNull"] = os.DevNull
	Pkg["Environ"] = os.Environ
	Pkg["ErrClosed"] = os.ErrClosed
	Pkg["ErrExist"] = os.ErrExist
	Pkg["ErrInvalid"] = os.ErrInvalid
	Pkg["ErrNotExist"] = os.ErrNotExist
	Pkg["ErrPermission"] = os.ErrPermission
	Pkg["Executable"] = os.Executable
	Pkg["Exit"] = os.Exit
	Pkg["Expand"] = os.Expand
	Pkg["ExpandEnv"] = os.ExpandEnv
	//Pkg["File"] = os.File
	//Pkg["FileInfo"] = os.FileInfo
	//Pkg["FileMode"] = os.FileMode
	Pkg["FindProcess"] = os.FindProcess
	Pkg["Getegid"] = os.Getegid
	Pkg["Getenv"] = os.Getenv
	Pkg["Geteuid"] = os.Geteuid
	Pkg["Getgid"] = os.Getgid
	Pkg["Getgroups"] = os.Getgroups
	Pkg["Getpagesize"] = os.Getpagesize
	Pkg["Getpid"] = os.Getpid
	Pkg["Getppid"] = os.Getppid
	Pkg["Getuid"] = os.Getuid
	Pkg["Getwd"] = os.Getwd
	Pkg["Hostname"] = os.Hostname
	Pkg["Interrupt"] = os.Interrupt
	Pkg["IsExist"] = os.IsExist
	Pkg["IsNotExist"] = os.IsNotExist
	Pkg["IsPathSeparator"] = os.IsPathSeparator
	Pkg["IsPermission"] = os.IsPermission
	Pkg["Kill"] = os.Kill
	Pkg["Lchown"] = os.Lchown
	Pkg["Link"] = os.Link
	//Pkg["LinkError"] = os.LinkError
	Pkg["LookupEnv"] = os.LookupEnv
	Pkg["Lstat"] = os.Lstat
	Pkg["Mkdir"] = os.Mkdir
	Pkg["MkdirAll"] = os.MkdirAll
	Pkg["ModeAppend"] = os.ModeAppend
	Pkg["ModeCharDevice"] = os.ModeCharDevice
	Pkg["ModeDevice"] = os.ModeDevice
	Pkg["ModeDir"] = os.ModeDir
	Pkg["ModeExclusive"] = os.ModeExclusive
	Pkg["ModeNamedPipe"] = os.ModeNamedPipe
	Pkg["ModePerm"] = os.ModePerm
	Pkg["ModeSetgid"] = os.ModeSetgid
	Pkg["ModeSetuid"] = os.ModeSetuid
	Pkg["ModeSocket"] = os.ModeSocket
	Pkg["ModeSticky"] = os.ModeSticky
	Pkg["ModeSymlink"] = os.ModeSymlink
	Pkg["ModeTemporary"] = os.ModeTemporary
	Pkg["ModeType"] = os.ModeType
	Pkg["NewFile"] = os.NewFile
	Pkg["NewSyscallError"] = os.NewSyscallError
	Pkg["O_APPEND"] = os.O_APPEND
	Pkg["O_CREATE"] = os.O_CREATE
	Pkg["O_EXCL"] = os.O_EXCL
	Pkg["O_RDONLY"] = os.O_RDONLY
	Pkg["O_RDWR"] = os.O_RDWR
	Pkg["O_SYNC"] = os.O_SYNC
	Pkg["O_TRUNC"] = os.O_TRUNC
	Pkg["O_WRONLY"] = os.O_WRONLY
	Pkg["Open"] = os.Open
	Pkg["OpenFile"] = os.OpenFile
	//Pkg["PathError"] = os.PathError
	Pkg["PathListSeparator"] = os.PathListSeparator
	Pkg["PathSeparator"] = os.PathSeparator
	Pkg["Pipe"] = os.Pipe
	//Pkg["ProcAttr"] = os.ProcAttr
	//Pkg["Process"] = os.Process
	//Pkg["ProcessState"] = os.ProcessState
	Pkg["Readlink"] = os.Readlink
	Pkg["Remove"] = os.Remove
	Pkg["RemoveAll"] = os.RemoveAll
	Pkg["Rename"] = os.Rename
	Pkg["SEEK_CUR"] = os.SEEK_CUR
	Pkg["SEEK_END"] = os.SEEK_END
	Pkg["SEEK_SET"] = os.SEEK_SET
	Pkg["SameFile"] = os.SameFile
	Pkg["Setenv"] = os.Setenv
	//Pkg["Signal"] = os.Signal
	Pkg["StartProcess"] = os.StartProcess
	Pkg["Stat"] = os.Stat
	Pkg["Stderr"] = os.Stderr
	Pkg["Stdin"] = os.Stdin
	Pkg["Stdout"] = os.Stdout
	Pkg["Symlink"] = os.Symlink
	//Pkg["SyscallError"] = os.SyscallError
	Pkg["TempDir"] = os.TempDir
	Pkg["Truncate"] = os.Truncate
	Pkg["Unsetenv"] = os.Unsetenv

}
