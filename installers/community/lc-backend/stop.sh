export VERSION=0.7.0
export USER_HOME=~
export LC_HOME=./LC
export LC_EXT=$LC_HOME/runtimeGOPath

if [ -n "$1" ]; then
    case "$1" in
    arm64)
    export ARCH=-$1
    ;;
    amd64)
    export ARCH=""
    ;;
    esac
else
	export ARCH=""
fi

docker-compose down