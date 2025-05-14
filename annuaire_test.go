package main

import "testing"

func TestAjouterEtRechercher(t *testing.T) {
	ann := NewAnnuaire()
	err := ann.Ajouter("Alice Dupont", "0123456789")
	if err != nil {
		t.Errorf("Erreur ajout: %v", err)
	}

	contact, err := ann.Rechercher("Alice Dupont")
	if err != nil {
		t.Error("Contact introuvable après ajout")
	}
	if contact.Tel != "0123456789" {
		t.Error("Numéro de téléphone incorrect")
	}
}

func TestAjouterDoublon(t *testing.T) {
	ann := NewAnnuaire()
	_ = ann.Ajouter("Bob Marley", "0102030405")
	err := ann.Ajouter("Bob Marley", "1111111111")
	if err == nil {
		t.Error("Ajout de doublon non détecté")
	}
}

func TestSupprimer(t *testing.T) {
	ann := NewAnnuaire()
	_ = ann.Ajouter("Test Delete", "0000000000")
	err := ann.Supprimer("Test Delete")
	if err != nil {
		t.Errorf("Erreur lors de la suppression: %v", err)
	}
}

func TestModifier(t *testing.T) {
	ann := NewAnnuaire()
	_ = ann.Ajouter("Charlie", "0811223344")
	err := ann.Modifier("Charlie", "0999888777")
	if err != nil {
		t.Errorf("Erreur modification: %v", err)
	}
	c, _ := ann.Rechercher("Charlie")
	if c.Tel != "0999888777" {
		t.Error("Le numéro n'a pas été modifié correctement")
	}
}
