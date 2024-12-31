# 🎨 ASCII Art Generator

A powerful and flexible **command-line tool** for generating stunning, stylized **ASCII art** with customizable borders, vibrant color schemes, and creative effects. Whether you're creating banners, adding a fun touch to your project, or just having fun, this tool has you covered!

---

## 🌟 Features

✨ **Multiple ASCII Art Styles**:

- Classic styles: Standard, Big, Slim, Small
- Boxed styles: Various border types like single-line, rounded, or dotted
- 3D effects: Shadows and depth for extra flair
- Decorative styles: Fancy borders, bubble letters, and more

🎨 **Rich Color Schemes**:

- Ocean (Blue & Cyan), Forest (Green shades), Sunset (Red & Yellow)
- Royal (Magenta), Neon, Monochrome, Rainbow, and more!

💡 **Additional Features**:

- Interactive and non-interactive modes
- File output support to save your art
- Preview mode to explore styles before choosing
- Continuous operation mode for creating multiple designs

---

## 🚀 Installation

### **Prerequisites**

- Go 1.16 or higher
- Required packages:
  ```bash
  go get github.com/common-nighthawk/go-figure
  go get github.com/fatih/color
  ```

### **Building from Source**

```bash
git clone [repository-url]
cd ascii-art-generator
go build -o ascii-art main.go
```

---

## 💻 Usage

### **Interactive Mode**

```bash
./ascii-art
```

### **Non-Interactive Mode**

```bash
./ascii-art -interactive=false "Your Text Here"
```

### **Command Line Options**

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

---

## 🎨 Style Categories

### **1. Classic**

- Standard ASCII art
- Big block letters
- Slim elegant letters
- Small compact letters

### **2. Boxed**

- Single-line borders
- Double-line borders
- Rounded corners
- Dotted borders

### **3. 3D Effects**

- Shadow effects
- Deep 3D
- Block 3D

### **4. Decorative**

- Wavy borders
- Star decorations
- Script style
- Bubble letters

---

## 🌈 Color Schemes

- **Ocean**: Blue & Cyan
- **Forest**: Green shades
- **Sunset**: Red & Yellow
- **Royal**: Magenta
- **Neon**: Bright neon vibes
- **Monochrome**: Sleek black-and-white
- **Rainbow**: A splash of every color

---

## 📝 Examples

```bash
# Generate ASCII art with default settings
./ascii-art
> Enter your text: Hello World

# Save to file
./ascii-art -output art.txt "Hello World"

# Use specific style and color
./ascii-art -category 2 -style 1 -colorscheme 3 "Hello World"

# List available styles
./ascii-art -list

# Alternative way to run the program

# Replace ./ascii-art with "go run main.go"

# Save to file
go run main.go -output art.txt "Hello World"

# Use specific style and color
go run main.go -category 2 -style 1 -colorscheme 3 "Hello World"

# List available styles
go run main.go -list
```

---

## 🐛 Known Issues

- **Color Selection**: Some terminals may require double input for color selection (Fixed 🥳).
- **Special Characters**: Certain characters may not render correctly in specific styles (Improving 🔨).
- **Find a bug?**: Let me know!
[report issue](https://github.com/Wosmos/assci-art/issues/new)

---

## 🔮 Future Improvements

- Additional border styles and decorations
- Custom color scheme creation
- Animation support for dynamic ASCII art
- Font combination options
- Template saving and loading for quick reuse

---

🌐 Let's connect 

[![Let's Connect](https://img.shields.io/badge/LEt's%20Connect-My%20Portfolio-purple?style=for-the-badge)](https://wasiff.vercel.app/)

Let’s turn your ideas into reality! 🌟

---

Thank you for checking out the **ASCII Art Generator**! If you enjoyed it, don’t forget to give this project a ⭐ on GitHub and share it with others!
