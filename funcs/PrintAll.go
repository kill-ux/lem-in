package lemIn

import "fmt"

func PrintAll() {
	fmt.Println(Data.Nants)
	for room, cord := range Data.Rooms {
		if room == Data.Start {
			fmt.Println("##start")
		} else if room == Data.End {
			fmt.Println("##end")
		}
		fmt.Println(room, cord[0], cord[1])
	}
	for _, link := range Data.Links {
		fmt.Println(link[0] + "-" + link[1])
	}
}
