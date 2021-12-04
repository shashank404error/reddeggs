package main

import (
	//"github.com/strike-official/go-sdk/strike"
	"os"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"	
	//"strings"
  
	"github.com/shashank404error/reddeggs/schema"
	//"github.com/shashank404error/reddeggs/database"
)

var base_url string = "https://efd6-2405-201-a40c-b0ba-386e-7bab-8f17-d115.ngrok.io"

func determineListenAddress() (string, error){
	port:=os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	return ":"+port,nil
} 

func main(){
	addr, err:=determineListenAddress()
	if err!=nil{
		log.Fatal(err)
	}
	r:=mux.NewRouter()
	//r.HandleFunc("/",index).Methods("POST")
	r.HandleFunc("/order",order).Methods("POST")
	http.Handle("/",r)
	if err:=http.ListenAndServe(addr,nil)
	err != nil{
		panic(err)
	}
}

func order(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

    var request schema.Strike_Meta_Request_Structure

    err := decoder.Decode(&request)
    if err != nil {
        panic(err)
    }

	fmt.Println(request)

	// name := request.Bybrisk_session_variables.Username
	// strike_object := strike.Create("pintrest",base_url+"/pintrest")
	
	// question_object := strike_object.Question("isShowMore").
	// QuestionCard().SetHeaderToQuestion(1, strike.FULL_WIDTH).AddTextRowToQuestion(strike.H4, "Hi "+name+", Here are some top memes from pintrest!", "#0088e3", false).
	// AddTextRowToQuestion(strike.H5, "⚠️Caution", "red", false).
	// AddTextRowToQuestion(strike.H5, "Content presented here are sole responsibility of the platfom pintrest and in no way are supported by the developer. Viewers discretion is advised.", "black", false)

	// // //Add answer
	// db := database.ConnectToRDS()
    // defer db.Close()

	// meme_url := database.GetMemes(db,"pintrest")
	// fmt.Println(meme_url)

	// answer_object := question_object.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION)
	
	// for _,url := range meme_url{

	// 	url = strings.Replace(url, "236x", "originals", 1)
	// 	answer_object = answer_object.AnswerCard().SetHeaderToAnswer(1,strike.FULL_WIDTH).
	// 	AddGraphicRowToAnswer(strike.PICTURE_ROW, []string{url},[]string{})
	// }

	// answer_object.AnswerCard().SetHeaderToAnswer(1,strike.HALF_WIDTH).
	// AddTextRowToAnswer(strike.H3, "Show More 👇", "#1064eb", true)


	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// //fmt.Println(string(strike_object.ToJson()))
	// w.Write(strike_object.ToJson())
}







