package pojo

import (
	"time"
)

type Matchhistoryeuro202181 struct {
	Id            int64     `xorm:"pk autoincr comment('主键Id') BIGINT"`
	Compid        int       `xorm:"INT"`
	Matchid       string    `xorm:"index VARCHAR(20)"`
	Sp3           float64   `xorm:"comment('Sp3') index(IDX_t_euro_his_Sp) DOUBLE"`
	Sp1           float64   `xorm:"comment('Sp1') index(IDX_t_euro_his_Sp) DOUBLE"`
	Sp0           float64   `xorm:"comment('Sp0') index(IDX_t_euro_his_Sp) DOUBLE"`
	Sk3           float64   `xorm:"DOUBLE"`
	Sk1           float64   `xorm:"DOUBLE"`
	Sk0           float64   `xorm:"DOUBLE"`
	Spayout       float64   `xorm:"comment('率') DOUBLE"`
	Sodddate      time.Time `xorm:"comment('数据时间') DATETIME"`
	Ep3           float64   `xorm:"comment('Ep3') index(IDX_t_euro_his_Ep3) DOUBLE"`
	Ep1           float64   `xorm:"comment('Ep1') index(IDX_t_euro_his_Ep3) DOUBLE"`
	Ep0           float64   `xorm:"comment('Ep0') index(IDX_t_euro_his_Ep3) DOUBLE"`
	Ek3           float64   `xorm:"DOUBLE"`
	Ek1           float64   `xorm:"DOUBLE"`
	Ek0           float64   `xorm:"DOUBLE"`
	Epayout       float64   `xorm:"comment('率') DOUBLE"`
	Eodddate      time.Time `xorm:"comment('数据时间') DATETIME"`
	Onep3         float64   `xorm:"comment('Ep3') index(IDX_t_euro_his_Onep) DOUBLE"`
	Onep1         float64   `xorm:"comment('Ep1') index(IDX_t_euro_his_Onep) DOUBLE"`
	Onep0         float64   `xorm:"comment('Ep0') index(IDX_t_euro_his_Onep) DOUBLE"`
	Onek3         float64   `xorm:"DOUBLE"`
	Onek1         float64   `xorm:"DOUBLE"`
	Onek0         float64   `xorm:"DOUBLE"`
	Onepayout     float64   `xorm:"comment('率') DOUBLE"`
	Oneodddate    time.Time `xorm:"comment('数据时间') DATETIME"`
	Twop3         float64   `xorm:"comment('Ep3') index(IDX_t_euro_his_Twop) DOUBLE"`
	Twop1         float64   `xorm:"comment('Ep1') index(IDX_t_euro_his_Twop) DOUBLE"`
	Twop0         float64   `xorm:"comment('Ep0') index(IDX_t_euro_his_Twop) DOUBLE"`
	Twok3         float64   `xorm:"DOUBLE"`
	Twok1         float64   `xorm:"DOUBLE"`
	Twok0         float64   `xorm:"DOUBLE"`
	Twopayout     float64   `xorm:"comment('率') DOUBLE"`
	Twoodddate    time.Time `xorm:"comment('数据时间') DATETIME"`
	Averagep3     float64   `xorm:"comment('Ep3') index(IDX_t_euro_his_Averagep) DOUBLE"`
	Averagep1     float64   `xorm:"comment('Ep1') index(IDX_t_euro_his_Averagep) DOUBLE"`
	Averagep0     float64   `xorm:"comment('Ep0') index(IDX_t_euro_his_Averagep) DOUBLE"`
	Averagek3     float64   `xorm:"DOUBLE"`
	Averagek1     float64   `xorm:"DOUBLE"`
	Averagek0     float64   `xorm:"DOUBLE"`
	Averagepayout float64   `xorm:"comment('赔付率') DOUBLE"`
	Createtime    time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') index DATETIME"`
	Updatetime    time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	Mark          string    `xorm:"comment('结果') VARCHAR(20)"`
	Ext           string    `xorm:"comment('扩展') VARCHAR(36)"`
}
