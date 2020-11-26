package front

import (
	"mindset.go/abacus/app/models"
	"mindset.go/abacus/app/tools/debug"
	"github.com/labstack/echo"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	_ "github.com/stripe/stripe-go/paymentintent"
	"net/http"
	"os"
	"time"
)

func OrderController(c echo.Context) error {
	debug.Log("front<OrderController>")

	db := models.GetDataBase()
	defer db.Close()
	debug.Log("<%v : db instance close\n", db)

	var site models.Site
	for _, s := range _sites {
		if s.Host == c.Request().Host {
			site = s
		}
	}

	orderType := c.FormValue("order_type")
	subscriptionItem := models.GetSubscriptionItem(orderType)

	student := models.StudentSession(db, c)
	if student.ID == 0 {
		return c.Redirect(http.StatusFound, "/connexion/")
	}

	stripe.Key = os.Getenv("STRIPE_SK")
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(subscriptionItem.Price * 100)),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
	}
	intent, err := paymentintent.New(params)
	if err != nil {
		println(err)
	}

	return c.Render(http.StatusOK, site.Host+"abacusfingers.local<order>", map[string]interface{}{
		"title":             "Commande",
		"student":           student,
		"stripe_sk":         os.Getenv("STRIPE_SK"),
		"stripe_pk":         os.Getenv("STRIPE_PK"),
		"client_secret":     intent.ClientSecret,
		"subscription_item": subscriptionItem,
	})
}

func POrderController(c echo.Context) error {
	debug.Log("front<POrderController>")
	db := models.GetDataBase()
	defer db.Close()
	debug.Log("<%v : db instance close\n", db)

	student := models.StudentSession(db, c)
	if student.ID == 0 {
		return c.Redirect(http.StatusFound, "/connexion/")
	}

	orderType := c.FormValue("order_type")
	subscriptionItem := models.GetSubscriptionItem(orderType)

	order := models.Order{
		StudentRefer:     student.ID,
		PaymentMethod:    "Visa",
		Price:            subscriptionItem.Price,
		SubscriptionItem: subscriptionItem.Type,
		CreatedAt:        time.Time{},
	}

	db.Save(&order)

	student.AddLicenceDays(db, int(subscriptionItem.DayDuration))

	return c.Redirect(http.StatusFound, "/mon-compte/")
}
