package questions

import "fmt"

func AskQuestions() {
	questionsArray := [5][2]string{
		{"What is the capital of France? \nA) Paris \nB) Rome \nC) Madrid", "A"},
		{"Which planet is known as the Red Planet? \nA) Venus \nB) Mars \nC) Jupiter", "B"},
		{"Who wrote 'Romeo and Juliet'? \nA) Mark Twain \nB) Charles Dickens \nC) William Shakespeare", "C"},
		{"What is the chemical symbol for water? \nA) CO2 \nB) H2O \nC) O2", "B"},
		{"In which year did the Titanic sink? \nA) 1912 \nB) 1920 \nC) 1898", "A"},
	} //each question is an array of 2 elements; question and answer
	score := 0
	for _, pair := range questionsArray {
		fmt.Println(pair[0])
		var user_answer string
		fmt.Scan(&user_answer)
		if user_answer == pair[1] {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect! The correct answer is:", pair[1])
		}
	}
	fmt.Printf("Your final score is: %d out of %d\n", score, len(questionsArray))
	fmt.Println("Thanks for playing the quiz game!")
}
