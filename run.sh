if [ $# -eq 0 ] 
then
    repo=$GOPATH
else 
    repo=$1
fi
repo=$repo/src/Moniport 

# Build tous les fichier Go du dossier internal/
for dir in $(find internal/* -type d)
do
    find $dir/* -type d > /tmp/find.txt
    if [ ! -s /tmp/find.txt ]
    then
        echo Build du package $dir 
        cd $repo/$dir
        go build
        cd $repo/
    fi
    if [ ! $? -eq 0 ]
    then
        echo Erreur lors du build du dossier $dir
        exit 1
    fi
done

# DOC POUR LANCER DANS UN NOUVEAU TERMINAL
# https://askubuntu.com/questions/484993/run-command-on-anothernew-terminal-window
# Lancement des programmes du dossier cmd/
cd $repo/cmd

# Récepteur

cd recepteur
go build 
if [ ! $? -eq 0 ]
then
    echo "Erreur lors du build du dossier recepteur"
    exit 1
else
    ./recepteur &
    echo Lancement du récepteur sur le processus $!
fi


# Lancement des capteurs

cd ../sensor
go build Moniport/cmd/sensor
cd ..

if [ ! $? -eq 0 ]
then
    echo "Erreur lors du build du dossier sensor"
    exit 1
else
    for config in config-files/*.json 
    do
        ./sensor/sensor -config $GOPATH/src/Moniport/cmd/$config &
        echo Lancement du capteur configuré dans le fichier $config : processus $!
    done
fi


