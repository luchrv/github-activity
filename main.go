package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"
)

func main() {
	args := os.Args
	userName := args[1]

	apiURL := "https://api.github.com/users/" + userName + "/events"

	req, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", req.StatusCode)
		return
	}

	var response Events
	err = json.NewDecoder(req.Body).Decode(&response)
	if err != nil {
		fmt.Printf("Error decoding JSON response: %v\n", err)
		return
	}

	// Display the feched activity in the terminal
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Print header
	fmt.Fprintln(w, "Event ID\tType\tRepo\tAction\tCreated At")
	// Print separator
	fmt.Fprintln(w, "--------\t----\t----\t------\t----------")

	for _, event := range response {
			var eventType string
			if event.Type != nil {
				eventType = *event.Type
			}

			var payloadAction string
			if event.Payload.Action != nil {
				payloadAction = *event.Payload.Action
			}

			fmt.Fprintln(
				w,
				event.ID+"\t"+
						eventType+"\t"+
						event.Repo.Name+"\t"+
							payloadAction+"\t"+
							event.CreatedAt.String(),
			)
		}

	err = w.Flush()
	if err != nil {
		fmt.Printf("Error flushing output: %v\n", err)
	}
}
