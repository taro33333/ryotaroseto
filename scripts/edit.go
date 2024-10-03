package scripts

import (
	"log"
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
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if _, err = f.Write([]byte(header() + docs_str)); err != nil {
		return err
	}

	return nil
}
