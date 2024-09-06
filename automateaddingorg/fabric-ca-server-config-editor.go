package main

import (
	"fmt"
	// "errors"
)

func fabricCaServerConfigEditor(Org1CA string, caorg1 string, country string, state string, city string, domain string) {
	filecont := `version: 1.2.0

port: 7054

debug: false

crlsizelimit: 512000

tls:
  enabled: true
  certfile:
  keyfile:
  clientauth:
    type: noclientcert
    certfiles:

ca:
  name: ` + Org1CA + `
  keyfile:
  certfile:
  chainfile:

crl:
  expiry: 24h

registry:
  maxenrollments: -1

  identities:
     - name: admin
       pass: adminpw
       type: client
       affiliation: ""
       attrs:
          hf.Registrar.Roles: "*"
          hf.Registrar.DelegateRoles: "*"
          hf.Revoker: true
          hf.IntermediateCA: true
          hf.GenCRL: true
          hf.Registrar.Attributes: "*"
          hf.AffiliationMgr: true

db:
  type: sqlite3
  datasource: fabric-ca-server.db
  tls:
      enabled: false
      certfiles:
      client:
        certfile:
        keyfile:

ldap:
   enabled: false
   url: ldap://<adminDN>:<adminPassword>@<host>:<port>/<base>
   tls:
      certfiles:
      client:
         certfile:
         keyfile:
   attribute:
      names: ['uid','member']
      converters:
         - name:
           value:
    
      maps:
         groups:
            - name:
              value:

affiliations:
   org1:
      - department1
      - department2
   org2:
      - department1
   org3:
      - department1

signing:
    default:
      usage:
        - digital signature
      expiry: 8760h
    profiles:
      ca:
         usage:
           - cert sign
           - crl sign
         expiry: 43800h
         caconstraint:
           isca: true
           maxpathlen: 0
      tls:
         usage:
            - signing
            - key encipherment
            - server auth
            - client auth
            - key agreement
         expiry: 8760h

csr:
   cn: ` + caorg1 + `
   names:
      - C: ` + country + `
        ST: ` + state + `
        L: ` + city + `
        O: ` + caorg1 + `
        OU: ` + caorg1 + `
   hosts:
     - localhost
     - ` + domain + `
     - ` + caorg1 + `
   ca:
      expiry: 131400h
      pathlength: 1

bccsp:
    default: SW
    sw:
        hash: SHA2
        security: 256
        filekeystore:
            keystore: msp/keystore


cacount:

cafiles:

intermediate:
  parentserver:
    url:
    caname:

  enrollment:
    hosts:
    profile:
    label:

  tls:
    certfiles:
    client:
      certfile:
      keyfile:
`

	// Call createFile to write content to the file
	output, err := createFile("fabric-ca-server-config.yaml", filecont)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	fmt.Println("File created successfully.")
}
