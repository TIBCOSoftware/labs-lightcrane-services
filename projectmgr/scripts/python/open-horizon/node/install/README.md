# Make sure env.cfg & node.policy.json is in the folder

# Setup
source env.cfg

# Install Agent
curl -sSL https://github.com/open-horizon/anax/releases/latest/download/agent-install.sh | bash -s -- -i anax: -k css: -c css: -w '*' -T 120

# Register Node
hzn register -m 'p50-ubuntu-syeven' -o 'myorg' -u 'admin:NRtypkkV2fO514hZcvVz0Z194AhiPz' -n 'p50-ubuntu-steven:' -t '120' --policy 'node.policy.json'