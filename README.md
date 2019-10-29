# :airplane: Moniport :airplane:

Initiation project to GO and MQTT

# Lancement du projet

Commandes à exécuter :

- mosquitto -v
- Lancer l'exécutable redis-server.exe
- ./run.sh <param> (donner le chemin d'accès au dossier Moniport. Si aucun argument n'est donné on considère que le projet est situé sous le chemin $GOPATH/src/Moniport)

:warning: Penser à arrêter les processus une fois le test terminé (ps + kill nbProcess) :skull_and_crossbones:

# Accès à l'API

Une API permet de visulaiser les données enregistrées sur le base REDIS. Deux requêtes sont disponibles :
* Visualisation des données d'un capteur entre deux bornes de temps : http://localhost:8080/measures/{airport}/{sensor}/{start}/{end} 
  * Airport = code IATA de l'aérport (par exemple NTE ou BES)
  * Sensor = type de capteur dont on veut les données (temp, press ou wind)
  * Start = date de début de récupération des données, au format "YYYY-MM-DD-hh-mm-ss"
  * End = date de fin de récupération des données, au format "YYYY-MM-DD-hh-mm-ss"
* Visualisation de la moyenne des données enregistrées sur une journée dans un aéroport : http://localhost:8080/avg-measures/{airport}/{date}
  * Airport = code IATA de l'aérport (par exemple NTE ou BES)
  * Date = jour pour lequel on souhaite connaître les moyennes des valeurs, au format "YYYY-MM-DD"

# Visualisation des données

Les données enregistrées sur la base REDIS peuvent être visualisées dans un navigateur à l'adresse http://localhost:8085/
