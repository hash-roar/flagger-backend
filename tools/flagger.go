package tools

func GetReputationLevel(ReputationValue int) int {
	if ReputationValue <= 0 {
		return 0
	} else if ReputationValue > 0 && ReputationValue <= 20 {
		return 1
	} else if ReputationValue > 21 && ReputationValue <= 60 {
		return 2
	} else if ReputationValue > 60 && ReputationValue <= 120 {
		return 3
	} else if ReputationValue > 120 && ReputationValue <= 240 {
		return 4
	}
	return 5
}
