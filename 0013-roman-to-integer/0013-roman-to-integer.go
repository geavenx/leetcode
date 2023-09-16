package main

func romanToInt(s string) int {
	romanToInt := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

    finalResult := 0
    previousValue := 0

    for i := len(s) - 1; i >= 0; i-- {
        currentValue := romanToInt[s[i]]

        if currentValue < previousValue {
           finalResult -= currentValue 
        } else {
            finalResult += currentValue
        }
        previousValue = currentValue
    }
    return finalResult
}
