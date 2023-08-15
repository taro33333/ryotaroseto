package scripts

import (
	"os"
)

const readmeFile = "README.md"
const head = `### Hi 😃
My name is Ryotaro Seto. I’m Japanese. I was born and raised in Japan.
### Skills

[<img src="./public/images/go.svg" width="40" height="40" />](https://golang.org/)
[<img src="./public/images/github.svg" width="40" height="40" />](https://github.com/)
[<img src="./public/images/kubernetes.svg" width="40" height="40" />](https://cloud.google.com/learn/what-is-kubernetes?hl=ja)
[<img src="./public/images/php.svg" width="40" height="40" />](https://php.org/)
[<img src="./public/images/python.svg" width="40" height="40" />](https://www.python.org/)
[<img src="./public/images/typescript.svg" width="40" height="40" />](https://www.typescriptlang.org/)`

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
