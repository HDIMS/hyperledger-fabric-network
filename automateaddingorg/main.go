package main

import (
	"os"
)

func createFile(name string, content string) (*os.File, error) {
	// Create the YAML file
	file, err := os.Create(name)
	if err != nil {
	}

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		file.Close()
		return nil, err
	}

	return file, nil
}

func main() {
	fabricCaServerConfigEditor("HOSP1CA", "ca-hosp1", "INDIA", "CG", "RAIPUR", "hehe.hosp1.com")
	caOrg1Editor("ca-hosp1", "hosp1")
	caOrg1ServiceEditor("ca-hosp1", "hosp1")
}
