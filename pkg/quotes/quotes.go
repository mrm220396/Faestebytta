package quotes

import "math/rand"

var quoteList = []string{
	"What’s the difference between an outlaw and an in-law? Outlaws are wanted.",
	"What do you call a sleeping bull? A Bulldozer",
	"Did you hear they arrested the devil? Yeah, they got him on possession.",
	"What did one DNA say to the other DNA? “Do these genes make me look fat?”",
	"My IQ test results came back. They were negative.",
	"What do you get when you cross a polar bear with a seal? A polar bear.",
	"Why can’t you trust an atom? Because they make up literally everything.",
	"Why was six afraid of seven? Because seven eight nine.",
	"What do you call a hippie’s wife? Mississippi.",
	"What’s the difference between an outlaw and an in-law? Outlaws are wanted.",
	"Scientists have recently discovered a food that greatly reduces sex drive. It’s called wedding cake.",
	"Before you marry a person, you should first make them use a computer with a slow Internet connection to see who they really are.",
	"I never knew what happiness was until I got married—and then it was too late.",
	"Some men say they don’t wear their wedding band because it cuts off circulation. Well, that’s the point, isn’t it?",
	"Advice to husbands: Try praising your wife now and then, even if it does startle her at first.",
	"Here's not a democracy!!",
	"What did the toaster say to the slice of bread? I want you inside me.",
	"Give it to me! Give it to me! she yelled. I'm so wet, give it to me now! She could scream all she wanted, but I was keeping the umbrella.",
}

// GetQuote returns a random quote from the list of quotes
func GetQuote() string {
	return quoteList[rand.Intn(len(quoteList))]
}
