package hotel

import "strings"

// InferBrand maps a hotel name to a sub-brand for multi-brand parent groups.
func InferBrand(parent, hotelName string) string {
	low := strings.ToLower(hotelName)
	switch parent {
	case "accor":
		return inferAccorBrand(low)
	case "ihg":
		return inferIHGBrand(low)
	case "hyatt":
		return inferHyattBrand(low)
	case "marriott":
		return inferMarriottBrand(low)
	case "hilton":
		return inferHiltonBrand(low)
	case "wyndham":
		return inferWyndhamBrand(low)
	case "bestwestern":
		return inferBestWesternBrand(low)
	case "radisson":
		return inferRadissonBrand(low)
	default:
		return parent
	}
}

func inferAccorBrand(low string) string {
	switch {
	case strings.Contains(low, "raffles"):
		return "Raffles"
	case strings.Contains(low, "fairmont"):
		return "Fairmont"
	case strings.Contains(low, "sofitel"):
		return "Sofitel"
	case strings.Contains(low, "mgallery"), strings.Contains(low, "m gallery"):
		return "MGallery"
	case strings.Contains(low, "pullman"):
		return "Pullman"
	case strings.Contains(low, "novotel"):
		return "Novotel"
	case strings.Contains(low, "mercure"):
		return "Mercure"
	case strings.Contains(low, "ibis budget"):
		return "Ibis Budget"
	case strings.Contains(low, "ibis styles"):
		return "Ibis Styles"
	case strings.Contains(low, "ibis"):
		return "Ibis"
	default:
		return "Accor"
	}
}

func inferIHGBrand(low string) string {
	switch {
	case strings.Contains(low, "six senses"):
		return "Six Senses"
	case strings.Contains(low, "intercontinental"):
		return "InterContinental"
	case strings.Contains(low, "kimpton"):
		return "Kimpton"
	case strings.Contains(low, "crowne plaza"):
		return "Crowne Plaza"
	case strings.Contains(low, "holiday inn express"):
		return "Holiday Inn Express"
	case strings.Contains(low, "holiday inn"):
		return "Holiday Inn"
	case strings.Contains(low, "hotel indigo"):
		return "Hotel Indigo"
	case strings.Contains(low, "vignette"):
		return "Vignette Collection"
	default:
		return "IHG"
	}
}

func inferHyattBrand(low string) string {
	switch {
	case strings.Contains(low, "andaz"):
		return "Andaz"
	case strings.Contains(low, "grand hyatt"):
		return "Grand Hyatt"
	case strings.Contains(low, "hyatt regency"):
		return "Hyatt Regency"
	case strings.Contains(low, "hyatt centric"):
		return "Hyatt Centric"
	case strings.Contains(low, "thompson"):
		return "Thompson Hotels"
	case strings.Contains(low, "dreams"):
		return "Dreams Resorts"
	case strings.Contains(low, "secrets"):
		return "Secrets Resorts"
	case strings.Contains(low, "zoëtry"), strings.Contains(low, "zoetry"):
		return "Zoëtry"
	case strings.Contains(low, "alua"):
		return "Alua Hotels"
	default:
		return "Hyatt"
	}
}

func inferMarriottBrand(low string) string {
	switch {
	case strings.Contains(low, "ritz-carlton"), strings.Contains(low, "ritz carlton"):
		return "The Ritz-Carlton"
	case strings.Contains(low, "st. regis"), strings.Contains(low, "st regis"):
		return "St. Regis"
	case strings.Contains(low, "w "):
		return "W Hotels"
	case strings.Contains(low, "jw marriott"):
		return "JW Marriott"
	case strings.Contains(low, "edition"):
		return "Edition"
	case strings.Contains(low, "luxury collection"):
		return "Luxury Collection"
	case strings.Contains(low, "westin"):
		return "Westin"
	case strings.Contains(low, "sheraton"):
		return "Sheraton"
	case strings.Contains(low, "le méridien"), strings.Contains(low, "le meridien"):
		return "Le Méridien"
	case strings.Contains(low, "renaissance"):
		return "Renaissance Hotels"
	case strings.Contains(low, "autograph"):
		return "Autograph Collection"
	case strings.Contains(low, "tribute portfolio"):
		return "Tribute Portfolio"
	case strings.Contains(low, "ac hotel"):
		return "AC Hotels"
	case strings.Contains(low, "aloft"):
		return "Aloft"
	case strings.Contains(low, "moxy"):
		return "Moxy"
	case strings.Contains(low, "courtyard"):
		return "Courtyard by Marriott"
	case strings.Contains(low, "residence inn"):
		return "Residence Inn"
	default:
		return "Marriott"
	}
}

func inferHiltonBrand(low string) string {
	switch {
	case strings.Contains(low, "waldorf"):
		return "Waldorf Astoria"
	case strings.Contains(low, "conrad"):
		return "Conrad"
	case strings.Contains(low, "doubletree"):
		return "DoubleTree by Hilton"
	case strings.Contains(low, "canopy"):
		return "Canopy by Hilton"
	case strings.Contains(low, "curio"):
		return "Curio Collection"
	case strings.Contains(low, "hampton"):
		return "Hampton by Hilton"
	default:
		return "Hilton"
	}
}

func inferWyndhamBrand(low string) string {
	switch {
	case strings.Contains(low, "ramada"):
		return "Ramada"
	case strings.Contains(low, "tryp"):
		return "Tryp"
	case strings.Contains(low, "dolce"):
		return "Dolce by Wyndham"
	default:
		return "Wyndham"
	}
}

func inferBestWesternBrand(low string) string {
	switch {
	case strings.Contains(low, "best western premier"):
		return "Best Western Premier"
	case strings.Contains(low, "best western plus"):
		return "Best Western Plus"
	default:
		return "Best Western"
	}
}

func inferRadissonBrand(low string) string {
	switch {
	case strings.Contains(low, "radisson collection"):
		return "Radisson Collection"
	case strings.Contains(low, "radisson blu"):
		return "Radisson Blu"
	case strings.Contains(low, "radisson red"):
		return "Radisson RED"
	case strings.Contains(low, "park inn"):
		return "Park Inn by Radisson"
	default:
		return "Radisson"
	}
}
