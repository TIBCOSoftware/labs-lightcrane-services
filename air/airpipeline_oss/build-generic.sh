
app_name=$1

go get -u github.com/lib/pq
go get -u github.com/jvanderl/flogo-components/activity/tcmpub

mkdir bin
go mod tidy

export GOOS=linux ;\
export GOARCH=386;\
go build -o app ./...;
mv ./app ./bin/app_386

export GOOS=linux ;\
export GOARCH=arm64;\
go build -o app ./...;
mv ./app ./bin/app_arm64

export GOOS=linux ;\
export GOARCH=amd64;\
go build -o app ./...;
mv ./app ./bin/app_amd64
