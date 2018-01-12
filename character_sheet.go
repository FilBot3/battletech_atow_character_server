package main

type CharacterSheet struct {
	Attributes struct {
		Body struct {
			Link  int64 `json:"link"`
			Score int64 `json:"score"`
			Xp    int64 `json:"xp"`
		} `json:"body"`
		Charisma struct {
			Link  int64 `json:"link"`
			Score int64 `json:"score"`
			Xp    int64 `json:"xp"`
		} `json:"charisma"`
		Dexterity struct {
			Link  int64 `json:"link"`
			Score int64 `json:"score"`
			Xp    int64 `json:"xp"`
		} `json:"dexterity"`
		Edge struct {
			Link  int64 `json:"link"`
			Score int64 `json:"score"`
			Xp    int64 `json:"xp"`
		} `json:"edge"`
		Intelligence struct {
			Link  int64 `json:"link"`
			Score int64 `json:"score"`
			Xp    int64 `json:"xp"`
		} `json:"intelligence"`
		Reflex struct {
			Link  int64 `json:"link"`
			Score int64 `json:"score"`
			Xp    int64 `json:"xp"`
		} `json:"reflex"`
		Strength struct {
			Link  int64 `json:"link"`
			Score int64 `json:"score"`
			Xp    int64 `json:"xp"`
		} `json:"strength"`
		Will struct {
			Link  int64 `json:"link"`
			Score int64 `json:"score"`
			Xp    int64 `json:"xp"`
		} `json:"will"`
	} `json:"attributes"`
	Biography []struct {
		Name  string `json:"name"`
		Notes string `json:"notes"`
	} `json:"biography"`
	CombatData struct {
		ConditionMonitor struct {
			FatigueDamage  int64 `json:"fatigue_damage"`
			StandardDamage int64 `json:"standard_damage"`
			Stun           int64 `json:"stun"`
			Unconcious     int64 `json:"unconcious"`
		} `json:"condition_monitor"`
		Movement struct {
			Climb    int64 `json:"climb"`
			Crawl    int64 `json:"crawl"`
			RunEvade int64 `json:"run_evade"`
			Sprint   int64 `json:"sprint"`
			Swim     int64 `json:"swim"`
			Walk     int64 `json:"walk"`
		} `json:"movement"`
		PersonalArmor []struct {
			BarB     int64  `json:"bar_b"`
			BarE     int64  `json:"bar_e"`
			BarM     int64  `json:"bar_m"`
			BarX     int64  `json:"bar_x"`
			Location string `json:"location"`
			Name     string `json:"name"`
			Type     string `json:"type"`
		} `json:"personal_armor"`
		Weapons []struct {
			Ammo  int64  `json:"ammo"`
			Name  string `json:"name"`
			Notes string `json:"notes"`
			Range struct {
				Close   int64 `json:"close"`
				Extreme int64 `json:"extreme"`
				Long    int64 `json:"long"`
				Medium  int64 `json:"medium"`
			} `json:"range"`
			SkillAp int64 `json:"skill_ap"`
			SkillBd int64 `json:"skill_bd"`
		} `json:"weapons"`
	} `json:"combat_data"`
	Inventory struct {
		Cbills int64 `json:"cbills"`
		Items  []struct {
			DataStats string `json:"data_stats"`
			Location  string `json:"location"`
			Name      string `json:"name"`
		} `json:"items"`
	} `json:"inventory"`
	PersonalData struct {
		Affiliation string `json:"affiliation"`
		Extra       string `json:"extra"`
		Eyes        string `json:"eyes"`
		Hair        string `json:"hair"`
		Height      int64  `json:"height"`
		Name        string `json:"name"`
		Player      string `json:"player"`
		Weight      int64  `json:"weight"`
	} `json:"personal_data"`
	Skills []struct {
		Level int64  `json:"level"`
		Links string `json:"links"`
		Name  string `json:"name"`
		TnC   string `json:"tn_c"`
		Xp    int64  `json:"xp"`
	} `json:"skills"`
	Traits []struct {
		Name    string `json:"name"`
		PageRef int64  `json:"page_ref"`
		Points  int64  `json:"points"`
		Xp      int64  `json:"xp"`
	} `json:"traits"`
	VehicleData []struct {
		VehicleMass      int64  `json:"vehicle_mass"`
		VehicleModelName string `json:"vehicle_model_name"`
		VehicleNotes     string `json:"vehicle_notes"`
		VehicleTraits    string `json:"vehicle_traits"`
		VehicleType      string `json:"vehicle_type"`
	} `json:"vehicle_data"`
}
