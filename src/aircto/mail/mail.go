package mail

import (
	DB "aircto/model"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
	"strings"
)

var auth smtp.Auth

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func PrepareToSendMail(issteDetailsRes DB.Issue, assigneeDetails DB.User, title string, message string) {
	fmt.Println("goroutine mail start...")

	// convert issue struct data into json
	arrIsu, _ := json.Marshal(issteDetailsRes)
	jsonIssueDetails := make(map[string]interface{})
	if err := json.Unmarshal([]byte(string(arrIsu)), &jsonIssueDetails); err != nil {
		panic(err)
	}
	// prepare mail connections
	smtpServer := "smtp.gmail.com"
	auth = smtp.PlainAuth(
		"",
		"launchyard.aircto@gmail.com",
		"launchyardtest",
		smtpServer,
	)

	// prepare mail template
	templateData := map[string]interface{}{
		"Name":    assigneeDetails.FirstName + " " + assigneeDetails.LastName,
		"Message": message,
		"Issue":   jsonIssueDetails,
	}

	bodyTemplate, err := ParseTemplate(Template, templateData)

	if err == nil {
		ok, _ := SendEmail(bodyTemplate, assigneeDetails.FirstName+" "+assigneeDetails.LastName, assigneeDetails.Email, title)
		fmt.Println("Mail send to the assignee...: ", ok)
	}

	fmt.Println("goroutine mail end...")

	return
}

func SendEmail(bodyTemplate string, toName string, toEmail string, title string) (bool, error) {
	addr := "smtp.gmail.com:587"
	from := mail.Address{"LaunchYard AirCTO", "launchyard.aircto@gmail.com"}
	to := mail.Address{toName, toEmail}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(bodyTemplate))

	if err := smtp.SendMail(addr, auth, from.Address, []string{to.Address}, []byte(message)); err != nil {
		return false, err
	}
	return true, nil
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	// function map to increment $key value
	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"inc": func(i string) string {
			return i
		},
	}

	t := template.Must(template.New("result").Funcs(funcMap).Parse(templateFileName))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
