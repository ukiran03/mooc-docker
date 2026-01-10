package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <user/repo> <docker-image>")
		return
	}

	repo := os.Args[1]
	image := os.Args[2]

	gitUrl := "https://github.com/" + repo + ".git"

	dirName := "repo-build"
	parts := strings.Split(repo, "/")
	if len(parts) > 1 {
		dirName = parts[1]
	}

	tempdir, err := os.MkdirTemp("", dirName)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempdir)

	if err = gitClone(gitUrl, tempdir); err != nil {
		fmt.Printf("Aborting: %v\n", err)
		return
	}

	if err = dockerLogin(); err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	if err = dockerBuildAndPush(tempdir, image); err != nil {
		fmt.Printf("Aborting: %v\n", err)
		return
	}
	fmt.Printf("\n%s: Pipeline completed successfully!\n", repo)
}

func dockerLogin() error {
	user := os.Getenv("DOCKER_USER")
	pass := os.Getenv("DOCKER_PWD")

	if user == "" || pass == "" {
		return fmt.Errorf("DOCKER_USER or DOCKER_PWD not set")
	}

	fmt.Printf("Logging in to Docker Hub as %s...\n", user)
	cmd := exec.Command("docker", "login", "-u", user, "--password-stdin")
	cmd.Stdin = strings.NewReader(pass)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker login failed: %w", err)
	}
	return nil
}

func gitClone(url, tempdir string) error {
	fmt.Printf("Cloning: %s into %s\n", url, tempdir)
	cmd := exec.Command("git", "clone", "--depth", "1", url, tempdir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone failed: %w", err)
	}
	return nil
}

func dockerBuildAndPush(contextDir, imgTag string) error {
	fmt.Printf("Building Docker image: [%s]...\n", imgTag)
	cmd := exec.Command("docker", "build", "-t", imgTag, ".")
	cmd.Dir = contextDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker build failed: %w", err)
	}

	fmt.Printf("Pushing Docker image: [%s]...\n", imgTag)
	pushCmd := exec.Command("docker", "push", imgTag)
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	if err := pushCmd.Run(); err != nil {
		return fmt.Errorf("docker push failed: %w", err)
	}
	return nil
}
