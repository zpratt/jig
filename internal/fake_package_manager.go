package internal

import (
	"k8s.io/utils/exec"
	"k8s.io/utils/exec/testing"
)

func MakeFakeCommand(output string) testingexec.FakeCmd {
	return testingexec.FakeCmd{
		CombinedOutputScript: []testingexec.FakeAction{
			func() ([]byte, []byte, error) {
				return []byte(output), nil, nil
			},
		},
	}
}

func MakeFakeExec(installedPackageManager string, command *testingexec.FakeCmd) testingexec.FakeExec {
	fakeExec := testingexec.FakeExec{
		LookPathFunc: func(command string) (string, error) {
			return installedPackageManager, nil
		},
		CommandScript: []testingexec.FakeCommandAction{
			func(cmd string, args ...string) exec.Cmd {
				return testingexec.InitFakeCmd(command, cmd, args...)
			},
		},
	}
	return fakeExec
}
