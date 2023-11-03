package proc

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gonum/stat"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"math"
	"regexp"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/module/elem/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/odds/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/vo"
	"time"
)

type EuroLastProcesser struct {
	service.CompService
	service2.EuroLastService
	service2.EuroHisService
	//入参
	//是否是单线程
	SingleThread  bool
	MatchLastList []*pojo.MatchLast
	//博彩公司对应的win007id
	CompWin007Ids      []string
	Win007idMatchidMap map[string]string
	MinuteDiff int
}

func GetEuroLastProcesser() *EuroLastProcesser {
	processer := &EuroLastProcesser{}
	processer.Init()
	return processer
}

func (this *EuroLastProcesser) Init() {
	//初始化参数值
	this.Win007idMatchidMap = map[string]string{}
}

func (this *EuroLastProcesser) Setup(temp *EuroLastProcesser) {
	//设置参数值
	this.CompWin007Ids = temp.CompWin007Ids
}

func (this *EuroLastProcesser) Startup() {

	var newSpider *spider.Spider
	processer := this
	newSpider = spider.NewSpider(processer, "EuroLastProcesser")
	if len(this.MatchLastList) == 0 {
		return
	}

	for i, v := range this.MatchLastList {
		if !this.SingleThread && i%1000 == 0 { //10000个比赛一个spider,一个赛季大概有30万场比赛,最多30spider
			//先将前面的spider启动
			newSpider.SetDownloader(down.NewMWin007Downloader())
			newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
			newSpider.SetSleepTime("rand", win007.SLEEP_RAND_S, win007.SLEEP_RAND_E)
			newSpider.SetThreadnum(1).Run()

			processer = GetEuroLastProcesser()
			processer.Setup(this)
			newSpider = spider.NewSpider(processer, "EuroLastProcesser"+strconv.Itoa(i))
		}

		temp_flag := v.Ext[win007.MODULE_FLAG]
		bytes, _ := json.Marshal(temp_flag)
		matchExt := new(pojo.MatchExt)
		json.Unmarshal(bytes, matchExt)
		win007_id := matchExt.Sid

		processer.Win007idMatchidMap[win007_id] = v.Id

		if win007_id == "" {
			win007_id = v.Id
		}
		//win007_id = "2322026"
		url := strings.Replace(win007.WIN007_EUROODD_URL_PATTERN, "${matchId}", win007_id, 1)
		newSpider = newSpider.AddUrl(url, "html")
	}

	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetSleepTime("rand", win007.SLEEP_RAND_S, win007.SLEEP_RAND_E)
	newSpider.SetThreadnum(1).Run()

}

func (this *EuroLastProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Error("URL:", request.Url, p.Errormsg())
		return
	}

	var hdata_str string
	p.GetHtmlParser().Find("script").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if hdata_str == "" && strings.Contains(text, "var hData") {
			hdata_str = text
		} else {
			return
		}
	})
	if hdata_str == "" {
		base.Log.Error("hdata_str:为空,URL:", request.Url)
		return
	}

	//base.Log.Info("hdata_str", hdata_str, "URL:", request.Url)
	// 获取script脚本中的，博彩公司信息
	hdata_str = strings.Replace(hdata_str, ";", "", 1)
	hdata_str = strings.Replace(hdata_str, "var hData = ", "", 1)
	if hdata_str == "" {
		base.Log.Info("hdata_str:解析失败,", hdata_str, "URL:", request.Url)
		return
	}
	this.hdata_process(request.Url, hdata_str)
}

func (this *EuroLastProcesser) hdata_process(url string, hdata_str string) {

	var hdata_list = make([]*vo.HData, 0)
	json.Unmarshal(([]byte)(hdata_str), &hdata_list)
	var regex_temp = regexp.MustCompile(`(\d+).htm`)
	win007Id := strings.Split(regex_temp.FindString(url), ".")[0]
	matchId := this.Win007idMatchidMap[win007Id]
	if matchId == "" {
		matchId = win007Id
	}

	//入库中
	//comp_list_slice := make([]interface{}, 0)
	//last_slice := make([]interface{}, 0)
	//last_update_slice := make([]interface{}, 0)
	var main_10 [10]float64
	var main_9 [9]float64
	var guest_10 [10]float64
	var guest_9 [9]float64
	var middle_10 [10]float64
	var middle_9 [9]float64
	//var CoreMainNorm float64
	//var CoreGuestNorm float64
	//var CoreMiddleNorm float64
	var CoreMainP float64
	var CoreGuestP float64
	var CoreMiddleP float64

	var IntMainP float64
	var IntGuestP float64
	var IntMiddleP float64

	var PrinMainP float64
	var PrinGuestP float64
	var PrinMiddleP float64

	var B365MainP float64
	var B365GuestP float64
	var B365MiddleP float64
	weights := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, v := range hdata_list {
		//comp := new(entity2.Comp)
		//comp.Name = v.Cn
		//comp.Type = 1
		//comp_exists := this.CompService.Exist(comp)
		//if !comp_exists {
		//	//comp.Id = bson.NewObjectId().Hex()
		//	comp.Id = strconv.Itoa(v.CId)
		//	comp_list_slice = append(comp_list_slice, comp)
		//}

		//判断公司ID是否在配置的波菜公司队列中
		if len(this.CompWin007Ids) > 0 {
			var equal bool
			for _, id := range this.CompWin007Ids {
				if strings.EqualFold(id, strconv.Itoa(v.CId)) {
					equal = true
					break
				}
			}
			if !equal {
				continue
			}
		}

		last := new(entity3.EuroLast)
		last.MatchId = matchId
		//last.CompId, _ = strconv.Atoi(comp.Id)
		last.CompId = v.CId

		last.Sp3 = v.Hw
		last.Sp1 = v.So
		last.Sp0 = v.Gw
		last.Ep3 = v.Rh
		last.Ep1 = v.Rs
		last.Ep0 = v.Rg
		//crown_545、li_82、weilian_225、weide_81、yi_90、li_474、12_18、bwin_255、365_281、prinn_177
		if v.CId == 545 {
			main_10[0] = last.Ep3
			main_9[0] = last.Ep3
			guest_10[0] = last.Ep0
			guest_9[0] = last.Ep0
			middle_10[0] = last.Ep1
			middle_9[0] = last.Ep1
		} else if v.CId == 82 {
			main_10[1] = last.Ep3
			main_9[1] = last.Ep3
			guest_10[1] = last.Ep0
			guest_9[1] = last.Ep0
			middle_10[1] = last.Ep1
			middle_9[1] = last.Ep1
		} else if v.CId == 115 {
			main_10[2] = last.Ep3
			main_9[2] = last.Ep3
			guest_10[2] = last.Ep0
			guest_9[2] = last.Ep0
			middle_10[2] = last.Ep1
			middle_9[2] = last.Ep1
		} else if v.CId == 81 {
			main_10[3] = last.Ep3
			main_9[3] = last.Ep3
			guest_10[3] = last.Ep0
			guest_9[3] = last.Ep0
			middle_10[3] = last.Ep1
			middle_9[3] = last.Ep1
		} else if v.CId == 90 {
			main_10[4] = last.Ep3
			main_9[4] = last.Ep3
			guest_10[4] = last.Ep0
			guest_9[4] = last.Ep0
			middle_10[4] = last.Ep1
			middle_9[4] = last.Ep1
		} else if v.CId == 659 {
			//yingli
			main_10[5] = last.Ep3
			main_9[5] = last.Ep3
			guest_10[5] = last.Ep0
			guest_9[5] = last.Ep0
			middle_10[5] = last.Ep1
			middle_9[5] = last.Ep1
		} else if v.CId == 18 {
			//12bet
			main_10[6] = last.Ep3
			main_9[6] = last.Ep3
			guest_10[6] = last.Ep0
			guest_9[6] = last.Ep0
			middle_10[6] = last.Ep1
			middle_9[6] = last.Ep1
			CoreMainP = last.Ep3
			CoreGuestP = last.Ep0
			CoreMiddleP = last.Ep1
		} else if v.CId == 255 {
			main_10[7] = last.Ep3
			main_9[7] = last.Ep3
			guest_10[7] = last.Ep0
			guest_9[7] = last.Ep0
			middle_10[7] = last.Ep1
			middle_9[7] = last.Ep1
			B365MainP = last.Ep3  //bwin
			B365GuestP = last.Ep0
			B365MiddleP = last.Ep1
		} else if v.CId == 474 { //liji
			main_10[8] = last.Ep3
			main_9[8] = last.Ep3
			guest_10[8] = last.Ep0
			guest_9[8] = last.Ep0
			middle_10[8] = last.Ep1
			middle_9[8] = last.Ep1
		} else if v.CId == 177 {
			main_10[9] = last.Ep3
			guest_10[9] = last.Ep0
			middle_10[9] = last.Ep1
			PrinMainP = last.Ep3
			PrinGuestP = last.Ep0
			PrinMiddleP = last.Ep1
		} else if v.CId == 281 {
			IntMainP = last.Ep3
			IntGuestP = last.Ep0
			IntMiddleP = last.Ep1
		}else if v.CId == 2 {
			//B365MainP = last.Ep3
			//B365GuestP = last.Ep0
			//B365MiddleP = last.Ep1
		}
		//last_slice = append(last_slice, last)

		//last_temp_id, last_exists := this.EuroLastService.Exist(last)
		//if !last_exists {
		//	last_slice = append(last_slice, last)
		//} else {
		//	last.Id = last_temp_id
		//	last_update_slice = append(last_update_slice, last)
		//}
	}

	if main_10[6] < 1.75 || guest_10[6] < 1.75 || middle_10[6] < 1.75 {
		his := new(pojo.MatchHis)
		last := new(pojo.MatchLast)
		his.Id = matchId
		last.Id = matchId
		this.EuroLastService.Del(last)
		this.EuroHisService.Del(his)
		base.Log.Info("del pay 小于1.75 2222222")
		return
	}
	//this.CompService.SaveList(comp_list_slice)
	//最后数据
	average_main_10 := mean(main_10)
	std_main_10 := stdDev(main_10, average_main_10)
	average_guest_10 := mean(guest_10)
	std_guest_10 := stdDev(guest_10, average_guest_10)
	average_middle_10 := mean(middle_10)
	std_middle_10 := stdDev(middle_10, average_middle_10)

	//average_main_9 := mean9(main_9)
	//std_main_9 := stdDev9(main_9, average_main_9)
	//average_guest_9 := mean9(guest_9)
	//std_guest_9 := stdDev9(guest_9, average_guest_9)
	//average_middle_9 := mean9(middle_9)
	//std_middle_9 := stdDev9(middle_9, average_middle_9)

	normdist_12bet_main_10 := decimal(normdist(main_10[6], average_main_10, std_main_10))
	//normdist_12bet_main_9 := decimal(normdist(main_9[6], average_main_9, std_main_9))

	normdist_12bet_guest_10 := decimal(normdist(guest_10[6], average_guest_10, std_guest_10))
	//normdist_12bet_guest_9 := decimal(normdist(guest_9[6], average_guest_9, std_guest_9))

	normdist_12bet_middle_10 := decimal(normdist(middle_10[6], average_middle_10, std_middle_10))
	//normdist_12bet_middle_9 := decimal(normdist(middle_9[6], average_middle_9, std_middle_9))

	base.Log.Info(" normdist_12bet_main_10: ", normdist_12bet_main_10, " normdist_12bet_guest_10:", normdist_12bet_guest_10, " normdist_12bet_middle_10:", normdist_12bet_middle_10)
	//base.Log.Info(" normdist_12bet_main_9: ", normdist_12bet_main_9, " normdist_12bet_guest_9:", normdist_12bet_guest_9, " normdist_12bet_middle_9:", normdist_12bet_middle_9)

	norm := new(entity3.Norm)
	norm.MatchId = matchId
	norm.CompId = 18
	norm.Ep3 = main_10[6]
	norm.Ep0 = guest_10[6]
	norm.Ep1 = middle_10[6]
	for _, v1 := range this.MatchLastList {
		if v1.Id == matchId {
			norm.MainName = v1.MainTeamId
			norm.GuestName = v1.GuestTeamId
			norm.MatchTime = v1.MatchDate
			break
		}
		continue
	}

	//baseMainMaxValue1 := decimal3(average_main_10 + 1.96 * std_main_10)
	//baseMainMaxValue2 := decimal3(average_main_10 + 2.58 * std_main_10)
	//baseMiddleMaxValue1 := decimal3(average_middle_10 + 1.96  * std_middle_10)
	//baseMiddleMaxValue2 := decimal3(average_middle_10 + 2.58 * std_middle_10)
	//baseGuestMaxValue1 := decimal3(average_guest_10 + 1.96  * std_guest_10)
	//baseGuestMaxValue2 := decimal3(average_guest_10 + 2.58 * std_guest_10)

	var mainFirstMax float64
	var mainSecondMax float64
	var middleFirstMax float64
	var middleSecondMax float64
	var guestFirstMax float64
	var guestSecondMax float64
	var firstFlag int
	var secondFlag int
	mainFirstMax = 0
	mainSecondMax = 0
	middleFirstMax = 0
	middleSecondMax = 0
	guestFirstMax = 0
	guestSecondMax = 0
	firstFlag = 0
	secondFlag = 0
	count := 0
	for flag, value := range main_10 {
		if value > 0 {
			count = count + 1
		}
		if(mainFirstMax < value){
			mainFirstMax = value
			firstFlag = flag
		}
	}
	for flag, value := range main_10 {
		if(mainSecondMax < value && value != mainFirstMax){
			mainSecondMax = value
			secondFlag = flag
		}
	}

	if(count<10){
		his := new(pojo.MatchHis)
		last := new(pojo.MatchLast)
		his.Id = matchId
		last.Id = matchId
		this.EuroLastService.Del(last)
		this.EuroHisService.Del(his)
		base.Log.Info("del count 小于10 11111111111")
		return
	}
	//if(mainFirstMax < IntMainP){
	//	mainFirstMax = IntMainP
	//	firstFlag = 10
	//}

	mainMaxNameString := ""
	if(firstFlag == 0){
		mainMaxNameString = mainMaxNameString + "12bet_"
	}else if(firstFlag == 1){
		mainMaxNameString = mainMaxNameString + "libo_"
	}else if(firstFlag == 2){
		mainMaxNameString = mainMaxNameString + "weilian_"
	}else if(firstFlag == 3){
		mainMaxNameString = mainMaxNameString + "weide_"
	}else if(firstFlag == 4){
		mainMaxNameString = mainMaxNameString + "yishibo_"
	}else if(firstFlag == 5){//474
		mainMaxNameString = mainMaxNameString + "yingli_"
	}else if(firstFlag == 6){//545
		mainMaxNameString = mainMaxNameString + "crown_"
	}else if(firstFlag == 7){//255
		mainMaxNameString = mainMaxNameString + "bwin_"
	}else if(firstFlag == 8){//104
		mainMaxNameString = mainMaxNameString + "int_"
	}else if(firstFlag == 9){//177
		mainMaxNameString = mainMaxNameString + "prinn_"
	}
	//else if(firstFlag == 10){
	//	mainMaxNameString = mainMaxNameString + "365_"
	//}

	if(secondFlag == 0){
		mainMaxNameString = mainMaxNameString + "12bet"
	}else if(secondFlag == 1){
		mainMaxNameString = mainMaxNameString + "libo"
	}else if(secondFlag == 2){
		mainMaxNameString = mainMaxNameString + "weilian"
	}else if(secondFlag == 3){
		mainMaxNameString = mainMaxNameString + "weide"
	}else if(secondFlag == 4){
		mainMaxNameString = mainMaxNameString + "yishibo"
	}else if(secondFlag == 5){//474
		mainMaxNameString = mainMaxNameString + "yingli"
	}else if(secondFlag == 6){//545
		mainMaxNameString = mainMaxNameString + "crown"
	}else if(secondFlag == 7){//255
		mainMaxNameString = mainMaxNameString + "bwin"
	}else if(secondFlag == 8){//104
		mainMaxNameString = mainMaxNameString + "int"
	}else if(secondFlag == 9){//177
		mainMaxNameString = mainMaxNameString + "prinn"
	}


	firstFlag = 0
	secondFlag = 0
	for flag, value := range middle_10 {
		if(middleFirstMax < value){
			middleFirstMax = value
			firstFlag = flag
		}
	}
	for flag, value := range middle_10 {
		if(middleSecondMax < value && value != middleFirstMax){
			middleSecondMax = value
			secondFlag = flag
		}
	}
	//if(middleFirstMax < IntMiddleP){
	//	middleFirstMax = IntMiddleP
	//	firstFlag = 10
	//}

	middleMaxNameString := ""
	if(firstFlag == 0){
		middleMaxNameString = middleMaxNameString + "12bet_"
	}else if(firstFlag == 1){
		middleMaxNameString = middleMaxNameString + "libo_"
	}else if(firstFlag == 2){
		middleMaxNameString = middleMaxNameString + "weilian_"
	}else if(firstFlag == 3){
		middleMaxNameString = middleMaxNameString + "weide_"
	}else if(firstFlag == 4){
		middleMaxNameString = middleMaxNameString + "yishibo_"
	}else if(firstFlag == 5){
		middleMaxNameString = middleMaxNameString + "yingli_"
	}else if(firstFlag == 6){//545
		middleMaxNameString = middleMaxNameString + "crown_"
	}else if(firstFlag == 7){//255
		middleMaxNameString = middleMaxNameString + "bwin_"
	}else if(firstFlag == 8){//104
		middleMaxNameString = middleMaxNameString + "int_"
	}else if(firstFlag == 9){//177
		middleMaxNameString = middleMaxNameString + "prinn_"
	}
	//else if(firstFlag == 10){
	//	middleMaxNameString = middleMaxNameString + "365_"
	//}

	if(secondFlag == 0){
		middleMaxNameString = middleMaxNameString + "12bet"
	}else if(secondFlag == 1){
		middleMaxNameString = middleMaxNameString + "libo"
	}else if(secondFlag == 2){
		middleMaxNameString = middleMaxNameString + "weilian"
	}else if(secondFlag == 3){
		middleMaxNameString = middleMaxNameString + "weide"
	}else if(secondFlag == 4){
		middleMaxNameString = middleMaxNameString + "yishibo"
	}else if(secondFlag == 5){
		middleMaxNameString = middleMaxNameString + "yingli"
	}else if(secondFlag == 6){//545
		middleMaxNameString = middleMaxNameString + "crown"
	}else if(secondFlag == 7){//255
		middleMaxNameString = middleMaxNameString + "bwin"
	}else if(secondFlag == 8){
		middleMaxNameString = middleMaxNameString + "int"
	}else if(secondFlag == 9){//177
		middleMaxNameString = middleMaxNameString + "prinn"
	}


	firstFlag = 0
	secondFlag = 0
	for flag, value := range guest_10 {
		if(guestFirstMax < value){
			guestFirstMax = value
			firstFlag = flag
		}
	}
	for flag, value := range guest_10 {
		if(guestSecondMax < value && value != guestFirstMax){
			guestSecondMax = value
			secondFlag = flag
		}
	}

	//if(guestFirstMax < IntGuestP){
	//	guestFirstMax = IntGuestP
	//	firstFlag = 10
	//}

	guestMaxNameString := ""
	if(firstFlag == 0){
		guestMaxNameString = guestMaxNameString + "12bet_"
	}else if(firstFlag == 1){
		guestMaxNameString = guestMaxNameString + "libo_"
	}else if(firstFlag == 2){
		guestMaxNameString = guestMaxNameString + "weilian_"
	}else if(firstFlag == 3){
		guestMaxNameString = guestMaxNameString + "weide_"
	}else if(firstFlag == 4){
		guestMaxNameString = guestMaxNameString + "yishibo_"
	}else if(firstFlag == 5){//474
		guestMaxNameString = guestMaxNameString + "yingli_"
	}else if(firstFlag == 6){//545
		guestMaxNameString = guestMaxNameString + "crown_"
	}else if(firstFlag == 7){//255
		guestMaxNameString = guestMaxNameString + "bwin_"
	}else if(firstFlag == 8){//281
		guestMaxNameString = guestMaxNameString + "int_"
	}else if(firstFlag == 9){//177
		guestMaxNameString = guestMaxNameString + "prinn_"
	}
	//else if(firstFlag == 10){
	//	guestMaxNameString = guestMaxNameString + "365_"
	//}

	if(secondFlag == 0){
		guestMaxNameString = guestMaxNameString + "12bet"
	}else if(secondFlag == 1){
		guestMaxNameString = guestMaxNameString + "libo"
	}else if(secondFlag == 2){
		guestMaxNameString = guestMaxNameString + "weilian"
	}else if(secondFlag == 3){
		guestMaxNameString = guestMaxNameString + "weide"
	}else if(secondFlag == 4){
		guestMaxNameString = guestMaxNameString + "yishibo"
	}else if(secondFlag == 5){//474
		guestMaxNameString = guestMaxNameString + "yingli"
	}else if(secondFlag == 6){//545
		guestMaxNameString = guestMaxNameString + "crown"
	}else if(secondFlag == 7){//255
		guestMaxNameString = guestMaxNameString + "bwin"
	}else if(secondFlag == 8){//281
		guestMaxNameString = guestMaxNameString + "int"
	}else if(secondFlag == 9){//177
		guestMaxNameString = guestMaxNameString + "prinn"
	}


	//baseMainMaxValue1 := decimal3(normdist(mainFirstMax, average_main_10, std_main_10))
	////baseMainMaxValue2 := decimal3(normdist(mainSecondMax, average_main_10, std_main_10))
	//baseMainMaxValue2 := decimal3(std_main_10)
	//baseMiddleMaxValue1 := decimal3(normdist(middleFirstMax, average_middle_10, std_middle_10))
	////baseMiddleMaxValue2 := decimal3(normdist(middleSecondMax, average_middle_10, std_middle_10))
	//baseMiddleMaxValue2 := decimal3(std_middle_10)
	//baseGuestMaxValue1 := decimal3(normdist(guestFirstMax, average_guest_10, std_guest_10))
	////baseGuestMaxValue2 := decimal3(normdist(guestSecondMax, average_guest_10, std_guest_10))
	//baseGuestMaxValue2 := decimal3(std_guest_10)

	mains := []float64{main_10[0], main_10[1], main_10[2], main_10[3], main_10[4], main_10[5], main_10[6], main_10[7], main_10[8], main_10[9]}
	middles := []float64{middle_10[0], middle_10[1], middle_10[2], middle_10[3], middle_10[4], middle_10[5], middle_10[6], middle_10[7], middle_10[8], middle_10[9]}
	guests := []float64{guest_10[0], guest_10[1], guest_10[2], guest_10[3], guest_10[4], guest_10[5], guest_10[6], guest_10[7], guest_10[8], guest_10[9]}

	baseMainMaxValue1 := decimal3(stat.Skew(mains,weights))
	baseMiddleMaxValue1 := decimal3(stat.Skew(middles,weights))
	baseGuestMaxValue1 := decimal3(stat.Skew(guests,weights))

	baseMainMaxValue2 := decimal3(stat.ExKurtosis(mains,weights))
	baseMiddleMaxValue2 := decimal3(stat.ExKurtosis(middles,weights))
	baseGuestMaxValue2 := decimal3(stat.ExKurtosis(guests,weights))

	norm.MainMaxName = mainMaxNameString
	norm.MiddleMaxName = middleMaxNameString
	norm.GuestMaxName = guestMaxNameString
	norm.BaseMainMaxValue1 = baseMainMaxValue1
	norm.BaseMainMaxValue2 = baseMainMaxValue2
	norm.BaseMiddleMaxValue1 = baseMiddleMaxValue1
	norm.BaseMiddleMaxValue2 = baseMiddleMaxValue2
	norm.BaseGuestMaxValue1 = baseGuestMaxValue1
	norm.BaseGuestMaxValue2 =baseGuestMaxValue2
	norm.MainFirstMax = mainFirstMax
	norm.MainSecondMax = mainSecondMax
	norm.MiddleFirstMax = middleFirstMax
	norm.MiddleSecondMax = middleSecondMax
	norm.GuestFirstMax = guestFirstMax
	norm.GuestSecondMax = guestSecondMax


	//norm.MainMaxAddStd = decimal3(mainFirstMax + 2 * std_main_10)
	//norm.MiddleMaxAddStd = decimal3(middleFirstMax + 2 * std_middle_10)
	//norm.GuestMaxAddStd = decimal3(guestFirstMax + 2 * std_guest_10)
	//norm.MainMeanAdd3Std = decimal3(average_main_10 + 3 * std_main_10)
	//norm.MiddleMeanAdd3Std = decimal3(average_middle_10 + 3 * std_middle_10)
	//norm.GuestMeanAdd3Std = decimal3(average_guest_10 + 3 * std_guest_10)

	norm.MatchTime = time.Now()
	norm.CompCount = count
	norm.CreateTime = time.Now()
	norm.OddTime = time.Now()
	norm.Main10Norm = normdist_12bet_main_10
	//norm.Main9Norm = normdist_12bet_main_9
	norm.Guest10Norm = normdist_12bet_guest_10
	//norm.Guest9Norm = normdist_12bet_guest_9
	norm.Middle10Norm = normdist_12bet_middle_10
	norm.DensityMain3 = decimal(normdistfalse(main_10[6], average_main_10, std_main_10))
	norm.DensityGuest3 = decimal(normdistfalse(guest_10[6], average_guest_10, std_guest_10))
	norm.DensityMiddle3 = decimal(normdistfalse(middle_10[6], average_middle_10, std_middle_10))
	//norm.Middle9Norm = normdist_12bet_middle_9
	if CoreMainP > 0 {
		norm.CoreMainNorm = decimal(normdist(CoreMainP, average_main_10, std_main_10))
		norm.CoreGuestNorm = decimal(normdist(CoreGuestP, average_guest_10, std_guest_10))
		norm.CoreMiddleNorm = decimal(normdist(CoreMiddleP, average_middle_10, std_middle_10))
		norm.CoreMainP = CoreMainP
		norm.CoreGuestP = CoreGuestP
		norm.CoreMiddleP = CoreMiddleP
	}
	if IntMainP > 0 {
		norm.IntMainP = IntMainP
		norm.IntGuestP = IntGuestP
		norm.IntMiddleP = IntMiddleP
		norm.Main9Norm = decimal(normdist(IntMainP, average_main_10, std_main_10))
		norm.Guest9Norm = decimal(normdist(IntGuestP, average_guest_10, std_guest_10))
		norm.Middle9Norm = decimal(normdist(IntMiddleP, average_middle_10, std_middle_10))
	}
	if PrinMainP > 0 {
		norm.PrinMainP = PrinMainP
		norm.PrinGuestP = PrinGuestP
		norm.PrinMiddleP = PrinMiddleP
		norm.PrinMainNorm = decimal(normdist(PrinMainP, average_main_10, std_main_10))
		norm.PrinGuestNorm = decimal(normdist(PrinGuestP, average_guest_10, std_guest_10))
		norm.PrinMiddleNorm = decimal(normdist(PrinMiddleP, average_middle_10, std_middle_10))

		norm.DensityMain1 = decimal(normdistfalse(PrinMainP, average_main_10, std_main_10))
		norm.DensityGuest1 = decimal(normdistfalse(PrinGuestP, average_guest_10, std_guest_10))
		norm.DensityMiddle1 = decimal(normdistfalse(PrinMiddleP, average_middle_10, std_middle_10))
	}

	if B365MainP > 0 {
		norm.B365MainP = B365MainP
		norm.B365GuestP = B365GuestP
		norm.B365MiddleP = B365MiddleP
		norm.B365MainNorm = decimal(normdist(B365MainP, average_main_10, std_main_10))
		norm.B365GuestNorm = decimal(normdist(B365GuestP, average_guest_10, std_guest_10))
		norm.B365MiddleNorm = decimal(normdist(B365MiddleP, average_middle_10, std_middle_10))

		norm.DensityMain2 = decimal(normdistfalse(B365MainP, average_main_10, std_main_10))
		norm.DensityGuest2 = decimal(normdistfalse(B365GuestP, average_guest_10, std_guest_10))
		norm.DensityMiddle2 = decimal(normdistfalse(B365MiddleP, average_middle_10, std_middle_10))
	}
	//norm.Id = 1

	//if((this.MinuteDiff>=0 && this.MinuteDiff<=10) || this.MinuteDiff>=27){

	//if(this.MinuteDiff>=0 && this.MinuteDiff<=10){
	//	base.Log.Info("Id:",norm.MatchId," liMain: ", norm.CoreMainNorm,"liGuest: ",norm.CoreGuestNorm, " crownMain:", norm.Main10Norm, " crownGuest:", norm.Guest10Norm," ylMain:", norm.Main9Norm, " ylGuest:", norm.Guest9Norm," bwMain:", norm.B365MainNorm, " bwGuest:", norm.B365GuestNorm)
	//
	//	if(norm.CoreMainNorm > 0.7 || norm.CoreGuestNorm>0.7 || norm.Main10Norm>0.7 || norm.Guest10Norm>0.7 || norm.Main9Norm> 0.7 || norm.Guest9Norm>0.7|| norm.B365MainNorm>0.7 || norm.B365GuestNorm>0.7 || norm.PrinMiddleNorm >= 0.98){
	//		//li/crown/yingli/bwin
	//	}else{
	//		his := new(pojo.MatchHis)
	//		last := new(pojo.MatchLast)
	//		his.Id = matchId
	//		last.Id = matchId
	//		this.EuroLastService.Del(last)
	//		this.EuroHisService.Del(his)
	//		return
	//	}
	//}

	this.EuroLastService.InsertNormInfo(norm)
	//this.EuroLastService.SaveList(last_slice)
	//this.EuroLastService.ModifyList(last_update_slice)
	//历史数据
	//his_slice := make([]interface{}, 0)
	//his_update_slice := make([]interface{}, 0)
	//last_all_slice := append(last_slice, last_update_slice)
	//for _, e := range last_all_slice {
	//	bytes, _ := json.Marshal(e)
	//	temp := new(entity3.EuroLast)
	//	json.Unmarshal(bytes, temp)
	//	if len(temp.MatchId) <= 0 {
	//		continue
	//	}
	//	his := new(entity3.EuroHis)
	//	his.EuroLast = *temp
	//
	//	his_temp_id, his_exists := this.EuroHisService.Exist(his)
	//	if !his_exists {
	//		his_slice = append(his_slice, his)
	//	} else {
	//		his.Id = his_temp_id
	//		his_update_slice = append(his_update_slice, his)
	//	}
	//}
	//this.EuroHisService.SaveList(his_slice)
	//this.EuroHisService.ModifyList(his_update_slice)

}

func (this *EuroLastProcesser) Finish() {
	base.Log.Info("euro 抓取解析完成 \r\n")

}

//func main() {
//	num := []float64{2.93, 2.75, 2.9, 2.8, 2.9, 2.88, 2.75, 2.75, 3.25, 2.72}
//	fmt.Println("The given array2 is:", num)
//	average := mean(num)
//
//	std := stdDev(num, average)
//
//	fmt.Println("The average of the above array is:", average)
//
//	fmt.Println("The Standard Deviation2 of the above array is:", std)
//}

func mean(data [10]float64) float64 {
	sum := 0.0
	count := 0
	for _, value := range data {

		if value > 0 {
			sum += value
			count = count + 1
		}
	}

	//return sum / float64(len(data))
	return sum / float64(count)
}

func stdDev(data [10]float64, mean float64) float64 {
	sum := 0.0
	count := 0
	for _, value := range data {
		if value > 0 {
			sum += math.Pow(value-mean, 2)
			count = count + 1
		}
	}

	//variance := sum / float64(len(data)-1)
	variance := sum / float64(count-1)

	return math.Sqrt(variance)
}

func mean9(data [9]float64) float64 {
	sum := 0.0
	count := 0
	for _, value := range data {

		if value > 0 {
			sum += value
			count = count + 1
		}
	}

	//return sum / float64(len(data))
	return sum / float64(count)
}

func stdDev9(data [9]float64, mean float64) float64 {
	sum := 0.0
	count := 0
	for _, value := range data {
		if value > 0 {
			sum += math.Pow(value-mean, 2)
			count = count + 1
		}
	}

	//variance := sum / float64(len(data)-1)
	variance := sum / float64(count-1)

	return math.Sqrt(variance)
}

func normdist(x float64, mean float64, stdev float64) float64 {
	x = (x - mean) / stdev
	var res float64
	if x == 0 {
		res = 0.5
	} else {
		oor2pi := 1 / (math.Sqrt(float64(2) * 3.14159265358979323846))
		t := 1 / (float64(1) + 0.2316419*math.Abs(x))
		t = t * oor2pi * math.Exp(-0.5*x*x) * (0.31938153 + t*(-0.356563782+t*(1.781477937+t*(-1.821255978+t*1.330274429))))
		if x >= 0 {
			res = float64(1) - t
		} else {
			res = t
		}
	}
	return res
}

func normdistfalse(x float64, mean float64, stdev float64) float64{
	exponent  := math.Pow((x-mean),2)/(2*math.Pow(stdev,2))

	return math.Exp(-exponent)/(math.Sqrt(2*math.Pi)*stdev)
}

func decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func decimal3(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", value), 64)
	return value
}
