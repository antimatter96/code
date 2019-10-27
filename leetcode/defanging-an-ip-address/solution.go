func defangIPaddr(address string) string {
    re := regexp.MustCompile(`\.`)
  return re.ReplaceAllLiteralString(address, "[.]")
}

func defangIPaddr(address string) string {    
  return strings.Replace(address, ".","[.]", -1)
}

