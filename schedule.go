package launch

import (
	"math/rand"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/elem/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
	"time"
)

/*func main() {
	Before_spider_euroLast()
	Spider_euroLast()
}*/

func Before_spider_euroLast() {
	//抓取前清空当前比较表
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_euro_last"})
}

//查询标识为win007,且欧赔未抓取的配置数据,指定菠菜公司
func Spider_euroLast() {
	datestr := time.Now().Format("2006-01-02")

	//timestr := time.Now().Format("15:04:05")
	//base.Log.Info("matchId1:", datestr, " timestr:", timestr)
	hour := time.Now().Hour()
	hourMarkStr1 := datestr + " " +  strconv.Itoa(hour) + ":30:00"
	hourMarkStr2 := datestr + " " + strconv.Itoa((hour + 1)) + ":00:00"
	hourMark1, _ := time.ParseInLocation("2006-01-02 15:04:05", hourMarkStr1, time.Local)
	hourMark2, _ := time.ParseInLocation("2006-01-02 15:04:05", hourMarkStr2, time.Local)

	base.Log.Info("hourMark1:", hourMark1, " hourMark2:", hourMark2)

	datestr = datestr + " 20:00:00"
	spiderDate, _ := time.ParseInLocation("2006-01-02 15:04:05", datestr, time.Local)
	if time.Now().After(spiderDate) {
		base.Log.Info("datestr:", datestr, " after spiderDate1:", spiderDate)
		time.Sleep(12 * time.Hour)
	}
	var minuteDiff int
	if(time.Now().After(hourMark1)){
		minuteDiff = 60-time.Now().Minute()
	}else{
		minuteDiff = 30-time.Now().Minute()
	}
	if(minuteDiff< 10){
		matchLastService := new(service2.MatchLastService)
		matchLasts := matchLastService.FindNotFinishedTen()
		var compIds []string
		val := utils.GetVal("spider", "euro_comp_ids")
		if len(val) < 0 {
			compService := new(service.CompService)
			compIds = compService.FindEuroIds()
		} else {
			compIds = strings.Split(val, ",")
		}
		processer := proc.GetEuroLastProcesser()
		processer.MatchLastList = matchLasts
		processer.CompWin007Ids = compIds
		processer.SingleThread = true
		processer.Startup()
	}else{
		matchLastService := new(service2.MatchLastService)
		matchLasts := matchLastService.FindNotFinished()
		var compIds []string
		val := utils.GetVal("spider", "euro_comp_ids")
		if len(val) < 0 {
			compService := new(service.CompService)
			compIds = compService.FindEuroIds()
		} else {
			compIds = strings.Split(val, ",")
		}
		processer := proc.GetEuroLastProcesser()
		processer.MatchLastList = matchLasts
		processer.CompWin007Ids = compIds
		processer.SingleThread = true
		processer.Startup()
	}

	//else if(minuteDiff>10 && minuteDiff<20){
	//	rand.Seed(time.Now().Unix())
	//	d := rand.Intn(130) + (minuteDiff-10)*60
	//	time.AfterFunc(time.Second*time.Duration(d), Spider_euroLast)
	//}

	if(minuteDiff>=27){
		rand.Seed(time.Now().Unix())
		d := rand.Intn(20) + 10*60
		time.AfterFunc(time.Second*time.Duration(d), Spider_euroLast)
	}else if(minuteDiff>=0 && minuteDiff<=10){
		rand.Seed(time.Now().Unix())
		d := rand.Intn(20) + 110
		time.AfterFunc(time.Second*time.Duration(d), Spider_euroLast)
	}else{
		rand.Seed(time.Now().Unix())
		d := rand.Intn(30) + (minuteDiff-10)*60
		time.AfterFunc(time.Second*time.Duration(d), Spider_euroLast)
	}

	//timer := time.AfterFunc(time.Minute*time.Duration(d), Spider_euroLast)
	//timer.Stop()
}

func Real() {
	matchLastService := new(service2.MatchLastService)
	matchLasts := matchLastService.FindReal()

	var compIds []string
	val := utils.GetVal("spider", "asia_comp_ids")
	if len(val) < 0 {
		//为空会抓取所有,这里没有必要配置所有的波菜公司ID
		compService := new(service.CompService)
		compIds = compService.FindEuroIds()
	} else {
		compIds = strings.Split(val, ",")
	}
	isHalf := utils.GetVal("spider", "isHalf")
	processer := proc.GetRealProcesser()
	processer.MatchLastList = matchLasts
	processer.CompWin007Ids = compIds
	processer.SingleThread = true
	processer.Ishalf = isHalf
	processer.Startup()

	//rand.Seed(time.Now().Unix())
	//d := rand.Intn(5) + 9
	//time.AfterFunc(time.Minute*time.Duration(d), Real)
}

func Spider_id(matchId string) {
	//matchLastService := new(service2.MatchLastService)
	dataList := make([]*pojo.MatchLast, 0)
	data := new(pojo.MatchLast)
	data.Id = matchId
	dataList = append(dataList, data)

	var compIds []string
	val := utils.GetVal("spider", "euro_comp_ids")
	if len(val) < 0 {
		//为空会抓取所有,这里没有必要配置所有的波菜公司ID
		compService := new(service.CompService)
		compIds = compService.FindEuroIds()
	} else {
		compIds = strings.Split(val, ",")
	}

	processer := proc.GetEuroLastProcesser()
	processer.MatchLastList = dataList
	processer.CompWin007Ids = compIds
	processer.SingleThread = true
	processer.Startup()

	//rand.Seed(time.Now().Unix())
	//d := rand.Intn(10) + 10
	//base.Log.Error("间隔时间:,", d)
	//time.AfterFunc(time.Minute*time.Duration(d), Spider_id())
}

//func Spider_euroLast_his(season string) {
//	matchLastService := new(service2.MatchHisService)
//	var matchLasts []*pojo.MatchLast
//	matchLasts = matchLastService.FindBySeason(season)
//
//	var compIds []string
//	//为空会抓取所有,这里没有必要配置所有的波菜公司ID
//	compService := new(service.CompService)
//	compIds = compService.FindEuroIds()
//
//	processer := proc.GetEuroLastProcesser()
//	processer.MatchLastList = matchLasts
//	processer.CompWin007Ids = compIds
//	processer.Startup()
//}

//查询标识为win007,且欧赔未抓取的配置数据,指定菠菜公司
//func Spider_euroLast_near() {
//	matchLastService := new(service2.MatchLastService)
//	matchLasts := matchLastService.FindNear()
//	if len(matchLasts) <= 0 {
//		return
//	}
//
//	var compIds []string
//	val := utils.GetVal("spider", "euro_comp_ids")
//	if len(val) < 0 {
//		//为空会抓取所有,这里没有必要配置所有的波菜公司ID
//		compService := new(service.CompService)
//		compIds = compService.FindEuroIds()
//	} else {
//		compIds = strings.Split(val, ",")
//	}
//
//	processer := proc.GetEuroLastProcesser()
//	processer.MatchLastList = matchLasts
//	processer.CompWin007Ids = compIds
//	processer.SingleThread = true
//	processer.Startup()
//}
