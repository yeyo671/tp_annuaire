package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Déclaration des flags pour les options de ligne de commande
	action := flag.String("action", "", "Action à effectuer: ajouter, supprimer, modifier, rechercher, lister")
	nom := flag.String("nom", "", "Nom complet du contact")
	tel := flag.String("tel", "", "Numéro de téléphone")
	flag.Parse()

	// Création d’un annuaire en mémoire
	ann := NewAnnuaire()

	ann.Ajouter("Charlie", "0811223344")
	ann.Ajouter("Alice", "0123456789")

	// Traitement selon l'action choisie
	switch *action {

	case "ajouter":
		// Vérifie que nom et téléphone sont fournis
		if *nom == "" || *tel == "" {
			fmt.Println("Nom et téléphone requis pour ajouter.")
			return
		}
		err := ann.Ajouter(*nom, *tel)
		if err != nil {
			fmt.Println("Erreur:", err)
		} else {
			fmt.Println("Contact ajouté avec succès.")
		}

	case "supprimer":
		err := ann.Supprimer(*nom)
		if err != nil {
			fmt.Println("Erreur:", err)
		} else {
			fmt.Println("Contact supprimé.")
		}

	case "modifier":
		err := ann.Modifier(*nom, *tel)
		if err != nil {
			fmt.Println("Erreur:", err)
		} else {
			fmt.Println("Contact modifié.")
		}

	case "rechercher":
		contact, err := ann.Rechercher(*nom)
		if err != nil {
			fmt.Println("Erreur:", err)
		} else {
			fmt.Printf("Contact trouvé: %s - %s\n", contact.Nom, contact.Tel)
		}

	case "lister":
		liste := ann.Lister()
		if len(liste) == 0 {
			fmt.Println("Aucun contact dans l'annuaire.")
			return
		}
		for _, c := range liste {
			fmt.Printf("- %s : %s\n", c.Nom, c.Tel)
		}

	default:
		// Aide si l’action n’est pas reconnue
		fmt.Println("Action inconnue. Utilisez --action ajouter|supprimer|modifier|rechercher|lister")
		os.Exit(1)
	}
}
