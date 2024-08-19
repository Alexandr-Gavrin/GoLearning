package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/cmplx"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/igrmk/treemap/v2"
)

var (
	module uint64     = 10000007
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Asin(-5 + 12i)
) // как объявлять глобальные переменные

const (
	tmpRune         rune = '▒'
	someInteresting int  = 228
) // Как объявлять глобальные константы

type longLONG uint64    // именованный тип данных (удобно когда функция должна принимать только этот тип данных)
type longerNum = uint64 // псевдоним

func ModifuSpaces(s, mode string) string {
	var newString string
	switch mode {
	case "dash":
		newString = strings.ReplaceAll(s, " ", "-")
		break
	case "underscore":
		newString = strings.ReplaceAll(s, " ", "_")
	default:
		newString = strings.ReplaceAll(s, " ", "*")
	}
	return newString
}

func sumInt(nums ...int) (sum, index int) {
	for _, elem := range nums {
		sum += elem
		index++
	}
	return
}

func Remove(nums []int, i int) []int {
	if i < 0 || i >= len(nums) {
		return nums
	}
	arr := make([]int, len(nums)-1)
	index := 0
	for k, el := range nums {
		if k != i {
			arr[index] = el
			index++
		}
	}
	return arr
}

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if userIDs == nil {
		return nil
	}
	if len(userIDs) <= 1 {
		return userIDs
	}
	uniq := make(map[int64]bool)
	var newArr []int64
	for _, el := range userIDs {
		uniq[el] = true
	}
	for key := range uniq {
		newArr = append(newArr, key)
	}
	sort.SliceStable(newArr, func(x, y int) bool { return int(newArr[x]) < int(newArr[y]) })
	return newArr
}

func UniqueUserIDs(userIDs []int64) []int64 {
	if len(userIDs) <= 1 {
		return userIDs
	}
	uniq := make(map[int64]bool)
	var newArr []int64
	for _, el := range userIDs {
		var ok bool
		_, ok = uniq[el]
		if !ok {
			newArr = append(newArr, el)
		}
		uniq[el] = true
	}
	return newArr
}

func latinLetters(s string) string {
	var newStr []string
	for _, el := range s {
		if unicode.Is(unicode.Latin, el) {
			newStr = append(newStr, string(el))
		}
	}
	return strings.Join(newStr, "")
}

func MergeNumberLists(numberLists ...[]int) []int {
	newArr := []int{}
	for _, elem := range numberLists {
		if elem != nil {
			newArr = append(newArr, elem...)
		}
	}
	return newArr
}

type Person struct {
	Age uint8
}

type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	mapYears := make(map[uint8]int)
	for _, el := range pl {
		mapYears[el.Age]++
	}
	return mapYears
}

type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// validation errors
var (
	errEmailRequired                = errors.New("email is required")
	errPasswordRequired             = errors.New("password is required")
	errPasswordConfirmationRequired = errors.New("password confirmation is required")
	errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	newUser := CreateUserRequest{}
	err := json.Unmarshal(requestBody, &newUser)
	if err == nil {
		if len(newUser.Email) == 0 {
			return CreateUserRequest{}, errEmailRequired
		} else if len(newUser.Password) == 0 {
			return CreateUserRequest{}, errPasswordRequired
		} else if len(newUser.PasswordConfirmation) == 0 {
			return CreateUserRequest{}, errPasswordConfirmationRequired
		} else if newUser.Password != newUser.PasswordConfirmation {
			return CreateUserRequest{}, errPasswordDoesNotMatch
		}
	}
	return newUser, err
}

type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
	return "validation error"
}

var (
	errBadConnection = errors.New("bad connection")
	errBadRequest    = errors.New("bad request")
)

const unknownErrorMsg = "unknown error"

func GetErrorMsg(err error) string {
	if errors.As(err, &nonCriticalError{}) {
		return ""
	} else if errors.Is(err, errBadConnection) {
		return errBadConnection.Error()
	} else if errors.Is(err, errBadRequest) {
		return errBadRequest.Error()
	} else {
		return "unknown error"
	}

}

type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

// errors
var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	job.IsFinished = true
	job.Merged = make(map[string]string)
	fl := false
	for _, el := range job.Dicts {
		if el == nil {
			fl = true
			continue
		}
		for key, val := range el {
			job.Merged[key] = val
		}
	}
	if len(job.Dicts) < 2 {
		job.IsFinished = true
		return job, errNotEnoughDicts
	} else if fl {
		return job, errNilDict
	} else {
		return job, nil
	}
}

func sum(arr []int, suml *int) {
	for _, el := range arr {
		(*suml) += el
	}
}

func MaxSum(nums1, nums2 []int) []int {
	sum1, sum2 := 0, 0
	go sum(nums1, &sum1)
	go sum(nums2, &sum2)
	time.Sleep(100 * time.Millisecond)
	if sum1 >= sum2 {
		return nums1
	} else {
		return nums2
	}
}

func SumWorker(numsCh chan []int, sumCh chan int) {

	for nums := range numsCh {
		sumCh <- takeSum(nums)
	}
}

func takeSum(arr []int) int {
	s := 0
	for _, el := range arr {
		s += el
	}
	return s
}

func main() {

	tr := treemap.New[int, string]()
	tr.Set(1, "World")
	tr.Set(0, "Hello")
	for it := tr.Iterator(); it.Valid(); it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
	// numsCh := make(chan []int)
	// sumCh := make(chan int)

	// go SumWorker(numsCh, sumCh)
	// numsCh <- nil
	// res := <-sumCh
	// fmt.Print(res)
}
