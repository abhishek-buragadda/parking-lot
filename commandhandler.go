package  main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"parkinglot/src/commands"
	"parkinglot/src/types"
	"strings"
)
func HandleFileScanner(){
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		handleCommand(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func HandleInterativeCommands(){
	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cmdString = strings.TrimSuffix(cmdString, "\n")
		if cmdString == "exit" {
			os.Exit(1)
		}
		handleCommand(cmdString)
	}
}


func handleCommand( command string  ){
	//println(command)
	parkingLot := types.GetParkingLotInstance()
	commandSplit := strings.Split(command, " ")
	parkingCommand := commandSplit[0]
	commands.GetCommandByName(parkingCommand).Execute(parkingLot , commandSplit[1:])
}