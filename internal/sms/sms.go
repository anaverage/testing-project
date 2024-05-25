package sms

import (
	"fmt"
	"io/ioutil"
	"log"
	"network-service/internal/entities"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/jinzhu/copier"
)

func GetResultSMSData(path string, wg *sync.WaitGroup) ([][]entities.SMSData, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("File %s does not exist", path)
	}

	first := smsReader(path)
	second := make([]entities.SMSData, len(first))
	err := copier.Copy(&second, &first)
	if err != nil {
		log.Print(err)
	}

	sort.SliceStable(first, func(i, j int) bool {
		return first[i].Provider < first[j].Provider
	})

	sort.SliceStable(second, func(i, j int) bool {
		return second[i].Country < second[j].Country
	})

	result := [][]entities.SMSData{
		first, second,
	}

	return result, nil
}

func smsReader(path string) []entities.SMSData {
	var result []entities.SMSData
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("File", path, "does not exist")
			return result
		} else {
			log.Println("Error opening file:", err)
			return result
		}
	}
	defer file.Close()

	reader, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Cannot read file:", err)
		return result
	}

	lines := strings.Split(string(reader), "\n")
	for _, value := range lines {
		splitVal := strings.Split(value, ";")
		if len(splitVal) == 4 {
			if checkSMS(splitVal) {
				res := entities.SMSData{
					Country:      CountryFromAlpha(splitVal[0]),
					Bandwidth:    splitVal[1],
					ResponseTime: splitVal[2],
					Provider:     splitVal[3],
				}
				result = append(result, res)
			}
		}
	}
	return result
}

func checkSMS(value []string) bool {
	if value[0] == CountryAlpha2()[value[0]] {
		percentValue, err := strconv.Atoi(value[1])
		if err != nil {
			log.Printf("The channel bandwidth value %v does not match the expected value.", value)
			return false
		}
		if -1 < percentValue && percentValue < 101 {
			_, err := strconv.Atoi(value[2])
			if err == nil {
				providers := map[string]string{"Topolo": "Topolo", "Rond": "Rond", "Kildy": "Kildy"}
				if value[3] == providers[value[3]] {
					return true
				} else {
					log.Printf("The provider's %v value does not match what is expected.", value)
					return false
				}
			} else {
				log.Printf("The response value in ms %v does not match the expected value.", value)
				return false
			}
		} else {
			log.Printf("The channel bandwidth value %v does not match the expected value.", value)
			return false
		}
	} else {
		log.Printf("The country value alpha-2 %v does not match the expected value.", value)
		return false
	}
}

func CountryFromAlpha(str string) string {
	list := map[string]string{"AD": "Andorra", "AE": "United Arab Emirates", "AF": "Afghanistan", "AG": "Antigua and Barbuda", "AI": "Anguilla", "AL": "Albania",
		"AM": "Armenia", "AO": "Angola", "AQ": "Antarctica", "AR": "Argentina", "AS": "American Samoa", "AT": "Austria", "AU": "Australia",
		"AW": "Aruba", "AX": "Åland Islands", "AZ": "Azerbaijan", "BA": "Bosnia and Herzegovina", "BB": "Barbados", "BD": "Bangladesh",
		"BE": "Belgium", "BF": "Burkina Faso", "BG": "Bulgaria", "BH": "Bahrain", "BI": "Burundi", "BJ": "Benin", "BL": "Saint Barthélemy",
		"BM": "Bermuda", "BN": "Brunei Darussalam", "BO": "Bolivia", "BQ": "Bonaire, Sint Eustatius and Saba", "BR": "Brazil", "BS": "Bahamas",
		"BT": "Bhutan", "BV": "Bouvet Island", "BW": "Botswana", "BY": "Belarus", "BZ": "Belize", "CA": "Canada", "CC": "Cocos (Keeling) Islands",
		"CD": "Congo, Democratic Republic of the", "CF": "Central African Republic", "CG": "Congo", "CH": "Switzerland", "CI": "Côte d'Ivoire",
		"CK": "Cook Islands", "CL": "Chile", "CM": "Cameroon", "CN": "China", "CO": "Colombia", "CR": "Costa Rica", "CU": "Cuba", "CV": "Cabo Verde",
		"CW": "Curaçao", "CX": "Christmas Island", "CY": "Cyprus", "CZ": "Czechia", "DE": "Germany", "DK": "Denmark", "DM": "Dominica",
		"DO": "Dominican Republic", "DZ": "Algeria", "EC": "Ecuador", "EE": "Estonia", "EG": "Egypt", "EH": "Western Sahara", "ER": "Eritrea",
		"ES": "Spain", "ET": "Ethiopia", "FI": "Finland", "FJ": "Fiji", "FK": "Falkland Islands (Malvinas)", "FM": "Micronesia (Federated States of)",
		"FO": "Faroe Islands", "FR": "France", "GA": "Gabon", "GB": "United Kingdom of Great Britain and Northern Ireland", "GD": "Grenada",
		"GE": "Georgia", "GF": "French Guiana", "GG": "Guernsey", "GH": "Ghana", "GI": "Gibraltar", "GL": "Greenland", "GM": "Gambia", "GN": "Guinea",
		"GP": "Guadeloupe", "GQ": "Equatorial Guinea", "GR": "Greece", "GS": "South Georgia and the South Sandwich Islands", "GT": "Guatemala",
		"GU": "Guam", "GW": "Guinea-Bissau", "GY": "Guyana", "HK": "Hong Kong", "HM": "Heard Island and McDonald Islands", "HN": "Honduras",
		"HR": "Croatia", "HT": "Haiti", "HU": "Hungary", "ID": "Indonesia", "IE": "Ireland", "IL": "Israel", "IM": "Isle of Man", "IN": "India",
		"IO": "British Indian Ocean Territory", "IQ": "Iraq", "IR": "Iran (Islamic Republic of)", "IS": "Iceland", "IT": "Italy", "JE": "Jersey",
		"JM": "Jamaica", "JO": "Jordan", "JP": "Japan", "KE": "Kenya", "KG": "Kyrgyzstan", "KH": "Cambodia", "KI": "Kiribati", "KM": "Comoros",
		"KN": "Saint Kitts and Nevis", "KP": "Korea (Democratic People's Republic of)", "KR": "Korea, Republic of", "KW": "Kuwait",
		"KY": "Cayman Islands", "KZ": "Kazakhstan", "LA": "Lao People's Democratic Republic", "LB": "Lebanon", "LC": "Saint Lucia",
		"LI": "Liechtenstein", "LK": "Sri Lanka", "LR": "Liberia", "LS": "Lesotho", "LT": "Lithuania", "LU": "Luxembourg", "LV": "Latvia",
		"LY": "Libya", "MA": "Morocco", "MC": "Monaco", "MD": "Moldova, Republic of", "ME": "Montenegro", "MF": "Saint Martin (French part)",
		"MG": "Madagascar", "MH": "Marshall Islands", "MK": "North Macedonia", "ML": "Mali", "MM": "Myanmar", "MN": "Mongolia", "MO": "Macao",
		"MP": "Northern Mariana Islands", "MQ": "Martinique", "MR": "Mauritania", "MS": "Montserrat", "MT": "Malta", "MU": "Mauritius",
		"MV": "Maldives", "MW": "Malawi", "MX": "Mexico", "MY": "Malaysia", "MZ": "Mozambique", "NA": "Namibia", "NC": "New Caledonia", "NE": "Niger",
		"NF": "Norfolk Island", "NG": "Nigeria", "NI": "Nicaragua", "NL": "Netherlands", "NO": "Norway", "NP": "Nepal", "NR": "Nauru", "NU": "Niue",
		"NZ": "New Zealand", "OM": "Oman", "PA": "Panama", "PE": "Peru", "PF": "French Polynesia", "PG": "Papua New Guinea", "PH": "Philippines",
		"PK": "Pakistan", "PL": "Poland", "PM": "Saint Pierre and Miquelon", "PN": "Pitcairn", "PR": "Puerto Rico", "PS": "Palestine, State of",
		"PT": "Portugal", "PW": "Palau", "PY": "Paraguay", "QA": "Qatar", "RE": "Réunion", "RO": "Romania", "RS": "Serbia", "RU": "Russian Federation",
		"RW": "Rwanda", "SA": "Saudi Arabia", "SB": "Solomon Islands", "SC": "Seychelles", "SD": "Sudan", "SE": "Sweden", "SG": "Singapore",
		"SH": "Saint Helena, Ascension and Tristan da Cunha", "SI": "Slovenia", "SJ": "Svalbard and Jan Mayen", "SK": "Slovakia",
		"SL": "Sierra Leone", "SM": "San Marino", "SN": "Senegal", "SO": "Somalia", "SR": "Suriname", "SS": "South Sudan", "ST": "Sao Tome and Principe",
		"SV": "El Salvador", "SX": "Sint Maarten (Dutch part)", "SY": "Syrian Arab Republic", "SZ": "Eswatini", "TC": "Turks and Caicos Islands",
		"TD": "Chad", "TF": "French Southern Territories", "TG": "Togo", "TH": "Thailand", "TJ": "Tajikistan", "TK": "Tokelau", "TL": "Timor-Leste",
		"TM": "Turkmenistan", "TN": "Tunisia", "TO": "Tonga", "TR": "Türkiye", "TT": "Trinidad and Tobago", "TV": "Tuvalu", "TW": "Taiwan, Province of China",
		"TZ": "Tanzania, United Republic of", "UA": "Ukraine", "UG": "Uganda", "UM": "United States Minor Outlying Islands",
		"US": "United States of America", "UY": "Uruguay", "UZ": "Uzbekistan", "VA": "Holy See", "VC": "Saint Vincent and the Grenadines",
		"VE": "Venezuela (Bolivarian Republic of)", "VG": "Virgin Islands (British)", "VI": "Virgin Islands (U.S.)", "VN": "Viet Nam",
		"VU": "Vanuatu", "WF": "Wallis and Futuna", "WS": "Samoa", "YE": "Yemen", "YT": "Mayotte", "ZA": "South Africa", "ZM": "Zambia", "ZW": "Zimbabwe"}
	return list[str]
}

func CountryAlpha2() map[string]string {
	str := "AD AE AF AG AI AL AM AO AQ AR AS AT AU AW AX AZ BA BB BD BE BF BG BH BI BJ BL BM BN BO BQ BR BS BT BV BW BY BZ CA CC CD CF CG CH CI CK CL CM CN CO CR CU CV CW CX CY CZ DE DJ DK DM DO DZ EC EE EG EH ER ES ET FI FJ FK FM FO FR GA GB GD GE GF GG GH GI GL GM GN GP GQ GR GS GT GU GW GY HK HM HN HR HT HU ID IE IL IM IN IO IQ IR IS IT JE JM JO JP KE KG KH KI KM KN KP KR KW KY KZ LA LB LC LI LK LR LS LT LU LV LY MA MC MD ME MF MG MH MK ML MM MN MO MP MQ MR MS MT MU MV MW MX MY MZ NA NC NE NF NG NI NL NO NP NR NU NZ OM PA PE PF PG PH PK PL PM PN PR PS PT PW PY QA RE RO RS RU RW SA SB SC SD SE SG SH SI SJ SK SL SM SN SO SR SS ST SV SX SY SZ TC TD TF TG TH TJ TK TL TM TN TO TR TT TV TW TZ UA UG UM US UY UZ VA VC VE VG VI VN VU WF WS YE YT ZA ZM ZW"
	alpha := make(map[string]string)
	splitStr := strings.Split(str, " ")
	for _, value := range splitStr {
		alpha[value] = value
	}
	return alpha
}
