package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/yu5429/mycoin/blockchain"
)

var templates *template.Template

const (
	templateDir string = "explorer/templates/"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", nil}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method { //http요청에 따라 다른 동작 수행
	case "GET": //add 랜더링
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST": //블록추가
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.Blockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect) //블록추가 후 홈으로 이동
	}
}

func Start(port int) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "/pages/*.gohtml"))     // /**/*.gohtml로 사용 모든 파일 선택이 안되서 2개로 나눠서 작성
	templates = template.Must(templates.ParseGlob(templateDir + "/partials/*.gohtml")) //Must는 에러가 있는지 확인하는 함수 발생시 panic처리
	handler.HandleFunc("/", home)                                                      //URL로 home 함수 호출
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler)) //string으로 넘겨준 포트번호를 통해 간단히 서버를 연다.
	//go는 Excetipon이 없어서 에러 발생시 log.Fatal을 통해 에러를 띄우고 panic 상태로 종료시켜줌
}
