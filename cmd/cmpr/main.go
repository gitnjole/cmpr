package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/gitnjole/cmpr/internal/api"
)

func toURLFriendly(movie string) string {
	return url.QueryEscape(strings.ToLower(movie))
}

func findCommonCast(cast1, cast2 []api.Actor) []CommonActors {
	castMap := make(map[int]ActorInfo)
	for _, actor := range cast1 {
		castMap[actor.ID] = ActorInfo{
			Name:      actor.Name,
			Character: actor.Character,
		}
	}

	var common []CommonActors
	for _, actor := range cast2 {
		if info, found := castMap[actor.ID]; found {
			common = append(common, CommonActors{
				Name:            info.Name,
				CharacterMovie1: info.Character,
				CharacterMovie2: actor.Character,
			})
		}
	}

	return common
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: cmpr <Movie1> <Movie2>")
		os.Exit(1)
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]

	cast1, movieTitle1, err := api.GetCast(toURLFriendly(arg1))
	if err != nil {
		fmt.Printf("Error fetching cast for `%s`: %s", arg1, err)
		os.Exit(1)
	}

	cast2, movieTitle2, err := api.GetCast(toURLFriendly(arg2))
	if err != nil {
		fmt.Printf("Error fetching cast for `%s`: %s", arg2, err)
		os.Exit(1)
	}

	commonCast := findCommonCast(cast1, cast2)

	if len(commonCast) < 1 {
		fmt.Println("\033[31mNo common actors found!\033[0m")
		os.Exit(0)
	}

	fmt.Printf("Common cast found: \n")
	for _, actor := range commonCast {
		fmt.Printf("\033[34m%s\033[0m\n", actor.Name)
		fmt.Printf("as \033[34m%s\033[0m in \033[32m%s\033[0m\n", actor.CharacterMovie1, movieTitle1)
		fmt.Printf("as \033[34m%s\033[0m in \033[32m%s\033[0m\n", actor.CharacterMovie2, movieTitle2)

		fmt.Println(strings.Repeat("-", 62))
	}
}

type ActorInfo struct {
	Name      string
	Character string
}

type CommonActors struct {
	Name            string
	CharacterMovie1 string
	CharacterMovie2 string
}
