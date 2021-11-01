package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/zpratt/jig/adapters"
	"gopkg.in/yaml.v2"
	utilsexec "k8s.io/utils/exec"
)

type JigConfig struct {
	Tools []string `yaml:"tools"`
}

func main() {
	desiredState, err := desiredStateFactory()
	if err != nil {
		log.Fatal("Unable to load OS specific setup:\n", err.Error())
	}
	jigConfig := parseConfig()
	notInstalled := findMissingPackages(jigConfig)

	if len(notInstalled) > 0 {
		for _, notInstalledTool := range notInstalled {
			desiredState.InstallPackage(notInstalledTool)
		}
	} else {
		log.Printf("all tools installed")
	}
}

func findMissingPackages(jigConfig JigConfig) []string {
	var notInstalled []string

	for _, tool := range jigConfig.Tools {
		_, err := exec.LookPath(tool)

		if err != nil {
			notInstalled = append(notInstalled, tool)
		}
	}
	return notInstalled
}

func parseConfig() JigConfig {
	config := JigConfig{}

	jigConfigFile := "jig.yaml"
	_, err := os.Stat(jigConfigFile)

	if err != nil {
		log.Fatalf("%s config does not exist", jigConfigFile)
	}

	file, _ := ioutil.ReadFile(jigConfigFile)
	err = yaml.Unmarshal(file, &config)

	if err != nil {
		log.Fatalf("failed to parse %s", jigConfigFile)
	}

	return config
}

func desiredStateFactory() (*DesiredState, error) {
	execInst := utilsexec.New()

	factories := map[string]func() *DesiredState{
		"darwin": func() *DesiredState {
			return &DesiredState{PlatformAdapter: adapters.NewDarwinAdapter(execInst)}
		},
		"linux": func() *DesiredState {
			return &DesiredState{PlatformAdapter: adapters.NewLinuxAdapter(execInst)}
		},
		"windows": func() *DesiredState {
			return &DesiredState{PlatformAdapter: adapters.NewWindowsAdapter(execInst)}
		},
	}

	if factory, ok := factories[runtime.GOOS]; ok {
		return factory(), nil
	} else {
		return nil, errors.New(fmt.Sprint("Unsupported OS: ", runtime.GOOS))
	}
}
