package main

import neweb_pay "neweb-pay"

// @title Line server
// @version 1.0
// @description
// @contact.name sherryMiet
// @host localhost:9002
// @schemes http,https
func main() {
	client := neweb_pay.NewClient("MS140665845", "LvMrtfcXgXLMENkkjo2Gn17p3NUvDmDd", "CH1wnkc5ZGGXpkpP")
	Info := neweb_pay.NewMPGGateWayTradeInfo()
	Info.WithOptional(neweb_pay.OptionValue{
		Email:       "sherry2000307@gmail.com",
		EmailModify: 0,
		LoginType:   0,
		Version:     "1.5",
	}).
		CreateOrder("MS140665845", 10, "test")

	html := client.MPGGateway(Info).DoTest()
	print(html)
}
