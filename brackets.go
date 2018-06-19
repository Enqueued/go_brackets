package main
import (
    "fmt"
)

var DefaultValue int = -1024

type fighter struct{
    L *fighter
    R *fighter
    Wins int
    Name string
    Empty bool
    Visited bool
}

func Insert (person *fighter, current *fighter){
    if (current == nil){
        return nil
    }

    switch{

    case person.Empty == true:
        current.R = &person
        current.L = &person
        return nil
    
    case person.Wins > current.Wins:
        if person.Wins == nil{
            current.R = &person
            return nil
        }
        return Insert(person, current.R)

    case person.Wins < current.Wins:
        if person.Wins == nil{
            current.L = &person
            return nil
        }
        return Insert(person, current.L)
    }
}

func (current *fighter) find(bruh string)(string, bool){
    if current == nil{
        return "", false
    }

    switch{
        case bruh == current.Name:
            return current, true
        case current.Visited != true
        case current.Empty == true:
            return current.L.find(s)
        default:
            return current.R.find(s)

    }
}

func main(){
    printf("Hello World~!")

}
