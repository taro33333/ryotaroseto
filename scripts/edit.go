package scripts

import (
	"os"
)

const readmeFile = "README.md"
const head = `### Hi 😃
My name is Ryotaro Seto. I’m Japanese. I was born and raised in Japan.
### Skills

[<img src="./public/images/go.svg" width="40" height="40" />](https://golang.org/)
[<img src="./public/images/github.svg" width="40" height="40" />](https://github.com/)`

const docs_str = "![](./profile-3d-contrib/profile-gitblock.svg)"

func Edit() error {
	f, err := os.Create(readmeFile)
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
	}()
	_, err = f.Write([]byte(head + docs_str))
	if err != nil {
		return err
	}

	return nil
}
