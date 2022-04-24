package handler

import (
	"encoding/json"
	"net/http"
	"scraper_trendyol/data_collector"
	"scraper_trendyol/excel_parser"
	"scraper_trendyol/pkg/helper"
	"scraper_trendyol/pkg/logging"
	"strconv"
	"sync/atomic"

	"github.com/sirupsen/logrus"
)
 
// type Ty_handler struct {
// 	// couch database client
// 	tt excel_parser.ExcelParser
// }

const (
	stateUnlocked uint32 = iota
	stateLocked
)

var (
	locker = stateUnlocked
	// tt     excel_parser.ExcelParser
)

func InitScraper(w http.ResponseWriter, r *http.Request) {

	// lock the request
	if !atomic.CompareAndSwapUint32(&locker, stateUnlocked, stateLocked) {
		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Scrape in progress!",
		})

		return
	}
	defer atomic.StoreUint32(&locker, stateUnlocked)

	keys := r.URL.Query()["product-limit"]

	productLimitStr := keys[0]

	logrus.Infoln("InitTrendyolScraper")
	logrus.Infoln("Total product limit: ", productLimitStr)

	productLimit, _ := strconv.Atoi(productLimitStr)

	helper.TotalProductLimit = productLimit
	helper.InsertedProductCount = 0

	bagisto, err := data_collector.NewBagisto()
	if err != nil {
		logging.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rootCategory, err := bagisto.GetRootCategory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scraper, err := data_collector.NewScraper()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scraper.BeginCollectingData(rootCategory)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg": "Trendyol scraper completed successfully!",
	})
}

func ParseLink(w http.ResponseWriter, r *http.Request) {

	link := r.URL.Query().Get("url")

	logrus.Info("link: ", link)

	linkParser := data_collector.NewLinkParser(link)
	productGroupId, err := linkParser.ParseLink()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg":            "Link parsed successfully",
		"productGroupId": strconv.Itoa(productGroupId),
	})
}

func ParseExcel(w http.ResponseWriter, r *http.Request) {
	ep, err := excel_parser.NewExcelParser()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": err.Error(),
		})

		return
	}

	err = ep.ParseExcelAndInsert()

	msg := "categories updated successfully"

	if err != nil {
		msg = err.Error()
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg": msg,
	})
}

func InitUpdater(w http.ResponseWriter, r *http.Request) {

	updater, err := data_collector.NewUpdater()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": err.Error(),
		})

		return
	}

	errUpdater := updater.InitUpdater()

	if errUpdater != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"msg": "updated products",
	})
}

// ----------------------------------------------------------------------------------------------------------------

func GetExcel(w http.ResponseWriter, r *http.Request) {

	// tt.GetExcelData()
	ep, err := excel_parser.NewExcelParser()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": err.Error(),
		})

		return
	}

	err = ep.GetExcelData()

	msg := "categories updated successfully"

	if err != nil {
		msg = err.Error()
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg": msg,
	})

}
