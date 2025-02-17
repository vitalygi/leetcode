package main

func maximumWealth(accounts [][]int) int {
	var richestCustomerWealth int
	for i := 0; i < len(accounts); i++ {
		for j := 1; j < len(accounts[i]); j++ {
			accounts[i][j] += accounts[i][j-1]
		}
		richestCustomerWealth = max(richestCustomerWealth, accounts[i][len(accounts[i])-1])
	}
	return richestCustomerWealth
}
