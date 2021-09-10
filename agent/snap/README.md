## To use the agent snap
1.- Install : sudo snap install lightcrane-air-agent_0.1.0_multi.snap --devmode

2.- Start : sudo lightcrane-air-agent -p=<platform i.e.,arm64 default:amd64> -v=<version i.e., 0.1.0> -n=<agent id, unique id per lightcrane deployer> -d=<home directory for agent downloading, absolute path> -s=<lightcrane ip i.e., 3.228.65.62>
  
3.- Stop : sudo lightcrane-air-agent -p=<platform i.e.,arm64 default:amd64> -v=<version i.e., 0.1.0> stop
  
4.- Uninstall : sudo snap remove lightcrane-air-agent
