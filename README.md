# OreSec Website
The OreSec website was in need of renewal/refurbishing. This was created without mozzarella and instead uses templated HTML with CSS for styling and a Go backend for an easy and secure http/https serving format.

## How to develop locally:

### Clone the repo

### install Go and the neccesary packages (formalize this process soonTM)

### install npm

Run this command to install the neccesary packages for tailwind and development

```bash
npm install
```
Run these to begin the server

```bash
go build oresec-site

go run oresec-site
```

To regenerate the CSS files when needed (i.e. Adding a tailwind css package that doesn't exist)

Run this
```bash
go generate
```

### Useful Resources
https://xeiaso.net/blog/using-tailwind-go/


