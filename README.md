# Annuaire CLI en Go

## Présentation

Ce projet est un annuaire en ligne de commande développé en Go. Il permet de gérer une liste de contacts (nom, prénom, numéro de téléphone) directement depuis le terminal, avec des commandes simples et rapides. Ce projet a été réalisé dans le cadre du TP1 de programmation Go.

## Fonctionnalités

- **Lister** tous les contacts
- **Ajouter** un contact
- **Rechercher** un contact par nom
- **Supprimer** un contact
- **Modifier** un contact
- **Tests unitaires** pour garantir la fiabilité du code

## Utilisation

Toutes les fonctionnalités s’utilisent via des flags en ligne de commande. Voici les commandes principales :

- **Ajouter un contact**
  ```
  go run main.go --action add --name "Charlie" --surname "Ronald" --tel "0811223344"
  ```

- **Lister tous les contacts**
  ```
  go run main.go --action list
  ```

- **Rechercher un contact**
  ```
  go run main.go --action search --name "John"    
  ```

- **Supprimer un contact**
  ```
  go run main.go --action remove --name "John" --surname "Doe"
  ```

- **Modifier un contact**
  ```
  go run main.go --action update --name "John" --surname "Doe" --tel "0" 
  ```

## Structure du projet

- **main.go** : Point d’entrée du programme, gestion des flags et des actions
- **annuaire.go** : Fonctions principales (ajouter, lister, rechercher, etc.)
- **annuaire_test.go** : Tests unitaires

## Tests

Le projet inclut au moins deux tests unitaires pour vérifier le bon fonctionnement des principales fonctionnalités.

## Bonus

- Import/export des contacts au format JSON (en option)
- Interface web rapide pour afficher l’annuaire (en option)

## Membres du groupe

- CHONG Jong Hoa

## Pour aller plus loin

Pour toute amélioration, suggestion ou retour, n’hésitez pas à ouvrir une issue ou une pull request sur ce dépôt.

---
