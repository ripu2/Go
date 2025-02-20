package voting

import "strconv"

func CanPersonVote(age *string) bool {
	var parsedValue, err = strconv.Atoi(*age)
	if err != nil {
		panic("Invalid Input")
	}
	if age == nil || parsedValue < 18 {
		return false
	}
	return true
}
