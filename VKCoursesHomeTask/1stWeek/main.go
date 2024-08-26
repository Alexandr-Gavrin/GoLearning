package main

import (
	"fmt"
	"io"
	"os"
)

func dirTreeLever(out io.Writer, path string, printMemory bool, levelStr string) error {
	files, err := os.ReadDir(path)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	for index, file := range files {
		var str string
		var fl bool

		if printMemory {
			if index != len(files)-1 {
				str = levelStr + fmt.Sprintf("├───%s", file.Name())
			} else {
				str = levelStr + fmt.Sprintf("└───%s", file.Name())

			}

		} else {
			fl = false
			for j := index + 1; j < len(files); j++ {
				if files[j].IsDir() {
					fl = true
					break
				}
			}
			if fl {
				str = levelStr + fmt.Sprintf("├───%s", file.Name())
			} else {
				str = levelStr + fmt.Sprintf("└───%s", file.Name())
			}
		}

		if file.IsDir() {
			fmt.Fprintln(out, str)
			if !printMemory {
				if fl {
					dirTreeLever(out, path+fmt.Sprintf("/%s", file.Name()), printMemory, levelStr+"│\t")
				} else {
					dirTreeLever(out, path+fmt.Sprintf("/%s", file.Name()), printMemory, levelStr+"\t")
				}
			} else {
				if index != len(files)-1 {
					dirTreeLever(out, path+fmt.Sprintf("/%s", file.Name()), printMemory, levelStr+"│\t")
				} else {
					dirTreeLever(out, path+fmt.Sprintf("/%s", file.Name()), printMemory, levelStr+"\t")
				}
			}

		} else {
			if printMemory {
				info, err := file.Info()
				if err != nil {
					return fmt.Errorf(err.Error())
				}
				if size := info.Size(); size != 0 {
					str += fmt.Sprintf(" (%db)", size)
				} else {
					str += " (empty)"
				}
				fmt.Fprintln(out, str)

			}

		}
	}
	return nil
}

func dirTree(output io.Writer, path string, printMemory bool) error {
	err := dirTreeLever(output, path, printMemory, "")
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
