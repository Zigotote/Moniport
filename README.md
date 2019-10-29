# :airplane: Moniport :airplane:

Initiation project to GO and MQTT

# Lancement du projet

Commandes à exécuter :

- mosquitto -v
- Lancer l'exécutable redis-server.exe
- ./run.sh <param> (donner le chemin d'accès au dossier Moniport. Si aucun argument n'est donné la variable $GOPATH est utilisée)

:warning: Penser à arrêter les processus une fois le test terminé (ps + kill nbProcess) :skull_and_crossbones:

# Visualisation des données

Les données enregistrées sur la base REDIS peuvent être visualisées dans un navigateur:
- ce déplacer dans le repertoire `cmd/templating` du projet
- executer `go run main.go`
- Les données sont sur : http://localhost:8085/
