# STOUK-GAME
An illegal random casinon online

si vous n'utiliser pas les conteneur docker pour faire tourner le projet ca va etre compliquer 
vous devrez parametré votre propre Base de données et modifier les premiere ligne du fichier createDB
[ici](./app/Stouk/database/createDB.go) la ligne 23

et il faudra installer cette librairie
go get -u github.com/go-sql-driver/mysql

sinon docker compose up -d dans le fichier et ca part (il faudra vous rendre dans le conteneur golang pour lancer le fichier main dans /app go run ./main/main.go

le site permet de crée des utilisateur pour jouer a un jeux de dé avec un ladder est un systeme de mmr (pour paser un user admin utiliser le pannel phpmyadmin)


il y a :
- systeme de connection via email et mdp stocké avec Bcrypt
- Systeme de cookie pour gerer les sessions
- page acceuil avec leaderboard
- systeme de boutique pour amelioré les dés
- bouton jouer
- gestion de la queue en direct avec du pairing dynamique
- gestion du ladder avec du mmr et de l'elo
- gestion des historiques des joueurs
