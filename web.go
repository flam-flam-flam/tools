package main

import (
	"database/sql"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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
	Flesh string
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
	MainMaxName    string
	MiddleMaxName  string
	GuestMaxName   string
	BaseMainMaxValue1 float64
	BaseMainMaxValue2 float64
	BaseMiddleMaxValue1 float64
	BaseMiddleMaxValue2 float64
	BaseGuestMaxValue1 float64
	BaseGuestMaxValue2 float64
	MainFirstMax float64
	MainSecondMax float64
	MiddleFirstMax float64
	MiddleSecondMax float64
	GuestFirstMax float64
	GuestSecondMax float64
	MainMaxAddStd      float64
	MiddleMaxAddStd    float64
	GuestMaxAddStd     float64
	MainMeanAdd3Std    float64
	MiddleMeanAdd3Std  float64
	GuestMeanAdd3Std   float64
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
	selDB, err := db.Query("SELECT la.Id,la.MainTeamId,la.GuestTeamId,la.MatchDate,tl.Name FROM foot.t_match_his la left join t_league tl  ON la.LeagueId=tl.Id WHERE MatchDate > DATE_SUB(NOW(),INTERVAL 150 MINUTE) AND MatchDate < DATE_ADD(NOW(), INTERVAL 147 MINUTE) ORDER BY MatchDate ASC")
	if err != nil {
		panic(err.Error())
	}
	match := matchHis{}
	res := []matchHis{}
	var flesh = strconv.FormatInt(time.Now().UnixNano(),10)
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
		match.Flesh = flesh
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

func HandleData(dataList []norm, matchId string,CompCount int){

	fruits := []string{}
	companysMain := []string{}
	companysMiddle := []string{}
	companysGuest := []string{}

	BaseMainMaxValue1 := make([]opts.LineData, 0)
	BaseMainMaxValue2 := make([]opts.LineData, 0)
	BaseMiddleMaxValue1 := make([]opts.LineData, 0)
	BaseMiddleMaxValue2 := make([]opts.LineData, 0)
	BaseGuestMaxValue1 := make([]opts.LineData, 0)
	BaseGuestMaxValue2 := make([]opts.LineData, 0)
	MainFirstMax := make([]opts.LineData, 0)
	MainSecondMax := make([]opts.LineData, 0)
	MiddleFirstMax := make([]opts.LineData, 0)
	MiddleSecondMax := make([]opts.LineData, 0)
	GuestFirstMax := make([]opts.LineData, 0)
	GuestSecondMax := make([]opts.LineData, 0)

	MainMaxAddStd  := make([]opts.LineData, 0)
	MiddleMaxAddStd := make([]opts.LineData, 0)
	GuestMaxAddStd := make([]opts.LineData, 0)
	MainMeanAdd3Std := make([]opts.LineData, 0)
	MiddleMeanAdd3Std := make([]opts.LineData, 0)
	GuestMeanAdd3Std := make([]opts.LineData, 0)


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


	itemsweideMainNorm := make([]opts.LineData, 0) //weide
	itemsweideMainP := make([]opts.LineData, 0) //weide
	itemsweideMiddleNorm := make([]opts.LineData, 0) //weide
	itemsweideMiddleP := make([]opts.LineData, 0) //weide
	itemsweideGuestNorm := make([]opts.LineData, 0) //weide
	itemsweideGuestP := make([]opts.LineData, 0) //weide


	itemsbwinMainNorm := make([]opts.LineData, 0) //bwin
	itemsbwinMainP := make([]opts.LineData, 0) //bwin
	itemsbwinMiddleNorm := make([]opts.LineData, 0) //bwin
	itemsbwinMiddleP := make([]opts.LineData, 0) //bwin
	itemsbwinGuestNorm := make([]opts.LineData, 0) //bwin
	itemsbwinGuestP := make([]opts.LineData, 0) //bwin

	var i int
	if len(dataList) < 15{
		if(len(dataList) == 0){
			return
		}
		i = len(dataList)-1
	}else {
		i = 14
	}


	for ; i >= 0 ; i-- {
		companysMain = append(companysMain, dataList[i].MainMaxName)
		companysMiddle = append(companysMiddle, dataList[i].MiddleMaxName)
		companysGuest = append(companysGuest, dataList[i].GuestMaxName)
		BaseMainMaxValue1= append(BaseMainMaxValue1,opts.LineData{Value: dataList[i].BaseMainMaxValue1})
		BaseMainMaxValue2= append(BaseMainMaxValue2,opts.LineData{Value: dataList[i].BaseMainMaxValue2})
		BaseMiddleMaxValue1 = append(BaseMiddleMaxValue1,opts.LineData{Value: dataList[i].BaseMiddleMaxValue1})
		BaseMiddleMaxValue2 = append(BaseMiddleMaxValue2,opts.LineData{Value: dataList[i].BaseMiddleMaxValue2})
		BaseGuestMaxValue1 = append(BaseGuestMaxValue1,opts.LineData{Value: dataList[i].BaseGuestMaxValue1})
		BaseGuestMaxValue2 = append(BaseGuestMaxValue2,opts.LineData{Value: dataList[i].BaseGuestMaxValue2})

		MainFirstMax = append(MainFirstMax,opts.LineData{Value: dataList[i].MainFirstMax})
		MainSecondMax = append(MainSecondMax,opts.LineData{Value: dataList[i].MainSecondMax})
		MiddleFirstMax = append(MiddleFirstMax,opts.LineData{Value: dataList[i].MiddleFirstMax})
		MiddleSecondMax = append(MiddleSecondMax,opts.LineData{Value: dataList[i].MiddleSecondMax})
		GuestFirstMax = append(GuestFirstMax,opts.LineData{Value: dataList[i].GuestFirstMax})
		GuestSecondMax = append(GuestSecondMax,opts.LineData{Value: dataList[i].GuestSecondMax})

		MainMaxAddStd = append(MainMaxAddStd,opts.LineData{Value: dataList[i].MainMaxAddStd})
		MiddleMaxAddStd = append(MiddleMaxAddStd,opts.LineData{Value: dataList[i].MiddleMaxAddStd})
		GuestMaxAddStd = append(GuestMaxAddStd,opts.LineData{Value: dataList[i].GuestMaxAddStd})
		MainMeanAdd3Std = append(MainMeanAdd3Std,opts.LineData{Value: dataList[i].MainMeanAdd3Std})
		MiddleMeanAdd3Std = append(MiddleMeanAdd3Std,opts.LineData{Value: dataList[i].MiddleMeanAdd3Std})
		GuestMeanAdd3Std = append(GuestMeanAdd3Std,opts.LineData{Value: dataList[i].GuestMeanAdd3Std})

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

		//weilian
		itemsweideMainNorm = append(itemsweideMainNorm, opts.LineData{Value: dataList[i].PrinMainNorm})
		itemsweideMiddleNorm = append(itemsweideMiddleNorm, opts.LineData{Value: dataList[i].PrinMiddleNorm})
		itemsweideGuestNorm = append(itemsweideGuestNorm, opts.LineData{Value: dataList[i].PrinGuestNorm})
		itemsweideMainP = append(itemsweideMainP, opts.LineData{Value: dataList[i].PrinMainP})
		itemsweideMiddleP = append(itemsweideMiddleP, opts.LineData{Value: dataList[i].PrinMiddleP})
		itemsweideGuestP = append(itemsweideGuestP, opts.LineData{Value: dataList[i].PrinGuestP})

		//bwin
		itemsbwinMainNorm = append(itemsbwinMainNorm, opts.LineData{Value: dataList[i].B365MainNorm})
		itemsbwinMiddleNorm = append(itemsbwinMiddleNorm, opts.LineData{Value: dataList[i].B365MiddleNorm})
		itemsbwinGuestNorm = append(itemsbwinGuestNorm, opts.LineData{Value: dataList[i].B365GuestNorm})
		itemsbwinMainP = append(itemsbwinMainP, opts.LineData{Value: dataList[i].B365MainP})
		itemsbwinMiddleP = append(itemsbwinMiddleP, opts.LineData{Value: dataList[i].B365MiddleP})
		itemsbwinGuestP = append(itemsbwinGuestP, opts.LineData{Value: dataList[i].B365GuestP})

	}

	page := components.NewPage()
	page.AddCharts(
		//lineMainN(items12BetMainNorm,fruits,itemsYiMainNorm,itemsYingLiMainNorm,CompCount,itemsweideMainNorm,itemsbwinMainNorm),
		//lineMainP(items12BetMainP,fruits,itemsYiMainP,itemsYingLiMainP,CompCount,itemsweideMainP,itemsbwinMainP),
		lineMainAnaly(companysMain,BaseMainMaxValue1,BaseMainMaxValue2,MainFirstMax,MainSecondMax,CompCount),
		lineMainMax(fruits,MainMeanAdd3Std,MainMaxAddStd,CompCount),
		//lineMiddleN(items12BetMiddleNorm,fruits,itemsYiMiddleNorm,itemsYingLiMiddleNorm,CompCount,itemsweideMiddleNorm,itemsbwinMiddleNorm),
		//lineMiddleP(items12BetMiddleP,fruits,itemsYiMiddleP,itemsYingLiMiddleP,CompCount,itemsweideMiddleP,itemsbwinMiddleP),
		lineMiddleAnaly(companysMiddle,BaseMiddleMaxValue1,BaseMiddleMaxValue2,MiddleFirstMax,MiddleSecondMax,CompCount),
		lineMiddleMax(fruits,MiddleMeanAdd3Std,MiddleMaxAddStd,CompCount),
		//lineGuestN(items12BetGuestNorm,fruits,itemsYiGuestNorm,itemsYingLiGuestNorm,CompCount,itemsweideGuestNorm,itemsbwinGuestNorm),
		//lineGuestP(items12BetGuestP,fruits,itemsYiGuestP,itemsYingLiGuestP,CompCount,itemsweideGuestP,itemsbwinGuestP),
		lineGuestAnaly(companysGuest,BaseGuestMaxValue1,BaseGuestMaxValue2,GuestFirstMax,GuestSecondMax,CompCount),
		lineGuestMax(fruits,GuestMeanAdd3Std,GuestMaxAddStd,CompCount),
	)
	str_name := "T-" + matchId
	//page.PageTitle = str_name
	page.PageTitle = dataList[0].MainName
	str_file := "D:/foot-master/foot-web/service/form/" + str_name + ".html"
	//f, err := os.Create("D:/foot-master/foot-web/service/form/line2.html")
	f, err := os.Create(str_file)
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

func lineMainMax(fruits []string, MeanAdd3Std []opts.LineData,MaxAddStd []opts.LineData,CompCount int) *charts.Line {

	line := charts.NewLine()
	title := "主meanAddStd_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: "dataMin", Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("MeanAdd3Std", MeanAdd3Std).
		AddSeries("MaxAddStd", MaxAddStd).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: true,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}
func lineMiddleMax(fruits []string, MeanAdd3Std []opts.LineData,MaxAddStd []opts.LineData,CompCount int) *charts.Line {

	line := charts.NewLine()
	title := "平meanAddStd_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: "dataMin", Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("MeanAdd3Std", MeanAdd3Std).
		AddSeries("MaxAddStd", MaxAddStd).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: true,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}
func lineGuestMax(fruits []string, MeanAdd3Std []opts.LineData,MaxAddStd []opts.LineData,CompCount int) *charts.Line {

	line := charts.NewLine()
	title := "客meanAddStd_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: "dataMin", Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("MeanAdd3Std", MeanAdd3Std).
		AddSeries("MaxAddStd", MaxAddStd).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: true,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}

func lineMainAnaly(fruits []string, BaseMaxValue1 []opts.LineData,BaseMaxValue2 []opts.LineData,FirstMax []opts.LineData,SecondMax []opts.LineData,CompCount int) *charts.Line {

	line := charts.NewLine()
	title := "主Base_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: "dataMin", Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("base1.62", BaseMaxValue1).
		AddSeries("maxFirst", FirstMax).
		AddSeries("maxSecond", SecondMax).
		AddSeries("base1.85", BaseMaxValue2).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: true,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}

func lineMiddleAnaly(fruits []string, BaseMaxValue1 []opts.LineData,BaseMaxValue2 []opts.LineData,FirstMax []opts.LineData,SecondMax []opts.LineData,CompCount int) *charts.Line {

	line := charts.NewLine()
	title := "平Base_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: "dataMin", Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("base1.62", BaseMaxValue1).
		AddSeries("maxFirst", FirstMax).
		AddSeries("maxSecond", SecondMax).
		AddSeries("base1.85", BaseMaxValue2).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: true,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}
func lineGuestAnaly(fruits []string, BaseMaxValue1 []opts.LineData,BaseMaxValue2 []opts.LineData,FirstMax []opts.LineData,SecondMax []opts.LineData,CompCount int) *charts.Line {

	line := charts.NewLine()
	title := "客Base_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: "dataMin", Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("base1.62", BaseMaxValue1).
		AddSeries("maxFirst", FirstMax).
		AddSeries("maxSecond", SecondMax).
		AddSeries("base1.85", BaseMaxValue2).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: true,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}

func lineMiddleP(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData,CompCount int,itemsWeide []opts.LineData,itemsBwin []opts.LineData) *charts.Line {

	line := charts.NewLine()
	title := "平赔_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: "dataMin", Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("12bet", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("crown", items12Bet).
		AddSeries("prinn", itemsWeide).
		AddSeries("bwin", itemsBwin).
		SetSeriesOptions(
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

func lineGuestP(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData,CompCount int,itemsWeide []opts.LineData,itemsBwin []opts.LineData) *charts.Line {

	line := charts.NewLine()
	title := "客赔_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		//charts.WithInitializationOpts(opts.Initialization{PageTitle:"Diagramms", Theme: types.ThemeWesteros, Width: "1200px", Height: "800px"}),
		//charts.WithLegendOpts(opts.Legend{Show: true, Left: "right", Orient: "vertical", Y: "100"}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: "dataMin", Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("12bet", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("crown", items12Bet).
		AddSeries("prinn", itemsWeide).
		AddSeries("bwin", itemsBwin).
		SetSeriesOptions(
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

func lineMainP(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData,CompCount int,itemsWeide []opts.LineData,itemsBwin []opts.LineData) *charts.Line {

	line := charts.NewLine()
	title := "主赔_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
			//Top:"0",
			//Right: "400",
		}),
		//charts.WithLegendOpts(opts.Legend{Show: true, Left: "right", Orient: "vertical", Y: "100"}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{Scale: true,Min: "dataMin", Max: "dataMax"}),
	)
	line.SetXAxis(fruits).
		AddSeries("12bet", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("crown", items12Bet).
		AddSeries("prinn", itemsWeide).
		AddSeries("bwin", itemsBwin).
		SetSeriesOptions(
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

func lineMiddleN(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData,CompCount int,itemsWeide []opts.LineData,itemsBwin []opts.LineData) *charts.Line {

	line := charts.NewLine()
	title := "平Norm_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
	)
	line.SetXAxis(fruits).
		AddSeries("12bet", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("crown", items12Bet).
		AddSeries("prinn", itemsWeide).
		AddSeries("bwin", itemsBwin).
		SetSeriesOptions(
		charts.WithLineChartOpts(opts.LineChart{
			ShowSymbol: true,
		}),
		charts.WithLabelOpts(opts.Label{
			Show: true,
		}),
	)
	return line
}

func lineGuestN(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData,CompCount int,itemsWeide []opts.LineData,itemsBwin []opts.LineData) *charts.Line {

	line := charts.NewLine()
	title := "客Norm_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		//charts.WithLineChartOpts(opts.LineChart{
		//
		//})
	)
	line.SetXAxis(fruits).
		AddSeries("12bet", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("crown", items12Bet).
		AddSeries("prinn", itemsWeide).
		AddSeries("bwin", itemsBwin).
		SetSeriesOptions(
		charts.WithLineChartOpts(opts.LineChart{
			ShowSymbol: true,
		}),
		charts.WithLabelOpts(opts.Label{
			Show: true,
		}),
	)
	return line
}

func lineMainN(items12Bet []opts.LineData,fruits []string,itemsYi []opts.LineData,itemsYing []opts.LineData,CompCount int,itemsWeide []opts.LineData,itemsBwin []opts.LineData) *charts.Line {

	line := charts.NewLine()
	title := "主Norm_"+ strconv.Itoa(CompCount)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		//charts.WithLegendOpts(opts.Legend{Left: "left",Top: "top"}), //legend是设定图例的
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeShine,
			Width: "720px",
			Height: "500px",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		//charts.WithGridOpts(opts.Grid{
		//	Left: "4%",
		//	Right: "5%",
		//	Bottom: "2%",
		//	Top: "15%",
		//	ContainLabel: true,
		//}),
		//charts.WithToolboxOpts(opts.Toolbox{
		//	Top: "4%",
		//	Left: "2%",
		//}),
	)
	line.SetXAxis(fruits).
		AddSeries("12bet", itemsYi).
		AddSeries("盈利", itemsYing).
		AddSeries("crown", items12Bet).
		AddSeries("prinn", itemsWeide).
		AddSeries("bwin", itemsBwin).
		SetSeriesOptions(
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
	sql := "SELECT MatchId,MainName,GuestName,CompCount,Main10Norm,Middle10Norm,Guest10Norm,Ep3,Ep1,Ep0,Main9Norm,Middle9Norm,Guest9Norm,OddTime,MatchTime,CoreMainNorm,CoreGuestNorm,CoreMiddleNorm,CoreMainP,CoreGuestP,CoreMiddleP,IntMainP,IntGuestP,IntMiddleP,PrinMainNorm,PrinGuestNorm,PrinMiddleNorm,PrinMainP,PrinGuestP,PrinMiddleP,B365MainNorm,B365GuestNorm,B365MiddleNorm,B365MainP,B365GuestP,B365MiddleP,DensityMain1,DensityGuest1,DensityMiddle1,DensityMain2,DensityGuest2,DensityMiddle2,DensityMain3,DensityGuest3,DensityMiddle3,MainMaxName,MiddleMaxName,GuestMaxName,BaseMainMaxValue1,BaseMainMaxValue2,BaseMiddleMaxValue1,BaseMiddleMaxValue2,BaseGuestMaxValue1 ,BaseGuestMaxValue2 ,MainFirstMax,MainSecondMax,MiddleFirstMax ,MiddleSecondMax,GuestFirstMax,GuestSecondMax,MainMaxAddStd,MiddleMaxAddStd,GuestMaxAddStd,MainMeanAdd3Std,MiddleMeanAdd3Std,GuestMeanAdd3Std from t_norm where MatchId = '" + nId + "' ORDER BY OddTime DESC"
	selDB, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	match := norm{}
	res := []norm{}
	for selDB.Next() {
		var MatchId, MainName, GuestName,MainMaxName,MiddleMaxName,GuestMaxName string
		var CompCount int
		var Main10Norm, Middle10Norm, Guest10Norm, MainP, MiddleP, GuestP, Main9Norm, Middle9Norm, Guest9Norm, CoreMainNorm, CoreGuestNorm, CoreMiddleNorm, CoreMainP, CoreGuestP, CoreMiddleP, IntMainP, IntGuestP, IntMiddleP,PrinMainNorm,PrinGuestNorm,PrinMiddleNorm,PrinMainP,PrinGuestP,PrinMiddleP,B365MainNorm,B365GuestNorm,B365MiddleNorm,B365MainP,B365GuestP,B365MiddleP,DensityMain1,DensityGuest1,DensityMiddle1,DensityMain2,DensityGuest2,DensityMiddle2,DensityMain3,DensityGuest3,DensityMiddle3,BaseMainMaxValue1,BaseMainMaxValue2,BaseMiddleMaxValue1,BaseMiddleMaxValue2,BaseGuestMaxValue1,BaseGuestMaxValue2,MainFirstMax,MainSecondMax,MiddleFirstMax,MiddleSecondMax,GuestFirstMax,GuestSecondMax,MainMaxAddStd,MiddleMaxAddStd,GuestMaxAddStd,MainMeanAdd3Std,MiddleMeanAdd3Std,GuestMeanAdd3Std float64
		var OddTime, MatchTime time.Time
		err = selDB.Scan(&MatchId, &MainName, &GuestName, &CompCount, &Main10Norm, &Middle10Norm, &Guest10Norm, &MainP, &MiddleP, &GuestP, &Main9Norm, &Middle9Norm, &Guest9Norm, &OddTime, &MatchTime, &CoreMainNorm, &CoreGuestNorm, &CoreMiddleNorm, &CoreMainP, &CoreGuestP, &CoreMiddleP, &IntMainP, &IntGuestP, &IntMiddleP, &PrinMainNorm,&PrinGuestNorm,&PrinMiddleNorm,&PrinMainP,&PrinGuestP,&PrinMiddleP,&B365MainNorm,&B365GuestNorm,&B365MiddleNorm,&B365MainP,&B365GuestP,&B365MiddleP,&DensityMain1,&DensityGuest1,&DensityMiddle1,&DensityMain2,&DensityGuest2,&DensityMiddle2,&DensityMain3,&DensityGuest3,&DensityMiddle3,&MainMaxName,&MiddleMaxName,&GuestMaxName,&BaseMainMaxValue1,&BaseMainMaxValue2,&BaseMiddleMaxValue1,&BaseMiddleMaxValue2,&BaseGuestMaxValue1,&BaseGuestMaxValue2,&MainFirstMax,&MainSecondMax,&MiddleFirstMax,&MiddleSecondMax,&GuestFirstMax,&GuestSecondMax,&MainMaxAddStd,&MiddleMaxAddStd,&GuestMaxAddStd,&MainMeanAdd3Std,&MiddleMeanAdd3Std,&GuestMeanAdd3Std)
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

		match.MainMaxName = MainMaxName
		match.MiddleMaxName = MiddleMaxName
		match.GuestMaxName = GuestMaxName
		match.BaseMainMaxValue1 = BaseMainMaxValue1
		match.BaseMainMaxValue2 = BaseMainMaxValue2
		match.BaseMiddleMaxValue1 = BaseMiddleMaxValue1
		match.BaseMiddleMaxValue2 = BaseMiddleMaxValue2
		match.BaseGuestMaxValue1 = BaseGuestMaxValue1
		match.BaseGuestMaxValue2 = BaseGuestMaxValue2
		match.MainFirstMax = MainFirstMax
		match.MainSecondMax = MainSecondMax
		match.MiddleFirstMax = MiddleFirstMax
		match.MiddleSecondMax = MiddleSecondMax
		match.GuestFirstMax = GuestFirstMax
		match.GuestSecondMax = GuestSecondMax
		match.MainMaxAddStd = MainMaxAddStd
		match.MiddleMaxAddStd = MiddleMaxAddStd
		match.GuestMaxAddStd = GuestMaxAddStd
		match.MainMeanAdd3Std = MainMeanAdd3Std
		match.MiddleMeanAdd3Std = MiddleMeanAdd3Std
		match.GuestMeanAdd3Std = GuestMeanAdd3Std
		res = append(res, match)
	}
	HandleData(res,nId,match.CompCount)
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
