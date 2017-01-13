package common

import (
	"os/exec"
	"time"
	"bytes"



)

type Invoke struct {}

func (i Invoke) Command(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	return
}

func CombinedOutputTimeout(c *exec.Cmd, timeout time.Duration) ([]byte, error) {
	var b bytes.Buffer
	c.Stdout = &b
	c.Stderr = &b
	if err := c.Start(); err != nil {
		return nil, err
	}
}

// 给定命令等待超时，假定命令已经开始。如果命令超时，尝试杀掉该进程。
func WaitTimeout(c *exec.Cmd, timeout time.Duration) error {
	timer := time.NewTimer(timeout)
	done := make(chan error)
}

