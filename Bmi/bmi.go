package Bmi

import (
	"fmt"
)

func BmiCalculator(weight, height float64) {
	bmi := calculateBMI(weight, height)
	message := getMessage(bmi)

	fmt.Printf("Right now your BMI is %.2f\n", bmi)
	fmt.Println(message)
}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func getMessage(bmi float64) string {
	if bmi < 18.5 {
		return "You are underweight, add more potato to the broth"
	} else if bmi < 25 {
		return "You have a normal weight, I have healthy envy of you"
	} else {
		return "You are overweight, I know, the pandemic has affected us all"
	}
}
