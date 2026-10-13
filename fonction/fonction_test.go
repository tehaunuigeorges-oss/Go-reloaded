package fonction_test

import (
	"Go-reloaded/fonction"
	"strings"
	"testing"
)

func TestAn(t *testing.T) {
	entree := "A Onix appeared in the cave"
	attendu := "An Onix appeared in the cave"

	resultat := fonction.An(strings.Fields(entree))
	resultatString := strings.Join(resultat, " ")

	if resultatString != attendu {
		t.Errorf("Erreur resultat attendu= %s", attendu)
	}
}

func TestBin(t *testing.T) {
	entree := "Brock caught 1010 (bin) Geodude in the cave"
	attendu := "Brock caught 10 Geodude in the cave"

	resultat := fonction.Bin(strings.Fields(entree))
	resultatString := strings.Join(resultat, " ")

	if resultatString != attendu {
		t.Errorf("Erreur resultat attendu= %s", attendu)
	}
}

func TestCap(t *testing.T) {
	entree := "ash (cap) left pallet town (cap, 2) with pikachu"
	attendu := "Ash left Pallet Town with pikachu"

	resultat := fonction.Cap(strings.Fields(entree))
	resultatString := strings.Join(resultat, " ")

	if resultatString != attendu {
		t.Errorf("Erreur resultat attendu= %s", attendu)
	}
}

func TestGuillemet(t *testing.T) {
	entree := "Ash shouted ' Pikachu ' and Pikachu answered ' pika pika chu '"
	attendu := "Ash shouted 'Pikachu' and Pikachu answered 'pika pika chu'"

	resultat := fonction.Guillemet(strings.Fields(entree))
	resultatString := strings.Join(resultat, " ")

	if resultatString != attendu {
		t.Errorf("Erreur resultat attendu= %s", attendu)
	}
}

func TestHex(t *testing.T) {
	entree := " Professor Oak gave me 1E (hex) Poke Balls"
	attendu := "Professor Oak gave me 30 Poke Balls"

	resultat := fonction.Hex(strings.Fields(entree))
	resultatString := strings.Join(resultat, " ")

	if resultatString != attendu {
		t.Errorf("Erreur resultat attendu= %s", attendu)
	}
}

func TestLow(t *testing.T) {
	entree := "CHARIZARD (low) USED FIRE BLAST (low, 2) ON BULBASAUR"
	attendu := "charizard USED fire blast ON BULBASAUR"

	resultat := fonction.Low(strings.Fields(entree))
	resultatString := strings.Join(resultat, " ")

	if resultatString != attendu {
		t.Errorf("Erreur resultat attendu= %s", attendu)
	}
}

func TestPonctuation(t *testing.T) {
	entree := "Ash said : Pikachu ,come here ; we must go ! Is that Mewtwo ? No ... wait !? It is a Ditto ."
	attendu := "Ash said: Pikachu, come here; we must go! Is that Mewtwo? No... wait!? It is a Ditto."

	resultat := fonction.Ponctuation(strings.Fields(entree))
	resultatString := strings.Join(resultat, " ")

	if resultatString != attendu {
		t.Errorf("Erreur resultat attendu= %s", attendu)
	}
}

func TestUp(t *testing.T) {
	entree := "pikachu (up) used thunder shock (up, 2) on the wild pidgey"
	attendu := "PIKACHU used THUNDER SHOCK on the wild pidgey"

	resultat := fonction.Up(strings.Fields(entree))
	resultatString := strings.Join(resultat, " ")

	if resultatString != attendu {
		t.Errorf("Erreur resultat attendu= %s", attendu)
	}
}
