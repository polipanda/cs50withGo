package week3

type sortFile string
type sortName string

const SORT1 sortFile = "sort1"
const SORT2 sortFile = "sort2"
const SORT3 sortFile = "sort3"
const BUBBLE sortName = "bubble sortFile"
const MERGE sortName = "merge sortFile"
const SELECTION sortName = "selection sortFile"

func returnAnswer(sort sortFile) sortName {
	switch sort {
	case SORT1:
		return BUBBLE
	case SORT2:
		return MERGE
	case SORT3:
		return SELECTION
	}

	return ""
}
