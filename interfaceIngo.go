package main
import ("fmt")

type SalaryCalculator interface{
	CalculateSalary() int
}

type Permanent struct{
	empId int
	basicpay int
	pf int
}

type Contract struct {
	empId int
	basicpay int
}

type FreeLance struct{
	empId int
	basicpay int
	bonusPerfoment int

}

func (f FreeLance) CalculateSalary() int{
	return f.basicpay + f.bonusPerfoment
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}	

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator){
	expense := 0
	for _, v :=range s{
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("total expense is: %d", expense)
}

func main(){
	pemp1 := Permanent{1,200,300}
	pemp2 := Permanent{2,300, 500}
	cemp1 := Contract{3, 500}
	femp1 := FreeLance{4, 600, 400}
	employess := []SalaryCalculator{pemp1, pemp2, cemp1, femp1}
	totalExpense(employess)
}