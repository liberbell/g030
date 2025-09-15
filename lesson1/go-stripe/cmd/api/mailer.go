package main

import "embed"

//go:embed templates
var emailTemplatesFS embed.FS

func (app *application) SendMail(from, to, subject, tmpl string, data interface{}) error {

	return nil
}
