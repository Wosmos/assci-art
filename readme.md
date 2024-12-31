# ASCII Art Generator

A powerful and flexible command-line tool for generating stylized ASCII art with various decorative borders and color schemes.

## 🌟 Features

- Multiple ASCII art styles across different categories:
  - Classic styles (Standard, Big, Slim, Small)
  - Boxed styles with various border types
  - 3D effects with shadows and depth
  - Decorative styles with fancy borders
- Rich color schemes including Ocean, Forest, Sunset, Royal, and more
- Interactive and non-interactive modes
- File output support
- Border decorations and special effects
- Preview mode for all available styles
- Continuous operation mode

## 🚀 Installation

### Prerequisites
- Go 1.16 or higher
- Required packages:
  ```bash
  go get github.com/common-nighthawk/go-figure
  go get github.com/fatih/color
  ```

### Building from Source
```bash
git clone [repository-url]
cd ascii-art-generator
go build -o ascii-art main.go
```

## 💻 Usage

### Interactive Mode
```bash
./ascii-art
```

### Non-interactive Mode
```bash
./ascii-art -interactive=false "Your Text Here"
```

### Command Line Options
```
-output string    Output file path (optional)
-color bool       Enable colored output (default: true)
-list            List all available styles
-preview         Preview all styles with sample text
-category int    Style category number
-style int       Style number within category
-colorscheme int Color scheme number
-interactive     Interactive mode (default: true)
```

## 🎨 Style Categories

1. **Classic**
   - Standard ASCII art
   - Big block letters
   - Slim elegant letters
   - Small compact letters

2. **Boxed**
   - Single-line borders
   - Double-line borders
   - Rounded corners
   - Dotted borders

3. **3D Effects**
   - Shadow effects
   - Deep 3D
   - Block 3D

4. **Decorative**
   - Wavy borders
   - Star decorations
   - Script style
   - Bubble letters

## 🌈 Color Schemes

- Ocean (Blue & Cyan)
- Forest (Green shades)
- Sunset (Red & Yellow)
- Royal (Magenta)
- Monochrome
- Neon
- Rainbow

## 🚪 Exit Options
- Type 'q' and press Enter to quit
- Press Ctrl+C to exit

## 📝 Examples

```bash
# Generate ASCII art with default settings
./ascii-art
> Enter your text: Hello World

# Save to file
./ascii-art -output art.txt "Hello World"

# Use specific style and color
./ascii-art -category 2 -style 1 -colorscheme 3 "Hello World"

# Preview all styles
./ascii-art -preview

# List available styles
./ascii-art -list
```

## 🐛 Known Issues

- Color selection may require double input in some terminals (Fixed 🥳)
- Some special characters may not render correctly in certain styles (Working on it 🔨)

## 🔮 Future Improvements

- Additional border styles and decorations
- Custom color scheme creation
- Animation support
- Font combination options
- Template saving and loading