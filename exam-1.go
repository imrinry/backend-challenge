package main

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
    59           [59]
   73 41        [73,41]
  52 40 53     [52,40,53]
 26 53 6 34   [26,53,6,34]


1. row 4 (26,53,6,34):
   sums = [26,53,6,34]

2. row 3 (52,40,53):
   - position 52: select [26,53] = 53  -> 52+53 = 105
   - position 40: select [53,6] = 53   -> 40+53 = 93
   - position 53: select [6,34] = 34   -> 53+34 = 87
   sums = [105,93,87]

3. row 2 (73,41):
   - position 73: select [105,93] = 105  -> 73+105 = 178
   - position 41: select [93,87] = 93    -> 41+93 = 134
   sums = [178,134]

4. row 1 (59):
   - position 59: select [178,134] = 178 -> 59+178 = 237

ans 237 (path: 59->73->52->53)
*/

func findSum(pyramid [][]int) int {

	sums := pyramid[len(pyramid)-1]

	for row := len(pyramid) - 2; row >= 0; row-- {

		newSums := make([]int, len(pyramid[row]))

		for i := 0; i < len(pyramid[row]); i++ {

			best := bigger(sums[i], sums[i+1])

			newSums[i] = pyramid[row][i] + best
		}

		sums = newSums
	}

	return sums[0]
}
