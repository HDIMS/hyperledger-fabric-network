package main

import "fmt"

func caOrg1Editor(caorg1 string, org1 string) {
	filecont := `apiVersion: apps/v1
	kind: Deployment
	metadata:
	  name: ` + caorg1 + `
	spec:
	  selector:
		matchLabels:
		  app: ` + caorg1 + `
	  replicas: 1
	  template:
		metadata:
		  labels:
			app: ` + caorg1 + `
		spec:
		  volumes:
			- name: data
			  persistentVolumeClaim:
				claimName: mypvc
		  containers:
			- name: ` + caorg1 + `
			  image: hyperledger/fabric-ca:1.4.9
			  imagePullPolicy: "Always"
			  command:
				[
				  "fabric-ca-server" ,
				  "start", "-b" ,"admin:adminpw","--port","7054", "-d"
				]
			  resources:
				requests:
				  memory: "300Mi"
				  cpu: "250m"
				limits:
				  memory: "400Mi"
				  cpu: "350m"
			  env:
			   - name: FABRIC_CA_SERVER_CA_NAME
				 value: ` + caorg1 + `
			   - name: FABRIC_CA_SERVER_TLS_ENABLED
				 value: "true"
			   - name: FABRIC_CA_SERVER_CSR_CN
				 value: "` + caorg1 + `"
			   - name: FABRIC_CA_SERVER_CSR_HOSTS
				 value: "` + caorg1 + `"
			  volumeMounts:
				- name: data
				  mountPath: /etc/hyperledger/fabric-ca-server
				  subPath: organizations/fabric-ca/` + org1 + ``

	// Call createFile to write content to the file
	output, err := createFile("ca-"+org1+".yaml", filecont)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	fmt.Println("File created successfully.")
}
