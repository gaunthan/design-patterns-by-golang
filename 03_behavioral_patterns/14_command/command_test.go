package command

func ExampleCommand() {
	key := "uid"
	commands := []Command{
		NewSetIntCommand(key, 123456),
		NewPrintValueCommand(key),
	}

	runCommands(commands...)

	// Output:
	// 123456
}

func runCommands(commands ...Command) {
	for _, c := range commands {
		c.Execute()
	}
}
