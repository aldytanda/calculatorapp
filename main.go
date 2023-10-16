package main

import (
	"bufio"
	"calculatorapp/calculator"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	calc := calculator.NewCalculator(0)

	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = runCmd(cmdString, calc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runCmd(str string, calc *calculator.Calculator) error {
	cmdString := strings.TrimSuffix(str, "\n")
	arrCommandStr := strings.Fields(cmdString)

	if arrCommandStr[0] == "exit" {
		os.Exit(0)
	}

	num, err := calc.Exec(arrCommandStr)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "%.1f\n", num)
	return nil
}
