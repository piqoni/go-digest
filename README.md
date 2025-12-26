# Intro

This repo shows how to use [Matcha](https://github.com/piqoni/matcha) and Github Pages/Actions to create your own news aggregator (either public or password-protected via [staticrypt](https://github.com/robinmoisson/staticrypt)).  

Demo URL (Go programming related feeds): https://piqoni.github.io/go-digest

<img width="718" height="807" alt="image" src="https://github.com/user-attachments/assets/c79be5d9-1df7-40f3-bb78-a164d50e2009" />

# How to make your own in less than 5 minutes

Substitute 'my-personal-digest' with the name you wish to call your digest, in below instructions:
 
```
git clone --depth 1 git@github.com:piqoni/go-digest my-personal-digest
cd my-personal-digest
rm matcha.db.enc docs/index.html
rm -rf .git
```

Go to https://github.com/new and create your my-personal-digest repository, and follow instructions to push the my-personal-digest directory as a new repository. (git init, commit, push)

Go to the my-personal-digest repository Settings -> Pages -> Branch: main -> /docs 

Go to the my-personal-digest repository Settings -> Secrets and Variables -> Actions
1. Mandatory: New Repository secret called **MATCHA_DB_KEY** which is used to encrypt the matcha database
2. Mandatory: Click Variables -> New Repository Variable -> **MATCHA_CONFIG_YAML** with value of a basic matcha configuration as below: 

Basic configuration, see [Matcha's repository](https://github.com/piqoni/matcha) for more options (like google_news_keywords etc)
```
markdown_dir_path: pre-docs
database_file_path: matcha.db
feeds:
  - http://hnrss.org/best
  - https://feeds.bbci.co.uk/news/rss.xml
```

3. Optional: if you want page to be **password protected**: Create a new Repository Secret called STATICRYPT_PASSWORD which will be the password that you will use to access your feed.

# Go Digest
Current Go Digest configuration/feeds subscriptions:
```yml
database_file_path: matcha.db
feeds:
  - https://www.alexedwards.net/static/feed.rss
  - https://www.reddit.com/r/golang/.rss
  - https://golangweekly.com/rss
  - https://blog.jetbrains.com/go/feed/
  - https://go.dev/blog/feed.atom
  - https://lobste.rs/t/go.rss
  - https://threedots.tech/index.xml
  - https://www.ardanlabs.com/index.xml
  - https://pliutau.com/index.xml
  - https://www.youtube.com/feeds/videos.xml?channel_id=UCO3LEtymiLrgvpb59cNsb8A
  - https://bitfieldconsulting.com/posts?format=rss
  - https://blog.boot.dev/golang/index.xml
  - https://antonz.org/tags/thank-go/index.xml
  - https://mshibanami.github.io/GitHubTrendingRSS/weekly/go.xml
google_news_keywords: golang
markdown_dir_path: pre-docs
```

