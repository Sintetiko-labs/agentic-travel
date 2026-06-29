package hotel
import "strings"
func InferBrand(parent, hotelName string) string {
 low := strings.ToLower(hotelName)
 switch parent {
 case "accor": return inferAccorBrand(low)
 case "ihg": return inferIHGBrand(low)
 case "hyatt": return inferHyattBrand(low)
 case "marriott": return inferMarriottBrand(low)
 case "hilton": return inferHiltonBrand(low)
 case "wyndham": return inferWyndhamBrand(low)
 case "bestwestern": return inferBestWesternBrand(low)
 case "radisson": return inferRadissonBrand(low)
 default: return parent
 }}
func inferAccorBrand(low string) string {
 switch {
 case strings.Contains(low,"raffles"): return "Raffles"
 case strings.Contains(low,"fairmont"): return "Fairmont"
 case strings.Contains(low,"sofitel"): return "Sofitel"
 case strings.Contains(low,"mgallery"): return "MGallery"
 case strings.Contains(low,"pullman"): return "Pullman"
 case strings.Contains(low,"novotel"): return "Novotel"
 case strings.Contains(low,"mercure"): return "Mercure"
 case strings.Contains(low,"ibis budget"): return "Ibis Budget"
 case strings.Contains(low,"ibis styles"): return "Ibis Styles"
 case strings.Contains(low,"ibis"): return "Ibis"
 default: return "Accor" }}
func inferIHGBrand(low string) string {
 switch {
 case strings.Contains(low,"six senses"): return "Six Senses"
 case strings.Contains(low,"intercontinental"): return "InterContinental"
 case strings.Contains(low,"kimpton"): return "Kimpton"
 case strings.Contains(low,"crowne plaza"): return "Crowne Plaza"
 case strings.Contains(low,"holiday inn express"): return "Holiday Inn Express"
 case strings.Contains(low,"holiday inn"): return "Holiday Inn"
 case strings.Contains(low,"hotel indigo"): return "Hotel Indigo"
 default: return "IHG" }}
func inferHyattBrand(low string) string {
 switch {
 case strings.Contains(low,"andaz"): return "Andaz"
 case strings.Contains(low,"grand hyatt"): return "Grand Hyatt"
 case strings.Contains(low,"hyatt regency"): return "Hyatt Regency"
 default: return "Hyatt" }}
func inferMarriottBrand(low string) string {
 switch {
 case strings.Contains(low,"ritz"): return "The Ritz-Carlton"
 case strings.Contains(low,"sheraton"): return "Sheraton"
 case strings.Contains(low,"westin"): return "Westin"
 case strings.Contains(low,"w "): return "W Hotels"
 case strings.Contains(low,"aloft"): return "Aloft"
 case strings.Contains(low,"moxy"): return "Moxy"
 case strings.Contains(low,"courtyard"): return "Courtyard by Marriott"
 default: return "Marriott" }}
func inferHiltonBrand(low string) string {
 switch {
 case strings.Contains(low,"waldorf"): return "Waldorf Astoria"
 case strings.Contains(low,"conrad"): return "Conrad"
 case strings.Contains(low,"doubletree"): return "DoubleTree by Hilton"
 case strings.Contains(low,"hampton"): return "Hampton by Hilton"
 default: return "Hilton" }}
func inferWyndhamBrand(low string) string {
 switch {
 case strings.Contains(low,"ramada"): return "Ramada"
 case strings.Contains(low,"tryp"): return "Tryp"
 default: return "Wyndham" }}
func inferBestWesternBrand(low string) string {
 switch {
 case strings.Contains(low,"premier"): return "Best Western Premier"
 case strings.Contains(low,"plus"): return "Best Western Plus"
 default: return "Best Western" }}
func inferRadissonBrand(low string) string {
 switch {
 case strings.Contains(low,"radisson blu"): return "Radisson Blu"
 case strings.Contains(low,"radisson red"): return "Radisson RED"
 case strings.Contains(low,"park inn"): return "Park Inn by Radisson"
 default: return "Radisson" }}
