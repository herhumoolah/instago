package main

import "fmt"

func main() {
	// Gets the command line options
	parseOptions()
	// Gets the config
	getConfig()
	// Tries to login
	login()

	if *post {
		// get data from file and  foreach id for get post syncFollowers()
		fmt.Println("post")
		instagramPost()
	} else if *stories {
		// get data from file and  foreach id for get stories loopTags()
		fmt.Println("stories")
	} else if *bdata {
		// get data from file and  foreach id for get bdata
		fmt.Println("bdata")
	} else if *all {
		// get data from file and  foreach id for get bdata, stories and post
		fmt.Println("all")
	}
}
