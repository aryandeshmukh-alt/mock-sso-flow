package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Token struct {
	IDToken     string
	AccessToken string
	Claims      map[string]string
	Expiry      time.Time
	Scopes      []string
}

var authCode string
var authCodeUsed bool
var state string
var issuedToken Token

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Mock SSO Authorization Code Flow (Enhanced) ===")

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Start SSO Login")
		fmt.Println("2. Exchange Auth Code for Token")
		fmt.Println("3. Verify Token")
		fmt.Println("4. Logout (Revoke Token)")
		fmt.Println("5. Exit")

		fmt.Print("Enter choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {

		case "1":
			startSSO(reader)

		case "2":
			exchangeToken(reader)

		case "3":
			verifyToken()

		case "4":
			logout()

		case "5":
			fmt.Println("Exiting SSO simulation.")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

// ---------- SSO Steps ----------

func startSSO(reader *bufio.Reader) {
	fmt.Println("\nRedirecting to Identity Provider...")

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	_, _ = reader.ReadString('\n') // ignored in mock

	state = "STATE123"
	authCode = "XYZ123"
	authCodeUsed = false

	fmt.Println("Login successful for user:", username)
	fmt.Println("State:", state)
	fmt.Println("Auth Code:", authCode)
}

func exchangeToken(reader *bufio.Reader) {
	if authCode == "" {
		fmt.Println("No auth code found. Please login first.")
		return
	}

	if authCodeUsed {
		fmt.Println("Auth code already used")
		return
	}

	fmt.Print("Enter auth code: ")
	code, _ := reader.ReadString('\n')
	code = strings.TrimSpace(code)

	fmt.Print("Enter state: ")
	inputState, _ := reader.ReadString('\n')
	inputState = strings.TrimSpace(inputState)

	if code != authCode || inputState != state {
		fmt.Println("Invalid auth code or state")
		return
	}

	issuedToken = Token{
		IDToken:     "abc.id.sig",
		AccessToken: "xyz.access.sig",
		Expiry:      time.Now().Add(30 * time.Second),
		Scopes:      []string{"profile", "email"},
		Claims: map[string]string{
			"sub":  "user123",
			"iss":  "mock-idp",
			"role": "user",
		},
	}

	authCodeUsed = true

	fmt.Println("Token issued successfully")
	fmt.Println("ID Token:", issuedToken.IDToken)
	fmt.Println("Access Token:", issuedToken.AccessToken)
	fmt.Println("Scopes:", issuedToken.Scopes)
	fmt.Println("Expires at:", issuedToken.Expiry.Format(time.RFC822))
}

func verifyToken() {
	if issuedToken.IDToken == "" || issuedToken.AccessToken == "" {
		fmt.Println("Invalid Token: Empty token")
		return
	}

	if time.Now().After(issuedToken.Expiry) {
		fmt.Println("Invalid Token: Token expired")
		return
	}

	if !strings.Contains(issuedToken.IDToken, ".") ||
		!strings.Contains(issuedToken.AccessToken, ".") {
		fmt.Println("Invalid Token: Bad structure")
		return
	}

	requiredClaims := []string{"sub", "iss", "role"}
	for _, claim := range requiredClaims {
		if issuedToken.Claims[claim] == "" {
			fmt.Println("Invalid Token: Missing claim", claim)
			return
		}
	}

	fmt.Println("Token Verified")
	fmt.Println("Claims:")
	for k, v := range issuedToken.Claims {
		fmt.Printf("- %s: %s\n", k, v)
	}
}

func logout() {
	issuedToken = Token{}
	authCode = ""
	authCodeUsed = false
	state = ""

	fmt.Println("Logged out successfully. Token revoked.")
}
