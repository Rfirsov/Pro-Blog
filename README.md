# Pro-Blog

> A modern, full-stack blogging platform built with **Nuxt 4**, **shadcn/ui**, and **Go (back-end)**. Designed for developer blogs, technical content, and community-driven publishing.

---

## ✅ Core Features

### 🧑‍💻 User System
- [x] Register, Login, Logout
- [x] JWT-based authentication (access & refresh)
- [x] Roles: User, Author, Admin
- [x] Secure password hashing (bcrypt)

### 📝 Blog Posts
- [x] Create, Edit, Delete Posts (Markdown)
- [x] Post slugs (`/blog/my-first-post`)
- [x] Draft vs Published state
- [] Cover image upload (with preview)
- [x] Tag support (filterable)

### 🧾 Content Display
- [] Public blog list page with pagination
- [] Post detail view
- [] Tag-based filtering
- [] SEO-friendly titles & meta descriptions
- [] RSS feed for all posts

---

## 💬 Community & Engagement
- [] Comments under each post (authenticated)
- [] Like posts (toggle)
- [] Bookmark posts (user saved list)
- [ ] Reply to comments (threaded replies)

---

## 🧑‍💼 Admin / CMS
- [] Admin dashboard
- [] Manage users (block, promote)
- [] Moderate posts (publish/delete)
- [ ] Manage comments (remove spam, etc.)
- [ ] Analytics overview (views, engagement)

---

## 🖥 Developer Tools
- [x] RESTful API with Swagger docs
- [x] Modular Go services with clean architecture
- [x] Docker + Docker Compose for local dev
- [] GitHub Actions (CI for testing + linting)
- [x] .env config management
- [ ] Unit & integration tests for critical paths

---

## 🌍 Deployment
- [] Nuxt + shadcn/ui frontend deployed on Vercel
- [] Go API + PostgreSQL hosted on Railway or Fly.io
- [] Auto deploy from `main` branch
- [] CDN caching + image optimization

---

## 🔮 Future Enhancements

| Feature                                    | Status     | Priority |
|--------------------------------------------|------------|----------|
| Email notifications (new post, replies)    | 🟡 Planned | Medium   |
| User avatars via Gravatar/upload           | 🟡 Planned | Low      |
| WYSIWYG/Markdown toggle editor             | 🟢 In Dev  | High     |
| Social sharing (X, LinkedIn, etc.)         | 🟡 Planned | Medium   |
| Dark mode theme switcher                   | 🟢 In Dev  | High     |

---

## 🧠 Tech Stack

| Layer        | Stack                              |
|--------------|------------------------------------|
| Front-End    | Nuxt 4, Tailwind CSS, shadcn/ui    |
| Back-End     | Go (Gin), PostgreSQL, JWT          |
| Deployment   | Vercel (Nuxt), Fly.io (API)        |
| CI/CD        | GitHub Actions, Docker             |
| Tooling      | Swagger, Postman, ESLint, Prettier |



## 📦 Swagger Docs Generation

To generate Swagger (OpenAPI 2.0) documentation for the ProBlog API, run:

```bash
swag init -g cmd/pro-blog-api/main.go

### 🔗 How to view Swagger UI

After starting the server, open [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) in your browser.
