package main

import (
	"fmt"
	"strings"
	"../classpath"
	"../rt/heap"

)


func main(){
	cmd :=parseCmd()
	if cmd.versionFlag{
		fmt.Printf("version 0.0.1")
	}else if cmd.helpFlag||cmd.class==""{
		printUsage()
	}else{
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd){
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, true)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

}

