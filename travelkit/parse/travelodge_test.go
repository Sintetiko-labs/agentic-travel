package parse

import "testing"

func TestHotelsFromTravelodgeSitemapLondon(t *testing.T) {
	xml := `<?xml version="1.0"?>
<urlset>
  <url><loc>https://www.travelodge.co.uk/uk/london/aldgate-east/index.html</loc></url>
  <url><loc>https://www.travelodge.co.uk/uk/greater-london/chiswick/index.html</loc></url>
  <url><loc>https://www.travelodge.co.uk/uk/manchester/piccadilly/index.html</loc></url>
</urlset>`
	rows := HotelsFromTravelodgeSitemap(xml, "London")
	if len(rows) != 2 {
		t.Fatalf("got %d rows, want 2", len(rows))
	}
	if rows[0].ID != "london/aldgate-east" {
		t.Fatalf("id=%q", rows[0].ID)
	}
	if rows[0].Name != "Travelodge Aldgate East" {
		t.Fatalf("name=%q", rows[0].Name)
	}
}
