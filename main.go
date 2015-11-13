package main

import "fmt"
import "math/rand"
import "time"
import "sort"
import "os"
import "os/exec"
import "runtime"

import "github.com/dustin/go-humanize"
import "github.com/divan/num2words"

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = clear["linux"]
	clear["windows"] = func() {
		cmd := exec.Command("cls") //Windows example it is untested, but I think its working
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		msg := fmt.Sprintf("Your platform, %s, is unsupported! I can't clear terminal screen :(", runtime.GOOS)
		panic(msg)
	}
}

func getAnswers(r1 *rand.Rand) (map[string]int, [4]int) {
	answers := make(map[string]int)
	var numbers [4]int
	for i := 0; i < 4; i++ {
		numbers[i] = r1.Intn(9999)
	}

	answers["a"] = numbers[0]
	answers["b"] = numbers[1]
	answers["c"] = numbers[2]
	answers["d"] = numbers[3]
	return answers, numbers
}

func chooseOneAndPrint(r1 *rand.Rand, numbers [4]int) int {
	// Pick the number to choose.
	key := r1.Intn(4)
	chosen := numbers[key]

	// Print the chosen's number
	fmt.Println(humanize.Comma(int64(chosen)))
	return chosen
}

func printChoices(answers map[string]int) {
	choices := []string{"a", "b", "c", "d"}

	sort.Strings(choices)

	// Do a multiple choice with words listed and A B C D
	for _, choice := range choices {
		text := num2words.Convert(answers[choice])
		fmt.Printf("%s.) %s \n", choice, text)
	}

	fmt.Println("q.) Quit")
}

func collectAnswer(answers map[string]int) int {
	// Collect Input from user
	ok := false

	var answer = make([]byte, 1)

	for ok == false {
		fmt.Print("Answer: ")
		os.Stdin.Read(answer)
		switch string(answer) {
		case "a":
			fallthrough
		case "b":
			fallthrough
		case "c":
			fallthrough
		case "d":
			ok = true
		case "q":
			os.Exit(0)
		}
	}

	return answers[string(answer)]
}

func play() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	answers, numbers := getAnswers(r1)
	chosen := chooseOneAndPrint(r1, numbers)
	printChoices(answers)
	answer := collectAnswer(answers)

	clearScreen()
	if chosen == answer {
		fmt.Println("Correct!")
		return 1
	}

	fmt.Println("Sorry... let's try another one.")
	return 0
}

func main() {
	clearScreen()

	score := 0
	for {
		fmt.Println("Current score: ", score)
		score += play()
	}
}
