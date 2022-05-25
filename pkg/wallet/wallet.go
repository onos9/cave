package wallet

import (
	"context"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/cave/pkg/database"
	"github.com/cave/pkg/mailer"
	"github.com/cave/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type Wallet struct {
}

const DOLLER_RATE = 400
const APPLICATION_FEE = 10
const TUITION_FEE = 1000

var ctx = context.Background()

func ProcessPayment(m fiber.Map, user *models.User) error {
	rdb := database.RedisClient(0)
	defer rdb.Close()

	html := m["html"].(string)
	t, err := getTransaction(html)
	if err != nil {
		return err
	}

	if _, ok := t["Narration"]; !ok {
		return fmt.Errorf("error: Narration not found")
	}

	naration := t["Narration"]
	reg := regexp.MustCompile("[0-9]+")
	code := reg.FindAllString(naration, -1)[1]
	paymentType := reg.FindAllString(naration, -1)[0]

	if _, ok := t["Amount"]; !ok {
		return fmt.Errorf("error: Amount not found")
	}

	amount := t["Amount"]
	a := strings.Replace(amount, ",", "", -1)
	a = strings.Split(a, ".")[0]

	amt, err := strconv.Atoi(a)
	if err != nil {
		return err
	}

	id, err := rdb.Get(ctx, code).Result()
	if err != nil {
		return err
	}

	if err := user.FetchByID(id); err != nil {
		return err
	}

	wallet := user.Wallet + (amt / DOLLER_RATE)
	if paymentType == "10" {
		wallet = wallet - APPLICATION_FEE
	}
	if paymentType == "12" {
		wallet = wallet - TUITION_FEE
	}

	if wallet < 0 {
		user.Wallet = wallet
		err = processIncompletePayment(user, amt)
		if err != nil {
			return err
		}
	}

	return nil
}

func processIncompletePayment(user *models.User, paid int) error {
	balance := math.Abs(math.Inf(user.Wallet))
	mail := fiber.Map{
		"fromAddress": "admin@adullam.ng",
		"toAddress":   user.Email,
		"subject":     "Adullam|Payment Confirmation",
		"content": fiber.Map{
			"filename": "repay.html",
			"balance":  balance,
			"due":      APPLICATION_FEE,
			"paid":     paid,
		},
	}

	m := new(mailer.Mail)
	_, err := m.SendMail(mail)
	if err != nil {
		return err
	}

	return nil
}
