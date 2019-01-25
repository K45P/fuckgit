package main
import "fmt"



func fibbonacci_sender(factor chan <- int, result <- chan int, finished chan bool){
	j := 0;
	prev1 := 0;
	prev2 := 1;
	
	factor <- prev1;
	factor <- prev2;

	for j = 0; j <10; j++{
		prev1 = prev2;
		prev2 = <-result;
		fmt.Println("sending factors");
		factor <- prev1;
		factor <- prev2;
		
	}
	finished <- true;
}

func fibbonacci_calc(factor chan int, result  chan int){
	fmt.Println("recieving factors");
	for{															//Vil kjøre loopen fram til finished = true
		msg1 := <-factor;
		msg2 := <-factor;
		fmt.Println("factors recieved, and they are: ", msg1, " , ", msg2);
		res := msg1 + msg2;
		fmt.Println("sending result: ", res);
		result <- res;
	}
	
	fmt.Println("result sent");

}


//tesp
//loooooool




func main() {
	factor := make(chan int, 2);
	result := make(chan int);
	finished := make(chan bool);
	
	go fibbonacci_sender(factor, result, finished);
	go fibbonacci_calc(factor, result);
	<-finished;
	fmt.Println("Program complete");
}
