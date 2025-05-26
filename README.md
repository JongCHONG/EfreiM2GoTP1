# Annuaire CLI & Web en Go

## Présentation

Ce projet est un annuaire en ligne de commande et en mode web, développé en Go. Il permet de gérer une liste de contacts (nom, prénom, numéro de téléphone) depuis le terminal ou via une interface web. Ce projet a été réalisé dans le cadre du TP1 de programmation Go.

## Fonctionnalités

- **Lister** tous les contacts
- **Ajouter** un contact
- **Rechercher** un contact par nom
- **Supprimer** un contact
- **Modifier** un contact
- **Vérifier** l’existence d’un contact (unicité du nom)
- **Exporter** l’annuaire au format JSON
- **Importer** un annuaire au format JSON
- **Serveur web** affichant l’annuaire sur une page HTML
- **Tests unitaires** pour garantir la fiabilité du code

## Utilisation

Toutes les fonctionnalités s’utilisent via des flags en ligne de commande. Voici les commandes principales :

- **Ajouter un contact**
  ```
  go run main.go --action add --name "Charlie" --surname "Durand" --tel "0811223344"
  ```

- **Lister tous les contacts**
  ```
  go run main.go --action list
  ```

- **Rechercher un contact**
  ```
  go run main.go --action search --name "Alice"
  ```

- **Supprimer un contact**
  ```
  go run main.go --action remove --name "Charlie" --surname "Durand"
  ```

- **Modifier un contact**
  ```
  go run main.go --action update --name "Charlie" --surname "Durand" --tel "0999888777"
  ```

- **Exporter l’annuaire en JSON**
  ```
  go run main.go --action export
  ```
  > Exporte tous les contacts dans le fichier `annuaire.json`.

- **Importer un annuaire depuis un fichier JSON**
  ```
  go run main.go --action import
  ```
  > Importe les contacts depuis le fichier `annuaire.json`.

- **Lancer le serveur web**
  ```
  go run main.go --action serve
  ```
  > Lance un serveur web sur [http://localhost:8080](http://localhost:8080) affichant l’annuaire sur une page HTML.

## Structure du projet

- **main.go** : Point d’entrée du programme, gestion des flags et des actions
- **annuaire/** : Fonctions principales (ajouter, lister, rechercher, etc.)
- **exportation/** : Fonctions d’export JSON
- **importation/** : Fonctions d’import JSON
- **index.html** : Template HTML pour l’affichage web
- **annuaire_test.go** : Tests unitaires

## Tests

Le projet inclut au moins deux tests unitaires pour vérifier le bon fonctionnement des principales fonctionnalités.

## Membres du groupe

CHONG Jong hoa

---
