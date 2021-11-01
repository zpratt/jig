package internal

import (
	"k8s.io/utils/exec"
	testingexec "k8s.io/utils/exec/testing"
)

func MakeFakeCommand(output string) *testingexec.FakeCmd {
	return &testingexec.FakeCmd{
		CombinedOutputScript: []testingexec.FakeAction{
			func() ([]byte, []byte, error) {
				return []byte(output), nil, nil
			},
		},
	}
}

func MakeFakeExec(installedPackageManager string, commands []*testingexec.FakeCmd) testingexec.FakeExec {
	var fakeCommands []testingexec.FakeCommandAction

	for _, fakeCommand := range commands {
		fakeCommands = append(fakeCommands,
			func(cmd string, args ...string) exec.Cmd {
				return testingexec.InitFakeCmd(fakeCommand, cmd, args...)
			})
	}

	fakeExec := testingexec.FakeExec{
		LookPathFunc: func(command string) (string, error) {
			return installedPackageManager, nil
		},
		CommandScript: fakeCommands,
	}
	return fakeExec
}
