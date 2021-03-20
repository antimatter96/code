func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}


scanner := bufio.NewScanner(os.Stdin)

for ;n > 0; n-- {
	scanner.Scan()
	temp := toInt(scanner.Bytes())
}
