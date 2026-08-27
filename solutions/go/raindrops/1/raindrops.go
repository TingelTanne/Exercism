package raindrops

import "strconv"

func Convert(number int) string {
	var sound = ""
	if(number % 3 == 0){
		sound += "Pling"
	}
	if(number % 5 == 0){
		sound += "Plang"
	}
	if(number % 7 == 0){
		sound += "Plong"
	}
	if(sound == ""){
		return strconv.Itoa(number)
	}else{
		return sound
	}
}
