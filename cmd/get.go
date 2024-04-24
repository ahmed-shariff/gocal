/*
Copyright © 2024 Shariff AM Faleel <shariff.mfa@outlook.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	// TODO: Update descriptions
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Args: cobra.ExactArgs(1),
	Run: getCalenderData,
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getCalenderData(cmd *cobra.Command, args []string) {
	srv := getCalenderService()

	minTime := time.Now().Format(time.RFC3339)
	maxTime := time.Now().Add(time.Hour * 24 * 30).Format(time.RFC3339)

        events, err := srv.Events.List("primary").ShowDeleted(false).
                SingleEvents(true).TimeMin(minTime).TimeMax(maxTime).OrderBy("StartTime").Do()
        if err != nil {
                log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
        }

        fmt.Println("Upcoming events:")
        if len(events.Items) == 0 {
                fmt.Println("No upcoming events found.")
        } else {
                for _, item := range events.Items {
                        date := item.Start.DateTime
                        if date == "" {
                                date = item.Start.Date
                        }
                        // fmt.Printf("%v (%v) ---- %v \n", item.Summary, date, item.Id)
			entry := CalendarEntry {
				"Id": item.Id,
				"Summary": item.Summary,
				"DateTime": date,
			}
			s, _ := json.Marshal(entry)
			fmt.Println(string(s))
                }
        }
}

type CalendarEntry map[string]string
