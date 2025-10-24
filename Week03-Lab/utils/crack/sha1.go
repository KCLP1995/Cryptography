package crack

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const targetSha1 = "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

func Sha1Cracking(wordlistPath, outputPath string) error {
	file, err := os.Open(wordlistPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		sum := sha1.Sum([]byte(line))
		hexSum := hex.EncodeToString(sum[:])

		if hexSum == targetSha1 {
			fmt.Printf("===> MATCH FOUND! Line %d: %q\n", lineNum, line)

			// append the match to the output file instead of overwriting
			f, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return fmt.Errorf("failed to open output file: %w", err)
			}
			defer f.Close()

			if _, err := f.WriteString("The result of cracking SHA1: " + line + "\n"); err != nil {
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
