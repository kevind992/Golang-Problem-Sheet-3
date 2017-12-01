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
	
	var userInputs [9]string
	userInputs[0] = "People say I look like both my mother and father."
	userInputs[1] = "Father was a teacher."
	userInputs[2] = "I was my father's favourite."
	userInputs[3] = "I'm looking forward to the weekend."
	userInputs[4] = "My grandfather was French!"
	userInputs[5] = "I am Happy."
	userInputs[6] = "I'm not happy with you responses."
	userInputs[7] = "I am not sure that you understand the effect that your questions are having on me."
	userInputs[8] = "I am supposed to just take what you're saying at face value?"

	for i:=0;i<9;i++{
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

	ranNum := rand.Intn(3)

	return response[ranNum]
		
}