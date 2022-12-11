package execute

import (
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func (e *Execute) Test(question string, stdin string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// prepare *exec.Cmd
	cmd := exec.CommandContext(ctx, "./"+question)
	cmd.Stderr = e.stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// create pipe
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer stdinPipe.Close()

	// error channel
	errChan := make(chan error)

	// cmd killer
	go func() {
		<-ctx.Done()
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}()

	// syscall listener
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sig
		cancel()
		errChan <- ErrAborted
	}()

	// TLE Timer
	go func() {
		time.Sleep(time.Second * 2)
		cancel()
		errChan <- ErrTimeout
	}()

	// create buffer for stdout
	buf := new(bytes.Buffer)

	// send stdin
	io.WriteString(stdinPipe, stdin)

	// execute target
	go func() {
		if err := cmd.Start(); err != nil {
			errChan <- err
		}

		// get stdout
		buf.ReadFrom(stdoutPipe)

		err := cmd.Wait()
		if err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	err = <-errChan
	if err != nil {
		return "", err
	}

	out := buf.String()
	return out, nil
}
