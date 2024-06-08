package modules

import (
	"net/http"
	"os"
	"strings"
)

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request: Only POST method is allowed", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request: Unable to parse form", http.StatusBadRequest)
		return
	}

	input := r.FormValue("input")
	option := r.FormValue("options")
	if option == "" {
		option = "standard"
	}

	fileContent, err := os.ReadFile("banners/" + option + ".txt")
	if err != nil {
		http.Error(w, "Internal Server Error: Error reading file", http.StatusInternalServerError)
		return
	}

	str := string(fileContent)
	str = strings.ReplaceAll(str, "\r", "")
	lines := strings.Split(str, "\n")
	splittedArg := strings.Split(input, "\n")

	result := GenerateASCIIArt(splittedArg, lines)
	data := struct {
		Input  string
		Option string
		Output string
	}{
		Input:  input,
		Option: option,
		Output: result,
	}
	err = tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error: Error executing template", http.StatusInternalServerError)
	}
}

func GenerateASCIIArt(splittedArg, lines []string) string {
	result := ""
	for _, line := range splittedArg {
		if line != "" {
			for i := 1; i <= 8; i++ {
				for _, char := range line {
					if int(char)-32 < 0 || int(char)-32 > 127 {
						continue
					} else {
						result += FindLine(lines, char, i)
					}
				}
				result += "\n"
			}
		}
	}
	return result
}

func FindLine(lines []string, char rune, pos int) string {
	index := (int(char)-32)*9 + pos
	return lines[index]
}
