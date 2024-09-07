package main

import (
	"fmt"
)

func caOrg1ServiceEditor(caorg1 string, org1 string, port string) {
	filecont := `apiVersion: v1
kind: Service
metadata:
  name: ` + caorg1 + `
  labels:
    app: ` + caorg1 + `
spec:
  type: ClusterIP
  selector:
    app: ` + caorg1 + `
  ports:
    - protocol: TCP
      targetPort: ` + port + `
      port: ` + port + `
`
	// Call createFile to write content to the file
	output, err := createFile("ca-"+org1+"-service.yaml", filecont)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	fmt.Println("File created successfully.")
}
