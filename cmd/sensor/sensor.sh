cd $GOPATH/src/Moniport/cmd
go build Moniport/cmd/sensor
for config in config-files/*.json 
do
    ./sensor/sensor -config $GOPATH/src/Moniport/cmd/$config  &
done