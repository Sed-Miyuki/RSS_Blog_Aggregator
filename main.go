package main

import (
	"log"
	"os"

	"github.com/Sed-Miyuki/RSS_Blog_Aggregator/internal/config"
	"github.com/Sed-Miyuki/RSS_Blog_Aggregator/internal/handlers"
)

var commands config.Commands
 
func main(){
	cfg,err:=config.Read()
	if err!=nil{
		log.Fatal(err)
	}
	state:=config.State{
		Config:&cfg,
	}
	commands=config.Commands{
		Commands:map[string]func(*config.State, config.Command) error{},
	}
	commands.Register("login",handlers.HandleLogin)
	if len(os.Args)<2{
		log.Fatal("usage: boot-dev-blog-aggregator <command> [args...]")
		return
	}
	name:=os.Args[1]
	args:=os.Args[2:]
	err=commands.Run(&state,config.Command{Name:name,Args:args})
	if err!=nil{
		log.Fatal(err)
	}
}