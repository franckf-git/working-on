package main

import (
	"flag"
	"fmt"
)

func main() {

	var verbose bool
	var download int
	var url string

	verboseDesc := "add verbose logs to ouput"
	flag.BoolVar(&verbose, "v", false, verboseDesc)
	flag.BoolVar(&verbose, "verbose", false, verboseDesc)
	flag.IntVar(&download, "d", 1, "number of parallels downloads")
	flag.IntVar(&download, "download", 1, "number of parallels downloads")
	flag.StringVar(&url, "u", "", "url to process")
	flag.StringVar(&url, "url", "", "url to process")

	otherWay := flag.Bool("o", false, "otherway to declare a flag variable")

	flag.Parse()
	args := flag.Args()

	fmt.Println("are you verbose ?", verbose)
	fmt.Println("how many downloads ?", download)
	fmt.Println("what url :", url)
	fmt.Println("other way, but must be call with dereference :", *otherWay)
	fmt.Println("extra args :", args)

	if verbose {
		fmt.Println(`
		Donut jelly sugar plum sweet jelly beans. Jujubes dessert tootsie roll cake gummies gummies cake. Apple pie pie danish chocolate cake liquorice caramels lollipop cake sweet. Cotton candy fruitcake cake sesame snaps gummies apple pie. Bonbon chocolate bar dessert powder sweet jelly. Chocolate bar chocolate cake powder lollipop carrot cake fruitcake. Bonbon danish lollipop sweet cotton candy. Sesame snaps gummies dragée dessert gummi bears candy macaroon pie. Sweet candy canes cookie halvah pie sesame snaps. Lollipop halvah danish carrot cake danish cotton candy cake halvah chocolate bar. Caramels apple pie tart bear claw pie ice cream sesame snaps. Muffin halvah dessert soufflé soufflé bonbon cheesecake apple pie. Cheesecake biscuit soufflé carrot cake danish. Apple pie bear claw gingerbread oat cake candy canes candy topping sesame snaps.
		`)
	}
	/*
		// more elaborate with parsing but flag.Func is in 1.16
		fs := flag.NewFlagSet("ExampleFunc", flag.ContinueOnError)
		fs.SetOutput(os.Stdout)
		var ip net.IP
		fs.Func("ip", "`IP address` to parse", func(s string) error {
			ip = net.ParseIP(s)
			if ip == nil {
				return errors.New("could not parse IP")
			}
			return nil
		})
		fs.Parse([]string{"-ip", "127.0.0.1"})
		fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())

		// 256 is not a valid IPv4 component
		fs.Parse([]string{"-ip", "256.0.0.1"})
		fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())
	*/
}
