import requests
import pathlib
import re
import os

# API endpoints
GLOBAL_INDEXER_ENDPOINT = os.getenv("GLOBAL_INDEXER_ENDPOINT", "https://gi.rss3.dev")

# Read README.md file
root = pathlib.Path(__file__).parent.resolve()
readme_path = root / "README.md"
with readme_path.open() as f:
    readme_content = f.read()

# Get network list
networks_response = requests.get(f"{GLOBAL_INDEXER_ENDPOINT}/nta/networks")
networks = networks_response.json().get("data", []) if networks_response.status_code == 200 else []

# Filter out 'rss' from networks
networks = [network for network in networks if network != "rss"]

# Initialize the table
table = f"| Network/Worker | {' | '.join(networks)} |\n"
table += "|" + "---|" * (len(networks) + 1) + "\n"

# Initialize workers dictionary
workers_dict = {}

# Iterate through each network and get the supported workers list
for network in networks:
    WORKERS_API = f"{GLOBAL_INDEXER_ENDPOINT}/nta/networks/{network}/list_workers"
    network_workers_response = requests.get(WORKERS_API)
    network_workers = network_workers_response.json().get("data", []) if network_workers_response.status_code == 200 else []

    for worker in network_workers:
        workers_dict.setdefault(worker, []).append(network)

# Sort worker names
sorted_workers = sorted(workers_dict.keys())

# Move "core" worker to the top
if "core" in sorted_workers:
    sorted_workers.remove("core")
    sorted_workers.insert(0, "core")

# Generate table rows
for worker in sorted_workers:
    # Add superscript to "core" worker
    display_name = f"**{worker}** [^1]" if worker == "core" else worker
    row = f"| {display_name} |"
    for network in networks:
        row += " ✓ |" if network in workers_dict[worker] else "   |"
    table += row + "\n"

# Replace the placeholder in README.md
pattern = r"<!-- network-worker table starts -->.*?<!-- network-worker table ends -->"
replacement = f"<!-- network-worker table starts -->\n{table}<!-- network-worker table ends -->"
updated_content = re.sub(pattern, replacement, readme_content, flags=re.DOTALL)

# Write back to README.md file
with readme_path.open("w") as f:
    f.write(updated_content)
