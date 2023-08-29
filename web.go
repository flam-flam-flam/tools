package main

import (
	"database/sql"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type matchHis struct {
	Id          string
	MainTeamId  string
	GuestTeamId string
	MatchDate   time.Time
	LeagueName  string
}

type norm struct {
	MatchId        string
	MainName       string
	GuestName      string
	CompCount      int
	Main10Norm     float64
	Middle10Norm   float64
	Guest10Norm    float64
	MainP          float64
	MiddleP        float64
	GuestP         float64
	IN3            float64
	IN1            float64
	IN0            float64
	IP3            float64
	IP1            float64
	IP0            float64
	CP3            float64
	CP1            float64
	CP0            float64
	CN3            float64
	CN1            float64
	CN0            float64
	PrinMainNorm   float64
	PrinGuestNorm  float64
	PrinMiddleNorm float64
	PrinMainP      float64
	PrinGuestP     float64
	PrinMiddleP    float64
	B365MainNorm   float64
	B365GuestNorm  float64
	B365MiddleNorm float64
	B365MainP      float64
	B365GuestP     float64
	B365MiddleP    float64
	DensityMain1   float64
	DensityGuest1  float64
	DensityMiddle1 float64
	DensityMain2   float64
	DensityGuest2  float64
	DensityMiddle2 float64
	DensityMain3   float64
	DensityGuest3  float64
	DensityMiddle3 float64
	OddTime        string
	MatchTime      time.Time
}

//MatchId,Pankou_1,CompCount_1,Main10_3_1,Main10_0_1,Pankou_2,CompCount_2,Main10_3_2,Main10_0_2,CreateTime
type normAsia struct {
	MatchId      string
	Pankou1      float64
	CompCount1   int
	Main10Norm1  float64
	Guest10Norm1 float64
	Pankou2      float64
	CompCount2   int
	Main10Norm2  float64
	Guest10Norm2 float64
	OddTime      time.Time
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "foot"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(10.100.93.53:3306)/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("D:/foot-master/foot-web/service/form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT la.Id,la.MainTeamId,la.GuestTeamId,la.MatchDate,tl.Name FROM foot.t_match_his la left join t_league tl  ON la.LeagueId=tl.Id WHERE MatchDate > DATE_SUB(NOW(),INTERVAL 2200 MINUTE) AND MatchDate < DATE_ADD(NOW(), INTERVAL 3 HOUR) ORDER BY MatchDate ASC")
	if err != nil {
		panic(err.Error())
	}
	match := matchHis{}
	res := []matchHis{}
	for selDB.Next() {
		var Id, MainTeamId, GuestTeamId, Name string
		var MatchDate time.Time
		err = selDB.Scan(&Id, &MainTeamId, &GuestTeamId, &MatchDate, &Name)
		if err != nil {
			panic(err.Error())
		}
		match.Id = Id
		match.MainTeamId = MainTeamId
		match.GuestTeamId = GuestTeamId
		match.MatchDate = MatchDate
		match.LeagueName = Name
		res = append(res, match)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

//type LineExamples struct{}
func  Examples() {
	page := components.NewPage()
	page.AddCharts(
		lineSmooth(),
		//lineMain(),
	)
	//page.PageTitle
	f, err := os.Create("D:/foot-master/foot-web/service/form/line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

func Smooth(w http.ResponseWriter, r *http.Request) {
	//nId := r.URL.Query().Get("id")
	//log.Println("INSERT: Name: "  +  " | City: ")
	pathstr := "D:/foot-master/foot-web/service/form/line2.html"
	t1 := template.Must(template.ParseFiles(pathstr))
	t1.Execute(w, "")
}

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 6; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(100)})
	}
	return items
}
func Line(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	str_name := "T-" + nId
	//str_file := "D:/foot-master/foot-web/service/form/" + str_name + ".html"
	//pathstr := "D:/foot-master/foot-web/service/form/line2.html"
	pathstr := "D:/foot-master/foot-web/service/form/"  + str_name + ".html"
	t1 := template.Must(template.ParseFiles(pathstr))
	t1.Execute(w, "")
}

func HandleData(dataList []norm, matchId string){

	fruits := []string{}
	items12BetMainNorm := make([]opts.LineData, 0) //12bet
	items12BetMainP := make([]opts.LineData, 0) //12bet
	itemsYiMainNorm := make([]opts.LineData, 0) //易时博
	itemsYiMainP := make([]opts.LineData, 0) //易时博
	itemsYingLiMainNorm := make([]opts.LineData, 0) //盈利
	itemsYingLiMainP := make([]opts.LineData, 0) //盈利

	items12BetMiddleNorm := make([]opts.LineData, 0) //12bet
	items12BetMiddleP := make([]opts.LineData, 0) //12bet
	itemsYiMiddleNorm := make([]opts.LineData, 0) //易时博
	itemsYiMiddleP := make([]opts.LineData, 0) //易时博
	itemsYingLiMiddleNorm := make([]opts.LineData, 0) //盈利
	itemsYingLiMiddleP := make([]opts.LineData, 0) //盈利

	items12BetGuestNorm := make([]opts.LineData, 0) //12bet
	items12BetGuestP := make([]opts.LineData, 0) //12bet
	itemsYiGuestNorm := make([]opts.LineData, 0) //易时博
	itemsYiGuestP := make([]opts.LineData, 0) //易时博
	itemsYingLiGuestNorm := make([]opts.LineData, 0) //盈利
	itemsYingLiGuestP := make([]opts.LineData, 0) //盈利
	var i int
	if len(dataList) < 10{
		i = len(dataList)-1
	}else {
		i = 9
	}

	//fruits = append(fruits, "000000")
	//items12BetMainNorm = append(items12BetMainNorm, opts.LineData{Value: dataList[i].Main10Norm})
	//itemsYiMainNorm = append(itemsYiMainNorm, opts.LineData{Value: dataList[i].CN3})
	//itemsYingLiMainNorm = append(itemsYingLiMainNorm, opts.LineData{Value: dataList[i].IN3})
	//items12BetMainP = append(items12BetMainP, opts.LineData{Value: 6.0})
	//itemsYiMainP = append(itemsYiMainP, opts.LineData{Value: 6.0})
	//itemsYingLiMainP = append(itemsYingLiMainP, opts.LineData{Value: 6.0})
	//
	//items12BetMiddleNorm = append(items12BetMiddleNorm, opts.LineData{Value: dataList[i].Middle10Norm})
	//itemsYiMiddleNorm = append(itemsYiMiddleNorm, opts.LineData{Value: dataList[i].CN1})
	//itemsYingLiMiddleNorm = append(itemsYingLiMiddleNorm, opts.LineData{Value: dataList[i].IN1})
	//items12BetMiddleP = append(items12BetMiddleP, opts.LineData{Value: 6.0})
	//itemsYiMiddleP = append(itemsYiMiddleP, opts.LineData{Value: 6.0})
	//itemsYingLiMiddleP = append(itemsYingLiMiddleP, opts.LineData{Value: 6.0})
	//
	//items12BetGuestNorm = append(items12BetGuestNorm, opts.LineData{Value: dataList[i].Guest10Norm})
	//itemsYiGuestNorm = append(itemsYiGuestNorm, opts.LineData{Value: dataList[i].CN0})
	//itemsYingLiGuestNorm = append(itemsYingLiGuestNorm, opts.LineData{Value: dataList[i].IN0})
	//items12BetGuestP = append(items12BetGuestP, opts.LineData{Value: 6.0})
	//itemsYiGuestP = append(itemsYiGuestP, opts.LineData{Value: 6.0})
	//itemsYingLiGuestP = append(itemsYingLiGuestP, opts.LineData{Value: 6.0})

	for ; i >= 0 ; i-- {
		fruits = append(fruits, dataList[i].OddTime)
		items12BetMainNorm = append(items12BetMainNorm, opts.LineData{Value: dataList[i].Main10Norm})
		itemsYiMainNorm = append(itemsYiMainNorm, opts.LineData{Value: dataList[i].CN3})
		itemsYingLiMainNorm = append(itemsYingLiMainNorm, opts.LineData{Value: dataList[i].IN3})
		items12BetMainP = append(items12BetMainP, opts.LineData{Value: dataList[i].MainP})
		itemsYiMainP = append(itemsYiMainP, opts.LineData{Value: dataList[i].CP3})
		itemsYingLiMainP = append(itemsYingLiMainP, opts.LineData{Value: dataList[i].IP3})

		items12BetMiddleNorm = append(items12BetMiddleNorm, opts.LineData{Value: dataList[i].Middle10Norm})
		itemsYiMiddleNorm = append(itemsYiMiddleNorm, opts.LineData{Value: dataList[i].CN1})
		itemsYingLiMiddleNorm = append(itemsYingLiMiddleNorm, opts.LineData{Value: dataList[i].IN1})
		items12BetMiddleP = append(items12BetMiddleP, opts.LineData{Value: dataList[i].MiddleP})
		itemsYiMiddleP = append(itemsYiMiddleP, opts.LineData{Value: dataList[i].CP1})
		itemsYingLiMiddleP = append(itemsYingLiMiddleP, opts.LineData{Value: dataList[i].IP1})

		items12BetGuestNorm = append(items12BetGuestNorm, opts.LineData{Value: dataList[i].Guest10Norm})
		itemsYiGuestNorm = append(itemsYiGuestNorm, opts.LineData{Value: dataList[i].CN0})
		itemsYingLiGuestNorm = append(itemsYingLiGuestNorm, opts.LineData{Value: dataList[i].IN0})
		items12BetGuestP = append(items12BetGuestP, opts.LineData{Value: dataList[i].GuestP})
		itemsYiGuestP = append(itemsYiGuestP, opts.LineData{Value: dataList[i].CP0})
		itemsYingLiGuestP = append(itemsYingLiGuestP, opts.LineData{Value: dataList[i].IP0})

	}

	page := components.NewPage()
	page.AddCharts(
		lineMainN(items12BetMainNorm,fruits,itemsYiMainNorm,itemsYingLiMainNorm),
		lineMainP(items12BetMainP,fruits,itemsYiMainP,itemsYingLiMainP),
		lineMiddleN(items12BetMiddleNorm,fruits,itemsYiMiddleNorm,itemsYingLiMiddleNorm),
		lineMiddleP(items12BetMiddleP,fruits,itemsYiMiddleP,itemsYingLiMiddleP),
		lineGuestN(items12BetGuestNorm,fruits,itemsYiGuestNorm,itemsYingLiGuestNorm),
		lineGuestP(items12BetGuestP,fruits,itemsYiGuestP,itemsYingLiGuestP),
	)
	str_name := "T-" + matchId
	page.PageTitle = str_name
	str_file := "D:/foot-master/foot-web/service/form/" + str_name + ".html"
	//f, err := os.Create("D:/foot-master/foot-web/service/form/line2.html")
	f, err := os.Create(str_file)
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

func lineMiddleP(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData) *charts.Line {

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "平赔",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: 1.9, Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("易时博", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("12Bet", items12Bet).SetSeriesOptions(
		charts.WithLineChartOpts(opts.LineChart{
			ShowSymbol: true,
		}),
		charts.WithLabelOpts(opts.Label{
			Show: true,
		}),
	)
	//line.Scale = true
	return line
}

func lineGuestP(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData) *charts.Line {

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "客赔",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		//charts.WithInitializationOpts(opts.Initialization{PageTitle:"Diagramms", Theme: types.ThemeWesteros, Width: "1200px", Height: "800px"}),
		//charts.WithLegendOpts(opts.Legend{Show: true, Left: "right", Orient: "vertical", Y: "100"}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: 1.9, Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("易时博", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("12Bet", items12Bet).SetSeriesOptions(
		charts.WithLineChartOpts(opts.LineChart{
			ShowSymbol: true,
		}),
		charts.WithLabelOpts(opts.Label{
			Show: true,
		}),
	)
	//line.Scale = true
	return line
}

func lineMainP(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData) *charts.Line {

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "主赔",
			//Top:"0",
			//Right: "400",
		}),
		//charts.WithLegendOpts(opts.Legend{Show: true, Left: "right", Orient: "vertical", Y: "100"}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: 1.9, Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("易时博", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("12Bet", items12Bet).SetSeriesOptions(
		charts.WithLineChartOpts(opts.LineChart{
			ShowSymbol: true,
		}),
		charts.WithLabelOpts(opts.Label{
			Show: true,
		}),
	)
	//line.Scale = true
	return line
}

func lineMiddleN(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData) *charts.Line {

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "平Norm",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
	)
	line.SetXAxis(fruits).
		AddSeries("易时博", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("12Bet", items12Bet).SetSeriesOptions(
		charts.WithLineChartOpts(opts.LineChart{
			ShowSymbol: true,
		}),
		charts.WithLabelOpts(opts.Label{
			Show: true,
		}),
	)
	return line
}

func lineGuestN(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData) *charts.Line {

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "客Norm",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			//Width: "720px",
			//Height: "500px",
		}),
		//charts.WithLineChartOpts(opts.LineChart{
		//
		//})
	)
	line.SetXAxis(fruits).
		AddSeries("易时博", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("12Bet", items12Bet).SetSeriesOptions(
		charts.WithLineChartOpts(opts.LineChart{
			ShowSymbol: true,
		}),
		charts.WithLabelOpts(opts.Label{
			Show: true,
		}),
	)
	return line
}

func lineMainN(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData) *charts.Line {

	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "主Norm",
		}),
		//charts.WithLegendOpts(opts.Legend{Left: "left",Top: "top"}), //legend是设定图例的
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Show:true,
			MinInterval: 0,
		}),
		//charts.WithToolboxOpts(opts.Toolbox{
		//	Top: "10px",
		//	Left: "0px",
		//}),
	)
	line.SetXAxis(fruits).
		AddSeries("易时博", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("12Bet", items12Bet).SetSeriesOptions(
		charts.WithLineChartOpts(opts.LineChart{
			ShowSymbol: true,
		}),
		charts.WithLabelOpts(opts.Label{
			Show: true,
		}),
	)
	return line
}


func lineSmooth() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "multi lines",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
		}),
	)
	fruits := []string{"Apple", "Banana", "Peach ", "Lemon", "Pear", "Cherry"}
	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		AddSeries("Category C", generateLineItems()).SetSeriesOptions(
		charts.WithLineChartOpts(opts.LineChart{
			ShowSymbol: true,
		}),
		charts.WithLabelOpts(opts.Label{
			Show: true,
		}),
	)
		//.SetSeriesOptions(charts.WithLineChartOpts(
		//
		//	opts.LineChart{
		//		Smooth: true,
		//	}),
		//)
	return line
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	sql := "SELECT MatchId,MainName,GuestName,CompCount,Main10Norm,Middle10Norm,Guest10Norm,Ep3,Ep1,Ep0,Main9Norm,Middle9Norm,Guest9Norm,OddTime,MatchTime,CoreMainNorm,CoreGuestNorm,CoreMiddleNorm,CoreMainP,CoreGuestP,CoreMiddleP,IntMainP,IntGuestP,IntMiddleP,PrinMainNorm,PrinGuestNorm,PrinMiddleNorm,PrinMainP,PrinGuestP,PrinMiddleP,B365MainNorm,B365GuestNorm,B365MiddleNorm,B365MainP,B365GuestP,B365MiddleP,DensityMain1,DensityGuest1,DensityMiddle1,DensityMain2,DensityGuest2,DensityMiddle2,DensityMain3,DensityGuest3,DensityMiddle3 from t_norm where MatchId = '" + nId + "' ORDER BY OddTime DESC"
	selDB, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	match := norm{}
	res := []norm{}
	for selDB.Next() {
		var MatchId, MainName, GuestName string
		var CompCount int
		var Main10Norm, Middle10Norm, Guest10Norm, MainP, MiddleP, GuestP, Main9Norm, Middle9Norm, Guest9Norm, CoreMainNorm, CoreGuestNorm, CoreMiddleNorm, CoreMainP, CoreGuestP, CoreMiddleP, IntMainP, IntGuestP, IntMiddleP,PrinMainNorm,PrinGuestNorm,PrinMiddleNorm,PrinMainP,PrinGuestP,PrinMiddleP,B365MainNorm,B365GuestNorm,B365MiddleNorm,B365MainP,B365GuestP,B365MiddleP,DensityMain1,DensityGuest1,DensityMiddle1,DensityMain2,DensityGuest2,DensityMiddle2,DensityMain3,DensityGuest3,DensityMiddle3 float64
		var OddTime, MatchTime time.Time
		err = selDB.Scan(&MatchId, &MainName, &GuestName, &CompCount, &Main10Norm, &Middle10Norm, &Guest10Norm, &MainP, &MiddleP, &GuestP, &Main9Norm, &Middle9Norm, &Guest9Norm, &OddTime, &MatchTime, &CoreMainNorm, &CoreGuestNorm, &CoreMiddleNorm, &CoreMainP, &CoreGuestP, &CoreMiddleP, &IntMainP, &IntGuestP, &IntMiddleP, &PrinMainNorm,&PrinGuestNorm,&PrinMiddleNorm,&PrinMainP,&PrinGuestP,&PrinMiddleP,&B365MainNorm,&B365GuestNorm,&B365MiddleNorm,&B365MainP,&B365GuestP,&B365MiddleP,&DensityMain1,&DensityGuest1,&DensityMiddle1,&DensityMain2,&DensityGuest2,&DensityMiddle2,&DensityMain3,&DensityGuest3,&DensityMiddle3)
		if err != nil {
			panic(err.Error())
		}
		namestr := []rune(MainName)
		match.MatchId = MatchId
		match.MainName = string(namestr[:3])
		match.GuestName = GuestName
		match.CompCount = CompCount
		match.Main10Norm = Main10Norm
		match.Middle10Norm = Middle10Norm
		match.Guest10Norm = Guest10Norm
		match.MainP = MainP
		match.MiddleP = MiddleP
		match.GuestP = GuestP
		match.IN3 = Main9Norm
		match.IN1 = Middle9Norm
		match.IN0 = Guest9Norm
		match.IP3 = IntMainP
		match.IP1 = IntMiddleP
		match.IP0 = IntGuestP
		match.OddTime = OddTime.Format("15:04:05")
		match.MatchTime = MatchTime
		match.CN3 = CoreMainNorm
		match.CN1 = CoreMiddleNorm
		match.CN0 = CoreGuestNorm
		match.CP3 = CoreMainP
		match.CP1 = CoreMiddleP
		match.CP0 = CoreGuestP
		match.PrinMainNorm = PrinMainNorm
		match.PrinGuestNorm = PrinGuestNorm
		match.PrinMiddleNorm = PrinMiddleNorm
		match.PrinMainP = PrinMainP
		match.PrinMiddleP = PrinMiddleP
		match.PrinGuestP = PrinGuestP

		match.B365MainNorm = B365MainNorm
		match.B365GuestNorm = B365GuestNorm
		match.B365MiddleNorm = B365MiddleNorm
		match.B365MainP = B365MainP
		match.B365MiddleP = B365MiddleP
		match.B365GuestP = B365GuestP

		match.DensityMain1 = DensityMain1
		match.DensityGuest1 = DensityGuest1
		match.DensityMiddle1 = DensityMiddle1
		match.DensityMain2 = DensityMain2
		match.DensityGuest2 = DensityGuest2
		match.DensityMiddle2 = DensityMiddle2
		match.DensityMain3 = DensityMain3
		match.DensityGuest3 = DensityGuest3
		match.DensityMiddle3 = DensityMiddle3
		res = append(res, match)
	}
	HandleData(res,nId)
	//if(len(res)==0){
	//	tmpl.ExecuteTemplate(w, "index", res)
	//}
	//tmpl.ExecuteTemplate(w, "Show", res)
	http.Redirect(w, r, "/", 301)
	defer db.Close()
}



func IndexAsia(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT la.Id,la.MainTeamId,la.GuestTeamId,la.MatchDate,tl.Name FROM foot.t_match_last la left join t_league tl  ON la.LeagueId=tl.Id WHERE MatchDate > DATE_SUB(NOW(),INTERVAL 90 MINUTE) AND MatchDate < DATE_ADD(NOW(), INTERVAL 3 HOUR) ORDER BY MatchDate ASC")
	if err != nil {
		panic(err.Error())
	}
	match := matchHis{}
	res := []matchHis{}
	for selDB.Next() {
		var Id, MainTeamId, GuestTeamId, Name string
		var MatchDate time.Time
		err = selDB.Scan(&Id, &MainTeamId, &GuestTeamId, &MatchDate, &Name)
		if err != nil {
			panic(err.Error())
		}
		match.Id = Id
		match.MainTeamId = MainTeamId
		match.GuestTeamId = GuestTeamId
		match.MatchDate = MatchDate
		match.LeagueName = Name
		res = append(res, match)
	}
	tmpl.ExecuteTemplate(w, "IndexAsia", res)
	defer db.Close()
}

func ShowAsia(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	sql := "SELECT MatchId,Pankou_1,CompCount_1,Main10_3_1,Main10_0_1,Pankou_2,CompCount_2,Main10_3_2,Main10_0_2,CreateTime from t_norm_asia where MatchId = '" + nId + "' ORDER BY CreateTime DESC"
	selDB, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	match := normAsia{}
	res := []normAsia{}
	for selDB.Next() {
		var MatchId string
		var CompCount1, CompCount2 int
		var Pankou1, Main10Norm1, Guest10Norm1, Pankou2, Main10Norm2, Guest10Norm2 float64
		var OddTime time.Time
		err = selDB.Scan(&MatchId, &Pankou1, &CompCount1, &Main10Norm1, &Guest10Norm1, &Pankou2, &CompCount2, &Main10Norm2, &Guest10Norm2, &OddTime)
		if err != nil {
			panic(err.Error())
		}
		match.MatchId = MatchId
		match.Pankou1 = Pankou1
		match.CompCount1 = CompCount1
		match.Main10Norm1 = Main10Norm1
		match.Guest10Norm1 = Guest10Norm1
		match.Pankou2 = Pankou2
		match.CompCount2 = CompCount2
		match.Main10Norm2 = Main10Norm2
		match.Guest10Norm2 = Guest10Norm2
		match.OddTime = OddTime

		res = append(res, match)
	}
	//if(len(res)==0){
	//	tmpl.ExecuteTemplate(w, "index", res)
	//}
	tmpl.ExecuteTemplate(w, "ShowAsia", res)
	defer db.Close()
}

//func Show(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	nId := r.URL.Query().Get("id")
//	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
//	if err != nil {
//		panic(err.Error())
//	}
//	emp := Employee{}
//	for selDB.Next() {
//		var id int
//		var name, city string
//		err = selDB.Scan(&id, &name, &city)
//		if err != nil {
//			panic(err.Error())
//		}
//		emp.Id = id
//		emp.Name = name
//		emp.City = city
//	}
//	tmpl.ExecuteTemplate(w, "Show", emp)
//	defer db.Close()
//}
//
//func New(w http.ResponseWriter, r *http.Request) {
//	tmpl.ExecuteTemplate(w, "New", nil)
//}
//
//func Edit(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	nId := r.URL.Query().Get("id")
//	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
//	if err != nil {
//		panic(err.Error())
//	}
//	emp := Employee{}
//	for selDB.Next() {
//		var id int
//		var name, city string
//		err = selDB.Scan(&id, &name, &city)
//		if err != nil {
//			panic(err.Error())
//		}
//		emp.Id = id
//		emp.Name = name
//		emp.City = city
//	}
//	tmpl.ExecuteTemplate(w, "Edit", emp)
//	defer db.Close()
//}
//
//func Insert(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	if r.Method == "POST" {
//		name := r.FormValue("name")
//		city := r.FormValue("city")
//		insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")
//		if err != nil {
//			panic(err.Error())
//		}
//		insForm.Exec(name, city)
//		log.Println("INSERT: Name: " + name + " | City: " + city)
//	}
//	defer db.Close()
//	http.Redirect(w, r, "/", 301)
//}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM t_match_his WHERE Id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE ", emp)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
func DeleteAsia(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM t_match_match WHERE Id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/asia", IndexAsia)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/showAsia", ShowAsia)
	//http.HandleFunc("/new", New)
	//http.HandleFunc("/edit", Edit)
	//http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/deleteAsia", DeleteAsia)
	http.HandleFunc("/smooth", Smooth)
	http.HandleFunc("/line", Line)
	Examples()
	http.ListenAndServe(":8088", nil)
}
