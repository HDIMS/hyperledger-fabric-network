import yaml
import sys

def add_organization(yaml_file, new_org_name):
    # Load existing YAML file
    with open(yaml_file, 'r') as file:
        config = yaml.safe_load(file)
    
    # Define the new organization
    new_org = {
        'Name': f'{new_org_name}MSP',
        'ID': f'{new_org_name}MSP',
        'MSPDir': f'../organizations/peerOrganizations/{new_org_name.lower()}.example.com/msp',
        'Policies': {
            'Readers': {
                'Type': 'Signature',
                'Rule': f"OR('{new_org_name}MSP.member')"
            },
            'Writers': {
                'Type': 'Signature',
                'Rule': f"OR('{new_org_name}MSP.member')"
            },
            'Admins': {
                'Type': 'Signature',
                'Rule': f"OR('{new_org_name}MSP.admin')"
            },
            'Endorsement': {
                'Type': 'Signature',
                'Rule': f"OR('{new_org_name}MSP.peer')"
            }
        },
        'AnchorPeers': [
            {
                'Host': f'peer0-{new_org_name.lower()}',
                'Port': 9051
            }
        ]
    }

    # Add new organization to the Organizations section
    config['Organizations'].append(new_org)

    # Update the Consortium to include the new organization
    consortium = config['Profiles']['TwoOrgsOrdererGenesis']['Consortiums']['SampleConsortium']
    if 'Organizations' not in consortium:
        consortium['Organizations'] = []
    consortium['Organizations'].append(new_org_name)

    # Update the Application Organizations
    profile = config['Profiles']['TwoOrgsChannel']
    application = profile['Application']
    if 'Organizations' not in application:
        application['Organizations'] = []
    application['Organizations'].append(new_org_name)

    # Save the updated YAML file
    with open(yaml_file, 'w') as file:
        yaml.dump(config, file, sort_keys=False)

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python add_organization.py <yaml_file> <new_org_name>")
        sys.exit(1)

    yaml_file = sys.argv[1]
    new_org_name = sys.argv[2]

    add_organization(yaml_file, new_org_name)
