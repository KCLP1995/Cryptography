package crack

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const targetMD5 = "6a85dfd77d9cb35770c9dc6728d73d3f"

func MD5Cracking(wordlistPath, outputPath string) error {
	file, err := os.Open(wordlistPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		sum := md5.Sum([]byte(line))
		hexSum := hex.EncodeToString(sum[:])

		if hexSum == targetMD5 {
			fmt.Printf("===> MATCH FOUND! Line %d: %q\n", lineNum, line)

			// write the match to a separate file
			err := os.WriteFile(outputPath, []byte("The result of cracking MD5: " + line + "\n"), 0644)
			if err != nil {
				return fmt.Errorf("failed to write match to file: %w", err)
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
