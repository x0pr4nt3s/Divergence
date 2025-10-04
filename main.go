package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

var (
	defaultGeneralWords = []string{
		"lima",
		"santiago",
		"bogota",
		"enero",
		"febrero",
		"marzo",
		"abril",
		"mayo",
		"junio",
		"julio",
		"agosto",
		"setiembre",
		"octubre",
		"noviembre",
		"diciembre",
		"january",
		"february",
		"march",
		"april",
		"may",
		"june",
		"july",
		"august",
		"september",
		"october",
		"november",
		"december",
		"verano",
		"invierno",
		"otono",
		"primavera",
		"sistemas",
		"enterprise",
		"soporte",
	}
	defaultTopYears = []string{
		"2040",
		"2035",
		"2030",
		"2025",
		"2023",
		"2022",
		"2021",
		"2020",
		"2019",
		"2018",
		"2017",
		"2016",
		"2015",
		"2014",
		"2013",
		"2012",
		"2011",
		"2010",
		"2009",
		"2008",
		"2007",
		"2006",
		"2005",
		"2004",
		"2003",
		"2002",
		"2001",
		"2000",
	}
	digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	lettersLower = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
	lettersUpper = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	allLetters = append(append([]string{}, lettersLower...), lettersUpper...)
	defaultSpecialChars = []string{
		"!", "@", "#", "$", "%", "&", "*", "(", ")", "-", "_", "+", "=",
		"[", "]", "{", "}", "|", "~", "`", ":", ";", "<", ">", ",", ".", "?", "/",
	}
	substitution = map[rune][]string{
		'a': []string{"a", "A", "4", "@"},
		'e': []string{"e", "E", "3"},
		'i': []string{"i", "I", "1", "!", "L"},
		'l': []string{"l", "L", "1", "!", "I"},
		'o': []string{"o", "O", "0"},
		's': []string{"s", "S", "5", "$"},
	}
	banner = `
　　∩.　　　∩.　　　∩.　　　∩.　　　∩.　　　∩.　　　∩.　　　∩.　　　　
  ノ×乂   ノ×乂   ノ×乂 　ノ×乂 　ノ×乂　 ノ×乂　 ノ×乂　 ノ×乂 　　
  |0┓爻|　|０爻|　|┏┓爻|　|┃┃爻|　|┏┓爻|　|┏┓爻|　|┏┓爻|　|┏┓爻| 　　
  |○┃爻|　|爻爻|　|┃┃爻|　|┗╋爻|　|┣┫爻|　|┗┓爻|　|┗┫爻|　|┣┓爻| 　　
  |0┃圭|　| * #|　|┗┛圭|　|/┃圭|　|┗┛圭|　|┗┛圭|　|┗┛圭|　|┗┛圭| 　　
  |圭圭|　|圭圭|　|圭圭|　|圭圭|　|圭圭|　|圭圭|　|圭圭|　|圭圭| 　　
 /|卅卅|廾|卅卅|廾|卅卅|廾|卅卅|廾|卅卅|廾|卅卅|廾|卅卅|廾|卅卅|且¯/|
|￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣||
|　　　　　　　　　- Wellcome to the Divergence -　　　　　　　　　||
|＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿|/
`
)

type config struct {
	rulePath        string
	specialPath     string
	outDir          string
	generalPath     string
	yearsPath        string
	specialCharsPath string
	maxSpecial      int
	maxTop          int
	maxRuleProduct  int
	maxResult       int
	noBanner        bool
	allowDuplicates bool
}

type emitter func(string) bool
type collector struct {
	file         *os.File
	writer       *bufio.Writer
	seen         map[string]struct{}
	values       []string
	limit        int
	limitReached bool
	err          error
}

type dedupeWriter struct {
	file            *os.File
	writer          *bufio.Writer
	seen            map[string]struct{}
	count           int
	limit           int
	limitReached    bool
	err             error
	allowDuplicates bool
}

type dictionaryData struct {
	specials     []string
	top          []string
	years        []string
	specialChars []string
}

type productStatus int
const (
	statusCompleted productStatus = iota
	statusRuleLimit
	statusStopAll
)
func main() {
	cfg := config{}
	flag.StringVar(&cfg.rulePath, "rule", "", "Ruta al archivo de reglas (obligatoria).")
	flag.StringVar(&cfg.specialPath, "specialword", "", "Ruta al archivo con palabras especiales (obligatoria).")
	flag.StringVar(&cfg.outDir, "out", ".", "Directorio donde se almacenaran los diccionarios generados.")
	flag.StringVar(&cfg.generalPath, "general", "", "Archivo con palabras generales adicionales (opcional).")
	flag.StringVar(&cfg.yearsPath, "years", "", "Archivo con anios adicionales (opcional).")
	flag.StringVar(&cfg.specialCharsPath, "special-chars", "", "Archivo con caracteres especiales adicionales (opcional).")
	flag.IntVar(&cfg.maxSpecial, "max-special", 0, "Limite de variantes unicas para special_words.txt (0 = sin limite).")
	flag.IntVar(&cfg.maxTop, "max-top", 0, "Limite de variantes unicas para top_special_words.txt (0 = sin limite).")
	flag.IntVar(&cfg.maxRuleProduct, "max-rule-product", 250000, "Limite de combinaciones generadas por cada regla (0 = sin limite).")
	flag.IntVar(&cfg.maxResult, "max-result", 0, "Limite de filas en resultado_wordlist.txt (0 = sin limite).")
	flag.BoolVar(&cfg.noBanner, "no-banner", false, "Oculta el banner de bienvenida.")
	flag.BoolVar(&cfg.allowDuplicates, "allow-duplicates", false, "Permite duplicados en resultado_wordlist.txt para reducir uso de memoria.")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Uso: %s -rule reglas.txt -specialword palabras.txt [opciones]\n\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()
	if cfg.rulePath == "" || cfg.specialPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
func run(cfg config) (err error) {
	if !cfg.noBanner {
		fmt.Println(banner)
	}
	if err := os.MkdirAll(cfg.outDir, 0o755); err != nil {
		return fmt.Errorf("crear directorio de salida: %w", err)
	}
	generalWords := append([]string{}, defaultGeneralWords...)
	if cfg.generalPath != "" {
	extraGeneral, err := readLines(cfg.generalPath)
		if err != nil {
			return fmt.Errorf("leer archivo de palabras generales: %w", err)
		}
		generalWords = append(generalWords, extraGeneral...)
	}
	years := append([]string{}, defaultTopYears...)
	if cfg.yearsPath != "" {
	extraYears, err := readLines(cfg.yearsPath)
		if err != nil {
			return fmt.Errorf("leer archivo de anios: %w", err)
		}
		years = append(years, extraYears...)
	}
	specialChars := append([]string{}, defaultSpecialChars...)
	if cfg.specialCharsPath != "" {
	extraSpecials, err := readLines(cfg.specialCharsPath)
		if err != nil {
			return fmt.Errorf("leer archivo de caracteres especiales: %w", err)
		}
		specialChars = append(specialChars, extraSpecials...)
	}
	specialInputs, err := readLines(cfg.specialPath)
	if err != nil {
		return fmt.Errorf("leer archivo de palabras especiales: %w", err)
	}
	rules, err := readLines(cfg.rulePath)
	if err != nil {
		return fmt.Errorf("leer archivo de reglas: %w", err)
	}
	specialCollector, err := newCollector(filepath.Join(cfg.outDir, "special_words.txt"), cfg.maxSpecial)
	if err != nil {
		return fmt.Errorf("crear special_words.txt: %w", err)
	}
	defer func() {
		if cerr := specialCollector.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	for _, word := range generalWords {
		if !emitCamelVariants(word, specialCollector.Emit) {
			break
		}
	}
	for _, word := range specialInputs {
		if !emitCamelVariants(word, specialCollector.Emit) {
			break
		}
	}
	topCollector, err := newCollector(filepath.Join(cfg.outDir, "top_special_words.txt"), cfg.maxTop)
	if err != nil {
		return fmt.Errorf("crear top_special_words.txt: %w", err)
	}
	defer func() {
		if cerr := topCollector.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	for _, word := range generalWords {
		if !emitTopVariants(word, topCollector.Emit) {
			break
		}
	}
	for _, word := range specialInputs {
		if !emitTopVariants(word, topCollector.Emit) {
			break
		}
	}
	resultWriter, err := newDedupeWriter(filepath.Join(cfg.outDir, "resultado_wordlist.txt"), cfg.maxResult, cfg.allowDuplicates)
	if err != nil {
		return fmt.Errorf("crear resultado_wordlist.txt: %w", err)
	}
	defer func() {
		if cerr := resultWriter.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	data := dictionaryData{
	specials:     specialCollector.Values(),
	top:          topCollector.Values(),
	years:        years,
	specialChars: specialChars,
}
	var limitedRules []string
RuleLoop:
	for _, rule := range rules {
		if rule == "" {
			continue
		}
		sources, err := buildRuleSources(rule, data)
		if err != nil {
			return fmt.Errorf("regla %q: %w", rule, err)
		}
		status := cartesianProduct(sources, cfg.maxRuleProduct, resultWriter.Emit)
		switch status {
		case statusRuleLimit:
			limitedRules = append(limitedRules, rule)
		case statusStopAll:
			break RuleLoop
		}
	}
	if resultWriter.Err() != nil {
		return resultWriter.Err()
	}
	if specialCollector.Err() != nil {
		return specialCollector.Err()
	}
	if topCollector.Err() != nil {
		return topCollector.Err()
	}
	fmt.Printf("special_words.txt listo (%d entradas).\n", len(specialCollector.Values()))
	fmt.Printf("top_special_words.txt listo (%d entradas).\n", len(topCollector.Values()))
	fmt.Printf("resultado_wordlist.txt listo (%d entradas).\n", resultWriter.Count())
	if specialCollector.LimitReached() && cfg.maxSpecial > 0 {
		fmt.Printf("Aviso: se alcanzo el limite de %d variantes en special_words.txt.\n", cfg.maxSpecial)
	}
	if topCollector.LimitReached() && cfg.maxTop > 0 {
		fmt.Printf("Aviso: se alcanzo el limite de %d variantes en top_special_words.txt.\n", cfg.maxTop)
	}
	if len(limitedRules) > 0 && cfg.maxRuleProduct > 0 {
		fmt.Printf("Aviso: se alcanzo el limite de %d combinaciones en las reglas: %s\n", cfg.maxRuleProduct, strings.Join(limitedRules, ", "))
	}
	if resultWriter.LimitReached() && cfg.maxResult > 0 {
		fmt.Printf("Aviso: se alcanzo el limite de %d lineas en resultado_wordlist.txt.\n", cfg.maxResult)
	}
	if cfg.allowDuplicates {
		fmt.Println("Duplicados permitidos en resultado_wordlist.txt (mapa en memoria desactivado).")
	}
	fmt.Printf("\nArchivos escritos en: %s\n", cfg.outDir)
	return nil
}
func newCollector(path string, limit int) (*collector, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &collector{
		file:   file,
		writer: bufio.NewWriter(file),
		seen:   make(map[string]struct{}),
		limit:  limit,
	}, nil
}
func (c *collector) Emit(word string) bool {
	if c.err != nil || word == "" {
		return c.err == nil
	}
	if _, ok := c.seen[word]; ok {
		return true
	}
	if c.limit > 0 && len(c.values) >= c.limit {
		c.limitReached = true
		return false
	}
	c.seen[word] = struct{}{}
	c.values = append(c.values, word)
	if _, err := c.writer.WriteString(word); err != nil {
		c.err = err
		return false
	}
	if err := c.writer.WriteByte('\n'); err != nil {
		c.err = err
		return false
	}
	return true
}
func (c *collector) Close() error {
	if c.writer != nil {
		if err := c.writer.Flush(); err != nil && c.err == nil {
			c.err = err
		}
	}
	if c.file != nil {
		if err := c.file.Close(); err != nil && c.err == nil {
			c.err = err
		}
	}
	return c.err
}
func (c *collector) Err() error {
	return c.err
}
func (c *collector) Values() []string {
	return c.values
}
func (c *collector) LimitReached() bool {
	return c.limitReached
}
func newDedupeWriter(path string, limit int, allowDup bool) (*dedupeWriter, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	var seen map[string]struct{}
	if !allowDup {
		seen = make(map[string]struct{})
	}
	return &dedupeWriter{
		file:            file,
		writer:          bufio.NewWriter(file),
		seen:            seen,
		limit:           limit,
		allowDuplicates: allowDup,
	}, nil
}
func (d *dedupeWriter) Emit(word string) bool {
	if d.err != nil || word == "" {
		return d.err == nil
	}
	if d.seen != nil {
		if _, ok := d.seen[word]; ok {
			return true
		}
	}
	if d.limit > 0 && d.count >= d.limit {
		d.limitReached = true
		return false
	}
	if d.seen != nil {
		d.seen[word] = struct{}{}
	}
	if _, err := d.writer.WriteString(word); err != nil {
		d.err = err
		return false
	}
	if err := d.writer.WriteByte('\n'); err != nil {
		d.err = err
		return false
	}
	d.count++
	return true
}
func (d *dedupeWriter) Close() error {
	if d.writer != nil {
		if err := d.writer.Flush(); err != nil && d.err == nil {
			d.err = err
		}
	}
	if d.file != nil {
		if err := d.file.Close(); err != nil && d.err == nil {
			d.err = err
		}
	}
	return d.err
}
func (d *dedupeWriter) Err() error {
	return d.err
}
func (d *dedupeWriter) Count() int {
	return d.count
}
func (d *dedupeWriter) LimitReached() bool {
	return d.limitReached
}
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 1024), 1024*1024)
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
func capitalize(word string) string {
	if word == "" {
		return ""
	}
	runes := []rune(strings.ToLower(word))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
func emitCamelVariants(word string, emit emitter) bool {
	perWord := make(map[string]struct{})
	localEmit := func(candidate string) bool {
		if candidate == "" {
			return true
		}
		if _, ok := perWord[candidate]; ok {
			return true
		}
		perWord[candidate] = struct{}{}
		return emit(candidate)
	}
	if !emitFromWord(word, localEmit) {
		return false
	}
	capWord := capitalize(word)
	if !localEmit(capWord) {
		return false
	}
	if capWord != word {
		if !emitFromWord(capWord, localEmit) {
			return false
		}
	}
	return true
}
func emitFromWord(word string, emit emitter) bool {
	options := buildOptions(word)
	if len(options) == 0 {
		return true
	}
	buffer := make([]byte, 0, len(word)*2)
	var walk func(int) bool
	walk = func(idx int) bool {
		if idx == len(options) {
			candidate := string(buffer)
			if !emit(candidate) {
				return false
			}
			upper := strings.ToUpper(candidate)
			if !emit(upper) {
				return false
			}
			lower := strings.ToLower(candidate)
			if !emit(lower) {
				return false
			}
			return true
		}
		for _, opt := range options[idx] {
			prevLen := len(buffer)
			buffer = append(buffer, opt...)
			if !walk(idx + 1) {
				return false
			}
			buffer = buffer[:prevLen]
		}
		return true
	}
	return walk(0)
}
func buildOptions(word string) [][]string {
	options := make([][]string, 0, len(word))
	for _, r := range word {
		if subs, ok := substitution[unicode.ToLower(r)]; ok {
			options = append(options, subs)
		} else {
			options = append(options, []string{string(r)})
		}
	}
	return options
}
func emitTopVariants(word string, emit emitter) bool {
	perWord := make(map[string]struct{})
	candidates := []string{word, capitalize(word), strings.ToUpper(word)}
	for _, candidate := range candidates {
		if candidate == "" {
			continue
		}
		if _, ok := perWord[candidate]; ok {
			continue
		}
		perWord[candidate] = struct{}{}
		if !emit(candidate) {
			return false
		}
	}
	return true
}
func cartesianProduct(sets [][]string, limit int, emit emitter) productStatus {
	if len(sets) == 0 {
		return statusCompleted
	}
	count := 0
	status := statusCompleted
	buffer := make([]byte, 0, 32)
	var walk func(int) bool
	walk = func(idx int) bool {
		if idx == len(sets) {
			if limit > 0 && count >= limit {
				status = statusRuleLimit
				return false
			}
			candidate := string(buffer)
			if !emit(candidate) {
				status = statusStopAll
				return false
			}
			count++
			return true
		}
		for _, value := range sets[idx] {
			prevLen := len(buffer)
			buffer = append(buffer, value...)
			if !walk(idx + 1) {
				return false
			}
			buffer = buffer[:prevLen]
			if status != statusCompleted {
				return false
			}
		}
		return true
	}
	if !walk(0) {
		return status
	}
	return status
}
func buildRuleSources(rule string, data dictionaryData) ([][]string, error) {
	sets := make([][]string, 0, len(rule))
	for _, ch := range rule {
		switch ch {
		case 'w':
			sets = append(sets, data.specials)
		case 't':
			sets = append(sets, data.top)
		case 'y':
			sets = append(sets, data.years)
		case 'n':
			sets = append(sets, digits)
		case 'l':
			sets = append(sets, allLetters)
		case 'M':
			sets = append(sets, lettersUpper)
		case 'm':
			sets = append(sets, lettersLower)
		case 's':
			sets = append(sets, data.specialChars)
		default:
			return nil, fmt.Errorf("caracter desconocido %q", string(ch))
		}
	}
	return sets, nil
}
