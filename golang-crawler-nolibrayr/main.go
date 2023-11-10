package main
import(
	"fmt"
	"net/http"
	"os"
    "strings"
	"golang.org/x/net/html"
)


func getHref(t html.Token)(ok bool ,href string){
	for _,a:=range t.Attr{
		if a.Key=="href"{
			href=a.Val
			ok=true
		}
	
	}
	return  
}

func crawl(url string ,ch chan string ,chFinished chan bool){
	resp,err:=http.Get(url)
	 defer func(){
		chFinished<-true
	 }()
	if err!=nil{
		fmt.Println("ERROR: Failed to crawl \""+url+"\"")
		return 
	}
	b := resp.Body
	defer b.Close()
	z :=html.NewTokenizer(b)
	for {
		tt:=z.Next()
		switch {
			case tt ==html.ErrorToken:
				return 
			case tt ==html.StartTagToken:
				t :=z.Token()
				isAnchor:=t.Data=="a"
				if !isAnchor{
					continue
				}
				ok, url:= getHref(t)
				if !ok{
					continue
				}
				hasProto:=strings.Index(url,"http")==0
				if hasProto{
					ch<-url
				}	
		}
	}
}
func main(){
	foundUrls:=make(map[string]bool)
	seedurls:=os.Args[1:]
	churls:=make(chan string)
	chfinished:=make(chan bool)

	for _,url:=range seedurls{
		go crawl(url,churls,chfinished)
	}
	for c:=0; c<len(seedurls);{
		select{
		case url :=<-churls:
			foundUrls[url]=true
		case<-chfinished:
			c++
		}
	}
	fmt.Println("\nFound",len(foundUrls),"unique urls:\n")

	for url,_:=range foundUrls{
		fmt.Println(" - "+url)
	}
	

}