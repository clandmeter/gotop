// +build linux,!arm64 openbsd,!arm64 freebsd darwin,!arm64
// +build linux,!riscv64 openbsd,!riscv64 freebsd darwin,!riscv64

package logging

import "syscall"
import "os"

func stderrToLogfile(logfile *os.File) {
	syscall.Dup2(int(logfile.Fd()), 2)
}
