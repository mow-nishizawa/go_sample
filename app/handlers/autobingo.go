package handlers

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var DrawNumber int
var DrawNumberList []int
var PlayerList []Player
var CardList [5][5]Card
var Terminate bool // false:実施中、true:終了

type Player struct {
	Name     string
	CardList [5][5]Card
	Message  string
}

type Card struct {
	CardNumber string
	Hit        int
}

func AutobingoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		players := getPlayers(r)

		if len(players) != 0 {
			PlayerList = createPlayerList(players)
		} else {
			PlayerList = []Player{}
		}

		param := getParam()
		log.Printf("%+v", param)

		t, err := template.ParseFiles("templates/autobingo.html")
		if err != nil {
			log.Fatal(err)
		}

		if err := t.Execute(w, param); err != nil {
			log.Fatal(err)
		}
	} else {
		DrawNumber = drawNumber()

		updateHit()
		// ビンゴ判定
		checkCardList()

		param := getParam()

		t, err := template.ParseFiles("templates/autobingo.html")
		if err != nil {
			log.Fatal(err)
		}

		if err := t.Execute(w, param); err != nil {
			log.Fatal(err)
		}
	}
}

// テンプレートに渡すパラメータを取得
func getParam() struct {
	PlayerList []Player
	DrawNumber int
	Terminate  bool
} {
	tmp := struct {
		PlayerList []Player
		DrawNumber int
		Terminate  bool
	}{
		PlayerList,
		DrawNumber,
		Terminate,
	}
	return tmp
}

// 入力された参加者を取得
func getPlayers(r *http.Request) []string {
	var players []string
	if len(r.FormValue("players")) != 0 {
		players = strings.Split(r.FormValue("players"), " ")
	}
	return players
}

// 参加者ごとにビンゴカードを作成
func createPlayerList(players []string) []Player {
	for count := 0; count < len(players); count++ {

		PlayerList = append(PlayerList, Player{Name: players[count], CardList: createCard()})
	}
	return PlayerList
}

// ビンゴカード作成
func createCard() [5][5]Card {
	numberList := createNumberList()

	var count int
	for row := 0; row < 5; row++ {
		for colum := 0; colum < 5; colum++ {
			if row == 2 && colum == 2 {
				CardList[row][colum].CardNumber = "FREE"
				CardList[row][colum].Hit = 1
				continue
			}
			CardList[row][colum].CardNumber = strconv.Itoa(numberList[count])
			count++
		}
	}
	return CardList
}

// ビンゴカードに使用する1-75のランダムな整数を作成
func createNumberList() []int {
	var numberList []int
	rand.Seed(time.Now().UnixNano())
	for {
		tmpNum := rand.Intn(76)
		if !include(numberList, tmpNum) && tmpNum != 0 {
			numberList = append(numberList, tmpNum)
		}
		if len(numberList) == 24 {
			break
		}
	}
	return numberList
}

// 1-75のランダムな整数を引いて保管する
func drawNumber() int {
	rand.Seed(time.Now().UnixNano())
	tmpNum := rand.Intn(76)
	for {
		if !include(DrawNumberList, tmpNum) && tmpNum != 0 {
			DrawNumberList = append(DrawNumberList, tmpNum)
			return tmpNum
		}
	}
}

// 重複チェック
func include(slice []int, target int) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}
	return false
}

// 引いた数字と一致するカードのHitを更新する
func updateHit() {
	drawNumber := strconv.Itoa(DrawNumber)
	for i, p := range PlayerList {
		for row := 0; row < 5; row++ {
			for colum := 0; colum < 5; colum++ {
				if p.CardList[row][colum].CardNumber == drawNumber {
					p.CardList[row][colum].Hit = 1
					PlayerList[i] = Player{Name: p.Name, CardList: p.CardList}
				}
			}
		}
	}
}

// リーチ・ビンゴの判定を行う
func checkCardList() {
	for i, p := range PlayerList {
		verticalCount := checkVertical(p.CardList)
		horizontalCount := checkHorizontal(p.CardList)
		obliqueDownRightCount := checkObliqueDownRight(p.CardList)
		obliqueDownLeftCount := checkObliqueDownLeft(p.CardList)
		if verticalCount == 5 || horizontalCount == 5 || obliqueDownRightCount == 5 || obliqueDownLeftCount == 5 {
			PlayerList[i] = Player{Name: p.Name, CardList: p.CardList, Message: p.Name + "さん、ビンゴです！"}
			Terminate = true
		} else if verticalCount == 4 || horizontalCount == 4 || obliqueDownRightCount == 4 || obliqueDownLeftCount == 4 {
			PlayerList[i] = Player{Name: p.Name, CardList: p.CardList, Message: p.Name + "さん、リーチです！"}
		}
	}
}

// 縦のチェック
func checkVertical(cardList [5][5]Card) int {
	var vertical0Count int
	var vertical1Count int
	var vertical2Count int
	var vertical3Count int
	var vertical4Count int
	for row := 0; row < 5; row++ {
		for colum := 0; colum < 5; colum++ {
			if cardList[colum][row].Hit == 1 {
				switch row {
				case 0:
					vertical0Count++
				case 1:
					vertical1Count++
				case 2:
					vertical2Count++
				case 3:
					vertical3Count++
				case 4:
					vertical4Count++
				}
			}
		}
	}
	if vertical0Count == 5 || vertical1Count == 5 || vertical2Count == 5 || vertical3Count == 5 || vertical4Count == 5 {
		return 5
	} else if vertical0Count == 4 || vertical1Count == 4 || vertical2Count == 4 || vertical3Count == 4 || vertical4Count == 4 {
		return 4
	}
	return 0
}

// 横のチェック
func checkHorizontal(cardList [5][5]Card) int {
	var horizontal0Count int
	var horizontal1Count int
	var horizontal2Count int
	var horizontal3Count int
	var horizontal4Count int
	for row := 0; row < 5; row++ {
		for colum := 0; colum < 5; colum++ {
			if cardList[row][colum].Hit == 1 {
				switch row {
				case 0:
					horizontal0Count++
				case 1:
					horizontal1Count++
				case 2:
					horizontal2Count++
				case 3:
					horizontal3Count++
				case 4:
					horizontal4Count++
				}
			}
		}
	}
	if horizontal0Count == 5 || horizontal1Count == 5 || horizontal2Count == 5 || horizontal3Count == 5 || horizontal4Count == 5 {
		return 5
	} else if horizontal0Count == 4 || horizontal1Count == 4 || horizontal2Count == 4 || horizontal3Count == 4 || horizontal4Count == 4 {
		return 4
	}
	return 0
}

// 右下がり斜めのチェック
func checkObliqueDownRight(cardList [5][5]Card) int {
	var obliqueDownRightCount int
	for i := 0; i < 5; i++ {
		if cardList[i][i].Hit == 1 {
			obliqueDownRightCount++
		}
	}
	return obliqueDownRightCount
}

// 左下がり斜めのチェック
func checkObliqueDownLeft(cardList [5][5]Card) int {
	var obliqueDownLeftCount int
	var colum int = 4
	for row := 0; row < 5; row++ {
		if cardList[row][colum].Hit == 1 {
			obliqueDownLeftCount++
		}
		colum--
	}
	return obliqueDownLeftCount
}
