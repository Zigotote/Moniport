cd $GOPATH/src/Moniport/cmd/sensor
go build Moniport/cmd/sensor

if [ ! $? -eq 0 ]
then
    echo "Erreur lors du build du dossier sensor"
fi

cd ..
for config in config-files/*.json 
do
    ./sensor/sensor -config $GOPATH/src/Moniport/cmd/$config  &
done