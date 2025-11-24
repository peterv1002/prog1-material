package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ShortestAbc.
MAX. PUNKTE: 10
*/

// ShortestAbc erwartet eine Liste von Strings und liefert
// das kürzeste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
//
// Hinweis: Die Funktion muss nur mit kurzen Strings der Länge < 100 funktionieren.
func ShortestAbc(list []string) string {
	Shortest := "gedsuzfgdsrhouögvhieorhgoiehewuhriehwurhewuihruhwqiorhwhriorhgopfjewriogheiorwjgfopierhgöihero"
	Ausgabe := 0
	for i := 0; i < len(list); i++ {
		if len(list[i]) >= 3 {
			if list[i][:3] == "abc" {
				if len(Shortest) >= len(list[i]) {
					Shortest = list[i]
					Ausgabe++
				}

			}
		}

	}
	if Ausgabe != 0 {
		return Shortest
	}
	return ""

}
