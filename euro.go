MinuteDiff int
processer.MinuteDiff = minuteDiff

if(this.MinuteDiff>=0 && this.MinuteDiff<=10) {
		if(norm.CoreMainNorm > 0.7 || norm.CoreGuestNorm>0.7 || norm.Main10Norm>0.7 || norm.Guest10Norm>0.7 || norm.Main9Norm> 0.7 || norm.Guest9Norm>0.7|| norm.B365MainNorm>0.7 || norm.B365GuestNorm>0.7){
			//li/crown/yingli/bwin
		}else{
			his := new(pojo.MatchHis)
			last := new(pojo.MatchLast)
			his.Id = matchId
			last.Id = matchId
			this.EuroLastService.Del(last)
			this.EuroHisService.Del(his)
			return
		}
	}



