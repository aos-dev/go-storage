package posixfs

import (
	"fmt"
	"path/filepath"
)

func (c *Client) createDir(path string) (err error) {
	errorMessage := "posixfs createDir [%s]: %w"

	rp := filepath.Join(c.workDir, path)

	err = c.osMkdirAll(rp, 0755)
	if err != nil {
		return fmt.Errorf(errorMessage, path, handleOsError(err))
	}
	return
}

func (c *Client) getAbsPath(path string) string {
	return filepath.Join(c.workDir, path)
}

func (c *Client) getDirPath(path string) string {
	return filepath.Join(c.workDir, filepath.Dir(path))
}
