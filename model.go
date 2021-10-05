package main

type Categories struct {
	Totalrecords int    `json:"totalRecords"`
	Defaultzip   string `json:"defaultzip"`
	Catarray     []struct {
		Name                 string `json:"name"`
		Seotoken             string `json:"seoToken"`
		Uniqueid             string `json:"uniqueId"`
		Parentcataloggroupid string `json:"parentCatalogGroupID"`
		Fullimage            string `json:"fullImage"`
		Thumbnail            string `json:"thumbnail"`
		Tag                  string `json:"tag"`
		Isleafnode           string `json:"isleafnode"`
		Subcatarray          []struct {
			Name                 string `json:"name"`
			Seotoken             string `json:"seoToken"`
			Uniqueid             string `json:"uniqueId"`
			Parentcataloggroupid string `json:"parentCatalogGroupID"`
			Count                string `json:"count"`
			Fullimage            string `json:"fullImage"`
			Thumbnail            string `json:"thumbnail"`
			Tag                  string `json:"tag"`
			Isleafnode           string `json:"isleafnode"`
			Subcatarray          []struct {
				Name                 string `json:"name"`
				Seotoken             string `json:"seoToken"`
				Uniqueid             string `json:"uniqueId"`
				Parentcataloggroupid string `json:"parentCatalogGroupID"`
				Fullimage            string `json:"fullImage"`
				Thumbnail            string `json:"thumbnail"`
				Tag                  string `json:"tag"`
				Isleafnode           string `json:"isleafnode"`
			} `json:"subCatArray,omitempty"`
		} `json:"subCatArray"`
	} `json:"catArray"`
	Marketingspotdata []struct {
		Categoryid   string `json:"categoryId"`
		Categoryname string `json:"categoryName"`
		Seotoken     string `json:"seoToken"`
		Tag          string `json:"tag"`
	} `json:"MarketingSpotData"`
}

type Products struct {
	Xselldata      []interface{} `json:"xSellData"`
	Totalrecords   int           `json:"totalRecords"`
	Suggestionview []struct {
		Buyable          string `json:"buyable"`
		Manufacturer     string `json:"manufacturer,omitempty"`
		Name             string `json:"name"`
		SeoTokenNtk      string `json:"seo_token_ntk"`
		Numberofskus     int    `json:"numberOfskus"`
		SeoMetaDesc      string `json:"seo_meta_desc"`
		SeoMetaTitle     string `json:"seo_meta_title"`
		Availabilitytype string `json:"availabilityType"`
		Categoryid       string `json:"categoryId"`
		Categoryname     string `json:"categoryname"`
		Isitemshareable  string `json:"isItemShareable"`
		Uniqueid         string `json:"uniqueId"`
		Targeturl        string `json:"targetUrl"`
		Skus             []struct {
			Availabilitytype string `json:"availabilityType"`
			Homeorpup        string `json:"homeorpup"`
			Cod              string `json:"cod"`
			Grocerytype      string `json:"groceryType"`
			Defaultrank      string `json:"defaultRank"`
			Imagekey         string `json:"imagekey"`
			Binaryimgcode    string `json:"binaryimgcode"`
			Imgcode          string `json:"imgcode"`
			PriceMrp         string `json:"price_MRP"`
			PriceSale        string `json:"price_SALE"`
			SavePrice        string `json:"save_price"`
			Invstatus        string `json:"invstatus"`
			Maxquantity      string `json:"maxQuantity"`
			Defining         []struct {
				Volume string `json:"volume"`
			} `json:"defining"`
			Skuuniqueid     string   `json:"skuUniqueID"`
			Articlenumber   string   `json:"articleNumber"`
			Buyable         string   `json:"buyable"`
			DefaultVariant  string   `json:"default_variant"`
			Bulkquantity    string   `json:"bulkQuantity"`
			Bulkthreshold   string   `json:"bulkThreshold"`
			Exclusive       string   `json:"exclusive"`
			Minbulkquantity string   `json:"minBulkQuantity"`
			Giftitem        string   `json:"giftItem"`
			Name            string   `json:"name"`
			Tags            []string `json:"tags"`
			Ribbon          string   `json:"ribbon"`
		} `json:"skus"`
		Templatetype string `json:"templatetype"`
	} `json:"suggestionView"`
	Breadcrumb []struct {
		Value string `json:"value"`
		Label string `json:"label"`
		Type  string `json:"type_"`
	} `json:"breadcrumb"`
	Plpbanners string        `json:"plpBanners"`
	Facetview  []interface{} `json:"facetview"`
}
