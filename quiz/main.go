package main

import (
	"fmt"
	"flag"
	"time"
	"bufio"
	"os"
	"strings"
)

func min(a, b int) int{
	if a<b{
		return a
	}
	return b
}
func ask(c chan string){
	var answer string
	fmt.Scan(&answer)
	c<-answer
}
func main(){

	filename:=flag.String("f","problems.csv","filename path")
	lim:=flag.Int("l",10,"how many quiz do you want")
	tim:=flag.Int("t",5,"timer of each problem")
	flag.Parse()
	
	//changed back to value from pointer so that it's easy to deal 
	var limit,times =*lim,time.Duration(*tim)
	limit=min(limit,13) // cuz our dataset only have 12 max questions	
	file,err:=os.Open(*filename);
	if(err!=nil){
		fmt.Println("Can't open this file")
	}
	count,correct:=0,0
	defer file.Close()
	fileScan:=bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)
	
	//file reading 

	for fileScan.Scan(){
		if(count>=limit){
			break;
		}
		
		line:=fileScan.Text()
		lines:=strings.Split(line,",");
		question,answer:=lines[0],lines[1]
		answer=strings.TrimSpace(answer)
		fmt.Println("sir please tell me",question)
		timer:=time.NewTimer(times*time.Second)
		c:=make(chan string)
		go ask(c)
		flag:=1
		select{
			case <-timer.C:
				fmt.Println("Time's up")
				flag=0
			case input:=<-c:
			if(answer==input) {
		//		fmt.Println("ok")
				correct+=1
				}
			}
		if(flag==0){
			break;
		}
		count+=1;
	}
	
	fmt.Println("Your point is",correct)
}
