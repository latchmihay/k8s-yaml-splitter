package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

// ---
// Taken directly from https://github.com/kubernetes/apimachinery/blob/master/pkg/util/yaml/decoder.go.

const (
	yamlSeparator = "\n---"
)

// splitYAMLDocument is a bufio.SplitFunc for splitting YAML streams into individual documents.
func splitYAMLDocument(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	sep := len([]byte(yamlSeparator))
	if i := bytes.Index(data, []byte(yamlSeparator)); i >= 0 {
		// We have a potential document terminator
		i += sep
		after := data[i:]
		if len(after) == 0 {
			// we can't read any more characters
			if atEOF {
				return len(data), data[:len(data)-sep], nil
			}
			return 0, nil, nil
		}
		if j := bytes.IndexByte(after, '\n'); j >= 0 {
			return i + j + 1, data[0 : i-sep], nil
		}
		return 0, nil, nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

// ---

type baseObject struct {
	bytes  []byte
	Kind   string `yaml:"kind"`
	ApiVer string `yaml:"apiVersion"`
	Meta   struct {
		Namespace   string            `yaml:"namespace"`
		Name        string            `yaml:"name"`
		Annotations map[string]string `yaml:"annotations,omitempty"`
	} `yaml:"metadata"`
}

func unmarshalObject(bytez []byte, dryRun bool, outputDir string) (error) {
	var base = baseObject{bytes: bytez}
	if err := yaml.Unmarshal(bytez, &base); err != nil {
		return makeUnmarshalObjectErr(err)
	}
	if len(base.Kind) > 0  && len(base.ApiVer) > 0 {
		fileName := fmt.Sprintf("%s-%s.yaml", base.Kind, base.Meta.Name)
		absolutePath := path.Join(outputDir, fileName)
		fmt.Printf("Found! type: %s | apiVersion: %s | name: %s | namespace: %s\n", base.Kind, base.ApiVer, base.Meta.Name, base.Meta.Namespace)
		if dryRun {
			fmt.Printf("==> DryRun: Writing %s\n", absolutePath)
			return nil
		}

		fmt.Printf("* Writing %s\n", absolutePath)
		// create file for writing
		f, err := os.Create(absolutePath)
		if err != nil {
			fmt.Println(err)
			fmt.Println("ERROR: Unable to create file " + absolutePath)
		}

		var out bytes.Buffer
		out.Write([]byte("---\n"))
		out.Write(base.bytes)

		// write to file
		byteWrote, err := out.WriteTo(f)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("* Wrote %d bytes to %s\n", byteWrote, absolutePath)
	}
	return nil
}

func makeUnmarshalObjectErr( err error) error {
	return errors.New("Could not parse. This likely means it is malformed YAML.")
}

func main() {
	dryRun := false
	if len(os.Args[1:]) <= 1 {
		fmt.Printf("Usage: %v %s %s\n", os.Args[0], "/path/to/combined-k8s.yaml", "/path/to/output/dir")
		fmt.Printf("Usage Dry Run: %v %s %s %s\n", os.Args[0], "/path/to/combined-k8s.yaml", "/path/to/output/dir", "-d")
		os.Exit(1)
	}

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		fmt.Printf("File %v does not exist on the system\n", os.Args[1])
		os.Exit(1)
	}
	if _, err := os.Stat(os.Args[2]); os.IsNotExist(err) {
		fmt.Printf("Directory %v does not exist on the system\n", os.Args[2])
		os.Exit(1)
	}

	if len(os.Args[1:]) == 3 && os.Args[3] == "-d" {
		dryRun = true
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	chunks := bufio.NewScanner(bytes.NewReader(content))
	initialBuffer := make([]byte, 4096)     // Matches startBufSize in bufio/scan.go
	chunks.Buffer(initialBuffer, 1024*1024) // Allow growth to 1MB
	chunks.Split(splitYAMLDocument)

	for chunks.Scan() {
		// It's not guaranteed that the return value of Bytes() will not be mutated later:
		// https://golang.org/pkg/bufio/#Scanner.Bytes
		// But we will be snaffling it away, so make a copy.
		bytes := chunks.Bytes()
		bytes2 := make([]byte, len(bytes), cap(bytes))
		copy(bytes2, bytes)
		if err := unmarshalObject(bytes2, dryRun, os.Args[2]); err != nil {
			fmt.Println(err, "parsing YAML doc")
		}
	}
}
