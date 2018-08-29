package example

import "github.com/hyhhhhhh/citest/awesome"

var result string

func Setup() {

    // Comment

    result = awesome.Smile()

}

func GetResult() string {

    /*
    Comment
    */

    return result

}
