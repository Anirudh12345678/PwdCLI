package cmd

import (
	"fmt"
	"math/rand/v2"
	"sync"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate random passwords",
	Long:  "Gnerate random passwords with customizable properties",
	Run:   generatePassword,
}

func generatePassword(cmd *cobra.Command, args []string) {
	wg := new(sync.WaitGroup)
	length, _ := cmd.Flags().GetInt("len")
	dig, _ := cmd.Flags().GetBool("digits")
	spl, _ := cmd.Flags().GetBool("spl")
	n, _ := cmd.Flags().GetInt("times")
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
	if n > 20 {
		fmt.Println("Cannot generate more than 20 passwords at once!!")
		return
	}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go generateNTimes(set, length, wg)
	}
	wg.Wait()
}

func generateNTimes(set string, length int, wg *sync.WaitGroup) {
	defer wg.Done()
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
	generateCmd.Flags().IntP("times", "t", 1, "Number of passwords to generate")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in your password")
	generateCmd.Flags().BoolP("spl", "s", false, "Include special characters")
}
