# Annuaire Téléphonique

Une application simple d'annuaire téléphonique en ligne de commande développée en Go.

## Auteurs

- Leo Tran
- Eliot Meurillon

## Prérequis

- Go 1.24.3 ou supérieur
- Aucune dépendance externe requise

## Installation

1. Clonez le dépôt :

   ```
   git clone https://github.com/yeyo671/tp_annuaire.git
   cd tp_annuaire
   ```

## Utilisation

L'application s'utilise en ligne de commande avec les flags suivants :

- `--action` : L'action à effectuer (ajouter, supprimer, modifier, rechercher, lister)
- `--nom` : Le nom complet du contact
- `--tel` : Le numéro de téléphone du contact

### Commandes disponibles

#### Ajouter un contact

```
go run . --action ajouter --nom "Nom Prénom" --tel "0123456789"
```

#### Supprimer un contact

```
go run . --action supprimer --nom "Nom Prénom"
```

#### Modifier un contact existant

```
go run . --action modifier --nom "Nom Prénom" --tel "9876543210"
```

#### Rechercher un contact

```
go run . --action rechercher --nom "Nom Prénom"
```

#### Lister tous les contacts

```
go run . --action lister
```

## Exemples d'utilisation

### Ajouter un nouveau contact

```
go run . --action ajouter --nom "John Doe" --tel "0612345678"
```

### Lister tous les contacts de l'annuaire

```
go run . --action lister
```

### Rechercher un contact spécifique

```
go run . --action rechercher --nom "John Doe"
```

## Exécuter les tests

Pour exécuter les tests unitaires :

```
go test
```

## Structure du projet

- `main.go` : Point d'entrée, gestion des commandes CLI
- `annuaire.go` : Implémentation des fonctionnalités de l'annuaire
- `annuaire_test.go` : Tests unitaires

## Note

Par défaut, l'application initialise un annuaire avec deux contacts: "Charlie" et "Alice" pour démonstration.
