package helpers

import (
	"errors"
	"fmt"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	ErrVerifier         = errors.New("verify email failed")
	ErrEmailSyntax      = errors.New("email address syntax is invalid")
	ErrDepositableEmail = errors.New("detect depositable email address")
	ErrNotReachable     = errors.New("email address is not reachable")
)

func VerifyEmailScheme(email string) error {
	verifier := emailverifier.NewVerifier()
	verifier = verifier.EnableDomainSuggest()
	// verifier = verifier.AddDisposableDomains([]string{"tractorjj.com"})
	// fmt.Println("verEmailPostHandler running")

	// email := c.PostForm("email")
	ret, err := verifier.Verify(email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return ErrVerifier
	}

	fmt.Println("email validation result", ret)
	fmt.Println("Email:", ret.Email, "\nReachable:", ret.Reachable, "\nSyntax:", ret.Syntax, "\nSMTP:", ret.SMTP, "\nGravatar:", ret.Gravatar, "\nSuggestion:", ret.Suggestion, "\nDisposable:", ret.Disposable, "\nRoleAccount:", ret.RoleAccount, "\nFree:", ret.Free, "\nHasMxRecords:", ret.HasMxRecords)

	// needs @ and . for starters
	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return ErrEmailSyntax
	}
	if ret.Disposable {
		fmt.Println("sorry, we do not accept disposable email addresses")
		// c.HTML(http.StatusBadRequest, "ver-email.go.html", gin.H{"message": "sorry, we do not accept disposable email addresses"})
		return ErrDepositableEmail
	}
	if ret.Suggestion != "" {
		fmt.Println("email address is not reachable, looking for ", ret.Suggestion, "instead?")
		return ErrNotReachable
	}
	// possible return string values: yes, no, unkown
	if ret.Reachable == "no" {
		fmt.Println("email address is not reachable")
		// c.HTML(http.StatusBadRequest, "ver-email.go.html", gin.H{"message": "email address was unreachable"})
		return ErrNotReachable
	}
	// check MX records so we know DNS setup properly to recieve emails
	// if !ret.HasMxRecords {
	// 	fmt.Println("domain entered not properly setup to recieve emails, MX record not found")
	// 	c.HTML(http.StatusBadRequest, "ver-email.go.html", gin.H{"message": "domain entered not properly setup to recieve emails, MX record not found"})
	// 	return
	// }
	// ... code to register user
	// c.HTML(http.StatusOK, "ver-email-result.go.html", gin.H{"email": email})
	return nil
}
