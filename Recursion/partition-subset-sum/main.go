// class Solution:
//     def canPartition(self, nums: List[int]) -> bool:
//         # Calculate the total sum of the input array
//         total_sum = sum(nums)
//
//         # If the total sum is odd, we cannot partition the array into two equal sum subsets
//         if total_sum % 2 == 1:
//             return False
//
//         # Calculate the target sum for each subset
//         target_sum = total_sum // 2
//
//         # Initialize a boolean list of size target_sum+1 to keep track of whether a sum can be formed using the input array
//         dp = [False] * (target_sum+1)
//
//         # We can always form a sum of 0 using the input array
//         dp[0] = True
//
//         # Loop through each element in the input array
//         for num in nums:
//             # Starting from the target sum, loop backwards through the dp list
//             for j in range(target_sum, num-1, -1):
//                 # If we can form a sum j-num using the previous elements in the input array,
//                 # we can also form a sum j using the current element
//                 dp[j] = dp[j] or dp[j-num]
//
//         # Return whether or not we can form a sum of target_sum using the input array
//         return dp[target_sum]

package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, i := range nums {
		sum += i
	}

	if sum%2 == 1 {
		return false
	}

	ts := sum / 2
	dp := make([]bool, ts+1)
	dp[0] = true

	for _, n := range nums {
		for j := ts; j > n-1; j-- {
			dp[j] = dp[j] || dp[j-n]
		}

		fmt.Println(dp)
	}
	fmt.Println("...............")
	fmt.Println(dp)
	return dp[ts]
}

func main() {
	n := []int{0, 1, 4, 5}
	fmt.Println(canPartition(n))
}
