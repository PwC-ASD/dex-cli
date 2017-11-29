dex-cli - Un client gRPC pour Dex
================================

Ecrit en Go, ce client offre des primitives de gestion de clients Dex et des URLs de redirection associées.

L'API Dex a été enrichie de fonctionnalités supplémentaires disponibles sur ce [fork](dex-github). Pour plus d'informations sur les dépendances du projet, se référer aux configurations Glide.

# Installation

Le client peut s'installer comme un package Go ou bien être compilé directement depuis les sources.

Utiliser la commande ```go get``` pour le téléchargement et l'installation du package et de ses dépendances:
```
go get github.com/PwC-ASD/dex-cli
```

Sinon, utiliser le Makefile à disposition pour la compilation des sources:
```
make
```

[dex-github]: https://github.com/PwC-ASD/dex
