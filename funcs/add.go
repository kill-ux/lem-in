package lemIn

import (
	"errors"
	"strconv"
	"strings"
)

func Add(line string, place string) error {
	whiteSplit := strings.Split(line, " ")
	if len(whiteSplit) == 3 && AddRoom(whiteSplit, place) {
		return nil
	} else if len(whiteSplit) == 1 && AddLink(whiteSplit[0]) {
		return nil
	}
	return errors.New("invalid room or link")
}

func AddRoom(room []string, place string) bool {
	if len(Data.Links) > 0 {
		return false
	}
	if room[0][0] == 'L' {
		return false
	}
	// coming back for visualize bonus(cordonnee verification)
	if _, ok := Data.Rooms[room[0]]; ok {
		return false
	}

	X, err1 := strconv.Atoi(room[1])
	Y, err2 := strconv.Atoi(room[2])
	if err1 != nil || err2 != nil || X < 0 || Y < 0 || !Equal([]int{X, Y}) {
		return false
	}
	if place == "start" {
		Data.Start = room[0]
	} else if place == "end" {
		Data.End = room[0]
	}
	Data.Rooms[room[0]] = []int{X, Y}
	return true
}

func Equal(parms []int) bool {
	for _, cord := range Data.Rooms {
		if cord[0] == parms[0] && cord[1] == parms[1] {
			return false
		}
	}
	return true
}

func AddLink(link string) bool {
	if !strings.Contains(link, "-") {
		return false
	}
	SplitHyphen := strings.Split(link, "-")
	if len(SplitHyphen) != 2 {
		return false
	}
	if SplitHyphen[0] == SplitHyphen[1] {
		return false
	}
	// verify duplicate links
	for _, lin := range Data.Links {
		if (SplitHyphen[0] == lin[0] && SplitHyphen[1] == lin[1]) || (SplitHyphen[0] == lin[1] && SplitHyphen[1] == lin[0]) {
			return false
		}
	}

	// verify rooms already exist

	if _, ok := Data.Rooms[SplitHyphen[0]]; !ok {
		return false
	}
	if _, ok := Data.Rooms[SplitHyphen[1]]; !ok {
		return false
	}
	Data.Links = append(Data.Links, SplitHyphen)
	return true
}
