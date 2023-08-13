package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var client = &http.Client{}

func main() {
	
	channels := []string{
		"https://t.me/s/V2rayNGn",
		"https://t.me/s/free4allVPN",
		"https://t.me/s/PrivateVPNs",
		"https://t.me/s/DirectVPN",
		"https://t.me/s/v2ray_outlineir",
		"https://t.me/s/NetAccount",
		"https://t.me/s/oneclickvpnkeys",
		"https://t.me/s/daorzadannet",
		"https://t.me/s/Outline_Vpn",
		"https://t.me/s/vpn_xw",
		"https://t.me/s/prrofile_purple",
		"https://t.me/s/ShadowSocks_s",
		"https://t.me/s/azadi_az_inja_migzare",
		"https://t.me/s/WomanLifeFreedomVPN",
		"https://t.me/s/customv2ray",
		"https://t.me/s/UnlimitedDev",
		"https://t.me/s/vmessorg",
		"https://t.me/s/v2rayNG_Matsuri",
		"https://t.me/s/v2rayngvpn",
		"https://t.me/s/vpn_ioss",
		"https://t.me/s/FalconPolV2rayNG",
		"https://t.me/s/v2rayNGNeT",
		"https://t.me/s/ShadowProxy66",
		"https://t.me/s/ipV2Ray",
		"https://t.me/s/kiava",
		"https://t.me/s/Helix_Servers",
		"https://t.me/s/PAINB0Y",
		"https://t.me/s/VpnProSec",
		"https://t.me/s/VlessConfig",
		"https://t.me/s/NIM_VPN_ir",
		"https://t.me/s/hashmakvpn",
		"https://t.me/s/iran_ray",
		"https://t.me/s/INIT1984",
		"https://t.me/s/ServerNett",
		"https://t.me/s/Pinkpotatocloud",
		"https://t.me/s/CloudCityy",
		"https://t.me/s/Qv2raychannel",
		"https://t.me/s/shopingv2ray",
		"https://t.me/s/xrayproxy",
		"https://t.me/s/Proxy_PJ",
		"https://t.me/s/V2rayng_Fast",
		"https://t.me/s/v2ray_swhil",
		"https://t.me/s/LoRd_uL4mo",
		"https://t.me/s/proxyymeliii",
		"https://t.me/s/MsV2ray",
		"https://t.me/s/free_v2rayyy",
		"https://t.me/s/v2ray1_ng",
		"https://t.me/s/vless_vmess",
		"https://t.me/s/MTConfig",
		"https://t.me/s/PNG_V2rayNG",
		"https://t.me/s/v2rayNG_VPNN",
		"https://t.me/s/vmess_vless_v2rayng",
		"https://t.me/s/Cov2ray",
		"https://t.me/s/V2RayTz",
		"https://t.me/s/VmessProtocol",
		"https://t.me/s/MehradLearn",
		"https://t.me/s/SafeNet_Server",
		"https://t.me/s/ovpn2",
		"https://t.me/s/lrnbymaa",
		"https://t.me/s/vpn_tehran",
		"https://t.me/s/OutlineVpnOfficial",
		"https://t.me/s/v2ray_vpn_ir",
		"https://t.me/s/v2_team",
		"https://t.me/s/ConfigsHUB",
		"https://t.me/s/V2rayngninja",
		"https://t.me/s/iSegaro",
		"https://t.me/s/bright_vpn",
		"https://t.me/s/talentvpn",
		"https://t.me/s/proxystore11",
		"https://t.me/s/yaney_01",
		"https://t.me/s/rayvps",
		"https://t.me/s/free1_vpn",
		"https://t.me/s/Parsashonam",
		"https://t.me/s/v2rayngvp",
		"https://t.me/s/Hope_Net",
		"https://t.me/s/VPNCLOP",
		"https://t.me/s/fnet00",
		"https://t.me/s/V2pedia",
		"https://t.me/s/molovpn",
		"https://t.me/s/melov2ray",
		"https://t.me/s/polproxy",
		"https://t.me/s/Outlinev2rayNG",
		"https://t.me/s/iP_CF",
		"https://t.me/s/VPNCUSTOMIZE",
		"https://t.me/s/MoV2ray",
		"https://t.me/s/Royalping_ir",
		"https://t.me/s/v2rayng_vpnrog",
		"https://t.me/s/v2rayng_config_amin",
		"https://t.me/s/rxv2ray",
		"https://t.me/s/Capital_NET",
		"https://t.me/s/VpnFreeSec",
		"https://t.me/s/lightning6",
		"https://t.me/s/WebShecan",
		"https://t.me/s/v2Line",
		"https://t.me/s/vmessiran",
		"https://t.me/s/Configforvpn01",
		"https://t.me/s/God_CONFIG",
		"https://t.me/s/Awlix_ir",
		"https://t.me/s/FreakConfig",
		"https://t.me/s/frev2ray",
		"https://t.me/s/vpnmasi",
		"https://t.me/s/BestV2rang",
		"https://t.me/s/TPvpn_Official",
		"https://t.me/s/AM_TEAMMM",
		"https://t.me/s/Lockey_vpn",
		"https://t.me/s/XsV2ray",
		"https://t.me/s/shh_proxy",
		"https://t.me/s/L_AGVPN13",
		"https://t.me/s/iTV2RAY",
		"https://t.me/s/V2rayNGmat",
		"https://t.me/s/ARv2ray",
		"https://t.me/s/V2parsin",
		"https://t.me/s/reality_daily",
		
			}

	configs := map[string]string{
		"ss":     "",
		"vmess":  "",
		"trojan": "",
		"vless":  "",
		"mixed":  "",
	}

	myregex := map[string]string{
		"ss":     `(.{3})ss:\/\/`,
		"vmess":  `vmess:\/\/`,
		"trojan": `trojan:\/\/`,
		"vless":  `vless:\/\/`,
	}

	for i := 0; i < len(channels); i++ {
		all_messages := false
		if strings.Contains(channels[i], "{all_messages}") {
			all_messages = true
			channels[i] = strings.Split(channels[i], "{all_messages}")[0]
		}

		req, err := http.NewRequest("GET", channels[i], nil)
		if err != nil {
			log.Fatalf("Error When requesting to: %d Error : %s", channels[i], err)
		}

		resp, err1 := client.Do(req)
		if err1 != nil {
			log.Fatal(err1)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		messages := doc.Find(".tgme_widget_message_wrap").Length()
		link, exist := doc.Find(".tgme_widget_message_wrap .js-widget_message").Last().Attr("data-post")
		if messages < 100 && exist == true {
			number := strings.Split(link, "/")[1]
			fmt.Println(number)

			doc = GetMessages(100, doc, number, channels[i])
		}

		var allConfigs []string // List to store all configurations for a channel

		if all_messages {
			fmt.Println(doc.Find(".js-widget_message_wrap").Length())
			doc.Find(".tgme_widget_message_text").Each(func(j int, s *goquery.Selection) {
				message_text := s.Text()
				lines := strings.Split(message_text, "\n")
				for a := 0; a < len(lines); a++ {
					for _, regex_value := range myregex {
						re := regexp.MustCompile(regex_value)
						lines[a] = re.ReplaceAllStringFunc(lines[a], func(match string) string {
							return "\n" + match
						})
					}
					for proto, _ := range configs {
						if strings.Contains(lines[a], proto) {
							configs["mixed"] += "\n" + lines[a] + "\n"
						}
					}
				}
			})
		} else {
			doc.Find("code,pre").Each(func(j int, s *goquery.Selection) {
				message_text := s.Text()
				lines := strings.Split(message_text, "\n")
				for a := 0; a < len(lines); a++ {
					for proto_regex, regex_value := range myregex {
						re := regexp.MustCompile(regex_value)
						lines[a] = re.ReplaceAllStringFunc(lines[a], func(match string) string {
							if proto_regex == "ss" {
								if match[:3] == "vme" {
									return "\n" + match
								} else if match[:3] == "vle" {
									return "\n" + match
								} else {
									return "\n" + match
								}
							} else {
								return "\n" + match
							}
						})

						if len(strings.Split(lines[a], "\n")) > 1 {
							myconfigs := strings.Split(lines[a], "\n")
							for i := 0; i < len(myconfigs); i++ {
								if myconfigs[i] != "" {
									re := regexp.MustCompile(regex_value)
									myconfigs[i] = strings.ReplaceAll(myconfigs[i], " ", "")
									match := re.FindStringSubmatch(myconfigs[i])
									if len(match) >= 1 {
										if proto_regex == "ss" {
											if match[1][:3] == "vme" {
												configs["vmess"] += "\n" + myconfigs[i] + "\n"
											} else if match[1][:3] == "vle" {
												configs["vless"] += "\n" + myconfigs[i] + "\n"
											} else {
												configs["ss"] += "\n" + myconfigs[i][3:] + "\n"
											}
										} else {
											configs[proto_regex] += "\n" + myconfigs[i] + "\n"
										}
									}
								}
							}
						}
					}
				}
			})
		}

		// Append the extracted configurations to the allConfigs list
		for _, config := range configs {
			allConfigs = append(allConfigs, config)
		}

		// Select the last 10 configurations from the list
		startIndex := len(allConfigs) - 10
		if startIndex < 0 {
			startIndex = 0
		}
		selectedConfigs := allConfigs[startIndex:]

		// Process and write the selected configurations
		for _, config := range selectedConfigs {
			WriteToFile(RemoveDuplicate(config), proto+"_iran.txt")
		}
	}
}

func WriteToFile(fileContent string, filePath string) {
	// Check if the file exists
	if _, err := os.Stat(filePath); err == nil {
		// If the file exists, clear its content
		err = ioutil.WriteFile(filePath, []byte{}, 0644)
		if err != nil {
			fmt.Println("Error clearing file:", err)
			return
		}
	} else if os.IsNotExist(err) {
		// If the file does not exist, create it
		_, err = os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	} else {
		// If there was some other error, print it and return
		fmt.Println("Error checking file:", err)
		return
	}

	// Write the new content to the file
	err := ioutil.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File written successfully")
}

func load_more(link string) *goquery.Document {
	req, _ := http.NewRequest("GET", link, nil)
	fmt.Println(link)
	resp, _ := client.Do(req)
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	return doc
}

func GetMessages(length int, doc *goquery.Document, number string, channel string) *goquery.Document {
	x := load_more(channel + "?before=" + number)

	html2, _ := x.Html()
	reader2 := strings.NewReader(html2)
	doc2, _ := goquery.NewDocumentFromReader(reader2)

	doc.Find("body").AppendSelection(doc2.Find("body").Children())

	newDoc := goquery.NewDocumentFromNode(doc.Selection.Nodes[0])

	messages := newDoc.Find(".tgme_widget_message_wrap").Length()

	fmt.Println(messages)
	if messages > length {
		return newDoc
	} else {
		num, _ := strconv.Atoi(number)
		n := num - 21
		if n > 0 {
			ns := strconv.Itoa(n)
			GetMessages(length, newDoc, ns, channel)
		} else {
			return newDoc
		}
	}

	return newDoc
}

func reverse(lines []string) []string {
	for i := 0; i < len(lines)/2; i++ {
		j := len(lines) - i - 1
		lines[i], lines[j] = lines[j], lines[i]
	}
	return lines
}

func RemoveDuplicate(config string) string {
	lines := strings.Split(config, "\n")

	// Use a map to keep track of unique lines
	uniqueLines := make(map[string]bool)

	// Loop over lines and add unique lines to map
	for _, line := range lines {
		if len(line) > 0 {
			uniqueLines[line] = true
		}
	}

	// Join unique lines into a string
	uniqueString := strings.Join(getKeys(uniqueLines), "\n")

	return uniqueString
}

func getKeys(m map[string]bool) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

