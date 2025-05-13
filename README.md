


# 🌐 yxanul: The Ultimate Web Scanning Toolkit

**yxanul** is a powerful, all-in-one web scanning toolkit designed to streamline reconnaissance and vulnerability assessment. Built with Go, it integrates cutting-edge tools for WAF detection, crawling, and scanning, making it a must-have for penetration testers, bug bounty hunters, and security enthusiasts.

---

## 🚀 Features

- **🔍 WAF Detection**: Automatically detect Web Application Firewalls (WAFs) using `wafw00f`.
- **🕸️ Intelligent Crawling**: Traverse JavaScript-heavy websites and extract actionable URLs with `Katana`.
- **🌐 Live Host Probing**: Validate and filter live URLs using `HTTPX`.
- **🛡️ Vulnerability Scanning**: Unleash the power of `Nuclei` with template-based vulnerability detection.
- **⚙️ Flexible Modes**: Choose between aggressive and standard scanning modes based on your needs.

---

## 📦 Installation

### Prerequisites

Ensure the following tools are installed on your system:
- [Go](https://go.dev/doc/install)
- [`wafw00f`](https://github.com/EnableSecurity/wafw00f)
- [`Katana`](https://github.com/projectdiscovery/katana)
- [`HTTPX`](https://github.com/projectdiscovery/httpx)
- [`Nuclei`](https://github.com/projectdiscovery/nuclei)

### Clone and Build

1. Clone the repository:
   ```bash
   git clone https://github.com/yxanul/yxanul.git
   cd yxanul
   ```

2. Build the binary:
   ```bash
   go build -o yxanul
   ```

3. Verify the binary is ready:
   ```bash
   ./yxanul -h
   ```

---

## 💻 Usage

Run **yxanul** with the following options:

### Basic Usage
```bash
./yxanul -url <target-url>
```

### Aggressive Mode
For faster scans:
```bash
./yxanul -url <target-url> -a
```

---

## 🛠️ How It Works

 **Step 1**: Detects WAF presence on the target website using `wafw00f`.
 
 **Step 2**: Crawls the site using `Katana`, extracting JavaScript-heavy links, sitemaps, and robots.txt.
 
 **Step 3**: Probes discovered URLs using `HTTPX` to identify live hosts.
 
 **Step 4**: Scans live URLs for vulnerabilities with `Nuclei`, leveraging its powerful template engine.

---

## ⚙️ Configuration

Customize the behavior of **yxanul** by modifying the integrated tool configurations:
- Adjust crawling depth, concurrency, or rate-limiting in `Katana`.
- Tune `Nuclei` to load custom templates with `-nt` or `-as` for specific scans.

---

## 📖 Example

Scanning the OWASP Juice Shop:
```bash
./yxanul -url https://juice-shop.herokuapp.com -a
```

Sample output:
```
WAF Detected: No WAF detected
Running Katana for crawling...
Running HTTPX to probe URLs...
Running Nuclei on probed URLs...
Scan completed successfully!
```

---

## 🧩 Contributing

Contributions are welcome! If you find a bug or have an idea for improvement:
1. Fork the repository.
2. Create a new branch (`feature/my-feature`).
3. Commit your changes.
4. Push the branch and create a pull request.

---

## 🛡️ License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

---

## ✨ Acknowledgments

- Thanks to [Project Discovery](https://projectdiscovery.io/) for their incredible tools.
- Built with ❤️ by [yxanul](https://github.com/yxanul).

---

## 🌟 Support

If you enjoy using **yxanul**, feel free to ⭐ the repo and share it with the community!

