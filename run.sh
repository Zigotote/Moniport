cd $GOPATH/src/Moniport/cmd/recepteur

# Lancement du recepteur

go build Moniport/cmd/recepteur
if [ ! $? -eq 0 ]
then
    echo "Erreur lors du build du dossier recepteur"
fi
./recepteur &
echo Lancement du récepteur sur le processus $!

# Lancement des capteurs

cd ../sensor
go build Moniport/cmd/sensor

if [ ! $? -eq 0 ]
then
    echo "Erreur lors du build du dossier sensor"
fi

cd ..
for config in config-files/*.json 
do
    ./sensor/sensor -config $GOPATH/src/Moniport/cmd/$config &
    echo Lancement du capteur configuré dans le fichier $config : processus $!
done
