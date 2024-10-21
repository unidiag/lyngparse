package lyngparse

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func LyngSat(path string) []string {

	url := "https://www.lyngsat.com" + path
	var cfg []string

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка при запросе страницы: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Ошибка: статус-код %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error", err)
		os.Exit(0)
	}
	html := strings.ReplaceAll(string(body), "\n", "")

	freqPolarityRegex := regexp.MustCompile(`<a[^>]*>(\d+)&nbsp;([HVRL])</a>.*?DVB-(S|S2)(?:<br>\S+)?<br>(\d+)<br>`)
	matches := freqPolarityRegex.FindAllStringSubmatch(html, -1)

	for _, match := range matches {
		frequency := match[1]      // Частота
		polarization := match[2]   // Поляризация
		modulationType := match[3] // Тип модуляции
		symbolRate := match[4]     // Скорость

		str := fmt.Sprintf("%.1f:%s:%s:%s:%s", extractSatelliteDegree(html), modulationType, frequency, polarization, symbolRate)
		if !in_array(str, cfg) {
			cfg = append(cfg, str)
		}
	}
	return cfg
}

func extractSatelliteDegree(title string) float64 {
	re := regexp.MustCompile(`<title>.*?(\d+\.?\d*)°\s?([EW]).*?</title>`)
	match := re.FindStringSubmatch(title)
	if len(match) == 3 {
		degree, err := strconv.ParseFloat(match[1], 64)
		if err != nil {
			return 0
		}
		if match[2] == "W" {
			degree = -degree
		}
		return degree
	}
	return 0
}

func in_array(str string, array []string) bool {
	for _, v := range array {
		if v == str {
			return true
		}
	}
	return false
}
