package GoGRF

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ensureDirectory(dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating directory %s: %w", dir, err)
	}
	return nil
}

func copyDirectory(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("error stating source directory: %v", err)
	}

	if !srcInfo.IsDir() {
		return fmt.Errorf("source is not a directory: %v", src)
	}

	if err := ensureDirectory(dst); err != nil {
		return err
	}

	srcDir, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening source directory: %v", err)
	}
	defer srcDir.Close()

	entries, err := srcDir.Readdir(-1)
	if err != nil {
		return fmt.Errorf("error reading source directory: %v", err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDirectory(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("error copying file: %v", err)
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("error stating source file: %v", err)
	}

	err = os.Chmod(dst, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("error setting destination file permissions: %v", err)
	}

	return nil
}
