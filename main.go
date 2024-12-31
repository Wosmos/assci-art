package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

// StyleCategory represents a category of text styles
type StyleCategory struct {
	name        string
	description string
	styles      []Style
}

// Style represents an ASCII art style configuration
type Style struct {
	name        string
	description string
	font        string    // Maps to go-figure font name
	decorator   Decorator // Optional decorator for additional styling
}

// Decorator provides additional styling to the ASCII art
type Decorator struct {
	top       string
	bottom    string
	left      string
	right     string
	corners   [4]string // TL, TR, BL, BR
	fill      string
	pre       func(string) string // Pre-processing function
	post      func(string) string // Post-processing function
}

// ColorScheme represents a color configuration
type ColorScheme struct {
	name       string
	primary    *color.Color
	secondary  *color.Color
	background *color.Color
}

// AppConfig holds the application configuration
type AppConfig struct {
	categories []StyleCategory
	colors     []ColorScheme
}

// Constants for frame patterns
const (
	framePatternTop    = "═══╡ %s ╞═══"
	framePatternBottom = "═════════════"
)

// Predefined decorators
var (
	boxDecorator = Decorator{
		top:     "─",
		bottom:  "─",
		left:    "│",
		right:   "│",
		corners: [4]string{"┌", "┐", "└", "┘"},
	}

	doubleBoxDecorator = Decorator{
		top:     "═",
		bottom:  "═",
		left:    "║",
		right:   "║",
		corners: [4]string{"╔", "╗", "╚", "╝"},
	}

	roundBoxDecorator = Decorator{
		top:     "─",
		bottom:  "─",
		left:    "│",
		right:   "│",
		corners: [4]string{"╭", "╮", "╰", "╯"},
	}

	dottedBoxDecorator = Decorator{
		top:     "┈",
		bottom:  "┈",
		left:    "┊",
		right:   "┊",
		corners: [4]string{"·", "·", "·", "·"},
	}

	stars3DDecorator = Decorator{
		top:     "★",
		bottom:  "★",
		left:    "★",
		right:   "★",
		corners: [4]string{"★", "★", "★", "★"},
		pre: func(s string) string {
			return addShadow(s)
		},
	}

	wavyDecorator = Decorator{
		top:     "～",
		bottom:  "～",
		left:    "※",
		right:   "※",
		corners: [4]string{"∿", "∿", "∿", "∿"},
	}
)

func newAppConfig() *AppConfig {
	return &AppConfig{
		categories: []StyleCategory{
			{
				name:        "Classic",
				description: "Traditional ASCII art styles",
				styles: []Style{
					{"Standard", "Classic ASCII art", "", Decorator{}},
					{"Big", "Large block letters", "big", Decorator{}},
					{"Slim", "Thin elegant letters", "slim", Decorator{}},
					{"Small", "Compact letters", "small", Decorator{}},
				},
			},
			{
				name:        "Boxed",
				description: "Styles with different types of borders",
				styles: []Style{
					{"Single Box", "Single-line border", "", boxDecorator},
					{"Double Box", "Double-line border", "", doubleBoxDecorator},
					{"Round Box", "Rounded corners", "", roundBoxDecorator},
					{"Dotted Box", "Dotted border style", "", dottedBoxDecorator},
				},
			},
			{
				name:        "3D Effects",
				description: "Three-dimensional looking styles",
				styles: []Style{
					{"Shadow", "Letters with shadow", "shadow", Decorator{}},
					{"Deep 3D", "Enhanced 3D effect", "standard", stars3DDecorator},
					{"Block 3D", "Solid 3D blocks", "block", Decorator{
						pre: func(s string) string { return addShadow(s) },
					}},
				},
			},
			{
				name:        "Decorative",
				description: "Fancy and ornamental styles",
				styles: []Style{
					{"Wavy", "Wavy border style", "", wavyDecorator},
					{"Stars", "Starred border", "", stars3DDecorator},
					{"Script", "Cursive style", "script", Decorator{}},
					{"Bubble", "Rounded bubble letters", "bubble", roundBoxDecorator},
				},
			},
		},
		colors: []ColorScheme{
			{
				"Ocean",
				color.New(color.FgBlue),
				color.New(color.FgCyan),
				color.New(color.FgHiBlue),
			},
			{
				"Forest",
				color.New(color.FgGreen),
				color.New(color.FgHiGreen),
				color.New(color.FgWhite),
			},
			{
				"Sunset",
				color.New(color.FgRed),
				color.New(color.FgYellow),
				color.New(color.FgHiRed),
			},
			{
				"Royal",
				color.New(color.FgMagenta),
				color.New(color.FgHiMagenta),
				color.New(color.FgWhite),
			},
			{
				"Monochrome",
				color.New(color.FgWhite),
				color.New(color.FgHiWhite),
				color.New(color.FgBlack),
			},
			{
				"Neon",
				color.New(color.FgHiGreen),
				color.New(color.FgHiYellow),
				color.New(color.FgHiCyan),
			},
			{
				"Rainbow",
				color.New(color.FgRed),
				color.New(color.FgGreen),
				color.New(color.FgBlue),
			},
		},
	}
}

func main() {
	config := newAppConfig()

	// Command line flags
	outputFile := flag.String("output", "", "Output file path (optional)")
	showColors := flag.Bool("color", true, "Enable colored output")
	listStyles := flag.Bool("list", false, "List all available styles")
	previewMode := flag.Bool("preview", false, "Preview all styles with sample text")
	categoryFlag := flag.Int("category", 0, "Style category number")
	styleFlag := flag.Int("style", 0, "Style number within category")
	colorFlag := flag.Int("colorscheme", 0, "Color scheme number")
	interactiveMode := flag.Bool("interactive", true, "Interactive mode")
	flag.Parse()

	printWelcomeBanner()

	if *listStyles {
		config.listAvailableStyles()
		return
	}

	if *previewMode {
		config.previewStyles()
		return
	}

	// Main program loop
	
for {
    if !*interactiveMode {
        text := strings.Join(flag.Args(), " ")
        if text == "" {
            fmt.Println("Error: No text provided in non-interactive mode")
            os.Exit(1)
        }
        processText(text, config, outputFile, showColors, categoryFlag, styleFlag, colorFlag)
        return
    }

    text := getUserInput()
    if strings.ToLower(strings.TrimSpace(text)) == "q" {
        fmt.Println("\nGoodbye! Thanks for using ASCII Art Generator! 😊✌️")
        return
    }

    processText(text, config, outputFile, showColors, categoryFlag, styleFlag, colorFlag)
}
}

func processText(text string, config *AppConfig, outputFile *string, showColors *bool, categoryFlag, styleFlag, colorFlag *int) {
    _, style := config.getStyleSelection(*categoryFlag, *styleFlag)
    colorScheme := config.getColorSelection(*colorFlag, *showColors)

    asciiArt := config.generateArt(text, style, colorScheme)

    if *outputFile != "" {
        if err := saveToFile(*outputFile, asciiArt); err != nil {
            fmt.Printf("Error saving to file: %v\n", err)
            os.Exit(1)
        }
        fmt.Printf("ASCII art saved to: %s\n", *outputFile)
    } else {
        fmt.Println("\nYour ASCII Art:")
        fmt.Println(asciiArt)
        
        // Add pause and prompt
        fmt.Print("\nPress Enter to continue or type 'q' to quit: ")
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        if strings.ToLower(strings.TrimSpace(input)) == "q" {
            fmt.Println("\nGoodbye! Thanks for using ASCII Art Generator!")
            os.Exit(0)
        }
    }
}

func printWelcomeBanner() {
	banner := figure.NewFigure("ASCII Art", "big", true).String()
	fmt.Println(color.CyanString(banner))
	fmt.Println(color.HiYellowString(fmt.Sprintf(framePatternTop, "Generator")))
	fmt.Println(color.HiYellowString(framePatternBottom))
	fmt.Println("\nType 'q' to quit or Ctrl+C to exit")
	fmt.Println()
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.GreenString("Enter your text: "))
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}
	return strings.TrimSpace(text)
}

func (config *AppConfig) generateArt(text string, style Style, colorScheme *ColorScheme) string {
	var asciiArt string
	if style.font != "" {
		asciiArt = figure.NewFigure(text, style.font, true).String()
	} else {
		asciiArt = text
	}

	if style.decorator.pre != nil {
		asciiArt = style.decorator.pre(asciiArt)
	}

	if style.decorator.top != "" || style.decorator.bottom != "" || style.decorator.left != "" || style.decorator.right != "" {
		asciiArt = applyDecorator(asciiArt, style.decorator)
	}

	if style.decorator.post != nil {
		asciiArt = style.decorator.post(asciiArt)
	}

	if colorScheme != nil {
		asciiArt = applyColorScheme(asciiArt, colorScheme)
	}

	return asciiArt
}

func applyDecorator(text string, d Decorator) string {
	lines := strings.Split(text, "\n")
	maxWidth := 0
	for _, line := range lines {
		width := utf8.RuneCountInString(line)
		if width > maxWidth {
			maxWidth = width
		}
	}

	result := d.corners[0] + strings.Repeat(d.top, maxWidth+2) + d.corners[1] + "\n"

	for _, line := range lines {
		padding := strings.Repeat(" ", maxWidth-utf8.RuneCountInString(line))
		result += d.left + " " + line + padding + " " + d.right + "\n"
	}

	result += d.corners[2] + strings.Repeat(d.bottom, maxWidth+2) + d.corners[3]

	return result
}

func addShadow(text string) string {
	lines := strings.Split(text, "\n")
	result := make([]string, len(lines))
	
	for i, line := range lines {
		if i < len(lines)-1 {
			shadowLine := strings.Map(func(r rune) rune {
				if r != ' ' {
					return '░'
				}
				return ' '
			}, line)
			result[i] = line + "\n" + strings.Repeat(" ", 2) + shadowLine
		} else {
			result[i] = line
		}
	}
	
	return strings.Join(result, "\n")
}

func applyColorScheme(text string, cs *ColorScheme) string {
	lines := strings.Split(text, "\n")
	var result []string
	
	for i, line := range lines {
		colorIndex := i % 3
		var colored string
		switch colorIndex {
		case 0:
			colored = cs.primary.Sprint(line)
		case 1:
			colored = cs.secondary.Sprint(line)
		case 2:
			colored = cs.background.Sprint(line)
		}
		result = append(result, colored)
	}
	
	return strings.Join(result, "\n")
}

func (config *AppConfig) listAvailableStyles() {
	fmt.Println(color.CyanString("\nAvailable Style Categories:"))
	for i, category := range config.categories {
		fmt.Printf("\n%d. %s - %s\n", i+1,
			color.BlueString(category.name),
			color.YellowString(category.description))
		for j, style := range category.styles {
			fmt.Printf("   %d.%d %s - %s\n", i+1, j+1,
				color.HiWhiteString(style.name),
				color.HiBlackString(style.description))
		}
	}

	fmt.Println(color.CyanString("\nAvailable Color Schemes:"))
	for i, scheme := range config.colors {
		fmt.Printf("%d. %s\n", i+1, scheme.primary.Sprint(scheme.name))
	}
}

func (config *AppConfig) previewStyles() {
	sampleText := "Hello!"
	fmt.Println(color.CyanString("\nStyle Previews:"))
	
	for _, category := range config.categories {
		fmt.Printf("\n%s - %s\n", 
			color.BlueString(category.name),
			color.YellowString(category.description))
		
		for _, style := range category.styles {
			fmt.Printf("\n%s (%s):\n",
				color.HiWhiteString(style.name),
				color.HiBlackString(style.description))
			fmt.Println(config.generateArt(sampleText, style, nil))
		}
	}
}

func (config *AppConfig) getStyleSelection(categoryFlag, styleFlag int) (StyleCategory, Style) {
	if categoryFlag > 0 && categoryFlag <= len(config.categories) {
		category := config.categories[categoryFlag-1]
		if styleFlag > 0 && styleFlag <= len(category.styles) {
			return category, category.styles[styleFlag-1]
		}
	}

	fmt.Println("\nAvailable style categories:")
	for i, category := range config.categories {
		fmt.Printf("\n%d. %s - %s\n", i+1,
			color.BlueString(category.name),
			color.YellowString(category.description))
		for j, style := range category.styles {
			fmt.Printf("   %d.%d %s - %s\n", i+1, j+1,
				color.CyanString(style.name),
				color.HiWhiteString(style.description))
		}
}

	var categoryChoice, styleChoice int
	for {
		fmt.Printf("\nSelect category (1-%d): ", len(config.categories))
		if _, err := fmt.Scanf("%d", &categoryChoice); err == nil && 
			categoryChoice >= 1 && categoryChoice <= len(config.categories) {
			break
		}
		fmt.Println(color.RedString("Invalid selection. Please try again."))
		bufio.NewReader(os.Stdin).ReadString('\n')
	}

	category := config.categories[categoryChoice-1]
	for {
		fmt.Printf("Select style (1-%d): ", len(category.styles))
		if _, err := fmt.Scanf("%d", &styleChoice); err == nil && 
			styleChoice >= 1 && styleChoice <= len(category.styles) {
			break
		}
		fmt.Println(color.RedString("Invalid selection. Please try again."))
		bufio.NewReader(os.Stdin).ReadString('\n')
	}

	return category, category.styles[styleChoice-1]
}

func (config *AppConfig) getColorSelection(colorFlag int, showColors bool) *ColorScheme {
	if !showColors {
		return nil
	}

	if colorFlag > 0 && colorFlag <= len(config.colors) {
		return &config.colors[colorFlag-1]
	}

	// Clear input buffer before color selection
	bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println("\nAvailable color schemes:")
	for i, scheme := range config.colors {
		fmt.Printf("%d. %s\n", i+1, scheme.primary.Sprint(scheme.name))
	}

	var choice int
	for {
		fmt.Printf("\nSelect color scheme (1-%d): ", len(config.colors))
		if _, err := fmt.Scanf("%d", &choice); err == nil && 
			choice >= 1 && choice <= len(config.colors) {
			break
		}
		fmt.Println(color.RedString("Invalid selection. Please try again."))
		bufio.NewReader(os.Stdin).ReadString('\n')
	}

	return &config.colors[choice-1]
}

func saveToFile(filepath string, content string) error {
	return os.WriteFile(filepath, []byte(content), 0644)
}