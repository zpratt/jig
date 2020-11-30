package adapters

import (
	utilsexec "k8s.io/utils/exec"
	testingexec "k8s.io/utils/exec/testing"
	"log"
	"testing"
)

func TestWindows_InstallPackage(t *testing.T) {
	installedPackageManager := "choco"

	chocoCommand := makeChocoCommand()

	fakeExec := testingexec.FakeExec{
		LookPathFunc: func(command string) (string, error) {
			return installedPackageManager, nil
		},
		CommandScript: []testingexec.FakeCommandAction{
			func(cmd string, args ...string) utilsexec.Cmd {
				return testingexec.InitFakeCmd(&chocoCommand, cmd, args...)
			},
		},
	}

	windowsAdapter := NewWindowsAdapter(&fakeExec)

	windowsAdapter.InstallPackage(installedPackageManager)

	chocoCommandString := chocoCommand.CombinedOutputLog[0][0]
	if chocoCommandString != installedPackageManager {
		log.Fatalf("choco not called")
	}

	chocoAction := chocoCommand.CombinedOutputLog[0][1]
	if chocoAction != "install" {
		log.Fatalf("choco install not called")
	}

	if chocoCommand.CombinedOutputCalls != 1 {
		log.Fatalf("choco install not called")
	}
}

func makeChocoCommand() testingexec.FakeCmd {
	return testingexec.FakeCmd{
		CombinedOutputScript: []testingexec.FakeAction{
			func() ([]byte, []byte, error) {
				chocoInstallOutputString := "0"
				return []byte(chocoInstallOutputString), nil, nil
			},
		},
	}
}
