package main

import (
	"errors"
	"fmt"
	"math"
)



func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	wordset := make(map[string]bool)
	output := make(map[string][2][2]int)
	offset := [][2]int{{1,0},{-1,0},{0,-1},{0,1},{1,1},{-1,1},{1,-1},{-1,-1}}
	maxLen := math.MinInt32
	for _,word := range words {
        wordset[word] = true
        if len(word) > maxLen {
            maxLen = len(word)
        }
    }
	
	for i := 0; i<len(puzzle); i++ {
        for j :=0;j<len(puzzle[0]); j++ {
            dynamString := [3][3]string{}
            for k:=0;k<maxLen;k++ {
                for _,off := range offset {
                    nr,nc := off[0]*k, off[1]*k
                    dsr, dsc := off[0], off[1]
                    if dsr == -1 {
                        dsr = 2
                    }
                    if dsc == -1 {
                        dsc = 2
                    }
                    if (i+nr >= 0 && i+nr < len(puzzle)) && (j+nc >= 0 && j+nc < len(puzzle[0])) {
                        dynamString[dsr][dsc] += string(puzzle[i+nr][j+nc])
                        _, ok := wordset[dynamString[dsr][dsc]]
                        if ok {
                            output[dynamString[dsr][dsc]] = [2][2]int{{j,i},{j+nc, i+nr}}
                        }
                    }
                }		
            }
        }
	}
	shouldError := false
	for searchword,_ := range wordset {
        _, ok := output[searchword]
        if (!ok) {
            output[searchword] = [2][2]int{[2]int{-1,-1}, [2]int{-1,-1}}
            shouldError = true
        }
    }
	if (shouldError) {
        return output,errors.New("this is an error object.")
    }
	return output,nil
}
func main() {
	puzzle :=  []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"}
	words :=  []string{"clojure", "elixir", "ecmascript", "rust", "java", "lua"}
		coords, _ := Solve(words, puzzle)
		for key,value := range coords {
			scoord, ecoord := value[0], value[1]
			fmt.Println(key,":",scoord,",",ecoord)
		}
}