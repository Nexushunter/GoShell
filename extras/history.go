package extras

import "errors"

// @author  Hunter Breathat
// @License Copyright (R) 2018 Hunter Breathat
// @repo	{github,gitlab}.com/NexisHunter/GoShell/extras/

// Contains the history of the inputted commands if any skips empty \r\n | \n
// | \r input as it does not count as a command

func LoadHistory() (history []string) {
	return history
}

func SaveHistory(history []string) (err error) {
	if emptyHistory(history) {
		return errors.New("cannot save empty history")
	}
	return nil
}

func emptyHistory(history []string) bool {
	return len(history) == 0
}
