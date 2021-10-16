package adapters

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/zpratt/jig/internal"
	"log"
	"testing"
)

func TestWindows_InstallPackage(t *testing.T) {
	installedPackageManager := "choco"
	packageToInstall := gofakeit.LetterN(10)
	shellOutput := "0"

	command := internal.MakeFakeCommand(shellOutput)
	fakeExec := internal.MakeFakeExec(installedPackageManager, &command)

	windowsAdapter := NewWindowsAdapter(&fakeExec)

	windowsAdapter.InstallPackage(packageToInstall)

	chocoCommandString := command.CombinedOutputLog[0][0]
	if chocoCommandString != installedPackageManager {
		log.Fatalf("choco not called")
	}

	chocoAction := command.CombinedOutputLog[0][1]
	if chocoAction != "install" {
		log.Fatalf("choco install not called")
	}

	if command.CombinedOutputCalls != 1 {
		log.Fatalf("choco install not called")
	}
}
