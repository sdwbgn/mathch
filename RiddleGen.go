package mathch

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var replMatch = [][]string{{"0", "6", "9"}, {"2", "3"}, {"3", "2", "5"}, {"4", "11"}, {"5", "3"}, {"6", "0", "9"}, {"9", "0", "6"},
	{"11", "4"}}
var remMatch = [][]string{{"6", "5"}, {"7", "1"}, {"8", "0", "6", "9"}, {"9", "5", "3"}, {"+", "-"}}
var addMatch = [][]string{{"0", "8"}, {"1", "7"}, {"3", "9"}, {"5", "6", "9"}, {"6", "8"}, {"9", "8"}, {"-", "+"}}

func Init() uint32 {
	rand.Seed(time.Now().Unix())
	return rand.Uint32()
}

func eval(str string) (int, error) {
	a := strings.Split(strings.Replace(strings.Replace(str, "-", "|-", -1), "+", "|+", -1), "|")
	if strings.Contains(str, "++") || strings.Contains(str, "--") || strings.Contains(str, "-+") || strings.Contains(str, "+-") {
		return 0, errors.New("double signs")
	}

	sum := 0
	for _, i := range a {
		val, err := strconv.Atoi(i)
		if err != nil {
			return 0, err
		}
		sum += val
	}
	return sum, nil
}

func MoveMatch(expr string, solved bool) ([]string, error) {
	var generated []string
	eqParts := strings.Split(expr, "=")
	if len(eqParts) != 2 {
		return []string{}, errors.New("invalid expression")
	}
	_, err := eval(eqParts[0])
	if err != nil {
		return []string{}, err
	}
	_, err = eval(eqParts[1])
	if err != nil {
		return []string{}, err
	}
	for i := 0; i < len(replMatch); i++ {
		curNum := strings.Index(expr[0:], replMatch[i][0])
		for curNum != -1 {
			for j := 1; j < len(replMatch[i]); j++ {
				changedStr := expr[:curNum] + replMatch[i][j] + expr[curNum+len(replMatch[i][0]):]
				eqParts := strings.Split(changedStr, "=")
				if len(eqParts) != 2 {
					continue
				}
				val1, err := eval(eqParts[0])
				if err != nil {
					continue
				}
				val2, err := eval(eqParts[1])
				if err != nil {
					continue
				}
				if (val1 == val2) == solved {
					generated = append(generated, changedStr)
				}
			}
			prCur := curNum + 1
			curNum = strings.Index(expr[curNum+1:], replMatch[i][0])
			if curNum != -1 {
				curNum += prCur
			}
		}
	}
	for i := 0; i < len(remMatch); i++ {
		curNum := strings.Index(expr[0:], remMatch[i][0])
		for curNum != -1 {
			for j := 1; j < len(remMatch[i]); j++ {
				changedStr := expr[:curNum] + remMatch[i][j] + expr[curNum+len(remMatch[i][0]):]
				for k := 0; k < len(addMatch); k++ {
					curNum2 := strings.Index(expr[0:], addMatch[k][0])
					for curNum2 != -1 {
						for l := 1; l < len(addMatch[k]); l++ {
							if curNum2 != curNum {
								changedStr2 := changedStr[:curNum2] + addMatch[k][l] + changedStr[curNum2+len(addMatch[k][0]):]
								eqParts := strings.Split(changedStr2, "=")
								if len(eqParts) != 2 {
									continue
								}
								val1, err := eval(eqParts[0])
								if err != nil {
									continue
								}
								val2, err := eval(eqParts[1])
								if err != nil {
									continue
								}
								if (val1 == val2) == solved {
									generated = append(generated, changedStr2)
								}
							}
						}
						prCur2 := curNum2 + 1
						curNum2 = strings.Index(expr[curNum2+1:], addMatch[k][0])
						if curNum2 != -1 {
							curNum2 += prCur2
						}
					}
				}

			}
			prCur := curNum + 1
			curNum = strings.Index(expr[curNum+1:], remMatch[i][0])
			if curNum != -1 {
				curNum += prCur
			}
		}
	}
	return generated, nil
}

func RandGenerate(diff int) string {
	var tasks []string
	for len(tasks) == 0 {
		expr := ""
		switch diff {
		case 1:
			if rand.Intn(2) == 0 {
				a := rand.Intn(10)
				b := rand.Intn(10)
				c := a + b
				expr = strconv.Itoa(a) + "+" + strconv.Itoa(b) + "=" + strconv.Itoa(c)
			} else {
				a := rand.Intn(10)
				b := rand.Intn(a + 1)
				c := a - b
				expr = strconv.Itoa(a) + "-" + strconv.Itoa(b) + "=" + strconv.Itoa(c)
			}
		case 2:
			if rand.Intn(2) == 0 {
				a := rand.Intn(90) + 10
				b := rand.Intn(90) + 10
				c := a + b
				expr = strconv.Itoa(a) + "+" + strconv.Itoa(b) + "=" + strconv.Itoa(c)
			} else {
				a := rand.Intn(90) + 10
				b := rand.Intn(a-9) + 10
				c := a - b
				expr = strconv.Itoa(a) + "-" + strconv.Itoa(b) + "=" + strconv.Itoa(c)
			}
		case 3:
			n := rand.Intn(4) + 3
			var ex []int
			ex = append(ex, rand.Intn(101))
			sum := ex[0]
			for i := 0; i < n; i++ {
				ex = append(ex, rand.Intn(201)-100)
				sum += ex[i+1]
			}
			if sum < 0 {
				val := rand.Intn(-sum+1) - sum
				sum += val
				ex = append(ex, val)
			}
			for _, i := range ex {
				if expr == "" {
					expr += strconv.Itoa(i)
					continue
				}
				if i > 0 {
					expr += "+" + strconv.Itoa(i)
				} else if i < 0 {
					expr += strconv.Itoa(i)
				} else if rand.Intn(2) == 0 {
					expr += "+0"
				} else {
					expr += "-0"
				}
			}
			expr += "=" + strconv.Itoa(sum)
		default:
			diff = rand.Intn(3) + 1
			return RandGenerate(diff)
		}
		var err error
		tasks, err = MoveMatch(expr, false)
		if err != nil {
			continue
		}
	}
	return tasks[rand.Intn(len(tasks))]
}

func Solve(riddle string, answer string) bool {
	solutions, err := MoveMatch(riddle, true)
	if err != nil {
		return false
	}
	for _, str := range solutions {
		if str == answer {
			return true
		}
	}
	return false
}
