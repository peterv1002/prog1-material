package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	indexErster := -187
	indexLetzter := -187
	ele := []string{}

	if len(list) == 0 {
		return []string{}
	}

	for i := 0; i < len(list); i++ {
		if list[i] == first {
			indexErster = i
		}
		if list[i] == last {
			indexLetzter = i
		}
	}
	if indexErster == -187 {
		return []string{}
	}
	if indexLetzter == -187 {
		return []string{}
	}
	if indexLetzter < indexErster {
		return []string{}
	}
	for i := 0; i < len(list); i++ {
		if i < indexErster {
			ele = append(ele, list[i])
		}
		if i > indexLetzter {
			ele = append(ele, list[i])
		}

	}
	return ele

}
