package main

import "go-mm-nrc-validator/validator"

func main() {
	v := validator.NewNRCValidator()
	v.ExtractNrcInfo("၁၂/မဂတ(နိုင်)၀၉၄၁၈၆")
	v.VerifyNRC("၄/ကပလ(နိုင်)၀၉၄၁၈၆")
	v.ValidateNRC("၄/ကပလ(နိုင်)၀၉၄၁၈၆")

}
