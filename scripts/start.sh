#!/bin/bash

network_type=${1:?}
os_type=${2:?}
arch_type=${3:?}
component_type=${4:?}

load_backend_offline() {
    pushd ./backend > /dev/null || exit 1
    ./load.sh || exit 2
    popd || exit 1
}

start_backend(){
    if [[ "${network_type}" == "offline" ]]; then
        load_backend_offline || exit 1
    fi
    pushd ./backend > /dev/null || exit 1
    ./start.sh || exit 2
    popd || exit 1
}

load_agent_offline() {
    pushd ./agent > /dev/null || exit 1
    ./load.sh || exit 2
    popd || exit 1
}

start_agent(){
    if [[ "${network_type}" == "offline" ]]; then
        load_agent_offline || exit 1
    fi
    pushd ./agent > /dev/null || exit 1
    ./start.sh || exit 2
    popd || exit 1
}

if [[ "${component_type}" == "backend" ]]; then
    start_backend || exit 1
elif [[ "${component_type}" == "agent" ]]; then
    start_agent || exit 1
fi






