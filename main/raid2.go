package main

import (
	"reflect"
	"os"
	"fmt"
	"strings"
	piscine".."
)

func IsValideInput(args []string) bool {
	lon := 0
	if len(args) == 0|| len(args) == 1 || (piscine.Sqrt(len(args)) == 0 && len(args) > 1) {
		return false
	}
	for _,res:= range args {
		if reflect.TypeOf(res) != reflect.TypeOf("a") {
			return false
		}else{
			if lon == 0 {
				lon =len(res)
			}else{
				if lon != len(res) {
					return false
				}
			}
			for _,item:= range res {
				if string(item) != "." && strings.Count(res,string(item)) > 1 {
					return false
				}else{
					
				}
			}
		}
	}
	return true
}

func printBoard(board []string) {
	fmt.Println("+-------+-------+-------+")
	for row,res:=range board {
		fmt.Print("| ")
		for col,item := range res{
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%s ", string(item))
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func EmptyCell(grille **[]string) bool {
	count := 0
	for i:=0;i < 9; i++ {
		for j:= 0; j<9; j++ {
			if grille[i][j] == "." {
				count++
			}
		}
	}
	return count == 0
}

func Backtracking(grille **[]string) bool {
	if !EmptyCell(*grille){
		return true
	}
	for i:=0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grille[i][j] == '.' {
				for candidate := 9; candidate >= 1; candidate-- {
					grille[i][j] = string(candidate)
					if IsValidGrille(grille) {
						if Backtracking(grille) {
							return true
						}
						grille[i][j] =  string(0)
					} else {
						grille[i][j] = string(0)
					}
				}
				return false
			}
		}
	}
	return false
}

func main(){
	args:= os.Args[1:] 
	if len(args) == 0 {
		fmt.Println("Error")
	}else{
		if IsValideInput(args){
			grille :=args
			Backtracking(grille)
			printBoard(grille)
		}else{
			fmt.Println("Error")
		}
	}
}
