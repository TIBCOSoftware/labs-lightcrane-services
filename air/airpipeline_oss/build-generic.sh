
app_name=$1

go get -u github.com/lib/pq
go get -u github.com/jvanderl/flogo-components/activity/tcmpub

flogo create -f ${app_name} app
cd app
cd src
go mod tidy
cd ..

export GOOS=linux ;\
export GOARCH=386;\
flogo build --verbose;
mv ./bin/app ./bin/app_386

export GOOS=linux ;\
export GOARCH=arm64;\
flogo build --verbose;
mv ./bin/app ./bin/app_arm64

export GOOS=linux ;\
export GOARCH=amd64;\
flogo build --verbose;
mv ./bin/app ./bin/app_amd64
