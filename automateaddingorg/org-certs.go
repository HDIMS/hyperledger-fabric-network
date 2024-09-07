package main

import (
	"fmt"
)

func createCerts_sh(domain string, caorg string, port string, hosp string) {
	filecont := `set -x
mkdir -p /organizations/peerOrganizations/` + domain + `/
export FABRIC_CA_CLIENT_HOME=/organizations/peerOrganizations/` + domain + `/	
fabric-ca-client enroll -u https://admin:adminpw@` + caorg + `:` + port + ` --caname ` + caorg + ` --tls.certfiles "/organizations/fabric-ca/` + hosp + `/tls-cert.pem"

echo 'NodeOUs:
Enable: true
	ClientOUIdentifier:
		Certificate: cacerts/` + caorg + `-` + port + `-` + caorg + `.pem
		OrganizationalUnitIdentifier: client
	  PeerOUIdentifier:
		Certificate: cacerts/` + caorg + `-` + port + `-` + caorg + `.pem
		OrganizationalUnitIdentifier: peer
	  AdminOUIdentifier:
		Certificate: cacerts/` + caorg + `-` + port + `-` + caorg + `.pem
		OrganizationalUnitIdentifier: admin
	  OrdererOUIdentifier:
		Certificate: cacerts/` + caorg + `-` + port + `-` + caorg + `.pem
		OrganizationalUnitIdentifier: orderer' > "/organizations/peerOrganizations/` + domain + `/msp/config.yaml"
	
	
	
	fabric-ca-client register --caname ` + caorg + ` --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles "/organizations/fabric-ca/` + hosp + `/tls-cert.pem"
	
	
	
	fabric-ca-client register --caname ` + caorg + ` --id.name user1 --id.secret user1pw --id.type client --tls.certfiles "/organizations/fabric-ca/` + hosp + `/tls-cert.pem"
	
	
	
	
	fabric-ca-client register --caname ` + caorg + ` --id.name ` + hosp + `admin --id.secret ` + hosp + `adminpw --id.type admin --tls.certfiles "/organizations/fabric-ca/` + hosp + `/tls-cert.pem"
	
	
	
	fabric-ca-client enroll -u https://peer0:peer0pw@` + caorg + `:` + port + ` --caname ` + caorg + ` -M "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/msp" --csr.hosts peer0.` + domain + ` --csr.hosts  peer0-` + hosp + ` --tls.certfiles "/organizations/fabric-ca/` + hosp + `/tls-cert.pem"
	
	
	
	cp "/organizations/peerOrganizations/` + domain + `/msp/config.yaml" "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/msp/config.yaml"
	
	
	
	fabric-ca-client enroll -u https://peer0:peer0pw@` + caorg + `:` + port + ` --caname ` + caorg + ` -M "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/tls" --enrollment.profile tls --csr.hosts peer0.` + domain + ` --csr.hosts  peer0-` + hosp + ` --csr.hosts ` + caorg + ` --csr.hosts localhost --tls.certfiles "/organizations/fabric-ca/` + hosp + `/tls-cert.pem"
	
	
	
	
	cp "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/tls/tlscacerts/"* "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/tls/ca.crt"
	cp "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/tls/signcerts/"* "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/tls/server.crt"
	cp "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/tls/keystore/"* "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/tls/server.key"
	
	mkdir -p "/organizations/peerOrganizations/` + domain + `/msp/tlscacerts"
	cp "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/tls/tlscacerts/"* "/organizations/peerOrganizations/` + domain + `/msp/tlscacerts/ca.crt"
	
	mkdir -p "/organizations/peerOrganizations/` + domain + `/tlsca"
	cp "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/tls/tlscacerts/"* "/organizations/peerOrganizations/` + domain + `/tlsca/tlsca.` + domain + `-cert.pem"
	
	mkdir -p "/organizations/peerOrganizations/` + domain + `/ca"
	cp "/organizations/peerOrganizations/` + domain + `/peers/peer0.` + domain + `/msp/cacerts/"* "/organizations/peerOrganizations/` + domain + `/ca/ca.` + domain + `-cert.pem"
	
	
	fabric-ca-client enroll -u https://user1:user1pw@` + caorg + `:` + port + ` --caname ` + caorg + ` -M "/organizations/peerOrganizations/` + domain + `/users/User1@` + domain + `/msp" --tls.certfiles "/organizations/fabric-ca/` + hosp + `/tls-cert.pem"
	
	cp "/organizations/peerOrganizations/` + domain + `/msp/config.yaml" "/organizations/peerOrganizations/` + domain + `/users/User1@` + domain + `/msp/config.yaml"
	
	fabric-ca-client enroll -u https://org1admin:org1adminpw@` + caorg + `:` + port + ` --caname ` + caorg + ` -M "/organizations/peerOrganizations/` + domain + `/users/Admin@` + domain + `/msp" --tls.certfiles "/organizations/fabric-ca/` + hosp + `/tls-cert.pem"
	
	cp "/organizations/peerOrganizations/` + domain + `/msp/config.yaml" "/organizations/peerOrganizations/` + domain + `/users/Admin@` + domain + `/msp/config.yaml"
	
	{ set +x; } 2>/dev/null
	`
	output, err := createFile(hosp+"-certs.sh", filecont)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	fmt.Println("File created successfully.")
}

func createJob_org(hosp string) {
	filecont := `apiVersion: batch/v1
kind: Job
metadata:
  name: create-certs-` + hosp + `
spec:
  parallelism: 1
  completions: 1
  template:
    metadata:
      name: create-certs-` + hosp + `
    spec:
      volumes:
        - name: fabricfiles
          persistentVolumeClaim:
            claimName: mypvc
      containers:
        - name: create-certs-` + hosp + `
          image: hyperledger/fabric-ca-tools:latest
          resources:
            requests:
              memory: "300Mi"
              cpu: "300m"
            limits:
              memory: "500Mi"
              cpu: "350m"
          volumeMounts:
            - mountPath: /organizations
              name: fabricfiles
              subPath: organizations
            - mountPath: /scripts
              name: fabricfiles
              subPath: scripts
          command:
            - /bin/sh
            - -c
            - |
              ./scripts/` + hosp + `-certs.sh
      restartPolicy: Never
`
	output, err := createFile(hosp+"-job.yaml", filecont)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	fmt.Println("File created successfully.")
}
