package main

import "fmt"

type project struct {
	name      string
	employers int
	budget    func(a int, b int) int
}

type salary struct {
	employer string
	wage     int
	salary   salCalc
}

type salCalc func(int, int) int

func main() {
	prj := project{
		name:      "x web page",
		employers: 6,
		budget: func(a int, b int) int {
			return a * b
		},
	}

	fmt.Println("project name : ", prj.name+" employers : ", prj.employers, " bugdet(usd) : ", prj.budget(prj.employers, 50))

	sal := salary{
		employer: "john do",
		wage:     20,
		salary: func(a int, b int) int {
			return a * b
		},
	}
	fmt.Println("employer : ", sal.employer, " salary(usd) : ", sal.salary(sal.wage, 30))
}
