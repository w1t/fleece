package docker

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/leastauthority/fleece/cmd/fleece/config"
)

func ContainerName(pkgPath, fuzzFuncName string) string {
	return fmt.Sprintf("%s_%s", filepath.Base(pkgPath), fuzzFuncName)
}

func GetGuestWorkdir(name string) (string, error) {
	fleeceDir, err := config.GetFleeceDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(".", fleeceDir, "workdirs", name), nil
}

func RunContainer(cwd string, args []string) error {
	args = append([]string{"run"}, args...)
	dockerCmd := exec.Command("docker", args...)
	dockerCmd.Dir = cwd
	dockerCmd.Stdout = os.Stdout
	dockerCmd.Stderr = os.Stderr
	return dockerCmd.Run()
}

func StopContainer(name string) error {
	cmd := exec.Command("docker", "stop", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
