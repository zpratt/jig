package adapters

import (
	"log"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/zpratt/jig/internal"
	testingexec "k8s.io/utils/exec/testing"
)

func TestDarwin_InstallPackage(t *testing.T) {
	installedPackageManager := "brew"
	packageToInstall := gofakeit.LetterN(10)
	shellOutput := "someoutput"

	var commands []*testingexec.FakeCmd
	commands = append(commands,
		internal.MakeFakeCommand(shellOutput),
		internal.MakeFakeCommand(shellOutput),
	)
	fakeExec := internal.MakeFakeExec(installedPackageManager, commands)

	darwinAdapter := NewDarwinAdapter(&fakeExec)

	darwinAdapter.InstallPackage(packageToInstall)

	actualBrewOutput := commands[1].CombinedOutputLog[0][0]
	if actualBrewOutput != installedPackageManager {
		log.Fatalf("%s not called", installedPackageManager)
	}

	brewAction := commands[1].CombinedOutputLog[0][1]
	if brewAction != "install" {
		log.Fatalf("%s install not called", installedPackageManager)
	}

	if commands[1].CombinedOutputCalls != 1 {
		log.Fatalf("%s install not called", installedPackageManager)
	}

}

func TestDarwin_UpdatePackageList(t *testing.T) {
}
