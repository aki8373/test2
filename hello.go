package main

// If you're just starting out please see the "hello" sample instead. While
// this one is relatively simple, if has quite a few extra flags that can
// be set to control how the code behaves at runtime. So, this is great
// for debugging and exploring those options - but not great if you want
// to see the bare minimum needed to start an app.

// The main purpose of this is to run an App (http server), however, it
// can also be used as a Batch Job if the JOB_INDEX env var is set - which
// is set by the Code Engine batch processor. This can be useful if you want
// the same code to be used for both Apps and Jobs. In this respect it's
// very similar to the app-n-job sample, but this has all of the interesting
// debug/configuration flags that can be tweaked.

import (
	"fmt"
)

func main(){
	fmt.Println("Good morning")	
}

