package sms

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"editory/editory_sms_service/genproto/sms_service"
	"editory/editory_sms_service/pkg/models"
	"editory/editory_sms_service/storage"

	_ "github.com/lib/pq" //db driver

	"editory/editory_sms_service/config"
)

//Daemon ...
type Daemon struct {
	Conf config.Config
	Strg storage.StorageI
}

var (
	table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
)

// Init initializes Deamon
func (dmn *Daemon) Init() {

	c := make(chan string)

	isShuttingDown := false
	started := false

OUTTER:
	for {
		fmt.Println("1")

		if isShuttingDown {
			break
		}
		select {
		case s, ok := <-c:

			if ok {
				fmt.Printf("\nservice is going down...\n")
				fmt.Println(fmt.Sprintf("\nReceived signal %x\n", s))
				isShuttingDown = true
				c = nil
				continue OUTTER
			}
		default:
			var smss []*sms_service.Sms
			fmt.Println("0")
			smss, err := dmn.Strg.Sms().GetNotSent(context.Background())
			if err != nil {
				log.Println("Error while checking status: ", err)
			}
			for _, sms := range smss {
				fmt.Println("0.1")
				fmt.Println(sms.PhoneNumber)
				fmt.Println(sms.Otp)
				err = dmn.sendSms(sms.Text+": "+sms.Otp, sms.PhoneNumber)
				if err != nil {
					log.Println("error while sending sms: ", err)

					err := dmn.Strg.Sms().IncrementSendCount(context.Background(), sms.Id)
					log.Println("error while incrementing sms send trying count ", err)
				} else {
					err = dmn.Strg.Sms().MakeSent(context.Background(), sms.Id)
					if err != nil {
						log.Println("error while updating status")
					}
				}
			}

			if !started {
				fmt.Print("\n\nSuccessfully\n\n")
				started = true
			}
			t := time.Now()
			fmt.Println(fmt.Sprint("Service is running...\n", t.Format(time.RFC3339)))
			time.Sleep(10000 * time.Millisecond)
		}
	}
	fmt.Print("\n\nThis service has finally shutted down\n\n")
}

// Send Mail...
func (dmn *Daemon) sendSms(text, phoneNumber string) error {
	fmt.Println("2")
	if len(phoneNumber) < 2 {
		return nil
	}
	phone := phoneNumber[1:]
	return sendWithPlayMobile(dmn, text, phone)
}

func sendWithPlayMobile(dmn *Daemon, text, phoneNumber string) error {

	//[0] text [1] code
	textAndCode := strings.Split(text, ":")

	cfg := config.Load()

	var body models.Body
	client := http.Client{}

	message := models.Message{
		Recipient: phoneNumber,
		MessageID: fmt.Sprintf("%s%s", "del", textAndCode[1]),
		SMS: models.SMS{
			Originator: dmn.Conf.PlayMobileOriginator,
			Content: models.Content{
				Text: text,
			},
		},
	}

	body.Messages = append(body.Messages, message)

	requestBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", dmn.Conf.PlayMobileUrl, bytes.NewBuffer(requestBody))

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	request.SetBasicAuth(cfg.PlayMobileLogin, cfg.PlayMobilePassword)

	res, err := client.Do(request)
	if err != nil {
		return err
	}
	fmt.Println("3")

	if res.StatusCode != 200 {
		return errors.New("error while sending sms")
	}

	return nil
}
