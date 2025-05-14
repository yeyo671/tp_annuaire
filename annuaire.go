package main

import (
	"errors"
	"strings"
)

// Contact représente une personne dans l'annuaire.
type Contact struct {
	Nom string // Nom complet (nom + prénom)
	Tel string // Numéro de téléphone
}

// Annuaire gère les contacts via une map avec nom normalisé comme clé.
type Annuaire struct {
	contacts map[string]Contact
}

// NewAnnuaire crée un nouvel annuaire vide.
func NewAnnuaire() *Annuaire {
	return &Annuaire{contacts: make(map[string]Contact)}
}

// normalizeName met le nom en minuscules et supprime les espaces inutiles.
func normalizeName(nom string) string {
	return strings.ToLower(strings.TrimSpace(nom))
}

// Ajouter ajoute un contact si le nom n'existe pas encore.
func (a *Annuaire) Ajouter(nom, tel string) error {
	key := normalizeName(nom)
	if _, existe := a.contacts[key]; existe {
		return errors.New("contact déjà existant")
	}
	a.contacts[key] = Contact{Nom: nom, Tel: tel}
	return nil
}

// Supprimer enlève un contact s'il existe.
func (a *Annuaire) Supprimer(nom string) error {
	key := normalizeName(nom)
	if _, existe := a.contacts[key]; !existe {
		return errors.New("contact introuvable")
	}
	delete(a.contacts, key)
	return nil
}

// Modifier change le numéro de téléphone d’un contact existant.
func (a *Annuaire) Modifier(nom, nouveauTel string) error {
	key := normalizeName(nom)
	if contact, existe := a.contacts[key]; existe {
		contact.Tel = nouveauTel
		a.contacts[key] = contact
		return nil
	}
	return errors.New("contact introuvable")
}

// Rechercher retourne un contact correspondant au nom donné.
func (a *Annuaire) Rechercher(nom string) (*Contact, error) {
	key := normalizeName(nom)
	if contact, existe := a.contacts[key]; existe {
		return &contact, nil
	}
	return nil, errors.New("contact introuvable")
}

// Lister retourne tous les contacts de l'annuaire.
func (a *Annuaire) Lister() []Contact {
	var liste []Contact
	for _, contact := range a.contacts {
		liste = append(liste, contact)
	}
	return liste
}
