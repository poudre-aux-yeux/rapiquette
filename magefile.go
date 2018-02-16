// +build mage

package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Aliases = map[string]interface{}{
	"build-ci":  BuildCI,
	"vendor-ci": VendorCI,
}

const (
	goexe  = "go"
	binary = "rapiquette"
)

// Get everything ready in for a fresh clone
func Setup() {
	fmt.Println("setting rapiquette up...")
	mg.Deps(GetGoGenerate, GetDep, GetGoimports)
	mg.Deps(Schema)
	mg.Deps(Vendor, Imports)
	fmt.Println("formatted the code")
	mg.Deps(Build)
	fmt.Println("generated the executable")
}

// Install golang/dep
func GetDep() error {
	if err := sh.Run(goexe, "get", "-u", "github.com/golang/dep/cmd/dep"); err != nil {
		fmt.Println("error installing golang/dep: ", err)
		return err
	}

	fmt.Println("installed golang/dep")
	return nil
}

// Install go generate
func GetGoGenerate() error {
	if err := sh.Run(goexe, "get", "-u", "github.com/jteeuwen/go-bindata/..."); err != nil {
		fmt.Println("error installing go generate")
		return err
	}

	fmt.Println("installed go generate")
	return nil
}

// Sync vendored dependencies
func Vendor() error {
	if err := sh.Run("dep", "ensure"); err != nil {
		fmt.Println("error syncing the deps: ", err)
		return err
	}

	fmt.Println("synced the dependencies")
	return nil
}

// Sync vendored dependencies without merging new deps
func VendorCI() error {
	if err := sh.Run("dep", "ensure", "-vendor-only"); err != nil {
		fmt.Println("error syncing the deps: ", err)
		return err
	}

	fmt.Println("synced the dependencies")
	return nil
}

// Generate the GraphQL schema
func Schema() error {
	cmd := exec.Command(goexe, "generate")
	cmd.Dir = "./schema"
	out, err := cmd.Output()
	if out != nil && len(out) > 0 && err != nil {
		fmt.Println(string(out), err)
	}
	if err == nil {
		fmt.Println("generated the GraphQL schema")
	}
	return err
}

// Build Build the app
func Build() error {
	return sh.Run(goexe, "build")
}

// Build the app for linux am64
func BuildCI() error {
	args := make(map[string]string)
	args["GOOS"] = "linux"
	args["GOARCH"] = "amd64"
	args["CGO_ENABLED"] = "0"
	return sh.RunWith(args, goexe, "build", "-o", binary)
}

// Build binary with race detector enabled
func BuildRace() error {
	mg.Deps(Vendor)
	return sh.Run(goexe, "build", "-race", "-o", binary)
}

// Install Install the binary
func Install() error {
	mg.Deps(Vendor)
	return sh.Run(goexe, "install", binary)
}

var docker = sh.RunCmd("docker")

// Build the Docker container
func Docker() error {
	err := docker("build", "-t", binary, ".")
	return err
}

// Run tests and linters
func Check() {
	if strings.Contains(runtime.Version(), "1.8") {
		// Go 1.8 doesn't play along with go test ./... and /vendor.
		fmt.Printf("Skip Check on %s\n", runtime.Version())
		return
	}
	mg.Deps(TestRace)
}

// Run tests
func Test() error {
	return sh.Run(goexe, "test", "./...")
}

// Run tests with race detector
func TestRace() error {
	return sh.Run(goexe, "test", "-race", "./...")
}

var pkgPrefixLen = len("github.com/poudre-aux-yeux/rapiquette")

func packages() ([]string, error) {
	s, err := sh.Output(goexe, "list", "./...")
	if err != nil {
		return nil, err
	}
	pkgs := strings.Split(s, "\n")
	for i := range pkgs {
		pkgs[i] = "." + pkgs[i][pkgPrefixLen:]
	}
	return pkgs, nil
}

// Install goimports
func GetGoimports() error {
	if err := sh.Run(goexe, "get", "-u", "golang.org/x/tools/cmd/goimports"); err != nil {
		fmt.Println("error installing go generate")
		return err
	}

	fmt.Println("installed goimports")
	return nil
}

// Run goimports linter
func Imports() error {
	pkgs, err := packages()
	if err != nil {
		return err
	}
	failed := false
	first := true
	for _, pkg := range pkgs {
		files, err := filepath.Glob(filepath.Join(pkg, "*.go"))
		if err != nil {
			return nil
		}
		for _, f := range files {
			// goimports doesn't exit with non-zero when it finds unformatted code
			// so we have to explicitly look for output, and if we find any, we
			// should fail this target.
			s, err := sh.Output("goimports", "-w", f)
			if err != nil {
				fmt.Printf("ERROR: running goimports on %q: %v\n", f, err)
				failed = true
			}
			if s != "" {
				if first {
					fmt.Println("The following files are not goimports'ed:")
					first = false
				}

				failed = true
				fmt.Println(s)
			}
		}
	}
	if failed {
		return errors.New("improperly formatted go files")
	}
	return nil
}

// Run golint linter
func Lint() error {
	pkgs, err := packages()
	if err != nil {
		return err
	}
	failed := false
	for _, pkg := range pkgs {
		// We don't actually want to fail this target if we find golint errors,
		// so we don't pass -set_exit_status, but we still print out any failures.
		if _, err := sh.Exec(nil, os.Stderr, nil, "golint", pkg); err != nil {
			fmt.Printf("ERROR: running go lint on %q: %v\n", pkg, err)
			failed = true
		}
	}
	if failed {
		return errors.New("errors running golint")
	}
	return nil
}

// Run go vet linter
func Vet() error {
	if err := sh.Run(goexe, "vet", "./..."); err != nil {
		return fmt.Errorf("error running govendor: %v", err)
	}
	return nil
}

// Start the databases with docker compose
func Databases() error {
	os.Setenv("RAQUETTE_HOST", "localhost:6380")
	os.Setenv("TENNIS_HOST", "localhost:6379")
	return sh.Run("docker-compose", "up", "tennis-redis", "raquette-redis")
}
