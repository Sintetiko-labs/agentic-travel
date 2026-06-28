package client

import ("fmt"; "strings"; "time")

func parseYMD(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" { return "", fmt.Errorf("empty date") }
	if len(s) == 10 && s[4] == '-' && s[7] == '-' { if _, err := time.Parse("2006-01-02", s); err != nil { return "", err }; return s, nil }
	for _, f := range []string{"2006-01-02", "02/01/2006", "01/02/2006", "20060102"} { if t, err := time.Parse(f, s); err == nil { return t.Format("2006-01-02"), nil } }
	return "", fmt.Errorf("invalid date %q (use YYYY-MM-DD)", s)
}
