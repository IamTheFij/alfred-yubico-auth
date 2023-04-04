//go:build mage

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deanishe/awgo/util"
	"github.com/deanishe/awgo/util/build"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	buildDir = "./build"
	distDir  = "./dist"
	binName  = "alfred-yubico-auth"
)

var (
	info *build.Info

	// Default mage target
	Default = Run

	// Output binary path
	binPath = filepath.Join(buildDir, binName)
)

func InfoWithVersion(v int) build.Option {
	return func(i *build.Info) {
		i.AlfredMajorVersion = v
	}
}

func init() {
	var err error
	if info, err = build.NewInfo(InfoWithVersion(5)); err != nil {
		panic(err)
	}
}

func CompileBinaryForArch(goos, goarch, outputPath string) error {
	env := info.Env()
	if goos != "" {
		env["GOOS"] = goos
	}
	if goarch != "" {
		env["GOARCH"] = goarch
	}

	err := sh.RunWith(env, "go", "build", "-o", outputPath, ".")
	if err != nil {
		return fmt.Errorf("error building %s %s binary %w", goos, goarch, err)
	}

	return nil
}

func CompileUniversalBinary() error {
	armPath := filepath.Join(buildDir, "alfred-yubico-auth_arm64")
	// amdPath := filepath.Join(buildDir, "alfred-yubico-auth_amd64")

	// NOTE: Universal binaries can't be compiled because we can't cross compile against the scard
	// libraries due to missing headers
	if err := CompileBinaryForArch("darwin", "arm64", armPath); err != nil {
		return fmt.Errorf("error compiling universal binary %w", err)
	}

	/*
	 * if err := CompileBinaryForArch("darwin", "amd64", amdPath); err != nil {
	 * 	return fmt.Errorf("error compiling universal binary %w", err)
	 * }
	 */

	if err := sh.RunWith(
		info.Env(),
		"lipo",
		"-create",
		"-output",
		binPath,
		armPath,
		// amdPath,
	); err != nil {
		return fmt.Errorf("failed combining binaries to universal %w", err)
	}

	return nil
}

// Build workflow.
func Build() error {
	mg.Deps(cleanBuild)
	fmt.Println("Building...")

	if err := CompileUniversalBinary(); err != nil {
		return fmt.Errorf("error building binary %w", err)
	}

	globs := build.Globs(
		"*.png",
		"info.plist",
		"README.md",
		"LICENSE.txt",
		"password-prompt.js",
	)

	return build.SymlinkGlobs(buildDir, globs...)
}

// Run workflow.
func Run() error {
	mg.Deps(Build)
	fmt.Println("Running...")

	return sh.RunWith(info.Env(), binPath)
}

// Dist packages workflow for distribution.
func Dist() error {
	mg.SerialDeps(Clean, Build)
	fmt.Println("Exporting dist...")

	p, err := build.Export(buildDir, distDir)
	if err != nil {
		return err
	}

	fmt.Printf("Exported %q\n", p)

	return nil
}

// Install symlinked workflow to Alfred.
func Install() error {
	mg.Deps(Build)
	fmt.Printf("Installing (linking) %q to %q...\n", buildDir, info.InstallDir)

	if err := sh.Rm(info.InstallDir); err != nil {
		return fmt.Errorf("error cleaning previously installed workflow: %w", err)
	}

	return build.Symlink(info.InstallDir, buildDir, true)
}

// InstallHooks will install pre-commit hooks.
func InstallHooks() error {
	return sh.RunV("pre-commit", "install", "--overwrite", "--install-hooks")
}

// Check will run all pre-commit hooks.
func Check() error {
	return sh.RunV("pre-commit", "run", "--all-files")
}

// Clean build files.
func Clean() error {
	fmt.Println("Cleaning...")
	mg.Deps(cleanBuild, cleanMage)

	return nil
}

// DistClean build files and distribution files.
func DistClean() error {
	mg.Deps(Clean, cleanDist)

	return nil
}

func cleanDir(name string) error {
	if !util.PathExists(name) {
		return nil
	}

	infos, err := ioutil.ReadDir(name)
	if err != nil {
		return fmt.Errorf("cleanDir could not read folder: %w", err)
	}

	for _, fi := range infos {
		if err := sh.Rm(filepath.Join(name, fi.Name())); err != nil {
			return fmt.Errorf("cleanDir could not remove file: %w", err)
		}
	}

	return nil
}

func cleanBuild() error { return cleanDir(buildDir) }
func cleanDist() error  { return cleanDir(distDir) }
func cleanMage() error  { return sh.Run("mage", "-clean") }
