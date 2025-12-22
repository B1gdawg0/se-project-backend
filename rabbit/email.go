package rabbit

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	EMAIL_SENDER = "your-email@example.com"
	EMAIL_PASSWORD = "your-email-password"
)


func InitEmailWorker() {

	go func() {
		consumeMessages()
		cleanupEmailWorker()
	}()
}

func consumeMessages() {
	msgs, err := rabbitCh.Consume(
		rabbitEmailQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	LogSeriousFailed(err, "Can't start consuming messages")

	log.Println("Waiting for messages...")

	for msg := range msgs {
		handleEmailMessage(msg)
	}
}

func handleEmailMessage(msg amqp.Delivery) {
	log.Printf("Received message: %s", msg.Body)

	order := new(response.GetOrderResponse)
	err := json.Unmarshal(msg.Body, order)
	if err != nil {
		log.Printf("Failed to deserialize message body: %v", err)
	}

	from := mail.NewEmail("Drink&Groove", "lerdphipat.k@ku.th")
    subject := "Receipt for Your Purchase at Drink&Groove"
    to := mail.NewEmail("Dear Customer", "lerdphipat.k@ku.th")

    plainTextContent := "Thank you for your order with Drink&Groove! Your receipt details are provided below. Enjoy your drinks!"

    orderDetails := ""
    totalAmount := 0.0
    for _, orderline := range order.Orderline {
		val, err := strconv.ParseFloat(orderline.Quantity, 64)
		if err != nil{
			return
		}

        itemTotal := orderline.Menu.Price * val
        totalAmount += itemTotal
        orderDetails += fmt.Sprintf(`
            <tr>
                <td>%s</td>
                <td>%s</td>
                <td>%.2f THB</td>
                <td>%.2f THB</td>
            </tr>
        `, orderline.Menu.Description, orderline.Quantity, orderline.Menu.Price, itemTotal)
    }

    emailContent := fmt.Sprintf(`
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <style>
                body { font-family: Arial, sans-serif; color: #333; margin: 20px; }
                .container { width: 100%%; max-width: 600px; margin: 0 auto; }
                .header { text-align: center; padding: 20px 0; background-color: #f4f4f4; }
                .details { padding: 20px; background-color: #ffffff; margin-top: 10px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
                .section { margin-bottom: 20px; }
                .footer { text-align: center; font-size: 12px; color: #777; margin-top: 20px; }
                table { width: 100%%; border-collapse: collapse; margin-top: 10px; }
                th, td { padding: 10px; border: 1px solid #ddd; text-align: left; }
                th { background-color: #f4f4f4; }
            </style>
        </head>
        <body>
            <div class="container">
                <div class="header">
                    <h1>Receipt for Your Order</h1>
                    <p>Hello, thank you for shopping with Drink&Groove.</p>
                </div>
                <div class="details">
                    <div class="section order-info">
                        <h2>Order Details</h2>
                        <p><strong>Order ID:</strong> %s</p>
                        <p><strong>Order Time:</strong> %s</p>
                    </div>
                    <table>
                        <thead>
                            <tr>
                                <th>Item</th>
                                <th>Quantity</th>
                                <th>Price (THB)</th>
                                <th>Total (THB)</th>
                            </tr>
                        </thead>
                        <tbody>
                            %s
                        </tbody>
                    </table>
                    <div class="section total">
                        <h3>Total Amount: %.2f THB</h3>
                    </div>
                </div>
                <div class="footer">
                    <p>If you have any questions, contact us at <a href="mailto:lerdphipat.k@ku.th">lerdphipat.k@ku.th</a>.</p>
                    <p>Thank you for choosing Drink&Groove!</p>
                </div>
            </div>
        </body>
        </html>`,
        order.ID,
        order.Time,
        orderDetails,
        totalAmount,
    )

    message := mail.NewSingleEmail(from, subject, to, plainTextContent, emailContent)

    // if order.Url != "" {
    //     attachment := mail.NewAttachment()
    //     attachment.SetContent(order.OURL)
    //     attachment.SetType("application/pdf")
    //     attachment.SetFilename("Receipt.pdf")
    //     attachment.SetDisposition("attachment")
    //     message.AddAttachment(attachment)
    // }

    client := sendgrid.NewSendClient("ASK_I_BEST_FOR_THIS")
    response, err := client.Send(message)
    if err != nil {
        log.Println("Error sending email:", err)
        return
    }

    log.Printf("Email sent! Status Code: %d\nResponse Body: %s\nResponse Headers: %v\n", response.StatusCode, response.Body, response.Headers)
}

// func sendEmail(to, subject, body string) error {
// 	return nil
// }

func cleanupEmailWorker() {
	CleanupRabbitMQ()
	log.Println("Email worker cleanup complete")
}
