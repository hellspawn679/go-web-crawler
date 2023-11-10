package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	
)
func main(){
	blogTitle,err:=Getlatestblogtitle("https://www.freecodecamp.org/news/best-blogging-platforms-for-developers/")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("blog titles: ",blogTitle)

}

func Getlatestblogtitle(url string) (string,error){
	resp,err:=http.Get(url)
	if err!=nil{
		return "",err
	}
	fmt.Println(resp.Body)
	defer resp.Body.Close()
	doc,err :=goquery.NewDocumentFromReader(resp.Body)
	fmt.Println(resp.Body)
	if err!=nil{
		return "",err
	}
	titles:=""
	doc.Find(".post-title").Each(func(i int , s *goquery.Selection){
		titles+="-"+s.Text()+"\n"
	})
	return titles,nil
}