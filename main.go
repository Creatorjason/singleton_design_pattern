package main


import (
	"fmt"
	"sync"
)

// Implementing the singleton design pattern
var lock = &sync.Mutex{}

// struct, you want a single instance of 
type single struct{}

// single instance of the struct
var singleInstance *single

// getInstance returns only a single instance of the struct
// every client that needs this instance will have to call this method to get it
func getInstance() *single{
	// check to see if the instance is already created
	if singleInstance == nil{
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil{
			fmt.Println("Creating single instance now")
			singleInstance = &single{}
		}else{
			fmt.Println("Single Instance already created-1")
		}
	}else{
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}


// Above code ensures that only one instance of the single struct is created. Some point worth noting. 
// 
// There is a check at the start for nil singleInstance. 
// This is to prevent the expensive lock operations every time getinstance() method is called. 
// If this check fails then it means that singleInstance is already created
// The singleInstance is created inside the lock.
// There is another check for nil singleIinstance after the lock is acquired. 
// This is to make sure that if more than one goroutine bypass the first check then only 
// one goroutine is able to create the singleton instance otherwise each of the goroutine will 
// create its own instance of the single struct.



func main() {
	hashid()
}

// TODO 
// Other methods
// Using INIT()

// init() function
// We can create a single instance inside the init function. 
// This is only applicable if the early initialization of the object is ok. 
// The init function is only called once per file in a package,  
// so we can be sure that only a single instance will be created.
 
// OR using sync.Once

var once sync.Once
func getInstanceUsingSyncOnce() *single{
	if singleInstance == nil{
		once.Do(func(){
			fmt.Println("Creating single instance now")
			singleInstance = &single{}
		} )

	}else{
		fmt.Println("Single instance already created - 2")
	}
	return singleInstance
}