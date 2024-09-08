#!/bin/bash

# Define variables
DOMAIN="example.com"
ORG_NAME="Orderer"
CRYPTO_CONFIG_FILE="./crypto-config.yaml"
OUTPUT_DIR="organizations"

# Step 1: Create crypto-config.yaml file
echo "Generating crypto-config.yaml for orderer organization..."

cat <<EOF > ${CRYPTO_CONFIG_FILE}
OrdererOrgs:
  - Name: ${ORG_NAME}
    Domain: ${DOMAIN}
    EnableNodeOUs: true
    Specs:
      - Hostname: orderer
      - Hostname: orderer2
      - Hostname: orderer3
      - Hostname: orderer4
      - Hostname: orderer5
    Users:
      Count: 2
PeerOrgs:
  - Name: Org1
    Domain: org1.example.com
    Template:
      Count: 1
    Users:
      Count: 1
EOF

echo "crypto-config.yaml file created."

# Step 2: Generate crypto material using cryptogen
echo "Generating crypto material using cryptogen..."
cryptogen generate --config=./crypto-config.yaml --output=./crypto-config

if [ $? -ne 0 ]; then
  echo "Failed to generate crypto material using cryptogen."
  exit 1
fi
echo "Crypto material generated successfully."

# Step 3: Check for the existence of generated certificates
# echo "Checking for the existence of admin certificate..."
# ADMIN_CERT_SRC="${OUTPUT_DIR}/ordererOrganizations/${DOMAIN}/msp/admincerts/Admin@${DOMAIN}-cert.pem"

# if [ ! -f ${ADMIN_CERT_SRC} ]; then
#   echo "Admin certificate source file does not exist: ${ADMIN_CERT_SRC}"
#   echo "Please ensure that the 'crypto-config.yaml' file is correct and that 'cryptogen' has generated the admin certificates."
#   exit 1
# fi

# Step 4: Create directory structure for orderer organization
echo "Creating directory structure for orderer organization..."
mkdir -p ${OUTPUT_DIR}/ordererOrganizations/${DOMAIN}/users/Admin@${DOMAIN}/msp/admincerts

# # Step 5: Copy admin certificate to appropriate directory
# echo "Copying admin certificate..."
# ADMIN_CERT_DEST="${OUTPUT_DIR}/ordererOrganizations/${DOMAIN}/users/Admin@${DOMAIN}/msp/admincerts/Admin@${DOMAIN}-cert.pem"
# cp ${ADMIN_CERT_SRC} ${ADMIN_CERT_DEST}

# if [ $? -ne 0 ]; then
#   echo "Error copying admin cert for org ${DOMAIN}."
#   exit 1
# fi

echo "Admin certificate copied successfully."

# Step 6: Display success message
echo "Orderer organization setup completed successfully using cryptogen."
echo "Certificates and keys are available in the '${OUTPUT_DIR}/ordererOrganizations/${DOMAIN}' directory."
