package parse
import ("regexp"; "strings")
var accorHotelCodeRE = regexp.MustCompile(`accorhotels/([a-z]{3}_[a-z]_[0-9]+)`)
func HotelsFromAccorSSR(html, baseURL string) []HotelLD {
 seen := map[string]bool{}; var out []HotelLD
 for _, m := range accorHotelCodeRE.FindAllStringSubmatch(html, -1) {
  code := m[1]; if seen[code] { continue }; seen[code]=true
  out = append(out, HotelLD{ID: code, Name: titleFromSlug(strings.ReplaceAll(code,"_"," ")), URL: absolutize(baseURL,"/hotels/"+code)})
 }
 return out
}
func HotelsFromAccorDestination(html, baseURL string) []HotelLD {
 if out := HotelsFromJSONLD(html, baseURL); len(out)>0 { return out }
 if out := HotelsFromAccorSSR(html, baseURL); len(out)>0 { return out }
 return HotelsFromBrandHomeLinks(html, baseURL, "all.accor.com")
}
