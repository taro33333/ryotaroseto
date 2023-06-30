package scripts

import (
	"os"
)

const readmeFile = "README.md"
const head = "### Hi 😃
My name is Ryotaro Seto. I’m Japanese. I was born and raised in Japan."
const docs_str = "![](./profile-3d-contrib/profile-gitblock.svg)"

func Edit() error {
	f, err := os.Create(readmeFile)
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
	}()
	_, err = f.Write([]byte(head+"/n"docs_str))
	if err != nil {
		return err
	}

	return nil
}
