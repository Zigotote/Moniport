if [ $# -eq 0 ] 
then
    repo=$GOPATH/src/Moniport
else 
    repo=$1
fi

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

# Lancement d'api et templating
prog=( "api" "templating" )
for i in "${prog[@]}"
do
    cd $i
    go build
    ./$i &
    if [ ! $? -eq 0 ]
    then
        echo Erreur lors du build du dossier $i
        exit 1
    fi
    echo Lancement de $i : processus $!
    cd ..
done

# Lancement des récepteurs et recepteur-csv

prog=( "recepteur" "recepteur-csv" )
airports=( "NTE" "BES" )

for p in "${prog[@]}"
do
    cd $p
    go build
    if [ ! $? -eq 0 ]
    then
        echo Erreur lors du build du dossier $p
        exit 1
    else
        for i in "${airports[@]}"
        do
            ./$p -config $i &
            echo Lancement du $p de l aéroport de $i : processus $!
        done
    fi
    cd ..
done

# Lancement des capteurs

cd sensor
go build 
cd ../..

if [ ! $? -eq 0 ]
then
    echo "Erreur lors du build du dossier sensor"
    exit 1
else
    for config in ressources/config-files/publishers-config/*.json 
    do
        ./cmd/sensor/sensor -config $repo/$config &
        echo Lancement du capteur configuré dans le fichier $config : processus $!
    done
fi


