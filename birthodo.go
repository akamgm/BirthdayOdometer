package birthodo

import (
	"fmt"
	"net/http"
	"time"
)

const CIRCUMFERENCE_OF_EARTH_ORBIT = 584000000
const EARTH_TRAVEL_PER_DAY = CIRCUMFERENCE_OF_EARTH_ORBIT / 365

func init() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/calc", calcHandler)
}

const birthdayForm = `
<html>
  <body>
    <form action="/calc" method="post">
      <div><input type="date" name="bday"/></div>
      <div><input type="submit" value="Calculate Miles"></div>
    </form>
  </body>
</html>
`

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, birthdayForm)
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	bday := r.FormValue("bday")
	milage, _ := milesTraveledSinceBirthday(bday)
	fmt.Fprintf(w, "You've traveled %d miles! Wow!", milage)
}

func milesTraveled(inDays int64) int64 {
	return EARTH_TRAVEL_PER_DAY * inDays
}

func milesTraveledSinceBirthday(inDate string) (int64, error) {
	birthday, err := time.Parse("2006-01-02", inDate)
	if err != nil {
		fmt.Errorf("Error: ", err)
		return 0, err
	}

	dur := time.Now().Sub(birthday)
	return milesTraveled(int64(dur.Hours() / 24)), nil

}
