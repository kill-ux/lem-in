package lemIn

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func ReadData(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	Head := false
	Tail := false
	for i := 0; scanner.Scan(); i++ {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if i == 0 {
			num, err := strconv.Atoi(line)
			if err != nil || num < 1 || num > 1100 {
				return errors.New("invalid number of ants")
			}
			Data.Nants = num
			continue
		}
		if line == "##start" {
			if len(Data.Start) == 0 {
				Head = true
			} else {
				return errors.New("multiple start rooms")
			}
			continue
		}

		if line == "##end" {
			if len(Data.End) == 0 {
				Tail = true
			} else {
				return errors.New("multiple end rooms")
			}
			continue
		}

		if line[0] == '#' {
			continue
		}
		if Head && !Tail {
			if err := Add(line, "start"); err != nil {
				return err
			}
			Head = false
			continue
		}
		if Tail && !Head {
			if err := Add(line, "end"); err != nil {
				return err
			}
			Tail = false
			continue
		}
		if err := Add(line, ""); err != nil {
			return err
		}
	}

	return nil
}
