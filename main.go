package main

import (
	"fmt"
	"log"
	"os"
)

// This could have been a shell script. Now it's cross-platform shell script.
// Except that it's not, because I don't know if windows has a PWD environment variable.
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <directory containing repos>\n", os.Args[0])
	}

	dir, err := os.ReadDir(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to open directory: %v\n", err)
	}

	if err := walkTree(os.Args[1], dir); err != nil {
		log.Fatalf("Problem enumerating directory: %v\n", err)
	}

	// TODO: Print the number of freed bytes and number of cleaned directories.
	fmt.Println("Done!")
}

// This seems like a very parallel problem.
// Goroutines might be overkill for this, but I still feel like using them.
func walkTree(prefix string, dir []os.DirEntry) error {
	for _, entry := range dir {
		if entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				return err
			}

			// Don't investigate hidden files.
			// Maybe a more robust filtering system would be nice.
			if info.Name()[0] == '.' {
				continue
			}

			nextPath := prefix + "/" + info.Name()
			subDir, err := os.ReadDir(nextPath)
			if err != nil {
				return err
			}

			err = walkTree(nextPath, subDir) // This is why it doesn't work silly!
			if err != nil {
				return err
			}
		} else {
			info, err := entry.Info()
			if err != nil {
				return err
			}

			// Writing this way so I can easily add more languages in the future
			name := info.Name()
			switch name {
			// Invariant: There should only ever be one of these in a given directory.
			case "Cargo.toml":
				return clean(prefix, name, "target")
			}
		}
	}

	return nil
}

func clean(dirName, projectFile, targetDir string) error {
	fmt.Printf("Found `%s` in `%s`, deleting target dir `%s`.\n", projectFile, dirName, targetDir)

	return os.RemoveAll(dirName + "/" + targetDir)
}
