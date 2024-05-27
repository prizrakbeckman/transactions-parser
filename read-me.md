## Steps to test
### Prerequesite step :
Please click on main in order to start the app in ing-transactions-history-servlet.go

Or after unzipping the folder, please run the following command :
go build
then start exe file generated.

After starting all containers in order to test and the main function:
1- To get a list of all transactions after a time period. (According to open api specs)
http://localhost:8081/company/transactions?after-timestamp=2022-01-01T00:01:36Z&company-iban=FI9209440025532454

2- to get the balance
http://localhost:8081/company/balance?after-timestamp=2022-01-01T00:01:36Z&company-iban=FI9209440025532454

3- to get all companies:
http://localhost:8081/companies

For the record :
- I chose go as a language since it's the one used at ing.
- I used go routines and channels, to improve performances in fetching  transactions from swift and sepa.
- Balance, is the balance of all transactions by company
- On startup currencies and companies are already fetched.
- I tried to use caching, but it didn't work on my machine as expected
- There are some transactions with exchanges rate non existing, so i excluded them, also transaction where
  valid ibans.
- Transaction struct has sender and payer, as well as receiver and beneficiary. (I wanted to use one struct, since go doesn't have a good inheritence feature.
- large db didn't work for me, so i couldn't work on peformance on it.
- I also added a unit test.