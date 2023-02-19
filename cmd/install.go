package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/magdyismail88/notebook/database"
	"github.com/magdyismail88/notebook/util"
)

func Install() error {
	if err := database.Run(); err != nil {
		return err
	}
	if err := buildFrontend(); err != nil {
		return err
	}
	if err := buildBackend(); err != nil {
		return err
	}
	return nil
}

func buildFrontend() error {
	fmt.Println("Node modules installing...")
	os.Chdir("frontend")
	a := exec.Command("npm", "install")
	if err := a.Run(); err != nil {
		return err
	}
	fmt.Println("Node modules installed.")
	b := exec.Command("npm", "run", "build")
	if err := b.Run(); err != nil {
		return err
	}
	fmt.Println("Built for production")
	os.Chdir("../")
	return nil
}

func buildBackend() error {
	var err error
	fmt.Println("Go modules installing...")
	// Install golang modules
	a := exec.Command("go", "mod", "tidy")
	if err := a.Run(); err != nil {
		return fmt.Errorf("build modules field: %s", err)
	}
	fmt.Println("Go modules installed.")
	os.Chdir("server")
	b := exec.Command("go", "build", ".")
	if err := b.Run(); err != nil {
		return fmt.Errorf("build app field!: %s", err)
	}
	os.Chdir("../")
	if err := util.MoveFile(path.Join("server", "server"), "bin/notebook-bin"); err != nil {
		return fmt.Errorf("Invalid move file")
	}
	if runtime.GOOS != "windows" {
		err = os.Chmod("bin/notebook-bin", 0700)
	}
	if err != nil {
		return err
	}
	return nil
}
