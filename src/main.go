package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// runCmd выполняет команду с указанными аргументами в заданной рабочей директории
func runCmd(name string, args []string, workDir string) error {
	cmd := exec.Command(name, args...)
	if workDir != "" {
		cmd.Dir = workDir
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Выполняется: %s %v (dir: %s)\n", name, args, workDir)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ошибка при выполнении %s %v: %w", name, args, err)
	}
	return nil
}

// doAll последовательно запускает все необходимые команды
func doAll(useFvm bool) error {
	flutterArgs := func(args ...string) (string, []string) {
		if useFvm {
			return "fvm", append([]string{"flutter"}, args...)
		}
		return "flutter", args
	}

	// Шаг 1: flutter clean
	cmd, args := flutterArgs("clean")
	if err := runCmd(cmd, args, ""); err != nil {
		return fmt.Errorf("ошибка при запуске flutter clean: %w", err)
	}

	// Шаг 2: flutter pub get
	cmd, args = flutterArgs("pub", "get")
	if err := runCmd(cmd, args, ""); err != nil {
		return fmt.Errorf("ошибка при запуске flutter pub get: %w", err)
	}

	iosDir := "ios"

	// Шаг 3.1: Удаляем Podfile.lock
	if err := runCmd("rm", []string{"-rf", "Podfile.lock"}, iosDir); err != nil {
		return fmt.Errorf("ошибка при удалении Podfile.lock: %w", err)
	}

	// Шаг 3.2: Удаляем Pods
	if err := runCmd("rm", []string{"-rf", "Pods"}, iosDir); err != nil {
		return fmt.Errorf("ошибка при удалении Pods: %w", err)
	}

	// Шаг 4: pod install --repo-update
	if err := runCmd("pod", []string{"install", "--repo-update"}, iosDir); err != nil {
		return fmt.Errorf("ошибка при запуске pod install --repo-update: %w", err)
	}

	return nil
}

func main() {
	useFvm := flag.Bool("fvm", false, "Use FVM to run flutter commands (fvm flutter ...)")
	flag.Parse()

	if err := doAll(*useFvm); err != nil {
		// Используем log.Fatal, чтобы вывести ошибку и завершить программу
		log.Fatal(err)
	}
}
