package main

func GetIndex() HtmlData {
	return HtmlData {
		Title: "Index",
		Layout: "layout.html",
		Imports: []string {
			"index/index.html",
		},
	}
}

func GetIndex_Partial() HtmlData {
	return HtmlData {
		Title: "",
		Layout: "",
		Imports: []string {
			"index/partial.html",
		},
	}
}
