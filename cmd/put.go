/*
Copyright © 2024 Shariff AM Faleel <shariff.mfa@outlook.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/api/calendar/v3"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run:  putCalenderData,
}

func init() {
	rootCmd.AddCommand(putCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func putCalenderData(cmd *cobra.Command, args []string) {
	var data CalendarEntry
	err_ := json.Unmarshal([]byte(strings.Join(args, "")), &data)
	if err_ != nil {
		log.Fatalf("Failed to decode json %v", err_)
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(data.Id) {
		log.Fatalf("Id has non alphanumeric chars. Got: %v", data.Id)
	}

	event := calendar.Event{
		Id: data.Id,
		Summary: data.Summary,
		Start: &calendar.EventDateTime{
			DateTime: data.StartDateTime,
		},
		End: &calendar.EventDateTime{
			DateTime: data.EndDateTime,
		},
	}

	var (
		err           error
		returnedEvent *calendar.Event
	)

	srv := getCalenderService()
	// '{\"Summary\":\"test event\", \"StartDateTime\":\"2024-04-25T12:00:00-07:00\", \"EndDateTime\":\"2024-04-25T13:00:00-07:00\", \"Id\":\"58b2ca69_4620_4c4f_8dj26_72b3375b5bee\"}'
	returnedEvent, err = srv.Events.Get("primary", event.Id).Do()

	if err != nil {
		returnedEvent, err = srv.Events.Insert("primary", &event).Do()
		if err != nil {
			log.Fatalf("Inserting Failed with error %v", err)
		}
	} else {
		returnedEvent, err = srv.Events.Update("primary", returnedEvent.Id, &event).Do()
		if err != nil {
			log.Fatalf("Updating Failed with error %v", err)
		}
	}
	fmt.Println(getJsonStringForEvent(*returnedEvent))
}
