#!/bin/bash

# Function to display usage
usage() {
    echo "Usage: $0 -h <hosp> -p <hospca> -c <cahosp> -r <country> -s <state> -t <city> -d <domain>"
    exit 1
}

# Initialize variables
hosp=""
hospca=""
cahosp=""
country=""
state=""
city=""
domain=""
port=""

# Parse command line options
while getopts "h:p:c:r:s:t:d:" opt; do
    case ${opt} in
        h )
            hosp=$OPTARG
            ;;
        p )
            hospca=$OPTARG
            ;;
        c )
            cahosp=$OPTARG
            ;;
        r )
            country=$OPTARG
            ;;
        s )
            state=$OPTARG
            ;;
        t )
            city=$OPTARG
            ;;
        d )
            domain=$OPTARG
            ;;
        po )
            port=$OPTARG
            ;;
        \? )
            usage
            ;;
    esac
done
shift $((OPTIND -1))

# Check if required arguments are set
if [ -z "$hosp" ] || [ -z "$hospca" ] || [ -z "$cahosp" ] || [ -z "$country" ] || [ -z "$state" ] || [ -z "$domain" ] || [-z "$port"]; then
    usage
fi

# Run the Go script with the provided arguments and capture the output
output=$(go run . "$hosp" "$hospca" "$cahosp" "$country" "$state" "$city" "$domain" "$port")

# Display the output
echo "$output"


# Create the directories if they do not exist
mkdir ../fabric-ca/$hosp

# Copy files to their respective directories
cp fabric-ca-server-config.yaml "../fabric-ca/$hosp/"
cp ca-"$hosp".yaml "../caserver_k8s/"
cp ca-"$hosp"-service.yaml "../caserver_k8s/"
cp "$hosp"-certs.sh "../scripts/"
cp "$hosp"-job.yaml "../certificates_k8s/"

rm fabric-ca-server-config.yaml
rm ca-"$hosp".yaml
rm ca-"$hosp"-service.yaml
rm "$hosp"-certs.sh
rm "$hosp"-job.yaml

# Display the details of the moved files
# echo "Details of the moved files:"
# ls -l "../fabric-ca/"
# ls -l "../caserver_k8s/"
# ls -l "../certificates_k8s"

kubectl apply -f ../caserver_k8s/ca-"$hosp".yaml
kubectl apply -f ../caserver_k8s/ca-"$hosp"-service.yaml

chmod +x ../scripts/"$hosp"-certs.sh

kubectl apply -f ../certificates_k8s/"$hosp"-job.yaml