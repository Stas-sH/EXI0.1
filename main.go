package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	//"myProj/xsd"
	//"myProj/xmlfmt"
)

func main() {
	/*
	   	xml1 := `<root><this><is>a</is><test /><message><!-- with comment --><org><cn>Some org-or-other</cn><ph>Wouldnt you like to know</ph></org><contact><fn>Pat</fn><ln>Califia</ln></contact></message></this></root>`
	   	x := xmlfmt.FormatXML(xml1, "\t", "  ")
	   	print(x)

	   	// If the XML Comments have nested tags in them
	   	xml1 = `<book> <author>Fred</author>
	   <!--
	   <price>20</price><currency>USD</currency>
	   -->
	    <isbn>23456</isbn> </book>`
	   	x = xmlfmt.FormatXML(xml1, "", "  ", true)
	   	print(x)
	*/

	file, err := os.Open("testXmlFile.txt")
	//var xml1 string
	if err != nil {
		err = errors.New("reading file error")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fle := bufio.NewScanner(file)
	for fle.Scan() {
		//xml1 = xml1 + fle.Text()
		fmt.Println(fle.Text())
	}
	//x := xmlfmt.FormatXML(xml1, "\t", "  ")
	//print(x)
	fmt.Printf("%T", file)
}
