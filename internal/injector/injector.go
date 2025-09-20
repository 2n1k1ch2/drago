package injector

import (
	"fmt"
	"os"
)

func Build(paths []string) {
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("open %s error: %v\n", path, err)
		}
		replacer := NewReplacer()

		err = replacer.Replace(file)
		if err != nil {
			fmt.Printf("replace %s error: %v\n", path, err)
		}

	}

}
