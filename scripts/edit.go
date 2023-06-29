package scripts

import (
	"os"
)

const readmeFile = "README.md"
const docs_str = "![](./profile-3d-contrib/profile-gitblock.svg)"

func Edit() error {
	f, err := os.Create(readmeFile)
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
	}()
	_, err = f.Write([]byte(docs_str))
	if err != nil {
		return err
	}

	return nil
}
