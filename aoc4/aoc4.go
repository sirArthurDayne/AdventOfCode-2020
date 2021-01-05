package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
    byr string;
    iyr string;
    eyr string;
    hgt string;
    hcl string;
    ecl string;
    pid string;
    cid string
}

func main() {
    /*Advent Of Code day4. solution by: sirArthurDayne */
    //FIRST PART
    inputData, err:= scanLines("data.txt")
    if err != nil {
        panic(err)
    }

    //1.1 clean data
    sanitizeData := getRawPassportData(inputData)
    //1.2 parse data into list of objects
    passportList := parsePasswords(sanitizeData)
    //1.3 validate passportList
    // fmt.Printf("valid passports: %v",totalValidPasswords(passportList))

    //SECOND PART
    fmt.Printf("valid passports: %v",totalValidPasswords2(passportList))
}

func scanLines(path string) ([]string, error) {

  file, err := os.Open(path)
  if err != nil {
     return nil, err
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  scanner.Split(bufio.ScanLines)

  var lines []string

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines, nil
}

func getRawPassportData(inputData []string) []string {
    var sanitizeData = []string{}
    currentPassportData := ""
    for _,line := range inputData {
        if line != "" {
            if currentPassportData != "" {
            currentPassportData += " "+ line
            }else {
                currentPassportData = line
            }
        }else {
            sanitizeData = append(sanitizeData, currentPassportData)
            currentPassportData = ""
        }
    }
    return sanitizeData
}

func parsePasswords(data []string) []Passport {
    var passportList = []Passport{}
    for _,line := range data {
        sliceBySpace := strings.Split(line, " ")
        var passport = Passport{}
        for i:=0; i<len(sliceBySpace); i++ {
            sliceByColon := strings.Split(sliceBySpace[i], ":")
            value:=strings.TrimSpace(sliceByColon[1])
            switch  keyName:= sliceByColon[0];keyName {
                case "byr": passport.byr = value
                case "iyr": passport.iyr = value
                case "eyr": passport.eyr = value
                case "hgt": passport.hgt = value
                case "hcl": passport.hcl = value
                case "ecl": passport.ecl = value
                case "pid": passport.pid = value
                case "cid": passport.cid = value
                default: fmt.Printf("unknow keyName: %s",keyName)
            }
        }
        //append
        passportList = append(passportList, passport)
    }
    return passportList
}

//FIRST PART
func totalValidPasswords(passportList []Passport) int {
    count := 0
    for _, passport := range passportList {
        if passport.byr == "" || passport.iyr == "" || passport.eyr == "" ||
        passport.hgt == "" || passport.hcl == "" || passport.ecl == "" ||
        passport.pid == "" {
            continue
        }
        count++
    }
    return count
}

func totalValidPasswords2(passportList []Passport) int {
    count := 0
    for _, passport := range passportList {
        if !validYearRange(passport.byr, 1920,2002) ||
        !validYearRange(passport.iyr,2010,2020) ||
        !validYearRange(passport.eyr,2020,2030)||
        !validHeight(passport.hgt) || !validHairColor(passport.hcl) ||
        !validEyeColor(passport.ecl) ||
        !validPID(passport.pid) {
            continue
        }
        count++
    }
    return count
}

func validYearRange(field string, lowerbound, highbound int) bool {
    if field == "" {
        return false
    }
    num, err := strconv.Atoi(field)
    if err != nil {
        panic(err)
    }
    if len(field) == 4 && (num >= lowerbound && num <= highbound) {
        return true
    }
    return false
}

func validHeight(field string) bool {
    if field == "" {
        return false
    }
    //get number and metric
    if strings.Contains(field, "cm") {
        valueHeight,err := strconv.Atoi(strings.Split(field,"cm")[0])
        if err != nil {
            panic(err)
        }
        if valueHeight >= 150 && valueHeight <= 193 {
            return true
        }
    } else if strings.Contains(field, "in") {
        valueHeight,err := strconv.Atoi(strings.Split(field,"in")[0])
        if err != nil {
            panic(err)
        }
        if valueHeight >= 59 && valueHeight <= 76 {
            return true
        }
    }
    return false
}

func validHairColor(field string) bool {
    if field == "" || !strings.Contains(field, "#"){
        return false
    }

    matched, err := regexp.MatchString(`[0-9a-z#]`, field)
    if matched && err == nil{
        return true
    }

    return false
}

func validEyeColor(field string) bool {

    switch field {
        case "amb": return true
        case "blu": return true
        case "brn": return true
        case "gry": return true
        case "grn": return true
        case "hzl": return true
        case "oth": return true
        default: return false
    }
}

func validPID(field string) bool {
    return len(field) == 9
}
