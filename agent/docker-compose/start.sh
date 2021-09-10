export COMPOSE_HTTP_TIMEOUT=200
export LC_HOME=must_set_working_folder
export ServiceLocatorPort=5408
export ServiceLocatorIP=3.228.65.62
export AgentID=must_set_your_unique_id
export Image=bigoyang/agent
#export Image=bigoyang/agent_arm64

docker-compose rm -f
docker-compose pull
docker-compose up -d --build
