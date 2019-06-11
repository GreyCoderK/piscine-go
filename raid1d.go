package piscine

import "fmt"

func Raid1c(x,y int) {
    for i:=0; i < y; i++ {//boucle pour la longueur
        for j:=0; j < x; j++ {//boucle pour la largeur
            if i == 0 {//première ligne
                if j == 0 || j == x - 1 {
            if j == 0{//première cologne
                fmt.Print("A")
            }
            if j == x-1{//dernière cologne
                if j!=0{
                    fmt.Print("C")
                }
                
            }                    
                }else{//colognes du milieu
                    fmt.Print("B")
                }
            }else if i == y - 1{//dernière ligne
        if j == 0 || j == x - 1 {
            if j == 0{//première cologne
                fmt.Print("A")
            }
            if j == x-1{//dernière cologne
                if j!=0{
                    fmt.Print("C")
                }
            }                    
                }else{//colognes du milieu
                    fmt.Print("B")
                }
            }else{//lignes du milieu
                if j == 0 || j == x - 1 {
                    fmt.Print("B")
                }else{
                    fmt.Print(" ")
                }
            }
            if j == x - 1 {
                fmt.Print("\n")        
            }
        }
    }

