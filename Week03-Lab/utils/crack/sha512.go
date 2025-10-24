package crack

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const targetSha512 = "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

func Sha512Cracking(wordlistPath, outputPath string) error {
	file, err := os.Open(wordlistPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		sum := sha512.Sum512([]byte(line))
		hexSum := hex.EncodeToString(sum[:])

		if hexSum == targetSha512 {
			fmt.Printf("===> MATCH FOUND! Line %d: %q\n", lineNum, line)

			// append the match to the output file instead of overwriting
			f, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return fmt.Errorf("failed to open output file: %w", err)
			}
			defer f.Close()

			if _, err := f.WriteString("The result of cracking SHA512: " + line + "\n"); err != nil {
				return fmt.Errorf("failed to append match to file: %w", err)
			}

			return nil // stop after first match
		}

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Println("No match found in the wordlist.")
	return nil
}
