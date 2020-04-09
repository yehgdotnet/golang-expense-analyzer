Why the script was developed:
	1) All expense analysis apps require force you to add each expense and define each expense into defined category. 
As there are many merchant names as and when you spend money, this make you waste time on - Right click > Set Category for each expense item. 
	2) Not tracking expenses will cause you to spend more and will make you remorseful in times of actual need as business nowadays make you easier and easier to spend money.


How to use:

Customise the codes to suit your needs. 

Basically folder structures:

	Folder: /budgets/ 
	Set your budgets in each files where overall.txt is the one you set the max money out. 

	Folder: /categories 
Set your usual spending merchants with regular expression syntax, separated by | characters

	Folder: /expenses/
	recurringexpenses.csv: Set your recurring expenses.
	other-bank-cards.csv:  CSV from other bank cards to include for analysis


Pre-requisites:
1. Install GoLang
2. TransactionHistory in CSV format which must contain at least one comma and amount spent (The format is same as your bank - you can login to your banking app and download it):

   Expense title,-amount.00
   e.g 
   Mellower Coffee,-5.00


How to execute:
In Terminal, run

% go build main.go
% ./main ./records/2019-12.csv [transaction type to show details = 3-letter word - gro,din,med,uti,onl,edu,tra,loa]

Notes:
You can store all transaction records in records/ folder for reference purposes. 
 

Knowledge Reference:
https://www.youneedabudget.com/



Cheers
Myo 