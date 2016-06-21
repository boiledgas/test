package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Notification struct {
	Title          string   `json:"title,omitempty"`          // notification title
	Body           string   `json:"body,omitempty"`           // notification body text
	Icon           string   `json:"icon,omitempty"`           // notification icon
	Sound          string   `json:"sound,omitempty"`          // sound to play when the device receives a notification
	Badge          string   `json:"badge,omitempty"`          // badge on the client app home icon
	Tag            string   `json:"tag,omitempty"`            // notification with the same tag is already being shown, the new notification replaces the existing one in the notification drawer
	Color          string   `json:"color,omitempty"`          // color of the icon, expressed in #rrggbb format
	Click_action   string   `json:"click_action,omitempty"`   // activity with a matching intent filter is launched when user clicks the notification
	Body_loc_key   string   `json:"body_loc_key,omitempty"`   // body string for localization. app's string resources
	Body_loc_args  []string `json:"body_loc_args,omitempty"`  // format arguments for the string resources
	Title_loc_key  string   `json:"title_loc_key,omitempty"`  // title string for localization. app's string resources
	Title_loc_args []string `json:"title_loc_args,omitempty"` // format specifiers in the title string
}

type FcmMessage struct {
	To               string   `json:"to,omitempty"`               // registration token, notification key, or topic
	Registration_ids []string `json:"registration_ids,omitempty"` // use this parameter only for multicast messaging, not for single recipients
	Condition        string   `json:"condition,omitempty"`        // conditions that determine the message target

	Collapse_key            string `json:"collapse_key,omitempty"`            // last message gets sent when delivery can be resumed, maximum of 4 different collapse keys
	Priority                string `json:"priority,omitempty"`                // normal, high
	Content_available       bool   `json:"content_available,omitempty"`       // set to true, an inactive client app is awoken
	Delay_while_idle        bool   `json:"delay_while_idle,omitempty"`        // indicates that the message should not be sent until the device becomes active
	Time_to_life            int    `json:"time_to_live,omitempty"`            // The maximum time to live supported is 4 weeks, and the default value is 4 weeks, 0 .. 2,419,200 sec
	Restricted_package_name string `json:"restricted_package_name,omitempty"` // This parameter specifies the package name of the application where the registration tokens must match in order to receive the message
	Dry_run                 bool   `json:"dry_run,omitempty"`                 // allows developers to test a request without actually sending a message
	// payload
	Data         interface{}  `json:"data,omitempty"`         // 4KB, custom keys. The key should not be a reserved word
	Notification Notification `json:"notification,omitempty"` // 2KB, predefined keys. Sruct Notification
}

func Notify(registrationId string, title string, message string) (err error) {
	fcmMessage := FcmMessage{
		To:       registrationId,
		Priority: "normal",
		Notification: Notification{
			Title: title,
			Body:  message,
			Icon:  "ic_launcher",
			Tag:   "message",
			Sound: "default",
		},
	}

	var bodyBytes []byte
	if bodyBytes, err = json.Marshal(fcmMessage); err != nil {
		return
	}

	body := bytes.NewReader(bodyBytes)
	req, _ := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", body)
	req.Header.Set("Authorization", "key=AIzaSyAKt2yDatfz3nKbE5l5l58_oXL40hGaZVo")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}

	defer resp.Body.Close()
	var resp_bytes []byte
	resp_bytes, err = ioutil.ReadAll(resp.Body)
	log.Println(string(resp_bytes))

	return
}

type request_login struct {
	RegistrationId string `json:"registrationId" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.POST("/login", func(c *gin.Context) {
		var req request_login
		if c.BindJSON(&req) != nil {
			c.JSON(http.StatusOK, gin.H{"status": "cannot bind"})
			return
		}

		if len(req.RegistrationId) > 0 {
			go func() {
				i := 0
				for {
					title := fmt.Sprintf("title %v", i)
					message := fmt.Sprintf("message %v", i)
					Notify(req.RegistrationId, title, message)
					time.Sleep(5 * time.Second)
					i++
				}
			}()
		}

		c.JSON(http.StatusOK, gin.H{"status": "hmmm"})
	})

	r.Run(":777")
}
