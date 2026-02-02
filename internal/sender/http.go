package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Nayati-Indonesia/spec-collector/models"
)

func HTTPSender(data models.SpecResponse) error {
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("failed marshal body: %v", err)
	}

	body := bytes.NewBuffer(b)
	resp, err := http.Post("https://ns.nayatisys.com/api/send-it-collector", "application/json; charset=utf-8", body)

	// TODO Temporary hardcoded URL API
	//resp, err := http.Post(cfg.APIUrl, "application/json; charset=utf-8", body)
	if err != nil {
		log.Fatalf("failed to send data: %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("Successfully sending data to the server! \n")
	} else {
		fmt.Printf("Server returned an error! \n")
	}

	return nil
}
