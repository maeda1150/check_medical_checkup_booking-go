package main

import (
	"github.com/sclevine/agouti"
	"log"
  "time"
)

func main() {

	log.Println("--- Create driver")
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			"--disable-gpu",
			"--window-size=1280,800",
			"--disable-application-cache",
			"--disable-infobars",
			"--no-sandbox",
			"--hide-scrollbars",
			"--enable-logging",
			"--log-level=0",
			"--v=99",
			"--single-process",
			"--homedir=/tmp",
			"--allow-insecure-localhost",
      "--enable-javascript",
		}),
		agouti.Debug,
	)
	if err := driver.Start(); err != nil {
		log.Fatal()
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("--- Access first page")
	err = page.Navigate("https://www.its-kenpo.or.jp/kanri/chokuei/yoyaku/index.html")
	if err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}
  page.Screenshot("Screenshot01.png")

  page.FindByLink("空き枠照会(個人予約)").Click()
  page.Screenshot("Screenshot02.png")

  err = page.FindByLink("同意する").Click()
  // err = page.FindByClass("wm_btn").Click()
	if err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}
  time.Sleep(2 * time.Second)

	// page = driver.JoinPage("https://sy.its-kenpo.or.jp/kenshinyoyaku/indiv_rsv_inquire")
  page.Screenshot("Screenshot03.png")

	page, err = driver.NewPage()
	err = page.Navigate("https://www.its-kenpo.or.jp/kanri/chokuei/yoyaku/index.html")
  page.Screenshot("Screenshot04.png")
  count, _ := page.WindowCount()
	log.Println(count)
}
