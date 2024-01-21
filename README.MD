# What 's it

This is a simple project to let you keep on doing something via [Bark](https://github.com/Finb/Bark)

# How To Use

## 🐳 Docker

```shell
docker build -t holdon .
docker run -e BARK_KEY=your_bark_key -e INTERVAL=60 -itd holdon
```

> notify every 60 minutes by default