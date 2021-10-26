//
//                       _oo0oo_
//                      o8888888o
//                      88" . "88
//                      (| -_- |)
//                      0\  =  /0
//                    ___/`---'\___
//                  .' \\|     |// '.
//                 / \\|||  :  |||// \
//                / _||||| -:- |||||- \
//               |   | \\\  -  /// |   |
//               | \_|  ''\---/''  |_/ |
//               \  .-\__  '-'  ___/-. /
//             ___'. .'  /--.--\  `. .'___
//          ."" '<  `.___\_<|>_/___.' >' "".
//         | | :  `- \`.;`\ _ /`;.`/ - ` : | |
//         \  \ `_.   \_ __\ /__ _/   .-` /  /
//     =====`-.____`.___ \_____/___.-`___.-'=====
//                       `=---='
//
//
//     ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//               佛祖保佑         永無 BUG
//
package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
)

func ticket(seed int64) (day int) {
	rand.Seed(seed)
	speed := rand.Intn(46) + 16	// 隨機速度
	day = int (math.Ceil(float64 (62100000 / (speed * 60 * 60 * 24))))
	return
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 隨機種子

	Spaceline := [3] string {	// Spaceline 航空公司
		"Space Adventures", 
		"SpaceX", 
		"Virgun Galactic",
	}
	TripType := [2] string {	// 隨機決定單程或往返
		"Round - trip",
		"One - way",
	}

	fmt.Printf("%-20s%-10s%-15s%-10s\n","Spaceline", "Days", "Trip Type", "Price")
	fmt.Println("==================================================")
	for i := 0; i < 9; i ++ {
		day := ticket(int64(i))
		TripTypeRand := TripType[rand.Intn(2)]
		price := 0
		if TripTypeRand == "One - way" {
			price = rand.Intn(14) +36
		} else {
			price = (rand.Intn(14) +36) * 2
		}
		fmt.Printf("%-20s%-10d%-15s%-10d\n", Spaceline[rand.Intn(3)], day, TripTypeRand, price)
	}
}