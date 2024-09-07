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
	hosp := os.Args[1]
	hospca := os.Args[2]
	cahosp := os.Args[3]
	country := os.Args[4]
	state := os.Args[5]
	city := os.Args[6]
	domain := os.Args[7]
	port := os.Args[8]
	fabricCaServerConfigEditor(hospca, cahosp, country, state, city, domain)
	caOrg1Editor(cahosp, hosp)
	caOrg1ServiceEditor(cahosp, hosp)
	createCerts_sh(domain, cahosp, port, hosp)
	createJob_org(hosp)
}
