package main

import "testing"

func Test_parseBlock(t *testing.T) {
	type args struct {
		block string
	}
	tests := []struct {
		name           string
		args           args
		wantTitle      string
		wantUrl        string
		wantAbstract   string
		wantCountLinks int
	}{
		{"test block",
			args{block: `<doc><title>Albedo</title><url>https://en.wikipedia.org/wiki/Albedo</url><abstract>Albedo (pronounced ; , meaning 'whiteness')</abstract><links><sublink linktype="nav"><anchor>Terrestrial albedo</anchor><link>https://en.wikipedia.org/wiki/Albedo#Terrestrial_albedo</link></sublink><sublink linktype="nav"><anchor>White-sky, black-sky, and blue-sky albedo</anchor><link>https://en.wikipedia.org/wiki/Albedo#White-sky,_black-sky,_and_blue-sky_albedo</link></sublink><sublink linktype="nav"><anchor>External links</anchor><link>https://en.wikipedia.org/wiki/Albedo#External_links</link></sublink></links></doc>`},
			"Albedo",
			"https://en.wikipedia.org/wiki/Albedo",
			"Albedo (pronounced ; , meaning 'whiteness')",
			3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTitle, gotUrl, gotAbstract, gotCountLinks := parseBlock(tt.args.block)
			if gotTitle != tt.wantTitle {
				t.Errorf("parseBlock() gotTitle = %v, want %v", gotTitle, tt.wantTitle)
			}
			if gotUrl != tt.wantUrl {
				t.Errorf("parseBlock() gotUrl = %v, want %v", gotUrl, tt.wantUrl)
			}
			if gotAbstract != tt.wantAbstract {
				t.Errorf("parseBlock() gotAbstract = %v, want %v", gotAbstract, tt.wantAbstract)
			}
			if gotCountLinks != tt.wantCountLinks {
				t.Errorf("parseBlock() gotCountLinks = %v, want %v", gotCountLinks, tt.wantCountLinks)
			}
		})
	}
}
