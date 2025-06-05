package main

import (
	"fmt"
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

func main() {
	// Шаг 1: flutter clean
	if err := runCmd("flutter", []string{"clean"}, ""); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Шаг 2: flutter pub get
	if err := runCmd("flutter", []string{"pub", "get"}, ""); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Рабочая директория для iOS
	iosDir := "ios"

	// Шаг 3: Удаление Podfile.lock и папки Pods в директории ios
	if err := runCmd("rm", []string{"-rf", "Podfile.lock"}, iosDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := runCmd("rm", []string{"-rf", "Pods"}, iosDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Шаг 4: pod install --repo-update
	if err := runCmd("pod", []string{"install", "--repo-update"}, iosDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
