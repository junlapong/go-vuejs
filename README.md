# Go Vuejs Example

Sample SPA application with Vuejs for web front-end and Golang for back-end

## Prerequisite

```
brew install go
brew install node
npm i -g yarn
npm i -g @vue/cli
```

## Run by clone this repository

```
git clone https://github.com/junlapong/go-vuejs.git
cd web && yarn && yarn build && cd -
go mod tidy
go run main.go
```

## Create new project

step by step for create new project

```
mkdir go-vuejs
cd go-vuejs
git init
```

### Create Vue Application

```
vue create web --no-git
```

#### Setup Options

```
Vue CLI v4.4.5
? Please pick a preset:
  default (babel, eslint)
❯ Manually select features

? Please pick a preset: Manually select features
? Check the features needed for your project:
 ◉ Babel
 ◯ TypeScript
 ◯ Progressive Web App (PWA) Support
 ◉ Router
 ◉ Vuex
 ◯ CSS Pre-processors
 ◉ Linter / Formatter
 ◉ Unit Testing
❯◉ E2E Testing

? Please pick a preset: Manually select features
? Check the features needed for your project: Babel, Router, Vuex, Linter, Unit, E2E
? Use history mode for router? (Requires proper server setup for index fallback in production) (Y/n) y

? Pick a linter / formatter config:
  ESLint with error prevention only
  ESLint + Airbnb config
  ESLint + Standard config
❯ ESLint + Prettier

? Pick additional lint features: (Press <space> to select, <a> to toggle all, <i> to invert selection)
❯◉ Lint on save
 ◯ Lint and fix on commit

? Pick a unit testing solution:
  Mocha + Chai
❯ Jest

? Pick an E2E testing solution: (Use arrow keys)
❯ Cypress (Chrome only)
  Nightwatch (WebDriver-based)

? Where do you prefer placing config for Babel, ESLint, etc.?
  In dedicated config files
❯ In package.json
```

create or update code as example in [App.vue](web/src/App.vue), [TestRoute.vue](web/src/components/TestRoute.vue), [index.js](web/src/router/index.js)

```
web
└── src
    ├── App.vue
    ├── components
    │   └── TestRoute.vue
    └── router
        └── index.js
```

#### Run Development Server

```
cd web && yarn serve && cd -
```

Then open http://localhost:8080/

#### Build Web for Production

```
cd web && yarn build && cd -
```

### Develop Go Application

```
go mod init github.com/junlapong/go-vuejs
touch main.go
```

Then write code in [main.go](main.go)

```
go run main.go
```

Then open http://localhost:8088/
