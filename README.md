# 🕸️ DOMXSS-CTF Challenges

A set of educational and competitive challenges on DOM-based XSS, designed to practice finding, exploiting, and automating web vulnerabilities. Each challenge is a standalone HTML file hosted on GitHub Pages, and solutions are verified using a custom Go checker.



## 🎯 Repository Goal

- Provide realistic, yet controlled DOM XSS scenarios of easy to moderate difficulty.
- Demonstrate data exfiltration techniques via client-side vectors (e.g., fetch).
- Provide automated verification of solutions using a Go checker.



## 🚀 Deployment

All challenges are pure HTML/JS/CSS and do not require a backend. This allows them to be instantly deployed via **[GitHub Pages](https://pages.github.com/)**:

1. Commit the HTML file to `challenges/`.
2. Enable GitHub Pages in your repository settings (usually from the `main` branch or the `/docs` folder).
3. The challenge becomes available at:
   `https://<your-login>.github.io/<repo>/challenges/public-echo.html`

No servers, SSL certificates, or configurations—everything out of the box.



## 🤖 Go Checker

Automatically checking solutions is performed using a simple but reliable **HTTP server in Go**, which:

- Accepts GET/POST requests with a flag (for example, via `fetch` from an XSS payload).
- Validates the flag format (`CTF{...}`).
- Can be integrated with a Telegram bot, Discord webhook, etc.



![Thank you for your attention to this repository](https://media1.tenor.com/m/yf2J9gTT3rQAAAAC/bye-bye.gif)

