package local

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func write(text, filepath string) error {
	fmt.Printf("writing file %v\n", filepath)
	b := []byte(text)
	err := ioutil.WriteFile(filepath, b, 0644)
	switch  {
	case err != nil:
		fmt.Printf("error encountered writing file %v, %v\n", filepath, err)
		return err
	default:
		fmt.Printf("file wrote successfully\n")
		return nil
	}
}

func read(filepath string) (string, error) {

	fmt.Printf("reading file %v\n", filepath)
	content, err := ioutil.ReadFile(filepath)

	switch {
	case err != nil:
		fmt.Printf("error encountered reading file %v, %v\n", filepath, err)
		return "", err
	default:
		fmt.Printf("file %v read, content: %v\n", filepath, string(content))
		return string(content), nil
	}
}

func EntryPoint(w http.ResponseWriter, r *http.Request) {
	text := "Hello, World!"
	filepath := "/tmp/file"
	err := write(text, filepath)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error: %v\n", err)))
		return
	}

	content, err := read(filepath)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error: %v\n", err)))
	}
	w.Write([]byte(fmt.Sprintf("read: %v\n", content)))
}