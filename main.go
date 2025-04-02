package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	usingCache := flag.String("using-cache", "yes", "Should cache be used (can be `yes` or `no`).")
	configFile := flag.String("config", "", "The path to a config file.")
	cacheFile := flag.String("cache-file", "", "The path to the cache file.")
	flag.Parse()

	// generate a temporary file pivot
	tmpFile := fmt.Sprintf("php-cs-%d.php", time.Now().UnixNano())
	tmpDir := os.TempDir()

	tmpFilePath := fmt.Sprintf("%s/%s", tmpDir, tmpFile)
	tmpFileHandle, err := os.Create(tmpFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// ensure the temporary file and directory are removed on exit
	defer func() {
		os.Remove(tmpFilePath)
		os.RemoveAll(tmpDir)
	}()

	// read the input from STDIN as provided by a text editor
	file, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	tmpFileHandle.Write(file)
	tmpFileHandle.Close()

	// here we need to yield any parameters, for now we rely on the
	// php-cs-fixer CLI to handle any validations and errors
	args := []string{"--quiet", "fix", "--using-cache=" + *usingCache}
	if *configFile != "" {
		args = append(args, fmt.Sprintf("--config=%s", *configFile))
	}
	if *cacheFile != "" {
		args = append(args, fmt.Sprintf("--cache-file=%s", *cacheFile))
	}
	args = append(args, tmpFilePath)

	cmd := exec.Command("php-cs-fixer", args...)

	// capture the output and error streams, in case of errors
	// we log them so the editor can display them to the user
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(output))
		log.Fatal(err)
	}

	// now we can load the temporary file into STDOUT
	// I am sure there is a better way to do this with less
	// allocation, but for now this works
	tmpFileHandle, err = os.Open(tmpFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer tmpFileHandle.Close()

	newFileContents, err := io.ReadAll(tmpFileHandle)
	if err != nil {
		log.Fatal(err)
	}

	// make it rain!
	fmt.Print(string(newFileContents))
}
