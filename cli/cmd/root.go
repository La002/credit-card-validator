package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

func init() {
	rootCmd.AddCommand(validateCmd)
}

var (
	rootCmd = &cobra.Command{
		Use:   "cc",
		Short: "Credit Card CLI",
		Long:  `A simple CLI to validate credit card numbers using the Luhn algorithm.`,
	}

	validateCmd = &cobra.Command{
		Use:   "validate <number>",
		Short: "Validate credit card numbers using the Luhn algorithm.",
		Run: func(cmd *cobra.Command, args []string) {
			number := args[0]
			c, err := SendRequest(number)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println(c)
		},
	}
)

func SendRequest(number string) (string, error) {
	url := fmt.Sprintf("http://localhost:8082/%s", number)

	resp, err := http.Get(url)
	if err != nil {
		return "Error in sending request:", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response:", err
	}

	return string(body), nil
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
