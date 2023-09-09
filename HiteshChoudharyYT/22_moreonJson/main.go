package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json videos!!")
	encodeJson()
	decodeJson()
}

func encodeJson() {
	lcoCources := []course{
		{"ReactJs BootCamp", 299, "learncodeonline.in", "xyz234", []string{"web-dev", "js"}},
		{"MERN BootCamp", 199, "learncodeonline.in", "abc123", []string{"fullstack", "mern"}},
		{"Angular BootCamp", 299, "learncodeonline.in", "pqr456", nil},
	}

	// package this data as a Json data
	// Marshall : struct data is converted into JSON
	// UnMarshall : JSON data to string
	// The methods returned data in byte format and we need to change returned data into JSON or String.
	finalJson, err := json.MarshalIndent(lcoCources, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", &finalJson)

	// Consume Json Data

}

func decodeJson() {
	jsonDatafromWeb := []byte(`
		{
			"coursename": "MERN BootCamp",
			"price": 199,
			"website": "learncodeonline.in",
			"tags": ["fullstack","mern"]
		}
	`)

	// Here I have a struct created and want the data to be converted in this format
	var lcoCourse course

	checkValid := json.Valid(jsonDatafromWeb)

	// %#v for printing a struct
	if checkValid {
		fmt.Println("Json was Valid.")
		json.Unmarshal(jsonDatafromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID!")
	}

	// Sometimes i dont want things in a struct just want key value pairs
	// Below i have guarantee key is string but value i am not sure so used interface
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDatafromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is: %v Val is: %v Type is: %T\n", k, v, v)
	}
}
