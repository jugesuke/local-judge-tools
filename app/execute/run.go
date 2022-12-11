package execute

import "os/exec"

func (e *Execute) Run(question string) error {
	cmd := exec.Command("./" + question)
	cmd.Stdout = e.stdout
	cmd.Stderr = e.stderr
	cmd.Stdin = e.stdin
	return cmd.Run()
}
