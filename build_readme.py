import requests
import pathlib
import re
import os

# API endpoint
NETWORK_CONFIG_ENDPOINT = os.getenv("NETWORK_CONFIG_ENDPOINT", "https://gi.rss3.io/nta")

# Read README.md file
root = pathlib.Path(__file__).parent.resolve()
readme_path = root / "README.md"
with readme_path.open() as f:
    readme_content = f.read()

# Get network and worker data
response = requests.get(f"{NETWORK_CONFIG_ENDPOINT}/networks/config")
data = response.json()['data']

# Extract networks and workers
networks = []
workers = set()

for category in ['decentralized', 'federated', 'rss']:
    if category == 'rss':
        networks.append(data[category]['id'])
        workers.add(data[category]['worker_configs']['worker']['value'])
    else:
        for network in data[category]:
            networks.append(network['id'])
            for worker_config in network['worker_configs']:
                workers.add(worker_config['worker']['value'])

# Sort networks and workers
networks.sort()
workers = sorted(list(workers))

# Move "core" worker to the top if it exists
if "core" in workers:
    workers.remove("core")
    workers.insert(0, "core")

# Initialize the table
table = f"| Network/Worker | {' | '.join(networks)} |\n"
table += "|" + "---|" * (len(networks) + 1) + "\n"

# Generate table rows
for worker in workers:
    # Add superscript to "core" worker
    display_name = f"**{worker}** [^1]" if worker == "core" else worker
    row = f"| {display_name} |"
    for network in networks:
        # Check if the worker exists for this network
        worker_exists = False
        for category in ['decentralized', 'federated', 'rss']:
            if category == 'rss':
                if network == data[category]['id'] and data[category]['worker_configs']['worker']['value'] == worker:
                    worker_exists = True
                    break
            elif network in [n['id'] for n in data[category]]:
                network_data = next(n for n in data[category] if n['id'] == network)
                if any(w['worker']['value'] == worker for w in network_data['worker_configs']):
                    worker_exists = True
                    break
        row += " ✓ |" if worker_exists else "   |"
    table += row + "\n"

# Replace the placeholder in README.md
pattern = r"<!-- network-worker table starts -->.*?<!-- network-worker table ends -->"
replacement = f"<!-- network-worker table starts -->\n{table}<!-- network-worker table ends -->"
updated_content = re.sub(pattern, replacement, readme_content, flags=re.DOTALL)

# Write back to README.md file
with readme_path.open("w") as f:
    f.write(updated_content)
