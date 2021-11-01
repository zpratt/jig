package adapters

import (
	"log"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/zpratt/jig/internal"
	testingexec "k8s.io/utils/exec/testing"
)

func TestWindows_InstallPackage(t *testing.T) {
	installedPackageManager := "choco"
	packageToInstall := gofakeit.LetterN(10)
	shellOutput := "0"

	var commands []*testingexec.FakeCmd
	commands = append(commands,
		internal.MakeFakeCommand(shellOutput),
		internal.MakeFakeCommand(shellOutput),
		internal.MakeFakeCommand(shellOutput),
	)
	fakeExec := internal.MakeFakeExec(installedPackageManager, commands)

	windowsAdapter := NewWindowsAdapter(&fakeExec)

	windowsAdapter.InstallPackage(packageToInstall)

	chocoCommandString := commands[2].CombinedOutputLog[0][0]
	if chocoCommandString != installedPackageManager {
		log.Fatalf("choco not called")
	}

	chocoAction := commands[2].CombinedOutputLog[0][1]
	if chocoAction != "install" {
		log.Fatalf("choco install not called")
	}

	if commands[2].CombinedOutputCalls != 1 {
		log.Fatalf("choco install not called")
	}
}
