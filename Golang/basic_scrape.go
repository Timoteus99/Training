package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//func openbrowser(url string) {
//	var err error
//
//	switch runtime.GOOS {
//	case "linux":
//		err = exec.Command("xdg-open", url).Start()
//	case "windows":
//		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
//	case "darwin":
//		err = exec.Command("open", url).Start()
//	default:
//		err = fmt.Errorf("unsupported platform")
//	}
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}

//func gethtml(url string) []byte {
//	//fmt.Printf("HTML code of %s ...\n", url)
//	resp, err := http.Get(url)
//	// handle the error if there is one
//	if err != nil {
//		panic(err)
//	}
// do this now so it won't be forgotten
//	defer resp.Body.Close()
// reads html as a slice of bytes
//	html, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		panic(err)
//	}
// show the HTML code as a string %s
//fmt.Printf("%s\n", html)
//	return html
//}

//func BytesToString(data []byte) string {
//converts []byte object to string
//	return string(data[:])
//}

//func main() {

//openbrowser("https://meteo.arso.gov.si/met/sl/warning/")
//gethtml("http://weevil.info/")
//BytesToString(gethtml("http://weevil.info/"))

//str1 := BytesToString(gethtml("https://www.w3schools.com/html/html_tables.asp"))

//fmt.Println(str1)

//re := regexp.MustCompile(`<td.*?>(.*)</td>`)

//result := []string

//submatchall := re.FindAllStringSubmatch(str1, -1)

//for _, element := range submatchall {
//(element[1])
//}
//print all the "text" from html

//fmt.Println(submatchall)

// reading entire file into memory
//	data, err := ioutil.ReadFile("checkouthtml.txt")
//	if err != nil {
//		fmt.Println("File reading error", err)
//		return
//	}
//fmt.Println("Contents of file:", string(data))

//StringData := string(data)

//	re2 := regexp.MustCompile(`<td.*?>(.*)</td>`)

//result := []string

//	submatchall2 := re2.FindAllStringSubmatch(string(data), -1)

//fmt.Println(submatchall2[0][0])

//	if strings.Contains(submatchall2[0][0], "checkout114") {
//		fmt.Println("Found checkout114 in file!")
//	}

//	data = nil

//}

func gethtml(url string) []byte {
	//fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)
	return html
}

//BytesToString converter
func BytesToString(data []byte) string {
	//converts []byte object to string
	return string(data[:])
}

func getUrl() string {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: " + args[0] + " [URL]")
		return ""
	}
	url := args[1]
	return url
}

func goGet(url string) {
	var rows [][]string
	var row []string

	data := BytesToString(gethtml(url))

	// data := `<html><head><META http-equiv="Content-Type" content="text/html; charset=utf-8"></head><body>

	// <div>

	// <div></div>
	// <div><img src="http://./Kiosk_files/loading4.gif" height="78" width="78"></div>

	// <div></div>

	// <table>
	// <tbody><tr><td width="48%">
	// <table>
	// 	<tbody><tr>
	// 		<th width="5%">Pool</th>
	// 		<th width="5%">✓<br>In</th>
	// 		<th width="8%">Room</th>
	// 		<th width="5%">Stay<br>Over</th>
	// 		<th width="5%">Out</th>
	// 		<th width="5%">Strip</th>
	// 		<th width="5%">Clean</th>
	// 		<th width="5%">Done</th>
	// 		<th width="5%">Pkg<br>Done</th>
	// 		<th width="5%">Q.C.</th>
	// 		<th width="5%">Early<br>Late</th>
	// 		<th>Extras</th>
	// 	</tr>
	// 		<tr>
	// 		<td> </td>

	// 		<td>✓</td>

	// 		<td><a>110</a><input type="hidden" name="custname110" value="PARRA, RAFAEL"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>VR</td>
	// 		<td>L</td>
	// 		<td>
	// 			<span></span>
	// 			<span>Belamere Love Locks, Rose Petals (Personal Message), Aromatherapy (3 pack) </span>
	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>111</a><input type="hidden" name="custname111" value="CHAN, STACIE"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>VR</td>
	// 		<td>E L</td>
	// 		<td>
	// 			<span></span>
	// 			<span>Rose Petals (Heart), Candles (6) </span>
	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>112</a><input type="hidden" name="custname112" value="HICKS, JOHNATHAN"></td>

	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>OR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>
	// 			<span>NO SERVICE POOL AT 12:00P </span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td> </td>

	// 		<td>✓</td>

	// 		<td><a>114</a><input type="hidden" name="custname114" value="CARTER, EDWARD"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>VR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span>Candles (6), Rose Petals (Heart), Aromatherapy (3 pack) </span>
	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>115</a><input type="hidden" name="custname115" value="BELLO, PRESTON"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>VR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span>Candles (12), Belamere Love Locks, Rose Petals (Heart) </span>
	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>116</a><input type="hidden" name="custname116" value="THOMAS, JAMESON"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>L</td>
	// 		<td>
	// 			<span></span>
	// 			<span>Romance Package (Sparkling Juice) </span>
	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>117</a><input type="hidden" name="custname117" value="BLACKIE, MELANIE"></td>

	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>OR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>
	// 			<span>TOWEL SWAP AT 12:15P POOL AT 12:15P </span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>118</a><input type="hidden" name="custname118" value="VELAZQUEZ, TANIA"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>VR</td>
	// 		<td>L</td>
	// 		<td>
	// 			<span></span>
	// 			<span>Belamere Love Locks, Rose Petals (Heart) </span>
	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>119</a><input type="hidden" name="custname119" value="GRAYSON-EVANS, KHALAILAH"></td>

	// 		<td>✓</td>
	// 		<td></td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td>OR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>
	// 			<span>FULL SERVICE AT 10:30A POOL AT 10:30A CHANGE SHEETS - ONLY - GUEST WILL BE IN ROOM </span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>120</a><input type="hidden" name="custname120" value="KING, CANDICE"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>
	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>121</a><input type="hidden" name="custname121" value="CLAYBOURN, COURTNEY"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>
	// 			<span>PAINTING </span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>122</a><input type="hidden" name="custname122" value="TURNER, TYLER"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>L</td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>
	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>123</a><input type="hidden" name="custname123" value="ADAMS, DOMINIQUE"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>VR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span>Rose Petals (Heart) </span>
	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span>MAINTENANCE IN ROOM</span>
	// 		</td>
	// 	</tr>
	// 		<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>124</a><input type="hidden" name="custname124" value="SIMPSON, MEDIA"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>L</td>
	// 		<td>
	// 			<span></span>
	// 			<span>Candles (6), Rose Petals (Heart) </span>
	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>
	// 		</td>
	// 	</tr>
	// 	</tbody></table>
	// </td>
	// <td width="1%"> </td>
	// <td width="48%">
	// <table>
	// 	<tbody><tr>
	// 		<th width="5%">Pool</th>
	// 		<th width="5%">✓<br>In</th>
	// 		<th width="8%">Room</th>
	// 		<th width="5%">Stay<br>Over</th>
	// 		<th width="5%">Out</th>
	// 		<th width="5%">Strip</th>
	// 		<th width="5%">Clean</th>
	// 		<th width="5%">Done</th>
	// 		<th width="5%">Pkg<br>Done</th>
	// 		<th width="5%">Q.C.</th>
	// 		<th width="5%">Early<br>Late</th>
	// 		<th>Extras</th>
	// 	</tr>
	// 			<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>125</a><input type="hidden" name="custname125" value=""></td>

	// 		<td> </td>
	// 		<td></td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>VR</td>
	// 		<td>E </td>
	// 		<td>
	// 			<span></span>
	// 			<span>Anniversary Package (Tams Backstage), Midnight Snack Basket </span>

	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>126</a><input type="hidden" name="custname126" value="MUSE, QUADISHA"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td>VR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>

	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>127</a><input type="hidden" name="custname127" value="WHITESIDE, LAKESHA"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>VR</td>
	// 		<td>E </td>
	// 		<td>
	// 			<span></span>
	// 			<span>Anniversary Package (Tams Backstage) </span>

	// 			<span>PUMP IN POOL ROOM </span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>128</a><input type="hidden" name="custname128" value="CHARLES, LAKKESIA"></td>

	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>OR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>

	// 			<span>TOWEL SWAP AT 10:00A POOL AT 10:00A RING IN\OUT
	// LEAVE TOWELS ON BANNISTER </span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>129</a><input type="hidden" name="custname129" value="WILLIAMS, CARMEN"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>VR</td>
	// 		<td>E L</td>
	// 		<td>
	// 			<span></span>
	// 			<span>Anniversary Package (Tupelos) </span>

	// 			<span>WON&#39;T BE HERE UNTIL 5PM </span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>130</a><input type="hidden" name="custname130" value="KHAN, BRIDGETTE"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>VR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span>Rose Petals (Heart) </span>

	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>131</a><input type="hidden" name="custname131" value="LOWE, TYSHAWNDRA"></td>

	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>OR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>

	// 			<span>TOWEL SWAP AT 10:45A POOL AT 10:45A </span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>132</a><input type="hidden" name="custname132" value="THOMAS, ERIN"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>L</td>
	// 		<td>
	// 			<span></span>
	// 			<span>Candles (6), Rose Petals (Heart), Aromatherapy (3 pack) (2) </span>

	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>133</a><input type="hidden" name="custname133" value="ZELLARS/ROSS, SHERITA/LISA"></td>

	// 		<td>✓</td>
	// 		<td></td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td>OR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>

	// 			<span>FULL SERVICE AT 12:30P POOL AT 12:30P </span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>134</a><input type="hidden" name="custname134" value="GAMBLE, SHAUNTAYE"></td>

	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td>OR</td>
	// 		<td></td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>

	// 			<span>NO SERVICE POOL AT 10:15A WOULD LIKE TRASH BAGS   ALSO REQUESTS TO USE A BROOM AND DUSTPAN BUT DOES NOT WANT ANYONE TO CLEAN </span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>135</a><input type="hidden" name="custname135" value="PICKENS, MICHELLE"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>L</td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>

	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>136</a><input type="hidden" name="custname136" value="MARTIN, BARYN"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td></td>
	// 		<td></td>
	// 		<td>L</td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>

	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>
	// 				<tr>
	// 		<td>P</td>

	// 		<td>✓</td>

	// 		<td><a>137</a><input type="hidden" name="custname137" value="MCMILLIAN, KAYONNA"></td>

	// 		<td> </td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td>✓</td>
	// 		<td></td>
	// 		<td>VR</td>
	// 		<td>E </td>
	// 		<td>
	// 			<span></span>
	// 			<span> </span>

	// 			<span></span>

	// 			<span></span>
	// 			<span></span>
	// 			<span></span>

	// 		</td>
	// 	</tr>

	// </tbody></table>
	// </td>
	// </tr>
	// <tr>
	// 	<td colspan="3" style="height:55px">
	// <table width="100%" style="height:55px">
	// 	<tbody><tr>
	// 		<td style="height:55px" width="20%"><button type="button">   Main Menu  </button></td>
	// 		<td width="10%" align="center">Cumming, Georgia</td>
	// 		<td width="10%" align="center">8/20/2020, 3:27:31 PM</td>
	// 		<td width="40%" align="center">0 Checkouts Left (0 regular / 0 late) - 7 stayovers - 8 left to service</td>
	// 		<td width="20%" style="height:55px" align="right"><button type="button">   Breakfast  </button></td>
	// 	</tr>
	// </tbody></table>
	// 	</td>
	// 	</tr>
	// 	</tbody></table>

	// <div>
	// <form name="0.1_lookupform" target="_blank" onsubmit="try {return window.confirm(&quot;This form may not function properly due to certain security constraints.\nContinue?&quot;);} catch (e) {return false;}">
	// <input type="hidden" name="roomno" value="">
	// <table width="100%">
	// <tbody><tr>
	// 	<td colspan="2">Room: <span></span>     Guest: <span></span>   <span></span></td>
	// </tr>
	// <tr>
	// 	<td colspan="2" style="height:20px"><hr size="1"></td>
	// </tr>
	// <tr>
	// 	<td colspan="2" style="height:20px"> </td>
	// </tr>
	// <tr>
	// 	<td width="45%" valign="top">

	// <table width="100%">

	// <tbody><tr>
	// 	<td align="right"><input type="checkbox" name="es_checkout"><input type="hidden" name="es_checkout2" value="0"></td>
	// 	<td>Checked Out</td>
	// </tr>
	// <tr>
	// 	<td align="right"><input type="checkbox" name="es_calledout"></td>
	// 	<td>Called Out</td>
	// </tr>
	// <tr>
	// 	<td colspan="2" style="height:20px"> </td>
	// </tr>
	// <tr>
	// 	<td align="right"><input type="checkbox" name="es_pooldone"></td>
	// 	<td>Pool Done</td>
	// </tr>
	// <tr>
	// 	<td align="right"><input type="checkbox" name="es_blower"></td>
	// 	<td>Blower/Ionizer In Room</td>
	// </tr>
	// <tr>
	// 	<td align="right"><input type="checkbox" name="es_maint"></td>
	// 	<td>Maintenance In Room</td>
	// </tr>
	// <tr>
	// 	<td colspan="2" style="height:20px"> </td>
	// </tr>
	// <tr>
	// 	<td align="right"><input type="checkbox" name="es_strip"></td>
	// 	<td>Stripped</td>
	// </tr>
	// <tr>
	// 	<td align="right"><input type="checkbox" name="es_clean"></td>
	// 	<td>Clean</td>
	// </tr>
	// <tr>
	// 	<td align="right"><input type="checkbox" name="es_done"></td>
	// 	<td>Done</td>
	// </tr>
	// <tr>
	// 	<td align="right"><input type="checkbox" name="es_pkgdone"></td>
	// 	<td>Packages Done</td>
	// </tr>
	// <tr>
	// 	<td align="right"><input type="checkbox" name="es_candleslit"></td>
	// 	<td>Candles Lit</td>
	// </tr>
	// <tr>
	// 	<td align="right"><input type="checkbox" name="es_qc"></td>
	// 	<td>Quality Checked</td>
	// </tr>
	// <tr>
	// 	<td colspan="2" style="height:20px"> </td>
	// </tr>
	// <tr>
	// 	<td align="right"><input type="checkbox" name="es_tourroom"></td>
	// 	<td>Tour Room</td>
	// </tr>
	// </tbody></table>
	// </td>
	// <td width="55%" valign="top">
	// <table width="100%">
	// <tbody><tr>
	// 	<td align="right"><input type="checkbox" name="es_checkedin"><input type="hidden" name="es_checkedin2" value="0"></td>
	// 	<td>Checked In (Today)</td>
	// </tr>
	// <tr>
	// 	<td colspan="2" style="height:20px"> </td>
	// </tr>
	// <tr>
	// 	<td colspan="2">
	// 		<select name="es_servicetype"><option selected value="FULL">FULL SERVICE</option><option value="NO">NO SERVICE</option><option value="TOWEL">TOWEL SWAP</option></select> at
	// 		<select name="es_serviceat">
	// 			<option value="N/A">N/A</option>
	// 			<option value="READY">READY</option>
	// 			<option value="WILL CALL">WILL CALL</option>
	// 			<option value="WITH BREAKFAST">BREAKFAST</option>
	// 			<option value="9:00A">9:00A</option>
	// 			<option value="9:15A">9:15A</option>
	// 			<option value="9:30A">9:30A</option>
	// 			<option value="9:45A">9:45A</option>
	// 			<option value="10:00A">10:00A</option>
	// 			<option value="10:15A">10:15A</option>
	// 			<option value="10:30A">10:30A</option>
	// 			<option value="10:45A">10:45A</option>
	// 			<option value="11:00A">11:00A</option>
	// 			<option value="11:15A">11:15A</option>
	// 			<option value="11:30A">11:30A</option>
	// 			<option value="11:45A">11:45A</option>
	// 			<option value="12:00P">12:00P</option>
	// 			<option value="12:15P">12:15P</option>
	// 			<option value="12:30P">12:30P</option>
	// 			<option value="12:45P">12:45P</option>
	// 			<option value="1:00P">1:00P</option>
	// 			<option value="1:15P">1:15P</option>
	// 			<option value="1:30P">1:30P</option>
	// 			<option value="1:45P">1:45P</option>
	// 			<option value="2:00P">2:00P</option>
	// 			<option value="2:15P">2:15P</option>
	// 			<option value="2:30P">2:30P</option>
	// 			<option value="2:45P">2:45P</option>
	// 			<option value="3:00P">3:00P</option>
	// 		</select>
	// 	</td>

	// </tr>
	// <tr>
	// 	<td colspan="2" style="height:20px"> </td>
	// </tr>
	// <tr>
	// 	<td colspan="2">
	// 		Pool Service at
	// 		<select name="es_poolat">
	// 			<option value="N/A">N/A</option>
	// 			<option value="READY">READY</option>
	// 			<option value="WILL CALL">WILL CALL</option>
	// 			<option value="WITH BREAKFAST">BREAKFAST</option>
	// 			<option value="9:00A">9:00A</option>
	// 			<option value="9:15A">9:15A</option>
	// 			<option value="9:30A">9:30A</option>
	// 			<option value="9:45A">9:45A</option>
	// 			<option value="10:00A">10:00A</option>
	// 			<option value="10:15A">10:15A</option>
	// 			<option value="10:30A">10:30A</option>
	// 			<option value="10:45A">10:45A</option>
	// 			<option value="11:00A">11:00A</option>
	// 			<option value="11:15A">11:15A</option>
	// 			<option value="11:30A">11:30A</option>
	// 			<option value="11:45A">11:45A</option>
	// 			<option value="12:00P">12:00P</option>
	// 			<option value="12:15P">12:15P</option>
	// 			<option value="12:30P">12:30P</option>
	// 			<option value="12:45P">12:45P</option>
	// 			<option value="1:00P">1:00P</option>
	// 			<option value="1:15P">1:15P</option>
	// 			<option value="1:30P">1:30P</option>
	// 			<option value="1:45P">1:45P</option>
	// 			<option value="2:00P">2:00P</option>
	// 			<option value="2:15P">2:15P</option>
	// 			<option value="2:30P">2:30P</option>
	// 			<option value="2:45P">2:45P</option>
	// 			<option value="3:00P">3:00P</option>
	// 		</select>
	// 	</td>

	// </tr>
	// <tr>
	// 	<td colspan="2" style="height:20px"></td>
	// </tr>
	// <tr>
	// 	<td colspan="2" align="right"><button style="border-radius:8px" type="button">Clear Service</button></td>
	// </tr>
	// <tr>
	// 	<td colspan="2" style="height:20px"></td>
	// </tr>
	// <tr>
	// 	<td colspan="2">Notes:<br><textarea name="es_servicenotes" style="width:100%;height:130px;text-transform:uppercase"></textarea></td>
	// </tr>
	// </tbody></table>
	// </td>
	// </tr>
	// </tbody></table>

	// </form>

	// <div style="height:5px"> </div>

	// <table width="100%">
	// 	<tbody><tr>
	// 		<td><button type="button" style="width:150px">Cancel</button></td>
	// 		<td align="right"><button type="button" style="width:150px">Save</button></td>
	// 	</tr>
	// </tbody></table>

	// </div>

	// </div></body></html>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}

	// Find each table
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				row = append(row, tablecell.Text())
				//fmt.Println(tablecell.Text())
			})
			rows = append(rows, row)
			//fmt.Println(len(row))

			row = nil
		})
	})
	var results []string
	//fmt.Println(rows[5]) // TUKI JE FORA!!!!
	for i, element := range rows {
		//fmt.Println(len(element))
		if len(element) == 12 && i < 33 {
			//fmt.Println(element)
			if element[4] == "✓" {
				//fmt.Println("Checkout found!")
				results = append(results, element[2])
			}
		}
	}
	fmt.Println(results)
	//fmt.Println("####### headings = ", len(headings), headings)
	//fmt.Println("####### rows = ", len(rows), rows)
	//fmt.Println(data)

	//var results []string
	//doc.Find("tr").Each(func(i int, s *goquery.Selection) {
	//fmt.Printf("Content of cell %d: %s\n", i, s.Text())
	//	results = append(results, s.Text())
	//})

}

func main() {
	start := time.Now()
	url := getUrl()
	if url == "" {
		return
	}
	goGet(url)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
