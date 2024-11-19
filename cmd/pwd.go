package cmd

import (
	"fmt"
	"math/rand/v2"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate random passwords",
	Long:  "Gnerate random passwords with customizable properties",
	Run:   generatePassword,
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("len")
	dig, _ := cmd.Flags().GetBool("digits")
	spl, _ := cmd.Flags().GetBool("spl")

	if length > 128 {
		fmt.Println("Password length cannot exceed 128")
		return
	}
	set := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if dig {
		set += "0123456789"
	}
	if spl {
		set += "!@#$%&"
	}

	password := make([]byte, length)
	for i := range password {
		password[i] = set[rand.IntN(len(set))]
	}

	str := fmt.Sprintf("%s", password)
	fmt.Println(str)
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("len", "l", 8, "Length")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in your password")
	generateCmd.Flags().BoolP("spl", "s", false, "Include special characters")
}
