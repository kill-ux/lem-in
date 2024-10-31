package lemIn

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Lem() error {
	// PrintAll()
	if len(Data.Start) == 0 {
		return errors.New("no start room found")
	}
	if len(Data.End) == 0 {
		return errors.New("no end room found")
	}

	for room := range Data.Rooms {
		for _, link := range Data.Links {
			if link[0] == room {
				Data.Realtions[room] = append(Data.Realtions[room], link[1])
			}
			if link[1] == room {
				Data.Realtions[room] = append(Data.Realtions[room], link[0])
			}
		}
	}
	Data.Paths = append(Data.Paths, []string{Data.Start})
	// fmt.Println(Data.Realtions)

	if !FindPaths() {
		return errors.New("no links")
	}
	PathGroups()

	PrintAll()
	fmt.Println()
	for _, group := range Data.PathGroups {
		GoAnts(group)
	}
	min := 0
	for step, _ := range Data.StepCalc {
		min = step
		break
	}
	for step, _ := range Data.StepCalc {
		if step < min {
			min = step
		}
	}

	fmt.Print(Data.StepCalc[min])
	return nil
}

func GoAnts(group [][]string) {
	Ants := GroupAnts(group)
	lin := make(map[string][]string)
	removed := make(map[string]string)
	for i, pathAnt := range Ants {
		for _, ant := range pathAnt {
			lin["L"+strconv.Itoa(ant)] = group[i][1:]
		}
	}

	fullRooms := ""
	steps := 0
	result := ""
	fullLinks := ""
	for TabsIsNil(lin) {
		steps++
		str := ""
		for key, value := range lin {
			if len(value) != 0 && !strings.Contains(fullRooms, " "+value[0]+" ") {
				if _, ok := removed[key]; ok {
					if !strings.Contains(fullLinks, " "+removed[key]+"-"+value[0]+" ") {
						fullLinks += " " + removed[key] + "-" + value[0] + " "
					} else {
						continue
					}
				} else {
					if !strings.Contains(fullLinks, " "+Data.Start+"-"+value[0]+" ") {
						fullLinks += " " + Data.Start + "-" + value[0] + " "
					} else {
						continue
					}
				}
				if value[0] != Data.End {
					fullRooms += " " + value[0] + " "
				}
				str += " " + key + "-" + value[0]
				removed[key] = lin[key][0]
				lin[key] = lin[key][1:]
			}
		}
		fullRooms = ""
		fullLinks = ""
		result += strings.TrimSpace(str) + "\n"
	}
	Data.StepCalc[steps] = result
}

func PathGroups() {
	tempath := Data.Paths
	for range Data.Paths {
		group := [][]string{tempath[0]}
		tempath = append(tempath[1:], tempath[0])
		for _, path := range tempath {
			fullRooms := ""
			for _, P := range group {
				for _, room := range P[1 : len(P)-1] {
					fullRooms += " " + room + " "
				}
			}
			roomExist := false
			for _, room := range path {
				if strings.Contains(fullRooms, room) {
					roomExist = true
				}
			}
			if !roomExist {
				group = append(group, path)
			}

		}
		Data.PathGroups = append(Data.PathGroups, group)
	}
}

func TabsIsNil(Tab map[string][]string) bool {
	for _, value := range Tab {
		if len(value) != 0 {
			return true
		}
	}
	return false
}

func GroupAnts(group [][]string) [][]int {
	Ants := [][]int{}
	for range group {
		Ants = append(Ants, []int{})
	}
	counter := 0
	for i := 0; i < len(group); {
		path := group[i]
		if (i < len(group)-1 && len(path)+len(Ants[i]) <= len(group[i+1])+len(Ants[i+1])) || i == len(group)-1 || len(group) == 1 {
			counter++
			Ants[i] = append(Ants[i], counter)
			if counter == Data.Nants {
				break
			}
			i = 0
		} else {
			i++
		}
	}
	return Ants
}

func FindPaths() bool {
	Tab := [][]string{}
start:
	var path []string
	if len(Data.Paths) != 0 {
		path = Data.Paths[0]
	} else {
		goto end
	}
	if path[len(path)-1] == Data.End {
		Tab = append(Tab, path)
		Data.Paths = Data.Paths[1:]
		goto start
	}
	for _, value := range Data.Realtions[path[len(path)-1]] {
		Add := true
		for _, room := range path {
			if room == value {
				Add = false
				break
			}
		}
		if !Add {
			continue
		}
		var tab []string
		tab = append(tab, path...)
		tab = append(tab, value)
		Data.Paths = append(Data.Paths, tab)
	}
	Data.Paths = Data.Paths[1:]
	goto start
end:
	// SortPath()
	Data.Paths = Tab
	if len(Data.Paths) == 0 {
		return false
	}
	return true
}

// func SortPath() {
// 	sort.Slice(Data.Paths, func(i, j int) bool {
// 		return len(Data.Paths[i]) < len(Data.Paths[j])
// 	})
// }
