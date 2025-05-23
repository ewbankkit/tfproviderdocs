package check

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type FileCheck interface {
	Run(string) error
	RunAll([]string) error
}

type FileOptions struct {
	BasePath string
}

// FullPath returns the full path of the file, combining path portions from opts.BasePath and path.
func (opts *FileOptions) FullPath(path string) string {
	if opts.BasePath != "" {
		return filepath.Join(opts.BasePath, path)
	}

	return path
}

// FileSizeCheck verifies that documentation file is below the Terraform Registry storage limit.
func FileSizeCheck(fullpath string) error {
	fi, err := os.Stat(fullpath)

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] File %s size: %d (limit: %d)", fullpath, fi.Size(), RegistryMaximumSizeOfFile)
	if fi.Size() >= int64(RegistryMaximumSizeOfFile) {
		return fmt.Errorf("exceeded maximum (%d) size of documentation file for Terraform Registry: %d", RegistryMaximumSizeOfFile, fi.Size())
	}

	return nil
}

const (
	FileIgnoreDSStore = `.DS_Store`
)

var IgnoreFiles = []string{
	FileIgnoreDSStore,
}

// FileIgnoreCheck checks if the file should be ignored.
func FileIgnoreCheck(path string) bool {
	fileName := filepath.Base(path)

	for _, igfi := range IgnoreFiles {
		if igfi == fileName {
			return true
		}
	}

	return false
}
