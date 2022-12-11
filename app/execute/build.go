package execute

import "os/exec"

func (e *Execute) Build(question string) error {
	cmd := exec.Command("gcc", "-o", question, question+".c", "-Wall", "-O")
	cmd.Stdout = e.stdout
	cmd.Stderr = e.stderr
	cmd.Stdin = e.stdin
	return cmd.Run()
}

func (e *Execute) BuildWithoutWarning(question string) error {
	cmd := exec.Command("gcc", "-o", question, question+".c", "-O", "-w")
	cmd.Stdout = e.stdout
	cmd.Stderr = e.stderr
	cmd.Stdin = e.stdin
	return cmd.Run()
}
