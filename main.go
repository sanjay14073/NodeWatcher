package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"github.com/fsnotify/fsnotify"
)

var cmd *exec.Cmd
var mutex sync.Mutex

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: nodewatcher <filename>")
	}
	filename := os.Args[1]

	// Start file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Add directories recursively to the watcher
	go watchDirectories(watcher)

	// Start the application
	restartApplication(filename)

	// Goroutine to listen for file events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("Event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("File modified:", event.Name)
					restartApplication(filename)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher error:", err)
			}
		}
	}()

	// Listen for user input
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if strings.TrimSpace(scanner.Text()) == "rs" {
				restartApplication(filename)
			}
		}
	}()

	// Block forever
	select {}
}

// Recursively add directories to the watcher
func watchDirectories(watcher *fsnotify.Watcher) {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			log.Println("Watching directory:", path)
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("Error adding directories:", err)
	}
}

// Restart the Node.js application
func restartApplication(filename string) {
	mutex.Lock()
	defer mutex.Unlock()

	// Kill the existing process if running
	if cmd != nil && cmd.Process != nil {
		log.Println("Stopping application...")
		err := cmd.Process.Kill()
		if err != nil {
			log.Println("Error stopping application:", err)
		}
		cmd.Wait()
	}

	// Start the new process
	log.Println("Starting application...")
	cmd = exec.Command("node", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal("Error starting application:", err)
	}
}
