package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"github.com/olekukonko/tablewriter"

)
var (
	fileCurrencyFormat string = "./config/currencyformat.txt"
	fileOtherBanksCard string = "./expenses/other-bank-cards.csv"
	fileRecurringExpenses string = "./expenses/recurringexpenses.csv"


	drinksTotalCost float64 = 0.00
	drinksBudget float64 = 0.00
	drinksRegex string = ""

	groceryTotalCost float64 = 0.00
	groceryBudget float64 = 0.00
	groceryRegex string = ""


	clothsTotalCost float64 = 0.00
	clothsBudget float64 = 0.00
	clothsRegex string = ""


	heaTotalCost float64 = 0.00
	heaBudget float64 = 0.00
	heaRegex string = ""


	onlineSubscriptionTotalCost float64 = 0.00
	onlineSubscriptionBudget float64 = 0.00
	onlineRegex string = ""


	eduTotalCost float64 = 0.00
	eduBudget float64 = 0.00
	eduRegex string = ""

	transportTotalCost float64 = 0.00
	transportBudget float64 = 0.00
	transportRegex string = ""


	utilityTotalCost float64 = 0.00
	utilityBudget float64 = 0.00
	utilityRegex string = ""

	loanTotalCost float64 = 0.00
	loanBudget float64 = 0.00
	loanRegex string = ""

	othersTotalCost float64 = 0.00
	othersBudget float64 = 0.00
	othersRegex string = ""


	insTotalCost float64 = 0.00
	insBudget float64 = 0.00
	insRegex string = ""

	liaTotalCost float64 = 0.00
	liaBudget float64 = 0.00
	liaRegex string = ""

	famTotalCost float64 = 0.00
	famBudget float64 = 0.00
	famRegex string = ""


	overallBudget float64 = 0.00
	allTotalCost float64 = 0.00
	essentialTotalCost float64 = 0.00
	nonEssentialTotalCost float64 = 0.00


	showTransType string = ""
	transType string = ""
	currencyFormat string = ""

)

func printTotal(){


	allTotalCost =
		othersTotalCost+
		clothsTotalCost+
		groceryTotalCost+
		utilityTotalCost+
		transportTotalCost+
		onlineSubscriptionTotalCost+
		heaTotalCost+
		drinksTotalCost+
		eduTotalCost+
		loanTotalCost+
		insTotalCost+
		liaTotalCost+
		famTotalCost


	essentialTotalCost =
		othersTotalCost+
		utilityTotalCost+
		transportTotalCost+
		onlineSubscriptionTotalCost+
		heaTotalCost+
		eduTotalCost+
		insTotalCost+
		liaTotalCost+
		famTotalCost

	nonEssentialTotalCost = clothsTotalCost+ groceryTotalCost+ drinksTotalCost+loanTotalCost



	fmt.Println("")

	data := [][]string{
		[]string{"Family/Parents",fmt.Sprintf("%0.2f",famTotalCost),fmt.Sprintf("%0.2f",famBudget-famTotalCost) },
		[]string{"Liabilities",fmt.Sprintf("%0.2f",liaTotalCost),fmt.Sprintf("%0.2f",liaBudget-liaTotalCost) },
		[]string{"Medicine/Health", fmt.Sprintf("%0.2f",heaTotalCost),fmt.Sprintf("%0.2f",heaBudget-heaTotalCost) },
		[]string{"Education", fmt.Sprintf("%0.2f",eduTotalCost),fmt.Sprintf("%0.2f",eduBudget-eduTotalCost)},

		[]string{"Insurance",fmt.Sprintf("%0.2f",insTotalCost),fmt.Sprintf("%0.2f",insBudget-insTotalCost) },

		[]string{"Dining/Others",fmt.Sprintf("%0.2f",othersTotalCost),fmt.Sprintf("%0.2f",othersBudget-othersTotalCost) },

		[]string{"Utility", fmt.Sprintf("%0.2f",utilityTotalCost),fmt.Sprintf("%0.2f",utilityBudget-utilityTotalCost)},

		[]string{"Online", fmt.Sprintf("%0.2f",onlineSubscriptionTotalCost),fmt.Sprintf("%0.2f",onlineSubscriptionBudget-onlineSubscriptionTotalCost)},
		[]string{"Transport", fmt.Sprintf("%0.2f",transportTotalCost),fmt.Sprintf("%0.2f",transportBudget-transportTotalCost)},


		[]string{"Loan",fmt.Sprintf("%0.2f",loanTotalCost),fmt.Sprintf("%0.2f",loanBudget-loanTotalCost) },


		[]string{"Coffee", fmt.Sprintf("%0.2f",drinksTotalCost), fmt.Sprintf("%0.2f",drinksBudget-drinksTotalCost)},
		[]string{"Grocery", fmt.Sprintf("%0.2f",groceryTotalCost),fmt.Sprintf("%0.2f",groceryBudget-groceryTotalCost)},
		[]string{"Cloths", fmt.Sprintf("%0.2f",clothsTotalCost),fmt.Sprintf("%0.2f",clothsBudget-clothsTotalCost)},

		[]string{"--------------","---------","----------" },
		[]string{"Total", fmt.Sprintf("%0.2f",allTotalCost),fmt.Sprintf("%0.2f",overallBudget-allTotalCost)},
		[]string{"--------------","---------","----------" },
		[]string{"Essential Total", fmt.Sprintf("%0.2f",essentialTotalCost),"-"},
		[]string{"--------------","---------","----------" },
		[]string{"Non-Essential Total", fmt.Sprintf("%0.2f",nonEssentialTotalCost),"-"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Area", "Spent", "Remaining"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output

	fmt.Println("")



	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)



	cyan := color.New(color.FgCyan)
	boldCyan := cyan.Add(color.Bold)

	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold)


	if drinksTotalCost > drinksBudget {
		boldRed.Println("[WARNING] Coffee/Tea is over budget: ",fmt.Sprintf("%0.2f",drinksBudget-drinksTotalCost))
	}
	if heaTotalCost > heaBudget {
		boldRed.Println("[WARNING] Medicine/Health is over budget: ",fmt.Sprintf("%0.2f",heaBudget-heaTotalCost))
	}
	if onlineSubscriptionTotalCost > onlineSubscriptionBudget {
		boldRed.Println("[WARNING] Online Spend is over budget: ",fmt.Sprintf("%0.2f",onlineSubscriptionBudget-onlineSubscriptionTotalCost))
	}
	if eduTotalCost > eduBudget {
		boldRed.Println("[WARNING] Education is over budget: ",fmt.Sprintf("%0.2f",eduBudget-eduTotalCost))
	}
	if transportTotalCost > transportBudget {
		boldRed.Println("[WARNING] Transport is over budget: ",fmt.Sprintf("%0.2f",transportBudget-transportTotalCost))
	}
	if utilityTotalCost > utilityBudget {
		boldRed.Println("[WARNING] Utility is over budget: ",fmt.Sprintf("%0.2f",utilityBudget-utilityTotalCost))
	}
	if loanTotalCost> loanBudget {
		boldRed.Println("[WARNING] Loan/Instalment is over budget: ",fmt.Sprintf("%0.2f",loanBudget-loanTotalCost))
	}
	if clothsTotalCost > clothsBudget {
		boldRed.Println("[WARNING] Cloths is over budget: ",fmt.Sprintf("%0.2f",clothsBudget-clothsTotalCost))
	}
	if othersTotalCost > othersBudget{
		boldRed.Println("[WARNING] Dining/Others is over budget: ",fmt.Sprintf("%0.2f",othersBudget-othersTotalCost))
	}
	if liaTotalCost > liaBudget{
		boldRed.Println("[WARNING] Liability is over budget: ",fmt.Sprintf("%0.2f",liaBudget-liaTotalCost))
	}
	if famTotalCost > famBudget{
		boldRed.Println("[WARNING] Family/Parent is over budget: ",fmt.Sprintf("%0.2f",famBudget-famTotalCost))
	}
	fmt.Println("")

	if allTotalCost> overallBudget {
		boldCyan.Println("[WARNING] Overall Budget is over allocated amount: ",currencyFormat,fmt.Sprintf("%0.2f out of %s %0.2f",overallBudget-allTotalCost,currencyFormat,overallBudget))
		fmt.Println("")
	}else{
		boldGreen.Println("[OK] Overall Budget is within control.\n     You've left with ",currencyFormat,fmt.Sprintf("%0.2f out of %s %0.2f",overallBudget-allTotalCost,currencyFormat,overallBudget))
		fmt.Println("")
	}



}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initializeVariables(){

	// current format
	currencyF,_ := ioutil.ReadFile(fileCurrencyFormat)
	currencyFormat = string(currencyF)


	// Initializing budgets
	driB,_ := ioutil.ReadFile("./budgets/drinks.txt")
	drinksBudget, _ = strconv.ParseFloat(string(driB),64)

	groB,_ := ioutil.ReadFile("./budgets/grocery.txt")
	groceryBudget, _ = strconv.ParseFloat(string(groB),64)

	cloR,_ := ioutil.ReadFile("./budgets/cloths.txt")
	clothsBudget, _ = strconv.ParseFloat(string(cloR),64)

	medB,_ := ioutil.ReadFile("./budgets/health.txt")
	heaBudget, _ = strconv.ParseFloat(string(medB),64)

	traB,_ := ioutil.ReadFile("./budgets/transport.txt")
	transportBudget, _ = strconv.ParseFloat(string(traB),64)

	utiB,_ := ioutil.ReadFile("./budgets/utility.txt")
	utilityBudget, _ = strconv.ParseFloat(string(utiB),64)

	loaB,_ := ioutil.ReadFile("./budgets/loans.txt")
	loanBudget, _ = strconv.ParseFloat(string(loaB),64)

	onlB,_ := ioutil.ReadFile("./budgets/online.txt")
	onlineSubscriptionBudget, _ = strconv.ParseFloat(string(onlB),64)

	eduB,_ := ioutil.ReadFile("./budgets/education.txt")
	eduBudget, _ = strconv.ParseFloat(string(eduB),64)

	othB,_ := ioutil.ReadFile("./budgets/others.txt")
	othersBudget, _ = strconv.ParseFloat(string(othB),64)

	ovaB,_ := ioutil.ReadFile("./budgets/overall.txt")
	overallBudget, _ = strconv.ParseFloat(string(ovaB),64)

	insB,_ := ioutil.ReadFile("./budgets/insurance.txt")
	insBudget, _ = strconv.ParseFloat(string(insB),64)

	liaB,_ := ioutil.ReadFile("./budgets/liability.txt")
	liaBudget, _ = strconv.ParseFloat(string(liaB),64)

	famB,_ := ioutil.ReadFile("./budgets/family.txt")
	famBudget, _ = strconv.ParseFloat(string(famB),64)

	// Initializing expense categories

	driR,_ := ioutil.ReadFile("./categories/drinks.txt")
	drinksRegex = string(driR)


	gro,_ := ioutil.ReadFile("./categories/grocery.txt")
	groceryRegex = string(gro)


	clo,_ := ioutil.ReadFile("./categories/cloths.txt")
	clothsRegex = string(clo)

	med,_ := ioutil.ReadFile("./categories/health.txt")
	heaRegex = string(med)

	tra,_ := ioutil.ReadFile("./categories/transport.txt")
	transportRegex = string(tra)

	uti,_ := ioutil.ReadFile("./categories/utility.txt")
	utilityRegex = string(uti)


	loa,_ := ioutil.ReadFile("./categories/loans.txt")
	loanRegex = string(loa)

	onl,_ := ioutil.ReadFile("./categories/online.txt")
	onlineRegex = string(onl)


	edu,_ := ioutil.ReadFile("./categories/education.txt")
	eduRegex = string(edu)

	ins,_ := ioutil.ReadFile("./categories/insurance.txt")
	insRegex = string(ins)

	lia,_ := ioutil.ReadFile("./categories/liability.txt")
	liaRegex = string(lia)

	fam,_ := ioutil.ReadFile("./categories/family.txt")
	famRegex = string(fam)
}

func printLine(){
	fmt.Println("--------------------------------------------------------")
}

func classify(expenseLine string){


	expenseCost := expenseLine[strings.Index(expenseLine,",-")+2:len(expenseLine)]
	expenseCostFloat, _ := strconv.ParseFloat(expenseCost,2)

	drinksMatched, _ := regexp.MatchString(drinksRegex, expenseLine)


	groceryMatched, _ := regexp.MatchString(groceryRegex, expenseLine)


	clothsMatched, _ := regexp.MatchString(clothsRegex, expenseLine)


	heaMatched, _ := regexp.MatchString(heaRegex, expenseLine)


	transportMatched, _ := regexp.MatchString(transportRegex, expenseLine)


	utilityMatched, _ := regexp.MatchString(utilityRegex, expenseLine)


	loanMatched, _ := regexp.MatchString(loanRegex, expenseLine)


	onlineSubscriptionMatched, _ := regexp.MatchString(onlineRegex, expenseLine)


	eduMatched, _ := regexp.MatchString(eduRegex, expenseLine)

	insMatched, _ := regexp.MatchString(insRegex, expenseLine)

	liaMatched, _ := regexp.MatchString(liaRegex, expenseLine)

	famMatched, _ := regexp.MatchString(famRegex, expenseLine)


	if drinksMatched == true {
		if showTransType == "dri"{
			fmt.Println(expenseLine)
		}

		drinksTotalCost = drinksTotalCost + expenseCostFloat

	}else if groceryMatched == true {
		if showTransType == "gro"{
			fmt.Println(expenseLine)
		}

		groceryTotalCost = groceryTotalCost  + expenseCostFloat
	}else if heaMatched == true {
		if showTransType == "med"{
			fmt.Println(expenseLine)
		}
		heaTotalCost = heaTotalCost  + expenseCostFloat
	}else if transportMatched == true {
		if showTransType == "tra"{
			fmt.Println(expenseLine)
		}
		transportTotalCost = transportTotalCost + expenseCostFloat
	}else if onlineSubscriptionMatched == true {
		if showTransType == "onl"{
			fmt.Println(expenseLine)
		}
		onlineSubscriptionTotalCost = onlineSubscriptionTotalCost  + expenseCostFloat
	}else if utilityMatched == true {
		if showTransType == "uti"{
			fmt.Println(expenseLine)
		}
		utilityTotalCost = utilityTotalCost + expenseCostFloat
	}else if clothsMatched == true {
		if showTransType == "clo"{
			fmt.Println(expenseLine)
		}
		clothsTotalCost = clothsTotalCost  + expenseCostFloat
	}else if eduMatched == true {
		if showTransType == "edu"{
			fmt.Println(expenseLine)
		}
		eduTotalCost = eduTotalCost  + expenseCostFloat
	}else if loanMatched == true {
		if showTransType == "loa"{
			fmt.Println(expenseLine)
		}
		loanTotalCost = loanTotalCost  + expenseCostFloat
	}else if insMatched== true {
		if showTransType == "ins"{
			fmt.Println(expenseLine)
		}
		insTotalCost = insTotalCost + expenseCostFloat
	}else if liaMatched== true {
		if showTransType == "lia"{
			fmt.Println(expenseLine)
		}
		liaTotalCost = liaTotalCost + expenseCostFloat
	}else if famMatched== true {
		if showTransType == "fam"{
			fmt.Println(expenseLine)
		}
		famTotalCost = famTotalCost + expenseCostFloat
	}else{
		if showTransType == "din"{
			fmt.Println(expenseLine)
		}
		othersTotalCost = othersTotalCost+ expenseCostFloat
	}


}

func main() {


	fmt.Println("")

	if len(os.Args) < 2{
		fmt.Println("go run main.go path/to/transactionfile [3-letter transaction type]")
		os.Exit(0)
	}
	transFile := os.Args[1]


	if len(os.Args) == 3{
		showTransType = os.Args[2]

		switch showTransType{
		case "dri":
			transType = "Coffee"
		case "gro":
			transType = "Grocery"
		case "din":
			transType = "Dining"
		case "med":
			transType = "Health"
		case "uti":
			transType = "Utility"
		case "onl":
			transType = "Online"
		case "clo":
			transType = "Cloths"
		case "edu":
			transType = "Education"
		case "tra":
			transType = "Transport"
		case "loa":
			transType = "Loan/Instalment"
		case "lia":
			transType = "Liabilities"
		case "fam":
			transType = "Family/Parents"

		}
		fmt.Println(fmt.Sprintf("Detailed Costs for %s",transType))
		printLine()
	}

	initializeVariables()

	////////////////// USER PROVIDED FILE /////////////////////////////////////////////
	file, err := os.Open(transFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expenseLine := strings.ToLower(scanner.Text())
		if strings.Contains(expenseLine, ",-") {
			classify(expenseLine)

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	////////////////// OTHER BANK CARDS /////////////////////////////////////////////

	file, err = os.Open(fileOtherBanksCard)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		expenseLine := strings.ToLower(scanner.Text())
		if strings.Contains(expenseLine, ",-") {
			classify(expenseLine)

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	////////////////// RECURRING EXPENSE FILES FROM OTHER BANKS/////////////////////////////////////////////
	file, err = os.Open(fileRecurringExpenses)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		expenseLine := strings.ToLower(scanner.Text())
		if strings.Contains(expenseLine, ",-") {
			classify(expenseLine)

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	printTotal()
}