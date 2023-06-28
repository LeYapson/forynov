
# FORYNOV


FORYNOV est un projet réalisé par MAERTEN Axel et YAPI Théau qui à pour but de créer un forum en ligne en utilisant le langage de progammation GOLANG et SQLITE pour la base de données.


## Installation

Pour installer le projet il faut :

- Cloner le repository depuis GITEA :
```bash
    git clone https://ytrack.learn.ynov.com/git/ytheau/forum.git
```
- Puis se placer dans le dossier forum :
```bash
  cd /[chemin vers le dossier]/forum
```
    
## Lancer le projet

Pour lancer le projet, il existe  2 possibilités :

```bash
  go run .
```

ou avec Docker :
```bash
    docker build . -t forynov
    docker run -p 8080:8080 forynov
```


## License

[YNOV](https://www.ynov.com)

