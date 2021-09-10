# Install Open-Horizon Agent
- Make sure env.cfg & node.policy.json is in the folder
- Setup 
source env.cfg
- Install Agent (make sure oh hub's 3090, 9443 port is accessible)
curl -sSL https://github.com/open-horizon/anax/releases/latest/download/agent-install.sh | bash -s -- -i anax: -k css: -c css: -w '*' -T 120
- Register Node
hzn register -m 'p50-ubuntu-syeven' -o 'myorg' -u 'admin:NRtypkkV2fO514hZcvVz0Z194AhiPz' -n 'p50-ubuntu-steven:' -t '120' --policy 'node.policy.json'
# Install RTSF
- Install VAS
git clone https://github.com/intel/video-analytics-serving.git
- Install loss-detection-app
git clone https://github.com/intel-iot-devkit/rtsf-at-checkout-reference-design.git
cp -r ../artifacts/rtsf-at-checkout-reference-design/loss-detection-app ./ 
- Inatall loss-detection-demo