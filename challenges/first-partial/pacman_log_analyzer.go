package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pacman struct {
	Installed int
	Removed   int
	Upgraded  int
}

type pacmanPackage struct {
	Name        string
	InstallDate string
	UpdateDate  string
	RemovalDate string
	Updates     int
}

func main() {
	fmt.Println("Pacman Log Analyzer")

	if len(os.Args) < 2 {
		fmt.Println("You must send at least one pacman log file to analize")
		fmt.Println("usage: ./pacman_log_analizer <logfile>")
		os.Exit(1)
	}

	// Your fun starts here.

	txtlines := getLines(os.Args[1])

	report, packages := getReport(txtlines)

	writeFile(report, packages, "packages_report.txt")
}

func getLines(name string) []string {

	fmt.Println("Opening file: " + name)

	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	return txtlines
}

func getReport(txtlines []string) (pacman, map[string]pacmanPackage) {

	report := pacman{0, 0, 0}
	packages := map[string]pacmanPackage{}
	for _, eachline := range txtlines {
		words := strings.Split(eachline, " ")
		switch words[3] {
		case "installed":
			date := strings.Replace(words[0], "[", "", -1) + " " + strings.Replace(words[1], "]", "", -1)
			packages[words[4]] = pacmanPackage{words[4], date, "-", "-", 0}
			report.Installed++
			break
		case "upgraded":
			upgraded := packages[words[4]]
			date := strings.Replace(words[0], "[", "", -1) + " " + strings.Replace(words[1], "]", "", -1)
			upgraded.UpdateDate = date
			upgraded.Updates++
			packages[words[4]] = upgraded
			report.Upgraded++
			break
		case "removed":
			removed := packages[words[4]]
			date := strings.Replace(words[0], "[", "", -1) + " " + strings.Replace(words[1], "]", "", -1)
			removed.UpdateDate = date
			removed.RemovalDate = date
			packages[words[4]] = removed
			report.Removed++
			break
		}
	}

	return report, packages
}

func writeFile(report pacman, packages map[string]pacmanPackage, name string) {

	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("failed writing file: %s", err)
	}
	defer f.Close()

	installed := strconv.Itoa(report.Installed)
	removed := strconv.Itoa(report.Removed)
	upgraded := strconv.Itoa(report.Upgraded)
	current := strconv.Itoa(report.Installed - report.Removed)

	w := bufio.NewWriter(f)
	_, err = w.WriteString("\nPacman Packages Report\n----------------------\n- Installed packages : " + installed + "\n- Removed packages   : " + removed + "\n- Upgraded packages  : " + upgraded + "\n- Current installed  : " + current + "\nList of packages\n----------------\n")
	if err != nil {
		log.Fatalf("failed writing file: %s", err)
	}
	//fmt.Print("\nPacman Packages Report\n----------------------\n- Installed packages : " + installed + "\n- Removed packages   : " + removed + "\n- Upgraded packages  : " + upgraded + "\n- Current installed  : " + current + "\nList of packages\n----------------\n")

	for _, value := range packages {
		_, err = w.WriteString("- Package Name          : " + value.Name + "\n\t- Install date      : " + value.InstallDate + "\n\t- Last update date  : " + value.UpdateDate + "\n\t- How many updates  : " + strconv.Itoa(value.Updates) + "\n\t- Removal date      : " + value.RemovalDate + "\n")
		if err != nil {
			log.Fatalf("failed writing file: %s", err)
		}
		//fmt.Print("- Package Name        : " + value.Name + "\n\t- Install date      : " + value.InstallDate + "\n\t- Last update date  : " + value.UpdateDate + "\n\t- How many updates  : " + strconv.Itoa(value.Updates) + "\n\t- Removal date      : " + value.RemovalDate + "\n")
	}

	defer w.Flush()
}
