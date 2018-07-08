package main

import(
	"../rt"
	"fmt"
	"../instructions/base"
	"../instructions"
	"../rt/heap"
)

func interpret(method *heap.Method ,logInst bool) {
	thread := rt.NewThread()
	frame := thread.NewFrame(method)
	fmt.Println("写入栈：",method.Name())
	thread.PushFrame(frame)
	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rt.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rt.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()

		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		//fmt.Println("Method Code:" ,frame.Method(),":" , pc )
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		if (logInst) {
			logInstruction(frame, inst)
		}
		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logFrames(thread *rt.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func logInstruction(frame *rt.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}