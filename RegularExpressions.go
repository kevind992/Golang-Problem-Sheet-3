package main

import(

	"fmt"
	"math/rand"
	"time"
	"regexp"
	"strings"
)

func main(){

	rand.Seed(time.Now().UTC().UnixNano())
	
	var userInputs [12]string
	userInputs[0] = "People say I look like both my mother and father."
	userInputs[1] = "Father was a teacher."
	userInputs[2] = "I was my father's favourite."
	userInputs[3] = "I'm looking forward to the weekend."
	userInputs[4] = "My grandfather was French!"
	userInputs[5] = "I am Happy."
	userInputs[6] = "I'm not happy with you responses."
	userInputs[7] = "I am not sure that you understand the effect that your questions are having on me."
	userInputs[8] = "I am supposed to just take what you're saying at face value?"
	userInputs[9] = "What time is it?"
	userInputs[10] = "What date is it?"
	userInputs[11] = "I want to go home."

	for i:=0;i<12;i++{
		response := ElizaResponce(userInputs[i])
		fmt.Println("Eliza: ",response)
		fmt.Println()
	}
}
func ElizaResponce(input string) string{
	
	// List of Response Reflections.
	reflections := [][]string{
		{"your", "my"},
		{"you’re", "i am"},
		{"I", "you"},
		{"you", "I"},
		{"me", "you"},
		{`you`, `me`},
		{`my`, `your`},
	}

	var response [3]string
	response[0] = "I'm not sure what you're trying to say. Could you explain it to me?"
	response[1] = "How does that make you feel?"
	response[2] = "Why do you say that?"

	fmt.Println("User: ",input)

	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don't you tell me more about your Father?"
	}

	re1 := regexp.MustCompile(`(?i)i(?:'|\sa)?m (.*)`)
	reflect := re1.FindStringSubmatch(input)
	
	if len(reflect) > 0 {

		split := strings.Split(reflect[1], " ")
		
		// Loop through each word, reflecting it if there is a match.
		for i, check := range split {
			for _, reflection := range reflections {
				if matched, _ := regexp.MatchString(reflection[0], check); matched {
					split[i] = reflection[1]
					break
				}
			}
		}

		//Join the string back together
		join := strings.Join(split, " ")

		return fmt.Sprintf("How do you know you are %s?", join)
	}


	//Question 6 - 3 new user inputs.
	if matched, _ := regexp.MatchString(`(?i).*\bdate\b.*`, input); matched {
		current_time := time.Now().Local()
		current_date := current_time.Format("01-02-2006")

		return fmt.Sprintf("The Current date is %s",current_date) 
	}
	if matched, _ := regexp.MatchString(`(?i).*\btime\b.*`, input); matched {

		current_day := time.Now()
		return fmt.Sprintf("The Current time is %s ", current_day ) 
	}

	re1 = regexp.MustCompile(`(?im)^\s*I want ([^\.!]*)[\.!]*\s*$`)
	if re1.MatchString(input){
		return re1.ReplaceAllString(input, fmt.Sprintf("Why do you want $1?"))
	}

	ranNum := rand.Intn(3)

	return response[ranNum]
		
}