package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	msg := ""
    msg += strings.Repeat("*", numStarsPerLine) + "\n"
    msg += welcomeMsg + "\n"
    msg += strings.Repeat("*", numStarsPerLine) 
    return msg
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg := strings.ReplaceAll(oldMsg, "*", "")
    msg = strings.TrimSpace(msg)
    return msg
}
